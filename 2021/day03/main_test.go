package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	// shared memory, prevents running in parallel
	var lines []string

	t.Run("ReadFile", func(t *testing.T) {
		lines = ReadFile("03-test-input.txt")
		assert.Equal(t, 12, len(lines))
	})

	t.Run("ParseBinary", func(t *testing.T) {
		assert.Equal(t, int64(22), ParseBinary("10110"))
		assert.Equal(t, int64(9), ParseBinary("01001"))
	})

	t.Run("Bitwise Not", func(t *testing.T) {
		assert.Equal(t, "01", Not("10"))
		assert.Equal(t, "110", Not("001"))
	})

	t.Run("PartOne", func(t *testing.T) {
		assert.Equal(t, int64(198), PartOne(lines))
	})

	t.Run("PartTwo", func(t *testing.T) {
		assert.Equal(t, int64(230), PartTwo(lines))
	})
}

func BenchmarkPartOne(b *testing.B) {
	INPUT_FILE := "03-input.txt"
	lines := ReadFile(INPUT_FILE)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		PartOne(lines)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	INPUT_FILE := "03-input.txt"
	lines := ReadFile(INPUT_FILE)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		PartTwo(lines)
	}
}
