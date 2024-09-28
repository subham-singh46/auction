package store

import "github.com/subham-singh46/auction/store/models"

type Storage interface {
	CreateUser(*models.User) (int, error)
	GetUsersByIds([]int) ([]*models.User, error)
	UpdatePassword(email, salt, pwHash string) error
	GetUsersByEmail([]string) ([]*models.User, error)
	AddTicket(*models.Ticket) (int, error)
	GetAllTickets(int, int) ([]*models.Ticket, error)
	GetTicketById(int) (*models.Ticket, error)
	GetTicketsByUserId(int) ([]*models.Ticket, error)
	AddNewBid(int, int, int, int) (int, error)
	GetUserBids(int) ([]*models.Bid, error)
}

type Connecter interface {
	Connect() (Storage, error)
}
