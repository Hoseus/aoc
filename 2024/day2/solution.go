package main

import (
	"aoc2024/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reports := parseInput()
	part1(reports)
	part2(reports)
}

func parseInput() [][]int {
	fileName := "./input"
	file, err := os.Open(fileName) // Read the entire file
	if err != nil {
		panic(fmt.Sprintf("Failure reading file: %s", fileName))
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	reports := make([][]int, 0)
	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		levels, _ := utils.StringsToNumbers(strings.Fields(line))

		reports = append(reports, levels)
		i++
	}

	return reports
}

func part1(reports [][]int) {
	count := 0
	for _, levels := range reports {
		if len(levels) <= 1 {
			count++
			continue
		}

		i := 1
		isAscending := levels[i-1] < levels[i]
		for i < len(levels) {
			prev := levels[i-1]
			current := levels[i]

			if !isValid(isAscending, prev, current) {
				break
			}

			i++
		}

		if i == len(levels) {
			count++
		}
	}

	fmt.Println(fmt.Sprintf("Day 2 Part 1 Result: %d", count))
}

func part2(reports [][]int) {
	count := 0
	for _, levels := range reports {
		if len(levels) <= 2 {
			count++
			continue
		}

		skipeableAscendingCount := 0
		invalidAscendingCount := 0
		skipeableDescendingCount := 0
		invalidDescendingCount := 0
		i := 1
		for i < len(levels)-1 && (invalidAscendingCount <= 1 || invalidDescendingCount <= 1) {
			prev := levels[i-1]
			current := levels[i]
			next := levels[i+1]

			isValidAscendingLeft := isValid(true, prev, current)

			isValidAscendingSkipped := isValid(true, prev, next)

			isValidAscendingRight := isValid(true, current, next)

			if !isValidAscendingLeft && i == 1 {
				invalidAscendingCount++
			}

			if !isValidAscendingRight {
				invalidAscendingCount++
			}

			if ((!isValidAscendingLeft || !isValidAscendingRight) && isValidAscendingSkipped) || (!isValidAscendingLeft && isValidAscendingRight && i == 1) || (isValidAscendingLeft && !isValidAscendingRight && i == len(levels)-2) {
				if !isValidAscendingLeft && !isValidAscendingRight {
					invalidAscendingCount--
				}
				skipeableAscendingCount++
			}

			isValidDescendingLeft := isValid(false, prev, current)

			isValidDescendingSkipped := isValid(false, prev, next)

			isValidDescendingRight := isValid(false, current, next)

			if !isValidDescendingLeft && i == 1 {
				invalidDescendingCount++
			}

			if !isValidDescendingRight {
				invalidDescendingCount++
			}

			if ((!isValidDescendingLeft || !isValidDescendingRight) && isValidDescendingSkipped) || (!isValidDescendingLeft && isValidDescendingRight && i == 1) || (isValidDescendingLeft && !isValidDescendingRight && i == len(levels)-2) {
				if !isValidDescendingLeft && !isValidDescendingRight {
					invalidDescendingCount--
				}
				skipeableDescendingCount++
			}

			i++
		}

		if i == len(levels)-1 && ((invalidAscendingCount == 0 || (invalidAscendingCount == 1 && skipeableAscendingCount >= 1)) || (invalidDescendingCount == 0 || (invalidDescendingCount == 1 && skipeableDescendingCount >= 1))) {
			count++
		}
	}

	fmt.Println(fmt.Sprintf("Day 2 Part 2 Result: %d", count))
}

func isValid(isAscending bool, left int, right int) bool {
	isValidDiff := utils.Abs(right-left) >= 1 && utils.Abs(right-left) <= 3

	if isAscending {
		return left < right && isValidDiff
	} else {
		return left > right && isValidDiff
	}
}
