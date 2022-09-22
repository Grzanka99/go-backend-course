package main

import (
	"log"
	"os"

	"github.com/grzanka99/backend-go/internal/application"
)

func main() {
	// err := run()
	var err error = application.Run()

	if err != nil {
		log.Println("Fatal error, application couldn't start", err)
		os.Exit(-1)
	}
}
