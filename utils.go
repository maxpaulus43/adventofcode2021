package main

import (
	"bufio"
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
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	result := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		check(err)
		result = append(result, num)
	}

	check(scanner.Err())

	return result
}

// wait for golang 1.18 to make something more generic

// func valuesFromFile(fileName string, conversionFunc func(string) (interface{}, error)) []interface{} {
// 	file, err := os.Open(fileName)
// 	check(err)
// 	defer file.Close()

// 	result := make([]interface{}, 0)
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		val, err := conversionFunc(scanner.Text())
// 		check(err)
// 		result = append(result, val)
// 	}

// 	check(scanner.Err())

// 	return result
// }
