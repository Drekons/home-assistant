package main

import (
	"context"
	"github.com/Drekons/home-assistant/backend/cmd/app"
	"github.com/Drekons/home-assistant/backend/internal/api"
	"github.com/Drekons/home-assistant/backend/internal/websocket"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()
	application := app.NewApp(ctx)
	defer application.Shutdown(ctx)

	handler := api.NewHandler(application.Deps)
	router := mux.NewRouter()

	// Add routes
	router.HandleFunc("/login", handler.Login).Methods("POST")
	router.HandleFunc("/logout", handler.Logout).Methods("POST")
	router.HandleFunc("/ws", websocket.HandleWebSocket)

	// Start the server
	log.Printf("Starting server on :%s", application.Deps.Config().Server.Port)
	log.Fatal(http.ListenAndServe(":"+application.Deps.Config().Server.Port, router))
}
