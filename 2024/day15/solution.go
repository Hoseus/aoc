package main

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"strings"
)

func main() {
	warehouseMap, robotStartLocation, robotMoves := parseInputPart1()
	part1(warehouseMap, robotStartLocation, robotMoves)

	warehouseMap, robotStartLocation, robotMoves = parseInputPart2()
	part2(warehouseMap, robotStartLocation, robotMoves)
}

func parseInputPart1() ([][]rune, utils.Tuple[int, int], string) {
	fileName := "./input"
	bytes, err := os.ReadFile(fileName) // Read the entire file
	if err != nil {
		panic(fmt.Sprintf("Failure reading file: %s", fileName))
	}
	inputParts := strings.Split(string(bytes), "\n\n")

	warehouseMapString := strings.Split(inputParts[0], "\n")
	warehouseMap := make([][]rune, len(warehouseMapString))
	for i, row := range warehouseMapString {
		warehouseMap[i] = []rune(row)
	}

	robotLocation := utils.Tuple[int, int]{Left: -1, Right: -1}
	for i := 0; i < len(warehouseMap); i++ {
		for j := 0; j < len(warehouseMap[0]); j++ {
			if warehouseMap[i][j] == '@' {
				robotLocation = utils.Tuple[int, int]{Left: i, Right: j}
				break
			}
		}

		if robotLocation.Left != -1 && robotLocation.Right != -1 {
			break
		}
	}

	robotMoves := strings.ReplaceAll(inputParts[1], "\n", "")

	return warehouseMap, robotLocation, robotMoves
}

func parseInputPart2() ([][]rune, utils.Tuple[int, int], string) {
	fileName := "./input"
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		panic(fmt.Sprintf("Failure reading file: %s", fileName))
	}
	inputParts := strings.Split(string(bytes), "\n\n")

	warehouseMapString := strings.Split(inputParts[0], "\n")
	warehouseMap := make([][]rune, len(warehouseMapString))
	for i, row := range warehouseMapString {
		for _, char := range row {
			if char == '#' {
				warehouseMap[i] = append(warehouseMap[i], '#', '#')
			} else if char == 'O' {
				warehouseMap[i] = append(warehouseMap[i], '[', ']')
			} else if char == '.' {
				warehouseMap[i] = append(warehouseMap[i], '.', '.')
			} else if char == '@' {
				warehouseMap[i] = append(warehouseMap[i], '@', '.')
			}
		}
	}

	robotLocation := utils.Tuple[int, int]{Left: -1, Right: -1}
	for i := 0; i < len(warehouseMap); i++ {
		for j := 0; j < len(warehouseMap[0]); j++ {
			if warehouseMap[i][j] == '@' {
				robotLocation = utils.Tuple[int, int]{Left: i, Right: j}
				break
			}
		}

		if robotLocation.Left != -1 && robotLocation.Right != -1 {
			break
		}
	}

	robotMoves := strings.ReplaceAll(inputParts[1], "\n", "")

	return warehouseMap, robotLocation, robotMoves
}

func part1(warehouseMap [][]rune, robotStartPosition utils.Tuple[int, int], robotMoves string) {
	currentPosition := robotStartPosition
	for _, direction := range robotMoves {
		currentPosition = movePart1(currentPosition, direction, &warehouseMap)
	}

	result := 0
	for i := range warehouseMap {
		for j, value := range warehouseMap[i] {
			if value == 'O' {
				result += 100*i + j
			}
		}
	}
	fmt.Println(fmt.Sprintf("Day 15 Part 1 Result: %d", result))
}

func part2(warehouseMap [][]rune, robotStartPosition utils.Tuple[int, int], robotMoves string) {
	currentPosition := robotStartPosition
	for _, direction := range robotMoves {
		currentPosition = movePart2(currentPosition, direction, &warehouseMap)
	}

	result := 0
	for i := range warehouseMap {
		for j, value := range warehouseMap[i] {
			if value == '[' {
				result += 100*i + j
			}
		}
	}
	fmt.Println(fmt.Sprintf("Day 15 Part 2 Result: %d", result))
}

func movePart1(start utils.Tuple[int, int], direction rune, warehouseMap *[][]rune) utils.Tuple[int, int] {
	derefWareHouseMap := *warehouseMap

	visitedPositions := utils.NewDeque[utils.Tuple[int, int]]()
	visitedPositions.PushBack(start)

	nextPosition := getNextPart1(start, direction)

	currentPosition := nextPosition
	for derefWareHouseMap[currentPosition.Left][currentPosition.Right] == 'O' {
		visitedPositions.PushBack(currentPosition)
		currentPosition = getNextPart1(currentPosition, direction)
	}

	for visitedPositions.Len() > 0 && derefWareHouseMap[currentPosition.Left][currentPosition.Right] == '.' {
		previousPosition, _ := visitedPositions.PopBack()
		derefWareHouseMap[currentPosition.Left][currentPosition.Right] = derefWareHouseMap[previousPosition.Left][previousPosition.Right]
		derefWareHouseMap[previousPosition.Left][previousPosition.Right] = '.'
		currentPosition = previousPosition
	}

	if derefWareHouseMap[start.Left][start.Right] == '@' {
		return start
	} else {
		return nextPosition
	}
}

