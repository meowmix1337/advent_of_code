package solutions

import (
	"bufio"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

type Day3 struct {
	InputFile string
	Logger    *log.Logger
}

func NewDay3Solver(inputFile string, logger *log.Logger) Solver {
	day3Solver := new(Day3)
	day3Solver.InputFile = inputFile
	day3Solver.Logger = logger

	return day3Solver
}

var lowerPoints = generatePoints(1, 26, 'a')
var upperPoints = generatePoints(27, 52, 'A')

func (d *Day3) Solve() (*Answers, error) {
	file, err := os.Open(d.InputFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	total := 0
	scanner := bufio.NewScanner(file)

	groups := make(map[int][]string)

	groupCounter := 0
	groupId := 1
	for scanner.Scan() {
		line := scanner.Text()

		if groupCounter == 3 {
			groupCounter = 0
			groupId++
		}

		// add line to group
		if _, ok := groups[groupId]; !ok {
			groups[groupId] = make([]string, 0)
			groups[groupId] = append(groups[groupId], line)
		} else {
			groups[groupId] = append(groups[groupId], line)
		}
		groupCounter++

		half := len(line)

		mid := half / 2
		// get first half
		firstHalf := line[:mid]
		// get second half
		secHalf := line[mid:]

		// get the letter that appears in both halfs
		letter := d.getCommonLetter(firstHalf, secHalf)

		d.calcPoints(letter, &total)
	}

	// solve part 2
	totalPart2 := 0
	for _, group := range groups {
		letter := d.getCommonLetterByGroup(group)
		d.calcPoints(letter, &totalPart2)
	}

	return &Answers{
		Answer1:  total,
		Answer2:  totalPart2,
		MetaData: nil,
	}, nil
}

func (d *Day3) getCommonLetterByGroup(group []string) string {
	first := d.getCount(group[0])
	second := d.getCount(group[1])
	third := d.getCount(group[2])

	for char := range first {
		// check if char exists in 2nd group
		if _, ok := second[char]; ok {
			// check if char exists in 3rd group
			if _, ok2 := third[char]; ok2 {
				return char
			}
		}
	}

	return ""
}

func (d *Day3) getCommonLetter(first, second string) string {
	firstCount := d.getCount(first)
	secondCount := d.getCount(second)

	for char := range firstCount {
		// check if the letter exists in second half
		if _, ok := secondCount[char]; ok {
			return char
		}
	}

	return ""
}

func (d *Day3) calcPoints(letter string, total *int) {
	if points, ok := lowerPoints[letter]; ok {
		d.Logger.Info(fmt.Sprintf("points: %v, letter: %v", points, letter))
		*total += points
	} else if points, ok := upperPoints[letter]; ok {
		d.Logger.Info(fmt.Sprintf("points: %v, letter: %v", points, letter))
		*total += points
	}
}

func (d *Day3) getCount(compartment string) map[string]int {
	count := make(map[string]int)

	for _, char := range compartment {
		charStr := string(char)
		if _, ok := count[charStr]; !ok {
			count[charStr] = 1
		} else {
			count[charStr]++
		}
	}

	return count
}

func generatePoints(start, end int, startLetter int) map[string]int {
	points := make(map[string]int)
	for i := start; i <= end; i++ {
		letter := rune(startLetter + i - start)
		points[fmt.Sprintf("%c", letter)] = i
	}
	return points
}
