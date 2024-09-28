package postgresDb

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/subham-singh46/auction/store"
)

var dsn = os.Getenv("DATABASE_URL")

func NewAuctionDbConnector() store.Connecter {
	return &PostgresDb{}
}

type PostgresDb struct {
	db *sql.DB
}

func (pg *PostgresDb) Connect() (store.Storage, error) {
	if pg.db == nil {
		var err error
		pg.db, err = initDb()
		if err != nil {
			return nil, err
		}
		return pg, nil
	}
	return pg, nil
}

func initDb() (*sql.DB, error) {

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	// @TODO Refactor to migration script
	db.Exec("CREATE TABLE IF NOT EXISTS users (user_id SERIAL PRIMARY KEY, name TEXT NOT NULL, salt TEXT NOT NULL, pw_hash TEXT NOT NULL, email TEXT NOT NULL UNIQUE, mobile TEXT NOT NULL UNIQUE, created_at TIMESTAMPTZ DEFAULT NOW());")
	query :=
		`
-- Function to set timestamps on insert and update
CREATE OR REPLACE FUNCTION set_timestamps()
RETURNS TRIGGER AS $$
BEGIN
    IF TG_OP = 'INSERT' THEN
        NEW.created_at = NOW();
    END IF;
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create tickets table and related objects
CREATE SEQUENCE IF NOT EXISTS tickets_id_seq;

CREATE TABLE IF NOT EXISTS tickets (
    id INTEGER NOT NULL DEFAULT nextval('tickets_id_seq'),
    user_id INTEGER NOT NULL,
    event_date DATE NOT NULL,
    venue TEXT NOT NULL,
    number_of_tickets INTEGER NOT NULL,
    seat_info JSONB,
    price NUMERIC(10, 2) NOT NULL,
    best_offer NUMERIC(10, 2),
    deadline TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE,
    PRIMARY KEY (id),
    CONSTRAINT tickets_number_of_tickets_check CHECK (number_of_tickets > 0)
);

CREATE TRIGGER set_tickets_timestamps
BEFORE INSERT OR UPDATE ON tickets
FOR EACH ROW
EXECUTE FUNCTION set_timestamps();

-- Create bids table and related objects
CREATE SEQUENCE IF NOT EXISTS bids_bid_id_seq;

CREATE TABLE IF NOT EXISTS bids (
    bid_id INTEGER NOT NULL DEFAULT nextval('bids_bid_id_seq'),
    bidder_id INTEGER NOT NULL,
    ticket_id INTEGER NOT NULL,
    owner_id INTEGER NOT NULL,
    bid_price INTEGER NOT NULL,
    og_price INTEGER NOT NULL,
    venue TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE,
    PRIMARY KEY (bid_id)
);

CREATE TRIGGER set_bids_timestamps
BEFORE INSERT OR UPDATE ON bids
FOR EACH ROW
EXECUTE FUNCTION set_timestamps();

-- Add foreign key constraints
ALTER TABLE bids
ADD CONSTRAINT bids_bidder_id_fkey FOREIGN KEY (bidder_id) REFERENCES users(user_id),
ADD CONSTRAINT bids_owner_id_fkey FOREIGN KEY (owner_id) REFERENCES users(user_id),
ADD CONSTRAINT bids_ticket_id_fkey FOREIGN KEY (ticket_id) REFERENCES tickets(id);
        `

	_, err = db.Exec(query)
	if err != nil {
		fmt.Println(err)
	}
	return db, nil
}
