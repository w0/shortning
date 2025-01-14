package server

import (
	"context"
	"log"

	"github.com/w0/shortning/internal/database"
)

func (s *Server) PruneByDays(days int32) {

	dbUrls, err := s.db.GetUrlsCreatedBefore(context.Background(), int32(days))

	if err != nil {
		log.Printf("error getting urls created before %d: %v", days, err)
		return
	}

	for url := range dbUrls {
		s.db.DeleteUrl(context.Background(), int32(url))
	}
}

func (s *Server) PruneByClick(clicks int32, days int32) {
	dbUrls, err := s.db.GetUrlsUnderClickCount(context.Background(), database.GetUrlsUnderClickCountParams{
		Clicks: clicks,
		Days:   days,
	})

	if err != nil {
		log.Printf("error getting urls by click: %v", err)
		return
	}

	for url := range dbUrls {
		s.db.DeleteUrl(context.Background(), int32(url))
	}
}
