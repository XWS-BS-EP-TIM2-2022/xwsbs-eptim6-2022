package handlers

import (
	"context"
	"encoding/base32"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/logger"
	"github.com/sirupsen/logrus"
	"math/rand"
	"net/http"
	"net/smtp"
	"reflect"
	"rsc.io/qr"
	"runtime"
	"strings"
	"time"
	"unicode"

	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/auth_service/startup/config"
	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/auth_service/store"
	authServicePb "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/proto/auth_service"
	profileGw "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/proto/profile_service"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgryski/dgoogauth"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ErrorMessage struct {
	Message string `json:"message"`
}
type JWTOptions struct {
	Username          string
	IsTokenNonExpired bool
}
type JWT struct {
	Token string `json:"token"`
}
type AuthHandler struct {
	UserStore                *store.UsersStore
	secretKey                []byte
	profileServiceGrpcClient profileGw.ProfileServiceClient
	log                      *logger.LoggerWrapper
	config                   config.Config
}

type ActivationRequest struct {
	token string `json:"token"`
}

func getConnection(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
func InitAuthHandler(serverConfig *config.Config, wrapper *logger.LoggerWrapper) (*AuthHandler, error) {
	userStore := store.InitUsersStore(*serverConfig)
	endpoint := fmt.Sprintf("%s:%s", serverConfig.ProfileServiceGrpcHost, serverConfig.ProfileServiceGrpcPort)
	conn, err := getConnection(endpoint)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Fatal error init profile service connection! %s", err.Error()))
	}
	client := profileGw.NewProfileServiceClient(conn)
	return &AuthHandler{UserStore: userStore, profileServiceGrpcClient: client,
		secretKey: []byte(serverConfig.SecretKey), log: wrapper, config: *serverConfig}, nil
}

func (ah *AuthHandler) LoginUser(user store.User) (JWT, error) {
	dbUser, err := ah.UserStore.FindByUsername(user.Username)
	if err != nil {
		return JWT{Token: ""}, errors.New(fmt.Sprintf("User with specified username not found. Username: %s", user.Username))
	}
	if dbUser.IsActivated == false {
		return JWT{Token: ""}, errors.New(fmt.Sprintf("User account not activated. Username: %s", user.Username))
	}

	if err = CheckPasswordHash(user.Password, dbUser.Password); err != nil {
		err = ah.HandleFailedLogin(dbUser)
		ah.log.Writeln(logger.LogMessage{Message: fmt.Sprintf("Login failed. %s", err.Error()), Component: getComponentName(CheckPasswordHash), Level: logrus.WarnLevel})
		return JWT{Token: ""}, err
	}

	if dbUser.Blocked {
		if time.Now().After(user.BlockedUntil) {
			ah.UserStore.ResetFailedLogForUser(user.Username)
		} else {
			return JWT{Token: ""}, errors.New(fmt.Sprintf("User account blocked. Username: %s", user.Username))
		}
	}
	ah.UserStore.ResetFailedLogForUser(user.Username)

	tokenStr, err := ah.GenerateJWT(JWTOptions{IsTokenNonExpired: false, Username: dbUser.Username})
	if err != nil {
		return JWT{Token: ""}, errors.New(fmt.Sprintf("Generation of JWT faild. %s", err.Error()))
	}
	return JWT{Token: tokenStr}, nil
}

func (ag *AuthHandler) AuthorizeJWT(w http.ResponseWriter, r *http.Request) {
	if r.Header["Authorization"] != nil {
		tokenStr := strings.Split(r.Header["Authorization"][0], " ")[1]
		ag.ValidateToken(tokenStr)
	}
}

func (ag *AuthHandler) ValidateToken(tokenStr string) (*store.User, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error")
		}
		return ag.secretKey, nil
	})
	if err != nil {
		ag.log.Writeln(logger.LogMessage{Message: fmt.Sprintf("Specified token is not valid."), Level: logrus.WarnLevel, Component: getComponentName(jwt.Parse)})
		return nil, err
	}
	if !token.Valid {
		ag.log.Writeln(logger.LogMessage{Message: fmt.Sprintf("Specified token is not valid."), Level: logrus.WarnLevel, Component: getComponentName(ag.ValidateToken)})
		return nil, errors.New("JWT not valid")
	}
	username := token.Claims.(jwt.MapClaims)["username"]
	return &store.User{Username: fmt.Sprintf("%v", username)}, nil

}
func (ag *AuthHandler) AddNewUser(user *store.User) error {
	if _, err := ag.UserStore.FindByUsername(user.Username); err == nil {
		return &store.RequestError{Err: errors.New("Username already in use"), StatusCode: 400}
	}
	if !isPasswordValid(user.Password) {
		return &store.RequestError{Err: errors.New("Bad password format"), StatusCode: 400}
	}
	user.Password, _ = HashPassword(user.Password)
	user.FailedLogins = 0

	user.IsActivated = false
	verificationHash, urlToken := ag.GenerateVerificationToken()
	user.VerificationToken = verificationHash
	user.TokenExpiration = time.Now().Add(time.Hour * 2)
	err := ag.SendEmail(user.Email, ag.MailActivationMessage(urlToken))
	if err != nil {
		return err
	}

	err = ag.UserStore.AddNew(user)
	if err != nil {
		return err
	}
	return nil
}

