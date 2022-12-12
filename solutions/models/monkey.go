package models

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

func NewMonkey(id int) *Monkey {
	return &Monkey{
		ID:    id,
		Items: make([]Item, 0),
	}
}

func (m *Monkey) Inspect(monkeys []*Monkey, skipBored bool, commonDivisor int) {
	for _, item := range m.Items {
		m.ItemsInspected++

		item.CalcNewLevel(m.Operation, m.OperationNumber)

		if !skipBored {
			item.MonkeyIsBored()
		} else {
			item.WorryLevel = item.WorryLevel % commonDivisor
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

func (m *Monkey) EmptyItems() {
	m.Items = []Item{}
}
