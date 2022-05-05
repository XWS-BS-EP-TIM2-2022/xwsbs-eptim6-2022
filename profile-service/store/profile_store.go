package store

type Profile struct {
	User      User
	Followers []User
	Following []User
	IsPublic  bool
}

// func getProfile(w http.ResponseWriter, r *http.Request) {
// 	user := User{Username: "lenche", Name: "Lenka Isidora", Surname: "Aleksic", Password: "123"}
// 	followers := []User{user}
// 	following := []User{user}
// 	isP := true
// 	profile := Profile{user, followers, following, isP}
// 	fmt.Println("Profile: ")
// 	json.NewEncoder(w).Encode(profile)
// }
