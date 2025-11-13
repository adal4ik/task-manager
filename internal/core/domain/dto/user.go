package dto

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterUser struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginUser struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
