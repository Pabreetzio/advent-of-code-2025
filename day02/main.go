package main

import (
	"advent-of-code-2025/common"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func solve(lines []string, isPart1 bool) int {
	// TODO: split lines by comma and name those strings as "ranges" a slice of strings
	rangeStrings := strings.Split(lines[1], ",")
	//create a range min and range max by splitting each range by "-"
	ranges := make([]Range, 0)
	for _, rangeStr := range rangeStrings {
		parts := strings.Split(rangeStr, "-")
		if len(parts) != 2 {
			continue
		}
		min, err1 := strconv.Atoi(parts[0])
		max, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			continue
		}
		ranges = append(ranges, Range{Min: min, Max: max})
	}
	// for each range, call sumEqualPartsInRange and keep a running total
	totalSum := 0
	for _, r := range ranges {
		totalSum += sumEqualPartsInRange(r, isPart1)
	}

	// return the total sum as a string
	return totalSum
}

// create a function that takes a Range object and iterates from the start of the range to the end of the range checking each integer to see if its equal parts, and keeps a running sum of every integer that meets those conditions
func sumEqualPartsInRange(r Range, isPart1 bool) int {
	sum := 0
	for i := r.Min; i <= r.Max; i++ {
		if isSplitEqualParts(i, isPart1) {
			sum += i
		}
	}
	return sum
}

// create a function that checks if an integer can be split and n equal parts and each part is the same sequence of digits
func isSplitEqualParts(num int, isPart1 bool) bool {
	if isPart1 {

		//the following only works if n is 2
		str := strconv.Itoa(num)
		if len(str)%2 != 0 {
			return false
		}
		mid := len(str) / 2
		return str[:mid] == str[mid:]
	} else {

		str := strconv.Itoa(num)
		length := len(str)
		if length == 1 {
			return false
		}
		for n := 1; n <= length/2; n++ {
			if length%n != 0 {
				continue
			}
			if n == 1 {
				partLength := 1
				partOne := str[:partLength]
				match := true
				for i := 1; i < length; i++ {
					partTwo := str[i*partLength : (i+1)*partLength]
					if partTwo != partOne {
						match = false
						break
					}
				}
				if match {
					fmt.Printf("%d\n", num)
					return true
				}
				continue
			} else {
				partLength := length / n
				partOne := str[:partLength]
				match := true
				for i := 1; i < n; i++ {
					partTwo := str[i*partLength : (i+1)*partLength]
					if partTwo != partOne {
						match = false
						break
					}
				}
				if match {
					fmt.Printf("%d\n", num)
					return true
				}
				continue
			}
		}
		return false
	}

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

	fmt.Println("Day 2")
	fmt.Println("-----")
	fmt.Printf("Part 1: %d\n", solve(lines, true))
	fmt.Printf("Part 2: %d\n", solve(lines, false))
}

type Range struct {
	Min int
	Max int
}
