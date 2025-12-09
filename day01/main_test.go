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
	lines, err := common.ReadLines("part1example.txt")
	if err != nil {
		t.Fatalf("Failed to read example input: %v", err)
	}

	result := part2(lines)
	expected := 8

	if result != expected {
		t.Errorf("part1() = %d, want %d", result, expected)
	}
}

func TestRotateLeft(t *testing.T) {
	result := RotateLeft(50, 68)
	expected := 82
	if result != expected {
		t.Errorf("part1() = %d, want %d", result, expected)
	}
}

func TestRotateRight(t *testing.T) {
	result := RotateRight(52, 48)
	expected := 0
	if result != expected {
		t.Errorf("part1() = %d, want %d", result, expected)
	}
}

func TestIsLeft(t *testing.T) {
	result := IsLeft("L48")
	expected := true
	if result != expected {
		t.Errorf("error")
	}
}

func TestNumberClicks(t *testing.T) {
	result := NumberClicks("L123")
	expected := 123
	if result != expected {
		t.Errorf("part1() = %d, want %d", result, expected)
	}
}
func TestRotateRightCounting0test1(t *testing.T) {
	_, result := RotateRightCounting0(0, 100)
	expected := 1
	if result != expected {
		t.Errorf("part1() = %d, want %d", result, expected)
	}
}
func TestRotateRightCounting0test2(t *testing.T) {
	_, result := RotateRightCounting0(0, 99)
	expected := 0
	if result != expected {
		t.Errorf("part1() = %d, want %d", result, expected)
	}
}
func TestRotateRightCounting0test3(t *testing.T) {
	_, result := RotateRightCounting0(0, 1)
	expected := 0
	if result != expected {
		t.Errorf("part1() = %d, want %d", result, expected)
	}
}
func TestRotateRightCounting0test4(t *testing.T) {
	_, result := RotateRightCounting0(1, 100)
	expected := 1
	if result != expected {
		t.Errorf("part1() = %d, want %d", result, expected)
	}
}
func TestRotateRightCountingtest5(t *testing.T) {
	_, result := RotateRightCounting0(99, 2)
	expected := 1
	if result != expected {
		t.Errorf("part1() = %d, want %d", result, expected)
	}
}
func TestRotateRightCountingtest6(t *testing.T) {
	_, result := RotateRightCounting0(99, 1)
	expected := 1
	if result != expected {
		t.Errorf("part1() = %d, want %d", result, expected)
	}
}

func TestRotateLeftCountingtest1(t *testing.T) {
	_, result := RotateLeftCounting0(1, 100)
	expected := 1
	if result != expected {
		t.Errorf("part1() = %d, want %d", result, expected)
	}
}
