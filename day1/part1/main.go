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
	lastMeasurement := 0
	increases := 0
	for i, measurement := range measurements {
		// Skip the first window
		if i > 0 && measurement > lastMeasurement {
			increases++
		}
		lastMeasurement = measurement
	}
	fmt.Println(increases)
}
