package main

import (
	"advent-of-code-2025/common"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func part1(lines []string, connections int) int {
	points := []Point3{}
	for _, line := range lines {
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])
		point := Point3{
			X: x,
			Y: y,
			Z: z,
		}
		points = append(points, point)
	}
	pairs := []Pair{}
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			pair := Pair{}
			pair.Points = []Point3{
				points[i],
				points[j],
			}
			pair.Distance = getDistance(pair.Points[0], pair.Points[1])
			pairs = append(pairs, pair)
		}
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Distance < pairs[j].Distance
	})

	circuits := []Circuit{}
	for i := 0; i < connections; i++ {
		point1CircuitNumber := -1
		point2CircuitNumber := -1
		for circuitNumber, circuit := range circuits {
			if circuit.JunctionBoxes[pairs[i].Points[0]] {
				point1CircuitNumber = circuitNumber
			}
			if circuit.JunctionBoxes[pairs[i].Points[1]] {
				point2CircuitNumber = circuitNumber
			}
		}
		if point1CircuitNumber != -1 && point1CircuitNumber == point2CircuitNumber {
			continue
		} else if point1CircuitNumber == -1 && point2CircuitNumber == -1 {
			newCircuit := Circuit{}
			newCircuit.JunctionBoxes = map[Point3]bool{}
			newCircuit.JunctionBoxes[pairs[i].Points[0]] = true
			newCircuit.JunctionBoxes[pairs[i].Points[1]] = true
			circuits = append(circuits, newCircuit)
		} else if point1CircuitNumber != -1 && point2CircuitNumber == -1 {
			circuits[point1CircuitNumber].JunctionBoxes[pairs[i].Points[1]] = true

		} else if point1CircuitNumber == -1 && point2CircuitNumber != -1 {
			circuits[point2CircuitNumber].JunctionBoxes[pairs[i].Points[0]] = true
		} else {
			newCircuit := Circuit{}
			newCircuit.JunctionBoxes = map[Point3]bool{}
			for junctionBox, _ := range circuits[point1CircuitNumber].JunctionBoxes {
				newCircuit.JunctionBoxes[junctionBox] = true
			}
			for junctionBox, _ := range circuits[point2CircuitNumber].JunctionBoxes {
				newCircuit.JunctionBoxes[junctionBox] = true
			}
			newCircuits := []Circuit{}
			for circuitNumber, circuit := range circuits {
				if circuitNumber != point1CircuitNumber && circuitNumber != point2CircuitNumber {
					newCircuits = append(newCircuits, circuit)
				}
			}
			newCircuits = append(newCircuits, newCircuit)
			circuits = newCircuits
		}
	}
	return productOfThreeLargestCircuits(circuits)
}

func getCircuit(circuits []Circuit, point Point3) Circuit {
	for _, circuit := range circuits {
		if circuit.JunctionBoxes[point] {
			return circuit
		}
	}
	newCircuit := Circuit{
		JunctionBoxes: map[Point3]bool{},
	}
	newCircuit.JunctionBoxes[point] = true
	return newCircuit
}

func getDistance(point1 Point3, point2 Point3) (distance int64) {
	x2 := int64(point1.X-point2.X) * int64(point1.X-point2.X)
	y2 := int64(point1.Y-point2.Y) * int64(point1.Y-point2.Y)
	z2 := int64(point1.Z-point2.Z) * int64(point1.Z-point2.Z)
	distance = x2 + y2 + z2
	return distance
}

func productOfThreeLargestCircuits(circuits []Circuit) (product int) {
	product = 1
	for i := 0; i < 3; i++ {
		newCircuits, largestLength := popLargestCircuit(circuits)
		product = product * largestLength
		circuits = newCircuits
	}

	return product
}
func popLargestCircuit(circuits []Circuit) (newCircuits []Circuit, largestLength int) {
	largestLength = 0
	var largestCircuit Circuit
	newCircuits = []Circuit{}
	for _, circuit := range circuits {
		if len(circuit.JunctionBoxes) > largestLength {
			//add the previously largest to new circuits
			if largestLength > 0 {
				newCircuits = append(newCircuits, largestCircuit)
			}
			//update largest with new largest
			largestCircuit = circuit
			//update length to new largest length
			largestLength = len(circuit.JunctionBoxes)
		} else {
			newCircuits = append(newCircuits, circuit)
		}
	}
	return newCircuits, largestLength
}

