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

	directoryStack := models.DirectoryStack{}

	startingStack := models.Directory{
		Name:        "/",
		Files:       make([]models.File, 0),
		Directories: make(map[string]*models.Directory),
		TotalSize:   0,
	}

	directoryStack.Push(&startingStack)
	rootDir := BuildFileStructure(file, directoryStack)

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

func BuildFileStructure(file *os.File, directoryStack models.DirectoryStack) *models.Directory {
	currentDir := directoryStack.Peek()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		cmd := strings.Split(line, " ")
		if cmd[0] == "$" {
			if cmd[1] == "cd" {
				// we don't care about root
				if cmd[2] == "/" {
					continue
				}

				if cmd[2] == ".." {
					// if .., pop from the stack
					directoryStack.Pop()
				} else {
					// go deeper from current directory
					directoryStack.Push(currentDir.Directories[cmd[2]])
				}
				// set the current directory to top of stack
				currentDir = directoryStack.Peek()
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

			totalStackSize := directoryStack.Size()

			for totalStackSize != 1 {
				// keep going back and add the size until we get to root directory (include root though)
				previousDirectory := directoryStack.Seek(totalStackSize)
				previousDirectory.TotalSize += size
				totalStackSize--
			}

		}
	}

	currentDir = directoryStack[0]
	return currentDir
}
