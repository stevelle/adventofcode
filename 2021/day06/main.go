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
	yesterdayState := calculateInitialState(rows[0])
	fmt.Println("** Result A: ", PartOne(yesterdayState))
	fmt.Println("** Result B: ", PartTwo(yesterdayState))
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

func calculateInitialState(input string) []int {
	initialPopulation := asInts(strings.Split(input, ","))
	todayState := make([]int, 9)
	for _, days := range initialPopulation {
		todayState[days] += 1
	}
	return todayState
}

func PartOne(todayState []int) int64 {
	return modelPopGrowth(80, todayState)
}

func PartTwo(todayState []int) int64 {
	return modelPopGrowth(256, todayState)
}

func modelPopGrowth(daysToModel int, todayState []int) int64 {
	for dayNum := 0; dayNum < daysToModel; dayNum++ {
		tomorrowState := make([]int, 9)

		// pad == Population At Day (x) in it's reproduction cycle
		for pad := 0; pad < 8; pad++ {
			tomorrowState[pad] += todayState[pad+1]
		}
		tomorrowState[6] += todayState[0]
		tomorrowState[8] += todayState[0]

		todayState = tomorrowState
	}
	return sum(todayState...)
}

// Utility functions
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

func sum(values ...int) int64 {
	sum := int64(0)
	for _, n := range values {
		sum += int64(n)
	}
	return sum
}
