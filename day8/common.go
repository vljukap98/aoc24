package day8

type Node struct {
	Char string
	X    int
	Y    int
}

func (n *Node) GetAntinode(node Node, w int, h int) *Node {
	if node.Char != n.Char {
		return nil
	}
	dx := node.X - n.X
	x := node.X + dx
	dy := node.Y - n.Y
	y := node.Y + dy

	if x < 0 {
		return nil
	}
	if y < 0 {
		return nil
	}
	if x >= w {
		return nil
	}
	if y >= h {
		return nil
	}

	return &Node{"#", x, y}
}

func (n *Node) GetAntinodes(node Node, w int, h int) []Node {
	if node.Char != n.Char {
		return make([]Node, 0)
	}
	dx := node.X - n.X
	dy := node.Y - n.Y
	x := node.X + dx
	y := node.Y + dy

	var antinodes []Node
	for x > -1 && y > -1 && x < w && y < h {
		antinodes = append(antinodes, Node{"#", x, y})
		x += dx
		y += dy
	}

	return antinodes
}

func contains(s []Node, e Node) bool {
	for _, a := range s {
		if a.Char == e.Char && a.X == e.X && a.Y == e.Y {
			return true
		}
	}
	return false
}

func contains2(s []Node, e Node) bool {
	for _, a := range s {
		if a.X == e.X && a.Y == e.Y {
			return true
		}
	}
	return false
}
