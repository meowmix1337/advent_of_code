package solutionsv2

import (
	"os"

	log "github.com/sirupsen/logrus"
)

type BaseSolver struct {
	Log       *log.Logger
	InputFile *os.File
}

func NewBaseSolver(log *log.Logger) *BaseSolver {
	return &BaseSolver{
		Log: log,
	}
}

type Answers struct {
	Part1Answer int `json:"part1_answer"`
	Part2Answer int `json:"part2_answer"`
}

type Solver interface {
	Solve() *Answers
}

func GetDaySolver(baseSolver *BaseSolver, day int, inputFile *os.File) Solver {
	var daySolver Solver
	switch day {
	case 1:
		daySolver = NewDay1Solver(baseSolver, inputFile)
	case 2:
		daySolver = NewDay2Solver(baseSolver, inputFile)
	case 3:
		daySolver = NewDay3Solver(baseSolver, inputFile)
	case 4:
		daySolver = NewDay4Solver(baseSolver, inputFile)
	case 5:
		daySolver = NewDay5Solver(baseSolver, inputFile)
	}

	return daySolver
}
