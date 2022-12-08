package models

type DirectoryStack []*Directory

func (s *DirectoryStack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *DirectoryStack) Push(directory *Directory) {
	*s = append(*s, directory)
}

func (s *DirectoryStack) Pop() *Directory {
	if s.IsEmpty() {
		return &Directory{}
	} else {
		idx := len(*s) - 1     // last idx is top
		directory := (*s)[idx] // get element
		*s = (*s)[:idx]        // remove the last element by doing some slicing
		return directory
	}
}

func (s *DirectoryStack) Peek() *Directory {
	return (*s)[len(*s)-1] // get element
}

func (s *DirectoryStack) Seek(howFarBack int) *Directory {
	if len(*s)-howFarBack < 0 {
		return nil
	}
	return (*s)[len(*s)-howFarBack] // get element
}

func (s *DirectoryStack) Size() int {
	return len(*s)
}
