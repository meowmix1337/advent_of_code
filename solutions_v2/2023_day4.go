package solutionsv2

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

type Day4Solver struct {
	*BaseSolver
}

func NewDay4Solver(baseSolver *BaseSolver, inputFile *os.File) Solver {
	day := &Day4Solver{
		BaseSolver: baseSolver,
	}
	day.InputFile = inputFile

	return day
}

type Cards []*Card

type Card struct {
	ID             int
	WinningNumbers []int
	Numbers        []int
	Matches        int
	Points         int
	NumCopies      int
}

func (cs *Cards) AppendCard(card *Card) {
	*cs = append(*cs, card)
}

func NewCard(idStr string, winningNumbers, numbers []string) *Card {
	id, _ := strconv.Atoi(idStr)
	card := &Card{
		ID:             id,
		WinningNumbers: convertStrToInts(winningNumbers),
		Numbers:        convertStrToInts(numbers),
		NumCopies:      1,
	}
	points, matches := card.GetPointsAndMatches()

	card.Points = points
	card.Matches = matches

	return card
}

func (c Card) GetPointsAndMatches() (int, int) {
	points := 0
	matches := 0
	matchMap := make(map[int]bool)
	for _, winNum := range c.WinningNumbers {
		for _, myNum := range c.Numbers {
			if winNum != myNum || matchMap[winNum] {
				continue
			}

			// match, increase match count and add to map
			if points >= 1 {
				points *= 2
			} else {
				points++
			}
			matches++
			matchMap[winNum] = true
			break
		}
	}
	return points, matches
}

func convertStrToInts(strings []string) []int {
	nums := make([]int, 0)
	for _, str := range strings {
		if str == "" {
			continue
		}
		num, _ := strconv.Atoi(str)
		nums = append(nums, num)
	}
	return nums
}

func (s *Day4Solver) Solve() *Answers {
	scanner := bufio.NewScanner(s.InputFile)

	total1 := 0
	total2 := 0

	cards := Cards{}
	for scanner.Scan() {
		line := scanner.Text()

		cardDetails := strings.Split(line, ":")
		cardID := strings.Split(cardDetails[0], " ")[1]
		cardNumbers := strings.Split(cardDetails[1], "|")
		winningNums := strings.Split(strings.Trim(cardNumbers[0], " "), " ")
		myNumbers := strings.Split(strings.Trim(cardNumbers[1], " "), " ")

		card := NewCard(cardID, winningNums, myNumbers)

		cards.AppendCard(card)
	}

	// part 1
	for _, card := range cards {
		total1 += card.Points
	}

	// part 2
	for idx, card := range cards {
		for i := 0; i < card.NumCopies; i++ {
			// reset the idx after we finished copying the first time
			cardIdxToCopy := idx
			matches := card.Matches
			// we need to copy the n next cards (matches) 1 time
			for matches != 0 {
				// add 1 copy to the next n cards
				cardIdxToCopy += 1
				cards[cardIdxToCopy].NumCopies++
				matches--
			}
		}
	}

	for _, card := range cards {
		total2 += card.NumCopies
	}

	s.Log.WithFields(log.Fields{
		"Part1": total1,
		"Part2": total2,
	}).Info("Answers")

	return &Answers{
		Part1Answer: total1,
		Part2Answer: total2,
	}
}
