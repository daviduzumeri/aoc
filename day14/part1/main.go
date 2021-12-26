package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type rule struct {
	from string
	to   string
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	template := ""
	rules := []rule{}
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		if template != "" {
			input := strings.Split(scanner.Text(), " -> ")
			// Build replacement string
			to := fmt.Sprintf("%v!%v!%v", string(input[0][0]), input[1], string(input[0][1]))
			rules = append(rules, rule{from: input[0], to: to})
		} else {
			template = scanner.Text()
		}
	}

	for i := 0; i < 10; i++ {
		for _, rule := range rules {
			for template != strings.ReplaceAll(template, rule.from, rule.to) {
				template = strings.ReplaceAll(template, rule.from, rule.to)
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
