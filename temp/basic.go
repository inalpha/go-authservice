package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// - User signed in and the service already exists
// - User signed in and they don’t have this service
// - User signed out and they don’t have an account at all
// - User signed out and they have an account that matches
// - User signed out and they use new service

// If the email address you pull from Facebook is not in your database :
// “We haven’t found an account linked to your Facebook profile”
// - I want to sign up using Facebook
// - I want to log into my account using another method

// If the email address you pull from Facebook is in your database, but registered with a different service :
// “Oops ! You don’t usually sign in using Facebook ! We’ve sent you an email to remind you which service you usually do use.”

func add(email, password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("unable to encrypt password")
	}
	//TODO: save user into storage and retrive id
	return string(hash), nil
}

func login(email, password string) (string, error) {
	//TODO: get user with that email, hash, id form store fail it not found
	hash := []byte("")
	if bcrypt.CompareHashAndPassword(hash, []byte(password)) != nil {
		return "", fmt.Errorf("wrong password")
	}
	return "id", nil
}

func loginWithProvider(provider, token, email string) (string, error) {
	//TODO get user with email if does not exist error
	return signupWithProvider(provider, token, email)

	//TODO user with email exist and provider does not exist
	//fmt.Errorf("user does not have provider linked")

	//TODO if user exisit and provider exist
	//TODO update token, refresh token

	// return "id", nil
}

func signupWithProvider(provider, token, email string) (string, error) {
	//TODO create new user in database
	addProvider(provider, token, "id")

	return "id", nil
}

func addProvider(provider, token, id string) error {
	//TODO Create auth provider for that user
	return nil
}
