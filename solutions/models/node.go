package models

type Node struct {
	ID        int
	X         int
	Y         int
	IsTail    bool
	Parent    *Node
	ChildNode *Node
}

func (h *Node) Move(direction string, mapArea Map) {
	switch direction {
	case "R":
		h.X++
	case "U":
		h.Y++
	case "L":
		h.X--
	case "D":
		h.Y--
	}

	// check if current location exists in map
	// if not exist, create a new location
	if location := mapArea.GetLocation(h.X, h.Y); location == nil {
		mapArea.AddLocation(h.X, h.Y, false)
	}

	h.ChildNode.MoveIfNeeded(mapArea)
}

// possible moves
// C = child, P = parent
// --------Move child to left of parent-----------
// .....      .....
// C....      .....
// C.P..  ->  .CP..
// C....      .....
// .....      .....
// ---------Move child to bottom of parent----------
// .....      .....
// .....      .....
// ..P..  ->  ..P..
// .....      ..C..
// .CCC.      .....
// ---------Move child to right of parent----------
// .....      .....
// ....C      .....
// ..P.C  ->  ..PC.
// ....C      .....
// .....      .....
// ---------Move child to top of parent----------
// .CCC.      .....
// .....      ..C..
// ..P..  ->  ..P..
// .....      .....
// .....      .....
// ---------Move child to bottom-left of parent----------
// .....      .....
// .....      .....
// ..P..  ->  ..P..
// .....      .C...
// C....      .....
// ---------Move child to bottom-right of parent----------
// .....      .....
// .....      .....
// ..P..  ->  ..P..
// .....      ...C.
// ....C      .....
// ---------Move child to top-right of parent----------
// ....C      .....
// .....      ...C.
// ..P..  ->  ..P..
// .....      .....
// .....      .....
// ----------Move child to top-left of parent---------
// C....      .....
// .....      .C...
// ..P..  ->  ..P..
// .....      .....
// .....      .....
// -------------------
func (n *Node) MoveIfNeeded(mapArea Map) {

	curNode := n
	for curNode != nil {
		// move to left side of head
		if curNode.MoveToLeftOfParent() {
			curNode.X = curNode.Parent.X - 1
			curNode.Y = curNode.Parent.Y
			// move to bottom side of head
		} else if curNode.MoveToBottomOfParent() {
			curNode.X = curNode.Parent.X
			curNode.Y = curNode.Parent.Y - 1
			// move to right side
		} else if curNode.MoveToRightOfParent() {
			curNode.X = curNode.Parent.X + 1
			curNode.Y = curNode.Parent.Y
			// move to top side
		} else if curNode.MoveToTopOfParent() {
			curNode.X = curNode.Parent.X
			curNode.Y = curNode.Parent.Y + 1
			// move to bottom left
		} else if curNode.MoveToBottomLeftOfParent() {
			curNode.X = curNode.Parent.X - 1
			curNode.Y = curNode.Parent.Y - 1
			// move to bottom right
		} else if curNode.MoveToBottomRightOfParent() {
			curNode.X = curNode.Parent.X + 1
			curNode.Y = curNode.Parent.Y - 1
			// move to top right
		} else if curNode.MoveToTopRightOfParent() {
			curNode.X = curNode.Parent.X + 1
			curNode.Y = curNode.Parent.Y + 1
			// move to top left
		} else if curNode.MoveToTopLeftOfParent() {
			curNode.X = curNode.Parent.X - 1
			curNode.Y = curNode.Parent.Y + 1
		}

		// we only keep track of the tail
		// check if location exists in map
		// if not, create a new location and mark it as visisted
		// if exists, mark location as visited
		if curNode.IsTail {
			if location := mapArea.GetLocation(curNode.X, curNode.Y); location == nil {
				mapArea.AddLocation(curNode.X, curNode.Y, true)
			} else {
				location.Visited = true
			}
		}

		// move on to next child if any
		curNode = curNode.ChildNode
	}
}

func (n *Node) MoveToLeftOfParent() bool {
	return (n.X == n.Parent.X-2 && n.Y == n.Parent.Y) || // 2 left
		(n.X == n.Parent.X-2 && n.Y == n.Parent.Y-1) || // 2 left 1 bottom
		(n.X == n.Parent.X-2 && n.Y == n.Parent.Y+1) // 2 left 1 top
}

func (n *Node) MoveToBottomOfParent() bool {
	return (n.X == n.Parent.X && n.Y == n.Parent.Y-2) || // 2 bottom
		(n.X == n.Parent.X-1 && n.Y == n.Parent.Y-2) || // 2 bottom 1 left
		(n.X == n.Parent.X+1 && n.Y == n.Parent.Y-2) // 2 bottom 1 right
}

func (n *Node) MoveToRightOfParent() bool {
	return (n.X == n.Parent.X+2 && n.Y == n.Parent.Y) || // 2 right
		(n.X == n.Parent.X+2 && n.Y == n.Parent.Y-1) || // 2 right 1 bottom
		(n.X == n.Parent.X+2 && n.Y == n.Parent.Y+1) // 2 right 1 top
}

func (n *Node) MoveToTopOfParent() bool {
	return (n.X == n.Parent.X && n.Y == n.Parent.Y+2) || // 2 top
		(n.X == n.Parent.X+1 && n.Y == n.Parent.Y+2) || // 2 top 1 right
		(n.X == n.Parent.X-1 && n.Y == n.Parent.Y+2) // 2 top 1 left
}

func (n *Node) MoveToBottomLeftOfParent() bool {
	return n.X == n.Parent.X-2 && n.Y == n.Parent.Y-2
}

func (n *Node) MoveToBottomRightOfParent() bool {
	return n.X == n.Parent.X+2 && n.Y == n.Parent.Y-2
}

func (n *Node) MoveToTopRightOfParent() bool {
	return n.X == n.Parent.X+2 && n.Y == n.Parent.Y+2
}

func (n *Node) MoveToTopLeftOfParent() bool {
	return n.X == n.Parent.X-2 && n.Y == n.Parent.Y+2
}
