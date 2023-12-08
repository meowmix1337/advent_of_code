package solutionsv2

import (
	"bufio"
	"os"
	"strings"
)

type Day8Solver struct {
	*BaseSolver
}

func NewDay8Solver(baseSolver *BaseSolver, inputFile *os.File) Solver {
	day := &Day8Solver{
		BaseSolver: baseSolver,
	}
	day.InputFile = inputFile

	return day
}

type Instructions struct {
	Movements []string
}

func NewInstructions(movements []string) *Instructions {
	instructions := new(Instructions)
	instructions.Movements = movements
	return instructions
}

type Network map[string]*Node

func NewNetwork() *Network {
	network := new(Network)
	*network = make(Network)
	return network
}

func (n Network) AddNode(position string, node *Node) {
	n[position] = node
}

func (n Network) GetNodeDirection(position string, direction string) string {
	if direction == "L" {
		return n[position].Left
	}
	return n[position].Right
}

type Node struct {
	Left  string
	Right string
}

func NewNode(left, right string) *Node {
	node := new(Node)
	node.Left = left
	node.Right = right

	return node
}

func (s Day8Solver) Solve() *Answers {
	scanner := bufio.NewScanner(s.InputFile)

	instructions := NewInstructions([]string{})
	network := NewNetwork()

	startingPositions := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		// if no instructions, then we are first line
		if len(instructions.Movements) == 0 {
			instructions.Movements = strings.Split(line, "")
			continue
		}

		// lets build out the network
		//left side
		position := line[:strings.Index(line, "=")-1]
		left := line[strings.Index(line, "(")+1 : strings.Index(line, ",")]
		right := line[strings.Index(line, ",")+2 : strings.Index(line, ")")]

		node := NewNode(left, right)

		network.AddNode(position, node)

		// part 2 needs all that end in A
		if strings.HasSuffix(position, "A") {
			startingPositions = append(startingPositions, position)
		}
	}

	steps1 := day8Part1(network, instructions.Movements)

	// part 2
	steps2 := day8Part2(network, startingPositions, instructions.Movements)

	return &Answers{
		Answer1: steps1,
		Answer2: steps2,
	}
}

func day8Part1(network *Network, instructions []string) int {
	steps := 0
	goal := "AAA"
	for goal != "ZZZ" {
		for _, direction := range instructions {
			goal = network.GetNodeDirection(goal, direction)

			if goal == "ZZZ" {
				steps++
				break
			}
			steps++
		}
	}
	return steps
}

func day8Part2(network *Network, startingPositions []string, instructions []string) int {
	directionIdx := 0
	numIterationsPerKey := make([]int, len(startingPositions))
	for idx := range startingPositions {
		numIterations := 0
		for {
			if directionIdx == len(instructions) {
				directionIdx = 0
			}

			if strings.HasSuffix(startingPositions[idx], "Z") {
				break
			}

			nextPos := network.GetNodeDirection(startingPositions[idx], instructions[directionIdx])
			startingPositions[idx] = nextPos

			directionIdx++
			numIterations++
		}
		numIterationsPerKey[idx] = numIterations
	}
	return leastCommonMultiple(numIterationsPerKey)
}

func leastCommonMultiple(numbers []int) int {
	lcm := numbers[0]
	for i := 0; i < len(numbers); i++ {
		num1 := lcm
		num2 := numbers[i]
		gcd := 1
		for num2 != 0 {
			temp := num2
			num2 = num1 % num2
			num1 = temp
		}
		gcd = num1
		lcm = (lcm * numbers[i]) / gcd
	}

	return lcm
}
