package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	sessions "github.com/kataras/go-sessions"
	"golang.org/x/oauth2"
)

var sess *sessions.Sessions

type Client struct {
	name string
}
type SessionsStore struct {
	store map[string]Client
	lock  sync.RWMutex
}

func (ss *SessionsStore) Get(id string) *Client {
	ss.lock.RLock()
	c, ok := ss.store[id]
	ss.lock.RUnlock()
	if !ok {
		return nil
	}
	return &c
}

func (ss *SessionsStore) Set(id string, c Client) {
	ss.lock.Lock()
	ss.store[id] = c
	ss.lock.Unlock()
}

var csess = SessionsStore{}

func main() {
	mux := http.NewServeMux()
	sess = sessions.New(sessions.Config{
		Cookie:  "_sessions",
		Expires: time.Hour * 2,
	})
	mux.HandleFunc("/", okHandler)
	// mux.HandleFunc("/login", handleGoogleLogin)
	// mux.HandleFunc("/callback", handleGoogleCallback)
	// mux.HandleFunc("/set", setSession)
	// mux.HandleFunc("/get", getSession)

	// mux.HandleFunc("/get", getSession)
	// mux.HandleFunc("/set", setSession)

	mux.Handle("/auth/", http.StripPrefix("/auth/", &Auth{
		sess: sessions.New(sessions.Config{
			Cookie:  "_sessions",
			Expires: time.Hour * 2,
		}),
	}))

	http.ListenAndServe(":4000", mux)
}

func setSession(w http.ResponseWriter, r *http.Request) {
	s := sess.Start(w, r)
	s.Set("name", "iris")
	w.Write([]byte(fmt.Sprintf("All ok session setted to: %s", s.GetString("name"))))
}
func getSession(w http.ResponseWriter, r *http.Request) {
	name := sess.Start(w, r).GetString("name")
	w.Write([]byte(fmt.Sprintf("The name on the /set was: %s", name)))
}
func okHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

var oauthStateString = "pseudo-random"

func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	content, err := getUserInfo(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	fmt.Fprintf(w, "Content: %s\n", content)
}
func getUserInfo(state string, code string) ([]byte, error) {
	if state != oauthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}
	return contents, nil
}
