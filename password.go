package main

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type PasswordStatus string

type Password struct {
	Hash      []byte
	Salt      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func New(password, salt string) (*Password, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &Password{
		Hash:      hash,
		Salt:      salt,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (p *Password) Compare(password string) error {
	return bcrypt.CompareHashAndPassword(
		p.Hash,
		[]byte(password+p.Salt),
	)
}

func (p *Password) ResetToken() (string, error) {
	return "", nil
}

func (p *Password) Reset(token, password string) error {
	return nil
}
