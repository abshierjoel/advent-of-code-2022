package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReturnsCorrectSeenSpacesForSampleLen2(t *testing.T) {
	result := knottyBoy(sampleInput(), 2)

	assert.Equal(t, result, 13)
}

func TestReturnsCorrectSeenSpacesForSampleLen10(t *testing.T) {
	result := knottyBoy(sampleInput(), 10)

	assert.Equal(t, result, 1)
}

func TestMegaSampleReturnsCorrectSeenSpaces(t *testing.T) {
	result := knottyBoy(megaSample(), 36)

	assert.Equal(t, result, 1)
}

func sampleInput() []string {
	return []string{
		"R 4",
		"U 4",
		"L 3",
		"D 1",
		"R 4",
		"D 1",
		"L 5",
		"R 2"}
}

func megaSample() []string {
	return []string{"R 5",
		"U 8",
		"L 8",
		"D 3",
		"R 17",
		"D 10",
		"L 25",
		"U 20"}
}
