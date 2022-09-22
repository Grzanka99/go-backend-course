package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/grzanka99/backend-go/internal/services/user"
)

func DefineRoutes() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/user/login", user.UserLogin).Methods(http.MethodPost)
	router.HandleFunc("/user/register", user.UserRegister).Methods(http.MethodPost)

	return router
}
