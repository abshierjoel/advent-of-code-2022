package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsInOrderReturnsPositioning(t *testing.T) {
	pairs := []Pair{
		{
			PacketA: &Packet{-1, []*Packet{{Value: 5}}, nil},
			PacketB: &Packet{-1, []*Packet{{Value: 10}}, nil},
		},
		{
			PacketA: &Packet{10, []*Packet{}, nil},
			PacketB: &Packet{20, []*Packet{}, nil},
		}}

	var results []int
	for i, pair := range pairs {
		if isInOrder(pair.PacketA, pair.PacketB) >= 0 {
			results = append(results, i)
		}
	}

	expected := []int{0, 1}
	assert.Equal(t, expected, results)

}

func TestIsInOrder(t *testing.T) {
	pairs := parsePackets(sample())

	var results []int
	for i, pair := range pairs {
		if isInOrder(pair.PacketA, pair.PacketB) >= 0 {
			results = append(results, i+1)
		}
	}

	expected := []int{1, 2, 4, 6}
	assert.Equal(t, expected, results)

}

func TestParsePackets(t *testing.T) {
	result := parsePackets(sample())

	assert.Equal(t, 7, len(result))
}

func TestParsePacket(t *testing.T) {
	input := "[1,2,3]"
	result := parsePacket(input)

	assert.Equal(t, 1, result.Elements[0].Value)
	assert.Equal(t, 2, result.Elements[1].Value)
	assert.Equal(t, 3, result.Elements[2].Value)
}
func TestParseDeepPacket(t *testing.T) {
	input := "[1,[2,[3,[4,[5,6,7]]]],8,9]"
	result := parsePacket(input)

	inner1 := result.Elements[1]
	inner2 := inner1.Elements[1]
	inner3 := inner2.Elements[1]
	inner4 := inner3.Elements[1]

	assert.Equal(t, 1, result.Elements[0].Value)
	assert.Equal(t, 2, inner1.Elements[0].Value)
	assert.Equal(t, 3, inner2.Elements[0].Value)
	assert.Equal(t, 4, inner3.Elements[0].Value)
	assert.Equal(t, 5, inner4.Elements[0].Value)
	assert.Equal(t, 6, inner4.Elements[1].Value)
	assert.Equal(t, 7, inner4.Elements[2].Value)
	assert.Equal(t, 8, result.Elements[2].Value)
	assert.Equal(t, 9, result.Elements[3].Value)
}

func sample() []string {
	return []string{
		"[1,1,3,1,1]",
		"[1,1,5,1,1]",
		"",
		"[[1],[2,3,4]]",
		"[[1],4]",
		"",
		"[9]",
		"[[8,7,6]]",
		"",
		"[[4,4],4,4]",
		"[[4,4],4,4,4]",
		"",
		"[7,7,7,7]",
		"[7,7,7]",
		"",
		"[]",
		"[3]",
		"",
		"[[[]]]",
		"[[]]",
		"",
		"[1,[2,[3,[4,[5,6,7]]]],8,9]",
		"[1,[2,[3,[4,[5,6,0]]]],8,9]",
	}
}
