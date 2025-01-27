package main

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	pageOrderingRules, pageUpdates := parseInput()
	part1(&pageOrderingRules, &pageUpdates)
	part2(&pageOrderingRules, &pageUpdates)
}

func parseInput() (map[int][]int, [][]int) {
	fileName := "./input"
	bytes, err := os.ReadFile(fileName) // Read the entire file
	if err != nil {
		panic(fmt.Sprintf("Failure reading file: %s", fileName))
	}
	pagesConfig := strings.Split(string(bytes), "\n\n")
	pageOrderingRulesTmp, pageUpdatesTmp := strings.Split(pagesConfig[0], "\n"), strings.Split(pagesConfig[1], "\n")

	pageOrderingRules := make(map[int][]int, len(pageOrderingRulesTmp))
	for i := 0; i < len(pageOrderingRulesTmp); i++ {
		split := strings.Split(pageOrderingRulesTmp[i], "|")
		before, _ := strconv.Atoi(split[0])
		after, _ := strconv.Atoi(split[1])
		pageOrderingRules[before] = append(pageOrderingRules[before], after)
	}

	for _, v := range pageOrderingRules {
		utils.QuickSort(&v, func(a, b int) bool { return a <= b })
	}

	pageUpdates := make([][]int, len(pageUpdatesTmp))
	for i := 0; i < len(pageUpdatesTmp); i++ {
		split := strings.Split(pageUpdatesTmp[i], ",")
		for j := 0; j < len(split); j++ {
			pageNumber, _ := strconv.Atoi(split[j])
			pageUpdates[i] = append(pageUpdates[i], pageNumber)
		}
	}

	return pageOrderingRules, pageUpdates
}

func part1(pageOrderingRules *map[int][]int, pagesUpdates *[][]int) {
	sum := 0
	for _, pageNumbers := range *pagesUpdates {
		valid := true
		for i := 0; i < len(pageNumbers)-1; i++ {
			current := pageNumbers[i]
			next := pageNumbers[i+1]
			orderingRules := (*pageOrderingRules)[current]

			foundNextPageIndex := utils.BinarySearch(&orderingRules, 0, len(orderingRules)-1, next)
			if foundNextPageIndex < 0 {
				valid = false
				break
			}
		}
		if valid {
			midIndex := (len(pageNumbers) - 1) / 2
			sum += pageNumbers[midIndex]
		}
	}
	fmt.Println(fmt.Sprintf("Day 4 Part 1 Result: %d", sum))
}

func part2(pageOrderingRules *map[int][]int, pagesUpdates *[][]int) {
	derefPageOrderingRules := *pageOrderingRules
	derefPageUpdates := *pagesUpdates

	sum := 0
	for _, pageNumbers := range derefPageUpdates {
		valid := true
		for i := 0; i < len(pageNumbers)-1; i++ {
			current := pageNumbers[i]
			next := pageNumbers[i+1]
			orderingRules := derefPageOrderingRules[current]

			foundNextPageIndex := utils.BinarySearch(&orderingRules, 0, len(orderingRules)-1, next)
			if foundNextPageIndex < 0 {
				valid = false
				break
			}
		}
		if !valid {
			utils.QuickSort(
				&pageNumbers,
				func(a, b int) bool {
					orderingRules := derefPageOrderingRules[a]
					return utils.BinarySearch(&orderingRules, 0, len(orderingRules)-1, b) >= 0
				},
			)
			midIndex := (len(pageNumbers) - 1) / 2
			sum += pageNumbers[midIndex]
		}
	}
	fmt.Println(fmt.Sprintf("Day 4 Part 1 Result: %d", sum))
}
