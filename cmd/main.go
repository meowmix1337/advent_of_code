package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
