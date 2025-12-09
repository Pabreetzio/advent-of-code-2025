package main

import (
	"advent-of-code-2025/common"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func attempt1part1(lines []string) int {
	cummulativeMaxJoltage := 0
	for lineNumber, line := range lines {
		lineMaxJoltage := getMaxJoltage(line, lineNumber)
		cummulativeMaxJoltage += lineMaxJoltage
	}
	return cummulativeMaxJoltage
}

func getMaxJoltage(line string, lineNumber int) int {
	maxJoltage := 0
	for i := 0; i < len(line)-1; i++ {
		batteryOneJoltage, _ := strconv.Atoi(string(line[i]))
		for j := i + 1; j < len(line); j++ {
			batteryTwoJoltage, _ := strconv.Atoi(string(line[j]))
			joltage := batteryOneJoltage*10 + batteryTwoJoltage
			if joltage > maxJoltage {
				maxJoltage = joltage
			}
		}
	}
	fmt.Printf("Line %d max joltage: %d\n", lineNumber+1, maxJoltage)
	return maxJoltage
}
func attempt1part2(lines []string) int {

	cummulativeMaxJoltage := 0
	for lineNumber, line := range lines {
		lineMaxJoltage := getMaxJoltagePart2(line, lineNumber)
		cummulativeMaxJoltage += lineMaxJoltage
	}
	return cummulativeMaxJoltage
}
func getMaxJoltagePart2(line string, lineNumber int) int {
	maxJoltage := 0

	reducedLine := line
	//while loop  while line length is greater than 12

	for lowestJoltage := 1; lowestJoltage < len(reducedLine)-11; lowestJoltage++ {
		//if the lenght of the reduced line is greater than 12 minus the count of batteries at the lowest joltage then delete all batteries with that joltage from the line
		batteriesAtLowestJoltage := 0
		for i := 0; i < len(reducedLine); i++ {
			batteryJoltage, _ := strconv.Atoi(string(reducedLine[i]))
			if batteryJoltage == lowestJoltage {
				batteriesAtLowestJoltage++
			}
		}
		if len(reducedLine)-batteriesAtLowestJoltage < 12 {
			//only delete the rightmost batteries with the lowest joltage until the line length is 12
			//len(reducedLine) - NumberOfBatteriesToDelete = 12 so NumberOfBatteriesToDelete = len(reducedLine) - 12
			numberOfBatteriesToDelete := len(reducedLine) - 12
			for n := 0; n < numberOfBatteriesToDelete; n++ {
				// Find the last occurrence of the lowest joltage and remove it
				lastIndex := strings.LastIndex(reducedLine, strconv.Itoa(lowestJoltage))
				if lastIndex != -1 {
					reducedLine = reducedLine[:lastIndex] + reducedLine[lastIndex+1:]
				}
			}
		} else {
			//remove the character for lowest joltage from the line
			reducedLine = strings.ReplaceAll(reducedLine, strconv.Itoa(lowestJoltage), "")
		}
	}
	for i := 0; i < 12; i++ {
		batteryJoltage, _ := strconv.Atoi(string(reducedLine[i]))
		maxJoltage += batteryJoltage * int(math.Pow(10, float64(12-i-1)))
	}
	fmt.Printf("Line %d max joltage: %d\n", lineNumber+1, maxJoltage)
	return maxJoltage
}

func attempt1main() {
	lines, err := common.ReadLines("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	fmt.Println("Day 1")
	fmt.Println("-----")
	fmt.Printf("Part 1: %d\n", attempt1part1(lines))
	fmt.Printf("Part 2: %d\n", attempt1part2(lines))
}
