package models

type Tree struct {
	Height int `json:"height"`
	X      int `json:"x"`
	Y      int `json:"y"`
}

type Forest struct {
	EdgeCount int      `json:"edgeCount"`
	Rows      int      `json:"rows"`
	Cols      int      `json:"cols"`
	Trees     [][]Tree `json:"trees"`
}

func (f *Forest) CountEdge() int {
	return f.Rows*f.Cols - (f.Rows-2)*(f.Cols-2)
}

func (f *Forest) CountVisibleTrees() int {
	visibleTrees := f.CountEdge()
	for _, treeRow := range f.Trees {
		for _, tree := range treeRow {
			if tree.isEdge(f.Rows, f.Cols) {
				continue
			}

			// check visibility on each tree
			if tree.IsVisible(*f) {
				visibleTrees++
			}
		}
	}

	return visibleTrees
}

func (f *Forest) GetHighestScenicScore() int {
	highestScore := 0

	for _, treeRow := range f.Trees {
		for _, tree := range treeRow {
			scenicScore := tree.CalculateScore(*f)

			if scenicScore > highestScore {
				highestScore = scenicScore
			}
		}
	}

	return highestScore
}

func (t *Tree) isEdge(forestMaxY, forestMaxX int) bool {
	return (t.Y == 0 || t.Y == forestMaxY-1) ||
		(t.X == 0 || t.X == forestMaxX-1)
}

func (t *Tree) CalculateScore(f Forest) int {
	_, topScore := t.IsVisibleTop(f, t.X, t.Y)
	_, rightScore := t.IsVisibleRight(f, t.X, t.Y)
	_, bottomScore := t.IsVisibleBottom(f, t.X, t.Y)
	_, leftScore := t.IsVisibleLeft(f, t.X, t.Y)

	return topScore * rightScore * bottomScore * leftScore
}

func (t *Tree) IsVisible(forest Forest) bool {

	topVis, _ := t.IsVisibleTop(forest, t.X, t.Y)
	rightVis, _ := t.IsVisibleRight(forest, t.X, t.Y)
	bottomVis, _ := t.IsVisibleBottom(forest, t.X, t.Y)
	leftVis, _ := t.IsVisibleLeft(forest, t.X, t.Y)

	if !topVis &&
		!rightVis &&
		!bottomVis &&
		!leftVis {
		return false
	}

	return true
}

func (t *Tree) IsVisibleTop(forest Forest, x, y int) (bool, int) {
	score := 0
	// check top y-1
	for y != 0 {
		// always increase the score, we'll return once we determine if we
		// hit the edge or a tree that is greater than equal to our height
		score++
		// We're not visible
		if t.Height <= forest.Trees[y-1][x].Height {
			return false, score
		}
		y--
	}

	return true, score
}

func (t *Tree) IsVisibleRight(forest Forest, x, y int) (bool, int) {
	score := 0
	// check top x+1
	for x != forest.Rows-1 {
		// always increase the score, we'll return once we determine if we
		// hit the edge or a tree that is greater than equal to our height
		score++
		// We're not visible
		if t.Height <= forest.Trees[y][x+1].Height {
			return false, score
		}
		x++
	}

	return true, score
}

func (t *Tree) IsVisibleBottom(forest Forest, x, y int) (bool, int) {
	score := 0
	// check top y+1
	for y != forest.Rows-1 {
		// always increase the score, we'll return once we determine if we
		// hit the edge or a tree that is greater than equal to our height
		score++
		// We're not visible
		if t.Height <= forest.Trees[y+1][x].Height {
			return false, score
		}
		y++
	}

	return true, score
}

func (t *Tree) IsVisibleLeft(forest Forest, x, y int) (bool, int) {
	score := 0
	// check top x-1
	for x != 0 {
		// always increase the score, we'll return once we determine if we
		// hit the edge or a tree that is greater than equal to our height
		score++
		// We're not visible
		if t.Height <= forest.Trees[y][x-1].Height {
			return false, score
		}
		x--
	}

	return true, score
}
