package server

import "net/http"

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/v1/shorten", s.NewShortUrl)
	mux.HandleFunc("GET /api/v1/shorten", s.GetShortUrl)
	mux.HandleFunc("GET /{id}", s.Redirect)

	return mux
}
