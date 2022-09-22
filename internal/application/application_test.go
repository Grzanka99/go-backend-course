package application_test

import (
	"testing"

	"github.com/grzanka99/backend-go/internal/application"
)

func TestRun(t *testing.T) {
	err := application.Run()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
}
