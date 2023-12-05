package solutionsv2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"

	log "github.com/sirupsen/logrus"
)

var spelledNumbers = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

var numberMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

type Day1Solver struct {
	*BaseSolver
}

func NewDay1Solver(baseSolver *BaseSolver, inputFile *os.File) Solver {
	day := &Day1Solver{
		BaseSolver: baseSolver,
	}
	day.InputFile = inputFile

	return day
}

func (s *Day1Solver) Solve() *Answers {
	scanner := bufio.NewScanner(s.InputFile)

	total1 := 0
	total2 := 0
	for scanner.Scan() {
		line := scanner.Text()

		total1 += part1(line)
		total2 += part2(line)
	}

	s.Log.WithFields(log.Fields{
		"Part1": total1,
		"Part2": total2,
	}).Info("Answers")

	return &Answers{
		Part1Answer: total1,
		Part2Answer: total2,
	}
}

func charIndexes(firstIdx, lastIdx int, line string) (int, int) {
	for idx, char := range line {
		if unicode.IsDigit(char) {
			if firstIdx == -1 || idx < firstIdx {
				firstIdx = idx
			}
			if lastIdx == -1 || idx > lastIdx {
				lastIdx = idx
			}
		}
	}

	return firstIdx, lastIdx
}

func wordIndexes(firstIdx, lastIdx int, line string) (int, int, string, string) {
	firstWord := ""
	lastWord := ""
	for _, word := range spelledNumbers {
		if idx := strings.Index(line, word); idx > -1 {
			if firstIdx == -1 || idx < firstIdx {
				firstIdx = idx
				firstWord = word
			}
		}
		if idx := strings.LastIndex(line, word); idx > -1 {
			// always use the last index found
			if idx > lastIdx {
				lastIdx = idx
				lastWord = word
			}
		}
	}

	return firstIdx, lastIdx, firstWord, lastWord
}

func getNumber(runeChar rune, word string) string {
	if unicode.IsLetter(runeChar) {
		return numberMap[word]
	}
	return string(runeChar)
}

func part1(line string) int {
	// find numbers
	numStr := ""
	total := 0

	firstIdx, lastIdx := charIndexes(-1, -1, line)
	numStr = strings.Join([]string{string(line[firstIdx]), string(line[lastIdx])}, "")

	num, _ := strconv.Atoi(numStr)
	total += num

	return total
}

func part2(line string) int {
	total := 0
	firstIdx, lastIdx, firstWord, lastWord := wordIndexes(-1, -1, line)
	firstIdx, lastIdx = charIndexes(firstIdx, lastIdx, line)

	firstNumStr := getNumber([]rune(line)[firstIdx], firstWord)
	lastNumStr := getNumber([]rune(line)[lastIdx], lastWord)

	numStr := fmt.Sprintf("%v%v", firstNumStr, lastNumStr)

	num, _ := strconv.Atoi(numStr)
	total += num

	return total
}
