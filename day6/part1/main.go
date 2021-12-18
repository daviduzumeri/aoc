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

const NUM_DAYS = 80

func main() {
	lanternfish := []int{}
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
		lanternfish = append(lanternfish, num)
	}
	for i := 0; i < NUM_DAYS; i++ {
		newFish := []int{}
		for j := 0; j < len(lanternfish); j++ {
			if lanternfish[j] == 0 {
				lanternfish[j] = 6
				newFish = append(newFish, 8)
			} else {
				lanternfish[j]--
			}
		}
		lanternfish = append(lanternfish, newFish...)
	}

	fmt.Println(len(lanternfish))
}
