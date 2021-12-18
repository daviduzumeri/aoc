package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const NUM_DAYS = 256

func main() {
	lfByDaysLeft := [9]int{}
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
		lfByDaysLeft[num]++
	}
	for i := 0; i < NUM_DAYS; i++ {
		newFish := lfByDaysLeft[0]
		lfByDaysLeft[0] = lfByDaysLeft[1]
		lfByDaysLeft[1] = lfByDaysLeft[2]
		lfByDaysLeft[2] = lfByDaysLeft[3]
		lfByDaysLeft[3] = lfByDaysLeft[4]
		lfByDaysLeft[4] = lfByDaysLeft[5]
		lfByDaysLeft[5] = lfByDaysLeft[6]
		lfByDaysLeft[6] = lfByDaysLeft[7] + newFish
		lfByDaysLeft[7] = lfByDaysLeft[8]
		lfByDaysLeft[8] = newFish
	}

	fmt.Println(lfByDaysLeft[0] +
		lfByDaysLeft[1] +
		lfByDaysLeft[2] +
		lfByDaysLeft[3] +
		lfByDaysLeft[4] +
		lfByDaysLeft[5] +
		lfByDaysLeft[6] +
		lfByDaysLeft[7] +
		lfByDaysLeft[8])
}
