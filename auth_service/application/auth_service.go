package application

import (
	"auth_service/store"
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
	return service.store.FindAll()
}
