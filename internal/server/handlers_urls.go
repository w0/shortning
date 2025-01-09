package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
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
		respondWithError(w, http.StatusBadRequest, "json decode error", err)
		return
	}

	dbUrl, err := s.db.NewUrl(req.Context(), u.Url)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to create short url", err)
	}

	res := resUrl{
		Route: base62.Encode(int(dbUrl.ID)),
	}

	respondWithJson(w, http.StatusCreated, res)
}

func (s *Server) GetShortUrl(w http.ResponseWriter, req *http.Request) {
	type reqJson struct {
		Url string `json:"url"`
	}

	decoder := json.NewDecoder(req.Body)
	data := reqJson{}

	err := decoder.Decode(&data)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "json decode error", err)
		return
	}

	url, err := url.Parse(data.Url)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "url invalid", err)
		return
	}

	id := base62.Decode(url.Path)

	dbUrl, err := s.db.GetUrl(req.Context(), int32(id))

	if err != nil {
		respondWithError(w, http.StatusNotFound, "entry not found", err)
		return
	}

	res := map[string]string{
		"url":       dbUrl.Url,
		"clicks":    fmt.Sprint(dbUrl.Clicks),
		"lastClick": dbUrl.UpdatedAt.Time.Format(time.RFC3339),
	}

	respondWithJson(w, http.StatusOK, res)
}

func (s *Server) Redirect(w http.ResponseWriter, req *http.Request) {
	id := req.PathValue("id")

	decodedId := base62.Decode(id)

	dbUrl, err := s.db.GetUrl(req.Context(), int32(decodedId))

	if err != nil {
		lp := filepath.Join("./web/template", "layout.html")
		fp := filepath.Join("./web/template", "404.html")

		tmpl, _ := template.ParseFiles(lp, fp)
		w.WriteHeader(http.StatusNotFound)
		tmpl.ExecuteTemplate(w, "layout", nil)
		return
	}

	err = s.db.AddClick(req.Context(), dbUrl.ID)

	if err != nil {
		log.Printf("failed to add click: %v", err)
	}

	http.Redirect(w, req, dbUrl.Url, http.StatusSeeOther)
}
