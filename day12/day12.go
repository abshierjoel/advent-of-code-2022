package main

import (
	"fmt"

	"github.com/abshierjoel/advent-of-code-2022/utils"
)

const filename = "mountains.txt"

type Node struct {
	Up      *Node
	Right   *Node
	Down    *Node
	Left    *Node
	Raw     rune
	Height  rune
	Visited bool
}

func main() {
	lines := utils.ReadLines(filename)
	grid, end := parseGrid(lines)
	grid = setNeighbors(grid)

	// Part 1
	pathLength := bfs('S', end, grid)
	fmt.Printf("The minimum required moves is %d\n", pathLength)

	// Part 2
	pathLengthB := bfs('a', end, grid)
	fmt.Printf("The minimum required moves is %d\n", pathLengthB)
}

func parseGrid(lines []string) ([][]*Node, *Node) {
	var grid [][]*Node
	var end *Node

	for _, line := range lines {
		var row []*Node

		for _, char := range line {
			node := Node{Height: char, Raw: char}
			if char == 'S' {
				node.Height = 'a'
			}
			if char == 'E' {
				node.Height = 'z'
				end = &node
			}
			row = append(row, &node)
		}

		grid = append(grid, row)
	}

	return grid, end
}

func findStartPoints(grid [][]*Node, key rune) []*Node {
	var startPoints []*Node
	for _, p := range grid {
		for _, r := range p {
			if r.Raw == key {
				startPoints = append(startPoints, r)
			}
		}
	}
	return startPoints
}

func bfs(startChar rune, end *Node, grid [][]*Node) int {
	var node *Node
	start := findStartPoints(grid, startChar)
	queue := start
	visited := start
	dist := make(map[*Node]int)

	for len(queue) > 0 {
		node, queue = pop(queue)
		node.visit()
		if node == end {
			break
		}

		for _, neighbor := range node.getNeighbors() {
			if !contains(visited, neighbor) {
				visited = append(visited, neighbor)
				queue = append(queue, neighbor)
				dist[neighbor] += dist[node] + 1
			}
		}
	}

	return dist[end]
}

func contains(list []*Node, node *Node) bool {
	for _, v := range list {
		if v == node {
			return true
		}
	}

	return false
}

func pop(queue []*Node) (*Node, []*Node) {
	return queue[0], queue[1:]
}

func (node *Node) visit() {
	node.Visited = true
}

func (node *Node) getNeighbors() []*Node {
	var neighbors []*Node
	if node.Right != nil {
		neighbors = append(neighbors, node.Right)
	}
	if node.Up != nil {
		neighbors = append(neighbors, node.Up)
	}
	if node.Left != nil {
		neighbors = append(neighbors, node.Left)
	}
	if node.Down != nil {
		neighbors = append(neighbors, node.Down)
	}

	return neighbors
}

func setNeighbors(grid [][]*Node) [][]*Node {
	for i, row := range grid {
		for j, node := range row {
			if j > 0 && grid[i][j-1].Height <= node.Height+1 {
				node.Left = grid[i][j-1]
			}
			if i > 0 && grid[i-1][j].Height <= node.Height+1 {
				node.Up = grid[i-1][j]
			}
			if j < len(row)-1 && grid[i][j+1].Height <= node.Height+1 {
				node.Right = grid[i][j+1]
			}
			if i < len(grid)-1 && grid[i+1][j].Height <= node.Height+1 {
				node.Down = grid[i+1][j]
			}
		}
	}
	return grid
}
