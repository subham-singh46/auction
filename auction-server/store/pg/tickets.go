package postgresDb

import (
	"encoding/json"
	"fmt"

	"github.com/hemantsharma1498/auction/store/models"
)

func (pg *PostgresDb) AddTicket(ticket *models.Ticket) (int, error) {
	seatInfoJson, err := json.Marshal(ticket.SeatInfo)
	fmt.Printf("%v", ticket)
	query :=
		`INSERT INTO tickets
            (user_id, event_date, venue, number_of_tickets, seat_info, price, best_offer, deadline)
            VALUES
            ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id;`
	var ticketId int
	err = pg.db.QueryRow(query, ticket.UserID, ticket.EventDate, ticket.Venue, ticket.NumberOfTickets, seatInfoJson, ticket.Price, ticket.BestOffer, ticket.Deadline).Scan(&ticketId)
	if err != nil {
		return -1, err
	}
	return ticketId, nil
}

func (pg *PostgresDb) GetAllTickets(limit, offset int) ([]*models.Ticket, error) {
	res := make([]*models.Ticket, 0)
	query :=
		`SELECT 
            id, venue, user_id, event_date, number_of_tickets, seat_info, price, best_offer, deadline
        FROM tickets
        LIMIT $1 OFFSET $2`

	rows, err := pg.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		ticket := &models.Ticket{}
		seatInfoJson := make([]byte, 0)
		err := rows.Scan(
			&ticket.TicketID,
			&ticket.Venue,
			&ticket.UserID,
			&ticket.EventDate,
			&ticket.NumberOfTickets,
			&seatInfoJson, // SeatInfo will be stored as a JSONB field (as []byte)
			&ticket.Price,
			&ticket.BestOffer,
			&ticket.Deadline,
		)
		if err != nil {
			return nil, err
		}
		if err = json.Unmarshal(seatInfoJson, &ticket.SeatInfo); err != nil {
			return nil, err
		}
		res = append(res, ticket)
	}
	return res, nil
}

func (pg *PostgresDb) GetTicketById(ticketID int) (*models.Ticket, error) {
	query := `SELECT id, user_id, event_date, number_of_tickets, seat_info, price, best_offer, deadline, created_at, updated_at 
			  FROM tickets WHERE id = $1`

	ticket := &models.Ticket{}
	var seatInfoJSON []byte
	// Querying the ticket and seat_info field
	err := pg.db.QueryRow(query, ticketID).Scan(
		&ticket.TicketID,
		&ticket.UserID,
		&ticket.EventDate,
		&ticket.NumberOfTickets,
		&seatInfoJSON, // SeatInfo will be stored as a JSONB field (as []byte)
		&ticket.Price,
		&ticket.BestOffer,
		&ticket.Deadline,
		&ticket.CreatedAt,
		&ticket.UpdatedAt,
	)
	if err != nil {
		return ticket, err
	}

	// Unmarshal the seat_info JSONB data into the SeatInfo struct
	if err := json.Unmarshal(seatInfoJSON, &ticket.SeatInfo); err != nil {
		return ticket, err
	}

	return ticket, nil
}

func (pg *PostgresDb) GetTicketsByUserId(ticketID int) ([]*models.Ticket, error) {
	res := make([]*models.Ticket, 0)
	query := `SELECT id, user_id, event_date, venue, number_of_tickets, seat_info, price, best_offer, deadline, created_at, updated_at 
			  FROM tickets WHERE user_id = $1`

	// Querying the ticket and seat_info field
	rows, err := pg.db.Query(query, ticketID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		ticket := &models.Ticket{}
		var seatInfoJSON []byte
		rows.Scan(
			&ticket.TicketID,
			&ticket.UserID,
			&ticket.EventDate,
			&ticket.Venue,
			&ticket.NumberOfTickets,
			&seatInfoJSON, // SeatInfo will be stored as a JSONB field (as []byte)
			&ticket.Price,
			&ticket.BestOffer,
			&ticket.Deadline,
			&ticket.CreatedAt,
			&ticket.UpdatedAt,
		)
		// Unmarshal the seat_info JSONB data into the SeatInfo struct
		if err := json.Unmarshal(seatInfoJSON, &ticket.SeatInfo); err != nil {
			return nil, err
		}
		res = append(res, ticket)
	}

	return res, nil
}
