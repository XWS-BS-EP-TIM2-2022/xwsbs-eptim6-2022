package mappers

import (
	usersServicePb "github.com/XWS-BS-EP-TIM2-2022/xwsbs-eptim6-2022/common/proto/profile_service"
	"profile-service/store"
)

func MapToUser(in *usersServicePb.User) *store.User {
	return &store.User{
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
	}
}
func MapToUsersResponse(users []store.User) *usersServicePb.UsersResponse {
	var response []*usersServicePb.User
	for _, user := range users {
		response = append(response, mapToUserPb(&user))
	}
	return &usersServicePb.UsersResponse{Users: response}
}
func mapToUserPb(usr *store.User) *usersServicePb.User {
	return &usersServicePb.User{
		Username:       usr.Username,
		Password:       usr.Password,
		Biography:      usr.Biography,
		BirthDate:      usr.BirthDate,
		Email:          usr.Email,
		Gender:         usr.Gender,
		IsPublic:       usr.IsPublic,
		Name:           usr.Name,
		Telephone:      usr.Telephone,
		Surname:        usr.Surname,
		Educations:     mapEducations(usr.Educations),
		Experiences:    mapExperiences(usr.Experiences),
		Skills:         mapSkills(usr.Skills),
		Followers:      mapFollowers(usr.Followers),
		Followings:     mapFollowings(usr.Followings),
		FollowRequests: mapFollowRequests(usr.FollowRequests),
		Interests:      mapInterests(usr.Interests),
	}
}
func MapToUserResponse(usr *store.User) *usersServicePb.UserResponse {
	return &usersServicePb.UserResponse{User: mapToUserPb(usr)}
}

func mapFollowRequests(requests []string) []*usersServicePb.Follower {
	var response []*usersServicePb.Follower
	for _, req := range requests {
		response = append(response, &usersServicePb.Follower{Username: req})
	}
	return response
}

func mapInterests(interests []store.Interest) []*usersServicePb.Interest {
	var response []*usersServicePb.Interest
	for _, req := range interests {
		response = append(response, &usersServicePb.Interest{Text: req.Text})
	}
	return response
}

func mapFollowings(followings []store.Following) []*usersServicePb.Follower {
	var response []*usersServicePb.Follower
	for _, req := range followings {
		response = append(response, &usersServicePb.Follower{Username: req.Username})
	}
	return response
}

func mapFollowers(followers []store.Follower) []*usersServicePb.Follower {
	var response []*usersServicePb.Follower
	for _, req := range followers {
		response = append(response, &usersServicePb.Follower{Username: req.Username})
	}
	return response
}

func mapSkills(skills []store.Skill) []*usersServicePb.Skill {
	var response []*usersServicePb.Skill
	for _, req := range skills {
		response = append(response, &usersServicePb.Skill{Text: req.Text})
	}
	return response
}

func mapExperiences(experiences []store.Experience) []*usersServicePb.Experience {
	var response []*usersServicePb.Experience
	for _, req := range experiences {
		response = append(response, &usersServicePb.Experience{Text: req.Text})
	}
	return response
}

func mapEducations(educations []store.Education) []*usersServicePb.Education {
	var educationResponse []*usersServicePb.Education
	for _, education := range educations {
		educationResponse = append(educationResponse, &usersServicePb.Education{Text: education.Text})
	}
	return educationResponse
}
