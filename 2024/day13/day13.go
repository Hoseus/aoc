package main

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type MachineConfiguration struct {
	AButton utils.Tuple[int, int]
	BButton utils.Tuple[int, int]
	Prize   utils.Tuple[int, int]
}

func main() {
	machineConfigurations := parseInput()

	part1(machineConfigurations)
	part2(machineConfigurations)
}

func parseInput() []MachineConfiguration {
	fileName := "./day13.input"
	bytes, err := os.ReadFile(fileName) // Read the entire file
	if err != nil {
		panic(fmt.Sprintf("Failure reading file: %s", fileName))
	}
	machineConfigurations := strings.Split(string(bytes), "\n\n")

	result := make([]MachineConfiguration, 0)
	for _, machineConfiguration := range machineConfigurations {
		buttonAMatch := regexp.MustCompile(`Button A: X\+(\d+), Y\+(\d+)\n`).FindStringSubmatch(machineConfiguration)
		buttonBMatch := regexp.MustCompile(`Button B: X\+(\d+), Y\+(\d+)`).FindStringSubmatch(machineConfiguration)
		prizeMatch := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`).FindStringSubmatch(machineConfiguration)

		aX, _ := strconv.Atoi(buttonAMatch[1])
		aY, _ := strconv.Atoi(buttonAMatch[2])

		bX, _ := strconv.Atoi(buttonBMatch[1])
		bY, _ := strconv.Atoi(buttonBMatch[2])

		pX, _ := strconv.Atoi(prizeMatch[1])
		pY, _ := strconv.Atoi(prizeMatch[2])

		result = append(
			result,
			MachineConfiguration{
				AButton: utils.Tuple[int, int]{Left: aX, Right: aY},
				BButton: utils.Tuple[int, int]{Left: bX, Right: bY},
				Prize:   utils.Tuple[int, int]{Left: pX, Right: pY},
			},
		)
	}

	return result
}

func part1(machineConfigurations []MachineConfiguration) {
	result := 0
	for _, machineConfiguration := range machineConfigurations {
		moves := findMoves(
			machineConfiguration.AButton.Left,
			machineConfiguration.AButton.Right,
			machineConfiguration.BButton.Left,
			machineConfiguration.BButton.Right,
			machineConfiguration.Prize.Left,
			machineConfiguration.Prize.Right,
		)
		result += moves.Left*3 + moves.Right
	}
	fmt.Println(fmt.Sprintf("Day 13 Part 1 Result: %d", result))
}

func part2(machineConfigurations []MachineConfiguration) {
	result := 0
	for _, machineConfiguration := range machineConfigurations {
		moves := findMoves(
			machineConfiguration.AButton.Left,
			machineConfiguration.AButton.Right,
			machineConfiguration.BButton.Left,
			machineConfiguration.BButton.Right,
			machineConfiguration.Prize.Left+10000000000000,
			machineConfiguration.Prize.Right+10000000000000,
		)
		result += moves.Left*3 + moves.Right
	}
	fmt.Println(fmt.Sprintf("Day 13 Part 2 Result: %d", result))
}

func findMoves(aX, aY, bX, bY, pX, pY int) utils.Tuple[int, int] {
	aXF := float64(aX)
	aYF := float64(aY)

	bXF := float64(bX)
	bYF := float64(bY)

	pXF := float64(pX)
	pYF := float64(pY)

	aXWithBY := aXF * bYF
	pXWithBY := pXF * bYF

	aYWithBX := aYF * bXF
	pYWithBX := pYF * bXF

	a := (pXWithBY - pYWithBX) / (aXWithBY - aYWithBX)
	b := (pYF - aYF*a) / bYF

	if a != float64(int(a)) || b != float64(int(b)) {
		return utils.Tuple[int, int]{}
	}

	return utils.Tuple[int, int]{Left: int(a), Right: int(b)}
}
