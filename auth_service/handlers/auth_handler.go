package handlers

import (
	"auth_service/store"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
	"time"
)

var secretString = []byte("secret_key") //TODO: Use ENV Variable
type ErrorMessage struct {
	Message string `json:"message"`
}
type JWT struct {
	Token string `json:"token"`
}
type AuthHandler struct {
	UserStore *store.UsersStore
}

func InitAuthHandler() *AuthHandler {
	userStore := store.InitUsersStore()
	return &AuthHandler{UserStore: userStore}
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
	tokenStr, err := GenerateJWT(dbUser)
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
		return secretString, nil
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
func (ag *AuthHandler) AddNewUser(w http.ResponseWriter, r *http.Request) {
	user, err := DecodeUser(r)
	if _, err := ag.UserStore.FindByUsername(user.Username); err == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorMessage{Message: "Username already in use"})
		return
	}
	if err != nil {
		println("Error while parsing json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ag.UserStore.AddNew(user)
	err = ag.notifyProfileServiceAboutRegistration(user)
	if err != nil {
		println("Error while parsing json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (ag *AuthHandler) notifyProfileServiceAboutRegistration(user store.User) error {
	fmt.Println("Post request")
	postBody, _ := json.Marshal(user)
	requestBody := bytes.NewBuffer(postBody)
	_, err := http.Post("http://localhost:8081/users", "application/json", requestBody)
	if err != nil {
		return err
	}
	return nil
}

func (ag *AuthHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users := ag.UserStore.FindAll()
	json.NewEncoder(w).Encode(users)
}

func DecodeUser(req *http.Request) (store.User, error) {
	var user store.User
	err := json.NewDecoder(req.Body).Decode(&user)
	return user, err
}

func GenerateJWT(dbUser store.User) (string, error) {
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
