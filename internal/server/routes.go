package server

import "net/http"

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", HelloWorld)
	mux.HandleFunc("POST /api/new", s.NewShortUrl)

	return mux
}
