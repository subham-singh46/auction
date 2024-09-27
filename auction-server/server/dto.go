package server

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRes struct {
	UserID int    `json:"userId"`
	Token  string `json:"token"`
}

type SignUpReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type SignUpRes struct {
	UserID int `json:"userId"`
}

type UpdatePasswordReq struct {
	Email       string `json:"email"`
	NewPassword string `json:"newPassword"`
}

type AddTicketReq struct {
	EventDate       string `json:"eventDate"`
	Venue           string `json:"Venue"`
	NumberOfTickets int    `json:"numberOfTickets"`
	Price           int    `json:"price"`
	SeatInfo        []Seat `json:"seatInfo"`
	Deadline        string `json:"deadline"`
}

type Seat struct {
	SeatNumber int    `json:"seatNumber"`
	Block      string `json:"block"`
	Level      int    `json:"level"`
}

type AddTicketRes struct {
	TicketId int `json:"ticketId"`
}

type GetAllTicketsReq struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type GetAllTicketsRes struct {
	Tickets []*Ticket
}

type Ticket struct {
	EventDate       string `json:"eventDate"`
	Venue           string `json:"Venue"`
	NumberOfTickets int    `json:"numberOfTickets"`
	Price           int    `json:"price"`
	SeatInfo        []Seat `json:"seatInfo"`
	Deadline        string `json:"deadline"`
	ListedBy        string `json:"listedBy"`
}
