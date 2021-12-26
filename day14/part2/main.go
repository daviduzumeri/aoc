package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	template := ""
	rules := map[string]*string{}
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		if template != "" {
			input := strings.Split(scanner.Text(), " -> ")
			// Build replacement string
			to := fmt.Sprintf("%v!%v!%v", string(input[0][0]), input[1], string(input[0][1]))
			rules[input[0]] = &to
		} else {
			template = scanner.Text()
		}
	}

	for i := 0; i < 40; i++ {
		fmt.Println(i)
		for j := 0; j < len(template)-1; j++ {
			fmt.Println(j)
			pair := string(template[j]) + string(template[j+1])
			if rules[pair] != nil {
				template = template[0:j] + *rules[pair] + template[j+2:]
			}
		}
		template = strings.ReplaceAll(template, "!", "")
	}

	charCount := map[rune]int{}
	for _, r := range template {
		charCount[r]++
	}
	maxCharCount := 0
	minCharCount := math.MaxInt
	for _, c := range charCount {
		if c > maxCharCount {
			maxCharCount = c
		}
		if c < minCharCount {
			minCharCount = c
		}
	}
	fmt.Println(maxCharCount - minCharCount)
}
