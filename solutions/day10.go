package solutions

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

type Day10 struct {
	InputFile string
	Logger    *log.Logger
}

func NewDay10Solver(inputFile string, logger *log.Logger) Solver {
	day10Solver := new(Day10)
	day10Solver.InputFile = inputFile
	day10Solver.Logger = logger

	return day10Solver
}

var CYCLES = map[int]bool{
	20:  true,
	60:  true,
	100: true,
	140: true,
	180: true,
	220: true,
}

type CRT struct {
	// DrawLocation always starts at 0 and +1 each cycle
	DrawLocation int `json:"pixelLocation"`
	// CurrentRow keeps track of which row the CRT is printing
	// Once the CPU hits 40/80/120/160/200/240, increment the row
	CurrentRow int          `json:"currentRow"`
	Pixels     [][]string   `json:"pixels"`
	CPU        *CPU         `json:"cpu"`
	CRTPixels  map[int]bool `json:"crtPixels"`
}

func NewCRT(crtPixels []int) *CRT {
	crt := new(CRT)
	crt.DrawLocation = 0
	crt.CurrentRow = 0
	crt.Pixels = make([][]string, 0)
	crt.CPU = NewCPU()
	crt.CRTPixels = make(map[int]bool)

	for _, pixels := range crtPixels {
		crt.CRTPixels[pixels] = true
	}
	return crt
}

type CPU struct {
	// RegisterVal is the sprite location which is the middle of the 3 pixels
	// ###
	// RegisterVal is 1 so the index is 1. which means registerVal-1 and registerVal+1 are the other 2 pixels
	RegisterVal    int `json:"registerVal"`
	Cycle          int `json:"cycle"`
	SignalStrength int `json:"signalStrength"`
}

func NewCPU() *CPU {
	return &CPU{
		RegisterVal:    1,
		Cycle:          0,
		SignalStrength: 0,
	}
}

func (c *CRT) runCycle(instructions []string) {
	xVal := 0
	if len(instructions) > 1 {
		xVal, _ = strconv.Atoi(instructions[1])
	}

	switch instructions[0] {
	case "noop":
		c.CPU.Cycle++
		if _, ok := CYCLES[c.CPU.Cycle]; ok {
			c.CPU.SignalStrength += c.CPU.RegisterVal * c.CPU.Cycle
		}
	case "addx":
		// addx requires 2 cycles in order to increase the value
		for i := 0; i < 2; i++ {
			c.CPU.Cycle++
			if _, ok := CYCLES[c.CPU.Cycle]; ok {
				c.CPU.SignalStrength += c.CPU.RegisterVal * c.CPU.Cycle
			}
			if i == 1 {
				c.CPU.RegisterVal += xVal
			}
		}
	}
}

func (d *Day10) Solve() (*Answers, error) {
	file, err := os.Open(d.InputFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	crt := NewCRT([]int{40, 80, 120, 160, 200, 240})

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		instructions := strings.Split(line, " ")

		crt.runCycle(instructions)
	}

	return &Answers{
		Answer1:  crt.CPU.SignalStrength,
		Answer2:  0,
		MetaData: crt,
	}, nil
}
