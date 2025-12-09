package main

import (
	"advent-of-code-2025/common"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func part1(lines []string) int {
	redTiles := getTilesFromInput(lines)
	perimeterSegments := getSegmentsFromRedTiles(redTiles)
	largestRectangleArea := getLargestRectangleArea(redTiles, perimeterSegments, false)
	return largestRectangleArea
}

func part2(lines []string) int {
	redTiles := getTilesFromInput(lines)
	perimeterSegments := getSegmentsFromRedTiles(redTiles)
	largestRectangleArea := getLargestRectangleArea(redTiles, perimeterSegments, true)
	return largestRectangleArea
}

func getTilesFromInput(segments []string) []Tile {
	tiles := []Tile{}
	for _, segment := range segments {
		coords := strings.Split(segment, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		tile := Tile{
			x: x,
			y: y,
		}
		tiles = append(tiles, tile)
	}
	return tiles
}

func getSegmentsFromRedTiles(tiles []Tile) []Segment {
	segments := []Segment{}
	for i := 0; i < len(tiles)-1; i++ {
		startTile := tiles[i]
		endTile := tiles[i+1]
		segment := Segment{
			start: startTile,
			end:   endTile,
		}
		segments = append(segments, segment)
	}
	//add line from last point to first point
	segments = append(segments, Segment{
		start: tiles[len(tiles)-1],
		end:   tiles[0],
	})
	return segments
}

func getLargestRectangleArea(tiles []Tile, segments []Segment, checkIfRedecoratable bool) int {
	largestRectangleArea := 0
	for _, corner1 := range tiles {
		for _, corner2 := range tiles {
			rectangleArea := getRectangleArea(corner1, corner2)
			if rectangleArea > largestRectangleArea {
				if !checkIfRedecoratable || isRectangleInsideRedGreenRegion(corner1, corner2, segments) {
					largestRectangleArea = rectangleArea
				}
			}
		}
	}
	return largestRectangleArea
}

func getRectangleArea(corner1 Tile, corner2 Tile) int {
	width := int(math.Abs(float64(corner1.x-corner2.x))) + 1
	height := int(math.Abs(float64(corner1.y-corner2.y))) + 1
	return width * height
}

func isRectangleInsideRedGreenRegion(corner1 Tile, corner2 Tile, segments []Segment) bool {
	rectangleDiagonal := Segment{
		start: corner1,
		end:   corner2,
	}
	for _, segment := range segments {
		if doSegmentsIntersect(rectangleDiagonal, segment) {
			return false
		}
	}
	return true
}

func doSegmentsIntersect(segment1 Segment, segment2 Segment) bool {
	return maxX(segment1) > minX(segment2) && minX(segment1) < maxX(segment2) &&
		maxY(segment1) > minY(segment2) && minY(segment1) < maxY(segment2)
}

type Tile struct {
	x int
	y int
}

type Segment struct {
	start Tile
	end   Tile
}

func maxX(line Segment) int {
	if line.start.x > line.end.x {
		return line.start.x
	}
	return line.end.x
}

func minX(line Segment) int {
	if line.start.x < line.end.x {
		return line.start.x
	}
	return line.end.x
}

func maxY(line Segment) int {
	if line.start.y > line.end.y {
		return line.start.y
	}
	return line.end.y
}

func minY(line Segment) int {
	if line.start.y < line.end.y {
		return line.start.y
	}
	return line.end.y
}

func main() {

	lines, err := common.ReadLines("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	exampleLines, err := common.ReadLines("example.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	fmt.Println("Day 9")
	fmt.Println("-----")
	fmt.Printf("Example Part 1: %d\n", part1(exampleLines))
	fmt.Printf("Example Part 2: %d\n", part2(exampleLines))
	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}

//part 1 answer = 4776100539
//part 2 =1476550548
