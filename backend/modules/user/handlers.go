package user

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
)

type UserSignupDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Signup(w http.ResponseWriter, r *http.Request) {
	payload, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Please send a valid json."))
		return
	}

	var user UserSignupDto
	ctx := r.Context()

	err = json.Unmarshal(payload, &user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Please send a valid json: email and password fields"))
		return
	}

	existentUser, err := findUserByEmail(ctx, user.Email)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	if existentUser != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User is already registered. Please sign in."))
		return
	}

	// Hash the password
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	err = insertUser(ctx, user.Email, string(password))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to create user's account."))
		return
	}

	w.WriteHeader(http.StatusCreated)
}

type UserSignInDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	payload, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Please send a valid json."))
		return
	}

	var userAuthData UserSignInDto

	err = json.Unmarshal(payload, &userAuthData)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Please send a valid json: email and password fields."))
		return
	}

	user, err := findUserByEmail(r.Context(), userAuthData.Email)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	if user == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found."))
		return
	}

	// TODO
}
