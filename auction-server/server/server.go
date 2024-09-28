package server

import (
	"log"
	"net/http"
	"os"

	"github.com/subham-singh46/auction/store"
)

type Server struct {
	Router *http.ServeMux
	store  store.Storage
}

func InitServer(store store.Storage) *Server {
	s := &Server{Router: http.NewServeMux(), store: store}
	s.Routes()
	return s
}

func (m *Server) Start(httpAddr string) error {
	log.Printf("Starting auction server at address: %s\n", httpAddr)
	port := ":" + os.Getenv("PORT")
	if port == ":" {
		port = httpAddr
	}
	if err := http.ListenAndServe(httpAddr, m.Router); err != nil {
		return err
	}
	return nil
}
