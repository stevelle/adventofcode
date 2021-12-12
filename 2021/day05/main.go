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
	rows := ReadFile("input.txt")
	fmt.Println("** Result A: ", PartOne(rows))
	fmt.Println("** Result B: ", PartTwo(rows))
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

func PartOne(rawLines []string) int {
	grid := Grid{}

	for _, rawLine := range rawLines {
		line := parseLine(rawLine)
		// only keep vertical or horizontal rows
		if line.isHorizontal() || line.isVertical() {
			grid.draw(line)
		}
	}

	return grid.scoreGrid()
}

func PartTwo(rawLines []string) int {
	grid := Grid{}

	for _, rawLine := range rawLines {
		line := parseLine(rawLine)
		grid.draw(line)
	}

	return grid.scoreGrid()
}

func parseLine(row string) Line {
	pairs := strings.Split(row, " -> ")
	line := make([]int, 0)
	for _, each := range pairs {
		line = append(line, asInts(strings.Split(each, ","))...)
	}
	return Line{Point{line[0], line[1]}, Point{line[2], line[3]}}
}

func asInts(inputs []string) []int {
	results := make([]int, len(inputs))
	for i, text := range inputs {
		results[i] = asInt(text)
	}
	return results
}

func asInt(value string) int {
	result, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
