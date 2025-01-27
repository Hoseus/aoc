package main

import (
	"aoc2024/utils"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	codes := parseInput()
	numericKeypad, directionalKeypad := buildNumericKeypad(), buildDirectionalKeypad()
	part1(codes, numericKeypad, directionalKeypad)
	part2(codes, numericKeypad, directionalKeypad)
}

func parseInput() [][]rune {
	fileName := "./input"
	bytes, err := os.ReadFile(fileName) // Read the entire file
	if err != nil {
		panic(fmt.Sprintf("Failure reading file: %s", fileName))
	}
	lines := strings.Split(string(bytes), "\n")

	codes := make([][]rune, len(lines))
	for i, row := range lines {
		codes[i] = []rune(row)
	}

	return codes
}

var directions = []utils.Tuple[rune, utils.Tuple[int, int]]{
	{Left: 'v', Right: utils.Tuple[int, int]{Left: 1, Right: 0}},
	{Left: '<', Right: utils.Tuple[int, int]{Left: 0, Right: -1}},
	{Left: '^', Right: utils.Tuple[int, int]{Left: -1, Right: 0}},
	{Left: '>', Right: utils.Tuple[int, int]{Left: 0, Right: 1}},
}

func buildNumericKeypad() map[rune][]utils.Tuple[rune, rune] {
	keypad := [][]rune{
		{'7', '8', '9'},
		{'4', '5', '6'},
		{'1', '2', '3'},
		{'-', '0', 'A'},
	}

	return buildKeypadAdjList(keypad)
}

func buildDirectionalKeypad() map[rune][]utils.Tuple[rune, rune] {
	keypad := [][]rune{
		{'-', '^', 'A'},
		{'<', 'v', '>'},
	}

	return buildKeypadAdjList(keypad)
}

func buildKeypadAdjList(keypad [][]rune) map[rune][]utils.Tuple[rune, rune] {
	rows := len(keypad)
	cols := len(keypad[0])

	adjacencyList := make(map[rune][]utils.Tuple[rune, rune])

	for i := range keypad {
		for j := range keypad[i] {
			key := keypad[i][j]
			if key == '-' {
				continue
			}
			adj := make([]utils.Tuple[rune, rune], 0)
			for _, direction := range directions {
				nextI := i + direction.Right.Left
				nextJ := j + direction.Right.Right
				if nextI >= 0 && nextI < rows && nextJ >= 0 && nextJ < cols {
					nextKey := keypad[nextI][nextJ]
					if nextKey != '-' {
						adj = append(adj, utils.Tuple[rune, rune]{Left: nextKey, Right: direction.Left})
					}
				}
			}
			adjacencyList[key] = adj
		}
	}

	return adjacencyList
}

var keypadPathsCache = make(map[utils.Tuple[rune, rune]][][]rune)

func part1(codes [][]rune, numericKeypad map[rune][]utils.Tuple[rune, rune], directionalKeypad map[rune][]utils.Tuple[rune, rune]) {
	result := 0
	memo := make(map[utils.Tuple[string, int]]int)
	for _, code := range codes {
		numPadSequenceLen := findNumPadSequenceLen(numericKeypad, directionalKeypad, code, 3, &memo)
		result += numPadSequenceLen * extractNumber(code)
	}
	fmt.Println(fmt.Sprintf("Day 21 Part 1 Result: %d", result))
}

func part2(codes [][]rune, numericKeypad map[rune][]utils.Tuple[rune, rune], directionalKeypad map[rune][]utils.Tuple[rune, rune]) {
	result := 0
	memo := make(map[utils.Tuple[string, int]]int)
	for _, code := range codes {
		numPadSequenceLen := findNumPadSequenceLen(numericKeypad, directionalKeypad, code, 26, &memo)
		result += numPadSequenceLen * extractNumber(code)
	}
	fmt.Println(fmt.Sprintf("Day 21 Part 2 Result: %d", result))
}

