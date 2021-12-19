package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	autocorrectScores := []int{}
scanLoop:
	for scanner.Scan() {
		row := scanner.Text()
		chunkStarts := []rune{}
		for _, r := range row {
			switch r {
			case ')':
				if len(chunkStarts) > 0 && chunkStarts[len(chunkStarts)-1] == '(' {
					chunkStarts = chunkStarts[:len(chunkStarts)-1]
				} else {
					continue scanLoop
				}
			case ']':
				if len(chunkStarts) > 0 && chunkStarts[len(chunkStarts)-1] == '[' {
					chunkStarts = chunkStarts[:len(chunkStarts)-1]
				} else {
					continue scanLoop
				}
			case '}':
				if len(chunkStarts) > 0 && chunkStarts[len(chunkStarts)-1] == '{' {
					chunkStarts = chunkStarts[:len(chunkStarts)-1]
				} else {
					continue scanLoop
				}
			case '>':
				if len(chunkStarts) > 0 && chunkStarts[len(chunkStarts)-1] == '<' {
					chunkStarts = chunkStarts[:len(chunkStarts)-1]
				} else {
					continue scanLoop
				}
			default:
				chunkStarts = append(chunkStarts, r)
			}
		}

		autocorrectScore := 0

		if len(chunkStarts) > 0 {
			for i := len(chunkStarts) - 1; i >= 0; i-- {
				cs := chunkStarts[i]
				autocorrectScore *= 5
				switch cs {
				case '(':
					autocorrectScore += 1
				case '[':
					autocorrectScore += 2
				case '{':
					autocorrectScore += 3
				case '<':
					autocorrectScore += 4
				}
			}
		}
		autocorrectScores = append(autocorrectScores, autocorrectScore)
	}
	sort.Ints(autocorrectScores)
	fmt.Println(autocorrectScores[len(autocorrectScores)/2])
}
