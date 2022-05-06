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
