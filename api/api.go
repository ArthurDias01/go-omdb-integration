package api

import (
	"encoding/json"
	"log/slog"
	"math/rand/v2"
	"net/http"
	"net/url"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type PostBody struct {
	URL string `json:"url"`
}

type Response struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func sendJSON(w http.ResponseWriter, resp Response, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(resp)
	if err != nil {
		slog.Error("Error marshalling response:", "error", err)
		sendJSON(w, Response{Error: "Internal server error"}, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(statusCode)
	if _, err := w.Write(data); err != nil {
		// fmt.Println("Error writing response:", err)
		slog.Error("Error writing response:", "error", err)
		return
	}
}

func NewHandler(db map[string]string) http.Handler {
	r := chi.NewMux()
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Post("/api/shorten", handlePost(db))
	r.Get("/{code}", handleGet(db))

	return r
}

func handlePost(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body PostBody
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			sendJSON(w, Response{Error: "Invalid request"}, http.StatusUnprocessableEntity)
			return
		}
		_, err := url.Parse(body.URL)
		if err != nil {
			sendJSON(w, Response{Error: "Invalid URL"}, http.StatusBadRequest)
			return
		}
		code := generateCode(db)
		db[code] = body.URL
		sendJSON(w, Response{Data: code}, http.StatusCreated)
	}
}

const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateCode(db map[string]string) string {
	const n = 8
	for {
		bytes := make([]byte, n)
		for i := range n {
			bytes[i] = characters[rand.IntN(len(characters))]
		}
		code := string(bytes)
		if _, exists := db[code]; !exists {
			return code
		}
	}
}

func handleGet(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "code")
		url, ok := db[code]
		if !ok {
			sendJSON(w, Response{Error: "Not found"}, http.StatusNotFound)
			return
		}
		http.Redirect(w, r, url, http.StatusPermanentRedirect)
	}
}
