package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func BuildRouter(uh UserHandler, oh OrderHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/users", uh.SignInUser).Methods("POST")
	r.HandleFunc("/orders", oh.AddProductsToOrder).Methods("POST")

	http.Handle("/", r)
	return r
}
