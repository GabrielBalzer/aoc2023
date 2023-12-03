package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/GabrielBalzer/aoc2023/util"
)

func solvePart1(lines []string) int {
	sumPartNumbers := 0
	var numberIndexesCurrentLine [][]int
	var symbolIndexesCurrentLine [][]int
	var symbolIndexesNextLine [][]int
	var symbolIndexesPreviousLine [][]int
	regexNumbers := regexp.MustCompile(`\w+`)
	regexSymbol := regexp.MustCompile("[^.0-9]")
	for i, line := range lines {
		symbolIndexesNextLine = [][]int{}
		symbolIndexesPreviousLine = [][]int{}

		numberIndexesCurrentLine = regexNumbers.FindAllStringIndex(line, -1)
		symbolIndexesCurrentLine = regexSymbol.FindAllStringIndex(line, -1)

		if len(lines) > i+1 {
			symbolIndexesNextLine = regexSymbol.FindAllStringIndex(lines[i+1], -1)
		}
		if i > 0 {
			symbolIndexesPreviousLine = regexSymbol.FindAllStringIndex(lines[i-1], -1)
		}

		for _, numberIndex := range numberIndexesCurrentLine {

			validPart := false
			for _, symbolIndex := range symbolIndexesCurrentLine {
				if numberIndex[1] == symbolIndex[0] {
					validPart = true
				}
				if numberIndex[0] == symbolIndex[1] {
					validPart = true
				}

			}
			validPart = checkSymbolsInAdjactentLine(symbolIndexesNextLine, numberIndex, validPart)
			validPart = checkSymbolsInAdjactentLine(symbolIndexesPreviousLine, numberIndex, validPart)
			if validPart {
				validPartNumber, _ := strconv.Atoi(line[numberIndex[0]:numberIndex[1]])
				sumPartNumbers += validPartNumber
			}
		}
	}

	return sumPartNumbers
}

func checkSymbolsInAdjactentLine(symbolIndexes [][]int, numberIndex []int, validPart bool) bool {
	for _, nextSymbolIndex := range symbolIndexes {
		if numberIndex[1] == nextSymbolIndex[0] {
			validPart = true
		}
		if numberIndex[0] == nextSymbolIndex[1] {
			validPart = true
		}
		if numberIndex[0] <= nextSymbolIndex[1] && nextSymbolIndex[0] <= numberIndex[1] {
			validPart = true
		}
	}
	return validPart
}

func main() {
	lines, err := util.Read_input("inputs/day3.txt")
	if err != nil {
		panic(err)
	}
	part1Result := solvePart1(lines)
	fmt.Printf("Answer for Part 1: %v\n", part1Result)

}
