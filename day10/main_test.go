package main

import (
	"advent-of-code-2025/common"
	"testing"
)

func TestPart1(t *testing.T) {
	lines, err := common.ReadLines("test_input.txt")
	if err != nil {
		t.Fatalf("Failed to read example input: %v", err)
	}

	result := part1(lines)
	expected := 7

	if result != expected {
		t.Errorf("part1() = %d, want %d", result, expected)
	}
}
func TestFirstLineJoltage(t *testing.T) {
	// First line: [##.....#.] (0,1,5,8) (1,6,7) (3,6,8) (1,3,6,7) (0,1,2,6,7) (1,2,3,5,7) (0,1,3,4,5,6,7) (1,2,4,5,7,8) (0,2,5,7,8) (1,2,3,5,7,8) {53,78,43,44,33,73,46,81,60}
	line := "[##.....#.] (0,1,5,8) (1,6,7) (3,6,8) (1,3,6,7) (0,1,2,6,7) (1,2,3,5,7) (0,1,3,4,5,6,7) (1,2,4,5,7,8) (0,2,5,7,8) (1,2,3,5,7,8) {53,78,43,44,33,73,46,81,60}"

	machine := parseSchematic(line)
	result := findFewestButtonClicksForJoltageMode(machine)

	// The maximum joltage requirement is 81, so the result must be at least 81
	minExpected := 81
	if result < minExpected {
		t.Errorf("findFewestButtonClicksForJoltageMode() = %d, want at least %d", result, minExpected)
	}

	// Also verify it's not 0 (which indicates no solution found)
	if result == 0 {
		t.Errorf("findFewestButtonClicksForJoltageMode() returned 0, indicating no solution found")
	}
}

func TestPart1ForReal(t *testing.T) {
	lines, err := common.ReadLines("input.txt")
	if err != nil {
		t.Fatalf("Failed to read example input: %v", err)
	}
	result := part1(lines)
	expected := 558
	if result != expected {
		t.Errorf("part1() = %d, want %d", result, expected)
	}
}
