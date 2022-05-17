package application

import (
	"fmt"
	"github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/profile-service/store"
)

type ProfileService struct {
	store store.ProfileStore
}

func NewProfileService(store store.ProfileStore) *ProfileService {
	return &ProfileService{
		store: store,
	}
}
func (service *ProfileService) AddUser(user store.User) {
	service.store.AddUser(user)
	fmt.Println("Added user:" + user.Username)
}
