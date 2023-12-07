package solutionsv2

import (
	"log"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

type Day6Solver struct {
	*BaseSolver
}

func NewDay6Solver(baseSolver *BaseSolver, inputFile *os.File) Solver {
	day := &Day6Solver{
		BaseSolver: baseSolver,
	}
	day.InputFile = inputFile

	return day
}

type Race struct {
	Time     int
	Distance int
}

func NewRace(time, distance int) *Race {
	return &Race{
		Time:     time,
		Distance: distance,
	}
}

func (r *Race) GetDifferentWinnings() int {
	// find minimum first
	minTimeHeld := -1
	// start at 1 since 0 will never win
	for milliseconds := 1; milliseconds < r.Time; milliseconds++ {
		distance := milliseconds * (r.Time - milliseconds)
		if distance > r.Distance {
			minTimeHeld = milliseconds
			break
		}
	}

	// raceTime - minTimeHeld = maxTimeHeld
	maxTimeHeld := r.Time - minTimeHeld
	// (maxTimeHeld - minTimeHeld) + 1 = totalDiffTimesHeld
	totalDiffTimesHeldToWin := (maxTimeHeld - minTimeHeld) + 1
	return totalDiffTimesHeldToWin
}

type Races []*Race

func (rs *Races) AppendRace(race *Race) {
	*rs = append(*rs, race)
}

func (s Day6Solver) Solve() *Answers {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	content, err := os.ReadFile("inputfiles/2023/day6.txt")
	if err != nil {
		log.Fatalln("failed to read file", err)
	}

	input := string(content)
	data := strings.Split(input, "\n")

	total1 := 1
	total2 := 1

	// idx 0 = time
	// idx 1 = distance to beat
	//Races := make([]Game, 0)
	timeData := strings.Trim(strings.Split(data[0], ":")[1], " ")
	distanceData := strings.Trim(strings.Split(data[1], ":")[1], " ")

	times := strings.Fields(timeData)
	distances := strings.Fields(distanceData)

	timePart2, _ := strconv.Atoi(strings.Replace(timeData, " ", "", -1))
	distancePart2, _ := strconv.Atoi(strings.Replace(distanceData, " ", "", -1))

	racePart2 := NewRace(timePart2, distancePart2)

	races := make(Races, 0)
	for i := 0; i < len(times); i++ {
		time, _ := strconv.Atoi(times[i])
		distance, _ := strconv.Atoi(distances[i])

		race := NewRace(time, distance)
		races.AppendRace(race)
	}

	for _, race := range races {
		// find minimum first
		total1 *= race.GetDifferentWinnings()
	}

	total2 = racePart2.GetDifferentWinnings()

	logger.Info("answers", "part 1", total1, "part 2", total2)

	return &Answers{
		Answer1: total1,
		Answer2: total2,
	}
}
