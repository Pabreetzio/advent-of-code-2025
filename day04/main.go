package main

import (
	"advent-of-code-2025/common"
	"fmt"
	"log"
)

func part1(lines []string) int {
	count := 0

	// Iterate through each position in the grid
	for row := 0; row < len(lines); row++ {
		for col := 0; col < len(lines[row]); col++ {
			// Check if current position is an '@'
			if lines[row][col] == '@' {
				neighbors := countNeighbors(lines, row, col)
				if neighbors < 4 {
					count++
				}
			}
		}
	}

	return count
}

// countNeighbors counts the number of '@' symbols in the 8 surrounding positions
func countNeighbors(lines []string, row, col int) int {
	count := 0

	// Define all 8 directions: up, down, left, right, and 4 diagonals
	directions := []struct{ dr, dc int }{
		{-1, -1}, {-1, 0}, {-1, 1}, // top-left, top, top-right
		{0, -1}, {0, 1}, // left, right
		{1, -1}, {1, 0}, {1, 1}, // bottom-left, bottom, bottom-right
	}

	// Check each direction
	for _, dir := range directions {
		newRow := row + dir.dr
		newCol := col + dir.dc

		// Check if the position is within bounds
		if newRow >= 0 && newRow < len(lines) &&
			newCol >= 0 && newCol < len(lines[newRow]) {
			if lines[newRow][newCol] == '@' {
				count++
			}
		}
	}

	return count
}

func part2(lines []string) int {
	// Create a mutable copy of the grid
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	totalReplaced := 0

	// Keep iterating until no more '@' can be replaced
	for {
		replacedThisRound := 0
		toReplace := []struct{ row, col int }{}

		// Find all '@' symbols with fewer than 4 neighbors
		for row := 0; row < len(grid); row++ {
			for col := 0; col < len(grid[row]); col++ {
				if grid[row][col] == '@' {
					neighbors := countNeighborsGrid(grid, row, col)
					if neighbors < 4 {
						toReplace = append(toReplace, struct{ row, col int }{row, col})
					}
				}
			}
		}

		// If no '@' to replace, we're done
		if len(toReplace) == 0 {
			break
		}

		// Replace all marked '@' with '.'
		for _, pos := range toReplace {
			grid[pos.row][pos.col] = '.'
			replacedThisRound++
		}

		totalReplaced += replacedThisRound
		fmt.Printf("Replaced %d '@' symbols this round (total: %d)\n", replacedThisRound, totalReplaced)
	}

	return totalReplaced
}

// countNeighborsGrid counts the number of '@' symbols in the 8 surrounding positions (grid version)
func countNeighborsGrid(grid [][]rune, row, col int) int {
	count := 0

	// Define all 8 directions: up, down, left, right, and 4 diagonals
	directions := []struct{ dr, dc int }{
		{-1, -1}, {-1, 0}, {-1, 1}, // top-left, top, top-right
		{0, -1}, {0, 1}, // left, right
		{1, -1}, {1, 0}, {1, 1}, // bottom-left, bottom, bottom-right
	}

	// Check each direction
	for _, dir := range directions {
		newRow := row + dir.dr
		newCol := col + dir.dc

		// Check if the position is within bounds
		if newRow >= 0 && newRow < len(grid) &&
			newCol >= 0 && newCol < len(grid[newRow]) {
			if grid[newRow][newCol] == '@' {
				count++
			}
		}
	}

	return count
}

func main() {
	lines, err := common.ReadLines("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	fmt.Println("Day 4")
	fmt.Println("-----")
	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}
