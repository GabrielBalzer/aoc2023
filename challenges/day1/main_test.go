package main

import "testing"

func TestExtractNumberFromString(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"1abc2 should be 12", "1abc2", 12},
		{"pqr3stu8vwx should be 38", "pqr3stu8vwx", 38},
		{"a1b2c3d4e5f should be 15", "a1b2c3d4e5f", 15},
		{"treb7uchet should be 77", "treb7uchet", 77},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractNumberFromString(tt.input)
			if result != tt.want {
				t.Errorf("got %v, want %v", result, tt.want)
			}
		})
	}
}

func TestExtractNumberAndStringNumberFromString(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		number int
	}{
		{"two1nine should be 29", "two1nine", 29},
		{"eightwothree should be 83", "eightwothree", 83},
		{"abcone2threexyz should be 13", "abcone2threexyz", 13},
		{"xtwone3four should be 24", "xtwone3four", 24},
		{"4nineeightseven2 should be 42", "4nineeightseven2", 42},
		{"zoneight234 should be 14", "zoneight234", 14},
		{"7pqrstsixteen should be 76", "7pqrstsixteen", 76},
		{"six1mpffbnbnnlxthree should be 63", "six1mpffbnbnnlxthree", 63},
		{"qpplcgvklzztqjkbbnfiveftcqmlrqnd824 should be 54", "qpplcgvklzztqjkbbnfiveftcqmlrqnd824", 54},
		{"fplzzhthreethree7sqnsmrljgsmnpktfkhzdpqfkone should be 31", "fplzzhthreethree7sqnsmrljgsmnpktfkhzdpqfkone", 31},
		{"pkpsgblkvsevennstsqvh3twofivekkhzzhqxjjsbvvvvnine should be 79", "pkpsgblkvsevennstsqvh3twofivekkhzzhqxjjsbvvvvnine", 79},
		{"nrzpqk3fivesldclpcbfmdtbbhpxonethreeeightwor should be 32", "nrzpqk3fivesldclpcbfmdtbbhpxonethreeeightwor", 32},
		{"896 should be 86", "896", 86},
		{"6f should be 66", "6f", 66},
		{"3seventwoseventcbsixllmhlmjbtqtfpfrjqbghd should be 36", "3seventwoseventcbsixllmhlmjbtqtfpfrjqbghd", 36},
		{"toneightonelb4threetmldjmvjlfrnmkeightctjxdqhrcjczcr should be 18", "toneightonelb4threetmldjmvjlfrnmkeightctjxdqhrcjczcr", 18},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractNumberAndStringNumberFromString(tt.input)
			if result != tt.number {
				t.Errorf("got %d, want %d for %s", result, tt.number, tt.name)
			}
		})
	}
}

func TestSolvePart1(t *testing.T) {
	test_input := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}
	desired_result := 142

	result := solvePart1(test_input)
	if result != desired_result {
		t.Errorf("Result was incorrect, got: %d for %s, want: %d", result, test_input, desired_result)
	}
}

func TestSolvePart2(t *testing.T) {
	test_input := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}
	desired_result := 281

	result := solvePart2(test_input)
	if result != desired_result {
		t.Errorf("Result was incorrect, got: %d for %s, want: %d", result, test_input, desired_result)
	}
}
