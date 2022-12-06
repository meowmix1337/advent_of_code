package solutions

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

type Day1MetaData struct {
	CalorieCount map[int]int `json:"calorieCount"`
}

type Day1 struct {
	InputFile string
}

func NewDay1Solver(inputFile string) Solver {
	day1Solver := new(Day1)
	day1Solver.InputFile = inputFile

	return day1Solver
}

func (d *Day1) Solve() (*Answers, error) {
	file, err := os.Open(d.InputFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	calories := make(map[int]int)
	sortedCalories := make([]int, 0)

	currentElf := 0
	mostCalories := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// it is a new elf if line is blank
		if line == "" {
			currentElf++
			continue
		}

		cals, _ := strconv.Atoi(line)

		// if elf doesn't exist in map, add it and the current value
		if _, ok := calories[currentElf]; !ok {
			calories[currentElf] = cals
		} else {
			calories[currentElf] += cals
		}

		// current elf's cals are the most
		if calories[currentElf] > mostCalories {
			mostCalories = calories[currentElf]
		}
	}

	for _, value := range calories {
		sortedCalories = append(sortedCalories, value)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sortedCalories)))

	topThree := sortedCalories[0] + sortedCalories[1] + sortedCalories[2]

	return &Answers{
		Answer1:  mostCalories,
		Answer2:  topThree,
		MetaData: calories,
	}, nil
}
