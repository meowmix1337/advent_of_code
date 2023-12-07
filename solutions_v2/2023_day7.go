package solutionsv2

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Day7Solver struct {
	*BaseSolver
}

func NewDay7Solver(baseSolver *BaseSolver, inputFile *os.File) Solver {
	day := &Day7Solver{
		BaseSolver: baseSolver,
	}
	day.InputFile = inputFile

	return day
}

var HandTypes = map[string]int{
	"5K": 7,
	"4K": 6,
	"FH": 5,
	"3K": 4,
	"2P": 3,
	"1P": 2,
	"HC": 1,
}

var LabelValuesMap = map[string]int{
	"A": 13,
	"K": 12,
	"Q": 11,
	"J": 0, // part 2 this is 0, part 1 is 10
	"T": 9,
	"9": 8,
	"8": 7,
	"7": 6,
	"6": 5,
	"5": 4,
	"4": 3,
	"3": 2,
	"2": 1,
}

var LabelValues = []string{
	"A",
	"K",
	"Q",
	"J",
	"T",
	"9",
	"8",
	"7",
	"6",
	"5",
	"4",
	"3",
	"2",
}

type CamelHands struct {
	Hands []*Hand
}

func (ch *CamelHands) AppendHand(hand *Hand) {
	ch.Hands = append(ch.Hands, hand)
}

type Hand struct {
	Hand            string
	Part2Hand       string
	Labels          []string
	Part2Labels     []string
	LabelCount      map[string]int
	JCount          int
	Rank            int
	Type            string
	Part2Type       string
	BidAmount       int
	TotalAmount     int
	HighCountLabel  string
	FirstPairLabel  string
	SecondPairLabel string
}

func NewHand(handDetails []string) *Hand {
	bidAmount, _ := strconv.Atoi(handDetails[1])
	return &Hand{
		Hand:        handDetails[0],
		Labels:      strings.Split(handDetails[0], ""),
		BidAmount:   bidAmount,
		Rank:        -1,
		TotalAmount: 1,
		JCount:      0,
		Part2Labels: make([]string, 0),
	}
}

func (h *Hand) DetermineHandType(part2 bool) {
	if !part2 {
		if h.IsFiveOfKind() {
			h.Type = "5K"
		} else if h.IsFourOfKind() {
			h.Type = "4K"
		} else if h.IsFullHouse() {
			h.Type = "FH"
		} else if h.IsThreeOfKind() {
			h.Type = "3K"
		} else if h.IsTwoPairs() {
			h.Type = "2P"
		} else if h.IsPair() {
			h.Type = "1P"
		} else {
			h.Type = "HC"
		}
		return
	}

	// after we determine the type
	// check the count of the # of Js
	// depending on the # of Js, add it to the type
	// i.e. 1P 22JJA with 2 Js -> 4K 2222A
	// i.e. 3K 222J3 with 1 J -> 4K 22223
	// i.e. 1P 1JJJJ with 4 Js -> 5K 11111
	// when determining types, ignore Js for now
	if h.IsFiveOfKind() {
		// by default, it always wins
		h.Type = "5K"
	} else if h.IsFourOfKind() {
		h.Part2Hand = strings.Replace(h.Hand, "J", h.HighCountLabel, h.JCount)
		if h.JCount == 1 {
			h.Type = "5K"
			return
		}
		h.Type = "4K"
	} else if h.IsFullHouse() {
		h.Type = "FH"
	} else if h.IsThreeOfKind() {
		h.Part2Hand = strings.Replace(h.Hand, "J", h.HighCountLabel, h.JCount)
		if h.JCount == 1 {
			h.Type = "4K"
			return
		} else if h.JCount == 2 {
			h.Type = "5K"
			return
		}
		h.Type = "3K"
	} else if h.IsTwoPairs() {
		// there can only be 1 J in this case
		// this would make a FH
		h.Part2Hand = strings.Replace(h.Hand, "J", h.GetHighestValue(), h.JCount)
		if h.JCount == 1 {
			h.Type = "FH"
			return
		}
		h.Type = "2P"
	} else if h.IsPair() {
		h.Part2Hand = strings.Replace(h.Hand, "J", h.HighCountLabel, h.JCount)
		if h.JCount == 1 {
			h.Type = "3K"
			return
		} else if h.JCount == 2 {
			h.Type = "4K"
			return
		} else if h.JCount == 3 {
			h.Type = "5K"
			return
		}
		h.Type = "1P"
	} else {
		h.Part2Hand = strings.Replace(h.Hand, "J", h.GetHighestValue(), h.JCount)
		if h.JCount == 1 {
			h.Type = "1P"
			return
		} else if h.JCount == 2 {
			h.Type = "3K"
			return
		} else if h.JCount == 3 {
			h.Type = "4K"
			return
		} else if h.JCount == 4 || h.JCount == 5 {
			h.Type = "5K"
			return
		}
		h.Type = "HC"
	}
	return
}

func (h *Hand) IsFullHouse() bool {
	// check if there is a pair and three of kind
	hasPair := false
	hasThree := false
	for _, count := range h.LabelCount {
		if count == 2 {
			hasPair = true
		}
		if count == 3 {
			hasThree = true
		}

		if hasThree && hasPair {
			return true
		}
	}

	return false
}

