package main

import (
	"strings"
)

func day11Part1() int {
	lines := linesFromFile("inputs/day11.txt")
	octopi := make(octopusGrid, 0, 10)
	for _, line := range lines {
		nums := stringsToInts(strings.Split(line, ""))
		octopi = append(octopi, nums)
	}

	totalFlashes := 0
	for i := 0; i < 100; i++ {
		n := octopi.step()
		totalFlashes += n
	}

	return totalFlashes
}

type octopusGrid [][]int

// returns how many flashes there were
func (g octopusGrid) step() int {
	for i := range g {
		for j := range (g)[0] {
			(g)[i][j] += 1
		}
	}

	didFlash := make([][]bool, len(g))
	for i := 0; i < len(g); i++ {
		didFlash[i] = make([]bool, len((g)[0]))
	}

	flashes := 0

	for g.shouldKeepFlashing() {
		for i := range g {
			for j := range (g)[0] {
				if (g)[i][j] > 9 {
					didFlash[i][j] = true
				}
			}
		}

		for i := range g {
			for j := range g[0] {
				if didFlash[i][j] && g[i][j] > 9 {
					g.flash(didFlash, i, j)
					flashes += 1
				}
			}
		}
	}
	return flashes
}

func (g octopusGrid) shouldKeepFlashing() bool {
	for i := range g {
		for j := range g[0] {
			if g[i][j] > 9 {
				return true
			}
		}
	}
	return false
}

func (g octopusGrid) flash(didFlash [][]bool, i, j int) {
	// update topleft
	if i > 0 && j > 0 && !didFlash[i-1][j-1] {
		g[i-1][j-1] += 1
	}
	// update topmid
	if i > 0 && !didFlash[i-1][j] {
		g[i-1][j] += 1
	}
	// update topright
	if i > 0 && j < len(g[0])-1 && !didFlash[i-1][j+1] {
		g[i-1][j+1] += 1
	}
	// update leftleft
	if j > 0 && !didFlash[i][j-1] {
		g[i][j-1] += 1
	}
	// update rightright
	if j < len(g[0])-1 && !didFlash[i][j+1] {
		g[i][j+1] += 1
	}
	// update bottomleft
	if i < len(g)-1 && j > 0 && !didFlash[i+1][j-1] {
		g[i+1][j-1] += 1
	}
	// update bottommid
	if i < len(g)-1 && !didFlash[i+1][j] {
		g[i+1][j] += 1
	}
	// update bottomright
	if i < len(g)-1 && j < len(g[0])-1 && !didFlash[i+1][j+1] {
		g[i+1][j+1] += 1
	}

	g[i][j] = 0
}

func day11Part2() int {
	lines := linesFromFile("inputs/day11.txt")
	octopi := make(octopusGrid, 0, 10)
	for _, line := range lines {
		nums := stringsToInts(strings.Split(line, ""))
		octopi = append(octopi, nums)
	}

	i := 1
	for {
		n := octopi.step()
		if n == len(octopi)*len(octopi[0]) {
			break
		}
		i++
	}

	return i
}
