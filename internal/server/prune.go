package server

import (
	"context"
	"log"
)

func (s *Server) PruneByDays(days int32) {

	dbUrls, err := s.db.GetUrlsCreatedBefore(context.Background(), int32(days))

	if err != nil {
		log.Printf("error get urls created before %d: %v", days, err)
	}

	for url := range dbUrls {
		s.db.DeleteUrl(context.Background(), int32(url))
	}
}
