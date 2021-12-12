package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCursor(t *testing.T) {
	t.Run("Reset", func(t *testing.T) {
		line := Line{Point{0, 0}, Point{3, 3}}
		cursor := Cursor{line, nil}

		cursor.reset()

		assert.True(t, cursor.isOpen())
		assert.Equal(t, 0, cursor.position.x)
		assert.Equal(t, 0, cursor.position.y)
	})

	t.Run("ResetOnCallToIter", func(t *testing.T) {
		line := Line{Point{0, 0}, Point{3, 3}}
		cursor := Cursor{line, nil}

		cursor.iter()

		assert.True(t, cursor.isOpen())
		assert.Equal(t, 0, cursor.position.x)
		assert.Equal(t, 0, cursor.position.y)
	})

	t.Run("TraversesDiagonalUpAndRight", func(t *testing.T) {
		line := Line{Point{0, 0}, Point{3, 3}}
		cursor := Cursor{line, &Point{0, 0}}

		cursor.iter()             // (1, 1)
		cursor.iter()             // (2, 2)
		position := cursor.iter() // (3, 3)

		assert.Equal(t, 3, position.x)
		assert.Equal(t, 3, position.y)

		assert.Equal(t, EXHAUSTED, cursor.iter())
		assert.False(t, cursor.isOpen())
	})

	t.Run("TraversesDiagonalDownAndRight", func(t *testing.T) {
		line := Line{Point{0, 3}, Point{3, 0}}
		cursor := Cursor{line, &Point{0, 3}}

		cursor.iter()             // (1, 2)
		cursor.iter()             // (2, 1)
		position := cursor.iter() // (3, 0)

		assert.Equal(t, 3, position.x)
		assert.Equal(t, 0, position.y)

		assert.Equal(t, EXHAUSTED, cursor.iter())
		assert.False(t, cursor.isOpen())
	})

	t.Run("TraversesDiagonalDownAndLeft", func(t *testing.T) {
		line := Line{Point{3, 3}, Point{0, 0}}
		cursor := Cursor{line, &Point{3, 3}}

		cursor.iter()             // (2, 2)
		cursor.iter()             // (1, 1)
		position := cursor.iter() // (0, 0)

		assert.Equal(t, 0, position.x)
		assert.Equal(t, 0, position.y)

		assert.Equal(t, EXHAUSTED, cursor.iter())
		assert.False(t, cursor.isOpen())
	})

	t.Run("TraversesDiagonalUpAndLeft", func(t *testing.T) {
		line := Line{Point{3, 0}, Point{0, 3}}
		cursor := Cursor{line, &Point{3, 0}}

		cursor.iter()             // (2, 1)
		cursor.iter()             // (1, 2)
		position := cursor.iter() // (0, 3)

		assert.Equal(t, 0, position.x)
		assert.Equal(t, 3, position.y)

		assert.Equal(t, EXHAUSTED, cursor.iter())
		assert.False(t, cursor.isOpen())
	})
}
