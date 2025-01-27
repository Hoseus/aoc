package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	towelPatterns, desiredDesigns := parseInput()
	part1(towelPatterns, desiredDesigns)
	part2(towelPatterns, desiredDesigns)
}

func parseInput() ([]string, []string) {
	fileName := "./input"
	bytes, err := os.ReadFile(fileName) // Read the entire file
	if err != nil {
		panic(fmt.Sprintf("Failure reading file: %s", fileName))
	}
	inputParts := strings.Split(string(bytes), "\n\n")

	towelPatterns := strings.Split(inputParts[0], ", ")
	desiredDesigns := strings.Split(inputParts[1], "\n")

	return towelPatterns, desiredDesigns
}

func part1(towelPatterns []string, desiredDesigns []string) {
	result := 0
	for _, desiredPattern := range desiredDesigns {
		if isDesignPossible(desiredPattern, towelPatterns) {
			result++
		}
	}
	fmt.Println(fmt.Sprintf("Day 19 Part 1 Result: %d", result))
}

func part2(towelPatterns []string, desiredDesigns []string) {
	result := 0
	for _, desiredPattern := range desiredDesigns {
		result += designTowelCombinations(desiredPattern, towelPatterns)
	}
	fmt.Println(fmt.Sprintf("Day 19 Part 2 Result: %d", result))
}

func isDesignPossible(desiredDesign string, towelPatterns []string) bool {
	if desiredDesign == "" {
		return true
	}

	for _, towelPattern := range towelPatterns {
		if strings.HasPrefix(desiredDesign, towelPattern) && isDesignPossible(desiredDesign[len(towelPattern):], towelPatterns) {
			return true
		}
	}

	return false
}

func designTowelCombinations(desiredDesign string, towelPatterns []string) int {
	asd := make(map[string]int)
	r := designTowelCombinationsWithMemoization(desiredDesign, towelPatterns, asd)
	return r
}

func designTowelCombinationsWithMemoization(desiredDesign string, towelPatterns []string, memo map[string]int) int {
	if desiredDesign == "" {
		return 1
	}

	if value, exists := memo[desiredDesign]; exists {
		return value
	}

	result := 0
	for _, towelPattern := range towelPatterns {
		if strings.HasPrefix(desiredDesign, towelPattern) {
			result += designTowelCombinationsWithMemoization(desiredDesign[len(towelPattern):], towelPatterns, memo)
		}
	}

	memo[desiredDesign] = result
	return result
}
