package postgresDb

import (
	"database/sql"
	"os"

	"github.com/hemantsharma1498/auction/store"
	_ "github.com/lib/pq"
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
	db.Exec("CREATE TABLE IF NOT EXISTS tickets (id SERIAL PRIMARY KEY, user_id INT NOT NULL, event_date DATE NOT NULL, number_of_tickets INT NOT NULL, seat_info JSONB, price NUMERIC(10, 2) NOT NULL, auction_end TIMESTAMPTZ NOT NULL, best_offer NUMERIC(10, 2), created_at TIMESTAMPTZ DEFAULT NOW(), updated_at TIMESTAMPZ);")
	db.Exec("CREATE TABLE IF NOT EXISTS tickets (id SERIAL PRIMARY KEY, user_id INT NOT NULL, event_date DATE NOT NULL, number_of_tickets INT NOT NULL, seat_info JSONB, price NUMERIC(10, 2) NOT NULL, auction_end TIMESTAMPTZ NOT NULL, best_offer NUMERIC(10, 2), created_at TIMESTAMPTZ DEFAULT NOW(), updated_at TIMESTAMPZ);")

	return db, nil
}
