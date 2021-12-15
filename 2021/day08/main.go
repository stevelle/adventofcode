package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type runeSort []rune

func (s runeSort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s runeSort) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s runeSort) Len() int {
	return len(s)
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
	sum := 0
	for _, row := range rows {
		idx := 2 + strings.Index(row, "|")
		output := row[idx:]
		digits := strings.Split(output, " ")
		sum += countKnown(digits)
	}
	return sum
}

func PartTwo(rows []string) int {
	sum := 0
	for _, row := range rows {
		idx := strings.Index(row, "|")
		patterns := strings.Split(row[:idx-1], " ")
		signals := decipher(patterns)

		output := row[idx+2:]
		digits := strings.Split(output, " ")
		sum += decode(digits, signals)
	}
	return sum
}

// count the easy-to-recognize digits based on number of segments
func countKnown(digits []string) int {
	sum := 0
	for _, digit := range digits {
		length := len(digit)
		switch length {
		case 2:
			sum++
		case 3:
			sum++
		case 4:
			sum++
		case 7:
			sum++
		}
	}
	return sum
}

// decipher digits by deduction, given patterns for each of 10 digits
func decipher(patterns []string) map[string]int {
	known := make([]string, 10)
	signals := make(map[string]int)
	for len(patterns) > 0 {
		digit := SortRunes(patterns[0])
		patterns = patterns[1:]

		length := len(digit)
		switch length {
		case 2:
			signals[digit] = 1
			known[1] = digit
		case 3:
			signals[digit] = 7
			known[7] = digit
		case 4:
			signals[digit] = 4
			known[4] = digit
		case 5:
			encodedDigit, requeue := decipher_fives(known, digit)
			// requeue this digit
			if requeue {
				patterns = append(patterns, digit)
			} else {
				signals[digit] = encodedDigit
				known[encodedDigit] = digit
			}
		case 6:
			encodedDigit, requeue := decipher_sixes(known, digit)
			// requeue this digit
			if requeue {
				patterns = append(patterns, digit)
			} else {
				signals[digit] = encodedDigit
				known[encodedDigit] = digit
			}
		case 7:
			signals[digit] = 8
			known[8] = digit
		}

	}
	return signals
}

func SortRunes(text string) string {
	runes := []rune(text)
	sort.Sort(runeSort(runes))
	return string(runes)
}

// decipher digits with 5 segments by deduction, given patterns for known digits
func decipher_fives(known []string, digit string) (int, bool) {
	if known[1] != "" && ContainsAll(digit, known[1]) {
		return 3, false
	} else if known[9] != "" {
		if ContainsAll(known[9], digit) {
			return 5, false
		} else {
			return 2, false
		}
	} else if known[6] != "" {
		if ContainsAll(known[6], digit) {
			return 5, false
		} else {
			return 2, false
		}
	}
	return 0, true
}

// decipher digits with 6 segments by deduction, given patterns for known digits
func decipher_sixes(known []string, digit string) (int, bool) {
	// retstructure me?
	if known[1] == "" {
		return 0, true
	} else if ContainsAll(digit, known[1]) {
		if known[4] == "" {
			if known[5] == "" {
				return 0, true
			} else if ContainsAll(digit, known[5]) {
				return 9, false
			} else {
				return 0, false
			}
		} else if ContainsAll(digit, known[4]) {
			return 9, false
		} else {
			return 0, false
		}
	} else {
		return 6, false
	}
}

// does the text contain each of the given characters?
func ContainsAll(text, chars string) bool {
	for _, r := range chars {
		if !strings.ContainsRune(text, r) {
			return false
		}
	}
	return true
}

// decode given digits, based on the dictionary of known signals
func decode(digits []string, signals map[string]int) int {
	output := 0
	for _, digit := range digits {
		output *= 10
		output += signals[SortRunes(digit)]
	}
	return output
}
