package main

import (
	"aoc2024/utils"
	"fmt"
	"maps"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type LogicGateType int

const (
	AND LogicGateType = iota
	OR
	XOR
	Failed
)

type LogicGate struct {
	logicGateType                         LogicGateType
	inputCable1, inputCable2, outputCable string
}

func main() {
	cableValues, logicGates := parseInput()
	part1(cableValues, logicGates)
	part2(cableValues, logicGates)
}

func parseInput() (map[string]int, []LogicGate) {
	fileName := "./input"
	bytes, err := os.ReadFile(fileName) // Read the entire file
	if err != nil {
		panic(fmt.Sprintf("Failure reading file: %s", fileName))
	}
	inputSections := strings.Split(string(bytes), "\n\n")

	cableValues := make(map[string]int)
	for _, line := range strings.Split(inputSections[0], "\n") {
		split := strings.Split(line, ": ")
		cableName := split[0]
		cableValue, _ := strconv.Atoi(split[1])
		cableValues[cableName] = cableValue
	}

	logicGates := make([]LogicGate, 0)
	for _, line := range strings.Split(inputSections[1], "\n") {
		split := strings.Split(line, " ")
		input1 := split[0]
		logicGateName := split[1]
		input2 := split[2]
		output := split[4]

		var logicGateType LogicGateType
		switch logicGateName {
		case "AND":
			logicGateType = AND
		case "OR":
			logicGateType = OR
		case "XOR":
			logicGateType = XOR
		default:
			panic("Something went wrong")
		}

		logicGates = append(logicGates, LogicGate{logicGateType: logicGateType, inputCable1: input1, inputCable2: input2, outputCable: output})
	}

	return cableValues, logicGates
}

func part1(cableValues map[string]int, logicGates []LogicGate) {
	zCablesCount := 0
	q := utils.NewDeque[LogicGate]()
	for _, logicGate := range logicGates {
		q.PushBack(logicGate)
		if strings.HasPrefix(logicGate.outputCable, "z") {
			zCablesCount++
		}
	}

	zCableResults := make([]int, zCablesCount)
	for q.Len() > 0 {
		logicGate, _ := q.PopFront()
		input1, input1Exists := cableValues[logicGate.inputCable1]
		input2, input2Exists := cableValues[logicGate.inputCable2]

		if input1Exists && input2Exists {
			output := executeLogicGate(input1, input2, logicGate.logicGateType)
			cableValues[logicGate.outputCable] = output
			if strings.HasPrefix(logicGate.outputCable, "z") {
				zCableIndex := extractNumber(logicGate.outputCable)
				zCableResults[zCableIndex] = output
			}
		} else {
			q.PushBack(logicGate)
		}

		output := executeLogicGate(input1, input2, logicGate.logicGateType)

		if strings.HasPrefix(logicGate.outputCable, "z") {
			zCableIndex := extractNumber(logicGate.outputCable)
			zCableResults[zCableIndex] = output
		}
	}

	binaryNumber := ""
	for i := zCablesCount - 1; i >= 0; i-- {
		binaryNumber += strconv.Itoa(zCableResults[i])
	}
	decimalNumber, _ := strconv.ParseInt(binaryNumber, 2, 64)
	fmt.Println(fmt.Sprintf("Day 24 Part 1 Result: %d", decimalNumber))
}

func part2(_ map[string]int, logicGates []LogicGate) {
	temp := make(map[string]bool)

	for _, logicGate := range logicGates {
		if strings.HasPrefix(logicGate.outputCable, "z") {
			zCableIndex := extractNumber(logicGate.outputCable)
			if logicGate.logicGateType != XOR && zCableIndex != 45 {
				temp[logicGate.outputCable] = true
			}
		} else if !isXOrY(logicGate.inputCable1) && !isXOrY(logicGate.inputCable2) && logicGate.inputCable1[0] != logicGate.inputCable2[0] && logicGate.logicGateType == XOR {
			temp[logicGate.outputCable] = true
		}

		if logicGate.logicGateType == XOR && isXOrY(logicGate.inputCable1) && isXOrY(logicGate.inputCable2) && logicGate.inputCable1[0] != logicGate.inputCable2[0] {
			isValid := false
			for _, logicGateDependency := range logicGates {
				if logicGateDependency.logicGateType == XOR && (logicGateDependency.inputCable1 == logicGate.outputCable || logicGateDependency.inputCable2 == logicGate.outputCable) {
					isValid = true
				}
			}
			if !isValid {
				temp[logicGate.outputCable] = true
			}
		}

		if logicGate.logicGateType == AND && isXOrY(logicGate.inputCable1) && isXOrY(logicGate.inputCable2) && logicGate.inputCable1[0] != logicGate.inputCable2[0] {
			isValid := false
			for _, logicGateDependency := range logicGates {
				if logicGateDependency.logicGateType == OR && (logicGateDependency.inputCable1 == logicGate.outputCable || logicGateDependency.inputCable2 == logicGate.outputCable) {
					isValid = true
				}
			}
			if !isValid {
				temp[logicGate.outputCable] = true
			}
		}
	}

	swappedCables := slices.Collect(maps.Keys(temp))
	slices.Sort(swappedCables)

	result := ""
	for _, swappedCable := range swappedCables {
		result += swappedCable + ","
	}

	fmt.Println(fmt.Sprintf("Day 24 Part 1 Result: %s", result[:len(result)-1]))
}

func executeLogicGate(input1, input2 int, logicGateType LogicGateType) int {
	switch logicGateType {
	case AND:
		return and(input1, input2)
	case OR:
		return or(input1, input2)
	case XOR:
		return xor(input1, input2)
	default:
		panic("Something went wrong")
	}
}

func and(input1, input2 int) int {
	if input1 == 1 && input2 == 1 {
		return 1
	}
	return 0
}

func or(input1, input2 int) int {
	if input1 == 1 || input2 == 1 {
		return 1
	}
	return 0
}

func xor(input1, input2 int) int {
	if input1 != input2 {
		return 1
	}
	return 0
}

func extractNumber(cableName string) int {
	regex := regexp.MustCompile("[1-9][0-9]*")
	match := regex.FindString(cableName)
	result, _ := strconv.Atoi(match)
	return result
}

func isXOrY(wire string) bool {
	wireNumber := extractNumber(wire)
	return (strings.HasPrefix(wire, "x") || strings.HasPrefix(wire, "y")) && wireNumber != 0
}
