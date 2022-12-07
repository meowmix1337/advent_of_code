package solutions

import (
	"bufio"
	"os"

	log "github.com/sirupsen/logrus"
)

type Day6 struct {
	InputFile string
	Logger    *log.Logger
}

type Day6Chunks struct {
	Answer1Chunk string `json:"chunk1"`
	Answer2Chunk string `json:"chunk2"`
}

func NewDay6Solver(inputFile string, logger *log.Logger) Solver {
	day6Solver := new(Day6)
	day6Solver.InputFile = inputFile
	day6Solver.Logger = logger

	return day6Solver
}

func (d *Day6) Solve() (*Answers, error) {
	file, err := os.Open(d.InputFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	dataStream := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		dataStream = scanner.Text()
	}

	answer1, chunk := d.solve(dataStream, 3)
	answer2, chunk2 := d.solve(dataStream, 13)

	return &Answers{
		Answer1: answer1,
		Answer2: answer2,
		MetaData: Day6Chunks{
			Answer1Chunk: chunk,
			Answer2Chunk: chunk2,
		},
	}, nil
}

func (d *Day6) solve(dataStream string, endIdx int) (int, string) {
	startIdx := 0
	chunk := dataStream[startIdx : endIdx+1]
	for d.hasRepeats(chunk) {
		// move the start and end points by 1
		startIdx++
		endIdx++
		chunk = dataStream[startIdx : endIdx+1]
	}

	return endIdx + 1, chunk
}

func (d *Day6) hasRepeats(chunk string) bool {
	set := make(map[string]int)
	for i := 0; i < len(chunk); i++ {
		c := chunk[i]
		if _, ok := set[string(c)]; !ok {
			set[string(c)] = 1
		} else {
			return true
		}
	}
	return false
}
