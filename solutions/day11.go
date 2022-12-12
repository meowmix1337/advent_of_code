package solutions

import (
	"bufio"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

type Day11 struct {
	InputFile string
	Logger    *log.Logger
}

type Monkey struct {
	ID              int                 `json:"id"`
	Items           []Item              `json:"items"`
	Operation       func(int, int) int  `json:"-"`
	OperationNumber int                 `json:"operationNumber"`
	Test            func(int, int) bool `json:"-"`
	DisivibleBy     int                 `json:"divisibleBy"`
	TestFailMonkey  int                 `json:"testFailMonkey"`
	TestPassMonkey  int                 `json:"testPassMonkey"`
	ItemsInspected  int                 `json:"itemsInspected"`
}

func (m *Monkey) Inspect(monkeys []*Monkey, skipBored bool) {
	for _, item := range m.Items {
		m.ItemsInspected++

		item.CalcNewLevel(m.Operation, m.OperationNumber)

		if !skipBored {
			item.MonkeyIsBored()
		} else {

		}

		if m.Test(item.WorryLevel, m.DisivibleBy) {
			m.ThrowToMonkey(m.TestPassMonkey, item, monkeys)
		} else {
			m.ThrowToMonkey(m.TestFailMonkey, item, monkeys)
		}
	}

	m.EmptyItems()
}

func (m *Monkey) ThrowToMonkey(throwTo int, item Item, monkeys []*Monkey) {
	monkeys[throwTo].Items = append(monkeys[throwTo].Items, item)
}

func NewMonkey(id int) *Monkey {
	return &Monkey{
		ID:    id,
		Items: make([]Item, 0),
	}
}

func (m *Monkey) EmptyItems() {
	m.Items = []Item{}
}

type Item struct {
	WorryLevel int `json:"worryLevel"`
}

func (i *Item) MonkeyIsBored() {
	i.WorryLevel = int(math.Floor(float64(i.WorryLevel) / 3))
}

func (i *Item) CalcNewLevel(operation func(int, int) int, operationNumber int) {
	// get new number by new = old * operation number
	// if operation number is 0, then use the worry level
	if operationNumber == 0 {
		operationNumber = i.WorryLevel
	}
	i.WorryLevel = operation(i.WorryLevel, operationNumber)
}

func NewItem(worryLevel int) *Item {
	return &Item{
		WorryLevel: worryLevel,
	}
}

type Monkeys []*Monkey

func (m *Monkeys) DoRound(skipBored bool) {
	for _, monkey := range *m {
		monkey.Inspect(*m, skipBored)
	}
}

func (m *Monkeys) AddMonkey(monkey *Monkey) {
	*m = append(*m, monkey)
}

func (m *Monkeys) CalculateMonkeyBusiness() int {
	itemsInspected := make([]int, 0)
	for _, monkey := range *m {
		itemsInspected = append(itemsInspected, monkey.ItemsInspected)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(itemsInspected)))

	return itemsInspected[0] * itemsInspected[1]
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

	for round := 1; round <= 20; round++ {
		monkeys1.DoRound(false)
	}

	for round := 1; round <= 10000; round++ {
		monkeys2.DoRound(true)
	}

	return &Answers{
		Answer1:  monkeys1.CalculateMonkeyBusiness(),
		Answer2:  monkeys2.CalculateMonkeyBusiness(),
		MetaData: monkeys1,
	}, nil
}

func (d *Day11) BuildMonkeys() (*Monkeys, error) {
	file, err := os.Open(d.InputFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// get digits
	re := regexp.MustCompile("[0-9]+")

	monkeys := &Monkeys{}
	var curMonkey *Monkey

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// New monkey
		if strings.Contains(line, "Monkey") {
			// this is the monkey we're gathering information on
			monkeyNum, _ := strconv.Atoi(re.FindAllString(line, -1)[0])
			monkey := NewMonkey(monkeyNum)
			monkeys.AddMonkey(monkey)
			curMonkey = monkey
		}

		// starting items
		if strings.Contains(line, "Starting items: ") {
			worryLevels := re.FindAllString(line, -1)
			for _, worryLevel := range worryLevels {
				worryLevelInt, _ := strconv.Atoi(worryLevel)
				newItem := NewItem(worryLevelInt)
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
