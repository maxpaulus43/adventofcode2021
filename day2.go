package main

import (
	"strconv"
	"strings"
)

const (
	DIR_FORWARD  = "forward"
	DIR_BACKWARD = "backward"
	DIR_DOWN     = "down"
	DIR_UP       = "up"
	DIR_LEFT     = "left"
	DIR_RIGHT    = "right"
)

type instruction struct {
	direction string
	distance  int
}

func instructionFromLine(line string) instruction {
	split := strings.Split(line, " ")
	d, err := strconv.Atoi(split[1])
	check(err)
	return instruction{direction: split[0], distance: d}
}

func makeInstructions() []instruction {
	lines := linesFromFile("inputs/day2.txt")
	instructions := make([]instruction, len(lines))
	for _, line := range lines {
		instructions = append(instructions, instructionFromLine(line))
	}
	return instructions
}

func day2Part1() int {
	x := 0
	y := 0 // y increases in the downward direction
	instructions := makeInstructions()

	for _, inst := range instructions {
		switch inst.direction {
		case DIR_FORWARD:
			x += inst.distance
		case DIR_UP:
			y -= inst.distance
		case DIR_DOWN:
			y += inst.distance
		}
	}

	return x * y
}

func day2Part2() int {
	x := 0
	y := 0
	aim := 0 // aim increases in the downward direction
	instructions := makeInstructions()

	for _, inst := range instructions {
		switch inst.direction {
		case DIR_FORWARD:
			{
				x += inst.distance
				y += aim * inst.distance
			}
		case DIR_UP:
			aim -= inst.distance
		case DIR_DOWN:
			aim += inst.distance
		}
	}

	return x * y
}
