package main

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	equations := parseInput()
	part1(&equations)
	part2(&equations)
}

func parseInput() []utils.Tuple[int, []int] {
	fileName := "./input"
	bytes, err := os.ReadFile(fileName) // Read the entire file
	if err != nil {
		panic(fmt.Sprintf("Failure reading file: %s", fileName))
	}
	plainTextEquations := strings.Split(string(bytes), "\n")

	equations := make([]utils.Tuple[int, []int], 0)

	for _, plainTextEquation := range plainTextEquations {
		plainTextSplitEquation := strings.Split(plainTextEquation, ":")
		testValue, _ := strconv.Atoi(plainTextSplitEquation[0])
		plainTextNumbers := strings.Fields(plainTextSplitEquation[1])
		numbers := make([]int, 0)
		for _, plainTextNumber := range plainTextNumbers {
			number, _ := strconv.Atoi(plainTextNumber)
			numbers = append(numbers, number)
		}
		equations = append(equations, utils.Tuple[int, []int]{testValue, numbers})
	}

	return equations
}

func part1(equations *[]utils.Tuple[int, []int]) {
	acc := 0
	for _, equation := range *equations {
		testValue := equation.Left
		numbers := equation.Right
		results := make([]int, 0)
		for _, number := range numbers {
			results = solvePossibilitiesPart1(results, number, testValue)
			if len(results) == 0 {
				break
			}
			if results[len(results)-1] == testValue {
				acc += testValue
				break
			}
		}
	}

	fmt.Println(fmt.Sprintf("Day 7 Part 1 Result: %d", acc))
}

func part2(equations *[]utils.Tuple[int, []int]) {
	acc := 0
	for _, equation := range *equations {
		testValue := equation.Left
		numbers := equation.Right
		results := make([]int, 0)
		for _, number := range numbers {
			results = solvePossibilitiesPart2(results, number, testValue)
			if len(results) == 0 {
				break
			}
			if results[len(results)-1] == testValue {
				acc += testValue
				break
			}
		}
	}

	fmt.Println(fmt.Sprintf("Day 7 Part 2 Result: %d", acc))
}

func solvePossibilitiesPart1(numbers []int, target, expectedResult int) []int {
	if len(numbers) == 0 {
		return []int{target}
	}

	results := make([]int, 0)
	for _, number := range numbers {
		add := performOperation(number, target, ADD, expectedResult, &results)
		if add == expectedResult {
			break
		}
		multi := performOperation(number, target, MULTI, expectedResult, &results)
		if multi == expectedResult {
			break
		}
	}
	return results
}

func solvePossibilitiesPart2(numbers []int, target, expectedResult int) []int {
	if len(numbers) == 0 {
		return []int{target}
	}

	results := make([]int, 0)
	for _, number := range numbers {
		add := performOperation(number, target, ADD, expectedResult, &results)
		if add == expectedResult {
			break
		}
		multi := performOperation(number, target, MULTI, expectedResult, &results)
		if multi == expectedResult {
			break
		}
		concat := performOperation(number, target, CONCAT, expectedResult, &results)
		if concat == expectedResult {
			break
		}
	}
	return results
}

func performOperation(a, b int, operator Operator, expectedResult int, resultAccumulator *[]int) int {
	result := 0
	switch operator {
	case MULTI:
		result = a * b
	case ADD:
		result = a + b
	case CONCAT:
		result, _ = strconv.Atoi(fmt.Sprintf("%d%d", a, b))
	}

	if result <= expectedResult {
		*resultAccumulator = append(*resultAccumulator, result)
	}

	return result
}

type Operator int

const (
	MULTI Operator = iota + 1
	ADD
	CONCAT
)
