package http

import (
	"encoding/json"
	"fmt"
	"github.com/alexander231/url-short/base62"
	"github.com/gorilla/mux"
	"io"
	"math/rand"
	"net/http"
	"time"
)

const shortURLVar = "shorturl"

func (s *Server) registerURLRoutes() {
	s.router.Handle("/{shorturl}", s.handleURLRedirect()).Methods("GET")

	s.router.Handle("/api/v1/shorturls", s.handleURLCreate()).Methods("POST")

	s.router.Handle("/api/v1/shorturls/{shorturl}", s.handleURLView()).Methods("GET")

}

func (s *Server) handleURLRedirect() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shortURL := mux.Vars(r)[shortURLVar]
		URL, err := s.URLService.Get(shortURL)

		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		http.Redirect(w, r, URL, http.StatusFound)
	}
}

func (s *Server) handleURLView() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shortURL := mux.Vars(r)[shortURLVar]
		fmt.Println(shortURL)
		URL, err := s.URLService.Get(shortURL)
		fmt.Println(URL)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondWithJSON(w, http.StatusOK, URL)
		return
	}
}

func (s *Server) handleURLCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		defer r.Body.Close()

		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		var urlReq URLRequest
		if err = json.Unmarshal(body, &urlReq); err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return

		}
		URL := urlReq.URL
		if URL == "" {
			respondWithError(w, http.StatusBadRequest, "URL can't be empty")
			return
		}
		shortURL := urlReq.ShortURL
		if shortURL == "" {
			shortURL = base62.Encode(rand.Uint64())
		}
		if err = s.URLService.Set(shortURL, URL, 60*time.Second); err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		fmt.Println(shortURL, URL)
		respondWithJSON(w, http.StatusOK, "OK")
		return
	}
}

type URLRequest struct {
	ShortURL string `json:"shortURL"`
	URL      string `json:"URL"`
}
