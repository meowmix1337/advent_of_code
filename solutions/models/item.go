package models

import "math"

type Item struct {
	WorryLevel int `json:"worryLevel"`
}

type ItemQueue []Item

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

func NewItemQueue() *ItemQueue {
	return &ItemQueue{}
}

func (q *ItemQueue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *ItemQueue) Dequeue() Item {
	item := (*q)[0]
	*q = (*q)[1:]
	return item
}

func (q *ItemQueue) Enqueue(item Item) {
	*q = append(*q, item)
}
