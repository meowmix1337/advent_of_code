package solutions

import (
	"os"

	log "github.com/sirupsen/logrus"

	"advent/solutions/models"
)

type Day12 struct {
	InputFile string
	Logger    *log.Logger
}

func NewDay12Solver(inputFile string, logger *log.Logger) Solver {
	day12Solver := new(Day12)
	day12Solver.InputFile = inputFile
	day12Solver.Logger = logger

	return day12Solver
}

func (d *Day12) Solve() (*Answers, error) {
	input, _ := os.ReadFile(d.InputFile)

	pathFinder := models.NewPathFinder()

	pathFinder.BuildHeightMap(string(input))

	startPath, shortestPath := pathFinder.PathFind()

	return &Answers{
		Answer1:  startPath,
		Answer2:  shortestPath,
		MetaData: nil,
	}, nil
}
