package http

import "net/http"

func (s *Server) registerURLRoutes() {
	s.router.Handle("/api/v1/shorturls", s.handleURLNew()).Methods("POST")
}

func (s *Server) handleURLNew() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		respondWithJSON(w, http.StatusOK, "OK")
	}
}
