package handlers

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"net/smtp"
	"strings"
	"time"
	"unicode"

	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/auth_service/startup/config"
	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/auth_service/store"
	authServicePb "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/proto/auth_service"
	profileGw "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/proto/profile_service"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ErrorMessage struct {
	Message string `json:"message"`
}
type JWT struct {
	Token string `json:"token"`
}
type AuthHandler struct {
	UserStore                *store.UsersStore
	secretKey                []byte
	profileServiceGrpcClient profileGw.ProfileServiceClient
}

type ActivationRequest struct {
	token string `json:"token"`
}

func getConnection(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
func InitAuthHandler(serverConfig *config.Config) *AuthHandler {
	userStore := store.InitUsersStore(serverConfig.MongoDbUri)
	endpoint := fmt.Sprintf("%s:%s", serverConfig.ProfileServiceGrpcHost, serverConfig.ProfileServiceGrpcPort)
	conn, err := getConnection(endpoint)
	if err != nil {
		fmt.Println("Fatal error init profile service connection!")
		return nil
	}
	client := profileGw.NewProfileServiceClient(conn)
	return &AuthHandler{UserStore: userStore, profileServiceGrpcClient: client, secretKey: []byte(serverConfig.SecretKey)}
}

func (ag *AuthHandler) LoginUserRequest(w http.ResponseWriter, r *http.Request) {
	user, err := DecodeUser(r)
	if err != nil {
		println("Error while parsing json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ag.LoginUser(user)
}
func (ah *AuthHandler) LoginUser(user store.User) (JWT, error) {
	dbUser, err := ah.UserStore.FindByUsername(user.Username)
	if err != nil {
		fmt.Println("User not found")
		return JWT{Token: ""}, err
	}
	if dbUser.IsActivated == false {
		err := errors.New("Account not activated!")
		return JWT{Token: ""}, err
	}
	if !CheckPasswordHash(user.Password, dbUser.Password) {
		err := ah.HandleFailedLogin(dbUser)
		return JWT{Token: ""}, err
	}
	if dbUser.Blocked {
		if time.Now().After(user.BlockedUntil) {
			ah.UserStore.ResetFailedLogForUser(user.Username)
		} else {
			return JWT{Token: ""}, errors.New("Blocked account")
		}
	}
	ah.UserStore.ResetFailedLogForUser(user.Username)

	tokenStr, err := GenerateJWT(dbUser, ah.secretKey)
	if err != nil {
		fmt.Printf("Token generation failed %s\n", err.Error())
		return JWT{Token: ""}, err
	}
	return JWT{Token: tokenStr}, nil
}

func (ag *AuthHandler) AuthorizeJWT(w http.ResponseWriter, r *http.Request) {
	if r.Header["Authorization"] != nil {
		tokenStr := strings.Split(r.Header["Authorization"][0], " ")[1]
		ag.ValidateToken(tokenStr)
	}
}

//TODO: Validarati korisnika u bazi dodatno mozda je u medjuvremenu obrisan ili mijenjao lozinku itd.
func (ag *AuthHandler) ValidateToken(tokenStr string) (*store.User, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error")
		}
		return ag.secretKey, nil
	})
	if err != nil {
		fmt.Println("Error")
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("JWT not valid")
	}
	fmt.Println("VALID")
	username := token.Claims.(jwt.MapClaims)["username"]
	str := fmt.Sprintf("%v", username)
	fmt.Println(str)
	return &store.User{Username: str}, nil

}
func (ag *AuthHandler) AddNewUser(user store.User) error {
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
	user.TokenExpiration = time.Now().AddDate(0, 0, 1)
	ag.SendEmail("andjela.ra28@gmail.com", ag.MailActivationMessage(urlToken))

	err := ag.UserStore.AddNew(user)
	if err != nil {
		return err
	}
	return nil
}

func (ag *AuthHandler) NotifyProfileServiceAboutRegistration(in *authServicePb.User) error {
	fmt.Println("Post request")
	_, err := ag.profileServiceGrpcClient.AddNewUser(context.TODO(), &profileGw.UserRequest{User: &profileGw.User{
		Username:  in.Username,
		Password:  in.Password,
		Biography: in.Biography,
		BirthDate: in.BirthDate,
		Email:     in.Email,
		Gender:    in.Gender,
		IsPublic:  in.IsPublic,
		Name:      in.Name,
		Telephone: in.Telephone,
		Surname:   in.Surname,
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

func GenerateJWT(dbUser store.User, secretString []byte) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["username"] = dbUser.Username
	claims["exp"] = time.Now().Add(time.Minute * 80).Unix()
	tokenStr, err := token.SignedString(secretString)
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

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
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
			return errors.New("Blocked account")
		}
	} else {
		if user.FailedLogins == 5 {
			ah.UserStore.BlockUser(user.Username)
			return errors.New("Blocked account")
		}
	}
	_ = ah.UserStore.UpdateFailedLogForUser(user.Username)
	return errors.New("Invalid credentials")
}

func (ah *AuthHandler) GenerateVerificationToken() (string, string) {
	randomBytes := make([]byte, 10)
	rand.Seed(time.Now().UnixNano())
	rand.Read(randomBytes)
	fmt.Println(randomBytes)
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
		return &user, errors.New("Wrong old password")
	} else {
		if !isPasswordValid(request.NewPassword) {
			return &user, errors.New("Incorrect password format")
		}
		newPassword, _ := HashPassword(request.NewPassword)
		err := ah.UserStore.UpdatePassword(request.Username, newPassword)
		if err != nil {
			return &user, err
		}
	}
	return &user, nil
}

