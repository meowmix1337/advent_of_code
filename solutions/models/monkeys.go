package models

import "sort"

type Monkeys struct {
	Monkeys       []*Monkey `json:"monkeys"`
	CommonDivisor int       `json:"commonDivisor"`
}

func NewMonkeys() *Monkeys {
	return &Monkeys{
		Monkeys:       make([]*Monkey, 0),
		CommonDivisor: 1,
	}
}

func (m *Monkeys) DoRound(skipBored bool) {
	for _, monkey := range m.Monkeys {
		monkey.Inspect(m.Monkeys, skipBored, m.CommonDivisor)
	}
}

func (m *Monkeys) AddMonkey(monkey *Monkey) {
	m.Monkeys = append(m.Monkeys, monkey)
}

func (m *Monkeys) CalculateMonkeyBusiness() int {
	itemsInspected := make([]int, 0)
	for _, monkey := range m.Monkeys {
		itemsInspected = append(itemsInspected, monkey.ItemsInspected)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(itemsInspected)))

	return itemsInspected[0] * itemsInspected[1]
}

func (m *Monkeys) DoRounds(rounds int, skipBored bool) {
	for round := 1; round <= rounds; round++ {
		m.DoRound(skipBored)
	}
}
