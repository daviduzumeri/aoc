package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type location struct {
	height int
	marked bool
}

type heightmap [][]*location

func (h heightmap) findBasin(x int, y int) int {
	loc := h[y][x]

	if loc.height == 9 || loc.marked {
		return 0
	}

	loc.marked = true

	size := 1
	if x < len(h[y])-1 && (h[y][x+1].height >= loc.height) {
		size += h.findBasin(x+1, y)
	}
	if y < len(h)-1 && (h[y+1][x].height >= loc.height) {
		size += h.findBasin(x, y+1)
	}
	if x > 0 && (h[y][x-1].height >= loc.height) {
		size += h.findBasin(x-1, y)
	}
	if y > 0 && (h[y-1][x].height >= loc.height) {
		size += h.findBasin(x, y-1)
	}
	return size
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	heightmap := heightmap{}
	for scanner.Scan() {
		heightmapRow := []*location{}
		row := scanner.Text()
		for _, r := range row {
			heightmapRow = append(heightmapRow, &location{height: int(r) - '0'})
		}
		heightmap = append(heightmap, heightmapRow)
	}

	basinSizes := []int{}
	for y, row := range heightmap {
		for x, loc := range row {
			if (x >= len(row)-1 || loc.height < heightmap[y][x+1].height) &&
				(y >= len(heightmap)-1 || loc.height < heightmap[y+1][x].height) &&
				(x <= 0 || loc.height < heightmap[y][x-1].height) &&
				(y <= 0 || loc.height < heightmap[y-1][x].height) {
				size := heightmap.findBasin(x, y)
				basinSizes = append(basinSizes, size)
			}
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))
	fmt.Println(basinSizes[0] * basinSizes[1] * basinSizes[2])
}
