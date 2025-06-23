package api

import (
	"github.com/Drekons/home-assistant/backend/cmd/app"
	"net/http"
)

type Handler struct {
	deps *app.Deps
}

func NewHandler(deps *app.Deps) *Handler {
	return &Handler{deps: deps}
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	// Implement logout logic here
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Logout successful"))
}
