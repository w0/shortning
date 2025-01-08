package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/w0/shortning/internal/base62"
)

func (s *Server) NewShortUrl(w http.ResponseWriter, req *http.Request) {
	type Url struct {
		Url string `json:"url"`
	}

	type resUrl struct {
		Route string `json:"Route"`
	}

	d := json.NewDecoder(req.Body)
	var u Url

	err := d.Decode(&u)

	if err != nil {
		log.Printf("NewShortURL: bad decode %v", err)
		return
	}

	log.Print(u.Url)
	dbUrl, err := s.db.NewUrl(req.Context(), u.Url)

	if err != nil {
		log.Printf("NewShortUrl: %v", err)
	}

	res := resUrl{
		Route: base62.Encode(int(dbUrl.ID)),
	}

	resJson, err := json.Marshal(&res)

	if err != nil {
		http.Error(w, "Failed to marshal response.", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(resJson)

}

func (s *Server) GetShortUrl(w http.ResponseWriter, req *http.Request) {
	type reqJson struct {
		Url string `json:"url"`
	}

	decoder := json.NewDecoder(req.Body)
	data := reqJson{}

	err := decoder.Decode(&data)

	if err != nil {
		http.Error(w, "failed to decode request", http.StatusInternalServerError)
		return
	}

	url, err := url.Parse(data.Url)

	if err != nil {
		http.Error(w, "failed to parse url", http.StatusInternalServerError)
		return
	}

	id := base62.Decode(url.Path)

	dbUrl, err := s.db.GetUrl(req.Context(), int32(id))

	if err != nil {
		http.Error(w, "db error", http.StatusInternalServerError)
		return
	}

	resp := map[string]string{
		"url":       dbUrl.Url,
		"clicks":    fmt.Sprint(dbUrl.Clicks),
		"lastClick": dbUrl.UpdatedAt.Time.Format(time.RFC3339),
	}

	jsonResp, err := json.Marshal(resp)

	if err != nil {
		http.Error(w, "failed to marshal json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)

}

func (s *Server) Redirect(w http.ResponseWriter, req *http.Request) {
	id := req.PathValue("id")

	decodedId := base62.Decode(id)

	dbUrl, err := s.db.GetUrl(req.Context(), int32(decodedId))

	if err != nil {
		http.Error(w, "id doesn't exist", http.StatusNotFound)
		return
	}

	err = s.db.AddClick(req.Context(), dbUrl.ID)

	if err != nil {
		log.Printf("failed to add click: %v", err)
	}

	http.Redirect(w, req, dbUrl.Url, http.StatusSeeOther)
}
