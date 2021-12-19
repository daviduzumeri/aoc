package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"unicode"
)

type cave struct {
	name    string
	isBig   bool
	visited bool
	caves   []*cave
}

type caveMap map[string]*cave

func (cm caveMap) getCave(name string) *cave {
	c, ok := cm[name]
	if !ok {
		cm[name] = &cave{
			name:  name,
			isBig: unicode.IsUpper(rune(name[0])),
		}
		c = cm[name]
	}
	return c

}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	caves := make(caveMap)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "-")
		c1 := caves.getCave(row[0])
		c2 := caves.getCave(row[1])
		c1.caves = append(c1.caves, c2)
		c2.caves = append(c2.caves, c1)
	}

	start := caves["start"]
}
