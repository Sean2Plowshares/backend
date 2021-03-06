package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/praelatus/backend/models"
	"github.com/praelatus/backend/mw"
	"github.com/praelatus/backend/store"
)

func initUserRoutes() {
	Router.Handle("/users/{username}", mw.Default(UpdateUser)).Methods("PUT")
	Router.Handle("/users/{username}", mw.Default(DeleteUser)).Methods("DELETE")
	Router.Handle("/users/{username}", mw.Default(GetUser)).Methods("GET")
	Router.Handle("/users", mw.Default(GetAllUsers)).Methods("GET")
	Router.Handle("/users", mw.Default(CreateUser)).Methods("POST")

	Router.Handle("/sessions", mw.Default(CreateSession)).Methods("POST")
	Router.Handle("/sessions", mw.Default(RefreshSession)).Methods("GET")
}

// TokenResponse is used when logging in or signing up, it will return a
// generated token plus the user model for use by the client.
type TokenResponse struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}

// GetUser will get a user from the database by the given username
func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	u := models.User{
		Username: vars["username"],
	}

	err := Store.Users().Get(&u)
	if err != nil {
		if err == store.ErrNotFound {
			w.WriteHeader(404)
			w.Write(apiError("No user exists with that username."))
			return
		}

		w.WriteHeader(500)
		w.Write(apiError(err.Error()))
		log.Println(err)
		return
	}

	u.Password = ""

	sendJSON(w, u)
}

// GetAllUsers will return the json encoded array of all users in the given
// store
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	u := mw.GetUser(r.Context())
	if u == nil {
		w.WriteHeader(403)
		w.Write(apiError("you must be logged in to view other users"))
		return
	}

	users, err := Store.Users().GetAll()
	if err != nil {
		w.WriteHeader(500)
		w.Write(apiError(err.Error()))
		log.Println(err)
		return
	}

	sendJSON(w, users)
}

// CreateUser will take the JSON given and attempt to
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var u models.User

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&u)
	if err != nil {
		w.WriteHeader(400)
		w.Write(apiError(err.Error()))
		log.Println(err)
		return
	}

	err = Store.Users().New(&u)
	if err != nil {
		if err == store.ErrDuplicateEntry {
			w.WriteHeader(400)
			w.Write(apiError(err.Error()))
			return
		}

		w.WriteHeader(500)
		w.Write(apiError(err.Error()))
		log.Println(err)
		return
	}

	token, err := mw.JWTSignUser(u)
	if err != nil {
		w.WriteHeader(500)
		w.Write(apiError(err.Error()))
		log.Println(err)
		return
	}

	sendJSON(w, TokenResponse{
		token,
		u,
	})
}

// UpdateUser will update a user in the database, it will reject the call if
// the user sending is not the user being updated or if the user sending is not
// a sys admin
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var u models.User

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&u)
	if err != nil {
		w.WriteHeader(400)
		w.Write(apiError(err.Error()))
		log.Println(err)
		return
	}

	err = Store.Users().Save(u)
	if err != nil {
		w.WriteHeader(500)
		w.Write(apiError(err.Error()))
		log.Println(err)
		return
	}

	sendJSON(w, u)
}

// DeleteUser will remove a user from the database by setting is_inactive = 1
// can only be used by sys admins
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var u models.User

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&u)
	if err != nil {
		w.WriteHeader(400)
		w.Write(apiError(err.Error()))
		log.Println(err)
		return
	}

	err = Store.Users().Remove(u)
	if err != nil {
		w.WriteHeader(500)
		w.Write(apiError(err.Error()))
		log.Println(err)
		return
	}

	w.Write([]byte(""))
}

// CreateSession will log in a user and create a jwt token for the current
// session
func CreateSession(w http.ResponseWriter, r *http.Request) {
	type loginRequest struct {
		Username string `json:"username"`
		Password string `json:"passoword"`
	}

	var l loginRequest

	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&l)
	if err != nil {
		w.WriteHeader(400)
		w.Write(apiError(err.Error()))
		log.Println(err)
		return
	}

	u := models.User{Username: l.Username}

	err = Store.Users().Get(&u)
	if err != nil {
		if err == store.ErrNotFound {
			w.WriteHeader(404)
			w.Write(apiError("No user exists with that username."))
			return
		}

		w.WriteHeader(500)
		w.Write(apiError(err.Error()))
		log.Println(err)
		return
	}

	if u.CheckPw([]byte(l.Password)) {
		u.Password = ""
		token, err := mw.JWTSignUser(u)
		if err != nil {
			w.WriteHeader(500)
			w.Write(apiError(err.Error()))
			log.Println(err)
			return

		}

		sendJSON(w, TokenResponse{
			token,
			u,
		})

		return
	}

	w.WriteHeader(401)
	w.Write(apiError("invalid password", "password"))
}

// RefreshSession will reset the expiration on the current jwt token
func RefreshSession(w http.ResponseWriter, r *http.Request) {
	u := mw.GetUser(r.Context())
	if u == nil {
		w.WriteHeader(401)
		w.Write(apiError("you must be logged in to refresh your session"))
		return
	}

	token, err := mw.JWTSignUser(*u)
	if err != nil {
		w.WriteHeader(500)
		w.Write(apiError(err.Error()))
		log.Println(err)
		return
	}

	w.Write([]byte(token))
}
