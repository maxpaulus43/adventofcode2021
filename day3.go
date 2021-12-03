package main

import (
	"strconv"
)

func day3Part1() int64 {
	nums := linesFromFile("inputs/day3.txt")
	numLen := len(nums[0])
	gammaRate := ""
	epsilonRate := ""

	bitCountForPos := make([]map[byte]int, numLen)
	for i := range bitCountForPos {
		bitCountForPos[i] = make(map[byte]int)
	}

	for pos := 0; pos < numLen; pos++ {
		count := bitCountForPos[pos]

		for _, num := range nums {
			bit := num[pos]
			count[bit]++
		}

		if count['1'] > count['0'] {
			gammaRate += "1"
			epsilonRate += "0"
		} else {
			gammaRate += "0"
			epsilonRate += "1"
		}
	}

	gammaRateDecimal, err := strconv.ParseInt(gammaRate, 2, 0)
	check(err)
	epsilonRateDecimal, err := strconv.ParseInt(epsilonRate, 2, 0)
	check(err)

	return gammaRateDecimal * epsilonRateDecimal
}

type conditionFunc func(int, int) bool

func day3Part2() int64 {
	numsFromFile := linesFromFile("inputs/day3.txt")

	findNumsSatifyingCondition := func(cond conditionFunc) string {
		nums := numsFromFile
		pos := 0

		for len(nums) > 1 {
			numsWithOnes := make([]string, 0)
			numsWithZeroes := make([]string, 0)

			for _, n := range nums {
				if n[pos] == '0' {
					numsWithZeroes = append(numsWithZeroes, n)
				} else {
					numsWithOnes = append(numsWithOnes, n)
				}
			}

			if cond(len(numsWithZeroes), len(numsWithOnes)) {
				nums = numsWithZeroes
			} else {
				nums = numsWithOnes
			}

			pos++
		}
		return nums[0]
	}

	o2Rating := findNumsSatifyingCondition(func(a int, b int) bool { return a > b })
	co2Rating := findNumsSatifyingCondition(func(a int, b int) bool { return a <= b })

	o2RatingDecimal, err := strconv.ParseInt(o2Rating, 2, 0)
	check(err)
	co2RatingDecimal, err := strconv.ParseInt(co2Rating, 2, 0)
	check(err)

	return o2RatingDecimal * co2RatingDecimal
}