func (ag *AuthHandler) NotifyProfileServiceAboutRegistration(ctx context.Context, in *authServicePb.User, dbUser *store.User) error {
	_, err := ag.profileServiceGrpcClient.AddNewUser(ctx, &profileGw.UserRequest{User: &profileGw.User{
		Username:  in.Username,
		Password:  dbUser.Password,
		Biography: in.Biography,
		BirthDate: in.BirthDate,
		Email:     in.Email,
		Gender:    in.Gender,
		IsPublic:  in.IsPublic,
		Name:      in.Name,
		Telephone: in.Telephone,
		Surname:   in.Surname,
		Role:      dbUser.Role,
	}})
	return err
}

func (ag *AuthHandler) GetAllUsers() []store.User {
	return ag.UserStore.FindAll()

}

func DecodeUser(req *http.Request) (store.User, error) {
	var user store.User
	err := json.NewDecoder(req.Body).Decode(&user)
	return user, err
}

func (ah *AuthHandler) GenerateJWT(options JWTOptions) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["username"] = options.Username
	if !options.IsTokenNonExpired {
		claims["exp"] = time.Now().Add(time.Minute * 80).Unix()
	}
	tokenStr, err := token.SignedString(ah.secretKey)
	if err != nil {
		fmt.Errorf("token signing error")
		return "", err
	}
	return tokenStr, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}

func isPasswordValid(password string) bool {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(password) >= 8 {
		hasMinLen = true
	}
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}

func (ah *AuthHandler) HandleFailedLogin(user store.User) error {
	if user.Blocked {
		if time.Now().After(user.BlockedUntil) {
			ah.UserStore.ResetFailedLogForUser(user.Username)
		} else {
			return errors.New(fmt.Sprintf("User account blocked. Username: %s", user.Username))
		}
	} else {
		if user.FailedLogins == 5 {
			ah.UserStore.BlockUser(user.Username)
			return errors.New(fmt.Sprintf("User account blocked. Username: %s", user.Username))
		}
	}
	_ = ah.UserStore.UpdateFailedLogForUser(user.Username)
	return errors.New(fmt.Sprintf("Invalid user password. Username: %s", user.Username))
}

func (ah *AuthHandler) GenerateVerificationToken() (string, string) {
	randomBytes := make([]byte, 10)
	rand.Seed(time.Now().UnixNano())
	rand.Read(randomBytes)
	encoded := base64.RawURLEncoding.EncodeToString(randomBytes)

	token, _ := bcrypt.GenerateFromPassword([]byte(encoded), 14)
	return string(token), encoded
}

func (ah *AuthHandler) ChangePassword(request store.ChangePasswordRequest) (*store.User, error) {
	user, err := ah.UserStore.FindByUsername(request.Username)
	if err != nil {
		return &user, err
	}
	er := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.OldPassword))
	if er != nil {
		return &user, errors.New(fmt.Sprintf("Specified old password is not corect."))
	} else {
		if !isPasswordValid(request.NewPassword) {
			return &user, errors.New(fmt.Sprintf("Incorrect format of new password."))
		}
		newPassword, _ := HashPassword(request.NewPassword)
		err := ah.UserStore.UpdatePassword(request.Username, newPassword)
		if err != nil {
			return &user, err
		}
	}
	return &user, nil
}

func (ah *AuthHandler) SendEmail(emailTo string, mailMessage []byte) error {
	from := config.NewConfig().EmailFrom
	password := config.NewConfig().EmailPassword
	to := []string{emailTo}
	host := config.NewConfig().EmailHost
	port := config.NewConfig().EmailPort
	address := host + ":" + port

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(address, auth, from, to, mailMessage)
	if err != nil {
		return err
	}
	return nil
}

func (ah *AuthHandler) MailActivationMessage(token string) []byte {
	confirmationLink := "https://" + config.NewConfig().FrontendUri + "/account-activation/" + token

	subject := "Subject: Account activation\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := ah.GenerateMailBody(confirmationLink, "Welcome to Dislinkt!", "Press the button below to activate your account", "Activate Account")
	message := []byte(subject + mime + body)
	return message
}

