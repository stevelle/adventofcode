package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type Point struct {
	x int
	y int
}

func main() {
	rows := ReadFile("input.txt")

	score := PartOne(rows)
	fmt.Println("** Result A: ", score)

	score = PartTwo(rows)
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

func PartOne(rows []string) int {
	board := constructBoard(rows)
	flashCount := 0
	for step := 1; step <= 100; step++ {
		flashes := 0
		upBoard, queue := powerUp(board)
		board, flashes = triggerFlashing(upBoard, queue)
		board = updateEnergy(board)
		flashCount += flashes
	}
	return flashCount
}

func PartTwo(rows []string) int {
	board := constructBoard(rows)
	for step := 1; true; step++ {
		upBoard, readyToFlash := powerUp(board)
		board, _ = triggerFlashing(upBoard, readyToFlash)
		board = updateEnergy(board)
		if allFlashed(board) {
			return step
		}
	}
	return 0
}

func constructBoard(rows []string) [][]int {
	board := make([][]int, 0)
	for _, line := range rows {
		currentRow := []int{}
		for _, r := range line {
			currentRow = append(currentRow, asInt(r))
		}
		board = append(board, currentRow)
	}
	return board
}

func asInt(value rune) int {
	result, err := strconv.Atoi(string(value))
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func powerUp(board [][]int) ([][]int, []Point) {
	readyToFlash := make([]Point, 0)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			board[i][j]++
			if board[i][j] > 9 {
				readyToFlash = append(readyToFlash, Point{j, i})
			}
		}
	}
	return board, readyToFlash
}

func triggerFlashing(board [][]int, workQueue []Point) ([][]int, int) {
	flashed := make([]Point, 0, 99)
	for len(workQueue) > 0 {
		next := workQueue[0]
		if board[next.y][next.x] > 9 {
			// exclude points which flashed from the rest of the search
			flashed = append(flashed, next)
			neighbors := getNeighbors(next.y, next.x, flashed)
			workQueue = append(workQueue, neighbors...)
			// disperse energy onto neighbors
			for _, n := range neighbors {
				board[n.y][n.x]++
			}
			// a negative score at this point indicates a flash this step
			board[next.y][next.x] = math.MinInt
		}
		workQueue = workQueue[1:]
	}

	// count how many flashes were triggered
	flashes := 0
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if board[i][j] < 0 {
				flashes++
			}
		}
	}
	return board, flashes
}

func getNeighbors(y, x int, excluded []Point) []Point {
	minY := y
	minX := x
	maxY := y
	maxX := x
	if y > 0 {
		minY--
	}
	if x > 0 {
		minX--
	}
	if y < 9 {
		maxY++
	}
	if x < 9 {
		maxX++
	}
	neighbors := []Point{}
	for i := minY; i <= maxY; i++ {
		for j := minX; j <= maxX; j++ {
			point := Point{j, i}
			if (i != y || j != x) && point.isNotIn(excluded) {
				neighbors = append(neighbors, point)
			}
		}
	}

	return neighbors
}

func (p *Point) isNotIn(excluded []Point) bool {
	for _, v := range excluded {
		if *p == v {
			return false
		}
	}
	return true
}

func updateEnergy(board [][]int) [][]int {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			// negative score indicates a flash at this point
			if board[i][j] < 0 {
				board[i][j] = 0
			}
		}
	}
	return board
}

func allFlashed(board [][]int) bool {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if board[i][j] != 0 {
				return false
			}
		}
	}
	return true
}
