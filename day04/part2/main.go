package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type BingoBoardValue struct {
	value  string
	marked bool
}

type BingoBoard struct {
	numbers [][]*BingoBoardValue
	winner  bool
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	calledNumbers := strings.Split(scanner.Text(), ",")
	scanner.Scan()

	boards := []*BingoBoard{}
	board := &BingoBoard{}
	for scanner.Scan() {
		if scanner.Text() == "" {
			boards = append(boards, board)
			board = &BingoBoard{}
			continue
		}
		fields := strings.Fields(scanner.Text())
		bingoCardValues := []*BingoBoardValue{}
		for _, field := range fields {
			bingoCardValues = append(bingoCardValues, &BingoBoardValue{value: field})
		}
		board.numbers = append(board.numbers, bingoCardValues)
	}
	boards = append(boards, board)

	numWinners := 0
	var latestWinningNumber string
	var latestWinner *BingoBoard
	for _, calledNumber := range calledNumbers {
		if numWinners == len(boards) {
			break
		}
		for _, board := range boards {
			if board.winner {
				continue
			}
			modified := false
			for _, row := range board.numbers {
				for _, data := range row {
					if data.value == calledNumber {
						data.marked = true
						modified = true
					}
				}
			}

			if modified {
				// Horizontal case
				for _, row := range board.numbers {
					allMarked := true
					for _, data := range row {
						if !data.marked {
							allMarked = false
							break
						}
					}
					if allMarked {
						board.winner = true
						latestWinner = board
						latestWinningNumber = calledNumber
					}
				}

				// Vertical case
				for i := 0; i < len(board.numbers[0]); i++ {
					allMarked := true
					for j := 0; j < len(board.numbers); j++ {
						if !board.numbers[j][i].marked {
							allMarked = false
							break
						}
					}
					if allMarked {
						board.winner = true
						latestWinner = board
						latestWinningNumber = calledNumber
					}
				}
			}
		}
	}
	unmarkedTotal := 0
	for _, row := range latestWinner.numbers {
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
	winningNumberValue, err := strconv.Atoi(latestWinningNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(unmarkedTotal * winningNumberValue)
}
