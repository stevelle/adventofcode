package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type Point struct {
	x int
	y int
}

type Range Point

type PointSet map[Point]struct{}

var EXISTS = struct{}{}

func main() {
	rows := ReadFile("input.txt")

	heightMap := generateMap(rows)
	lowPoints := findLowPoints(heightMap)

	score := scorePartOne(lowPoints)
	fmt.Println("** Result A: ", score)

	score = PartTwo(heightMap, lowPoints)
	fmt.Println("** Result B: ", score)
}

func ReadFile(inputName string) []string {
	file, err := os.Open(inputName)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func PartTwo(heightMap [][]int, lowPoints map[Point]int) int {
	largestBasins := make([]int, 0)
	for point := range lowPoints {
		size := sizeOfBasin(point, heightMap)
		largestBasins = append(largestBasins, size)
		sort.Sort(sort.Reverse(sort.IntSlice(largestBasins)))
	}

	return largestBasins[0] * largestBasins[1] * largestBasins[2]
}

// Supporting Part A
func generateMap(rows []string) [][]int {
	newMap := make([][]int, len(rows))
	for y, row := range rows {
		newMap[y] = make([]int, len(row))
		for x := range row {
			newMap[y][x] = asInt(row[x : x+1])
		}
	}
	return newMap
}

func asInt(value string) int {
	result, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func findLowPoints(heightMap [][]int) map[Point]int {
	results := map[Point]int{}
	for y, row := range heightMap {
		for x := range row {
			height := heightMap[y][x]
			leastAdjacent := findLeastAdjacent(x, y, heightMap)
			if leastAdjacent > height {
				results[Point{x, y}] = height
			}
		}
	}
	return results
}

func findLeastAdjacent(x, y int, heightMap [][]int) int {
	left := 10
	if x != 0 {
		left = heightMap[y][x-1]
	}
	right := 10
	if x+1 < len(heightMap[0]) {
		right = heightMap[y][x+1]
	}
	up := 10
	if y != 0 {
		up = heightMap[y-1][x]
	}
	down := 10
	if y+1 < len(heightMap) {
		down = heightMap[y+1][x]
	}
	least := right
	if left < right {
		least = left
	}
	if up < least {
		least = up
	}
	if down < least {
		least = down
	}
	return least
}

func scorePartOne(lowPoints map[Point]int) int {
	score := 0
	for _, val := range lowPoints {
		score += 1 + val
	}
	return score
}

// Supporting Part B
func sizeOfBasin(start Point, heightMap [][]int) int {
	visitedSet := PointSet{}
	size := 0

	mapRange := Range{len(heightMap[0]), len(heightMap)}

	// a basic breadth-first search will solve this
	// basins are surrounded by 9s, making the end test simple
	workQueue := PointSet{start: EXISTS}
	for len(workQueue) > 0 {
		next := pop(&workQueue)
		visitedSet[next] = EXISTS

		if heightMap[next.y][next.x] < 9 {
			enqueueAdjacentPoints(next, &workQueue, visitedSet, mapRange)
			size++
		}
	}
	return size
}

func pop(pendingSet *PointSet) Point {
	var next Point
	for k := range *pendingSet {
		delete(*pendingSet, k)
		next = k
		break
	}
	return next
}

func enqueueAdjacentPoints(point Point, pending *PointSet, exclude PointSet, mapRange Range) {
	adjacents := make([]Point, 0)
	if point.x != 0 {
		adjacents = append(adjacents, Point{point.x - 1, point.y})
	}
	if point.x+1 < mapRange.x {
		adjacents = append(adjacents, Point{point.x + 1, point.y})
	}
	if point.y != 0 {
		adjacents = append(adjacents, Point{point.x, point.y - 1})
	}
	if point.y+1 < mapRange.y {
		adjacents = append(adjacents, Point{point.x, point.y + 1})
	}

	for _, next := range adjacents {
		if _, ok := exclude[next]; !ok {
			(*pending)[next] = EXISTS
		}
	}
}
