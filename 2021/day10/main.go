package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

const P_OPEN = '('
const P_CLOSE = ')'
const B_OPEN = '['
const B_CLOSE = ']'
const C_OPEN = '{'
const C_CLOSE = '}'
const A_OPEN = '<'
const A_CLOSE = '>'

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
	scoreMap := map[rune]int{P_CLOSE: 3, B_CLOSE: 57, C_CLOSE: 1197, A_CLOSE: 25137}

	score := 0
	for _, row := range rows {
		if corrupt, violation := isCorrupted(row); corrupt {
			score += scoreMap[violation]
		}
	}
	return score
}

func PartTwo(rows []string) int {
	scores := []int{}
	for _, row := range rows {
		if ok, _ := isCorrupted(row); ok {
			continue
		}
		scores = append(scores, fixIncomplete(row))
	}
	sort.Ints(scores)

	middle := len(scores) / 2
	return scores[middle]
}

// has unmatched closing character
func isCorrupted(line string) (bool, rune) {
	stack := ""
	var found rune
	for _, symbol := range line {
		switch symbol {
		case P_OPEN:
			stack = push(symbol, stack)
		case B_OPEN:
			stack = push(symbol, stack)
		case C_OPEN:
			stack = push(symbol, stack)
		case A_OPEN:
			stack = push(symbol, stack)
		case P_CLOSE:
			found, stack = pop(stack)
			if P_OPEN != found {
				return true, P_CLOSE
			}
		case B_CLOSE:
			found, stack = pop(stack)
			if B_OPEN != found {
				return true, B_CLOSE
			}
		case C_CLOSE:
			found, stack = pop(stack)
			if C_OPEN != found {
				return true, C_CLOSE
			}
		case A_CLOSE:
			found, stack = pop(stack)
			if A_OPEN != found {
				return true, A_CLOSE
			}
		}
	}
	return false, 0
}

func fixIncomplete(row string) int {
	scoreMap := map[rune]int{P_CLOSE: 1, B_CLOSE: 2, C_CLOSE: 3, A_CLOSE: 4}
	score := 0

	stack := ""
	for _, symbol := range row {
		switch symbol {
		case P_OPEN:
			stack = push(symbol, stack)
		case B_OPEN:
			stack = push(symbol, stack)
		case C_OPEN:
			stack = push(symbol, stack)
		case A_OPEN:
			stack = push(symbol, stack)
		case P_CLOSE:
			_, stack = pop(stack)
		case B_CLOSE:
			_, stack = pop(stack)
		case C_CLOSE:
			_, stack = pop(stack)
		case A_CLOSE:
			_, stack = pop(stack)
		}
	}

	for _, unmatched := range stack {
		switch unmatched {
		case P_OPEN:
			score = score*5 + scoreMap[P_CLOSE]
		case B_OPEN:
			score = score*5 + scoreMap[B_CLOSE]
		case C_OPEN:
			score = score*5 + scoreMap[C_CLOSE]
		case A_OPEN:
			score = score*5 + scoreMap[A_CLOSE]
		}
	}
	return score
}

func push(symbol rune, stack string) string {
	return string(symbol) + stack
}

func pop(stack string) (rune, string) {
	return []rune(stack)[0], stack[1:]
}
