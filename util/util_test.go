package util

import (
	"reflect"
	"testing"
)

func TestReadInput(t *testing.T) {
	// Test case 1: Provide an existing file
	expected := []string{"line1", "line2", "line3"}
	result, err := Read_input("testfile.txt")

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	// Test case 2: Provide a non-existing file
	_, err = Read_input("nonexistent.txt")

	if err == nil {
		t.Error("Expected an error, but got none")
	}
}

func TestDeleteNonZeroElements(t *testing.T) {
	input := []int{0, 5, 0, 0, 3, 0, 0, 0, 9, 12, 0, 0}
	expected := []int{5, 3, 9, 12}

	result := DeleteNonZeroElements(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
