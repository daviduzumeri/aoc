package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

type cave struct {
	name  string
	caves []*cave
}

type caveMap map[string]*cave

func (cm caveMap) getCave(name string) *cave {
	c, ok := cm[name]
	if !ok {
		cm[name] = &cave{name: name}
		c = cm[name]
	}
	return c

}

func (cm caveMap) findPaths(c *cave, path []string) [][]string {
	paths := [][]string{}
	if c == nil {
		c = cm["start"]
	}
	if c.name == "end" {
		return [][]string{append(path, c.name)}
	}
	if unicode.IsLower(rune(c.name[0])) {
		for _, pc := range path {
			if pc == c.name {
				return [][]string{}
			}
		}
	}
	for _, cp := range c.caves {
		paths = append(paths, cm.findPaths(cp, append(path, c.name))...)
	}
	return paths
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

	paths := caves.findPaths(nil, []string{})

	fmt.Println(len(paths))
}
