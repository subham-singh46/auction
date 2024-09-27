package server

import (
	"net/http"

	"github.com/hemantsharma1498/auction/pkg/utils"
	"github.com/hemantsharma1498/auction/store/models"
)

func (s *Server) AddTicket(w http.ResponseWriter, r *http.Request) {
	d := &AddTicketReq{}
	if err := utils.DecodeReqBody(r, d); err != nil {
		utils.WriteResponse(w, err, "Encountered an error. Please try again", http.StatusInternalServerError)
		return
	}
	userId := r.Context().Value("UserID").(int)
	_, err := s.store.GetUsersByIds([]int{userId})
	if err != nil {
		utils.WriteResponse(w, err, "unable to find user against given userId", http.StatusBadRequest)
		return
	}
	seatInfo := make([]*models.SeatInfo, 0)
	for _, seat := range d.SeatInfo {
		seat := &models.SeatInfo{
			SeatNumber: seat.SeatNumber,
			Block:      seat.Block,
			Level:      seat.Level,
		}
		seatInfo = append(seatInfo, seat)
	}
	eventDate, err := utils.IsoDateToTime(d.EventDate)
	if err != nil {
		utils.WriteResponse(w, err, "incorrect date format", http.StatusBadRequest)
		return
	}
	deadline, err := utils.IsoDateToTime(d.Deadline)
	if err != nil {
		utils.WriteResponse(w, err, "incorrect date format", http.StatusBadRequest)
		return
	}
	ticket := &models.Ticket{
		UserID:          userId,
		NumberOfTickets: d.NumberOfTickets,
		EventDate:       eventDate,
		Venue:           d.Venue,
		SeatInfo:        seatInfo,
		Price:           float64(d.Price),
		BestOffer:       float64(d.Price),
		Deadline:        deadline,
	}

	ticketId, err := s.store.AddTicket(ticket)
	if err != nil {
		utils.WriteResponse(w, err, "Encountered an error. Please try again", http.StatusInternalServerError)
		return
	}
	res := &AddTicketRes{TicketId: ticketId}
	utils.WriteResponse(w, nil, res, http.StatusCreated)
	return
}

func (s *Server) GetAllTickets(w http.ResponseWriter, r *http.Request) {
	d := &GetAllTicketsReq{}
	if err := utils.DecodeReqBody(r, d); err != nil {
		utils.WriteResponse(w, err, "Error occurred. Please try again later", http.StatusInternalServerError)
		return
	}
	tickets, err := s.store.GetAllTickets(d.Limit, d.Offset)
	if err != nil {
		utils.WriteResponse(w, err, "Error occurred. Please try again later", http.StatusInternalServerError)
		return
	}

	res := &GetAllTicketsRes{}

	for _, t := range tickets {
		users, err := s.store.GetUsersByIds([]int{t.UserID})
		if err != nil {
			utils.WriteResponse(w, err, "Error occurred. Please try again later", http.StatusInternalServerError)
			return
		}
		seats := make([]Seat, 0)
		for _, s := range t.SeatInfo {
			seatInfo := Seat{
				SeatNumber: s.SeatNumber,
				Block:      s.Block,
				Level:      s.Level,
			}
			seats = append(seats, seatInfo)
		}

		ticket := &Ticket{
			TicketID:        t.TicketID,
			EventDate:       t.EventDate.String(),
			UserID:          t.UserID,
			Venue:           t.Venue,
			NumberOfTickets: t.NumberOfTickets,
			HighestBid:      int(t.BestOffer),
			Price:           int(t.Price),
			SeatInfo:        seats,
			Deadline:        t.Deadline.String(),
			ListedBy:        users[0].Name,
		}
		res.Tickets = append(res.Tickets, ticket)
	}

	utils.WriteResponse(w, nil, res, http.StatusOK)
	return

}

func (s *Server) GetUserListing(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("UserID").(int)
	tickets, err := s.store.GetTicketsByUserId(userId)
	if err != nil {
		utils.WriteResponse(w, err, "Error occurred. Please try again later", http.StatusInternalServerError)
		return
	}

	res := &GetAllTicketsRes{}

	for _, t := range tickets {
		users, _ := s.store.GetUsersByIds([]int{t.UserID})
		seats := make([]Seat, 0)
		for _, s := range t.SeatInfo {
			seatInfo := Seat{
				SeatNumber: s.SeatNumber,
				Block:      s.Block,
				Level:      s.Level,
			}
			seats = append(seats, seatInfo)
		}
		ticket := &Ticket{
			TicketID:        t.TicketID,
			EventDate:       t.EventDate.String(),
			Venue:           t.Venue,
			UserID:          t.UserID,
			NumberOfTickets: t.NumberOfTickets,
			HighestBid:      int(t.BestOffer),
			Price:           int(t.Price),
			SeatInfo:        seats,
			Deadline:        t.Deadline.String(),
			ListedBy:        users[0].Name,
		}
		res.Tickets = append(res.Tickets, ticket)
	}

	utils.WriteResponse(w, nil, res, http.StatusOK)
	return
}
