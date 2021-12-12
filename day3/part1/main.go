package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	numbers := []int{}
	for _, cstr := range strings.Split(string(file), "\n") {
		if cstr == "" {
			break
		}

		n, err := strconv.ParseInt(cstr, 2, 0)
		if err != nil {
			log.Fatal(err)
		}

		numbers = append(numbers, int(n))
	}
	maxNumber := 0
	for _, num := range numbers {
		if num > maxNumber {
			maxNumber = num
		}
	}
	numDigits := len(strconv.FormatInt(int64(maxNumber), 2))
	numOnes := make([]int, numDigits)
	for _, num := range numbers {
		power := 1
		for i := 0; i < numDigits; i++ {
			if num&power > 0 {
				numOnes[i]++
			}
			power *= 2
		}
	}
	bits := make([]bool, numDigits)

	for i, numOne := range numOnes {
		if numOne > len(numbers)/2 {
			bits[len(bits)-(i+1)] = true
		}
	}

	gamma := 0
	epsilon := 0
	for i, bit := range bits {
		if bit {
			gamma += int(math.Pow(2, float64(len(bits)-i-1)))
		} else {
			epsilon += int(math.Pow(2, float64(len(bits)-i-1)))

		}
	}
	fmt.Println(gamma * epsilon)

}
