package application

import (
	"fmt"
	"net/http"

	"github.com/grzanka99/backend-go/internal/controller"
)

var PORT int16 = 3000

func Run() error {
	println("Starting server on port: ", PORT)

	router := controller.DefineRoutes()

	err := http.ListenAndServe(":3000", router)

	if err != nil {
		fmt.Printf("Unexpected error: %s", err)
	}

	return nil
}
