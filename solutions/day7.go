package solutions

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"advent/solutions/models"
)

type Day7 struct {
	InputFile string
	Logger    *log.Logger
}

func NewDay7Solver(inputFile string, logger *log.Logger) Solver {
	day7Solver := new(Day7)
	day7Solver.InputFile = inputFile
	day7Solver.Logger = logger

	return day7Solver
}

func (d *Day7) Solve() (*Answers, error) {
	file, err := os.Open(d.InputFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	root := &models.Directory{
		Name:        "/",
		Files:       make([]models.File, 0),
		Directories: make(map[string]*models.Directory),
		TotalSize:   0,
		Parent:      nil,
	}

	rootDir := BuildFileStructure(file, root)

	// total needed to meet 30,000,000
	// total space 70,000,000
	totalSpaceLeft := 70000000 - rootDir.TotalSize
	required := 30000000 - totalSpaceLeft

	answer1, greaterThanRequired := rootDir.SumAtMost100000AndFindRequired(required)

	sort.Ints(greaterThanRequired)
	answer2 := greaterThanRequired[0]

	return &Answers{
		Answer1:  answer1,
		Answer2:  answer2,
		MetaData: rootDir,
	}, nil
}

func BuildFileStructure(file *os.File, root *models.Directory) *models.Directory {
	currentDir := root

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		cmd := strings.Split(line, " ")
		if cmd[0] == "$" {
			if cmd[1] == "cd" {
				if cmd[2] == "/" {
					currentDir = root
				} else if cmd[2] == ".." {
					// go back one
					currentDir = currentDir.Parent
				} else {
					// set parent before we move down one level
					currentDir.Directories[cmd[2]].Parent = currentDir
					currentDir = currentDir.Directories[cmd[2]]
				}
			}

			// do LS action...noop
			continue
		}

		if cmd[0] == "dir" {
			// create blank directory
			currentDir.Directories[cmd[1]] = models.NewDirectory(cmd[1])
		} else {
			// add the file to current directory
			size, _ := strconv.Atoi(cmd[0])
			file := models.File{
				Size:     size,
				Filename: cmd[1],
			}
			currentDir.Files = append(currentDir.Files, file)
			currentDir.TotalSize += size

			// keep adding the size to parent directories until we get to root (add size to root though)
			temp := currentDir
			for temp.Parent != nil {
				parent := temp.Parent
				parent.TotalSize += size
				temp = temp.Parent
			}
		}
	}

	return root
}
