package main

import (
	"sort"
	"strconv"
	"strings"
)

func day8Part1() int {
	lines := linesFromFile("inputs/day8.txt")
	outputDigits := make([][]string, 0)
	for _, l := range lines {
		split := strings.Split(l, " | ")
		outputDigits = append(outputDigits, strings.Split(split[1], " "))
	}

	result := 0
	for _, digits := range outputDigits {
		for _, digit := range digits {
			l := len(digit)
			if l == 2 || l == 3 || l == 4 || l == 7 {
				result += 1
			}
		}
	}

	return result
}

func day8Part2() int {
	lines := linesFromFile("inputs/day8.txt")

	result := 0
	for _, l := range lines {
		split := strings.Split(l, " | ")

		keyDigits := strings.Split(split[0], " ")
		keyDigitsSorted := make([]string, 0, len(keyDigits))
		for _, elem := range keyDigits {
			r := runes(elem)
			sort.Sort(r)
			keyDigitsSorted = append(keyDigitsSorted, string(r))
		}

		outputDigits := strings.Split(split[1], " ")
		outputDigitsSorted := make([]string, 0, len(outputDigits))
		for _, elem := range outputDigits {
			r := runes(elem)
			sort.Sort(r)
			outputDigitsSorted = append(outputDigitsSorted, string(r))
		}

		result += parseDigitsWithKeyDigits(outputDigitsSorted, keyDigitsSorted)
	}

	return result
}

func parseDigitsWithKeyDigits(outputDigits, keyDigits []string) int {
	digitForKey := make(map[string]string)
	keyForDigit := make(map[string]string)

	link := func(digitStr string, numStr string) {
		digitForKey[digitStr] = numStr
		keyForDigit[numStr] = digitStr
	}

	for _, digitStr := range keyDigits {
		switch len(digitStr) {
		case 2:
			link(digitStr, "1")
		case 3:
			link(digitStr, "7")
		case 4:
			link(digitStr, "4")
		case 7:
			link(digitStr, "8")
		}
	}

	for _, digitStr := range keyDigits {
		switch {
		case len(digitStr) == 5 && encompasses(digitStr, keyForDigit["1"]):
			link(digitStr, "3")
		case len(digitStr) == 6 && !encompasses(digitStr, keyForDigit["1"]):
			link(digitStr, "6")
		case len(digitStr) == 6 && encompasses(digitStr, keyForDigit["4"]):
			link(digitStr, "9")
		}
	}

	for _, digitStr := range keyDigits {
		switch {
		case len(digitStr) == 5 && encompasses(keyForDigit["6"], digitStr):
			link(digitStr, "5")
		case len(digitStr) == 6 && digitForKey[digitStr] == "":
			link(digitStr, "0")
		}
	}

	for _, digitStr := range keyDigits {
		if digitForKey[digitStr] == "" {
			link(digitStr, "2")
		}
	}

	result := ""
	for _, digitStr := range outputDigits {
		result += digitForKey[digitStr]
	}
	n, err := strconv.Atoi(result)
	check(err)
	return n
}

func encompasses(outer, inner string) bool {
	for _, innerC := range inner {
		doesContain := false
		for _, outerC := range outer {
			if outerC == innerC {
				doesContain = true
				break
			}
		}
		if !doesContain {
			return false
		}
	}
	return true
}

type runes []rune

func (r runes) Len() int           { return len(r) }
func (r runes) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r runes) Less(i, j int) bool { return r[i] < r[j] }
