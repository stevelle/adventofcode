package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputFile := "input.txt"
	PartOne(inputFile)
	PartTwo(inputFile)
}

func PartOne(inputFilename string) {
	lines := ReadFile(inputFilename)

	var depth int64
	var position int64

	for _, line := range lines {
		x, y := AsCoordinates(line)
		position += int64(x)
		depth += int64(y)
	}
	fmt.Println("** Result A: ", depth*position)
}

func PartTwo(inputFilename string) {
	lines := ReadFile(inputFilename)

	var depth int
	var position int
	var aim int

	for _, line := range lines {
		x, a := AsCoordinates(line)
		position += x
		depth += x * aim
		aim += a
	}
	fmt.Println("** Result B: ", depth*position)
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

func AsInt(value string) int {
	result, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func AsCoordinates(text string) (x int, y int) {
	tokens := strings.Split(text, " ")
	direction := tokens[0]
	distance := AsInt(tokens[1])

	switch direction {
	case "up":
		return 0, -distance
	case "down":
		return 0, distance
	case "forward":
		return distance, 0
	default:
		log.Fatal("unknown direction: ", direction)
	}
	return 0, 0
}
