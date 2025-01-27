package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	memory := parseInput()

	part1(memory)
	part2(memory)
}

func parseInput() string {
	fileName := "./input"
	bytes, err := os.ReadFile(fileName) // Read the entire file
	if err != nil {
		panic(fmt.Sprintf("Failure reading file: %s", fileName))
	}

	return string(bytes)
}

func part1(memory string) {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	mulInstructions := re.FindAllStringSubmatch(memory, -1)

	sum := 0
	for _, mulInstruction := range mulInstructions {
		n1, _ := strconv.Atoi(mulInstruction[1])
		n2, _ := strconv.Atoi(mulInstruction[2])
		sum += n1 * n2
	}
	fmt.Println(fmt.Sprintf("Day 3 Part 1 Result: %d", sum))
}

func part2(memory string) {
	re := regexp.MustCompile(`(do)\(\)|(don't)\(\)|(mul)\((\d{1,3}),(\d{1,3})\)`)
	instructions := re.FindAllStringSubmatch(memory, -1)

	sum := 0
	enabled := true
	for _, instruction := range instructions {
		doInstructionName := instruction[1]
		dontInstructionName := instruction[2]
		mulInstructionName := instruction[3]

		if doInstructionName != "" {
			enabled = true
		} else if dontInstructionName != "" {
			enabled = false
		} else if mulInstructionName != "" && enabled {
			n1, _ := strconv.Atoi(instruction[4])
			n2, _ := strconv.Atoi(instruction[5])
			sum += n1 * n2
		}

	}
	fmt.Println(fmt.Sprintf("Day 3 Part 2 Result: %d", sum))
}
