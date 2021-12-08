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
	var measurements []int
	for _, strMeasurement := range strings.Split(string(file), "\n") {
		if strMeasurement == "" {
			continue
		}
		measurement, err := strconv.Atoi(strMeasurement)
		if err != nil {
			log.Fatal(err)
		}
		measurements = append(measurements, measurement)
	}
	var lastWindow int
	var increases int
	for i, num := range measurements {
		if i < len(measurements)-2 {
			window := num + measurements[i+1] + measurements[i+2]
			if i > 0 && window > lastWindow {
				increases++
			}
			lastWindow = window
		}
	}
	fmt.Println(increases)
}
