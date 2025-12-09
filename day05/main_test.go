package main

import (
	"advent-of-code-2025/common"
	"testing"
)

func TestPart1(t *testing.T) {
	lines, err := common.ReadLines("part1example.txt")
	if err != nil {
		t.Fatalf("Failed to read example input: %v", err)
	}

	result := part1(lines)
	expected := 3

	if result != expected {
		t.Errorf("part1() = %d, want %d", result, expected)
	}
}
func TestPart2(t *testing.T) {
	lines, err := common.ReadLines("part2example.txt")
	if err != nil {
		t.Fatalf("Failed to read example input: %v", err)
	}

	result := part2(lines)
	expected := 14

	if result != expected {
		t.Errorf("part1() = %d, want %d", result, expected)
	}
}