func (ah *AuthHandler) SendEmail(emailTo string, mailMessage []byte) {
	from := config.NewConfig().EmailFrom
	password := config.NewConfig().EmailPassword
	to := []string{emailTo}
	host := config.NewConfig().EmailHost
	port := config.NewConfig().EmailPort
	address := host + ":" + port

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(address, auth, from, to, mailMessage)
	if err != nil {
		panic(err)
	}
}

func (ah *AuthHandler) MailActivationMessage(token string) []byte {
	confirmationLink := "http://localhost:4200/account-activation/" + token

	subject := "Account activation\n"
	body := "Please click on following link to confirm your email <a href='" + confirmationLink + "'>link</a>"
	message := []byte(subject + body)
	return message
}

func (ah *AuthHandler) ActivateAccount(token string) error {
	user := ah.UserStore.FindByToken(token)
	if user == (store.User{}) {
		return errors.New("Activation failed, invalid token!")
	}

	if user.TokenExpiration.Before(time.Now()) {
		return errors.New("Activation token expired!")
	}

	validateToken := bcrypt.CompareHashAndPassword([]byte(user.VerificationToken), []byte(token))
	if validateToken != nil {
		return errors.New("Activation failed, invalid token!")
	}

	err := ah.UserStore.ActivateAccount(user.Username)
	return err
}

func (ah *AuthHandler) AccountRecoveryEmail(email string) error {
	user, noUserFound := ah.UserStore.FindByEmail(email)

	if noUserFound != nil {
		return errors.New("Account with this email doesn't exist!")
	}

	verificationHash, urlToken := ah.GenerateVerificationToken()

	invalidToken := ah.UserStore.RefreshToken(user.Username, verificationHash)
	if invalidToken != nil {
		return errors.New("Error occured while updating token!")
	}

	resetLink := "http://localhost:4200/set-password/" + urlToken
	subject := "Dislinkt password reset\n"
	body := "Please click on following link to set your new account password <a href='" + resetLink + "'>link</a>"
	message := []byte(subject + body)
	ah.SendEmail(user.Email, message)
	return nil
}

func (ah *AuthHandler) SendPasswordlessLoginEmail(email string) error {
	user, noUserFound := ah.UserStore.FindByEmail(email)

	if noUserFound != nil {
		return errors.New("Account with this email doesn't exist!")
	}

	verificationHash, urlToken := ah.GenerateVerificationToken()
	invalidToken := ah.UserStore.RefreshToken(user.Username, verificationHash)
	if invalidToken != nil {
		return errors.New("Error occured while updating token!")
	}

	resetLink := "http://localhost:4200/passwordless/" + urlToken
	subject := "Dislinkt password reset\n"
	body := "Please click on following link to log into your account <a href='" + resetLink + "'>link</a>"
	message := []byte(subject + body)
	ah.SendEmail(user.Email, message)
	return nil
}

func (ah *AuthHandler) PasswordlessLogin(token string) (JWT, error) {
	user := ah.UserStore.FindByToken(token)
	if user == (store.User{}) {
		return JWT{}, errors.New("Activation failed, invalid token!")
	}

	if user.TokenExpiration.Before(time.Now()) {
		return JWT{}, errors.New("Activation token expired!")
	}

	validateToken := bcrypt.CompareHashAndPassword([]byte(user.VerificationToken), []byte(token))
	if validateToken != nil {
		return JWT{}, errors.New("Activation failed, invalid token!")
	}

	jwt, err2 := ah.LoginChecks(user)
	if err2 == nil {
		return JWT{}, err2
	}
	return jwt, nil
}

func (ah *AuthHandler) LoginChecks(user store.User) (JWT, error) {
	if user.IsActivated == false {
		err := errors.New("Account not activated!")
		return JWT{Token: ""}, err
	}
	if user.Blocked {
		if time.Now().After(user.BlockedUntil) {
			ah.UserStore.ResetFailedLogForUser(user.Username)
		} else {
			return JWT{Token: ""}, errors.New("Blocked account")
		}
	}
	ah.UserStore.ResetFailedLogForUser(user.Username)

	tokenStr, err := GenerateJWT(user, ah.secretKey)
	if err != nil {
		fmt.Printf("Token generation failed %s\n", err.Error())
		return JWT{Token: ""}, err
	}
	return JWT{Token: tokenStr}, nil
}
