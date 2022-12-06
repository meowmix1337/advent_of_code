package solutions

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

type Day4 struct {
	InputFile string
	Logger    *log.Logger
}

type LineData struct {
	Min1 int
	Min2 int
	Max1 int
	Max2 int
}

func NewDay4Solver(inputFile string, logger *log.Logger) Solver {
	day4Solver := new(Day4)
	day4Solver.InputFile = inputFile
	day4Solver.Logger = logger

	return day4Solver
}

func (d *Day4) Solve() (*Answers, error) {
	file, err := os.Open(d.InputFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	part1Total := 0
	part2Total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		groupings := strings.Split(line, ",")

		d.Logger.Info(groupings)

		lineData := d.makeGroup(groupings[0], groupings[1])

		// single number and within range of second range
		if lineData.singleNumInRange() || lineData.AtLeastOneRangeFullyContains() {
			part1Total++
		}

		if lineData.OverlapsAtAll() {
			part2Total++
		}

	}

	return &Answers{
		Answer1:  part1Total,
		Answer2:  part2Total,
		MetaData: nil,
	}, nil
}

func (d *Day4) makeGroup(group1, group2 string) LineData {
	range1 := d.getRange(group1)
	range2 := d.getRange(group2)

	return LineData{
		Min1: range1[0],
		Max1: range1[1],
		Min2: range2[0],
		Max2: range2[1],
	}
}

func (d *Day4) getRange(group string) []int {
	rangeStr := strings.Split(group, "-")
	rangeInt := make([]int, len(rangeStr))

	for i, s := range rangeStr {
		rangeInt[i], _ = strconv.Atoi(s)
	}

	return rangeInt
}

func (d *LineData) singleNumInRange() bool {
	return (d.Min1 == d.Max1 && d.Min1 >= d.Min2 && d.Min1 <= d.Max2) ||
		(d.Min2 == d.Max2 && d.Min2 >= d.Min1 && d.Min2 <= d.Max1)
}

func (d *LineData) AtLeastOneRangeFullyContains() bool {
	return (d.Min1 >= d.Min2 &&
		d.Min1 <= d.Max2 &&
		d.Max1 <= d.Max2 &&
		d.Max1 >= d.Min2) ||

		(d.Min2 >= d.Min1 &&
			d.Min2 <= d.Max1 &&
			d.Max2 <= d.Max1 &&
			d.Max2 >= d.Min1)
}

func (d *LineData) OverlapsAtAll() bool {
	return d.Min1 <= d.Max2 && d.Min2 <= d.Max1
}
