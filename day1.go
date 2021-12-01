package main

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed inputs/day1.txt
var day1Txt string

func makeDepths() []int {
	depths := make([]int, 0)
	for _, val := range strings.Split(day1Txt, "\n") {
		n, _ := strconv.Atoi(val)
		depths = append(depths, n)
	}
	return depths
}

func day1Part1() int {
	depths := makeDepths()

	depthIncreased := 0
	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			depthIncreased++
		}
	}
	return depthIncreased
}

func day1Part2() int {
	depths := makeDepths()

	depthIncreased := 0
	for i := 3; i < len(depths); i++ {
		prevSum := depths[i-3] + depths[i-2] + depths[i-1]
		currSum := depths[i-2] + depths[i-1] + depths[i]

		if currSum > prevSum {
			depthIncreased++
		}
	}

	return depthIncreased
}
