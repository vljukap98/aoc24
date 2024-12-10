package day8

import (
	"aoc24/util"
	"fmt"
)

func Day8() {
	mat, _ := util.ReadMat("./day8/day8-input.txt")
	var nodes []Node
	for y := 0; y < len(mat); y++ {
		for x := 0; x < len(mat[y]); x++ {
			if mat[y][x] != "." {
				nodes = append(nodes, Node{mat[y][x], x, y})
			}
		}
	}

	h := len(mat)
	w := len(mat[0])

	part1(nodes, w, h)
	part2(mat, nodes, w, h)
}

func part1(nodes []Node, w int, h int) {
	var antinodes []Node
	for _, n1 := range nodes {
		for _, n2 := range nodes {
			if n1.X == n2.X && n1.Y == n2.Y {
				continue
			}
			antinode := n1.GetAntinode(n2, w, h)
			if antinode != nil && !contains(antinodes, *antinode) {
				antinodes = append(antinodes, *antinode)
			}
		}
	}

	fmt.Println(len(antinodes))
}

func part2(mat [][]string, nodes []Node, w int, h int) {
	var antinodes []Node
	for _, n1 := range nodes {
		for _, n2 := range nodes {
			if n1.X == n2.X && n1.Y == n2.Y {
				continue
			}
			newAntinodes := n1.GetAntinodes(n2, w, h)
			for _, antinode := range newAntinodes {
				if !contains(antinodes, antinode) &&
					!contains2(nodes, antinode) {
					antinodes = append(antinodes, antinode)
				}
			}
		}
	}

	fmt.Println(len(antinodes) + len(nodes))
}
