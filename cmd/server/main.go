package main

import (
	"context"
	"net/http"

	"github.com/eif-courses/restaurant/internal/config"
	"github.com/eif-courses/restaurant/internal/handlers"
	"github.com/eif-courses/restaurant/internal/logger"
	"github.com/eif-courses/restaurant/internal/repository"
	"github.com/eif-courses/restaurant/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {

	log := logger.NewLogger()

	defer log.Sync()

	err := godotenv.Load()
	if err != nil {
		log.Fatalw("Error loading .env file", "error", err)
	}

	cfg := config.NewConfig()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	db, err := pgxpool.New(context.Background(), cfg.DatabaseUrl)

	if err != nil {
		log.Fatalw("Failed to connect to db!", "error", err)
	}
	defer db.Close()

	log.Infow("Database successfully connected!")

	queries := repository.New(db)

	foodService := services.NewFoodService(queries)
	foodHandler := handlers.NewFoodHandler(foodService)

	r.Get("/foods", foodHandler.GetFoodList)
	r.Post("/foods", foodHandler.InsertFruit)
	r.Get("/hello", foodHandler.GetHelloText)

	http.ListenAndServe(":8080", r)

}
