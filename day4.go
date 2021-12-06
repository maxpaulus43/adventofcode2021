package main

// import (
// 	"strings"
// )

// type spot struct {
// 	value    int
// 	isMarked bool
// }

// type pos struct {
// 	row int
// 	col int
// }

// type board struct {
// 	inner               [][]spot
// 	leastRecentlyMarked pos
// }

// func (b board) mark(p pos) {
// 	b.inner[p.row][p.col].isMarked = true
// 	b.leastRecentlyMarked = p
// }

// func (b board) checkForWin() bool {
// 	columnDidFail := make([]bool, len(b.inner))
// 	for i := range b.inner {
// 		for j := range b.inner {
// 			if !b.inner[i][j].isMarked {
// 				columnDidFail[j] = true
// 				break
// 			}
// 			if j == len(b.inner)-1 {
// 				return true // the whole row was marked
// 			}
// 		}
// 	}

// 	for _, c := range columnDidFail {
// 		if !c {
// 			return true // the whole column did not fail
// 		}
// 	}

// 	return false
// }

// func day4Part1() int {
// 	lines := linesFromFile("inputs/day4.txt")
// 	nums := stringsToInts(strings.Split(lines[0], ","))
// 	lines = lines[2:]

// 	boards := make([]board, 0, len(lines)/6)
// 	currentBoard := board{inner: make([][]spot, 0)}
// 	currentRow := 0
// 	for _, line := range lines {
// 		if line == "\n" {
// 			boards = append(boards, currentBoard)
// 			currentBoard = board{inner: make([][]spot, 0)}
// 			currentRow = 0
// 		} else {
// 			nums := stringsToInts(strings.Split(line, " "))
// 			spots := make()
// 			currentBoard.inner = append(currentBoard.inner, row)
// 		}
// 	}
// }
