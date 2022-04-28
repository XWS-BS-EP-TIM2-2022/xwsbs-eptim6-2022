package store

import "fmt"

type User struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
}

type UsersStore struct {
	Users []User
}

func (us *UsersStore) AddNew(u User) {
	us.Users = append(us.Users, u)
}
func (us *UsersStore) FindByUsername(username string) User {
	var user User
	for i := range us.Users {
		fmt.Println(i)
		user = us.Users[i]
	}
	return user
}
