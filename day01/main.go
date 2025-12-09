package main

import (
	"advent-of-code-2025/common"
	"fmt"
	"log"
	"strconv"
)

func part1(lines []string) int {
	position := 50
	timesAt0 := 0
	for _, line := range lines {
		clicks := NumberClicks(line)
		if IsLeft(line) {
			position = RotateLeft(position, clicks)
		} else {
			position = RotateRight(position, clicks)
		}
		if position == 0 {
			timesAt0++
		}
	}
	return timesAt0
}

func part2(lines []string) int {
	position := 50
	timesAt0 := 0
	for lineNumber, line := range lines {
		clicks := NumberClicks(line)
		clicksTo0 := 0
		if IsLeft(line) {
			position, clicksTo0 = RotateLeftCounting0(position, clicks)
		} else {
			position, clicksTo0 = RotateRightCounting0(position, clicks)
		}
		timesAt0 += clicksTo0
		fmt.Printf("%d:%s - %d\n", lineNumber, line, clicksTo0)
	}
	return timesAt0
}

func main() {
	lines, err := common.ReadLines("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	fmt.Println("Day 1")
	fmt.Println("-----")
	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}

func RotateLeft(startingPosition int, clicks int) (newPosition int) {
	newPosition = (startingPosition - clicks) % 100
	if newPosition < 0 {
		newPosition = 100 + newPosition
	}
	return newPosition
}

func RotateLeftCounting0(startingPosition int, clicks int) (newPosition int, timesAt0 int) {
	newPosition = (startingPosition - clicks) % 100
	rotations := clicks / 100
	timesAt0 = rotations
	if newPosition < 0 {
		newPosition = 100 + newPosition
	}
	if newPosition > startingPosition && startingPosition != 0 {
		timesAt0++
	}
	if newPosition == 0 && startingPosition != 0 {
		timesAt0++
	}
	return newPosition, timesAt0
}
func IsLeft(line string) bool {
	return line[0] == 'L'
}
func NumberClicks(line string) int {
	clickString := line[1:]
	clickInt, _ := strconv.Atoi(clickString)
	return clickInt
}

func RotateRight(startingPosition int, clicks int) (newPosition int) {
	newPosition = (startingPosition + clicks) % 100
	return newPosition
}
func RotateRightCounting0(startingPosition int, clicks int) (newPosition int, timesAt0 int) {
	newPosition = (startingPosition + clicks) % 100
	rotations := clicks / 100
	timesAt0 = rotations
	if startingPosition > newPosition && newPosition != 0 {
		timesAt0++
	}
	if newPosition == 0 && startingPosition != 0 {
		timesAt0++
	}
	return newPosition, timesAt0
}
