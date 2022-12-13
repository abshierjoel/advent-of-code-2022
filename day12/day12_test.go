package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBFSReturnsCorrectValueFromSingleStart(t *testing.T) {
	grid, end := parseGrid(sample())
	grid = setNeighbors(grid)
	result := bfs('S', end, grid)

	assert.Equal(t, 31, result)
}

func TestBFSReturnsCorrectValueFromMultipleStarts(t *testing.T) {
	grid, end := parseGrid(sample())
	grid = setNeighbors(grid)
	result := bfs('a', end, grid)

	assert.Equal(t, 29, result)
}

func sample() []string {
	return []string{
		"Sabqponm",
		"abcryxxl",
		"accszExk",
		"acctuvwj",
		"abdefghi",
	}
}
