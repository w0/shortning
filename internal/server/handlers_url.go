package server

import (
	"encoding/json"
	"log"
	"net/http"
)

type Url struct {
	Url string `json:"url"`
}

func (s *Server) NewShortUrl(w http.ResponseWriter, req *http.Request) {
	d := json.NewDecoder(req.Body)
	var u Url

	err := d.Decode(&u)

	if err != nil {
		log.Printf("NewShortURL: bad decode %v", err)
		return
	}

	log.Print(u.Url)
	_, err = s.db.NewUrl(req.Context(), u.Url)

	if err != nil {
		log.Printf("NewShortUrl: %v", err)
	}

}

func HelloWorld(w http.ResponseWriter, req *http.Request) {
	msg := map[string]string{"hello": "world"}
	jsonRes, err := json.Marshal(msg)

	if err != nil {
		http.Error(w, "Failed to marshal response.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(jsonRes); err != nil {
		log.Printf("Failed to write response: %v", err)
	}

}
