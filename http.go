package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Handler struct {
}

const contentTypeHTML = "text/html; charset=utf-8"
const contentTypeJWT = "application/jwt"
const contentTypeJSON = "application/json"
const contentTypePlain = "text/plain"

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, contentTypeJSON) {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	var m map[string]string
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m); err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	fmt.Println(m)
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
