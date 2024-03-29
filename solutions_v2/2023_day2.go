package solutionsv2

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

// greater than these values:
// 12 red cubes, 13 green cubes, and 14 blue cubes
const (
	MaxReds   = 12
	MaxGreens = 13
	MaxBlues  = 14
)

type Day2Solver struct {
	*BaseSolver
}

func NewDay2Solver(baseSolver *BaseSolver, inputFile *os.File) Solver {
	day := &Day2Solver{
		BaseSolver: baseSolver,
	}
	day.InputFile = inputFile

	return day
}

type Game struct {
	ID       int
	Subset   []Set
	MaxRed   int
	MaxBlue  int
	MaxGreen int
}

type Set struct {
	Red   int
	Green int
	Blue  int
}

func (g *Game) isValid() bool {
	// if any subset has a value greater than specified, then it isn't valid
	for _, subset := range g.Subset {
		if subset.Red > MaxReds || subset.Blue > MaxBlues || subset.Green > MaxGreens {
			return false
		}
	}
	return true
}

func (g *Game) multipleCubes() int {
	return g.MaxGreen * g.MaxBlue * g.MaxRed
}

func (s *Day2Solver) Solve() *Answers {
	scanner := bufio.NewScanner(s.InputFile)

	total1 := 0
	total2 := 0

	games := make([]Game, 0)
	for scanner.Scan() {
		gameLine := scanner.Text()

		// get game ID
		// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
		gameDetails := strings.Split(gameLine, ":")
		gameID, _ := strconv.Atoi(strings.Split(gameDetails[0], " ")[1])
		subsetOfGames := strings.Trim(gameDetails[1], " ")

		game := Game{
			ID:       gameID,
			Subset:   make([]Set, 0),
			MaxBlue:  1,
			MaxGreen: 1,
			MaxRed:   1,
		}

		setOfCubes := strings.Split(subsetOfGames, ";")
		for _, set := range setOfCubes {
			// create the set of cubes
			cubesPerGame := strings.Split(set, ",")
			subset := Set{}
			for _, cubes := range cubesPerGame {
				cubes = strings.Trim(cubes, " ")
				cubeDetails := strings.Split(cubes, " ")
				number, _ := strconv.Atoi(cubeDetails[0])
				color := cubeDetails[1]

				switch color {
				case "red":
					subset.Red = number
					if number > game.MaxRed {
						game.MaxRed = number
					}
				case "blue":
					subset.Blue = number
					if number > game.MaxBlue {
						game.MaxBlue = number
					}
				case "green":
					subset.Green = number
					if number > game.MaxGreen {
						game.MaxGreen = number
					}
				}

			}

			game.Subset = append(game.Subset, subset)
		}

		games = append(games, game)
	}

	for _, game := range games {
		// determine if game is valid
		if game.isValid() {
			total1 += game.ID
		}

		total2 += game.multipleCubes()
	}

	s.Log.WithFields(log.Fields{
		"Part1": total1,
		"Part2": total2,
	}).Info("Answers")

	return &Answers{
		Answer1: total1,
		Answer2: total2,
	}
}
