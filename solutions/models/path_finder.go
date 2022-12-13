package models

import "strings"

type Point struct {
	X int
	Y int
}

func NewPoint(x, y int) Point {
	return Point{X: x, Y: y}
}

func (p *Point) Add(p2 Point) Point {
	return Point{
		X: p.X + p2.X,
		Y: p.Y + p2.Y,
	}
}

type PointQueue []Point

func NewPointQueue() *PointQueue {
	return &PointQueue{}
}

func (q *PointQueue) Enqueue(p Point) {
	*q = append(*q, p)
}

func (q *PointQueue) Dequeue() Point {
	point := (*q)[0]
	*q = (*q)[1:]
	return point
}

func (q *PointQueue) IsEmpty() bool {
	return len(*q) == 0
}

type PathFinder struct {
	// a height map
	HeightMap map[Point]rune `json:"heightMap"`
	// keeps track of the shortest point to end for best signal
	ShortestPath *Point `json:"shortestPath"`
	// keeps track of the distances from each point to end
	Distances map[Point]int `json:"distances"`
	// the queue for path finding
	Queue      PointQueue `json:"pointQueue"`
	StartPoint Point      `json:"startPoint"`
	EndPoint   Point      `json:"endPoint"`
}

func NewPathFinder() *PathFinder {
	return &PathFinder{
		HeightMap:    make(map[Point]rune),
		ShortestPath: nil,
		Distances:    make(map[Point]int),
		Queue:        make([]Point, 0),
	}
}

func (pf *PathFinder) BuildHeightMap(input string) {

	// iterate through each row and individual rune
	// build a height map so we have easy access to heights
	for y, row := range strings.Fields(input) {
		for x, char := range row {
			pf.HeightMap[NewPoint(x, y)] = char

			// set the starting and end points
			if char == 'S' {
				pf.StartPoint = NewPoint(x, y)
			} else if char == 'E' {
				pf.EndPoint = NewPoint(x, y)
			}
		}
	}

	// we know start is a and end is z, set those heights
	pf.HeightMap[pf.StartPoint] = 'a'
	pf.HeightMap[pf.EndPoint] = 'z'

	// we start at the end, make our way to the start
	pf.Distances[pf.EndPoint] = 0

	pf.Queue = append(pf.Queue, pf.EndPoint)
}

func (pf *PathFinder) HasSeen(nextPoint Point) bool {
	_, seen := pf.Distances[nextPoint]
	return seen
}

func (pf *PathFinder) IsHeightValid(nextPoint Point) bool {
	_, valid := pf.HeightMap[nextPoint]
	return valid
}

func (pf *PathFinder) PathFind() (int, int) {
	// work our way from end to start
	// last node we get to should be the start so we'll have added the total distance from end to start
	for !pf.Queue.IsEmpty() {
		curPoint := pf.Queue.Dequeue()

		// keep a pointer to `a` so we can keep track of the shortest path
		if pf.HeightMap[curPoint] == 'a' && pf.ShortestPath == nil {
			pf.ShortestPath = &curPoint
		}

		// go through each neighbor
		for _, neighbor := range []Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
			nextPoint := curPoint.Add(neighbor)

			// if this is a new point and is a valid height AND the current point's height is less than or equal to the next height (+1),
			// we can increment the current point + 1 to the next point's distance
			if !pf.HasSeen(nextPoint) && pf.IsHeightValid(nextPoint) && pf.HeightMap[curPoint] <= pf.HeightMap[nextPoint]+1 {
				// add the distance
				pf.Distances[nextPoint] = pf.Distances[curPoint] + 1
				pf.Queue.Enqueue(nextPoint)
			}
		}
	}

	return pf.Distances[pf.StartPoint], pf.Distances[*pf.ShortestPath]
}
