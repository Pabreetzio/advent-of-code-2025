package main

import (
	"advent-of-code-2025/common"
	"fmt"
	"log"
	"strings"
)

func part1(lines []string) int {
	devices := parseInput(lines)
	numberOfPathsOut := getNumberOfOutputsForSource(devices, "you", "out")
	return numberOfPathsOut
}

func getNumberOfOutputsForSource(devices map[string]*Device, sourceId string, destinationId string) (numberOfOutputs int) {
	device := devices[sourceId]
	return getNumberOfOutputs(device, destinationId)
}

func getNumberOfOutputs(device *Device, destinationId string) (numberOfOutputs int) {
	numberOfPathsOut := 0
	for _, outputDevice := range device.outputs {
		if outputDevice.name == "out" {
			numberOfPathsOut++
		} else {
			additionalPathsOut := getNumberOfOutputs(outputDevice, destinationId)
			numberOfPathsOut += additionalPathsOut
		}
	}
	return numberOfPathsOut
}

func parseInput(lines []string) map[string]*Device {
	devices := map[string]*Device{}
	out := Device{
		name:    "out",
		pathsTo: map[string]int{"out": 1},
	}
	devices["out"] = &out
	for _, line := range lines {
		deviceName := strings.Split(line, ":")[0]
		device := Device{
			name:    deviceName,
			outputs: []*Device{},
			pathsTo: map[string]int{deviceName: 1},
		}
		devices[device.name] = &device
	}
	for _, line := range lines {
		deviceName := strings.Split(line, ":")[0]
		device := devices[deviceName]
		outputs := strings.Split(strings.TrimSpace(strings.Split(line, ":")[1]), " ")
		for _, outputName := range outputs {
			output := devices[outputName]
			device.outputs = append(device.outputs, output)
		}
		devices[device.name] = device
	}
	return devices
}

type Device struct {
	name     string
	outputs  []*Device
	inputs   []*Device
	pathsOut *int
	pathsTo  map[string]int
}

func part2(lines []string) int {
	devices := parseInput(lines)
	numberOfPathsOut := getNumberOfOutputsThroughDacAndFft(devices)
	return numberOfPathsOut
}
func getNumberOfOutputsThroughDacAndFft(devices map[string]*Device) int {
	addInputs(devices)
	devices["out"].pathsOut = new(int)
	*devices["out"].pathsOut = 1

	devices["out"].pathsTo["out"] = 1

	populatePathsTo(devices["dac"], "out")
	populatePathsTo(devices["fft"], "dac")
	populatePathsTo(devices["svr"], "fft")
	return devices["svr"].pathsTo["fft"] * devices["fft"].pathsTo["dac"] * devices["dac"].pathsTo["out"]
}
func populatePathsOut(device *Device) {
	if device.pathsOut != nil {
		return
	}
	pathsOut := 0
	for _, outputDevice := range device.outputs {
		if outputDevice.pathsOut == nil {
			populatePathsOut(outputDevice)
		}
		pathsOut += *outputDevice.pathsOut
	}
	device.pathsOut = &pathsOut
}

func populatePathsTo(Device *Device, destinationId string) {
	if _, exists := Device.pathsTo[destinationId]; exists {
		return
	}
	pathsTo := 0
	for _, outputDevice := range Device.outputs {
		if outputDevice.pathsTo[destinationId] == 0 {
			populatePathsTo(outputDevice, destinationId)
		}
		pathsTo += outputDevice.pathsTo[destinationId]
	}
	Device.pathsTo[destinationId] = pathsTo
}

func addInputs(devices map[string]*Device) {
	for _, device := range devices {
		device.inputs = []*Device{}
	}
	for _, device := range devices {
		for _, outputDevice := range device.outputs {
			outputDevice.inputs = append(outputDevice.inputs, device)
		}
	}
}

func main() {
	lines, err := common.ReadLines("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	fmt.Println("Day 11")
	fmt.Println("-----")
	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}
