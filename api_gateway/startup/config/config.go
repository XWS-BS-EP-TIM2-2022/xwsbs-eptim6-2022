package config

import (
	"regexp"
)

type SecurityPermissions struct {
	permissions map[string][]string
}

type Config struct {
	Port                string
	AuthHost            string
	AuthPort            string
	PostsHost           string
	PostsPort           string
	ProfileHost         string
	ProfilePort         string
	SecurityPermissions *SecurityPermissions
}
type Permission string

const (
	CREATE_POSTS Permission = "CREATE_POSTS"
	VIEW_POSTS   Permission = "VIEW_POSTS"
	UPDATE_POSTS Permission = "UPDATE_POSTS"
	DELETE_POSTS Permission = "DELETE_POSTS"

	CREATE_USER Permission = "CREATE_USER"
	UPDATE_USER Permission = "UPDATE_USER"
	VIEW_USER   Permission = "VIEW_USER"
	DELETE_USER Permission = "DELETE_USER"
)

func (secPermissions *SecurityPermissions) ValidateUnauthorizedRequest(request string) bool {
	for _, values := range secPermissions.permissions {
		for i := 0; i < len(values); i++ {
			reg := regexp.MustCompile(values[i])
			matchString := reg.FindString(request)
			if matchString != "" {
				return false
			}
		}
	}
	return true
}

func (secPermissions *SecurityPermissions) ValidatePermission(userPermissions []string, request string) bool {
	if secPermissions.ValidateUnauthorizedRequest(request) {
		return true
	}
	for i := 0; i < len(userPermissions); i++ {
		requests := secPermissions.permissions[userPermissions[i]]
		for i := 0; i < len(requests); i++ {
			reg := regexp.MustCompile(requests[i])
			matchString := reg.FindString(request)
			if matchString != "" {
				return true
			}
		}
	}
	return false
}

func NewConfig() *Config {
	permissions := SecurityPermissions{permissions: map[string][]string{}}
	permissions.permissions[string(VIEW_USER)] = []string{"GET/api/auth/users", "GET/api/users"}
	permissions.permissions[string(CREATE_POSTS)] = []string{"POST/api/posts"}
	permissions.permissions[string(VIEW_POSTS)] = []string{"GET/api/posts"}
	permissions.permissions[string(UPDATE_USER)] = []string{"PUT/api/users/experience", "PUT/users/follow/[a-zA-Z0-9]+"}
	permissions.permissions[string(DELETE_USER)] = []string{"DELETE/api/users/[a-zA-Z0-9]+"}
	return &Config{
		SecurityPermissions: &permissions,
		Port:                "5000",
		AuthHost:            "localhost",
		AuthPort:            "5001",
		PostsHost:           "localhost",
		PostsPort:           "5002",
		ProfileHost:         "localhost",
		ProfilePort:         "5003",
	}
	//return &Config{
	//	SecurityPermissions: &permissions,
	//	Port:                os.Getenv("GATEWAY_PORT"),
	//	AuthHost:            os.Getenv("AUTH_SERVICE_HOST"),
	//	AuthPort:            os.Getenv("AUTH_SERVICE_PORT"),
	//	PostsHost:           os.Getenv("POSTS_SERVICE_HOST"),
	//	PostsPort:           os.Getenv("POSTS_SERVICE_PORT"),
	//	ProfileHost:         os.Getenv("PROFILE_SERVICE_HOST"),
	//	ProfilePort:         os.Getenv("PROFILE_SERVICE_PORT"),
	//}
}