func (h *Hand) IsTwoPairs() bool {
	// check if there is a pair and three of kind
	firstPair := false
	secondPair := false
	for label, count := range h.LabelCount {
		if count == 2 && !firstPair {
			h.FirstPairLabel = label
			firstPair = true
			continue
		}
		if count == 2 && !secondPair {
			h.SecondPairLabel = label
			secondPair = true
			continue
		}
	}

	if firstPair && secondPair {
		return true
	}

	return false
}

// IsPair look for exactly 2
func (h *Hand) IsPair() bool {
	for label, count := range h.LabelCount {
		if count == 2 {
			h.HighCountLabel = label
			return true
		}
	}
	return false
}

// IsThreeOfKind look for exactly 3
func (h *Hand) IsThreeOfKind() bool {
	for label, count := range h.LabelCount {
		if count == 3 {
			h.HighCountLabel = label
			return true
		}
	}
	return false
}

// IsFourOfKind look for exactly 4
func (h *Hand) IsFourOfKind() bool {
	for label, count := range h.LabelCount {
		if count == 4 {
			h.HighCountLabel = label
			return true
		}
	}
	return false
}

// IsFiveOfKind look for exactly 5
func (h *Hand) IsFiveOfKind() bool {
	for label, count := range h.LabelCount {
		if count == 5 {
			h.HighCountLabel = label
			return true
		}
	}
	return false
}

func (h *Hand) CountLabels(part2 bool) {
	labelCount := make(map[string]int)

	for _, char := range h.Labels {
		// we can ignore J for initial count as we'll need to replace them to make the strongest hand
		if part2 && char == "J" {
			// count the number of Js
			h.JCount++
			continue
		}
		labelCount[char]++
	}

	h.LabelCount = labelCount
}

func customSort(arr []*Hand) {
	sort.Slice(arr, func(i, j int) bool {
		hand1 := arr[i]
		hand2 := arr[j]

		// Custom comparison logic based on specific rules for each character
		for idx := 0; idx < len(hand1.Hand) && idx < len(hand2.Hand); idx++ {
			char1 := string(hand1.Hand[idx])
			char2 := string(hand2.Hand[idx])
			switch {
			case LabelValuesMap[char1] > LabelValuesMap[char2]:
				return true
			case LabelValuesMap[char1] < LabelValuesMap[char2]:
				return false
			}
		}

		// If the strings are the same until a certain length, sort by length
		return len(hand1.Hand) < len(hand2.Hand)
	})
}

func (h *Hand) GetHighestValue() string {
	highestValue := -1
	highestLabel := ""
	for _, char := range h.Labels {
		// ignore J
		if char == "J" {
			continue
		}
		if LabelValuesMap[char] > highestValue {
			highestValue = LabelValuesMap[char]
			highestLabel = char
		}
	}
	return highestLabel
}

func (s Day7Solver) Solve() *Answers {
	scanner := bufio.NewScanner(s.InputFile)

	total1 := 0
	total2 := 0

	camelHands := CamelHands{}
	for scanner.Scan() {
		line := scanner.Text()
		handDetails := strings.Split(line, " ")

		hand := NewHand(handDetails)
		hand.CountLabels(true)
		hand.DetermineHandType(true)

		camelHands.AppendHand(hand)
	}

	// group the hands now
	groupedTypes := make(map[string][]*Hand)
	for _, hand := range camelHands.Hands {
		if _, exists := groupedTypes[hand.Type]; !exists {
			groupedTypes[hand.Type] = make([]*Hand, 0)
			groupedTypes[hand.Type] = append(groupedTypes[hand.Type], hand)
		} else {
			groupedTypes[hand.Type] = append(groupedTypes[hand.Type], hand)
		}
	}

	// now we need to determine the rank by going through each group
	for _, group := range groupedTypes {
		// we need to sort each hand in each group based on rules
		customSort(group)
	}

	sortedHands := make([]*Hand, 0)
	sortedHands = append(sortedHands, groupedTypes["5K"]...)
	sortedHands = append(sortedHands, groupedTypes["4K"]...)
	sortedHands = append(sortedHands, groupedTypes["FH"]...)
	sortedHands = append(sortedHands, groupedTypes["3K"]...)
	sortedHands = append(sortedHands, groupedTypes["2P"]...)
	sortedHands = append(sortedHands, groupedTypes["1P"]...)
	sortedHands = append(sortedHands, groupedTypes["HC"]...)

	// give rank
	for idx, hand := range sortedHands {
		s.Log.Info(fmt.Sprintf("%v - %v", hand.Hand, hand.Part2Hand))
		hand.Rank = len(sortedHands) - idx
		hand.TotalAmount *= hand.BidAmount
	}

	for _, hand := range sortedHands {
		total1 += hand.TotalAmount * hand.Rank
	}

	s.Log.Info("answers", "part 1", total1, "part 2", total2)

	return &Answers{
		Answer1: total1,
		Answer2: total2,
	}
}
