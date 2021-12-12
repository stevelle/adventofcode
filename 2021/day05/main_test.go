package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	// shared memory, prevents running in parallel
	var lines []string

	t.Run("ReadFile", func(t *testing.T) {
		lines = ReadFile("test-input.txt")
		assert.Equal(t, 10, len(lines))
	})

	t.Run("PartOne", func(t *testing.T) {
		assert.Equal(t, 5, PartOne(lines))
	})

	t.Run("PartTwo", func(t *testing.T) {
		assert.Equal(t, 12, PartTwo(lines))
	})
}

func BenchmarkPartOne(b *testing.B) {
	INPUT_FILE := "input.txt"
	lines := ReadFile(INPUT_FILE)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		PartOne(lines)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	INPUT_FILE := "input.txt"
	lines := ReadFile(INPUT_FILE)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		PartTwo(lines)
	}
}