func findNumPadSequenceLen(numericKeypad map[rune][]utils.Tuple[rune, rune], directionalKeypad map[rune][]utils.Tuple[rune, rune], numericCode []rune, indirectionCount int, memo *map[utils.Tuple[string, int]]int) int {
	if result, exists := (*memo)[utils.Tuple[string, int]{Left: string(numericCode), Right: indirectionCount}]; exists {
		return result
	}

	result := 0
	currentKey := 'A'
	for _, numKey := range numericCode {
		pathsToNumKey := findShortestPaths(numericKeypad, currentKey, numKey)
		minDirSequence := -1
		if len(pathsToNumKey) == 0 {
			minDirSequence = 1
		}
		for _, pathToNumKey := range pathsToNumKey {
			pathToNumKey = append(pathToNumKey, 'A')
			dirPadSequenceLen := findDirPadSequenceLen(directionalKeypad, pathToNumKey, indirectionCount-1, memo)
			if minDirSequence < 0 || dirPadSequenceLen < minDirSequence {
				minDirSequence = dirPadSequenceLen
			}
		}
		result += minDirSequence
		currentKey = numKey
	}

	(*memo)[utils.Tuple[string, int]{Left: string(numericCode), Right: indirectionCount}] = result
	return result
}

func findDirPadSequenceLen(dirKeypad map[rune][]utils.Tuple[rune, rune], dirKeySequence []rune, indirectionCount int, memo *map[utils.Tuple[string, int]]int) int {
	if result, exists := (*memo)[utils.Tuple[string, int]{Left: string(dirKeySequence), Right: indirectionCount}]; exists {
		return result
	}

	if indirectionCount == 0 {
		return len(dirKeySequence)
	}

	dirPadDirSequence := 0
	currentKey := 'A'
	for _, dirKey := range dirKeySequence {
		pathsToDirKey := findShortestPaths(dirKeypad, currentKey, dirKey)
		minDirSequence := -1
		if len(pathsToDirKey) == 0 {
			minDirSequence = 1
		}
		for _, pathToDirKey := range pathsToDirKey {
			pathToDirKey = append(pathToDirKey, 'A')
			dirPadSequence := findDirPadSequenceLen(dirKeypad, pathToDirKey, indirectionCount-1, memo)
			if minDirSequence < 0 || dirPadSequence < minDirSequence {
				minDirSequence = dirPadSequence
			}
		}
		dirPadDirSequence += minDirSequence
		currentKey = dirKey
	}

	(*memo)[utils.Tuple[string, int]{Left: string(dirKeySequence), Right: indirectionCount}] = dirPadDirSequence
	return dirPadDirSequence
}

func findShortestPaths(keypad map[rune][]utils.Tuple[rune, rune], startKey, endKey rune) [][]rune {
	type Position struct {
		Key  rune
		Path []rune
	}

	if result, exists := keypadPathsCache[utils.Tuple[rune, rune]{Left: startKey, Right: endKey}]; exists {
		return result
	}

	if startKey == endKey {
		return make([][]rune, 0)
	}

	minHeap := utils.NewHeap[utils.Tuple[Position, int]](
		func(a, b utils.Tuple[Position, int]) bool {
			return a.Right < b.Right
		},
	)
	scores := make(map[rune]int)
	for key := range keypad {
		scores[key] = math.MaxInt
	}

	minHeap.PushT(utils.Tuple[Position, int]{Left: Position{Key: startKey, Path: make([]rune, 0)}, Right: 0})
	scores[startKey] = 0

	paths := make([][]rune, 0)
	for minHeap.Len() > 0 {
		currentPosition := minHeap.PopT()
		currentKey := currentPosition.Left.Key
		currentPath := currentPosition.Left.Path
		currentScore := currentPosition.Right

		if currentKey == endKey && currentScore <= scores[currentKey] {
			paths = append(paths, currentPath)
		}

		for _, adj := range keypad[currentKey] {
			nextKey := adj.Left
			adjDirection := adj.Right
			adjScore := currentScore + 1
			if len(currentPath) > 0 && adjDirection != currentPath[len(currentPath)-1] {
				adjScore = adjScore + 2
			}

			if adjScore <= scores[nextKey] {
				scores[nextKey] = adjScore
				nextPath := make([]rune, len(currentPath)+1)
				copy(nextPath, currentPath)
				nextPath[len(currentPath)] = adjDirection
				minHeap.PushT(
					utils.Tuple[Position, int]{
						Left:  Position{Key: nextKey, Path: nextPath},
						Right: adjScore,
					},
				)
			}
		}
	}

	keypadPathsCache[utils.Tuple[rune, rune]{Left: startKey, Right: endKey}] = paths
	return paths
}

func extractNumber(code []rune) int {
	codeString := string(code)
	regex := regexp.MustCompile("[1-9][0-9]*")
	match := regex.FindString(codeString)
	result, _ := strconv.Atoi(match)
	return result
}
