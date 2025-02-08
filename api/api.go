package api

import (
	"encoding/json"
	"go-first-big-project/api/omdb"
	"log/slog"
	"net/http"

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

func NewHandler(apiKey string) http.Handler {
	r := chi.NewMux()
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Get("/", handleSearchMovie(apiKey))

	return r
}

func handleSearchMovie(apiKey string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		search := r.URL.Query().Get("search")
		if search == "" {
			sendJSON(w, Response{Error: "Missing search parameter"}, http.StatusBadRequest)
			return
		}
		resp, err := omdb.Search(apiKey, search)
		if err != nil {
			sendJSON(w, Response{Error: err.Error()}, http.StatusBadGateway)
			return
		}
		if len(resp.Search) == 0 {
			sendJSON(w, Response{Data: omdb.Result{
				TotalResults: "0",
				Response:     "true",
			}}, http.StatusOK)
			return
		}

		sendJSON(w, Response{Data: resp}, http.StatusOK)
	}
}
