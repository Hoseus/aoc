package main

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"strings"
)

func main() {
	gardenMap := parseInput()

	part1(gardenMap)
	part2(gardenMap)
}

func parseInput() []string {
	fileName := "./day12.input"
	bytes, err := os.ReadFile(fileName) // Read the entire file
	if err != nil {
		panic(fmt.Sprintf("Failure reading file: %s", fileName))
	}
	gardenMap := strings.Split(string(bytes), "\n")

	return gardenMap
}

func part1(gardenMap []string) {
	start := utils.Tuple[int, int]{}
	visited := map[utils.Tuple[int, int]]bool{}
	plants := utils.NewDeque[utils.Tuple[int, int]]()
	plants.PushBack(start)
	result := 0
	for plants.Len() > 0 {
		plant, _ := plants.PopFront()
		if !visited[plant] {
			fenceResult, fenceAdj := solveFencePart1(plant, &gardenMap, &visited)
			result += fenceResult
			for _, plantAdj := range fenceAdj {
				plants.PushBack(plantAdj)
			}
		}
	}

	fmt.Println(fmt.Sprintf("Day 12 Part 1 Result: %d", result))
}

func part2(gardenMap []string) {
	start := utils.Tuple[int, int]{}
	visited := map[utils.Tuple[int, int]]bool{}
	plants := utils.NewDeque[utils.Tuple[int, int]]()
	plants.PushBack(start)
	result := 0
	for plants.Len() > 0 {
		plant, _ := plants.PopFront()
		if !visited[plant] {
			fenceResult, fenceAdj := solveFencePart2(plant, &gardenMap, &visited)
			result += fenceResult
			for _, plantAdj := range fenceAdj {
				plants.PushBack(plantAdj)
			}
		}
	}

	fmt.Println(fmt.Sprintf("Day 12 Part 2 Result: %d", result))
}

func getAdj(point utils.Tuple[int, int], gardenMap *[]string) []utils.Tuple[int, int] {
	var adjPossibilities = []utils.Tuple[int, int]{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	derefGardenMap := *gardenMap

	rows := len(derefGardenMap)
	cols := len(derefGardenMap[0])

	row := point.Left
	col := point.Right

	adj := make([]utils.Tuple[int, int], 0)
	for _, adjPossibility := range adjPossibilities {
		adjRow := row + adjPossibility.Left
		adjCol := col + adjPossibility.Right
		if adjRow >= 0 && adjRow < rows && adjCol >= 0 && adjCol < cols {
			adj = append(adj, utils.Tuple[int, int]{Left: adjRow, Right: adjCol})
		}
	}
	return adj
}

func getVerticeCount(point utils.Tuple[int, int], gardenMap *[]string) int {
	var adjPossibilities = []utils.Tuple[int, int]{
		{-1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
		{0, -1},
		{-1, -1},
	}
	possibilitiesSize := len(adjPossibilities)

	derefGardenMap := *gardenMap

	rows := len(derefGardenMap)
	cols := len(derefGardenMap[0])

	row := point.Left
	col := point.Right
	plantType := rune(derefGardenMap[point.Left][point.Right])

	verticeCount := 0
	for i := 0; i < possibilitiesSize; i += 2 {
		adjDir := adjPossibilities[i%possibilitiesSize]
		nextAdjDiagonalDir := adjPossibilities[(i+1)%possibilitiesSize]
		nextAdjDir := adjPossibilities[(i+2)%possibilitiesSize]

		adjRow := row + adjDir.Left
		adjCol := col + adjDir.Right
		nextAdjDiagonalRow := row + nextAdjDiagonalDir.Left
		nextAdjDiagonalCol := col + nextAdjDiagonalDir.Right
		nextAdjRow := row + nextAdjDir.Left
		nextAdjCol := col + nextAdjDir.Right

		if (!isInBounds(adjRow, adjCol, rows, cols) && !isInBounds(nextAdjRow, nextAdjCol, rows, cols)) ||
			(isInBounds(adjRow, adjCol, rows, cols) && !isInBounds(nextAdjRow, nextAdjCol, rows, cols) &&
				rune(derefGardenMap[adjRow][adjCol]) != plantType) ||
			(!isInBounds(adjRow, adjCol, rows, cols) && isInBounds(nextAdjRow, nextAdjCol, rows, cols) &&
				rune(derefGardenMap[nextAdjRow][nextAdjCol]) != plantType) ||
			((isInBounds(adjRow, adjCol, rows, cols) && isInBounds(nextAdjRow, nextAdjCol, rows, cols)) &&
				((rune(derefGardenMap[adjRow][adjCol]) != plantType &&
					rune(derefGardenMap[nextAdjRow][nextAdjCol]) != plantType) ||
					(rune(derefGardenMap[adjRow][adjCol]) == plantType &&
						rune(derefGardenMap[nextAdjRow][nextAdjCol]) == plantType &&
						rune(derefGardenMap[nextAdjDiagonalRow][nextAdjDiagonalCol]) != plantType))) {
			verticeCount++
		}
	}

	return verticeCount
}

func isInBounds(x, y int, m, n int) bool {
	return x >= 0 && x < m && y >= 0 && y < n
}

func solveFencePart1(start utils.Tuple[int, int], gardenMap *[]string, visited *map[utils.Tuple[int, int]]bool) (int, []utils.Tuple[int, int]) {
	derefGardenMap := *gardenMap
	derefVisited := *visited

	plantType := rune(derefGardenMap[start.Left][start.Right])

	plantCount := 0
	wallCount := 0
	fenceAdj := make([]utils.Tuple[int, int], 0)

	q := utils.NewDeque[utils.Tuple[int, int]]()
	q.PushBack(start)
	derefVisited[start] = true

	for q.Len() > 0 {
		curr, _ := q.PopFront()

		plantCount++
		wallCount += 4

		for _, adj := range getAdj(curr, gardenMap) {
			plantTypeAdj := rune(derefGardenMap[adj.Left][adj.Right])
			if plantType == plantTypeAdj {
				wallCount--
			}
			if !derefVisited[adj] {
				if plantType == plantTypeAdj {
					derefVisited[adj] = true
					q.PushBack(adj)
				} else {
					fenceAdj = append(fenceAdj, adj)
				}
			}
		}
	}

	return plantCount * wallCount, fenceAdj
}

func solveFencePart2(start utils.Tuple[int, int], gardenMap *[]string, visited *map[utils.Tuple[int, int]]bool) (int, []utils.Tuple[int, int]) {
	derefGardenMap := *gardenMap
	derefVisited := *visited

	plantType := rune(derefGardenMap[start.Left][start.Right])

	plantCount := 0
	verticesCount := 0
	fenceAdj := make([]utils.Tuple[int, int], 0)

	q := utils.NewDeque[utils.Tuple[int, int]]()
	q.PushBack(start)
	derefVisited[start] = true

	for q.Len() > 0 {
		curr, _ := q.PopFront()

		plantCount++
		verticesCount += getVerticeCount(curr, gardenMap)
		for _, adj := range getAdj(curr, gardenMap) {
			plantTypeAdj := rune(derefGardenMap[adj.Left][adj.Right])
			if !derefVisited[adj] {
				if plantType == plantTypeAdj {
					derefVisited[adj] = true
					q.PushBack(adj)
				} else {
					fenceAdj = append(fenceAdj, adj)
				}
			}
		}
	}

	return plantCount * verticesCount, fenceAdj
}