func (ah *AuthHandler) ActivateAccount(token string) (string, error) {
	user := ah.UserStore.FindByToken(token)
	if user == (store.User{}) {
		return "", errors.New(fmt.Sprintf("User account activation failed. Invalid token! Username: %s", user.Username))
	}

	if user.TokenExpiration.Before(time.Now()) {
		return "", errors.New(fmt.Sprintf("User account activation failed. Token expired! Username: %s", user.Username))
	}

	validateToken := bcrypt.CompareHashAndPassword([]byte(user.VerificationToken), []byte(token))
	if validateToken != nil {
		return "", errors.New(fmt.Sprintf("User account activation failed. Invalid token! Username: %s", user.Username))
	}

	err := ah.UserStore.ActivateAccount(user.Username)
	return user.Username, err
}

func (ah *AuthHandler) AccountRecoveryEmail(email string) (store.User, error) {
	user, noUserFound := ah.UserStore.FindByEmail(email)

	if noUserFound != nil {
		return user, errors.New(fmt.Sprintf("Account with this email doesn't exist! Email: %s", email))
	}

	verificationHash, urlToken := ah.GenerateVerificationToken()

	err := ah.UserStore.RefreshToken(user.Username, verificationHash)
	if err != nil {
		return user, err
	}

	resetLink := "https://" + config.NewConfig().FrontendUri + "/set-password/" + urlToken
	subject := "Subject: Account recovery\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := ah.GenerateMailBody(resetLink, "Account Recovery", "Press the button below to reset password", "Reset Password")
	message := []byte(subject + mime + body)
	err = ah.SendEmail(user.Email, message)
	if err != nil {
		return store.User{}, err
	}
	return user, nil
}

func (ah *AuthHandler) Enable2FA(username string) ([]byte, error) {
	randomStr := randStr(6)
	secret := base32.StdEncoding.EncodeToString([]byte(randomStr))
	authLink := "otpauth://totp/SocketLoop?secret=" + secret + "&issuer=SocketLoop"
	code, err := qr.Encode(authLink, qr.L)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	imgByte := code.PNG()

	err = ah.UserStore.Enable2FA(username, secret)
	if err != nil {
		return nil, err
	}
	return imgByte, nil
}

func (ah *AuthHandler) Validate2FA(username string, token string) bool {
	user, err := ah.UserStore.FindByUsername(username)
	if err != nil {
		return false
	}
	otpConfig := &dgoogauth.OTPConfig{
		Secret:      strings.TrimSpace(user.Secret),
		WindowSize:  3,
		HotpCounter: 0,
	}
	trimmedToken := strings.TrimSpace(token)

	ok, err := otpConfig.Authenticate(trimmedToken)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return ok
}

func (ah *AuthHandler) ResetPassword(token string, newPassword string) (string, error) {
	user := ah.UserStore.FindByToken(token)
	if user == (store.User{}) {
		return user.Username, errors.New(fmt.Sprintf("User account recovery failed. Invalid token!"))
	}

	if !isPasswordValid(newPassword) {
		return user.Username, errors.New(fmt.Sprintf("User account recovery failed. Incorrect password format"))
	}

	password, _ := HashPassword(newPassword)
	err := ah.UserStore.UpdatePassword(user.Username, password)
	return user.Username, err
}

func (ah *AuthHandler) SendPasswordlessLoginEmail(email string) (string, error) {
	user, noUserFound := ah.UserStore.FindByEmail(email)

	if noUserFound != nil {
		return "", errors.New(fmt.Sprintf("Account with this email doesn't exist! Email: %s", email))
	}

	verificationHash, urlToken := ah.GenerateVerificationToken()
	err := ah.UserStore.RefreshToken(user.Username, verificationHash)
	if err != nil {
		return user.Username, err
	}

	resetLink := "https://" + config.NewConfig().FrontendUri + "/passwordless/" + urlToken
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: Dislinkt passwordless login\n"
	body := ah.GenerateMailBody(resetLink, "Dislinkt passwordless login", "Press the button below to login!", "Login")
	message := []byte(subject + mime + body)
	err = ah.SendEmail(user.Email, message)
	if err != nil {
		return user.Username, err
	}
	return user.Username, nil
}

