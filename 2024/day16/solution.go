package main

import (
	"aoc2024/utils"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

type Position struct {
	Coordinates utils.Tuple[int, int]
	Direction   rune
}

func main() {
	labyrinthMap, startPosition, endPosition := parseInput()
	part1(labyrinthMap, startPosition, endPosition)
	part2(labyrinthMap, startPosition, endPosition)
}

func parseInput() ([][]rune, utils.Tuple[int, int], utils.Tuple[int, int]) {
	fileName := "./input"
	bytes, err := os.ReadFile(fileName) // Read the entire file
	if err != nil {
		panic(fmt.Sprintf("Failure reading file: %s", fileName))
	}
	lines := strings.Split(string(bytes), "\n")

	startPosition := utils.Tuple[int, int]{}
	endPosition := utils.Tuple[int, int]{}
	labyrinthMap := make([][]rune, len(lines))
	for i, row := range lines {
		for j, char := range row {
			if char == 'S' {
				startPosition = utils.Tuple[int, int]{Left: i, Right: j}
			}

			if char == 'E' {
				endPosition = utils.Tuple[int, int]{Left: i, Right: j}
			}
		}
		labyrinthMap[i] = []rune(row)
	}

	return labyrinthMap, startPosition, endPosition
}

func part1(labyrinthMap [][]rune, startPosition utils.Tuple[int, int], endPosition utils.Tuple[int, int]) {
	fmt.Println(fmt.Sprintf("Day 16 Part 1 Result: %d", findLowestScore(&labyrinthMap, startPosition, '>', endPosition)))
}

func part2(labyrinthMap [][]rune, startPosition utils.Tuple[int, int], endPosition utils.Tuple[int, int]) {
	shortestPaths := findShortestPaths(&labyrinthMap, startPosition, '>', endPosition)
	uniqueCoordinates := make(map[utils.Tuple[int, int]]bool)
	for _, shortestPath := range shortestPaths {
		for _, coordinates := range shortestPath {
			uniqueCoordinates[coordinates] = true
		}
	}
	fmt.Println(fmt.Sprintf("Day 16 Part 2 Result: %d", len(uniqueCoordinates)))
}

func findLowestScore(labyrinthMap *[][]rune, startCoordinates utils.Tuple[int, int], direction rune, endCoordinates utils.Tuple[int, int]) int {
	startPosition := Position{Coordinates: startCoordinates, Direction: direction}
	startScoredPosition := utils.Tuple[Position, int]{
		Left:  startPosition,
		Right: 0,
	}

	minHeap := utils.NewHeap[utils.Tuple[Position, int]](
		func(a, b utils.Tuple[Position, int]) bool {
			return a.Right < b.Right
		},
	)
	scores := make(map[Position]int)

	scores[startPosition] = 0
	minHeap.PushT(startScoredPosition)
	for minHeap.Len() > 0 {
		currentScoredPosition := minHeap.PopT()
		currentPosition := currentScoredPosition.Left
		currentScore := currentScoredPosition.Right

		if currentPosition.Coordinates == endCoordinates {
			return currentScore
		}

		for _, adjScoredPosition := range getAdj(currentScoredPosition, labyrinthMap) {
			adjPosition := adjScoredPosition.Left
			adjScore := adjScoredPosition.Right
			if _, exists := scores[adjPosition]; !exists {
				scores[adjPosition] = math.MaxInt
			}
			if adjScore < scores[adjPosition] {
				scores[adjPosition] = adjScore
				minHeap.PushT(adjScoredPosition)
			}
		}
	}

	return -1
}

func findShortestPaths(labyrinthMap *[][]rune, startCoordinates utils.Tuple[int, int], direction rune, endCoordinates utils.Tuple[int, int]) [][]utils.Tuple[int, int] {
	startPosition := Position{Coordinates: startCoordinates, Direction: direction}
	startScoredPosition := utils.Tuple[[]Position, int]{
		Left:  []Position{startPosition},
		Right: 0,
	}

	minHeap := utils.NewHeap[utils.Tuple[[]Position, int]](
		func(a, b utils.Tuple[[]Position, int]) bool {
			return a.Right < b.Right
		},
	)
	scores := make(map[Position]int)

	resultPaths := make(map[int][][]utils.Tuple[int, int])
	scores[startPosition] = 0
	minHeap.PushT(startScoredPosition)
	for minHeap.Len() > 0 {
		currentScoredPath := minHeap.PopT()
		currentPath := currentScoredPath.Left
		currentScore := currentScoredPath.Right
		currentPosition := currentPath[len(currentPath)-1]
		currentScoredPosition := utils.Tuple[Position, int]{
			Left:  currentPosition,
			Right: currentScoredPath.Right,
		}

		if currentPosition.Coordinates == endCoordinates {
			if _, exists := resultPaths[currentScore]; !exists && len(resultPaths) == 1 {
				break
			}

			path := make([]utils.Tuple[int, int], 0)
			for _, position := range currentPath {
				path = append(path, position.Coordinates)
			}
			resultPaths[currentScore] = append(resultPaths[currentScore], path)
		}

		for _, adjScoredPosition := range getAdj(currentScoredPosition, labyrinthMap) {
			adjPosition := adjScoredPosition.Left
			adjScore := adjScoredPosition.Right
			if _, exists := scores[adjPosition]; !exists {
				scores[adjPosition] = math.MaxInt
			}
			if adjScore <= scores[adjPosition] {
				scores[adjPosition] = adjScore
				adjPath := append([]Position{}, currentPath...)
				adjPath = append(adjPath, adjPosition)
				adjScoredPath := utils.Tuple[[]Position, int]{
					Left:  adjPath,
					Right: adjScoredPosition.Right,
				}
				minHeap.PushT(adjScoredPath)
			}
		}
	}

	minScore := math.MaxInt
	for k := range resultPaths {
		minScore = min(minScore, k)
	}

	return resultPaths[minScore]
}

func getAdj(scoredPosition utils.Tuple[Position, int], labyrinthMap *[][]rune) []utils.Tuple[Position, int] {
	moves := []utils.Tuple[int, int]{
		{Left: -1, Right: 0},
		{Left: 0, Right: 1},
		{Left: 1, Right: 0},
		{Left: 0, Right: -1},
	}
	directions := []rune{'^', '>', 'v', '<'}
	directionsCount := len(directions)

	derefMap := *labyrinthMap
	rows := len(derefMap)
	cols := len(derefMap[0])
	position := scoredPosition.Left
	score := scoredPosition.Right

	result := make([]utils.Tuple[Position, int], 0)

	straightDirectionIndex := slices.Index(directions, position.Direction)
	straightMove := moves[straightDirectionIndex]
	newPosition := Position{
		Coordinates: utils.Tuple[int, int]{
			Left:  position.Coordinates.Left + straightMove.Left,
			Right: position.Coordinates.Right + straightMove.Right,
		},
		Direction: position.Direction,
	}
	if newPosition.Coordinates.Left >= 0 && newPosition.Coordinates.Left < rows &&
		newPosition.Coordinates.Right >= 0 && newPosition.Coordinates.Right < cols &&
		derefMap[newPosition.Coordinates.Left][newPosition.Coordinates.Right] != '#' {
		result = append(
			result,
			utils.Tuple[Position, int]{Left: newPosition, Right: score + 1},
		)
	}

	leftDirection := directions[((straightDirectionIndex-1)%directionsCount+directionsCount)%directionsCount]
	newPosition = Position{
		Coordinates: position.Coordinates,
		Direction:   leftDirection,
	}
	result = append(
		result,
		utils.Tuple[Position, int]{Left: newPosition, Right: score + 1000},
	)

	rightDirection := directions[((straightDirectionIndex+1)%directionsCount+directionsCount)%directionsCount]
	newPosition = Position{
		Coordinates: position.Coordinates,
		Direction:   rightDirection,
	}
	result = append(
		result,
		utils.Tuple[Position, int]{Left: newPosition, Right: score + 1000},
	)

	return result
}

func printMap(labyrinthMap [][]rune, position Position) {
	labyrinthMap[position.Coordinates.Left][position.Coordinates.Right] = position.Direction
	for _, row := range labyrinthMap {
		fmt.Println(string(row))
	}
}

func printMap2(labyrinthMap [][]rune, positions map[utils.Tuple[int, int]]bool) {
	for position := range positions {
		labyrinthMap[position.Left][position.Right] = 'O'
	}
	for _, row := range labyrinthMap {
		fmt.Println(string(row))
	}
}
