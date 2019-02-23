package main

import (
	"time"

	"github.com/gofrs/uuid"
)

type AccountStatus string

const (
	ActiveAccount   AccountStatus = "active"
	DisabledAccount AccountStatus = "disabled"
)

type Account struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Status    AccountStatus
	Emails    []Email
	Providers []Provider
	Password  *Password
}

func NewAccount(email, password string) (*Account, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return nil, nil
	}
	p, err := New(password, "")

	if err != nil {
		return nil, nil
	}

	return &Account{
		ID:       id,
		Password: p,
	}, nil
}

func authorize(email, password string) error {
	return nil
}

func createWithOAuth(provider string, token string, userID string) error {
	return nil
}

func validateEmail(token string) error {
	return nil
}

func resetPassword(token string) error {
	return nil
}

func getPasswordResetTokenForEmail(email string) error {
	return nil
}

func getPasswordResetTokenForID(id string) error {
	return nil
}
