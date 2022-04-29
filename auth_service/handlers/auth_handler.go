package handlers

import (
	"auth_service/store"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
	"time"
)

var secretString = []byte("secret_key") //TODO: Use ENV Variable

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

func (ag *AuthHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	user, err := DecodeUser(r)
	if err != nil {
		println("Error while parsing json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	dbUser := ag.UserStore.FindByUsername(user.Username)
	if dbUser.Password == user.Password {
		tokenStr, err := GenerateJWT(dbUser)
		if err != nil {
			fmt.Printf("Token generation failed %s\n", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode(JWT{Token: tokenStr})
	} else {
		fmt.Println("Login failed")
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (ag *AuthHandler) AuthorizeJWT(w http.ResponseWriter, r *http.Request) {
	if r.Header["Authorization"] != nil {
		tokenStr := strings.Split(r.Header["Authorization"][0], " ")[1]
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("error")
			}
			return secretString, nil
		})
		if err != nil {
			fmt.Println("Error")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if token.Valid {
			fmt.Println("VALID")
			json.NewEncoder(w).Encode("valid")
		}
	}
}

func (ag *AuthHandler) AddNewUser(w http.ResponseWriter, r *http.Request) {
	user, err := DecodeUser(r)
	if err != nil {
		println("Error while parsing json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ag.UserStore.AddNew(user)
	fmt.Println("Post request")
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
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	tokenStr, err := token.SignedString(secretString)
	if err != nil {
		fmt.Errorf("token signing error")
		return "", err
	}
	return tokenStr, nil
}
