package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/GabrielBalzer/aoc2023/util"
)

var (
	value_map = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
)

func extractNumberFromString(line string) int {
	var numbers []int
	for _, char := range line {
		number, _ := strconv.Atoi(string(char))
		if number != 0 {
			numbers = append(numbers, number)
		}
	}
	result, _ := strconv.Atoi(strconv.Itoa(numbers[0]) + strconv.Itoa(numbers[len(numbers)-1]))
	return result
}

func extractNumberAndStringNumberFromString(line string) int {
	numbers := make([]int, len(line))

	for i, char := range line {
		number, _ := strconv.Atoi(string(char))
		if number != 0 {
			numbers[i] = number
		}
	}
	for key, value := range value_map {
		re := regexp.MustCompile(key)
		matches := re.FindAllStringIndex(line, -1)
		if matches != nil {
			for _, match := range matches {
				numbers[match[0]] = value
			}
		}
	}

	numbers = util.DeleteNonZeroElements(numbers)

	result, _ := strconv.Atoi(strconv.Itoa(numbers[0]) + strconv.Itoa(numbers[len(numbers)-1]))
	return result
}

func solvePart1(lines []string) int {
	var sum int
	for _, line := range lines {
		sum += extractNumberFromString(line)
	}
	return sum
}

func solvePart2(lines []string) int {
	var sum int
	for _, line := range lines {
		sum += extractNumberAndStringNumberFromString(line)
	}
	return sum
}

func main() {
	lines, err := util.Read_input("day1.txt")
	if err != nil {
		panic(err)
	}
	part1_result := solvePart1(lines)
	fmt.Printf("Answer for Part 1: %v\n", part1_result)

	part2_result := solvePart2(lines)
	fmt.Printf("Answer for Part 2: %v\n", part2_result)
}
