package main

const inputFile = "inputs/day1.txt"

func day1Part1() int {
	depths := numsFromFile(inputFile)

	depthIncreased := 0
	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			depthIncreased++
		}
	}
	return depthIncreased
}

func day1Part2() int {
	depths := numsFromFile(inputFile)

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
