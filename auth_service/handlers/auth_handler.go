package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/auth_service/startup/config"
	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/auth_service/store"
	authServicePb "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/proto/auth_service"
	profileGw "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/proto/profile_service"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"strings"
	"time"
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
	if dbUser.Password != user.Password {
		fmt.Println("User not found")
		return JWT{Token: ""}, err
	}
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
	ag.UserStore.AddNew(user)
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
