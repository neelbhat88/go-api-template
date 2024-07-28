package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"time"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	err := godotenv.Load()
	if err != nil {
		log.Warn().Msg("No .env file found")
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
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")

		if name == "world" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Name cannot be 'world'"))
			return
		}
		w.Write([]byte(fmt.Sprintf("Hello, %s!", name)))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Info().Str("port", port).Msg("Application started")

	http.ListenAndServe(fmt.Sprintf(":%v", port), r)
}
