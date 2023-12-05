package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	log "github.com/sirupsen/logrus"

	"advent/solutions"
	solutionsv2 "advent/solutions_v2"
	"advent/util/responses"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

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

	solverService := solutionsv2.NewBaseSolver(logger)

	r.Route("/advent/", func(r chi.Router) {
		r.Get("/2022/day/{day}", func(w http.ResponseWriter, r *http.Request) {
			dayNumber, err := strconv.Atoi(chi.URLParam(r, "day"))
			if err != nil {
				responses.Error(w, http.StatusInternalServerError, err)
			}

			year := 2022

			inputFile := fmt.Sprintf("inputfiles/%v/day%v.txt", year, dayNumber)

			logger.Infof("input file: %v; year: %v; day: %v", inputFile, year, dayNumber)

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

		r.Get("/2023/day/{day}", func(w http.ResponseWriter, r *http.Request) {
			dayNumber, err := strconv.Atoi(chi.URLParam(r, "day"))
			if err != nil {
				responses.Error(w, http.StatusInternalServerError, err)
			}

			year := 2023

			inputFile := fmt.Sprintf("inputfiles/%v/day%v.txt", year, dayNumber)

			logger.Infof("input file: %v; year: %v; day: %v", inputFile, year, dayNumber)

			file, err := os.Open(inputFile)
			if err != nil {
				responses.Error(w, http.StatusInternalServerError, err)
				return
			}
			defer file.Close()

			daySolver := solutionsv2.GetDaySolver(solverService, dayNumber, file)
			if daySolver == nil {
				http.Error(w, fmt.Sprintf("day %v solver not found!", dayNumber), http.StatusNotFound)
				return
			}

			dayAnswers := daySolver.Solve()

			responses.JSON(w, http.StatusOK, dayAnswers)
		})
	})

	// start http server
	logger.Info("Server is running on port 8084!")
	http.ListenAndServe(":8084", r)
}
