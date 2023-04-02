package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	log "github.com/sirupsen/logrus"

	"advent/solutions"
	"advent/util/responses"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	logger := log.New()
	logger.Formatter = &log.JSONFormatter{}
	logger.Out = os.Stdout
	logger.Level = log.DebugLevel

	r.Route("/advent", func(r chi.Router) {
		r.Get("/day/{day}", func(w http.ResponseWriter, r *http.Request) {
			dayNumber, err := strconv.Atoi(chi.URLParam(r, "day"))
			if err != nil {
				responses.Error(w, http.StatusInternalServerError, err)
			}

			inputFile := fmt.Sprintf("../inputfiles/day%v/input.txt", dayNumber)

			daySolver := solutions.GetDaySolver(dayNumber, inputFile, logger)
			if daySolver == nil {
				http.Error(w, fmt.Sprintf("day %v solver not found!", dayNumber), http.StatusNotFound)
				return
			}

			dayAnswers, err := daySolver.Solve()
			if err != nil {
				responses.Error(w, http.StatusInternalServerError, err)
				return
			}

			responses.JSON(w, http.StatusOK, dayAnswers)
		})
	})

	// start http server
	logger.Info("Server is running!")
	http.ListenAndServe(":8084", r)
}
