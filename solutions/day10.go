package solutions

import (
	"bufio"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"advent/solutions/models"
)

type Day10 struct {
	InputFile string
	Logger    *log.Logger
}

type Answer2 struct {
	Row1 string `json:"row1"`
	Row2 string `json:"row2"`
	Row3 string `json:"row3"`
	Row4 string `json:"row4"`
	Row5 string `json:"row5"`
	Row6 string `json:"row6"`
}

func NewDay10Solver(inputFile string, logger *log.Logger) Solver {
	day10Solver := new(Day10)
	day10Solver.InputFile = inputFile
	day10Solver.Logger = logger

	return day10Solver
}

func (d *Day10) Solve() (*Answers, error) {
	file, err := os.Open(d.InputFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	crt := models.NewCRT([]int{39, 79, 119, 159, 199, 239})

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		instructions := strings.Split(line, " ")

		crt.RunCycle(instructions)
	}

	return &Answers{
		Answer1: crt.CPU.SignalStrength,
		Answer2: Answer2{
			Row1: strings.Join(crt.Pixels[0], ""),
			Row2: strings.Join(crt.Pixels[1], ""),
			Row3: strings.Join(crt.Pixels[2], ""),
			Row4: strings.Join(crt.Pixels[3], ""),
			Row5: strings.Join(crt.Pixels[4], ""),
			Row6: strings.Join(crt.Pixels[5], ""),
		},
		MetaData: crt,
	}, nil
}
