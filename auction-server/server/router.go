package server

import (
	"net/http"

	middleware "github.com/hemantsharma1498/auction/pkg/auth-middleware"
)

func (s *Server) Routes() {
	s.Router.HandleFunc("/sign-up", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
		s.SignUp(w, r)
	})
	s.Router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
		s.Login(w, r)
	})
	s.Router.Handle("/update-password", middleware.JWTMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		s.UpdatePassword(w, r)
	})))
	s.Router.Handle("/add-ticket", middleware.JWTMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		s.UpdatePassword(w, r)
	})))
}
