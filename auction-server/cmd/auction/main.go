package main

import (
	"log"

	"github.com/subham-singh46/auction/server"
	postgresDb "github.com/subham-singh46/auction/store/pg"
)

const httpAddress = ":3100"

func main() {
	log.Printf("Initialising members server")

	log.Printf("Connecting to database...")

	store, err := postgresDb.NewAuctionDbConnector().Connect()
	if err != nil {
		log.Panicf("Unable to connect to db, error: %s\n", err)
	}
	log.Printf("Db connection established")

	s := server.InitServer(store)
	if err = s.Start(httpAddress); err != nil {
		log.Panicf("Failed to initialise server at %s, error: %s\n", httpAddress, err)
	}
}
