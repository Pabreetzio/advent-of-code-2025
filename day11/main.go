package main

import (
	"advent-of-code-2025/common"
	"fmt"
	"log"
)

func part1(lines []string) int {
	// TODO: Implement part 1
	return 0
}

func part2(lines []string) int {
	// TODO: Implement part 2
	return 0
}

func main() {
	lines, err := common.ReadLines("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	fmt.Println("Day 11")
	fmt.Println("-----")
	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}
