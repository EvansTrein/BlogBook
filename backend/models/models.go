package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `gorm:"not null"`
	Email string `gorm:"not null;unique"`
	Hash  string `gorm:"hash" json:"-"`
}

type DataUser struct {
	Name string  `json:"name"`
	Email string `json:"email" binding:"required,email"`
	// requires a valid password with a minimum length of 8 characters
	Password string `json:"password" binding:"required,min=8"`
}

type UpdDataUser struct {
    Name     string `json:"name" binding:"required"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"omitempty,min=8"`
}

type Tokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type ResponceError struct {
	ErrText string
}

type ResponseMessage struct {
	Message string
}