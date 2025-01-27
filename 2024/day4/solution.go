package main

import (
	"aoc2024/utils"
	"bufio"
	"fmt"
	"os"
)

func main() {
	wordSearch := parseInput()
	part1(wordSearch)
	part2(wordSearch)
}

func parseInput() []string {
	fileName := "./input"
	file, err := os.Open(fileName) // Read the entire file
	if err != nil {
		panic(fmt.Sprintf("Failure reading file: %s", fileName))
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	wordSearch := make([]string, 0)
	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		wordSearch = append(wordSearch, line)
		i++
	}

	return wordSearch
}

func part1(wordSearch []string) {
	ROWS := len(wordSearch)
	COLS := len(wordSearch[0])
	word := "XMAS"
	count := 0

	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			char := wordSearch[i][j]
			if char == 'X' {
				if isWordPresent(word, utils.Tuple[int, int]{i, j}, func(t utils.Tuple[int, int]) utils.Tuple[int, int] { return utils.Tuple[int, int]{t.Left + 1, t.Right} }, wordSearch) {
					count++
				}
				if isWordPresent(word, utils.Tuple[int, int]{i, j}, func(t utils.Tuple[int, int]) utils.Tuple[int, int] { return utils.Tuple[int, int]{t.Left, t.Right + 1} }, wordSearch) {
					count++
				}
				if isWordPresent(word, utils.Tuple[int, int]{i, j}, func(t utils.Tuple[int, int]) utils.Tuple[int, int] {
					return utils.Tuple[int, int]{t.Left + 1, t.Right + 1}
				}, wordSearch) {
					count++
				}
				if isWordPresent(word, utils.Tuple[int, int]{i, j}, func(t utils.Tuple[int, int]) utils.Tuple[int, int] { return utils.Tuple[int, int]{t.Left - 1, t.Right} }, wordSearch) {
					count++
				}
				if isWordPresent(word, utils.Tuple[int, int]{i, j}, func(t utils.Tuple[int, int]) utils.Tuple[int, int] { return utils.Tuple[int, int]{t.Left, t.Right - 1} }, wordSearch) {
					count++
				}
				if isWordPresent(word, utils.Tuple[int, int]{i, j}, func(t utils.Tuple[int, int]) utils.Tuple[int, int] {
					return utils.Tuple[int, int]{t.Left - 1, t.Right - 1}
				}, wordSearch) {
					count++
				}
				if isWordPresent(word, utils.Tuple[int, int]{i, j}, func(t utils.Tuple[int, int]) utils.Tuple[int, int] {
					return utils.Tuple[int, int]{t.Left + 1, t.Right - 1}
				}, wordSearch) {
					count++
				}
				if isWordPresent(word, utils.Tuple[int, int]{i, j}, func(t utils.Tuple[int, int]) utils.Tuple[int, int] {
					return utils.Tuple[int, int]{t.Left - 1, t.Right + 1}
				}, wordSearch) {
					count++
				}
			}
		}
	}
	fmt.Println(fmt.Sprintf("Day 4 Part 1 Result: %d", count))
}

