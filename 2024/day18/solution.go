package main

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytePositions := parseInput()
	part1(bytePositions, 71, 71, 0, 1024, utils.Tuple[int, int]{Left: 0, Right: 0}, utils.Tuple[int, int]{Left: 70, Right: 70})
	part2(bytePositions, 71, 71, 1025, len(bytePositions), utils.Tuple[int, int]{Left: 0, Right: 0}, utils.Tuple[int, int]{Left: 70, Right: 70})
}

func parseInput() []utils.Tuple[int, int] {
	fileName := "./input"
	bytes, err := os.ReadFile(fileName) // Read the entire file
	if err != nil {
		panic(fmt.Sprintf("Failure reading file: %s", fileName))
	}

	result := make([]utils.Tuple[int, int], 0)
	for _, line := range strings.Split(string(bytes), "\n") {
		splitLine := strings.Split(line, ",")
		col, _ := strconv.Atoi(splitLine[0])
		row, _ := strconv.Atoi(splitLine[1])
		result = append(result, utils.Tuple[int, int]{Left: row, Right: col})
	}

	return result
}

func part1(bytePositions []utils.Tuple[int, int], rows, cols int, simulateStart, simulateLimit int, start utils.Tuple[int, int], end utils.Tuple[int, int]) {
	memorySpace := buildMemorySpace(bytePositions[simulateStart:simulateLimit], rows, cols)

	q := utils.NewDeque[utils.Tuple[utils.Tuple[int, int], int]]()
	visited := make(map[utils.Tuple[int, int]]bool)

	q.PushBack(utils.Tuple[utils.Tuple[int, int], int]{Left: start})
	visited[start] = true

	result := -1
	for q.Len() > 0 {
		current, _ := q.PopFront()
		currentPosition := current.Left
		count := current.Right

		if currentPosition == end {
			result = count
			break
		}

		for _, adjPosition := range getAdj(currentPosition, &memorySpace) {
			if !visited[adjPosition] {
				q.PushBack(utils.Tuple[utils.Tuple[int, int], int]{Left: adjPosition, Right: count + 1})
				visited[adjPosition] = true
			}
		}
	}

	fmt.Println(fmt.Sprintf("Day 18 Part 1 Result: %d", result))
}

func part2(bytePositions []utils.Tuple[int, int], rows, cols int, simulateStart, simulateLimit int, start utils.Tuple[int, int], end utils.Tuple[int, int]) {
	memorySpace := buildMemorySpace(bytePositions[:simulateStart-1], rows, cols)
	result := utils.Tuple[int, int]{}
	for i := simulateStart - 1; i < simulateLimit; i++ {
		bytePosition := bytePositions[i]
		memorySpace[bytePosition.Left][bytePosition.Right] = '#'

		q := utils.NewDeque[utils.Tuple[utils.Tuple[int, int], int]]()
		visited := make(map[utils.Tuple[int, int]]bool)

		q.PushBack(utils.Tuple[utils.Tuple[int, int], int]{Left: start})
		visited[start] = true

		endReached := false
		for q.Len() > 0 {
			current, _ := q.PopFront()
			currentPosition := current.Left
			count := current.Right

			if currentPosition == end {
				endReached = true
				break
			}

			for _, adjPosition := range getAdj(currentPosition, &memorySpace) {
				if !visited[adjPosition] {
					q.PushBack(utils.Tuple[utils.Tuple[int, int], int]{Left: adjPosition, Right: count + 1})
					visited[adjPosition] = true
				}
			}
		}

		if !endReached {
			result = bytePositions[i]
			break
		}
	}

	fmt.Println(fmt.Sprintf("Day 18 Part 2 Result: %d,%d", result.Right, result.Left))
}

func buildMemorySpace(bytePositions []utils.Tuple[int, int], rows, cols int) [][]rune {
	result := make([][]rune, rows)
	for i := range result {
		result[i] = make([]rune, cols)
		for j := range result[i] {
			result[i][j] = '.'
		}
	}

	for _, bytePosition := range bytePositions {
		result[bytePosition.Left][bytePosition.Right] = '#'
	}

	return result
}

func printMemorySpace(memorySpace [][]rune) {
	for _, row := range memorySpace {
		fmt.Println(string(row))
	}
}

func getAdj(position utils.Tuple[int, int], memorySpace *[][]rune) []utils.Tuple[int, int] {
	var adjPossibilities = []utils.Tuple[int, int]{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	derefMemorySpace := *memorySpace

	rows := len(derefMemorySpace)
	cols := len(derefMemorySpace[0])

	row := position.Left
	col := position.Right

	adj := make([]utils.Tuple[int, int], 0)
	for _, adjPossibility := range adjPossibilities {
		adjRow := row + adjPossibility.Left
		adjCol := col + adjPossibility.Right
		if adjRow >= 0 && adjRow < rows && adjCol >= 0 && adjCol < cols && derefMemorySpace[adjRow][adjCol] != '#' {
			adj = append(adj, utils.Tuple[int, int]{Left: adjRow, Right: adjCol})
		}
	}
	return adj
}
