package models

import (
	"fmt"
)

type Location struct {
	X       int
	Y       int
	Visited bool
}

type Map struct {
	Path map[string]*Location `json:"path"`
}

func NewLocation(x, y int, visited bool) *Location {
	return &Location{
		X:       x,
		Y:       y,
		Visited: visited,
	}
}

func (m *Map) CountUniquelyVisited() int {
	count := 0
	for _, location := range m.Path {
		if location.Visited {
			count++
		}
	}
	return count
}

func (m *Map) AddLocation(x, y int, visited bool) {
	m.Path[fmt.Sprintf("%v_%v", x, y)] = NewLocation(x, y, visited)
}

func (m *Map) GetLocation(x, y int) *Location {
	return m.Path[fmt.Sprintf("%v_%v", x, y)]
}
