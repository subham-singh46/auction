package store

import "github.com/hemantsharma1498/auction/store/models"

type Storage interface {
	CreateUser(*models.User) (int, error)
	GetUsersByIds([]int) ([]*models.User, error)
	GetUsersByEmail([]string) ([]*models.User, error)
	UpdatePassword(email, salt, pwHash string) error
	AddTicket(*models.Ticket) (int, error)
	GetAllTickets(int, int) ([]*models.Ticket, error)
	GetTicketById(int) (*models.Ticket, error)
	GetTicketsByUserId(int) ([]*models.Ticket, error)
}

type Connecter interface {
	Connect() (Storage, error)
}
