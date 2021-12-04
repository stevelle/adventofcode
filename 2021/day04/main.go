package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const MARKED = "x"
const BINGO = 5

func main() {
	lines := ReadFile("input.txt")
	fmt.Println("** Result A: ", PartOne(lines))
	fmt.Println("** Result B: ", PartTwo(lines))
}

func PartOne(lines []string) int64 {
	queue := strings.Split(lines[0], ",")
	boards := constructBoards(lines[1:])
	boardNum, lastDraw := findWinningBoard(boards, queue)
	return score(boards[boardNum], lastDraw)
}

func PartTwo(lines []string) int64 {
	queue := strings.Split(lines[0], ",")
	boards := constructBoards(lines[1:])

	boardNum := -1
	lastDraw := ""
	var lastBoard [][]string
	for len(boards) > 0 {
		boardNum, lastDraw = findWinningBoard(boards, queue)
		lastBoard = boards[boardNum]
		boards = remove(boards, boardNum)
	}
	return score(lastBoard, lastDraw)
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

func constructBoards(rows []string) [][][]string {
	boards := make([][][]string, 0)
	var currentBoard [][]string
	for _, line := range rows {
		// blank lines indicate time for a new board
		if len(strings.TrimSpace(line)) == 0 {
			if len(currentBoard) == 5 {
				boards = append(boards, currentBoard)
			}
			currentBoard = make([][]string, 0)
			continue
		}
		currentBoard = append(currentBoard, strings.Fields(line))
	}
	// be sure to add the last board built
	boards = append(boards, currentBoard)
	return boards
}

func findWinningBoard(boards [][][]string, queue []string) (int, string) {
	for step, draw := range queue {
		for b, board := range boards {
			for r, row := range board {
				for c, cell := range row {
					if draw == cell {
						boards[b][r][c] = MARKED
						if step >= BINGO {
							if isWinner(board) {
								return b, draw
							}
						}
					}
				}
			}
		}
	}
	return -1, ""
}

func isWinner(board [][]string) bool {
	rowScore := []int{0, 0, 0, 0, 0}
	colsScore := []int{0, 0, 0, 0, 0}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if board[i][j] == MARKED {
				colsScore[j] += 1
				rowScore[i] += 1
			}
			if colsScore[j] == BINGO {
				return true
			}
		}
		if rowScore[i] == BINGO {
			return true
		}
	}
	return false
}

func score(board [][]string, draw string) int64 {
	sum := int64(0)
	for _, row := range board {
		for _, cell := range row {
			if cell != MARKED {
				sum += asInt64(cell)
			}
		}
	}
	return sum * asInt64(draw)
}

func asInt64(value string) int64 {
	result, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal(err)
	}
	return int64(result)
}

func remove(slice [][][]string, index int) [][][]string {
	slice[index] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
