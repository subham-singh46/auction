package server

import "net/http"

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
}
