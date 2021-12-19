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
			var yStart, yFinish int
			// Vertical line
			if y2 >= y1 {
				yStart, yFinish = y1, y2
			} else {
				yStart, yFinish = y2, y1
			}
			for y := yStart; y <= yFinish; y++ {
				floorVents[y][x1]++
			}
		} else if y1 == y2 {
			var xStart, xFinish int
			// Horizontal line
			if x2 >= x1 {
				xStart, xFinish = x1, x2
			} else {
				xStart, xFinish = x2, x1
			}
			for x := xStart; x <= xFinish; x++ {
				floorVents[y1][x]++
			}
		} else {
			// Diagonal line
			if x2 > x1 {
				if y2 > y1 {
					for x, y := x1, y1; x <= x2 && y <= y2; x, y = x+1, y+1 {
						floorVents[y][x]++
					}
				} else {
					for x, y := x1, y1; x <= x2 && y >= y2; x, y = x+1, y-1 {
						floorVents[y][x]++
					}
				}
			} else {
				if y2 > y1 {
					for x, y := x1, y1; x >= x2 && y <= y2; x, y = x-1, y+1 {
						floorVents[y][x]++
					}
				} else {
					for x, y := x1, y1; x >= x2 && y >= y2; x, y = x-1, y-1 {
						floorVents[y][x]++
					}
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
