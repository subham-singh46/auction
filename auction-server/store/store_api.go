package store

import "github.com/hemantsharma1498/auction/store/models"

type Storage interface {
	CreateUser(*models.User) error
	GetUsersByEmail([]string) ([]*models.User, error)
	UpdatePassword(email, salt, pwHash string) error
}

type Connecter interface {
	Connect() (Storage, error)
}
