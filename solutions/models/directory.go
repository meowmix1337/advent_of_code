package models

type File struct {
	Filename string `json:"filename"`
	Size     int    `json:"size"`
}

type Directory struct {
	TotalSize   int                   `json:"totalSize"`
	Name        string                `json:"name"`
	Files       []File                `json:"files"`
	Directories map[string]*Directory `json:"directories"`
}

func NewDirectory(name string) *Directory {
	return &Directory{
		Name:        name,
		Files:       make([]File, 0),
		Directories: make(map[string]*Directory),
		TotalSize:   0,
	}
}

func (d *Directory) SumAtMost100000AndFindRequired(required int) (int, []int) {
	totalSize := 0
	greaterThanRequired := make([]int, 0)

	return Traverse(*d, totalSize, required, greaterThanRequired)
}

func Traverse(dir Directory, totalSize, required int, greaterThanRequired []int) (int, []int) {
	if dir.TotalSize <= 100000 {
		totalSize += dir.TotalSize
	}

	if dir.TotalSize >= required {
		greaterThanRequired = append(greaterThanRequired, dir.TotalSize)
	}

	if dir.hasDirectories() {
		for _, dir := range dir.Directories {
			totalSize, greaterThanRequired = Traverse(*dir, totalSize, required, greaterThanRequired)
		}
	}

	return totalSize, greaterThanRequired
}

func (d *Directory) hasDirectories() bool {
	return len(d.Directories) > 0
}
