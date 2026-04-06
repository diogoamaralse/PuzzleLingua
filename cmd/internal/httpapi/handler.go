package httpapi

import (
	"PuzzleLingua/cmd/internal/data"
	"PuzzleLingua/cmd/internal/game"
	"encoding/json"
	"net/http"
)

type Handler struct {
	service *game.Service
	mux     *http.ServeMux
}

func NewHandler() http.Handler {
	h := &Handler{
		service: game.NewService(data.DefaultPuzzles()),
		mux:     http.NewServeMux(),
	}

	h.routes()
	return h.withCORS(h.mux)
}

func (h *Handler) routes() {
	h.mux.HandleFunc("/api/health", h.handleHealth)
	h.mux.HandleFunc("/api/round", h.handleRound)
	h.mux.HandleFunc("/api/check", h.handleCheck)
}

func (h *Handler) handleHealth(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (h *Handler) handleRound(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	puzzle := h.service.RandomPuzzle()
	writeJSON(w, http.StatusOK, game.RoundResponse{Puzzle: puzzle})
}

func (h *Handler) handleCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	var req game.CheckRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid json"})
		return
	}

	result := h.service.Check(req)
	writeJSON(w, http.StatusOK, result)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func (h *Handler) withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
