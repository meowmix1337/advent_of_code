package solutionsv2

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Day9Solver struct {
	*BaseSolver
}

func NewDay9Solver(baseSolver *BaseSolver, inputFile *os.File) Solver {
	day := &Day9Solver{
		BaseSolver: baseSolver,
	}
	day.InputFile = inputFile

	return day
}

func allZeros(nums []int) bool {
	for _, num := range nums {
		if num != 0 {
			return false
		}
	}
	return true
}

func nextPredictionValue(currentPrediction []int) int {
	// we really only need to add the last value of each row until we hit 0
	result := currentPrediction[len(currentPrediction)-1]

	for !allZeros(currentPrediction) {
		nextPrediction := make([]int, 0)
		for idx := range currentPrediction {
			// exit case, we've reach the end of the array
			if idx == len(currentPrediction)-1 {
				break
			}

			// calc the difference and add to the next prediction
			num1 := currentPrediction[idx]
			num2 := currentPrediction[idx+1]
			delta := num2 - num1
			nextPrediction = append(nextPrediction, delta)
		}
		// swap things
		currentPrediction = nextPrediction

		// add the last number to the result since that's the only number we care
		result += currentPrediction[len(currentPrediction)-1]
	}

	return result
}

func (s Day9Solver) Solve() *Answers {
	scanner := bufio.NewScanner(s.InputFile)

	total1 := 0
	total2 := 0

	data := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()

		rowData := make([]int, 0)

		nums := strings.Split(line, " ")
		for _, numStr := range nums {
			num, _ := strconv.Atoi(numStr)
			rowData = append(rowData, num)
		}

		data = append(data, rowData)
	}

	for _, row := range data {
		total1 += nextPredictionValue(row)
	}

	// part 2, just reverse it
	for _, row := range data {
		slices.Reverse(row)
	}

	for _, row := range data {
		total2 += nextPredictionValue(row)
	}

	return &Answers{
		Answer1: total1,
		Answer2: total2,
	}
}
