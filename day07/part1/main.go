package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	crabs := []int{}
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	inputs := strings.FieldsFunc(scanner.Text(), func(r rune) bool {
		return !unicode.IsDigit(r)
	})
	for _, i := range inputs {
		num, err := strconv.Atoi(i)
		if err != nil {
			log.Fatal(err)
		}
		crabs = append(crabs, num)
	}

	minCrab := math.MaxInt
	maxCrab := 0
	for _, crab := range crabs {
		if minCrab > crab {
			minCrab = crab
		}
		if maxCrab < crab {
			maxCrab = crab
		}
	}

	minFuel := math.MaxInt
	for i := minCrab; i <= maxCrab; i++ {
		fuel := 0
		for _, crab := range crabs {
			fuel += int(math.Abs(float64(crab - i)))
		}
		if fuel < minFuel {
			minFuel = fuel
		}
	}

	fmt.Println(minFuel)
}
