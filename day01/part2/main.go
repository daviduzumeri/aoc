package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	measurements := []int{}
	strMeasurements := strings.Split(string(file), "\n")
	for _, mstr := range strMeasurements {
		if mstr == "" {
			break
		}
		m, err := strconv.Atoi(mstr)
		if err != nil {
			log.Fatal(err)
		}
		measurements = append(measurements, m)
	}
	lastWindow := 0
	increases := 0
	for i := 0; i < len(measurements)-2; i++ {
		window := measurements[i] + measurements[i+1] + measurements[i+2]

		// Skip the first window
		if i > 0 && window > lastWindow {
			increases++
		}
		lastWindow = window
	}
	fmt.Println(increases)
}
