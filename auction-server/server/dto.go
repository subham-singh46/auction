package server

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type UpdatePasswordReq struct {
	Email       string `json:"email"`
	NewPassword string `json:"newPassword"`
}

type AddTicketReq struct {
	ConcertDate     string `json:"concertDate"`
	NumberOfTickets int    `json:"numberOfTickets"`
	Seats           []Seat `json:"details"`
	Deadline        string `json:"deadline"`
}

type Seat struct {
	SeatNumber int    `json:"seatNumber"`
	Block      string `json:"block"`
	Level      string `json:"level"`
}
