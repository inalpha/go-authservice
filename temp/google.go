package main

import (
	"os"

	"golang.org/x/oauth2/google"

	"golang.org/x/oauth2"
)

var googleOauthConfig *oauth2.Config

func init() {
	println(os.Getenv("GOOGLE_CLIENT_ID"))
	println(os.Getenv("GOOGLE_CLIENT_SECRET"))
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:4000/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}
