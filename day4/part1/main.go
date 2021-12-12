package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type bingoBoardValue struct {
	value  string
	marked bool
}

type bingoBoard [][]*bingoBoardValue

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	boards := []bingoBoard{}
	board := bingoBoard{}
	calledNumbers := strings.Split(scanner.Text(), ",")
	scanner.Scan()
	for scanner.Scan() {
		if scanner.Text() == "" {
			boards = append(boards, board)
			board = bingoBoard{}
			continue
		}
		fields := strings.Fields(scanner.Text())
		bingoCardValues := []*bingoBoardValue{}
		for _, field := range fields {
			bingoCardValues = append(bingoCardValues, &bingoBoardValue{value: field})
		}
		board = append(board, bingoCardValues)
	}
	boards = append(boards, board)

	winningNumber := ""
	winner := bingoBoard{}
mainLoop:
	for _, calledNumber := range calledNumbers {
		for _, board := range boards {
			modified := false
			for _, row := range board {
				for _, data := range row {
					if data.value == calledNumber {
						data.marked = true
						modified = true
					}
				}
			}

			if modified {
				// Horizontal case
				for _, row := range board {
					allMarked := true
					for _, data := range row {
						if !data.marked {
							allMarked = false
							break
						}
					}
					if allMarked {
						winner = board
						winningNumber = calledNumber
						break mainLoop
					}
				}

				// Vertical case
				for i := 0; i < len(board[0]); i++ {
					allMarked := true
					for j := 0; j < len(board); j++ {
						if !board[j][i].marked {
							allMarked = false
							break
						}
					}
					if allMarked {
						winner = board
						winningNumber = calledNumber
						break mainLoop
					}
				}
			}
		}
	}
	unmarkedTotal := 0
	for _, row := range winner {
		for _, data := range row {
			if !data.marked {
				numValue, err := strconv.Atoi(data.value)
				if err != nil {
					log.Fatal(err)
				}
				unmarkedTotal += numValue
			}
		}
	}
	winningNumberValue, err := strconv.Atoi(winningNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(unmarkedTotal * winningNumberValue)
}
