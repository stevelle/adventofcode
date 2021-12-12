package main

type Grid map[Point]int

func (grid *Grid) draw(line Line) {
	cursor := Cursor{line, nil}
	for cursor.reset(); cursor.isOpen(); cursor.iter() {
		grid.incr(*cursor.position)
	}
}

func (grid *Grid) incr(point Point) {
	(*grid)[point] += 1
}

func (grid *Grid) scoreGrid() int {
	sum := 0

	for _, score := range *grid {
		if score > 1 {
			sum++
		}
	}
	return sum
}
