package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type octopusGrid [][]*octopus

type octopus struct {
	energyLevel int
	flashed     bool
}

func (og octopusGrid) increment(x int, y int) int {
	if x < 0 || y < 0 {
		return 0
	}
	if x > 9 || y > 9 {
		return 0
	}
	og[y][x].energyLevel++
	return og.checkFlashes(x, y)
}

func (og octopusGrid) checkFlashes(x int, y int) int {
	if x < 0 || y < 0 {
		return 0
	}
	if x > 9 || y > 9 {
		return 0
	}
	o := og[y][x]
	if o.energyLevel > 9 && !o.flashed {
		numFlashes := 1
		o.flashed = true
		numFlashes += og.increment(x+1, y+1)
		numFlashes += og.increment(x+1, y)
		numFlashes += og.increment(x+1, y-1)
		numFlashes += og.increment(x, y+1)
		numFlashes += og.increment(x, y-1)
		numFlashes += og.increment(x-1, y+1)
		numFlashes += og.increment(x-1, y)
		numFlashes += og.increment(x-1, y-1)
		return numFlashes
	}
	return 0
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	og := octopusGrid{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		or := []*octopus{}
		row := scanner.Text()
		for _, r := range row {
			or = append(or, &octopus{energyLevel: (int(r) - '0')})
		}
		og = append(og, or)
	}

	numFlashes := 0
	for step := 1; ; step++ {
		for _, or := range og {
			for _, o := range or {
				o.energyLevel++
				o.flashed = false
			}
		}
		for y, or := range og {
			for x := range or {
				numFlashes += og.checkFlashes(x, y)
			}
		}
		allFlashed := true
		for _, or := range og {
			for _, o := range or {
				if o.flashed {
					o.energyLevel = 0
				} else {
					allFlashed = false
				}
			}
		}
		if allFlashed {
			fmt.Println(step)
			break
		}
	}
}
