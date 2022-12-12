package models

type Monkey struct {
	ID              int                 `json:"id"`
	Items           *ItemQueue          `json:"items"`
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
		Items: NewItemQueue(),
	}
}

func (m *Monkey) Inspect(monkeys []*Monkey, skipBored bool, commonDivisor int) {
	for !m.Items.IsEmpty() {
		item := m.Items.Dequeue()
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
}

func (m *Monkey) ThrowToMonkey(throwTo int, item Item, monkeys []*Monkey) {
	monkeys[throwTo].Items.Enqueue(item)
}
