package solutionsv2

import (
	"bufio"
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
	"J": 10,
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

var LabelValues = []string{"A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2"}

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

func (h *Hand) HandleJokerWildCard(hand, highestCardLabel string, jokerCount int) {
	h.Part2Hand = strings.Replace(hand, "J", highestCardLabel, jokerCount)
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
		h.Type = h.GetJokerFourOfKindHandType()
	} else if h.IsFullHouse() {
		h.Type = "FH"
	} else if h.IsThreeOfKind() {
		h.Part2Hand = strings.Replace(h.Hand, "J", h.HighCountLabel, h.JCount)
		h.Type = h.GetJokerThreeOfKindHandType()
	} else if h.IsTwoPairs() {
		// there can only be 1 J in this case
		// this would make a FH
		h.Part2Hand = strings.Replace(h.Hand, "J", h.GetHighestValue(part2), h.JCount)
		h.Type = h.GetJokerTwoPairHandType()
	} else if h.IsPair() {
		h.Part2Hand = strings.Replace(h.Hand, "J", h.HighCountLabel, h.JCount)
		h.Type = h.GetJokerPairHandType()
	} else {
		h.Part2Hand = strings.Replace(h.Hand, "J", h.GetHighestValue(part2), h.JCount)
		h.Type = h.GetJokerHighHandType()
	}
}

func (h *Hand) GetJokerFourOfKindHandType() string {
	switch h.JCount {
	case 1:
		return "5K"
	default:
		return "4K"
	}
}

func (h *Hand) GetJokerThreeOfKindHandType() string {
	switch h.JCount {
	case 1:
		return "4K"
	case 2:
		return "5K"
	default:
		return "3K"
	}
}

func (h *Hand) GetJokerTwoPairHandType() string {
	switch h.JCount {
	case 1:
		return "FH"
	default:
		return "2P"
	}
}

func (h *Hand) GetJokerPairHandType() string {
	switch h.JCount {
	case 1:
		return "3K"
	case 2:
		return "4K"
	case 3:
		return "5K"
	default:
		return "1P"
	}
}

func (h *Hand) GetJokerHighHandType() string {
	switch h.JCount {
	case 1:
		return "1P"
	case 2:
		return "3K"
	case 3:
		return "4K"
	case 4:
		fallthrough
	case 5:
		return "5K"
	default:
		return "HC"
	}
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

func (s Day7Solver) sortHand(arr []*Hand, labelValues map[string]int) {
	sort.Slice(arr, func(i, j int) bool {
		hand1 := arr[i]
		hand2 := arr[j]

		// Custom comparison logic based on specific rules for each character
		for idx := 0; idx < len(hand1.Hand) && idx < len(hand2.Hand); idx++ {
			char1 := string(hand1.Hand[idx])
			char2 := string(hand2.Hand[idx])

			switch {
			case labelValues[char1] > labelValues[char2]:
				return true
			case labelValues[char1] < labelValues[char2]:
				return false
			}
		}

		// If the strings are the same until a certain length, sort by length
		return len(hand1.Hand) < len(hand2.Hand)
	})
}

func (h *Hand) GetHighestValue(part2 bool) string {
	highestValue := -1
	highestLabel := ""
	for _, char := range h.Labels {
		// ignore J
		if part2 && char == "J" {
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
	camelHands2 := CamelHands{}
	for scanner.Scan() {
		line := scanner.Text()
		handDetails := strings.Split(line, " ")

		hand := NewHand(handDetails)
		hand.CountLabels(false)
		hand.DetermineHandType(false)

		hand2 := NewHand(handDetails)
		hand2.CountLabels(true)
		hand2.DetermineHandType(true)

		camelHands.AppendHand(hand)
		camelHands2.AppendHand(hand2)
	}

	// group the hands now
	part1GroupByType := s.groupHandsByType(camelHands)
	part2GroupByType := s.groupHandsByType(camelHands2)

	// now we need to determine the rank by going through each group
	sortedPart1GroupByType := s.sortGroupedHandsByType(part1GroupByType, LabelValuesMap)
	LabelValuesMap["J"] = 0
	sortedPart2GroupByType := s.sortGroupedHandsByType(part2GroupByType, LabelValuesMap)

	sortedHandsPart1 := s.generateSortedHands(sortedPart1GroupByType)
	sortedHandsPart2 := s.generateSortedHands(sortedPart2GroupByType)

	// give rank
	s.giveRank(sortedHandsPart1)
	s.giveRank(sortedHandsPart2)

	total1 = s.calculateWinnings(sortedHandsPart1)
	total2 = s.calculateWinnings(sortedHandsPart2)

	s.Log.Info("answers", "part 1 ", total1, "part 2 ", total2)

	return &Answers{
		Answer1: total1,
		Answer2: total2,
	}
}

func (s Day7Solver) groupHandsByType(camelHands CamelHands) map[string][]*Hand {
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

	return groupedTypes
}

func (s Day7Solver) sortGroupedHandsByType(groupedHandsByType map[string][]*Hand, labelValues map[string]int) map[string][]*Hand {
	// now we need to determine the rank by going through each group
	for _, group := range groupedHandsByType {
		// we need to sort each hand in each group based on rules
		s.sortHand(group, labelValues)
	}

	return groupedHandsByType
}

func (s Day7Solver) generateSortedHands(sortedGroupedHandsByType map[string][]*Hand) []*Hand {
	sortedHands := make([]*Hand, 0)
	sortedHands = append(sortedHands, sortedGroupedHandsByType["5K"]...)
	sortedHands = append(sortedHands, sortedGroupedHandsByType["4K"]...)
	sortedHands = append(sortedHands, sortedGroupedHandsByType["FH"]...)
	sortedHands = append(sortedHands, sortedGroupedHandsByType["3K"]...)
	sortedHands = append(sortedHands, sortedGroupedHandsByType["2P"]...)
	sortedHands = append(sortedHands, sortedGroupedHandsByType["1P"]...)
	sortedHands = append(sortedHands, sortedGroupedHandsByType["HC"]...)

	return sortedHands
}

func (s Day7Solver) giveRank(sortedHands []*Hand) {
	// give rank
	for idx, hand := range sortedHands {
		hand.Rank = len(sortedHands) - idx
		hand.TotalAmount *= hand.BidAmount
	}
}

func (s Day7Solver) calculateWinnings(sortedHands []*Hand) int {
	winnings := 0
	for _, hand := range sortedHands {
		winnings += hand.TotalAmount * hand.Rank
	}
	return winnings
}
