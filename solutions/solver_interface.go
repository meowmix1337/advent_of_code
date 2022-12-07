package solutions

import (
	log "github.com/sirupsen/logrus"
)

type Answers struct {
	Answer1  interface{} `json:"answer1"`
	Answer2  interface{} `json:"answer2"`
	MetaData interface{} `json:"metaData"`
}

type Solver interface {
	Solve() (*Answers, error)
}

func GetDaySolver(dayNumber int, inputFile string, logger *log.Logger) Solver {
	var daySolver Solver
	switch dayNumber {
	case 1:
		daySolver = NewDay1Solver(inputFile)
	case 2:
		daySolver = NewDay2Solver(inputFile)
	case 3:
		daySolver = NewDay3Solver(inputFile, logger)
	case 4:
		daySolver = NewDay4Solver(inputFile, logger)
	case 5:
		daySolver = NewDay5Solver(inputFile, logger)
	}

	return daySolver
}
