package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type FrequencyTable map[int]int

func main() {
	rows := ReadFile("input.txt")
	initialPositions := asInts(strings.Split(rows[0], ","))

	position, bestFuelCost := PartOne(initialPositions)
	fmt.Println("** Result A: ", position, bestFuelCost)

	position, bestFuelCost = PartTwo(initialPositions)
	fmt.Println("** Result B: ", position, bestFuelCost)
}

func PartOne(initialPositions []int) (int, int) {
	return bestFuelCost(initialPositions, func(left int, right int) int {
		if left < right {
			return right - left
		}
		return left - right
	})
}

func PartTwo(initialPositions []int) (int, int) {
	return bestFuelCost(initialPositions, func(left int, right int) int {
		if left < right {
			return sumOfFirstNInts(right - left)
		}
		return sumOfFirstNInts(left - right)
	})
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

// Supporting functions
func bestFuelCost(initialPositions []int, pointCostFunction func(int, int) int) (position int, cost int) {
	bestTarget := 0
	bestCost := math.MaxInt
	for i := 0; i < len(initialPositions); i++ {
		currentCost := 0
		for _, val := range initialPositions {
			currentCost += pointCostFunction(val, i)
		}

		if currentCost < bestCost {
			bestCost = currentCost
			bestTarget = i
		}
	}
	return bestTarget, bestCost
}

// Util functions
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

func sumOfFirstNInts(n int) int {
	return n * (n + 1) / 2
}
