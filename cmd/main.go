package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	log "github.com/sirupsen/logrus"

	"advent/controller"
	solutionsv2 "advent/solutions_v2"
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

	controller2022 := controller.NewController2022(logger)
	solverCtrl := controller.NewSolverController(logger, solverService)

	r.Route("/"+controller.APIV1+"/advent/", func(r chi.Router) {
		r.Get("/2022/{day}", controller2022.Solve2022Day)
		r.Get("/{year}/{day}", solverCtrl.SolveDay)
	})

	// start http server
	logger.Info("Server is running on port 8084!")
	http.ListenAndServe(":8084", r)
}
