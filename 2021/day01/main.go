package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	inputFile := "input.txt"
	PartOne(inputFile)
	PartTwo(inputFile)
}

func PartOne(inputFilename string) {
	lines := ReadFile(inputFilename)

	// preload the first line
	first, lines := PopStr(lines)
	prev := AsInt(first)

	counter := 0
	for _, line := range lines {
		next := AsInt(line)

		if prev < next {
			counter++
		}
		prev = next
	}

	fmt.Println("** Increases: ", counter)
}

func PartTwo(inputFilename string) {
	lines := ReadFile(inputFilename)

	// preload the window
	window := make([]int, 3)
	for i := 0; i < 3; i++ {
		var line string
		line, lines = PopStr(lines)
		window[i] = AsInt(line)
	}

	counter := 0
	for _, line := range lines {
		prevSum := Sum(window...)
		_, window = SlideInt(AsInt(line), window)

		newSum := Sum(window...)
		if prevSum < newSum {
			counter++
		}
	}

	fmt.Println("** Increases: ", counter)
}

func Sum(values ...int) int {
	sum := 0
	for _, n := range values {
		sum += n
	}
	return sum
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

func PopStr(from []string) (string, []string) {
	return from[0], from[1:]
}

func SlideInt(value int, onto []int) (int, []int) {
	window := append(onto, value)
	return window[0], window[1:]
}

func AsInt(value string) int {
	result, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
