package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/GabrielBalzer/aoc2023/util"
)

var (
	possible_part1 = map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
)

func solvePart1(lines []string) int {
	sum_ids := 0

	for line_number, line := range lines {
		game_id := line_number + 1
		game_data := strings.Split(line, ": ")[1]
		sets := strings.Split(game_data, "; ")
		game_possible := true
		for _, set := range sets {
			set_possible := map[string]bool{
				"red":   true,
				"green": true,
				"blue":  true,
			}
			colors := strings.Split(set, ", ")
			for _, color := range colors {

				color_value_string_part := strings.Split(color, " ")
				color_value, err := strconv.Atoi(color_value_string_part[0])
				if err != nil {
					panic(err)
				}
				for color_name, color_possible := range possible_part1 {
					if strings.Contains(color, color_name) {
						if !(color_value <= color_possible) {
							set_possible[color_name] = false
						}
					}
				}
			}
			for _, v := range set_possible {
				if !v {
					game_possible = false
					break
				}
			}
		}
		if game_possible {
			sum_ids += game_id
		}
	}
	return sum_ids
}

func main() {
	lines, err := util.Read_input("inputs/day2.txt")
	if err != nil {
		panic(err)
	}
	part1_result := solvePart1(lines)
	fmt.Printf("Answer for Part 1: %v\n", part1_result)
}
