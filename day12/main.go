package main

import (
	"advent-of-code-2025/common"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func part1(lines []string) int {
	// TODO: Implement part 1
	presentShapes, regions := parseInput(lines)
	howManyOfTheRegionsCanFitPresents := 0
	for _, region := range regions {
		if regionCanFitPresents(region, presentShapes) {
			howManyOfTheRegionsCanFitPresents++
		}
	}
	return howManyOfTheRegionsCanFitPresents
}

func regionCanFitPresents(region Region, presentShapes []PresentShape) bool {
	if region.SanityCheck == "INSANELY SMALL" {
		return false
	} else if region.SanityCheck == "SPACIOUS" {
		return true
	} else if region.SanityCheck == "MAYBE" {
		//throw error saying I don't f* know how to figure out how to make it fit but it might
		log.Fatalf("Region size %dx%d with required presents %v is marked MAYBE, and I don't know how to handle that yet", region.SizeX, region.SizeY, region.RequiredPresents)
	}
	return false
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

	fmt.Println("Day 12")
	fmt.Println("-----")
	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}

func parseInput(lines []string) ([]PresentShape, []Region) {
	sections := [][]string{}
	section := []string{}
	for _, line := range lines {
		if line == "" {
			sections = append(sections, section)
			section = []string{}
		} else {
			section = append(section, line)
		}
	}
	sections = append(sections, section)
	//all but last section contain present shapes
	presentShapes := parsePresentShapes(sections[:len(sections)-1])
	// the last section is a list of regions
	regions := parseRegions(sections[len(sections)-1])
	doSanityCheck(regions, presentShapes)
	for _, region := range regions {
		fmt.Println(region)
	}
	return presentShapes, regions
}

// if packed perfecly and there still would be no room then set region SanityCheck to "INSANELY SMALL"
// if there is enough room for each present without any overlap then set region SanityCheck to "SPACIOUS"
// else set region SanityCheck to "MAYBE"
func doSanityCheck(regions []Region, presentShapes []PresentShape) {

	presentAreas := []int{}
	for _, presentShape := range presentShapes {
		presentArea := 0
		for _, row := range presentShape {
			for _, cell := range row {
				if cell {
					presentArea++
				}
			}
		}
		presentAreas = append(presentAreas, presentArea)
	}
	for i, region := range regions {
		//each present will at the worst case still fit into a 3x3 area
		spaciousPresentSize := 0
		for _, requiredPresent := range region.RequiredPresents {
			spaciousPresentSize += 9 * requiredPresent
		}
		totalPresentArea := 0
		for requirePresentIndex, requiredPresent := range region.RequiredPresents {
			totalPresentArea += presentAreas[requirePresentIndex] * requiredPresent
		}
		regionArea := region.SizeX * region.SizeY
		if totalPresentArea > regionArea {
			regions[i].SanityCheck = "INSANELY SMALL"
		} else if regionArea >= spaciousPresentSize {
			regions[i].SanityCheck = "SPACIOUS"
		} else {
			regions[i].SanityCheck = "MAYBE"
		}
	}
}

func parsePresentShapes(sections [][]string) []PresentShape {
	presentShapes := []PresentShape{}
	for _, section := range sections {
		shape := PresentShape{}
		for lineNumber, line := range section {
			if lineNumber == 0 {
				//expect format to be "presentIndex:" and we can ignore it
				continue
			}
			row := []bool{}
			for _, char := range line {
				if char == '#' {
					row = append(row, true)
				} else {
					row = append(row, false)
				}
			}
			shape = append(shape, row)
		}
		presentShapes = append(presentShapes, shape)
	}
	return presentShapes
}

func parsePresentShape(lines []string) PresentShape {
	shape := PresentShape{}
	for _, line := range lines {
		row := []bool{}
		for _, char := range line {
			if char == '#' {
				row = append(row, true)
			} else {
				row = append(row, false)
			}
		}
		shape = append(shape, row)
	}
	return shape
}

func parseRegions(lines []string) []Region {
	regions := []Region{}
	for _, line := range lines {
		region := parseRegion(line)
		regions = append(regions, region)
	}
	return regions
}

func parseRegion(line string) Region {
	region := Region{}
	sidesOfLine := strings.Split(line, ": ")
	dimensions := strings.Split(sidesOfLine[0], "x")
	region.SizeX, _ = strconv.Atoi(dimensions[0])
	region.SizeY, _ = strconv.Atoi(dimensions[1])
	requiredPresentsStrings := strings.Split(sidesOfLine[1], " ")
	for _, presentCountString := range requiredPresentsStrings {
		presentCount, _ := strconv.Atoi(presentCountString)
		region.RequiredPresents = append(region.RequiredPresents, presentCount)
	}
	return region
}

type PresentShape [][]bool
type Region struct {
	SizeX            int
	SizeY            int
	RequiredPresents []int
	SanityCheck      string
}

func (region Region) String() string {
	requiredPresentsString := ""
	for i, count := range region.RequiredPresents {
		if i > 0 {
			requiredPresentsString += " "
		}
		requiredPresentsString += strconv.Itoa(count)
	}
	return strconv.Itoa(region.SizeX) + "x" + strconv.Itoa(region.SizeY) + ": " + requiredPresentsString + " - " + region.SanityCheck
}
