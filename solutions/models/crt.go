package models

import "strconv"

var CPU_CYCLES = map[int]bool{
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
	CurrentRow int `json:"currentRow"`
	// CRT screen pixels
	Pixels [][]string `json:"pixels"`
	// The CRT's CPU
	CPU *CPU `json:"cpu"`
}

func NewCRT(crtPixels []int) *CRT {
	crt := new(CRT)
	crt.DrawLocation = 0
	crt.CurrentRow = 0
	crt.Pixels = make([][]string, 6)

	for i := range crt.Pixels {
		crt.Pixels[i] = make([]string, 40)
	}

	crt.CPU = NewCPU()
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

func (c *CRT) RunCycle(instructions []string) {
	xVal := 0
	if len(instructions) > 1 {
		xVal, _ = strconv.Atoi(instructions[1])
	}

	switch instructions[0] {
	case "noop":
		c.CPU.Cycle++
		if _, ok := CPU_CYCLES[c.CPU.Cycle]; ok {
			c.CPU.SignalStrength += c.CPU.RegisterVal * c.CPU.Cycle
		}
		c.Draw()
	case "addx":
		// addx requires 2 cycles in order to increase the value
		for i := 0; i < 2; i++ {
			c.CPU.Cycle++
			c.Draw()
			if _, ok := CPU_CYCLES[c.CPU.Cycle]; ok {
				c.CPU.SignalStrength += c.CPU.RegisterVal * c.CPU.Cycle
			}
			if i == 1 {
				c.CPU.RegisterVal += xVal
			}
		}
	}
}

func (c *CRT) Draw() {
	// print # if current location is overlapping
	// print . if current location is not overlapping any of the 3 pixels of where RegisterVal is
	if c.IsOverlapping() {
		c.Pixels[c.CurrentRow][c.DrawLocation] = "#"
	} else {
		c.Pixels[c.CurrentRow][c.DrawLocation] = "."
	}

	// if cycle matches the end of the row, increment the CRT's current row
	if c.DrawLocation == len(c.Pixels[c.CurrentRow])-1 {
		c.CurrentRow++
		c.DrawLocation = 0
		return
	}

	// increase the drawLocation
	c.DrawLocation++
}

func (c *CRT) IsOverlapping() bool {
	// check if CRT's current pixel is overlapping the registerVal, registerVal-1 or registerVal+2
	return c.DrawLocation >= c.CPU.RegisterVal-1 && c.DrawLocation <= c.CPU.RegisterVal+1
}
