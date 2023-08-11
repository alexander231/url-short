package http

import (
	"encoding/json"
	shorturl "github.com/alexander231/url-short"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	router *mux.Router

	Addr string

	URLService shorturl.URLService
}

func NewServer(service shorturl.URLService) *Server {
	s := &Server{router: mux.NewRouter(),
		URLService: service}

	s.registerURLRoutes()

	return s
}

func (s *Server) ListenAndServe() error {
	return http.ListenAndServe(s.Addr, s.router)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	res, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	w.Write(res)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}
