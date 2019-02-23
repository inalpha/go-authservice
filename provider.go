package main

import (
	"time"

	"github.com/gofrs/uuid"
)

type Provider struct {
	ID           uuid.UUID
	AccessToken  string
	TokenType    string
	RefreshToken string
	Expiry       time.Time
}
