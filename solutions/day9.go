package solutions

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"advent/solutions/models"
)

type Day9 struct {
	InputFile string
	Logger    *log.Logger
}

type Day9MetaData struct {
	Part1Map models.Map `json:"map1"`
	Part2Map models.Map `json:"map2"`
}

var SECTIONS = 9

func NewDay9Solver(inputFile string, logger *log.Logger) Solver {
	day9Solver := new(Day9)
	day9Solver.InputFile = inputFile
	day9Solver.Logger = logger

	return day9Solver
}

func (d *Day9) Solve() (*Answers, error) {
	file, err := os.Open(d.InputFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	part1Head, part1MapArea := d.CreatePart1()
	part2Head, part2MapArea := d.CreatePart2()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		instructions := strings.Split(line, " ")
		direction := instructions[0]
		steps, _ := strconv.Atoi(instructions[1])

		// loop through each step in the same direction
		for steps != 0 {
			// part 1
			part1Head.Move(direction, part1MapArea)
			part2Head.Move(direction, part2MapArea)
			steps--
		}
	}
	return &Answers{
		Answer1: part1MapArea.CountUniquelyVisited(),
		Answer2: part2MapArea.CountUniquelyVisited(),
		MetaData: Day9MetaData{
			Part1Map: part1MapArea,
			Part2Map: part2MapArea,
		},
	}, nil
}

func (d *Day9) CreatePart1() (*models.Node, models.Map) {
	mapArea := models.Map{
		Path: make(map[string]*models.Location),
	}

	startX := 0
	startY := 0

	// add starting location
	mapArea.AddLocation(startX, startY, true)

	// head
	head := &models.Node{
		X:  startX,
		Y:  startY,
		ID: 0,
	}

	// tail
	tail := &models.Node{
		IsTail: true,
		X:      startX,
		Y:      startY,
		Parent: head,
		ID:     1,
	}

	head.ChildNode = tail

	return head, mapArea
}

func (d *Day9) CreatePart2() (*models.Node, models.Map) {
	mapArea := models.Map{
		Path: make(map[string]*models.Location),
	}

	startX := 0
	startY := 0

	// add starting location
	mapArea.AddLocation(startX, startY, true)

	// head
	head := &models.Node{
		X:  startX,
		Y:  startY,
		ID: 0,
	}

	tempNextNode := head
	// create 1-9 nodes, node 9 is tail
	for i := 0; i < SECTIONS; i++ {
		node := &models.Node{
			X:      startX,
			Y:      startY,
			Parent: tempNextNode,
			ID:     i + 1,
		}
		if i == SECTIONS-1 {
			node.IsTail = true
		}
		tempNextNode.ChildNode = node
		tempNextNode = node
	}

	return head, mapArea
}
