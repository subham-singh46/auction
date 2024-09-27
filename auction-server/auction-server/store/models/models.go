package models

import (
	"time"
)

type User struct {
	UserID    int       `json:"user_id" db:"user_id"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	Mobile    string    `json:"mobile" db:"mobile"`
	Salt      string    `json:"salt" db:"salt"`
	PwHash    string    `json:"pw_hash" db:"pw_hash"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

type Ticket struct {
	TicketID        int         `json:"ticketId" db:"ticket_id"`   // Ticket ID (Primary Key)
	UserID          int         `json:"userId" db:"user_id"`       // The user who is selling the ticket
	EventDate       time.Time   `json:"eventDate" db:"event_date"` // The date of the event
	Venue           string      `json:"venue" db:"venue"`
	NumberOfTickets int         `json:"numOfTickets" db:"number_of_tickets"` // Number of tickets available for sale
	SeatInfo        []*SeatInfo `json:"seatInfo" db:"seat_info"`             // JSONB field to store seat info (e.g., seats, block, level)
	Price           float64     `json:"price" db:"price"`                    // Ticket price
	BestOffer       float64     `json:"bestOffer" db:"best_offer"`           // Best offer (can be NULL, so we use a pointer)
	Deadline        time.Time   `json:"deadline" db:"auction_end"`           // Auction end timestamp
	CreatedAt       time.Time   `json:"createdAt" db:"created_at"`           // Record creation timestamp
	UpdatedAt       time.Time   `json:"updatedAt" db:"updated_at"`
}

type SeatInfo struct {
	SeatNumber int    `json:"seatNumber"`
	Block      string `json:"block"`
	Level      int    `json:"level"`
}

type Bid struct {
	BidId         int       `json:"BidId"`
	TicketId      int       `json:"ticketId"`
	BidPrice      int       `json:"bidPrice"`
	OriginalPrice int       `json:"originalPrice"`
	Venue         string    `json:"venue"`
	OwnerId       int       `json:"ownerId"`
	BidderId      int       `json:"bidderId"`
	CreatedAt     time.Time `json:"createdAt" db:"created_at"` // Record creation timestamp
	UpdatedAt     time.Time `json:"updatedAt" db:"updated_at"`
}
