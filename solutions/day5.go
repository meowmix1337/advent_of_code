package solutions

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"advent/solutions/models"
)

type Day5 struct {
	InputFile string
	Logger    *log.Logger
}

var regex = regexp.MustCompile(`(move\s|from\s|to\s)`)

type Stacks map[int]*models.Stack

func NewDay5Solver(inputFile string, logger *log.Logger) Solver {
	day5Solver := new(Day5)
	day5Solver.InputFile = inputFile
	day5Solver.Logger = logger

	return day5Solver
}

func (d *Day5) Solve() (*Answers, error) {

	// we know the stacks so just hardcode it
	stack1 := &models.Stack{"Q", "W", "P", "S", "Z", "R", "H", "D"}
	stack2 := &models.Stack{"V", "B", "R", "W", "Q", "H", "F"}
	stack3 := &models.Stack{"C", "V", "S", "H"}
	stack4 := &models.Stack{"H", "F", "G"}
	stack5 := &models.Stack{"P", "G", "J", "B", "Z"}
	stack6 := &models.Stack{"Q", "T", "J", "H", "W", "F", "L"}
	stack7 := &models.Stack{"Z", "T", "W", "D", "L", "V", "J", "N"}
	stack8 := &models.Stack{"D", "T", "Z", "C", "J", "G", "H", "F"}
	stack9 := &models.Stack{"W", "P", "V", "M", "B", "H"}

	stacks := Stacks{
		1: stack1,
		2: stack2,
		3: stack3,
		4: stack4,
		5: stack5,
		6: stack6,
		7: stack7,
		8: stack8,
		9: stack9,
	}

	stacks2 := make(map[int]*models.Stack, len(stacks))
	for id, stack := range stacks {
		stackCopy := *stack
		stacks2[id] = &stackCopy
	}

	file, err := os.Open(d.InputFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		move, from, to := d.getActions(line)

		d.doMoveFromTo(move, stacks[from], stacks[to])

		d.doMoveFromTo9001(move, stacks2[from], stacks2[to])
	}

	d.Logger.Info(stacks2)

	return &Answers{
		Answer1: strings.Join([]string{
			stacks[1].Peek(),
			stacks[2].Peek(),
			stacks[3].Peek(),
			stacks[4].Peek(),
			stacks[5].Peek(),
			stacks[6].Peek(),
			stacks[7].Peek(),
			stacks[8].Peek(),
			stacks[9].Peek(),
		}, ""),
		Answer2: strings.Join([]string{
			stacks2[1].Peek(),
			stacks2[2].Peek(),
			stacks2[3].Peek(),
			stacks2[4].Peek(),
			stacks2[5].Peek(),
			stacks2[6].Peek(),
			stacks2[7].Peek(),
			stacks2[8].Peek(),
			stacks2[9].Peek(),
		}, ""),
		MetaData: stacks,
	}, nil
}

func (d *Day5) getActions(cmd string) (int, int, int) {
	replaceCmd := regex.ReplaceAllString(cmd, "")
	cleanCmd := strings.Split(replaceCmd, " ")

	move, _ := strconv.Atoi(cleanCmd[0])
	from, _ := strconv.Atoi(cleanCmd[1])
	to, _ := strconv.Atoi(cleanCmd[2])

	return move, from, to
}

func (d *Day5) doMoveFromTo(move int, fromStk, toStk *models.Stack) {

	// pop # of move from the from stack and push to the to stack
	for i := 0; i < move; i++ {
		topElement := fromStk.Pop()
		toStk.Push(topElement)
	}
}

func (d *Day5) doMoveFromTo9001(move int, fromStk, toStk *models.Stack) {
	// we need to keep order the same so pop it to a temp stack to pop back into to stack
	tempStack := models.Stack{}
	d.doMoveFromTo(move, fromStk, &tempStack)
	d.doMoveFromTo(move, &tempStack, toStk)
}