func movePart2(start utils.Tuple[int, int], direction rune, warehouseMap *[][]rune) utils.Tuple[int, int] {
	derefWareHouseMap := *warehouseMap

	visitedPositions := utils.NewDeque[[]utils.Tuple[int, int]]()
	visitedPositions.PushBack([]utils.Tuple[int, int]{start})

	startNextPosition := getNextPart1(start, direction)

	currentPositions := getNextPart2([]utils.Tuple[int, int]{start}, direction, warehouseMap)
	for containsABigBox(currentPositions, warehouseMap) && !containsAnObstruction(currentPositions, warehouseMap) {
		visitedPositions.PushBack(currentPositions)
		currentPositions = getNextPart2(currentPositions, direction, warehouseMap)
	}

	for visitedPositions.Len() > 0 && areAllEmpty(currentPositions, warehouseMap) {
		previousPositions, _ := visitedPositions.PopBack()
		for _, previousPosition := range previousPositions {
			currentPosition := getNextPart1(previousPosition, direction)
			derefWareHouseMap[currentPosition.Left][currentPosition.Right] = derefWareHouseMap[previousPosition.Left][previousPosition.Right]
			derefWareHouseMap[previousPosition.Left][previousPosition.Right] = '.'
		}
		currentPositions = previousPositions
	}

	if derefWareHouseMap[start.Left][start.Right] == '@' {
		return start
	} else {
		return startNextPosition
	}
}

func getNextPart1(position utils.Tuple[int, int], direction rune) utils.Tuple[int, int] {
	positionChange := map[rune]utils.Tuple[int, int]{
		'^': {Left: -1, Right: 0},
		'<': {Left: 0, Right: -1},
		'>': {Left: 0, Right: 1},
		'v': {Left: 1, Right: 0},
	}[direction]

	return utils.Tuple[int, int]{
		Left:  position.Left + positionChange.Left,
		Right: position.Right + positionChange.Right,
	}
}

func getNextPart2(positions []utils.Tuple[int, int], direction rune, warehouseMap *[][]rune) []utils.Tuple[int, int] {
	positionChangeOptions := map[string]utils.Tuple[int, int]{
		"^":   {Left: -1, Right: 0},
		"<":   {Left: 0, Right: -1},
		">":   {Left: 0, Right: 1},
		"v":   {Left: 1, Right: 0},
		"^+<": {Left: -1, Right: -1},
		"^+>": {Left: -1, Right: 1},
		"v+<": {Left: 1, Right: -1},
		"v+>": {Left: 1, Right: 1},
	}
	directionString := string(direction)

	neighbours := make(map[utils.Tuple[int, int]]bool)
	for _, position := range positions {
		positionChange := positionChangeOptions[directionString]
		nextPosition := utils.Tuple[int, int]{
			Left:  position.Left + positionChange.Left,
			Right: position.Right + positionChange.Right,
		}
		if (*warehouseMap)[nextPosition.Left][nextPosition.Right] != '.' {
			neighbours[nextPosition] = true
			if direction == '^' || direction == 'v' {
				if (*warehouseMap)[nextPosition.Left][nextPosition.Right] == ']' {
					positionChangeDiagonalLeft := positionChangeOptions[directionString+"+<"]
					nextPositionLeftNeighbour := utils.Tuple[int, int]{
						Left:  position.Left + positionChangeDiagonalLeft.Left,
						Right: position.Right + positionChangeDiagonalLeft.Right,
					}
					neighbours[nextPositionLeftNeighbour] = true
				} else if (*warehouseMap)[nextPosition.Left][nextPosition.Right] == '[' {
					positionChangeDiagonalRight := positionChangeOptions[directionString+"+>"]
					nextPositionRightNeighbour := utils.Tuple[int, int]{
						Left:  position.Left + positionChangeDiagonalRight.Left,
						Right: position.Right + positionChangeDiagonalRight.Right,
					}
					neighbours[nextPositionRightNeighbour] = true
				}
			}
		}
	}

	result := make([]utils.Tuple[int, int], 0)
	for neighbour := range neighbours {
		result = append(result, neighbour)
	}

	return result
}

func containsAnObstruction(positions []utils.Tuple[int, int], warehouseMap *[][]rune) bool {
	containsObstruction := false
	i := 0
	for i < len(positions) && !containsObstruction {
		position := positions[i]
		if (*warehouseMap)[position.Left][position.Right] == '#' {
			containsObstruction = true
		}
		i++
	}

	return containsObstruction
}

func containsABigBox(positions []utils.Tuple[int, int], warehouseMap *[][]rune) bool {
	containsBox := false
	i := 0
	for i < len(positions) && !containsBox {
		position := positions[i]
		if (*warehouseMap)[position.Left][position.Right] == '[' ||
			(*warehouseMap)[position.Left][position.Right] == ']' {
			containsBox = true
		}
		i++
	}

	return containsBox
}

func areAllEmpty(positions []utils.Tuple[int, int], warehouseMap *[][]rune) bool {
	allEmpty := true
	i := 0
	for i < len(positions) && allEmpty {
		position := positions[i]
		if (*warehouseMap)[position.Left][position.Right] != '.' {
			allEmpty = false
		}
		i++
	}

	return allEmpty
}
