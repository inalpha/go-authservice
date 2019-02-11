package main

import (
	"fmt"
	"net/http"
	"strings"

	sessions "github.com/kataras/go-sessions"
)

type Auth struct {
	sess  *sessions.Sessions
	store AccountStore
}

func (a *Auth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	if r.Method != "POST" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	if r.URL.Path == "login" {
		a.login(w, r)
		return
	}
	if r.URL.Path == "logout" {
		a.logout(w, r)
		return
	}
	if r.URL.Path == "signup" {
		a.signup(w, r)
		return
	}

	http.Error(w, http.StatusText(400), http.StatusBadRequest)
	return
}

func (a *Auth) signup(w http.ResponseWriter, r *http.Request) {
	s := a.sess.Start(w, r)
	s.Set("name", "iris")
	a.store.Add()
	w.Write([]byte(fmt.Sprintf("All ok session setted to: %s", s.GetString("name"))))
}

func (a *Auth) login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("All ok session setted to: %s", a.sess.Start(w, r).GetString("name"))))
}

func (a *Auth) logout(w http.ResponseWriter, r *http.Request) {
	s := a.sess.Start(w, r)
	s.Delete("name")
	w.Write([]byte("Deleted"))
}

type AccountStore interface {
	Add(email, password string) error
	Get(email, password string) error
}

type InMemory struct {
	store map[string]string
}

func (im *InMemory) Add(email, password string) error {
	return nil
}

func (im *InMemory) Get(email, password string) error {
	return nil
}
