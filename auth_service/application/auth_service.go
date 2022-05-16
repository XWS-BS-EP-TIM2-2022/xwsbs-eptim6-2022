package application

import (
	"fmt"
	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/auth_service/store"
)

type AuthService struct {
	store store.AuthStore
}

func NewAuthService(store store.AuthStore) *AuthService {
	return &AuthService{
		store: store,
	}
}

func (service *AuthService) FindByUsername(username string) (store.User, error) {
	return service.store.FindByUsername(username)
}

func (service *AuthService) FindAll() []store.User {
	fmt.Println("FindAll auth_service")
	return service.store.FindAll()
}

func (service *AuthService) AddNewUser(user store.User) {
	service.store.AddNew(user)
	fmt.Println("Added user:" + user.Username)
}

func (service *AuthService) LoginUser(user store.User) string {
	token, _ := service.store.LoginUser(user)
	fmt.Println("Logged in user:" + user.Username)
	fmt.Println("token: " + token)
	return token
}
