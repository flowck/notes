package http

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"notes/internal/common"
	"notes/internal/user/service"
)

type SignupDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserHttpHandler struct {
	service *service.UserService
}

func NewUserHttpHandler(service *service.UserService) UserHttpHandler {
	if service == nil {
		panic("user service is missing.")
	}

	return UserHttpHandler{service}
}

func (h *UserHttpHandler) Signup(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		common.SendResponse(w, "Please send a valid json.", http.StatusBadRequest)
		return
	}

	var authPayload AuthPayload
	err = json.Unmarshal(body, &authPayload)

	if err != nil {
		common.SendResponse(w, "Please send a valid json: email and password fields", http.StatusBadRequest)
		return
	}

	err = h.service.SignUp(r.Context(), authPayload.Email, authPayload.Password)

	if err != nil {
		common.SendResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *UserHttpHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	payload, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Please send a valid json."))
		return
	}

	var userAuthData AuthPayload

	err = json.Unmarshal(payload, &userAuthData)

	if err != nil {
		common.SendResponse(w, "Please send a valid json: email and password fields.", http.StatusBadRequest)
		return
	}

	token, err := h.service.SignIn(r.Context(), userAuthData.Email, userAuthData.Password)

	if err != nil {
		log.Printf("Unable to find run query to find user: %s", err)
		common.SendResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	common.WriteJSON(w, struct {
		Token string `json:"token"`
	}{Token: token}, 200)
}
