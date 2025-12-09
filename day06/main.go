package main

import (
	"advent-of-code-2025/common"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func part1(lines []string) int64 {
	mathProblems := parseInput(lines)
	sumOfSolutions := int64(0)
	for _, problem := range mathProblems {
		solution := solveProblem(problem)
		sumOfSolutions += solution
	}
	return sumOfSolutions
}

func parseInput(lines []string) (mathProblems []MathProblem) {
	operationLineNumber := 4
	for lineNumber, line := range lines {
		fields := strings.Fields(line)
		if lineNumber == 0 {
			mathProblems = make([]MathProblem, len(fields))
			for problemNumber, number := range fields {
				numbers := make([]string, 4)
				numbers[lineNumber] = number
				newProblem := MathProblem{
					numbers: numbers,
				}
				mathProblems[problemNumber] = newProblem
			}
		}
		if lineNumber < operationLineNumber {
			for problemNumber, number := range fields {
				mathProblems[problemNumber].numbers[lineNumber] = number
			}
		}
		if lineNumber == operationLineNumber {
			for problemNumber, operatorString := range fields {
				mathProblems[problemNumber].operation = operatorString[0]
			}
		}
	}
	return mathProblems
}

func solveProblem(mathProblem MathProblem) (solution int64) {
	solution = 0
	if mathProblem.operation == "+"[0] {
		solution = 0
		for _, number := range mathProblem.numbers {
			numberInt, _ := strconv.Atoi(number)
			solution = solution + int64(numberInt)
		}
	}
	if mathProblem.operation == "*"[0] {
		solution = 1
		for _, number := range mathProblem.numbers {
			numberInt, _ := strconv.Atoi(number)
			solution = solution * int64(numberInt)
		}
	}
	return solution
}

func part2(lines []string) int64 {
	mathProblems := parseInputPart2(lines)
	sumOfSolutions := int64(0)
	for _, problem := range mathProblems {
		solution := solveProblem(problem)
		sumOfSolutions += solution
	}
	return sumOfSolutions
}

func parseInputPart2(lines []string) (mathProblems []MathProblem) {
	maxLineLength := 0
	for _, line := range lines {
		if len(line) > maxLineLength {
			maxLineLength = len(line)
		}
	}
	mathProblems = make([]MathProblem, 0)
	mathProblem := MathProblem{}
	for column := 0; column < maxLineLength; column++ {
		blankColumn := true
		number := ""
		if lines[0][column] != ' ' {
			number += string(lines[0][column])
			blankColumn = false
		}
		if lines[1][column] != ' ' {
			number += string(lines[1][column])
			blankColumn = false
		}
		if lines[2][column] != ' ' {
			number += string(lines[2][column])
			blankColumn = false
		}
		if lines[3][column] != ' ' {
			number += string(lines[3][column])
			blankColumn = false
		}

		if lines[4][column] != ' ' {
			mathProblem.operation += lines[4][column]
			blankColumn = false
		}
		if blankColumn {
			mathProblems = append(mathProblems, mathProblem)
			mathProblem = MathProblem{}
			mathProblem.numbers = []string{}
		} else {
			mathProblem.numbers = append(mathProblem.numbers, number)
		}
	}
	mathProblems = append(mathProblems, mathProblem)

	return mathProblems
}

// func solveProblemPart2(mathProblem MathProblem) (solution int64) {
// 	return 0
// }

func main() {
	lines, err := common.ReadLines("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	fmt.Println("Day 6")
	fmt.Println("-----")
	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}

type MathProblem struct {
	numbers   []string
	operation byte
}
