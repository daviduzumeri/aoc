package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type coordinate struct {
	x int
	y int
}

type fold struct {
	axis  string
	value int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	coordinateSection := true
	coordinates := []coordinate{}
	folds := []fold{}
	maxXValue := 0
	maxYValue := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			coordinateSection = false
		} else if coordinateSection {
			input := strings.Split(scanner.Text(), ",")
			x, err := strconv.Atoi(input[0])
			if err != nil {
				log.Fatal(err)
			}
			y, err := strconv.Atoi(input[1])
			if err != nil {
				log.Fatal(err)
			}
			if x > maxXValue {
				maxXValue = x
			}
			if y > maxYValue {
				maxYValue = y
			}
			coordinates = append(coordinates, coordinate{x: x, y: y})
		} else {
			input := strings.Split(strings.Split(scanner.Text(), " ")[2], "=")
			value, err := strconv.Atoi(input[1])
			if err != nil {
				log.Fatal(err)
			}
			folds = append(folds, fold{axis: input[0], value: value})
		}
	}

	dots := make([][]bool, maxYValue+1)
	for i := range dots {
		dots[i] = make([]bool, maxXValue+1)
	}

	for _, c := range coordinates {
		dots[c.y][c.x] = true
	}

	// Only do first fold (for now)
	for _, fold := range folds {
		switch fold.axis {
		case "x":
			for x := fold.value + 1; x < maxXValue+1; x++ {
				for y := 0; y < maxYValue+1; y++ {
					if dots[y][x] {
						dots[y][fold.value-(x-fold.value)] = true
					}
				}
			}
			for i, r := range dots {
				dots[i] = r[0:fold.value]
			}
			maxXValue = fold.value - 1
		case "y":
			for y := fold.value + 1; y < maxYValue+1; y++ {
				for x := 0; x < maxXValue+1; x++ {
					if dots[y][x] {
						dots[fold.value-(y-fold.value)][x] = true
					}
				}
			}
			dots = dots[0:fold.value]
			maxYValue = fold.value - 1
		default:
			log.Fatalf("Nonexistent axis %v", fold.axis)
		}
	}
	for _, r := range dots {
		for _, d := range r {
			if d {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
