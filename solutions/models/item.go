package models

import "math"

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
