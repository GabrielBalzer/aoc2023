package main

import "testing"

func TestExtractNumberFromString(t *testing.T) {
	test_cases := map[string]int{
		"1abc2":       12,
		"pqr3stu8vwx": 38,
		"a1b2c3d4e5f": 15,
		"treb7uchet":  77,
	}

	for test_word, number := range test_cases {
		result := extractNumberFromString(test_word)
		if result != number {
			t.Errorf("Result was incorrect, got: %d for %s, want: %d", result, test_word, number)
		}
	}
}

func TestExtractNumberAndStringNumberFromString(t *testing.T) {
	test_cases := map[string]int{
		"two1nine":                            29,
		"eightwothree":                        83,
		"abcone2threexyz":                     13,
		"xtwone3four":                         24,
		"4nineeightseven2":                    42,
		"zoneight234":                         14,
		"7pqrstsixteen":                       76,
		"six1mpffbnbnnlxthree":                63,
		"qpplcgvklzztqjkbbnfiveftcqmlrqnd824": 54,
		"fplzzhthreethree7sqnsmrljgsmnpktfkhzdpqfkone":      31,
		"pkpsgblkvsevennstsqvh3twofivekkhzzhqxjjsbvvvvnine": 79,
		"nrzpqk3fivesldclpcbfmdtbbhpxonethreeeightwor":      32,
		"896": 86,
		"6f":  66,
		"3seventwoseventcbsixllmhlmjbtqtfpfrjqbghd":            36,
		"toneightonelb4threetmldjmvjlfrnmkeightctjxdqhrcjczcr": 18,
	}

	for test_word, number := range test_cases {
		result := extractNumberAndStringNumberFromString(test_word)
		if result != number {
			t.Errorf("Result was incorrect, got: %d for %s, want: %d", result, test_word, number)
		}
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
