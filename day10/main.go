package main

import (
	"advent-of-code-2025/common"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func part1(lines []string) int {
	machineSchematics := parseInput(lines)
	totalClicks := 0
	for _, machine := range machineSchematics {
		fewestClicks := findFewestButtonClicksForLightMode(machine)
		fmt.Printf("Fewest Clicks: %d\n", fewestClicks)
		totalClicks += fewestClicks
	}
	return totalClicks
}

func part2(lines []string) int {
	machineSchematics := parseInput(lines)
	totalClicks := 0
	for _, machine := range machineSchematics {
		fewestClicks := findFewestButtonClicksForJoltageMode(machine)
		fmt.Printf("Fewest Clicks: %d\n", fewestClicks)
		totalClicks += fewestClicks
	}
	return totalClicks
}

func findFewestButtonClicksForLightMode(machine MachineSchematic) int {
	initialState := MachineState{IndicatorLights: make([]bool, len(machine.IndicatorLights))}
	currentStates := []MachineState{initialState}
	exploredStates := map[string]bool{
		getIndicatorLightStateKey(initialState): true,
	}

	clicks := 0
	for len(currentStates) > 0 {
		nextStates := []MachineState{}

		for _, currentState := range currentStates {
			for _, button := range machine.Buttons {
				newState := applyButtonToIndicatorLights(currentState, button)
				stateKey := getIndicatorLightStateKey(newState)

				if exploredStates[stateKey] {
					continue
				}

				if indicatorLightsMatch(machine.IndicatorLights, newState.IndicatorLights) {
					return clicks + 1
				}

				nextStates = append(nextStates, newState)
				exploredStates[stateKey] = true
			}
		}

		currentStates = nextStates
		clicks++
	}

	return 0
}

func findFewestButtonClicksForJoltageMode(machine MachineSchematic) int {
	coefficientMatrix := buildCoefficientMatrix(machine)
	joltageRequirements := make([]int, len(machine.JoltageRequirements))
	copy(joltageRequirements, machine.JoltageRequirements)

	convertToRowEchelonForm(&coefficientMatrix, &joltageRequirements)

	return solveForMinimumButtonPresses(coefficientMatrix, joltageRequirements)
}

func buildCoefficientMatrix(machine MachineSchematic) [][]int {
	numRegisters := len(machine.JoltageRequirements)
	numButtons := len(machine.Buttons)

	matrix := make([][]int, numRegisters)
	for registerIndex := 0; registerIndex < numRegisters; registerIndex++ {
		matrix[registerIndex] = make([]int, numButtons)
	}

	for buttonIndex, button := range machine.Buttons {
		for _, registerIndex := range button.Toggle {
			matrix[registerIndex][buttonIndex] = 1
		}
	}

	return matrix
}

func convertToRowEchelonForm(matrix *[][]int, rightHandSide *[]int) {
	numRows := len(*matrix)
	if numRows == 0 {
		return
	}
	numColumns := len((*matrix)[0])

	currentRow := 0
	currentColumn := 0

	for currentRow < numRows && currentColumn < numColumns {
		pivotRow := findPivotRow(*matrix, currentRow, currentColumn)

		if pivotRow == -1 {
			currentColumn++
			continue
		}

		if pivotRow != currentRow {
			swapRows(matrix, rightHandSide, currentRow, pivotRow)
		}

		eliminateBelowPivot(matrix, rightHandSide, currentRow, currentColumn)

		currentRow++
		currentColumn++
	}
}

func findPivotRow(matrix [][]int, startRow int, column int) int {
	for row := startRow; row < len(matrix); row++ {
		if matrix[row][column] != 0 {
			return row
		}
	}
	return -1
}

func swapRows(matrix *[][]int, rightHandSide *[]int, row1 int, row2 int) {
	(*matrix)[row1], (*matrix)[row2] = (*matrix)[row2], (*matrix)[row1]
	(*rightHandSide)[row1], (*rightHandSide)[row2] = (*rightHandSide)[row2], (*rightHandSide)[row1]
}

func eliminateBelowPivot(matrix *[][]int, rightHandSide *[]int, pivotRow int, pivotColumn int) {
	pivotValue := (*matrix)[pivotRow][pivotColumn]
	numRows := len(*matrix)
	numColumns := len((*matrix)[0])

	for row := pivotRow + 1; row < numRows; row++ {
		if (*matrix)[row][pivotColumn] == 0 {
			continue
		}

		currentValue := (*matrix)[row][pivotColumn]

		// To preserve integer arithmetic: row = row * pivotValue - pivotRow * currentValue
		for column := pivotColumn; column < numColumns; column++ {
			(*matrix)[row][column] = (*matrix)[row][column]*pivotValue - (*matrix)[pivotRow][column]*currentValue
		}
		(*rightHandSide)[row] = (*rightHandSide)[row]*pivotValue - (*rightHandSide)[pivotRow]*currentValue
	}
}

func solveForMinimumButtonPresses(matrix [][]int, joltageRequirements []int) int {
	numRows := len(matrix)
	if numRows == 0 {
		return 0
	}
	numColumns := len(matrix[0])

	pivotColumns, freeVariableColumns := identifyPivotAndFreeVariables(matrix)
	upperBound := calculateSearchUpperBound(joltageRequirements, freeVariableColumns)

	minimumClicks := -1
	buttonPresses := make([]int, numColumns)

	searchFreeVariableCombinations(0, freeVariableColumns, upperBound, buttonPresses,
		matrix, joltageRequirements, pivotColumns, &minimumClicks)

	if minimumClicks == -1 {
		return 0
	}
	return minimumClicks
}

func identifyPivotAndFreeVariables(matrix [][]int) ([]int, []int) {
	numRows := len(matrix)
	numColumns := len(matrix[0])

	pivotColumns := make([]int, numRows)
	isPivot := make([]bool, numColumns)

	for row := 0; row < numRows; row++ {
		pivotColumns[row] = -1
		for column := 0; column < numColumns; column++ {
			if matrix[row][column] != 0 {
				pivotColumns[row] = column
				isPivot[column] = true
				break
			}
		}
	}

	freeVariableColumns := []int{}
	for column := 0; column < numColumns; column++ {
		if !isPivot[column] {
			freeVariableColumns = append(freeVariableColumns, column)
		}
	}

	return pivotColumns, freeVariableColumns
}

func calculateSearchUpperBound(joltageRequirements []int, freeVariableColumns []int) int {
	maxJoltage := 0
	for _, requirement := range joltageRequirements {
		if requirement > maxJoltage {
			maxJoltage = requirement
		}
	}

	upperBound := maxJoltage
	if len(freeVariableColumns) == 0 {
		upperBound = 0
	} else if upperBound > 200 {
		upperBound = 200
	}

	return upperBound
}

func searchFreeVariableCombinations(
	freeVariableIndex int,
	freeVariableColumns []int,
	upperBound int,
	buttonPresses []int,
	matrix [][]int,
	joltageRequirements []int,
	pivotColumns []int,
	minimumClicks *int,
) {
	if freeVariableIndex == len(freeVariableColumns) {
		solution := solveWithBackSubstitution(buttonPresses, matrix, joltageRequirements, pivotColumns)
		if solution == nil {
			return
		}

		totalClicks := sumButtonPresses(solution)
		if *minimumClicks == -1 || totalClicks < *minimumClicks {
			*minimumClicks = totalClicks
		}
		return
	}

	freeVariableColumn := freeVariableColumns[freeVariableIndex]
	for value := 0; value <= upperBound; value++ {
		buttonPresses[freeVariableColumn] = value
		searchFreeVariableCombinations(freeVariableIndex+1, freeVariableColumns, upperBound,
			buttonPresses, matrix, joltageRequirements, pivotColumns, minimumClicks)
	}
}

func solveWithBackSubstitution(
	buttonPresses []int,
	matrix [][]int,
	joltageRequirements []int,
	pivotColumns []int,
) []int {
	numRows := len(matrix)
	numColumns := len(matrix[0])

	solution := make([]int, numColumns)
	copy(solution, buttonPresses)

	for row := numRows - 1; row >= 0; row-- {
		pivotColumn := pivotColumns[row]

		if pivotColumn == -1 {
			if joltageRequirements[row] != 0 {
				return nil // Inconsistent system
			}
			continue
		}

		sum := 0
		for column := pivotColumn + 1; column < numColumns; column++ {
			sum += matrix[row][column] * solution[column]
		}

		remaining := joltageRequirements[row] - sum
		pivotValue := matrix[row][pivotColumn]

		if pivotValue == 0 || remaining%pivotValue != 0 {
			return nil // No integer solution
		}

		solution[pivotColumn] = remaining / pivotValue
		if solution[pivotColumn] < 0 {
			return nil // Negative button presses not allowed
		}
	}

	return solution
}

func sumButtonPresses(buttonPresses []int) int {
	total := 0
	for _, presses := range buttonPresses {
		total += presses
	}
	return total
}

func getIndicatorLightStateKey(state MachineState) string {
	key := ""
	for _, light := range state.IndicatorLights {
		if light {
			key += "#"
		} else {
			key += "."
		}
	}
	return key
}

func indicatorLightsMatch(goalState []bool, currentState []bool) bool {
	for index, goalLight := range goalState {
		if goalLight != currentState[index] {
			return false
		}
	}
	return true
}

func applyButtonToIndicatorLights(state MachineState, button Button) MachineState {
	newState := MachineState{
		IndicatorLights: make([]bool, len(state.IndicatorLights)),
	}
	copy(newState.IndicatorLights, state.IndicatorLights)

	for _, lightIndex := range button.Toggle {
		newState.IndicatorLights[lightIndex] = !newState.IndicatorLights[lightIndex]
	}

	return newState
}
func parseInput(lines []string) []MachineSchematic {
	machineSchematics := []MachineSchematic{}
	for _, line := range lines {
		schematic := parseSchematic(line)
		machineSchematics = append(machineSchematics, schematic)
	}
	return machineSchematics
}

func parseSchematic(line string) MachineSchematic {
	parts := strings.Split(line, " ")
	indicatorLights := parseIndicatorLights(parts[0])
	buttons := parseButtons(parts[1 : len(parts)-1])
	joltageRequirements := parseJoltageRequirements(parts[len(parts)-1])

	return MachineSchematic{
		IndicatorLights:     indicatorLights,
		Buttons:             buttons,
		JoltageRequirements: joltageRequirements,
	}
}

func parseIndicatorLights(part string) []bool {
	// Format: [##.....#.], where # is on and . is off
	indicatorLights := []bool{}
	content := part[1 : len(part)-1] // Strip brackets

	for _, char := range content {
		indicatorLights = append(indicatorLights, char == '#')
	}

	return indicatorLights
}

func parseButtons(parts []string) []Button {
	buttons := []Button{}
	for _, part := range parts {
		button := parseButton(part)
		buttons = append(buttons, button)
	}
	return buttons
}

func parseButton(buttonSchematic string) Button {
	// Format: (0,1,5,8)
	content := buttonSchematic[1 : len(buttonSchematic)-1] // Strip parentheses

	if content == "" {
		return Button{Toggle: []int{}}
	}

	lightIndices := strings.Split(content, ",")
	toggles := []int{}

	for _, lightIndex := range lightIndices {
		index, _ := strconv.Atoi(lightIndex)
		toggles = append(toggles, index)
	}

	return Button{Toggle: toggles}
}

func parseJoltageRequirements(part string) []int {
	// Format: {53,78,43,44,33,73,46,81,60}
	content := part[1 : len(part)-1] // Strip curly braces
	joltageStrings := strings.Split(content, ",")
	joltageRequirements := []int{}

	for _, joltageString := range joltageStrings {
		joltage, _ := strconv.Atoi(joltageString)
		joltageRequirements = append(joltageRequirements, joltage)
	}

	return joltageRequirements
}

func main() {
	lines, err := common.ReadLines("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	fmt.Println("Day 10")
	fmt.Println("-----")
	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}

type MachineSchematic struct {
	IndicatorLights     []bool
	Buttons             []Button
	JoltageRequirements []int
}

type Button struct {
	Toggle []int
}

type MachineState struct {
	IndicatorLights []bool
}
