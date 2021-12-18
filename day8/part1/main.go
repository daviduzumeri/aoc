package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	matches := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		for i := 11; i < 15; i++ {
			length := len(fields[i])
			if length == 2 || length == 3 || length == 4 || length == 7 {
				matches++
			}
		}
	}
	fmt.Println(matches)
}
