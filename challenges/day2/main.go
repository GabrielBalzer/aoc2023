package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/GabrielBalzer/aoc2023/util"
)

var (
	possiblePart1 = map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
)

func solvePart1(lines []string) int {
	sumIDs := 0

	for lineNumber, line := range lines {
		gameID := lineNumber + 1
		gameData := strings.Split(line, ": ")[1]
		sets := strings.Split(gameData, "; ")
		gamePossible := true
		for _, set := range sets {
			setPossible := map[string]bool{
				"red":   true,
				"green": true,
				"blue":  true,
			}
			colors := strings.Split(set, ", ")
			for _, color := range colors {

				colorValueStringPart := strings.Split(color, " ")
				colorValue, err := strconv.Atoi(colorValueStringPart[0])
				if err != nil {
					panic(err)
				}
				for colorName, colorPossible := range possiblePart1 {
					if strings.Contains(color, colorName) {
						if !(colorValue <= colorPossible) {
							setPossible[colorName] = false
						}
					}
				}
			}
			for _, v := range setPossible {
				if !v {
					gamePossible = false
					break
				}
			}
		}
		if gamePossible {
			sumIDs += gameID
		}
	}
	return sumIDs
}

func solvePart2(lines []string) int {
	sumProducts := 0
	for _, line := range lines {
		colorCounts := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		gameData := strings.Split(line, ": ")[1]
		sets := strings.Split(gameData, "; ")
		for _, set := range sets {

			colors := strings.Split(set, ", ")
			for _, color := range colors {

				colorValueStringPart := strings.Split(color, " ")
				colorValue, err := strconv.Atoi(colorValueStringPart[0])
				if err != nil {
					panic(err)
				}
				for colorName, colorMinimal := range colorCounts {
					if strings.Contains(color, colorName) {
						if colorMinimal < colorValue {
							colorCounts[colorName] = colorValue
						}
					}
				}
			}
		}
		sumProducts += colorCounts["red"] * colorCounts["green"] * colorCounts["blue"]

	}
	return sumProducts
}

func main() {
	lines, err := util.Read_input("inputs/day2.txt")
	if err != nil {
		panic(err)
	}
	part1Result := solvePart1(lines)
	fmt.Printf("Answer for Part 1: %v\n", part1Result)
	part2Result := solvePart2(lines)
	fmt.Printf("Answer for Part 2: %v\n", part2Result)
}
