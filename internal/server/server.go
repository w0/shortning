package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	_ "github.com/joho/godotenv/autoload"
	"github.com/w0/shortning/internal/database"
)

type Server struct {
	port string
	db   *database.Queries
}

func NewServer() *http.Server {
	port := os.Getenv("PORT")
	conn, err := pgx.Connect(context.Background(), os.Getenv("GOOSE_DBSTRING"))

	if err != nil {
		log.Fatal("Failed to open db connection")
	}

	NewServer := &Server{
		port: port,
		db:   database.New(conn),
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
