package main

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	stones := parseInput()

	part1(stones, 25)
	part2(stones, 75)
}

func parseInput() map[string]int {
	fileName := "./input"
	bytes, err := os.ReadFile(fileName) // Read the entire file
	if err != nil {
		panic(fmt.Sprintf("Failure reading file: %s", fileName))
	}
	stones := make(map[string]int)

	for _, stone := range strings.Split(string(bytes), " ") {
		stones[stone] = 1
	}

	return stones
}

func part1(stones map[string]int, blinks int) {
	for i := 0; i < blinks; i++ {
		stones = blink(stones)
	}
	result := 0
	for _, count := range stones {
		result += count
	}
	fmt.Println(fmt.Sprintf("Day 11 Part 1 Result: %d", result))
}

func part2(stones map[string]int, blinks int) {
	for i := 0; i < blinks; i++ {
		stones = blink(stones)
	}
	result := 0
	for _, count := range stones {
		result += count
	}
	fmt.Println(fmt.Sprintf("Day 11 Part 2 Result: %d", result))
}

func blink(stones map[string]int) map[string]int {
	result := make(map[string]int)
	for stone, count := range stones {
		if stone == "0" {
			result["1"] += count
		} else if utils.Even(len(stone)) {
			half := len(stone) / 2
			n1, _ := strconv.Atoi(stone[0:half])
			n2, _ := strconv.Atoi(stone[half:])
			result[strconv.Itoa(n1)] += count
			result[strconv.Itoa(n2)] += count
		} else {
			n, _ := strconv.Atoi(stone)
			result[strconv.Itoa(n*2024)] += count
		}
	}
	return result
}
