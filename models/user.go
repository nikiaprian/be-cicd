package models

import (
	"time"
)

type User struct {
	ID        int        `json:"id"`
	Email     string     `json:"email"`
	Password  string     `json:"password,omitempty"`
	Username  string     `json:"username"`
	Name      *string    `json:"name,omitempty"`
	Photo     *string    `json:"photo"`
	Role      *string    `json:"role,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserRegisterResponse struct {
	Message string `json:"message"`
}

type UserRegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required,min=8,max=16,alphanum"`
	Password string `json:"password" validate:"required,min=8,max=16"`
}

type UserLoginResponse struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}

type UserGoogleResponse struct {
	Email    string `json:"email"`
	GoogleId string `json:"id"`
}

type UserUpdateProfileRequest struct {
	FileName *string `json:"file"`
	Username string  `json:"username,required,min=8,max=16,alphanum"`
}
