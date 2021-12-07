package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func linesFromFile(fileName string) []string {
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	result := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	check(scanner.Err())

	return result
}

func numsFromFile(fileName string) []int {
	return stringsToInts(linesFromFile(fileName))
}

func stringsToInts(strings []string) []int {
	result := make([]int, 0, len(strings))
	for _, s := range strings {
		n, err := strconv.Atoi(s)
		check(err)
		result = append(result, n)
	}
	return result
}

func abs(n int) int {
	return int(math.Abs(float64(n)))
}

func sum(nums ...int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func avg(nums ...int) float64 {
	return float64(sum(nums...)) / float64(len(nums))
}
