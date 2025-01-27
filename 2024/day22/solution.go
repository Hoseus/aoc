package main

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	secretNumbers := parseInput()
	part1(secretNumbers)
	part2(secretNumbers)
}

func parseInput() []int {
	fileName := "./input"
	bytes, err := os.ReadFile(fileName) // Read the entire file
	if err != nil {
		panic(fmt.Sprintf("Failure reading file: %s", fileName))
	}
	lines := strings.Split(string(bytes), "\n")

	secretNumbers := make([]int, len(lines))
	for i, line := range lines {
		secretNumbers[i], _ = strconv.Atoi(line)
	}

	return secretNumbers
}

func part1(secretNumbers []int) {
	result := 0
	for _, secretNumber := range secretNumbers {
		result += getNSecretNumber(secretNumber, 2000)
	}
	fmt.Println(fmt.Sprintf("Day 22 Part 1 Result: %d", result))
}

func part2(secretNumbers []int) {
	pricesForChanges := make(map[string]int)
	for _, secretNumber := range secretNumbers {
		priceChanges := getNPriceChanges(secretNumber, 2000)
		calculatePricesForChanges(priceChanges, &pricesForChanges)
	}

	result := 0
	for _, priceForChanges := range pricesForChanges {
		if priceForChanges > result {
			result = priceForChanges
		}
	}
	fmt.Println(fmt.Sprintf("Day 22 Part 2 Result: %d", result))
}

func getNSecretNumber(secretNumber int, n int) int {
	current := secretNumber
	for i := 0; i < n; i++ {
		current = getNextSecretNumber(current)
	}
	return current
}

func getNPriceChanges(secretNumber int, n int) []utils.Tuple[int, int] {
	priceChanges := make([]utils.Tuple[int, int], n)

	previousPrice := getPrice(secretNumber)
	current := getNextSecretNumber(secretNumber)
	for i := 0; i < n; i++ {
		price := getPrice(current)
		priceChange := price - previousPrice
		priceChanges[i] = utils.Tuple[int, int]{Left: price, Right: priceChange}
		current = getNextSecretNumber(current)
		previousPrice = price
	}
	return priceChanges
}

func getNextSecretNumber(secretNumber int) int {
	step1 := prune(mix(secretNumber, secretNumber*64))
	step2 := prune(mix(step1, step1/32))
	step3 := prune(mix(step2, step2*2048))
	return step3
}

func mix(secretNumber int, value int) int {
	return secretNumber ^ value
}

func prune(secretNumber int) int {
	mod := 16777216
	return (secretNumber%mod + mod) % mod
}

func calculatePricesForChanges(priceChanges []utils.Tuple[int, int], pricesForChanges *map[string]int) {
	derefPricesByVariation := *pricesForChanges

	changesOccurred := make(map[string]bool)
	for i := 3; i < len(priceChanges); i++ {
		key := strconv.Itoa(priceChanges[i-3].Right) + strconv.Itoa(priceChanges[i-2].Right) + strconv.Itoa(priceChanges[i-1].Right) + strconv.Itoa(priceChanges[i].Right)
		value := priceChanges[i].Left
		if !changesOccurred[key] {
			derefPricesByVariation[key] += value
			changesOccurred[key] = true
		}
	}
}

func getPrice(secretNumber int) int {
	secretNumberString := strconv.Itoa(secretNumber)
	price, _ := strconv.Atoi(secretNumberString[len(secretNumberString)-1:])
	return price
}
