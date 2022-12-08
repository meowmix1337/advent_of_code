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
	switch dayNumber {
	case 1:
		return NewDay1Solver(inputFile)
	case 2:
		return NewDay2Solver(inputFile)
	case 3:
		return NewDay3Solver(inputFile, logger)
	case 4:
		return NewDay4Solver(inputFile, logger)
	case 5:
		return NewDay5Solver(inputFile, logger)
	case 6:
		return NewDay6Solver(inputFile, logger)
	case 7:
		return NewDay7Solver(inputFile, logger)
	case 8:
		return NewDay8Solver(inputFile, logger)
	}

	return nil
}