func part2(wordSearch []string) {
	ROWS := len(wordSearch)
	COLS := len(wordSearch[0])
	prefix := "AM"
	sufix := "AS"
	count := 0

	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			char := wordSearch[i][j]
			if char == 'A' {
				if isWordPresent(prefix, utils.Tuple[int, int]{i, j}, func(t utils.Tuple[int, int]) utils.Tuple[int, int] {
					return utils.Tuple[int, int]{t.Left + 1, t.Right + 1}
				}, wordSearch) &&
					isWordPresent(sufix, utils.Tuple[int, int]{i, j}, func(t utils.Tuple[int, int]) utils.Tuple[int, int] {
						return utils.Tuple[int, int]{t.Left - 1, t.Right - 1}
					}, wordSearch) &&
					isWordPresent(prefix, utils.Tuple[int, int]{i, j}, func(t utils.Tuple[int, int]) utils.Tuple[int, int] {
						return utils.Tuple[int, int]{t.Left + 1, t.Right - 1}
					}, wordSearch) &&
					isWordPresent(sufix, utils.Tuple[int, int]{i, j}, func(t utils.Tuple[int, int]) utils.Tuple[int, int] {
						return utils.Tuple[int, int]{t.Left - 1, t.Right + 1}
					}, wordSearch) {
					count++
				}
				if isWordPresent(sufix, utils.Tuple[int, int]{i, j}, func(t utils.Tuple[int, int]) utils.Tuple[int, int] {
					return utils.Tuple[int, int]{t.Left + 1, t.Right + 1}
				}, wordSearch) &&
					isWordPresent(prefix, utils.Tuple[int, int]{i, j}, func(t utils.Tuple[int, int]) utils.Tuple[int, int] {
						return utils.Tuple[int, int]{t.Left - 1, t.Right - 1}
					}, wordSearch) &&
					isWordPresent(sufix, utils.Tuple[int, int]{i, j}, func(t utils.Tuple[int, int]) utils.Tuple[int, int] {
						return utils.Tuple[int, int]{t.Left + 1, t.Right - 1}
					}, wordSearch) &&
					isWordPresent(prefix, utils.Tuple[int, int]{i, j}, func(t utils.Tuple[int, int]) utils.Tuple[int, int] {
						return utils.Tuple[int, int]{t.Left - 1, t.Right + 1}
					}, wordSearch) {
					count++
				}
				if isWordPresent(prefix, utils.Tuple[int, int]{i, j}, func(t utils.Tuple[int, int]) utils.Tuple[int, int] {
					return utils.Tuple[int, int]{t.Left + 1, t.Right + 1}
				}, wordSearch) &&
					isWordPresent(sufix, utils.Tuple[int, int]{i, j}, func(t utils.Tuple[int, int]) utils.Tuple[int, int] {
						return utils.Tuple[int, int]{t.Left - 1, t.Right - 1}
					}, wordSearch) &&
					isWordPresent(sufix, utils.Tuple[int, int]{i, j}, func(t utils.Tuple[int, int]) utils.Tuple[int, int] {
						return utils.Tuple[int, int]{t.Left + 1, t.Right - 1}
					}, wordSearch) &&
					isWordPresent(prefix, utils.Tuple[int, int]{i, j}, func(t utils.Tuple[int, int]) utils.Tuple[int, int] {
						return utils.Tuple[int, int]{t.Left - 1, t.Right + 1}
					}, wordSearch) {
					count++
				}
				if isWordPresent(sufix, utils.Tuple[int, int]{i, j}, func(t utils.Tuple[int, int]) utils.Tuple[int, int] {
					return utils.Tuple[int, int]{t.Left + 1, t.Right + 1}
				}, wordSearch) &&
					isWordPresent(prefix, utils.Tuple[int, int]{i, j}, func(t utils.Tuple[int, int]) utils.Tuple[int, int] {
						return utils.Tuple[int, int]{t.Left - 1, t.Right - 1}
					}, wordSearch) &&
					isWordPresent(prefix, utils.Tuple[int, int]{i, j}, func(t utils.Tuple[int, int]) utils.Tuple[int, int] {
						return utils.Tuple[int, int]{t.Left + 1, t.Right - 1}
					}, wordSearch) &&
					isWordPresent(sufix, utils.Tuple[int, int]{i, j}, func(t utils.Tuple[int, int]) utils.Tuple[int, int] {
						return utils.Tuple[int, int]{t.Left - 1, t.Right + 1}
					}, wordSearch) {
					count++
				}
			}
		}
	}
	fmt.Println(fmt.Sprintf("Day 4 Part 2 Result: %d", count))
}

func isWordPresent(word string, start utils.Tuple[int, int], nextIndexFunc func(utils.Tuple[int, int]) utils.Tuple[int, int], wordSearch []string) bool {
	ROWS := len(wordSearch)
	COLS := len(wordSearch[0])
	wordLen := len(word)

	wordSearchIndex := start
	if wordSearchIndex.Left < 0 || wordSearchIndex.Left >= ROWS || wordSearchIndex.Right < 0 || wordSearchIndex.Right >= COLS {
		return false
	}

	i := 0
	for i < wordLen && word[i] == wordSearch[wordSearchIndex.Left][wordSearchIndex.Right] {
		i++
		wordSearchIndex = nextIndexFunc(wordSearchIndex)
		if i < wordLen && (wordSearchIndex.Left < 0 || wordSearchIndex.Left >= ROWS || wordSearchIndex.Right < 0 || wordSearchIndex.Right >= COLS) {
			return false
		}
	}

	return i == wordLen
}
