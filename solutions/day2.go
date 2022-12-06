package solutions

import (
	"bufio"
	"os"
	"strings"
)

type Day2MetaData struct {
}

var actionPoints = map[string]int{
	"Y": 2, // paper
	"X": 1, // rock
	"Z": 3, // scissors
}

var WIN_CONDITIONS = map[string]bool{
	"YA": true, // paper beats rock
	"XC": true, // rock beats scissors
	"ZB": true, // scissors beats paper
}

var LOSE_CONDITIONS = map[string]bool{
	"YC": true, // paper loses to scissors
	"XB": true, // rock loses to paper
	"ZA": true, // scissors loses to rock
}

var I_NEED_TO_LOSE = map[string]string{
	"A": "Z", // opp chooses rock, i need scissors
	"B": "X", // opp choose paper, i need rock
	"C": "Y", // opp chooses scissors, i need paper
}

var I_NEED_TO_WIN = map[string]string{
	"A": "Y", // opp chooses rock, i need paper
	"B": "Z", // opp choose paper, i need scissors
	"C": "X", // opp chooses scissors, i need rock
}

var WHAT_I_NEED_TO_DRAW = map[string]string{
	"A": "X", // rock
	"B": "Y", // paper
	"C": "Z", // scissors
}

const WIN = 6
const LOSE = 0
const DRAW = 3

type Day2 struct {
	InputFile string
}

func NewDay2Solver(inputFile string) Solver {
	day2Solver := new(Day2)
	day2Solver.InputFile = inputFile

	return day2Solver
}

func (d *Day2) Solve() (*Answers, error) {
	part1, err := d.part1()
	if err != nil {
		return nil, err
	}

	part2, err := d.part2()
	if err != nil {
		return nil, err
	}

	return &Answers{
		Answer1:  part1,
		Answer2:  part2,
		MetaData: nil,
	}, nil
}

func (d *Day2) part1() (int, error) {
	file, err := os.Open(d.InputFile)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	totalScore := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		ourActions := strings.Fields(line)
		opponent := ourActions[0]
		me := ourActions[1]

		action := me + opponent

		if _, ok := WIN_CONDITIONS[action]; ok { // win conditions
			totalScore += WIN + actionPoints[me]
		} else if _, ok := LOSE_CONDITIONS[action]; ok { // lose conditions
			totalScore += LOSE + actionPoints[me]
		} else { // draw conditions
			totalScore += DRAW + actionPoints[me]
		}
	}

	return totalScore, nil
}

func (d *Day2) part2() (int, error) {
	file, err := os.Open(d.InputFile)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	totalScore := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		ourActions := strings.Fields(line)
		opponent := ourActions[0]
		me := ourActions[1]

		if me == "X" { // lose
			totalScore += LOSE + actionPoints[I_NEED_TO_LOSE[opponent]]
		} else if me == "Y" { // draw
			totalScore += DRAW + actionPoints[WHAT_I_NEED_TO_DRAW[opponent]]
		} else if me == "Z" { // win
			totalScore += WIN + actionPoints[I_NEED_TO_WIN[opponent]]
		}
	}

	return totalScore, nil
}
