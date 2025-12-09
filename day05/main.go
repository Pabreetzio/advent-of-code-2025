package main

import (
	"advent-of-code-2025/common"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func part1(lines []string) int {
	freshIngredients := 0
	freshIngredientIdRanges := []FreshRange{}
	availableIngredientIds := []int{}
	freshIngredientIdRanges, availableIngredientIds = parseInput(lines)
	freshIngredients = countFreshIngredients(freshIngredientIdRanges, availableIngredientIds)

	return freshIngredients
}

func countFreshIngredients(freshIngredientIdRanges []FreshRange, availableIngredientIds []int) int {
	freshIngredients := 0
	for _, availableIngredientId := range availableIngredientIds {
		isFresh := false
		for _, freshIngredientIdRange := range freshIngredientIdRanges {
			if availableIngredientId >= freshIngredientIdRange.start && availableIngredientId <= freshIngredientIdRange.end {
				isFresh = true
				break
			}
		}
		if isFresh {
			freshIngredients++
		}
	}
	return freshIngredients
}

func parseInput(lines []string) ([]FreshRange, []int) {
	freshIngredientIdRanges := []FreshRange{}
	availableIngredientIds := []int{}
	isFirstPartOfInput := true
	for _, line := range lines {
		if line == "" {
			isFirstPartOfInput = false
			continue
		}
		if isFirstPartOfInput {
			freshRange := strings.Split(line, "-")
			start, _ := strconv.Atoi(freshRange[0])
			end, _ := strconv.Atoi(freshRange[1])
			freshIngredientIdRanges = append(freshIngredientIdRanges, FreshRange{start: start, end: end})
		} else {
			availableIngredientId, _ := strconv.Atoi(line)
			availableIngredientIds = append(availableIngredientIds, availableIngredientId)
		}
	}
	return freshIngredientIdRanges, availableIngredientIds
}

func part2(lines []string) int {
	freshIngredientIdRanges := []FreshRange{}
	freshIngredientIdRanges, _ = parseInput(lines)
	collapsedRanges := collapseRanges(freshIngredientIdRanges)
	numberCollapsed := len(freshIngredientIdRanges) - len(collapsedRanges)
	for numberCollapsed > 0 {
		freshIngredientIdRanges = collapsedRanges
		collapsedRanges = collapseRanges(freshIngredientIdRanges)
		numberCollapsed = len(freshIngredientIdRanges) - len(collapsedRanges)
	}
	countOfAllAvailableFreshIds := addRangeLengths(collapsedRanges)

	return countOfAllAvailableFreshIds
}

func collapseRanges(uncollapsedRanges []FreshRange) []FreshRange {
	collapsedRanges := []FreshRange{}
	for _, uncollapsedRange := range uncollapsedRanges {
		isStandalone := true
		for collapsedRangeIndex, collapsedRange := range collapsedRanges {
			//if ranges don't overlap then the start and end of range will both be less than range2 start or both be greater than range2 end
			if (uncollapsedRange.start < collapsedRange.start && uncollapsedRange.end < collapsedRange.start) || (uncollapsedRange.start > collapsedRange.end && uncollapsedRange.end > collapsedRange.end) {
				//do nothing
			} else {
				newRange := collapseOverlappingRange(uncollapsedRange, collapsedRange)
				collapsedRanges[collapsedRangeIndex] = newRange
				isStandalone = false
				break
			}
		}
		if isStandalone {
			collapsedRanges = append(collapsedRanges, uncollapsedRange)
		}
	}
	return collapsedRanges
}

func collapseOverlappingRange(range1 FreshRange, range2 FreshRange) (newRange FreshRange) {
	if range1.start < range2.start {
		newRange.start = range1.start
	} else {
		newRange.start = range2.start
	}

	if range1.end > range2.end {
		newRange.end = range1.end
	} else {
		newRange.end = range2.end
	}
	return newRange
}

func addRangeLengths(ranges []FreshRange) int {
	cummulativeRangeLength := 0
	for _, freshRange := range ranges {
		cummulativeRangeLength += rangeLength(freshRange)
	}
	return cummulativeRangeLength
}

func rangeLength(r FreshRange) int {
	return r.end - r.start + 1
}

func main() {
	lines, err := common.ReadLines("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	fmt.Println("Day 5")
	fmt.Println("-----")
	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}

type FreshRange struct {
	start int
	end   int
}
