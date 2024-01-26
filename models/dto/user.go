package dto

type UserLoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserLoginResponse struct {
	UserData UserDataResponse
	Token    string
}

type UserDataResponse struct {
	Id    int
	Name  string
	Email string
	Role  string
}
