package main

import (
	"time"
)

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//TODO: might need email and id of the user from the service
type OAuth struct {
	Email        string    `json:"email"`
	UserID       string    `json:"user_id"`
	Service      string    `json:"service"`
	TokenType    string    `json:"token_type,omitempty"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token,omitempty"`
	Expiry       time.Time `json:"expiry,omitempty"`
}

type OAuthStore interface {
	Find(userID string, service string) (*OAuth, error)
	Create(oauth *OAuth) error
	Delete(id string) error
}

type UserStore interface {
	ByID(id string) (*User, error)
	ByEmail(email string) (*User, error)

	Create(user *User) error
	Update(user *User) error
	Delete(id string) error
}
