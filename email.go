package main

import "time"

type EmailStatus string

const (
	ActiveEmail   EmailStatus = "active"
	DisabledEmail EmailStatus = "disabled"
)

type Email struct {
	Email     string
	Status    EmailStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (e *Email) ActivationToken() (string, error) {
	return "", nil
}

func (e *Email) Activate(token string) error {
	return nil
}
