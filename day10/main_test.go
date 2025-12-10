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
func TestPart2(t *testing.T) {
	lines, err := common.ReadLines("test_input.txt")
	if err != nil {
		t.Fatalf("Failed to read example input: %v", err)
	}

	result := part2(lines)
	expected := 33

	if result != expected {
		t.Errorf("part1() = %d, want %d", result, expected)
	}
}
