package postgresDb

import (
	"fmt"

	"github.com/subham-singh46/auction/store/models"
)

func (pg *PostgresDb) AddNewBid(bidPrice, ownerId, ticketId, bidderId int) (int, error) {
	query :=
		`INSERT INTO bids(owner_id, ticket_id, bidder_id, bid_price) VALUES($1, $2, $3, $4)`

	var bidId int
	tx, err := pg.db.Begin()
	if err != nil {
		return -1, err
	}
	fmt.Println(bidPrice, ownerId, ticketId, bidderId)
	err = tx.QueryRow(query, ownerId, ticketId, bidderId, bidPrice).Scan(&bidId)
	_, err = tx.Exec("UPDATE tickets set best_offer = $1 where id = $2", bidPrice, ticketId)
	if err != nil {
		return -1, err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	return bidId, nil
}

func (pg *PostgresDb) GetUserBids(userId int) ([]*models.Bid, error) {
	res := make([]*models.Bid, 0)

	query :=
		`SELECT bid_id, owner_id, ticket_id, bidder_id, bid_price, og_price, venue, created_at FROM bids WHERE bidder_id = $1`
	rows, err := pg.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		bid := &models.Bid{}
		err := rows.Scan(&bid.BidId, &bid.OwnerId, &bid.TicketId, &bid.BidderId, &bid.BidPrice, &bid.OriginalPrice, &bid.Venue, &bid.CreatedAt)
		if err != nil {
			return nil, err
		}
		res = append(res, bid)
	}

	return res, nil
}
