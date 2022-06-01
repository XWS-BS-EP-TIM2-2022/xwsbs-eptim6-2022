package consts

type Permission string
type Role string

const (
	ADMIN         Role = "ADMIN"
	USER          Role = "USER"
	COMPANY_OWNER Role = "COMPANY_OWNER"
)
const (
	CREATE_POSTS Permission = "CREATE_POSTS"
	VIEW_POSTS   Permission = "VIEW_POSTS"
	UPDATE_POSTS Permission = "UPDATE_POSTS"
	DELETE_POSTS Permission = "DELETE_POSTS"

	CREATE_USER Permission = "CREATE_USER"
	UPDATE_USER Permission = "UPDATE_USER"
	VIEW_USER   Permission = "VIEW_USER"
	DELETE_USER Permission = "DELETE_USER"

	CREATE_JOB_OFFER Permission = "CREATE_JOB_OFFER"
	UPDATE_JOB_OFFER Permission = "UPDATE_JOB_OFFER"
	VIEW_JOB_OFFER   Permission = "VIEW_JOB_OFFER"
	DELETE_JOB_OFFER Permission = "DELETE_JOB_OFFER"
)
