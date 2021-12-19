package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	var floorVents [1000][1000]int
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		coordinates := strings.FieldsFunc(scanner.Text(), func(r rune) bool {
			return !unicode.IsDigit(r)
		})
		x1, err := strconv.Atoi(coordinates[0])
		if err != nil {
			log.Fatal(err)
		}
		y1, err := strconv.Atoi(coordinates[1])
		if err != nil {
			log.Fatal(err)
		}
		x2, err := strconv.Atoi(coordinates[2])
		if err != nil {
			log.Fatal(err)
		}
		y2, err := strconv.Atoi(coordinates[3])
		if err != nil {
			log.Fatal(err)
		}
		if x1 == x2 {
			// Horizontal line
			if y2 >= y1 {
				for i := y1; i <= y2; i++ {
					floorVents[i][x1]++
				}
			} else {
				for i := y1; i >= y2; i-- {
					floorVents[i][x1]++
				}
			}
		} else if y1 == y2 {
			// Vertical line
			if x2 >= x1 {
				for i := x1; i <= x2; i++ {
					floorVents[y1][i]++
				}
			} else {
				for i := x1; i >= x2; i-- {
					floorVents[y1][i]++
				}
			}
		}
	}

	numOverlaps := 0
	for _, row := range floorVents {
		for _, point := range row {
			if point >= 2 {
				numOverlaps++
			}
		}
	}
	fmt.Println(numOverlaps)
}
