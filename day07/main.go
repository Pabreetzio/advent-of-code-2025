package main

import (
	"advent-of-code-2025/common"
	"fmt"
	"log"
	"strings"
)

func part1(lines []string) int {
	numberOfTimesSplit := 0
	width := len(lines[0])
	startingPosition := strings.Index(lines[0], "S")
	current := make([]bool, width)
	current[startingPosition] = true
	for _, line := range lines {
		next := make([]bool, width)
		for positionNumber, splitterIndicator := range line {
			if current[positionNumber] && splitterIndicator == '^' {
				numberOfTimesSplit++
				if positionNumber > 0 {
					next[positionNumber-1] = true
				}
				if positionNumber < width-2 {
					next[positionNumber+1] = true
				}

			} else if current[positionNumber] {
				next[positionNumber] = true
			}
		}
		current = next
	}
	return numberOfTimesSplit
}

func part2(lines []string) int {
	numberOfTimelines := 0
	width := len(lines[0])
	startingPosition := strings.Index(lines[0], "S")
	current := make([]int, width)
	current[startingPosition] = 1
	for _, line := range lines {
		next := make([]int, width)
		for positionNumber, splitterIndicator := range line {
			if current[positionNumber] > 0 && splitterIndicator == '^' {
				if positionNumber > 0 {
					next[positionNumber-1] = next[positionNumber-1] + current[positionNumber]
				}
				if positionNumber < width-1 {
					next[positionNumber+1] = next[positionNumber+1] + current[positionNumber]
				}

			} else if current[positionNumber] > 0 {
				next[positionNumber] = next[positionNumber] + current[positionNumber]
			}
		}
		current = next
	}
	for _, tachyonsTimelines := range current {
		numberOfTimelines += tachyonsTimelines
	}
	return numberOfTimelines
}

func main() {

	lines, err := common.ReadLines("part1example.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	fmt.Println("Day 7")
	fmt.Println("-----")

	fmt.Printf("Part 1 Example: %d\n", part1(lines))
	fmt.Printf("Part 2 Example: %d\n", part2(lines))
	lines, err = common.ReadLines("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}
