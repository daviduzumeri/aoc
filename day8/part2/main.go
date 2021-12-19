package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func diffStrings(s1 string, s2 string) string {
	var returnValue []rune
	for _, r := range s1 {
		if !strings.ContainsRune(s2, r) {
			returnValue = append(returnValue, r)
		}
	}
	return string(returnValue)
}

func findSegmentIndex(s string, segments [10]string) (int, error) {
	for i, segment := range segments {
		if segment == s {
			return i, nil
		}
	}
	return -1, errors.New("string not in segment array")
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	total := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		segments := [10]string{}
		sortedFields := [14]string{}
		i := 0
		for _, field := range fields {
			if field == "|" {
				continue
			}
			field = sortString(field)
			sortedFields[i] = field
			i++
			switch len(field) {
			case 2:
				segments[1] = field
			case 3:
				segments[7] = field
			case 4:
				segments[4] = field
			case 7:
				segments[8] = field
			}
		}

		segmentBD := diffStrings(segments[4], segments[1])

		for _, field := range sortedFields {
			switch len(field) {
			case 5:
				// 2, 3, 5
				if len(diffStrings(field, segments[1])) == 3 {
					segments[3] = field
				} else if len(diffStrings(field, segmentBD)) == 3 {
					segments[5] = field
				} else {
					segments[2] = field
				}
			case 6:
				// 0, 6, 9
				if len(diffStrings(field, segments[4])) == 2 {
					segments[9] = field
				} else if len(diffStrings(field, segments[1])) == 4 {
					segments[0] = field
				} else {
					segments[6] = field
				}
			}
		}

		amount := ""
		for i := 10; i < 14; i++ {
			segmentIndex, err := findSegmentIndex(sortedFields[i], segments)
			if err != nil {
				log.Fatal(err)
			}
			amount += fmt.Sprintf("%d", segmentIndex)
		}
		amountInt, err := strconv.Atoi(amount)
		if err != nil {
			log.Fatal(err)
		}
		total += amountInt
	}
	fmt.Println(total)
}
