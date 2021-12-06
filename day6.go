package main

import (
	"fmt"
	"strings"
)

type timer struct {
	timer int
}

func (t timer) String() string {
	return fmt.Sprintf("%v", t.timer)
}

func (t *timer) tick() {
	t.timer--
}

func (t *timer) reset() {
	t.timer = 6
}

func parseFile(fileName string) []timer {
	nums := stringsToInts(strings.Split(linesFromFile(fileName)[0], ","))
	result := make([]timer, 0, len(nums))
	for _, n := range nums {
		result = append(result, timer{timer: n})
	}
	return result
}

func day6Part1(days int) int {
	fish := parseFile("inputs/day6.txt")

	for i := 0; i < days; i++ {
		fishLen := len(fish)
		for j := 0; j < fishLen; j++ {
			f := &fish[j]
			f.tick()
			if f.timer < 0 {
				f.reset()
				fish = append(fish, timer{timer: 8})
			}
		}
	}
	return len(fish)
}

func day6Part2(days int) int {
	fish := stringsToInts(strings.Split(linesFromFile("inputs/day6.txt")[0], ","))
	count := make([]int, 9)

	// count how many of each kind of fish there are in the beginning
	for _, f := range fish {
		count[f] += 1
	}

	// count[x - 1] on day 2 is equal to count[x] on day 1. the count array shifts left every day
	for i := 0; i < days; i++ {
		resetCount := count[0]
		for j := 1; j < 9; j++ {
			count[j-1] = count[j]
		}
		// fish reset back to value 6
		count[6] += resetCount
		// fish who reset spawn a new fish with value 8
		count[8] = resetCount
	}

	sum := 0
	for k := range count {
		sum += count[k]
	}
	return sum
}
