package main

import (
	"math"
	"sort"
	"strings"
)

func day7Part1() int {
	crabPositions := stringsToInts(strings.Split(linesFromFile("inputs/day7.txt")[0], ","))

	sort.Ints(crabPositions)
	medianPos := crabPositions[len(crabPositions)/2-1]

	fuelNeeded := 0
	for _, pos := range crabPositions {
		diff := int(math.Abs(float64(pos - medianPos)))
		fuelNeeded += diff
	}
	return fuelNeeded
}

func day7Part2() int {
	crabPositions := stringsToInts(strings.Split(linesFromFile("inputs/day7.txt")[0], ","))
	avg := avg(crabPositions...)
	truncatedAvg := int(avg)
	roundedAvg := int(math.Round(avg))

	fuelNeeded := 0
	fuelNeeded2 := 0
	for _, pos := range crabPositions {
		dist := abs(pos - truncatedAvg)
		fuelForDist := (dist + 1) * dist / 2
		fuelNeeded += fuelForDist

		dist = abs(pos - roundedAvg)
		fuelForDist = (dist + 1) * dist / 2
		fuelNeeded2 += fuelForDist
	}

	if fuelNeeded2 < fuelNeeded {
		fuelNeeded = fuelNeeded2
	}

	return fuelNeeded
}
