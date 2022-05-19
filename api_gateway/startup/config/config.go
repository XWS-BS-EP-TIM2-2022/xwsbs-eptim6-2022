package config

import "os"

type SecurityPermissions struct {
	permissions map[string]string
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
	if _, ok := secPermissions.permissions[request]; !ok {
		return true
	}
	return false
}

func (secPermissions *SecurityPermissions) ValidatePermission(userPermissions []string, request string) bool {
	if _, ok := secPermissions.permissions[request]; !ok {
		return true
	}
	for i := 0; i < len(userPermissions); i++ {
		if userPermissions[i] == secPermissions.permissions[request] {
			return true
		}
	}
	return false
}

func NewConfig() *Config {
	permissions := SecurityPermissions{permissions: map[string]string{}}
	permissions.permissions["GET/api/auth/users"] = string(VIEW_USER)
	permissions.permissions["POST/api/posts"] = string(CREATE_POSTS)
	permissions.permissions["GET/api/posts"] = string(VIEW_POSTS)
	permissions.permissions["PUT/api/users/experience"] = string(UPDATE_USER)
	permissions.permissions["PUT/users/follow/{id}"] = string(UPDATE_USER) //TODO: VIDITI KAKO KORISTITI REGEX

	return &Config{
		SecurityPermissions: &permissions,
		Port:                "5000",      ///os.Getenv("GATEWAY_PORT"),
		AuthHost:            "localhost", //os.Getenv("AUTH_SERVICE_HOST"),
		AuthPort:            "8080",      //os.Getenv("AUTH_SERVICE_PORT"),
		PostsHost:           os.Getenv("ORDERING_SERVICE_HOST"),
		PostsPort:           os.Getenv("ORDERING_SERVICE_PORT"),
		ProfileHost:         os.Getenv("SHIPPING_SERVICE_HOST"),
		ProfilePort:         os.Getenv("SHIPPING_SERVICE_PORT"),
	}
}
