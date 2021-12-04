package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	lines := ReadFile("input.txt")
	fmt.Println("** Result A: ", PartOne(lines))
	fmt.Println("** Result B: ", PartTwo(lines))
}

func PartOne(lines []string) int64 {
	gammaString := ""
	for i := 0; i < len(lines[0]); i++ {
		gammaString += string(valueToKeep(i, lines, true))
	}
	return ParseBinary(gammaString) * ParseBinary(Not(gammaString))
}

func PartTwo(oxyRows []string) int64 {
	// copy the input rows before we start filtering
	co2Rows := make([]string, len(oxyRows))
	copy(co2Rows, oxyRows)

	// filter by applying bit criteria
	oxygen := filter(oxyRows, true)
	co2 := filter(co2Rows, false)

	return ParseBinary(oxygen) * ParseBinary(co2)
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

func ParseBinary(value string) int64 {
	result, err := strconv.ParseInt(value, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func Not(value string) string {
	result := ""
	for i := 0; i < len(value); i++ {
		if value[i] == '0' {
			result += "1"
		} else {
			result += "0"
		}
	}
	return result
}

func filter(rows []string, keepMostCommon bool) string {
	for colNum := 0; colNum < len(rows[0]); colNum++ {
		valueToKeep := valueToKeep(colNum, rows, keepMostCommon)

		// decrementing is slightly easier to read than incrementing
		for rowNum := len(rows) - 1; rowNum >= 0; rowNum-- {
			if rows[rowNum][colNum] != valueToKeep {
				rows = remove(rows, rowNum)
			}
		}
		if len(rows) == 1 {
			break
		}
	}
	return rows[0]
}

func valueToKeep(colNum int, rows []string, keepMostCommon bool) byte {
	threshold := (len(rows) + 1) / 2
	columnSum := calculateColumnSum(colNum, rows)
	if keepMostCommon {
		// if the most common value in this column is 0
		if columnSum < threshold {
			return '0'
		}
		return '1'
	} else {
		// if the most common value in this column is 0
		if columnSum < threshold {
			return '1'
		}
		return '0'
	}
}

func calculateColumnSum(colNum int, rows []string) int {
	sum := 0
	for _, row := range rows {
		if row[colNum] == '1' {
			sum += 1
		}
	}
	return sum
}

// don't need to preserve ordering of rows
func remove(slice []string, index int) []string {
	slice[index] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
