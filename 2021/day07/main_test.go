package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	// shared memory, prevents running in parallel
	var lines []string
	var initialPositions []int

	t.Run("ReadFile", func(t *testing.T) {
		lines = ReadFile("test-input.txt")
		assert.Equal(t, 1, len(lines))
	})

	t.Run("ParseFile", func(t *testing.T) {
		initialPositions = asInts(strings.Split(lines[0], ","))
	})

	t.Run("PartOne", func(t *testing.T) {
		_, cost := PartOne(initialPositions)
		assert.Equal(t, 37, cost)
	})

	t.Run("PartTwo", func(t *testing.T) {
		_, cost := PartTwo(initialPositions)
		assert.Equal(t, 168, cost)
	})
}

func BenchmarkPartOne(b *testing.B) {
	INPUT_FILE := "input.txt"
	lines := ReadFile(INPUT_FILE)
	initialPositions := asInts(strings.Split(lines[0], ","))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		PartOne(initialPositions)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	INPUT_FILE := "input.txt"
	lines := ReadFile(INPUT_FILE)
	initialPositions := asInts(strings.Split(lines[0], ","))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		PartTwo(initialPositions)
	}
}
