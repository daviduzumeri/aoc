package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	errorScore := 0
	for scanner.Scan() {
		row := scanner.Text()
		chunkStarts := []rune{}
	rowLoop:
		for _, r := range row {
			switch r {
			case ')':
				if len(chunkStarts) > 0 && chunkStarts[len(chunkStarts)-1] == '(' {
					chunkStarts = chunkStarts[:len(chunkStarts)-1]
				} else {
					errorScore += 3
					break rowLoop
				}
			case ']':
				if len(chunkStarts) > 0 && chunkStarts[len(chunkStarts)-1] == '[' {
					chunkStarts = chunkStarts[:len(chunkStarts)-1]
				} else {
					errorScore += 57
					break rowLoop
				}
			case '}':
				if len(chunkStarts) > 0 && chunkStarts[len(chunkStarts)-1] == '{' {
					chunkStarts = chunkStarts[:len(chunkStarts)-1]
				} else {
					errorScore += 1197
					break rowLoop
				}
			case '>':
				if len(chunkStarts) > 0 && chunkStarts[len(chunkStarts)-1] == '<' {
					chunkStarts = chunkStarts[:len(chunkStarts)-1]
				} else {
					errorScore += 25137
					break rowLoop
				}
			default:
				chunkStarts = append(chunkStarts, r)
			}
		}
	}
	fmt.Println(errorScore)
}
