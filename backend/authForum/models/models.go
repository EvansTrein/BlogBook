package models

type User struct {
	Id           uint   `json:"id"`
	UserName     string `json:"name"`
	UserEmail    string `json:"email"`
	PasswordHash string `json:"password"`
}

type DataUser struct {
	Name     string `json:"name"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type UpdDataUser struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"omitempty,min=8"`
}

type RespActiveUser struct {
	Id         uint   `json:"id"`
	UserName   string `json:"name"`
	UserEmail  string `json:"email"`
	UserTokens Tokens
}

type Tokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type ResponceError struct {
	ErrMess string
	Err     string
}

type ResponseMessage struct {
	Message string
}
