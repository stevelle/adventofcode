package main

type Point struct {
	x int
	y int
}

type Line struct {
	start Point
	end   Point
}

func (line *Line) x1() int {
	return line.start.x
}

func (line *Line) x2() int {
	return line.end.x
}

func (line *Line) y1() int {
	return line.start.y
}

func (line *Line) y2() int {
	return line.end.y
}

func (line *Line) isHorizontal() bool {
	return line.y1() == line.y2()
}

func (line *Line) isVertical() bool {
	return line.x1() == line.x2()
}

func (line *Line) startpoint() Point {
	return Point{line.x1(), line.y1()}
}

func (line *Line) endpoint() Point {
	return Point{line.x2(), line.y2()}
}
