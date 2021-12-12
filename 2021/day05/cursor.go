package main

import "math"

var EXHAUSTED = Point{math.MinInt, math.MinInt}

type Cursor struct {
	line     Line
	position *Point
}

func (cursor *Cursor) iter() Point {
	if cursor.position == nil {
		cursor.reset()
		return *cursor.position
	}
	if *cursor.position == cursor.line.endpoint() {
		cursor.position = &EXHAUSTED
		return EXHAUSTED
	}

	if cursor.position.x < cursor.line.x2() {
		cursor.position.x++
	} else if cursor.position.x > cursor.line.x2() {
		cursor.position.x--
	}

	if cursor.position.y < cursor.line.y2() {
		cursor.position.y++
	} else if cursor.position.y > cursor.line.y2() {
		cursor.position.y--
	}

	return *cursor.position
}

func (cursor *Cursor) reset() Point {
	start := cursor.line.startpoint()
	cursor.position = &start
	return *cursor.position
}

func (cursor *Cursor) isOpen() bool {
	return cursor.position != nil && cursor.position != &EXHAUSTED
}
