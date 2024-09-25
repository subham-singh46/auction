package postgresDb

import (
	"database/sql"
	"os"

	"github.com/hemantsharma1498/auction/store"
	_ "github.com/lib/pq"
)

var dsn = os.Getenv("Dsn")

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

	return db, nil
}
