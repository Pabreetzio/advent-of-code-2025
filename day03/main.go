package main

import (
	"advent-of-code-2025/common"
	"fmt"
	"log"
	"strconv"
)

func part1(lines []string) int {
	numberOfBatteries := 2
	maximumJoltageSum := 0
	for _, bank := range lines {
		maximumJoltage := getMaximumJoltage(bank, numberOfBatteries)
		maximumJoltageSum += maximumJoltage
	}
	return maximumJoltageSum
}

func getMaximumJoltage(bank string, numberOfBatteries int) int {
	numberOfBatteriesLeft := numberOfBatteries
	bankRemaining := bank
	selectedBank := ""
	for numberOfBatteriesLeft > 0 {
		largestBattery := 0
		largestBattery, bankRemaining = setLargestBattery(bankRemaining, numberOfBatteriesLeft)
		selectedBank += strconv.Itoa(largestBattery)
		numberOfBatteriesLeft--
	}
	fmt.Printf("%v\n", selectedBank)
	maximumJoltage, _ := strconv.Atoi(selectedBank)
	return maximumJoltage
}

func setLargestBattery(bank string, numberOfBatteriesLeft int) (int, string) {
	largestBattery := 0
	largestBatteryIndex := -1
	for i, char := range bank {
		if len(bank)-i < numberOfBatteriesLeft {
			break
		}
		batteryJoltage, _ := strconv.Atoi(string(char))
		if batteryJoltage > largestBattery {
			largestBattery = batteryJoltage
			largestBatteryIndex = i
		}
	}
	bankRemaining := bank[largestBatteryIndex+1:]
	return largestBattery, bankRemaining
}
func part2(lines []string) int {
	numberOfBatteries := 12
	maximumJoltageSum := 0
	for _, bank := range lines {
		maximumJoltage := getMaximumJoltage(bank, numberOfBatteries)
		maximumJoltageSum += maximumJoltage
	}
	return maximumJoltageSum
}

func main() {
	lines, err := common.ReadLines("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	testLines, err := common.ReadLines("test_input.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	fmt.Println("Day 3")
	fmt.Println("-----")

	fmt.Printf("Part 1 test: %d\n", part1(testLines))
	fmt.Printf("Part 2 test: %d\n", part2(testLines))

	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}
