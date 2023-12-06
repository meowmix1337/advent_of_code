package controller

import (
	solutionsv2 "advent/solutions_v2"
	"advent/util/responses"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
)

type SolverController struct {
	*Controller
	SolverService *solutionsv2.BaseSolver
}

func NewSolverController(log *log.Logger, solverService *solutionsv2.BaseSolver) *SolverController {
	return &SolverController{
		Controller:    NewController(log),
		SolverService: solverService,
	}
}

func (c *SolverController) SolveDay(w http.ResponseWriter, r *http.Request) {
	year, err := strconv.Atoi(chi.URLParam(r, "year"))
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
	}

	dayNumber, err := strconv.Atoi(chi.URLParam(r, "day"))
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
	}

	inputFile := fmt.Sprintf("inputfiles/%v/day%v.txt", year, dayNumber)

	c.Log.Infof("input file: %v; year: %v; day: %v", inputFile, year, dayNumber)

	file, err := os.Open(inputFile)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer file.Close()

	daySolver := solutionsv2.GetDaySolver(c.SolverService, dayNumber, file)
	if daySolver == nil {
		http.Error(w, fmt.Sprintf("day %v solver not found!", dayNumber), http.StatusNotFound)
		return
	}

	dayAnswers := daySolver.Solve()

	responses.JSON(w, http.StatusOK, dayAnswers)
}
