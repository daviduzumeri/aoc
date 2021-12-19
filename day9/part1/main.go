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
	heightmap := [][]int{}
	for scanner.Scan() {
		heightmapRow := []int{}
		row := scanner.Text()
		for _, r := range row {
			heightmapRow = append(heightmapRow, int(r)-'0')
		}
		heightmap = append(heightmap, heightmapRow)
	}

	risk := 0
	for y, row := range heightmap {
		for x, location := range row {
			if (x >= len(row)-1 || location < heightmap[y][x+1]) &&
				(y >= len(heightmap)-1 || location < heightmap[y+1][x]) &&
				(x <= 0 || location < heightmap[y][x-1]) &&
				(y <= 0 || location < heightmap[y-1][x]) {
				risk += location + 1
			}
		}
	}
	fmt.Println(risk)
}
