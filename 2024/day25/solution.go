package main

import (
	"fmt"
	"os"
	"strings"
)

type Lock struct {
	rows, cols int
	heights    []int
}

type Key struct {
	rows, cols int
	heights    []int
}

func main() {
	locks, keys := parseInput()
	part1(locks, keys)
	part2()
}

func parseInput() ([]Lock, []Key) {
	fileName := "./input"
	bytes, err := os.ReadFile(fileName) // Read the entire file
	if err != nil {
		panic(fmt.Sprintf("Failure reading file: %s", fileName))
	}
	inputSections := strings.Split(string(bytes), "\n\n")

	locks := make([]Lock, 0)
	keys := make([]Key, 0)

	for _, inputSection := range inputSections {
		lines := strings.Split(inputSection, "\n")
		rows := len(lines)
		cols := len(lines[0])

		heights := make([]int, cols)
		for i := 0; i < cols; i++ {
			countHashChar := 0
			for j := 0; j < rows; j++ {
				if lines[j][i] == '#' {
					countHashChar++
				}
			}
			heights[i] = countHashChar - 1
		}

		if lines[0] == strings.Repeat("#", cols) {
			locks = append(locks, Lock{rows: rows, cols: cols, heights: heights})
		} else {
			keys = append(keys, Key{rows: rows, cols: cols, heights: heights})
		}
	}

	return locks, keys
}

func part1(locks []Lock, keys []Key) {
	result := 0
	for _, lock := range locks {
		for _, key := range keys {
			keyFits := true
			if lock.rows == key.rows && lock.cols == key.cols {
				for i := 0; i < len(key.heights); i++ {
					if (lock.rows - 1 - lock.heights[i]) <= key.heights[i] {
						keyFits = false
						break
					}
				}
			}
			if keyFits {
				result++
			}
		}
	}
	fmt.Println(fmt.Sprintf("Day 25 Part 1 Result: %d", result))
}

func part2() {
	fmt.Println(fmt.Sprintf("Day 25 Part 2 Result: %d", 0))
}
