package storage

import "github.com/grzanka99/backend-go/internal/services/user"

type Database struct {
	User []user.UserInformation
}
