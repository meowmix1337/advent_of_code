package solutionsv2

import (
	"math"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

const StartingConnectingPipe rune = 'F'
const EndingConnectingPipe rune = 'L'

const (
	UP    int = 0
	RIGHT int = 1
	DOWN  int = 2
	LEFT  int = 3
)

type Day10Solver struct {
	*BaseSolver
}

func NewDay10Solver(baseSolver *BaseSolver, inputFile *os.File) Solver {
	day := &Day10Solver{
		BaseSolver: baseSolver,
	}
	day.InputFile = inputFile

	return day
}

type Tile struct {
	X int
	Y int
}

func NewTile(x, y int) Tile {
	return Tile{x, y}
}

// checks the next Tile up, down, left, right
func (t *Tile) NextTile(tile Tile) Tile {
	return Tile{
		t.X + tile.X,
		t.Y + tile.Y,
	}
}

func (t *Tile) CanConnectToNextTile(nextTile Tile, ls map[Tile]rune, direction int) bool {
	currentPipe := string(ls[*t])
	nextPipe := string(ls[nextTile])

	// ground, do nothing
	if nextPipe == "." || currentPipe == "S" {
		return false
	}

	// 0 up, 1 right, 2 down, 3 left

	// 7
	if currentPipe == "7" {
		if direction != DOWN && direction != LEFT {
			return false
		}
		if direction == DOWN {
			if nextPipe == "|" || nextPipe == "J" || nextPipe == "L" {
				return true
			}
		}
		if direction == LEFT {
			if nextPipe == "L" || nextPipe == "-" || nextPipe == "F" {
				return true
			}
		}
	}

	if currentPipe == "F" {
		if direction != DOWN && direction != RIGHT {
			return false
		}
		if direction == DOWN {
			if nextPipe == "|" || nextPipe == "J" || nextPipe == "L" {
				return true
			}
		}
		if direction == RIGHT {
			if nextPipe == "J" || nextPipe == "-" || nextPipe == "7" {
				return true
			}
		}
	}

	if currentPipe == "L" {
		if direction != RIGHT && direction != UP {
			return false
		}
		if direction == UP {
			if nextPipe == "|" || nextPipe == "F" || nextPipe == "7" {
				return true
			}
		}

		if direction == RIGHT {
			if nextPipe == "J" || nextPipe == "-" || nextPipe == "7" {
				return true
			}
		}
	}

	if currentPipe == "J" {
		if direction != LEFT && direction != UP {
			return false
		}
		if direction == UP {
			if nextPipe == "|" || nextPipe == "F" || nextPipe == "7" {
				return true
			}
		}

		if direction == LEFT {
			if nextPipe == "L" || nextPipe == "F" || nextPipe == "-" {
				return true
			}
		}
	}

	if currentPipe == "|" {
		if direction != DOWN && direction != UP {
			return false
		}
		if direction == DOWN {
			if nextPipe == "|" || nextPipe == "L" || nextPipe == "J" {
				return true
			}
		}

		if direction == UP {
			if nextPipe == "|" || nextPipe == "F" || nextPipe == "7" {
				return true
			}
		}
	}

	if currentPipe == "-" {
		if direction != LEFT && direction != RIGHT {
			return false
		}
		if direction == LEFT {
			if nextPipe == "-" || nextPipe == "L" || nextPipe == "F" {
				return true
			}
		}

		if direction == RIGHT {
			if nextPipe == "-" || nextPipe == "J" || nextPipe == "7" {
				return true
			}
		}

	}

	return false
}

type TileQueue []Tile

func NewTileQueue() *TileQueue {
	return &TileQueue{}
}

func (q *TileQueue) Enqueue(tile Tile) {
	*q = append(*q, tile)
}

func (q *TileQueue) Dequeue() Tile {
	tile := (*q)[0]
	*q = (*q)[1:]
	return tile
}

func (q *TileQueue) IsEmpty() bool {
	return len(*q) == 0
}

type Landscape struct {
	Log        *log.Logger
	Landscape  map[Tile]rune
	Queue      TileQueue
	StartPoint Tile
	Visited    map[Tile]bool
	Steps      int
}

func NewLandscape(log *log.Logger) *Landscape {
	return &Landscape{
		Log:       log,
		Landscape: make(map[Tile]rune),
		Visited:   make(map[Tile]bool),
		Queue:     make([]Tile, 0),
		Steps:     0,
	}
}

func (ls *Landscape) BuildLandscape(input string) {
	for y, row := range strings.Fields(input) {
		for x, char := range row {
			ls.Landscape[NewTile(x, y)] = char

			if char == 'S' {
				ls.StartPoint = NewTile(x, y)
			}
		}
	}

	ls.Steps++
	ls.Queue.Enqueue(ls.StartPoint)
}

func (ls *Landscape) HasVisited(nextTile Tile) bool {
	_, visited := ls.Visited[nextTile]
	return visited
}

func (ls *Landscape) FindPath() {
	for !ls.Queue.IsEmpty() {
		curTile := ls.Queue.Dequeue()

		// ls.Log.Info(string(ls.Landscape[curTile]))

		// 0 UP, 1 RIGHT, 2 DOWN, 3 LEFT
		for direction, neighbor := range []Tile{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
			nextTile := curTile.NextTile(neighbor)

			// exit case
			// we know the ending pipe and we know we need to go up to finish
			if nextTile == ls.StartPoint && ls.Landscape[curTile] == EndingConnectingPipe && direction == UP {
				break
			}

			// starting point, specific input can only go left
			if curTile == ls.StartPoint && ls.Landscape[nextTile] == StartingConnectingPipe && direction == LEFT {
				ls.Visited[nextTile] = true
				ls.Queue.Enqueue(nextTile)
				ls.Steps++
				break
			}

			if !ls.HasVisited(nextTile) && curTile.CanConnectToNextTile(nextTile, ls.Landscape, direction) {
				ls.Visited[nextTile] = true
				ls.Queue.Enqueue(nextTile)
				ls.Steps++
				break
			}
		}
	}
}

func (s *Day10Solver) Solve() *Answers {

	content, err := os.ReadFile("inputfiles/2023/day10.txt")
	if err != nil {
		log.Fatalln("failed to open file", err)
	}

	landscape := NewLandscape(s.Log)
	landscape.BuildLandscape(string(content))

	landscape.FindPath()

	s.Log.Info("answers", "part 1", math.Ceil(float64(landscape.Steps/2)), "part 2", 0)

	return &Answers{
		Answer1: landscape.Steps / 2,
	}
}
