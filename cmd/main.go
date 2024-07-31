package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	v1 "github.com/neelbhat88/go-api-template/cmd/domain/v1"
	"github.com/neelbhat88/go-api-template/internal/apimiddleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Warn().Msg("No .env file found")
	}

	if os.Getenv("ENV") == "local" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	//var dbConfig postgres.DatabaseConfig
	//err = cleanenv.ReadEnv(&dbConfig)
	//if err != nil {
	//	log.Error().Err(err).Msg("Failed to read DatabaseConfig from config.env")
	//}

	//var appConfig AppConfig
	//err = cleanenv.ReadEnv(&appConfig)
	//if err != nil {
	//	log.Error().Err(err).Msg("Failed to read AppConfig from config.env")
	//}

	//ms := DatabaseMigrationSource{}
	//db, err := postgres.InitializeDB(dbConfig, ms)
	//if err != nil {
	//	log.Fatal().Err(err).Msg("Failed to initialize DB")
	//}
	//defer db.Close()

	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(apimiddleware.RequestResponseLogger)
	r.Use(middleware.Heartbeat("/ping"))
	r.Use(apimiddleware.Recoverer)
	r.Use(middleware.Compress(5))
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/", v1.Root)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Info().Str("port", port).Msg("Application started")

	http.ListenAndServe(fmt.Sprintf(":%v", port), r)
}
