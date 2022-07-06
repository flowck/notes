package user

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
	"net/http"
	"notes/infra"
	"time"
)

type SignupDto struct {
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

	var user SignupDto
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

type SignInDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CustomClaims struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	jwt.RegisteredClaims
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	payload, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Please send a valid json."))
		return
	}

	var userAuthData SignInDto

	err = json.Unmarshal(payload, &userAuthData)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Please send a valid json: email and password fields."))
		return
	}

	user, err := findUserByEmail(r.Context(), userAuthData.Email)

	if err != nil {
		log.Printf("Unable to find run query to find user: %s", err)
		infra.SendResponse(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if user == nil {
		infra.SendResponse(w, "User not found.", http.StatusNotFound)
		return
	}

	// Compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userAuthData.Password))

	if err != nil {
		infra.SendResponse(w, "User's password do not match.", http.StatusUnauthorized)
		return
	}

	// TODO
	// Generate JWT token

	// defaultClaims :=

	claims := CustomClaims{
		user.FirstName,
		user.LastName,
		user.Email,
		jwt.RegisteredClaims{
			Subject:   "",
			ID:        user.Id,
			Issuer:    "NOTES_SERVICE",
			Audience:  []string{"default"},
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 5 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, signErr := token.SignedString([]byte(infra.Cfg.JwtSigningKey))

	if signErr != nil {
		log.Printf("Could not sign the JWT: %s", signErr)
		infra.SendResponse(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	fmt.Println("Token", ss)

	res, _ := json.Marshal(struct {
		Token string `json:"token"`
	}{Token: ss})

	w.Write(res)
}
