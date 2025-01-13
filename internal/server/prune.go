package server

import (
	"context"
	"log"
)

func (s *Server) PruneOldLinks(days int32) {

	dbUrls, err := s.db.GetUrlsCreatedBefore(context.Background(), int32(days))

	if err != nil {
		log.Printf("PruneOldLinks: %v", err)
	}

	for i := range dbUrls {
		log.Printf("ID: %d", i)
	}
}
