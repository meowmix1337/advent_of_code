package models

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() string {
	if s.IsEmpty() {
		return ""
	} else {
		idx := len(*s) - 1   // last idx is top
		element := (*s)[idx] // get element
		*s = (*s)[:idx]      // remove the last element by doing some slicing
		return element
	}
}

func (s *Stack) Peek() string {
	return (*s)[len(*s)-1] // get element
}
