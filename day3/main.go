package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// NOTE: Bitwise operators was not, in retrospect, the sane way to approach this problem, but I stuck with it out of stubbornness.

func findNumber(numbers []int, moreCommon bool, numDigits int) int {
	if len(numbers) == 1 {
		return numbers[0]
	}

	// Build match array for this sub-array
	maxNumber := 0
	for _, num := range numbers {
		if num > maxNumber {
			maxNumber = num
		}
	}
	if numDigits == -1 {
		numDigits = len(strconv.FormatInt(int64(maxNumber), 2))
	}
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
	newArray := []int{}
	bits := make([]bool, numDigits)

	for i, numOne := range numOnes {
		if float64(numOne) >= float64(len(numbers))/2 {
			bits[len(bits)-(i+1)] = true
		}
	}

	returnArray := make([]string, numDigits)
	for i, bit := range bits {
		if bit {
			if moreCommon {
				returnArray[i] = "1"
			} else {
				returnArray[i] = "0"
			}
		} else {
			if moreCommon {
				returnArray[i] = "0"
			} else {
				returnArray[i] = "1"
			}
		}
	}
	returnNumber, err := strconv.ParseInt(strings.Join(returnArray, ""), 2, 0)
	if err != nil {
		log.Fatal(err)
	}
	maxBit := int(math.Pow(2, float64(numDigits-1)))

	for _, number := range numbers {
		if int(returnNumber)&maxBit > 0 && number&maxBit > 0 {
			newArray = append(newArray, number)
		}
		if int(returnNumber)&maxBit == 0 && number&maxBit == 0 {
			newArray = append(newArray, number)
		}
	}
	return findNumber(newArray, moreCommon, numDigits-1)
}

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
	fmt.Println(findNumber(numbers, true, -1) * findNumber(numbers, false, -1))
}
