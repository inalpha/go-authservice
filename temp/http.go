package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type Handler struct {
}

const contentTypeHTML = "text/html; charset=utf-8"
const contentTypeJWT = "application/jwt"
const contentTypeJSON = "application/json"
const contentTypePlain = "text/plain"

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	var m LoginRequest
	if err := m.FromJSON(r); err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	fmt.Println(m)
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, contentTypeJSON) {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
}

type LoginRequest struct {
	Email    string
	Password string
	Remember bool
}

func (lr *LoginRequest) FromJSON(r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(lr); err != nil {
		return err
	}
	if err := lr.validate(); err != nil {
		return err
	}
	return nil
}

func (lr *LoginRequest) FromPOST(r *http.Request) error {
	lr.Email = r.PostForm.Get("email")
	lr.Password = r.PostForm.Get("password")
	lr.Remember, _ = strconv.ParseBool(r.PostForm.Get("remember"))
	if err := lr.validate(); err != nil {
		return err
	}
	return nil
}

var rxEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func (lr *LoginRequest) validate() (err error) {
	if len(lr.Email) > 254 || !rxEmail.MatchString(lr.Email) {
		return fmt.Errorf("error: please use a valid email address")
	}
	if len(lr.Password) < 8 {
		return fmt.Errorf("error: please use a strong password")
	}
	return nil
}

type SignUpRequest struct {
	Email    string
	Password string
	Repeat   string
}

func getJSONCredentials(r *http.Request) (string, string, error) {
	m := map[string]string{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m); err != nil {
		// http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return "", "", err
	}
	return m["username"], m["password"], nil
}

func getPOSTCredentials(r *http.Request) (string, string, error) {
	return r.PostForm.Get("username"), r.PostForm.Get("password"), nil
}
