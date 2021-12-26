package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type cavern [][]int

func (c cavern) lowestTotalRisk(x int, y int, risk int) int {
	paths := []int{}
	if x == 0 && y == 0 {
		c[0][0] = 0
	}
	if x == len(c[0])-1 && y == len(c)-1 {
		return c[y][x] + risk
	}
	if x < len(c[0])-1 {
		paths = append(paths, c.lowestTotalRisk(x+1, y, c[y][x]+risk))
	}
	if y < len(c)-1 {
		paths = append(paths, c.lowestTotalRisk(x, y+1, c[y][x]+risk))
	}
	minRisk := math.MaxInt
	for _, p := range paths {
		if p < minRisk {
			minRisk = p
		}
	}
	return minRisk
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	c := cavern{}
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), "")
		row := []int{}
		for _, risk := range input {
			riskInt, err := strconv.Atoi(risk)
			if err != nil {
				log.Fatal(err)
			}
			row = append(row, riskInt)
		}
		c = append(c, row)
	}
	fmt.Println(c.lowestTotalRisk(0, 0, 0))
}
