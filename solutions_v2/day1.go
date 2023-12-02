package solutionsv2

import (
	"bufio"
	"fmt"
	"log"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"unicode"
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

// TODO extract all this crap and create a base solver class that has a logger and input file
// Then an interface that requires each day to implement Part1 and Part2 functions
func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total1 := 0
	total2 := 0
	for scanner.Scan() {
		line := scanner.Text()

		total1 += part1(line)
		total2 += part2(line)
	}

	logger.Info("answers", "part 1", total1, "part 2", total2)
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
