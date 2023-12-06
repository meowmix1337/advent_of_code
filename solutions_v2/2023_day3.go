package solutionsv2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"

	log "github.com/sirupsen/logrus"
)

type Day3Solver struct {
	*BaseSolver
}

func NewDay3Solver(baseSolver *BaseSolver, inputFile *os.File) Solver {
	day := &Day3Solver{
		BaseSolver: baseSolver,
	}
	day.InputFile = inputFile

	return day
}

type Number struct {
	Value    string
	StartCol int
	EndCol   int
	Row      int
}

type Symbol struct {
	Type string
	Row  int
	Col  int
}

type AdjacentNumbers struct {
	Numbers []*Number
}

func NewAdjacentNumbers() *AdjacentNumbers {
	return &AdjacentNumbers{make([]*Number, 0)}
}

func (a *AdjacentNumbers) add(number *Number) {
	a.Numbers = append(a.Numbers, number)
}

func (a *AdjacentNumbers) len() int {
	return len(a.Numbers)
}

func (a *AdjacentNumbers) multiple() int {
	total := 1
	for _, num := range a.Numbers {
		total *= num.getNumberValue()
	}

	return total
}

func (a *AdjacentNumbers) reset() bool {
	if a.len() > 2 {
		a.Numbers = []*Number{}
		return true
	}
	return false
}

func NewSymbol(symbolType string, row, col int) Symbol {
	return Symbol{
		Type: symbolType,
		Row:  row,
		Col:  col,
	}
}

func NewNumber(initialValue string, startCol, row int) *Number {
	return &Number{
		Value:    initialValue,
		StartCol: startCol,
		EndCol:   0,
		Row:      row,
	}
}

func (n *Number) max() bool {
	return len(n.Value) == 3
}

func (n *Number) appendValue(value string) {
	n.Value += value
}

func (n *Number) getNumberValue() int {
	num, _ := strconv.Atoi(n.Value)
	return num
}

func (n *Number) isAdjacentToSymbol(symbolCol, symbolRow int) bool {
	// check if the symbol is adjacent to the number
	// if the number row is outside +/- 1 from symbol or if the number isn't even in the same row, then we can just skip it since it isn't adjacent
	// Ex. ..123..
	//     .......
	//     ...$...
	//     .......
	//     ..123..
	if symbolRow+1 != n.Row && symbolRow-1 != n.Row && symbolRow != n.Row {
		return false
	}

	// if the number column is outside the symbol's +/- 1, then we aren't adjacent
	// Ex. ..123.$.456
	if symbolCol+1 != n.StartCol && symbolCol-1 != n.EndCol && !(symbolCol <= n.EndCol && symbolCol >= n.StartCol) {
		return false
	}

	return true
}

func (s Day3Solver) Solve() *Answers {
	scanner := bufio.NewScanner(s.InputFile)

	total1 := 0
	total2 := 0

	numbers := make([]*Number, 0)
	symbols := make([]Symbol, 0)
	row := 0
	for scanner.Scan() {
		line := scanner.Text()

		// go through each character in a line
		var potentialNumber *Number
		for col, char := range line {
			c := fmt.Sprintf("%c", char)
			if c == "." {
				if potentialNumber != nil {
					// We need to handle len 2 numbers, we've already handled len 3 during isDigit
					// if we have a potential number, and we hit a period, it means we have len 2 and should set the end col
					potentialNumber.EndCol = col - 1
					numbers = append(numbers, potentialNumber)
					potentialNumber = nil
				}
				continue
			}

			// symbol
			if unicode.IsSymbol(char) || unicode.IsPunct(char) || unicode.IsControl(char) {
				// reset number
				if potentialNumber != nil {
					// same thing with symbol
					potentialNumber.EndCol = col - 1
					numbers = append(numbers, potentialNumber)
					potentialNumber = nil
				}
				symbol := NewSymbol(c, row, col)
				symbols = append(symbols, symbol)
				continue
			}

			// number
			if unicode.IsDigit(char) {
				if potentialNumber == nil {
					potentialNumber = NewNumber(c, col, row)
				} else {
					// otherwise, we add the value
					potentialNumber.appendValue(c)
					if potentialNumber.max() {
						potentialNumber.EndCol = col
						numbers = append(numbers, potentialNumber)
						potentialNumber = nil
					}
				}
			}

			// handle the case where we have a potential number with len of 2, but we reached the end of the line
			if potentialNumber != nil && col+1 == len(line) && len(potentialNumber.Value) == 2 {
				potentialNumber.EndCol = len(line) - 1
				numbers = append(numbers, potentialNumber)
				potentialNumber = nil
			}
		}

		row++
	}

	// we now have all the info we need for part 1
	for _, symbol := range symbols {
		for _, number := range numbers {
			if !number.isAdjacentToSymbol(symbol.Col, symbol.Row) {
				continue
			}
			total1 += number.getNumberValue()
		}
	}

	// part 2
	// we only care about * and there needs to be exactly TWO adjacent numbers
	for _, symbol := range symbols {
		if symbol.Type != "*" {
			continue
		}

		adjacentNumbers := NewAdjacentNumbers()
		for _, number := range numbers {
			if !number.isAdjacentToSymbol(symbol.Col, symbol.Row) {
				continue
			}

			adjacentNumbers.add(number)

			// if we have more than 2, we can move on to the next *
			if adjacentNumbers.reset() {
				break
			}
		}
		// exactly 2 adjacent numbers for this *
		if adjacentNumbers.len() == 2 {
			total2 += adjacentNumbers.multiple()
		}
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
