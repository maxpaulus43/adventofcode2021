package main

import (
	"bufio"
	_ "embed"
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

// wait for golang 1.18 to make something more generic

// func mapValuesFromFile[T any](fileName string, fn func(string) (T, error)) []T {
// 	file, err := os.Open(fileName)
// 	check(err)
// 	defer file.Close()

// 	result := make([]T, 0)
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		val, err := fn(scanner.Text())
// 		check(err)
// 		result = append(result, val)
// 	}

// 	check(scanner.Err())

// 	return result
// }

// func Map[T any, K any](fn func(T) K, arr []T) []K {
// 	result := make([]K, 0, len(arr))
// 	for _, elem := range arr {
// 		result = append(result, fn(elem))
// 	}
// 	return result
// }
