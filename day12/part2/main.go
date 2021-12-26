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

type path struct {
	caves         []string
	revisitedCave bool
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

func (cm caveMap) findPaths(c *cave, p path) []path {
	paths := []path{}
	if c == nil {
		c = cm["start"]
	}
	if c.name == "end" {
		return []path{
			{
				caves:         append(p.caves, c.name),
				revisitedCave: p.revisitedCave,
			},
		}
	}
	if unicode.IsLower(rune(c.name[0])) {
		for _, pc := range p.caves {
			if pc == c.name {
				if p.revisitedCave || pc == "start" {
					return []path{}
				} else {
					p.revisitedCave = true
				}
			}
		}
	}
	for _, cp := range c.caves {
		caves := append(p.caves, c.name)
		newPaths := cm.findPaths(
			cp,
			path{
				caves:         caves,
				revisitedCave: p.revisitedCave,
			},
		)
		paths = append(paths, newPaths...)
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

	paths := caves.findPaths(nil, path{})

	fmt.Println(len(paths))
}
