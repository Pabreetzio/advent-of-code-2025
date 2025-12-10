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
	clicks := 0
	for _, machine := range machineSchematics {
		fewestClicks := findFewestButtonClicksForLightMode(machine)
		fmt.Printf("Fewest Clicks: %d\n", fewestClicks)
		clicks += fewestClicks
	}
	return clicks
}

func findFewestButtonClicksForLightMode(machine MachineSchematic) int {
	initialState := MachineState{IndicatorLights: make([]bool, len(machine.IndicatorLights))}
	currentMachineStates := []MachineState{
		initialState,
	}
	exploredStates := map[string]bool{
		getStateKey(initialState): true,
	}
	newMachineStates := []MachineState{}
	unsolved := true
	clicks := 0
	for unsolved {
		for _, machineState := range currentMachineStates {
			for _, button := range machine.Buttons {
				newMachineState := generateNewState(machineState, button)
				stateKey := getStateKey(newMachineState)
				if !exploredStates[stateKey] {
					if isSolved(machine.IndicatorLights, newMachineState.IndicatorLights) {
						unsolved = false
						return clicks + 1
					}
					newMachineStates = append(newMachineStates, newMachineState)
					exploredStates[stateKey] = true
				}
			}
		}
		currentMachineStates = newMachineStates
		newMachineStates = []MachineState{}
		clicks++
	}
	return 0
}

// this isn't performant enough to work yet, checking  in to go to bed.
func findFewestButtonClicksForJoltageMode(machine MachineSchematic) int {
	initialState := make([]int, len(machine.JoltageRequirements)) //all zeros
	currentMachineStates := [][]int{
		initialState,
	}
	newMachineStates := [][]int{}
	unsolved := true
	clicks := 0
	for unsolved {
		for _, machineState := range currentMachineStates {
			for _, button := range machine.Buttons {
				newMachineState := generateNewJoltageState(machineState, button)
				//stateKey := getJoltageStateKey(newMachineState)
				//if !exploredStates[stateKey] {
				if isJoltageSolved(machine.JoltageRequirements, newMachineState) {
					unsolved = false
					return clicks + 1
				}
				busted := getBusted(machine.JoltageRequirements, newMachineState)
				if busted {
					continue
				} else {
					newMachineStates = append(newMachineStates, newMachineState)
				}
				//exploredStates[stateKey] = true
				//}
			}
		}
		currentMachineStates = newMachineStates
		newMachineStates = [][]int{}
		clicks++
	}
	return 0
}

func getBusted(goalJoltage []int, currentJoltage []int) bool {
	for index, joltage := range currentJoltage {
		if joltage > goalJoltage[index] {
			return true
		}
	}
	return false
}

func isJoltageSolved(goalJoltage []int, currentJoltage []int) bool {
	for index, joltage := range goalJoltage {
		if joltage != currentJoltage[index] {
			return false
		}
	}
	return true
}

func generateNewJoltageState(state []int, button Button) []int {
	newState := make([]int, len(state))
	copy(newState, state)
	for _, toggle := range button.Toggle {
		newState[toggle]++
	}
	return newState
}

func getStateKey(state MachineState) string {
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

func isSolved(goalState []bool, litLights []bool) bool {
	for lightIndex, light := range goalState {
		if light != litLights[lightIndex] {
			return false
		}
	}
	return true
}

func generateNewState(state MachineState, button Button) MachineState {
	newState := MachineState{
		IndicatorLights: make([]bool, len(state.IndicatorLights)),
	}
	copy(newState.IndicatorLights, state.IndicatorLights)
	for _, toggle := range button.Toggle {
		newState.IndicatorLights[toggle] = !newState.IndicatorLights[toggle]
	}
	return newState
}

func part2(lines []string) int {
	machineSchematics := parseInput(lines)
	clicks := 0
	for _, machine := range machineSchematics {
		fewestClicks := findFewestButtonClicksForJoltageMode(machine)
		fmt.Printf("Fewest Clicks: %d\n", fewestClicks)
		clicks += fewestClicks
	}
	return clicks
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
	//is in format [##.....#.], where # is on and . is off
	indicatorLights := []bool{}
	for _, char := range part[1 : len(part)-1] {
		if char == '#' {
			indicatorLights = append(indicatorLights, true)
		} else {
			indicatorLights = append(indicatorLights, false)
		}
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
	//is in format (0,1,5,8)
	//strip off parentheses
	buttonSchematic = buttonSchematic[1 : len(buttonSchematic)-1]
	lightIndices := strings.Split(buttonSchematic, ",")
	toggles := []int{}
	for _, lightIndex := range lightIndices {
		index, _ := strconv.Atoi(lightIndex)
		toggles = append(toggles, index)
	}
	return Button{
		Toggle: toggles,
	}
}

func parseJoltageRequirements(part string) []int {
	//is in format {53,78,43,44,33,73,46,81,60}
	//strip off curly braces
	part = part[1 : len(part)-1]
	joltageStrings := strings.Split(part, ",")
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
