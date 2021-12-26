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
	rules := map[string]string{}
	elements := map[string]int{}
	pairs := map[string]int{}
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		if template != "" {
			input := strings.Split(scanner.Text(), " -> ")
			rules[input[0]] = input[1]
		} else {
			template = scanner.Text()
		}
	}

	for i := range template {
		elements[string(template[i])]++
		if i == len(template)-1 {
			break
		}
		pairs[string(template[i])+string(template[i+1])]++
	}

	for i := 0; i < 40; i++ {
		newPairs := map[string]int{}
		for k, v := range pairs {
			newPairs[k] = v
		}
		for from, to := range rules {
			newPairs[from[:1]+to] += pairs[from]
			newPairs[to+from[1:]] += pairs[from]
			elements[to] += pairs[from]
			newPairs[from] -= pairs[from]
		}
		pairs = newPairs
	}

	maxCharCount := 0
	minCharCount := math.MaxInt
	for _, c := range elements {
		if c > maxCharCount {
			maxCharCount = c
		}
		if c < minCharCount {
			minCharCount = c
		}
	}
	fmt.Println(maxCharCount - minCharCount)
}
