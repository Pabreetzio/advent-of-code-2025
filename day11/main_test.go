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
	expected := 5

	if result != expected {
		t.Errorf("part1() = %d, want %d", result, expected)
	}
}

func TestPartForReal(t *testing.T) {
	lines, err := common.ReadLines("input.txt")
	if err != nil {
		t.Fatalf("Failed to read example input: %v", err)
	}

	result := part1(lines)
	expected := 552

	if result != expected {
		t.Errorf("part1() = %d, want %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	lines, err := common.ReadLines("test_input_part2.txt")
	if err != nil {
		t.Fatalf("Failed to read example input: %v", err)
	}

	result := part2(lines)
	expected := 2

	if result != expected {
		t.Errorf("part2() = %d, want %d", result, expected)
	}
}

func TestFromDac(t *testing.T) {
	lines, err := common.ReadLines("test_input_part2.txt")
	if err != nil {
		t.Fatalf("Failed to read example input: %v", err)
	}
	devices := parseInput(lines)
	result := getNumberOfOutputsForSource(devices, "dac", "out")
	expected := 2

	if result != expected {
		t.Errorf("part1() = %d, want %d", result, expected)
	}
}

func TestFromDacForReal(t *testing.T) {
	lines, err := common.ReadLines("input.txt")
	if err != nil {
		t.Fatalf("Failed to read example input: %v", err)
	}
	devices := parseInput(lines)
	result := getNumberOfOutputsForSource(devices, "dac", "out")
	expected := 3681

	if result != expected {
		t.Errorf("part1() = %d, want %d", result, expected)
	}
}

func TestFromSvrToDacForReal(t *testing.T) {
	lines, err := common.ReadLines("input.txt")
	if err != nil {
		t.Fatalf("Failed to read example input: %v", err)
	}
	devices := parseInput(lines)
	result := getNumberOfOutputsForSource(devices, "dac", "fft")
	expected := 3681

	if result != expected {
		t.Errorf("part1() = %d, want %d", result, expected)
	}
}

func TestFromFftForReal(t *testing.T) {
	lines, err := common.ReadLines("input.txt")
	if err != nil {
		t.Fatalf("Failed to read example input: %v", err)
	}
	devices := parseInput(lines)
	result := getNumberOfOutputsForSource(devices, "fft", "dac")
	expected := 2

	if result != expected {
		t.Errorf("part1() = %d, want %d", result, expected)
	}
}

// func TestSplitOfOutputs(t *testing.T) {

// 	line := "aaaa: you hhh"
// 	result := strings.Split(strings.TrimSpace(strings.Split(line, ":")[1]), " ")
// 	expected := []string{"you", "hhh"}
// 	//compare two arrays
// 	if
// 	if result != expected {
// 		t.Errorf("part1() = %d, want %d", result, expected)
// 	}
// }
