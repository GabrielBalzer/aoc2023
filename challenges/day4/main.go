package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/GabrielBalzer/aoc2023/util"
)

func solvePart1(lines []string) int {
	cardWorthSum := 0
	for _, line := range lines {
		cardWorth := 0
		cardData := strings.Split(line, ": ")

		numberDataSingleSpace := strings.Join(strings.Fields(cardData[1]), " ")
		numberData := strings.Split(numberDataSingleSpace, " | ")
		winningNumbers := numberData[0]
		actualNumbers := numberData[1]
		winningNumberSlice := convertStringToIntSlice(winningNumbers)
		actualNumberSlice := convertStringToIntSlice(actualNumbers)
		for _, actualNumber := range actualNumberSlice {
			if contains(winningNumberSlice, actualNumber) {
				if cardWorth == 0 {
					cardWorth = 1
				} else {
					cardWorth *= 2
				}
			}
		}
		cardWorthSum += cardWorth
	}
	return cardWorthSum
}

func convertStringToIntSlice(winningNumbers string) []int {
	StringSlice := strings.Split(winningNumbers, " ")
	NumberSlice := make([]int, len(StringSlice))
	for i, s := range StringSlice {
		NumberSlice[i], _ = strconv.Atoi(s)
	}
	return NumberSlice
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	lines, err := util.Read_input("inputs/day4.txt")
	if err != nil {
		panic(err)
	}
	part1Result := solvePart1(lines)
	fmt.Printf("Answer for Part 1: %v\n", part1Result)
	// part2Result := solvePart2(lines)
	// fmt.Printf("Answer for Part 2: %v\n", part2Result)
}