func (ah *AuthHandler) PasswordlessLogin(token string) (JWT, error) {
	user := ah.UserStore.FindByToken(token)
	if user == (store.User{}) {
		return JWT{}, errors.New("Passwordless login failed, invalid token!")
	}

	if user.TokenExpiration.Before(time.Now()) {
		return JWT{}, errors.New("Verification token expired!")
	}

	validateToken := bcrypt.CompareHashAndPassword([]byte(user.VerificationToken), []byte(token))
	if validateToken != nil {
		return JWT{}, errors.New("Passwordless login failed,, invalid token!")
	}

	if user.IsActivated == false {
		return JWT{Token: ""}, errors.New(fmt.Sprintf("User account not activated. Username: %s", user.Username))
	}

	if user.Blocked {
		if time.Now().After(user.BlockedUntil) {
			ah.UserStore.ResetFailedLogForUser(user.Username)
		} else {
			return JWT{Token: ""}, errors.New(fmt.Sprintf("User account is blocked. Username: %s", user.Username))
		}
	}
	ah.UserStore.ResetFailedLogForUser(user.Username)

	tokenStr, err := ah.GenerateJWT(JWTOptions{IsTokenNonExpired: false, Username: user.Username})
	if err != nil {
		return JWT{Token: ""}, errors.New(fmt.Sprintf("Generation of JWT faild. %s", err.Error()))
	}
	return JWT{Token: tokenStr}, nil
}

func randStr(strSize int) string {
	dictionary := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, strSize)
	rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}
	return string(bytes)
}

func (ah *AuthHandler) GenerateMailBody(url string, heading string, text string, buttonText string) string {
	body := "<html>\n" +
		"<body style=\"margin: 0 !important; padding: 0 !important;\">\n" +
		"<table border=\"0\" cellpadding=\"0\" cellspacing=\"0\" width=\"100%\">\n" +
		"<tr>\n<td bgcolor=\"#5ac5e6\" align=\"center\">\n" +
		"<table border=\"0\" cellpadding=\"0\" cellspacing=\"0\" width=\"100%\" style=\"max-width: 600px;\">\n" +
		"<tr>\n" +
		"<td align=\"center\" valign=\"top\" style=\"padding: 40px 10px 40px 10px;\"></td>\n" +
		"</tr>\n    </table>\n</td>\n" +
		"</tr>\n" +
		"<tr>\n<td bgcolor=\"#5ac5e6\" align=\"center\" style=\"padding: 0px 10px 0px 10px;\">\n" +
		"<table border=\"0\" cellpadding=\"0\" cellspacing=\"0\" width=\"100%\" style=\"max-width: 600px;\">\n" +
		"<tr>\n" +
		"<td bgcolor=\"#ffffff\" align=\"center\" valign=\"top\"\nstyle=\"padding: 40px 20px 20px 20px; " +
		"border-radius: 4px 4px 0px 0px; color: #111111; font-family: Helvetica, Arial, sans-serif; font-size: 48px; " +
		"font-weight: 400; letter-spacing: 4px; line-height: 48px;\">\n<h1 style=\"font-size: 48px; font-weight: 400; " +
		"margin: 2;\">" + heading + "</h1>\n" +
		"</td>\n" +
		"</tr>\n" +
		"</table>\n</td>\n" +
		"</tr>\n" +
		"<tr>\n<td bgcolor=\"#5ac5e6\" align=\"center\" style=\"padding: 0px 10px 0px 10px;\">\n" +
		"<table border=\"0\" cellpadding=\"0\" cellspacing=\"0\" width=\"100%\" style=\"max-width: 600px;\">\n" +
		"<tr>\n" +
		"<td bgcolor=\"#ffffff\" align=\"center\"\nstyle=\"padding: 20px 30px 40px 30px; color: #666666; " +
		"font-family: Helvetica, Arial, sans-serif; font-size: 18px; font-weight: 400; line-height: 25px;\">\n" +
		"<p style=\"margin: 0; font-size: 20px;\">" + text + "</p>\n" +
		"</td>\n" +
		"</tr>\n" +
		"<tr>\n" +
		"<td bgcolor=\"#ffffff\" align=\"left\">\n<table width=\"100%\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n" +
		"<tr>\n" +
		"<td bgcolor=\"#ffffff\" align=\"center\" style=\"padding: 20px 30px 60px 30px;\">\n" +
		"<table border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n<tr>\n" +
		"<td align=\"center\" style=\"border-radius: 3px; border: none;\"\n" +
		"bgcolor=\"#5ac5e6\">\n" +
		"<a href=\"" + url + "\"\n" +
		"style=\"font-size: 20px; font-family: Helvetica, Arial, sans-serif; color: #ffffff; text-decoration: none; " +
		"color: #ffffff; text-decoration: none; padding: 15px 25px; border-radius: 2px; display: inline-block;\">" + buttonText +
		"</a>\n" +
		"</td>\n</tr>\n" +
		"</table>\n" +
		"</td>\n" +
		"</tr>\n</table>\n" +
		"</td>\n" +
		"</tr></table><br><br></body></html>"

	return body
}

func getComponentName(methode interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(methode).Pointer()).Name()
}
