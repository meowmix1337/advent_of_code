package solutionsv2

import (
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"
)

type Day5Solver struct {
	*BaseSolver
}

func NewDay5Solver(baseSolver *BaseSolver, inputFile *os.File) Solver {
	day := &Day5Solver{
		BaseSolver: baseSolver,
	}
	day.InputFile = inputFile

	return day
}

type Seed int

type Seed2 struct {
	SeedStart int
	SeedRange int
}

type Seeds []Seed

func (s *Seeds) setSeeds(seeds []string) {
	for _, seedStr := range seeds {
		seedInt, _ := strconv.Atoi(seedStr)
		seed := Seed(seedInt)
		*s = append(*s, seed)
	}
}

type DataMap struct {
	LineData []*LineData
}

type LineData struct {
	DestinationStart int
	DestinationEnd   int
	SourceStart      int
	SourceEnd        int
	Length           int
}

func NewDataMap() *DataMap {
	return &DataMap{
		LineData: make([]*LineData, 0),
	}
}

func NewLineData(destinationStart, destinationEnd, sourceStart, sourceEnd, length int) *LineData {
	return &LineData{
		DestinationStart: destinationStart,
		DestinationEnd:   destinationEnd,
		SourceStart:      sourceStart,
		SourceEnd:        sourceEnd,
		Length:           length,
	}
}

func (s *Day5Solver) Solve() *Answers {

	content, err := os.ReadFile("inputfiles/2023/day5.txt")
	if err != nil {
		log.Fatalln("failed to read file", err)
	}

	input := string(content)
	data := splitByEmptyNewline(input)

	seeds := Seeds{}
	maps := make([]*DataMap, 0)
	var seedsTwo []Seed2
	for _, line := range data {
		// get seeds first
		if strings.Contains(line, "seeds:") {
			seedsStr := strings.Trim(strings.Split(line, ":")[1], " ")
			seeds.setSeeds(strings.Split(seedsStr, " "))
			seedsTwo = setSeeds2(strings.Split(seedsStr, " "))
			continue
		}

		mapData := getMap(line)
		maps = append(maps, mapData)
	}

	seedToSoil := maps[0]
	soilToFert := maps[1]
	fertToWater := maps[2]
	waterToLight := maps[3]
	lightToTemp := maps[4]
	tempToHumid := maps[5]
	humidToLocation := maps[6]

	// part 1
	lowestLocation := 0
	for _, seed := range seeds {

		soil := getValue(seedToSoil, int(seed))
		fert := getValue(soilToFert, soil)
		water := getValue(fertToWater, fert)
		light := getValue(waterToLight, water)
		temp := getValue(lightToTemp, light)
		humid := getValue(tempToHumid, temp)
		location := getValue(humidToLocation, humid)

		if lowestLocation == 0 {
			lowestLocation = location
		} else if location < lowestLocation {
			lowestLocation = location
		}
	}

	// part 2
	lowestLocation2 := 0

	// seed and location
	seedMap := make(map[int]int)
	var wg sync.WaitGroup
	for _, seedRange := range seedsTwo {
		wg.Add(1)

		go func(seedRange Seed2) {
			for seed := seedRange.SeedStart; seed <= seedRange.SeedStart+seedRange.SeedRange; seed++ {
				if _, exists := seedMap[seed]; exists {
					continue
				}
				soil := getValue(seedToSoil, seed)
				fert := getValue(soilToFert, soil)
				water := getValue(fertToWater, fert)
				light := getValue(waterToLight, water)
				temp := getValue(lightToTemp, light)
				humid := getValue(tempToHumid, temp)
				location := getValue(humidToLocation, humid)

				if lowestLocation2 == 0 {
					lowestLocation2 = location
				} else if location < lowestLocation2 {
					lowestLocation2 = location
				}
			}
			wg.Done()
		}(seedRange)
	}
	wg.Wait()

	s.Log.WithFields(log.Fields{
		"part 1": lowestLocation,
		"part 2": lowestLocation2,
	}).Info("Answers")

	return &Answers{
		Answer1: lowestLocation,
		Answer2: lowestLocation2,
	}
}

func setSeeds2(seeds []string) []Seed2 {
	seedsTwo := make([]Seed2, 0)

	for i := 0; i < len(seeds); i += 2 {
		seedStart, _ := strconv.Atoi(seeds[i])
		seedRange, _ := strconv.Atoi(seeds[i+1])
		seedsTwo = append(seedsTwo, Seed2{
			SeedStart: seedStart,
			SeedRange: seedRange,
		})
	}

	return seedsTwo
}

func getValue(mapData *DataMap, key int) int {
	// check if key is within range of source
	mapping := -1
	for _, lineData := range mapData.LineData {
		// outside range, don't care about it
		if key < lineData.SourceStart || key > lineData.SourceEnd {
			continue
		}

		delta := lineData.DestinationStart - lineData.SourceStart

		// we know which line has the destination mapping
		mapping = key + delta
	}

	if mapping == -1 {
		return key
	}

	return mapping
}

func splitByEmptyNewline(str string) []string {
	strNormalized := regexp.
		MustCompile("\r\n").
		ReplaceAllString(str, "\n")

	return regexp.
		MustCompile(`\n\s*\n`).
		Split(strNormalized, -1)

}

func getMap(line string) *DataMap {
	// remove header title, so we only have the list of maps
	dataMap := NewDataMap()
	listOfMap := strings.Trim(strings.Split(line, ":")[1], "\n")
	mapData := strings.Split(listOfMap, "\n")
	for _, d := range mapData {
		ranges := strings.Split(d, " ")

		// idx 2 = length
		length, _ := strconv.Atoi(ranges[2])

		// idx 0 = destination start
		destStart, _ := strconv.Atoi(ranges[0])

		// idx 1 = source start
		sourceStart, _ := strconv.Atoi(ranges[1])

		dataMap.LineData = append(dataMap.LineData, NewLineData(destStart, destStart+length-1, sourceStart, sourceStart+length-1, length))
	}

	return dataMap
}
