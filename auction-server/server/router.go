package server

import (
	"net/http"

	middleware "github.com/subham-singh46/auction/pkg/auth-middleware"
)

/*
*ConfirmSale
 */

func (s *Server) Routes() {
	s.Router.HandleFunc("/sign-up", func(w http.ResponseWriter, r *http.Request) {
		enableCors(w, r)
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
		s.SignUp(w, r)
	})
	s.Router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		enableCors(w, r)
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
		s.Login(w, r)
	})
	s.Router.Handle("/update-password", middleware.JWTMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		enableCors(w, r)
		if r.Method != http.MethodPut {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		s.UpdatePassword(w, r)
	})))
	s.Router.Handle("/add-ticket", middleware.JWTMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		enableCors(w, r)
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		s.AddTicket(w, r)
	})))
	s.Router.Handle("/get-all-tickets", middleware.JWTMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		enableCors(w, r)
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		s.GetAllTickets(w, r)
	})))
	s.Router.Handle("/get-user-listing", middleware.JWTMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		s.GetUserListing(w, r)
	})))
	s.Router.Handle("/place-bid", middleware.JWTMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		s.AddNewBid(w, r)
	})))
	s.Router.Handle("/get-user-bids", middleware.JWTMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		s.GetUserBids(w, r)
	})))
}

func enableCors(w http.ResponseWriter, r *http.Request) {
	// Set the CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Origin, Accept, Authorization")

	// If the method is OPTIONS, respond with 200 OK and return
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
}

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set the CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Origin, Accept, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true") // Optional, if you allow credentials like cookies or authorization headers

		// If the method is OPTIONS, respond with 200 OK and return
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Proceed to the next middleware or handler
		next.ServeHTTP(w, r)
	})
}
