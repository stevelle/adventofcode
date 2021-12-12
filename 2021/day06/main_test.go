package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	// shared memory, prevents running in parallel
	var lines []string
	var initialState []int

	t.Run("ReadFile", func(t *testing.T) {
		lines = ReadFile("test-input.txt")
		assert.Equal(t, 1, len(lines))
	})

	t.Run("ParseFile", func(t *testing.T) {
		initialState = calculateInitialState(lines[0])
	})

	t.Run("PartOne", func(t *testing.T) {
		assert.Equal(t, int64(5934), PartOne(initialState))
	})

	t.Run("PartTwo", func(t *testing.T) {
		assert.Equal(t, int64(26984457539), PartTwo(initialState))
	})
}

func BenchmarkPartOne(b *testing.B) {
	INPUT_FILE := "input.txt"
	lines := ReadFile(INPUT_FILE)
	initialState := calculateInitialState(lines[0])
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		PartOne(initialState)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	INPUT_FILE := "input.txt"
	lines := ReadFile(INPUT_FILE)
	initialState := calculateInitialState(lines[0])
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		PartTwo(initialState)
	}
}
