package main

import (
	"log"

	jwt "gopkg.in/square/go-jose.v2/jwt"
)

func main() {
	password, err := New("Hello", "1234")
	if err != nil {
		log.Fatal(err)
	}
	token, err := password.ResetToken()
	if err != nil {
		log.Fatal(err)
	}
	jwt.Signed
	log.Println(token)
}