func part2(lines []string, numberOfPoints int) int64 {
	points := []Point3{}
	for _, line := range lines {
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])
		point := Point3{
			X: x,
			Y: y,
			Z: z,
		}
		points = append(points, point)
	}
	pairs := []Pair{}
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			pair := Pair{}
			pair.Points = []Point3{
				points[i],
				points[j],
			}
			pair.Distance = getDistance(pair.Points[0], pair.Points[1])
			pairs = append(pairs, pair)
		}
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Distance < pairs[j].Distance
	})

	circuits := []Circuit{}
	isNotLastPair := true
	lastPair := Pair{}
	i := 0
	for isNotLastPair {
		point1CircuitNumber := -1
		point2CircuitNumber := -1
		for circuitNumber, circuit := range circuits {
			if circuit.JunctionBoxes[pairs[i].Points[0]] {
				point1CircuitNumber = circuitNumber
			}
			if circuit.JunctionBoxes[pairs[i].Points[1]] {
				point2CircuitNumber = circuitNumber
			}
		}
		if point1CircuitNumber != -1 && point1CircuitNumber == point2CircuitNumber {
			i++
			continue
		} else if point1CircuitNumber == -1 && point2CircuitNumber == -1 {
			newCircuit := Circuit{}
			newCircuit.JunctionBoxes = map[Point3]bool{}
			newCircuit.JunctionBoxes[pairs[i].Points[0]] = true
			newCircuit.JunctionBoxes[pairs[i].Points[1]] = true
			circuits = append(circuits, newCircuit)
		} else if point1CircuitNumber != -1 && point2CircuitNumber == -1 {
			circuits[point1CircuitNumber].JunctionBoxes[pairs[i].Points[1]] = true

		} else if point1CircuitNumber == -1 && point2CircuitNumber != -1 {
			circuits[point2CircuitNumber].JunctionBoxes[pairs[i].Points[0]] = true
		} else {
			newCircuit := Circuit{}
			newCircuit.JunctionBoxes = map[Point3]bool{}
			for junctionBox, _ := range circuits[point1CircuitNumber].JunctionBoxes {
				newCircuit.JunctionBoxes[junctionBox] = true
			}
			for junctionBox, _ := range circuits[point2CircuitNumber].JunctionBoxes {
				newCircuit.JunctionBoxes[junctionBox] = true
			}
			newCircuits := []Circuit{}
			for circuitNumber, circuit := range circuits {
				if circuitNumber != point1CircuitNumber && circuitNumber != point2CircuitNumber {
					newCircuits = append(newCircuits, circuit)
				}
			}
			newCircuits = append(newCircuits, newCircuit)
			circuits = newCircuits
		}
		if len(circuits) > 0 && len(circuits[0].JunctionBoxes) == numberOfPoints {
			isNotLastPair = false
			lastPair = pairs[i]
		}
		i++
	}
	return productOfLastPairOnXAxis(lastPair)
}

func productOfLastPairOnXAxis(pair Pair) int64 {
	return int64(pair.Points[0].X) * int64(pair.Points[1].X)
}
func main() {
	lines, err := common.ReadLines("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	example, err := common.ReadLines("part1example.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	fmt.Println("Day 8")
	fmt.Println("-----")
	fmt.Printf("Part 1 Example: %d\n", part1(example, 10))
	fmt.Printf("Part 1: %d\n", part1(lines, 1000))
	fmt.Printf("Part 2 Example: %d\n", part2(example, 20))
	fmt.Printf("Part 2: %d\n", part2(lines, 1000))
}

type Circuit struct {
	JunctionBoxes map[Point3]bool
}

type Point3 struct {
	X int
	Y int
	Z int
}

type Pair struct {
	Points   []Point3
	Distance int64
}
