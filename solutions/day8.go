package solutions

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"advent/solutions/models"
)

type Day8 struct {
	InputFile string
	Logger    *log.Logger
}

func NewDay8Solver(inputFile string, logger *log.Logger) Solver {
	day8Solver := new(Day8)
	day8Solver.InputFile = inputFile
	day8Solver.Logger = logger

	return day8Solver
}

func (d *Day8) Solve() (*Answers, error) {
	file, err := os.Open(d.InputFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	forest := models.Forest{}

	row := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		rowOfTrees := make([]models.Tree, 0)
		for col, height := range strings.Split(line, "") {
			h, _ := strconv.Atoi(height)
			tree := models.Tree{
				Height: h,
				X:      col,
				Y:      row,
			}

			rowOfTrees = append(rowOfTrees, tree)
		}
		forest.Trees = append(forest.Trees, rowOfTrees)
		row++
	}

	forest.Rows = row
	forest.Cols = len(forest.Trees)

	return &Answers{
		Answer1:  forest.CountVisibleTrees(),
		Answer2:  forest.GetHighestScenicScore(),
		MetaData: forest,
	}, nil
}
