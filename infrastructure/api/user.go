package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/charly3pins/eShop/application"
)

type UserHandler struct {
	UserService application.UserService
}

type SignInUserBody struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func (uh UserHandler) SignInUser(w http.ResponseWriter, r *http.Request) {
	body := &SignInUserBody{}
	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	req := application.SignInUserRequest{
		Email: body.Email,
		Name:  body.Name,
	}
	if _, err := uh.UserService.SignInUser(req); err != nil {
		log.Printf("error signing in a new user: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
