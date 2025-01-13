package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
	_ "github.com/joho/godotenv/autoload"
	"github.com/w0/shortning/internal/database"
)

type Server struct {
	port      string
	db        *database.Queries
	pruneDays int32
}

func NewServer() *http.Server {
	port := os.Getenv("PORT")
	conn, err := pgx.Connect(context.Background(), os.Getenv("GOOSE_DBSTRING"))

	if err != nil {
		log.Fatal("Failed to open db connection")
	}

	pruneDays, err := strconv.Atoi(os.Getenv("PRUNE_DAYS"))

	if err != nil {
		log.Printf("[WARNING] %v", err)
	}

	NewServer := &Server{
		port:      port,
		db:        database.New(conn),
		pruneDays: int32(pruneDays),
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	go dbBackgroundTasks(*NewServer, 24*time.Second)

	return server
}

func respondWithError(w http.ResponseWriter, status int, msg string, err error) {
	if err != nil {
		log.Println(msg)
	}

	type errRes struct {
		Error string `json:"error"`
	}

	e := errRes{
		Error: fmt.Sprintf("%s: %v", msg, err),
	}

	respondWithJson(w, status, e)
}

func respondWithJson(w http.ResponseWriter, status int, res interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json, err := json.Marshal(res)

	if err != nil {
		log.Printf("respondWithJson: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(json)
}

func dbBackgroundTasks(s Server, interval time.Duration) {
	ticker := time.NewTicker(interval)

	done := make(chan bool)

	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			s.PruneOldLinks(s.pruneDays)
		}
	}
}
