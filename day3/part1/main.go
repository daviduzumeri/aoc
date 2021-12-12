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

	gammaArray := make([]string, numDigits)
	epsilonArray := make([]string, numDigits)
	for i, bit := range bits {
		if bit {
			gammaArray[i] = "1"
			epsilonArray[i] = "0"
		} else {
			gammaArray[i] = "0"
			epsilonArray[i] = "1"

		}
	}
	gamma, err := strconv.ParseInt(strings.Join(gammaArray, ""), 2, 0)
	if err != nil {
		log.Fatal(err)
	}
	epsilon, err := strconv.ParseInt(strings.Join(epsilonArray, ""), 2, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(gamma * epsilon)

}
