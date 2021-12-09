package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type command struct {
	direction string
	units     int
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	commands := []command{}
	strCommands := strings.Split(string(file), "\n")
	for _, cstr := range strCommands {
		if cstr == "" {
			break
		}

		direction := strings.Split(cstr, " ")[0]
		units, err := strconv.Atoi(strings.Split(cstr, " ")[1])
		if err != nil {
			log.Fatal(err)
		}

		command := command{
			direction: direction,
			units:     units,
		}
		commands = append(commands, command)
	}

	hpos, depth, aim := 0, 0, 0
	for _, cmd := range commands {
		switch cmd.direction {
		case "forward":
			hpos += cmd.units
			depth += aim * cmd.units
		case "down":
			aim += cmd.units
		case "up":
			aim -= cmd.units
		default:
			log.Fatalf("Unrecognized direction %s", cmd.direction)
		}
	}

	fmt.Println(hpos * depth)
}
