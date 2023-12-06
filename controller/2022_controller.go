package controller

import (
	"advent/solutions"
	"advent/util/responses"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
)

type Controller2022 struct {
	*Controller
}

func NewController2022(log *log.Logger) *Controller2022 {
	return &Controller2022{
		Controller: NewController(log),
	}
}

func (c *Controller2022) Solve2022Day(w http.ResponseWriter, r *http.Request) {
	dayNumber, err := strconv.Atoi(chi.URLParam(r, "day"))
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
	}

	year := 2022

	inputFile := fmt.Sprintf("inputfiles/%v/day%v.txt", year, dayNumber)

	c.Log.Infof("input file: %v; year: %v; day: %v", inputFile, year, dayNumber)

	daySolver := solutions.GetDaySolver(dayNumber, inputFile, c.Log)
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
}
