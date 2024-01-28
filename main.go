package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"log/slog"
	"net/http"
	"os"
	"server/internal/database"
	"server/internal/handlers"
	"server/internal/store"
)

func main() {
  env := os.Getenv("ENV")
  if env != "PROD"{
    err := godotenv.Load()
    if err != nil {
      fmt.Println("Failed to read env var file")
      panic(err)
    }
  }
	textHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{})
	logger := slog.New(textHandler)
	postgresPool := database.NewPostgresPool()
	redisClient := database.RedisClient()
	users := store.NewUsers(postgresPool)
	otpCodes := store.NewOTPCodes(redisClient)
	sessions := store.NewSessionsStore(redisClient)
	authHandler := handlers.NewAuthHandler(otpCodes, users, sessions, logger)
  webhookHandler := handlers.NewWebhookHandler()
	r := chi.NewRouter()

  r.Get("/hook", webhookHandler.Verify)
  r.Post("/hook", webhookHandler.Handle)

	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", authHandler.Register)
	})
	fmt.Println("Server launched on port 8081")
  err := http.ListenAndServe(":8081", r)
	if err != nil {
		panic(err)
	}
}
