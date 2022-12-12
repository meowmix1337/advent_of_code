package solutions

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"advent/solutions/models"
)

type Day11 struct {
	InputFile string
	Logger    *log.Logger
}

func NewDay11Solver(inputFile string, logger *log.Logger) Solver {
	day11Solver := new(Day11)
	day11Solver.InputFile = inputFile
	day11Solver.Logger = logger

	return day11Solver
}

func (d *Day11) Solve() (*Answers, error) {

	monkeys1, _ := d.BuildMonkeys()
	monkeys2, _ := d.BuildMonkeys()

	monkeys1.DoRounds(20, false)
	monkeys2.DoRounds(10000, true)

	return &Answers{
		Answer1:  monkeys1.CalculateMonkeyBusiness(),
		Answer2:  monkeys2.CalculateMonkeyBusiness(),
		MetaData: monkeys1,
	}, nil
}

func (d *Day11) BuildMonkeys() (*models.Monkeys, error) {
	file, err := os.Open(d.InputFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// get digits
	re := regexp.MustCompile("[0-9]+")

	monkeys := models.NewMonkeys()
	var curMonkey *models.Monkey

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// New monkey
		if strings.Contains(line, "Monkey") {
			// this is the monkey we're gathering information on
			monkeyNum, _ := strconv.Atoi(re.FindAllString(line, -1)[0])
			monkey := models.NewMonkey(monkeyNum)
			monkeys.AddMonkey(monkey)
			curMonkey = monkey
		}

		// starting items
		if strings.Contains(line, "Starting items: ") {
			worryLevels := re.FindAllString(line, -1)
			for _, worryLevel := range worryLevels {
				worryLevelInt, _ := strconv.Atoi(worryLevel)
				newItem := models.NewItem(worryLevelInt)
				curMonkey.Items = append(curMonkey.Items, *newItem)
			}
		}

		// set operation
		if strings.Contains(line, "Operation:") {
			numbers := re.FindAllString(line, -1)
			if len(numbers) > 0 {
				operationNumber, _ := strconv.Atoi(numbers[0])
				curMonkey.OperationNumber = operationNumber
			}

			if strings.Contains(line, "*") {
				curMonkey.Operation = func(worryLevel, operationNumber int) int {
					return worryLevel * operationNumber
				}
				test := curMonkey.Operation(10, curMonkey.OperationNumber)
				d.Logger.Info(test)
			} else if strings.Contains(line, "+") {
				curMonkey.Operation = func(worryLevel, operationNumber int) int {
					return worryLevel + operationNumber
				}
				test := curMonkey.Operation(10, curMonkey.OperationNumber)
				d.Logger.Info(test)
			}
		}

		// set test and monkeys to throw to
		if strings.Contains(line, "Test:") {
			divisibleBy, _ := strconv.Atoi(re.FindAllString(line, -1)[0])
			curMonkey.DisivibleBy = divisibleBy
			curMonkey.Test = func(numToTest, divisibleBy int) bool {
				return numToTest%divisibleBy == 0
			}
			monkeys.CommonDivisor *= divisibleBy
		}

		if strings.Contains(line, "If true:") {
			throwTo, _ := strconv.Atoi(re.FindAllString(line, -1)[0])
			curMonkey.TestPassMonkey = throwTo
		}

		if strings.Contains(line, "If false:") {
			throwTo, _ := strconv.Atoi(re.FindAllString(line, -1)[0])
			curMonkey.TestFailMonkey = throwTo
		}
	}

	return monkeys, nil
}
