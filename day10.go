package main

import "sort"

func isOpeningBracket(r rune) bool {
	return r == '{' || r == '[' || r == '(' || r == '<'
}

func day10part1() int {
	lines := linesFromFile("inputs/day10.txt")
	scoreForRune := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	partnerForRune := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
		'>': '<',
	}

	score := 0
	for _, line := range lines {
		stk := make(stack, 0)
		for _, r := range line {
			if isOpeningBracket(r) {
				stk.push(r)
			} else if stk.peek() == partnerForRune[r] {
				stk.pop()
			} else {
				score += scoreForRune[r]
				break
			}
		}
	}

	return score
}

func day10part2() int {
	lines := linesFromFile("inputs/day10.txt")
	scoreForRune := map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}
	partnerForRune := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
		')': '(',
		']': '[',
		'}': '{',
		'>': '<',
	}

	scores := make([]int, 0)
outer:
	for _, line := range lines {
		score := 0
		stk := make(stack, 0)
		for _, r := range line {
			if isOpeningBracket(r) {
				stk.push(r)
			} else if stk.peek() == partnerForRune[r] {
				stk.pop()
			} else {
				// error found, skip the whole line
				continue outer
			}
		}
		completionString := make([]rune, 0, len(stk))
		for len(stk) > 0 {
			completionString = append(completionString, partnerForRune[stk.pop()])
		}
		for _, r := range completionString {
			score *= 5
			score += scoreForRune[r]
		}
		scores = append(scores, score)
	}

	sort.Ints(scores)

	return scores[len(scores)/2]
}
