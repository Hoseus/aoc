package main

import (
	"aoc2024/utils"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	Coordinates utils.Tuple[int, int]
	Score       int
}

func main() {
	raceTrackMap, noCheatsPath := parseInput()
	part1(raceTrackMap, noCheatsPath)
	part2(raceTrackMap, noCheatsPath)
}

func parseInput() ([][]rune, []Position) {
	fileName := "./input"
	bytes, err := os.ReadFile(fileName) // Read the entire file
	if err != nil {
		panic(fmt.Sprintf("Failure reading file: %s", fileName))
	}
	lines := strings.Split(string(bytes), "\n")

	startPosition := utils.Tuple[int, int]{}
	endPosition := utils.Tuple[int, int]{}
	raceTrackMap := make([][]rune, len(lines))
	for i, row := range lines {
		for j, char := range row {
			if char == 'S' {
				startPosition = utils.Tuple[int, int]{Left: i, Right: j}
			}

			if char == 'E' {
				endPosition = utils.Tuple[int, int]{Left: i, Right: j}
			}
		}
		raceTrackMap[i] = []rune(row)
	}

	return raceTrackMap, findNoCheatsPath(&raceTrackMap, startPosition, endPosition)
}

func part1(raceTrackMap [][]rune, noCheatsPath []Position) {
	fmt.Println(fmt.Sprintf("Day 16 Part 1 Result: %d", len(findCheats(&raceTrackMap, &noCheatsPath, 2, 100))))
}

func part2(raceTrackMap [][]rune, noCheatsPath []Position) {
	fmt.Println(fmt.Sprintf("Day 16 Part 2 Result: %d", len(findCheats(&raceTrackMap, &noCheatsPath, 20, 100))))
}

func findNoCheatsPath(raceTrackMap *[][]rune, startCoordinates utils.Tuple[int, int], endCoordinates utils.Tuple[int, int]) []Position {
	derefRaceTrackMap := *raceTrackMap

	q := utils.NewDeque[Position]()
	visited := make(map[utils.Tuple[int, int]]bool)

	q.PushBack(Position{Coordinates: startCoordinates, Score: 0})
	visited[startCoordinates] = true

	path := make([]Position, 0)
	for q.Len() > 0 {
		currentPosition, _ := q.PopFront()
		currentCoordinates := currentPosition.Coordinates
		currentScore := currentPosition.Score

		path = append(path, currentPosition)

		if currentCoordinates == endCoordinates {
			break
		}

		for _, adjCoordinates := range getAdj(currentCoordinates, raceTrackMap) {
			if derefRaceTrackMap[adjCoordinates.Left][adjCoordinates.Right] != '#' && !visited[adjCoordinates] {
				adjPosition := Position{Coordinates: adjCoordinates, Score: currentScore + 1}
				q.PushBack(adjPosition)
				visited[adjCoordinates] = true
				break
			}
		}
	}

	return path
}

func findCheats(raceTrackMap *[][]rune, noCheatPath *[]Position, cheatDuration int, timeToSave int) []utils.Tuple[utils.Tuple[int, int], utils.Tuple[int, int]] {
	derefNoCheatPath := *noCheatPath

	noCheatCoordinatesScores := make(map[utils.Tuple[int, int]]int)
	for _, position := range derefNoCheatPath {
		noCheatCoordinatesScores[position.Coordinates] = position.Score
	}
	noCheatEndScore := derefNoCheatPath[len(derefNoCheatPath)-1].Score

	result := make([]utils.Tuple[utils.Tuple[int, int], utils.Tuple[int, int]], 0)
	for _, position := range derefNoCheatPath {
		currentCoordinates := position.Coordinates
		currentPossibleCheats := findCheatsFromPosition(raceTrackMap, position, cheatDuration)
		for _, possibleCheat := range currentPossibleCheats {
			endCheatCoordinates := possibleCheat.Right.Coordinates
			endCheatScore := possibleCheat.Right.Score
			if noCheatEndScore-(endCheatScore+(noCheatEndScore-noCheatCoordinatesScores[endCheatCoordinates])) >= timeToSave {
				result = append(result, utils.Tuple[utils.Tuple[int, int], utils.Tuple[int, int]]{Left: currentCoordinates, Right: endCheatCoordinates})
			}
		}
	}

	return result
}

func findCheatsFromPosition(raceTrackMap *[][]rune, position Position, cheatDuration int) []utils.Tuple[Position, Position] {
	derefRaceTrackMap := *raceTrackMap

	minHeap := utils.NewHeap[utils.Tuple[Position, int]](func(a, b utils.Tuple[Position, int]) bool {
		return a.Left.Score < b.Left.Score
	})
	scores := make(map[utils.Tuple[int, int]]int)

	scores[position.Coordinates] = position.Score
	minHeap.PushT(utils.Tuple[Position, int]{Left: position, Right: 0})

	result := make([]utils.Tuple[Position, Position], 0)
	for minHeap.Len() > 0 {
		currentCountedPosition := minHeap.PopT()
		currentPosition := currentCountedPosition.Left
		currentCoordinates := currentPosition.Coordinates
		currentValue := derefRaceTrackMap[currentCoordinates.Left][currentCoordinates.Right]
		currentScore := currentPosition.Score
		currentCheatCount := currentCountedPosition.Right

		if currentCheatCount > cheatDuration {
			continue
		}

		if currentValue != '#' {
			result = append(result, utils.Tuple[Position, Position]{Left: position, Right: currentPosition})
		}

		for _, adjCoordinates := range getAdj(currentCoordinates, raceTrackMap) {
			adjScore := currentScore + 1

			if _, exists := scores[adjCoordinates]; !exists {
				scores[adjCoordinates] = math.MaxInt
			}
			if adjScore < scores[adjCoordinates] {
				scores[adjCoordinates] = adjScore
				minHeap.PushT(utils.Tuple[Position, int]{Left: Position{Coordinates: adjCoordinates, Score: adjScore}, Right: currentCheatCount + 1})
			}
		}
	}

	return result
}

func getAdj(coordinates utils.Tuple[int, int], raceTrackMap *[][]rune) []utils.Tuple[int, int] {
	moves := []utils.Tuple[int, int]{
		{Left: -1, Right: 0},
		{Left: 0, Right: 1},
		{Left: 1, Right: 0},
		{Left: 0, Right: -1},
	}

	derefMap := *raceTrackMap
	rows := len(derefMap)
	cols := len(derefMap[0])

	result := make([]utils.Tuple[int, int], 0)
	for _, move := range moves {
		newCoordinates := utils.Tuple[int, int]{
			Left:  coordinates.Left + move.Left,
			Right: coordinates.Right + move.Right,
		}
		if newCoordinates.Left >= 0 && newCoordinates.Left < rows &&
			newCoordinates.Right >= 0 && newCoordinates.Right < cols {
			result = append(result, newCoordinates)
		}
	}

	return result
}

func printMap(raceTrackMap [][]rune, cheats []utils.Tuple[Position, Position]) {
	mapToPrint := make([][]rune, len(raceTrackMap))
	for i, row := range raceTrackMap {
		newRow := make([]rune, len(row))
		copy(newRow, row)
		mapToPrint[i] = newRow
	}

	fmt.Println()
	for _, cheat := range cheats {
		fromValue := mapToPrint[cheat.Left.Coordinates.Left][cheat.Left.Coordinates.Right]
		destValue := mapToPrint[cheat.Right.Coordinates.Left][cheat.Right.Coordinates.Right]
		if fromValue == 'S' || fromValue == 's' {
			mapToPrint[cheat.Left.Coordinates.Left][cheat.Left.Coordinates.Right] = 's'
		} else if fromValue == 'E' || fromValue == 'e' {
			mapToPrint[cheat.Left.Coordinates.Left][cheat.Left.Coordinates.Right] = 'e'
		} else {
			mapToPrint[cheat.Left.Coordinates.Left][cheat.Left.Coordinates.Right] = 'X'
		}

		if destValue == 'S' || destValue == '$' {
			mapToPrint[cheat.Right.Coordinates.Left][cheat.Right.Coordinates.Right] = '$'
		} else if destValue == 'E' || destValue == '3' {
			mapToPrint[cheat.Right.Coordinates.Left][cheat.Right.Coordinates.Right] = '3'
		} else {
			mapToPrint[cheat.Right.Coordinates.Left][cheat.Right.Coordinates.Right] = 'O'
		}
	}

	for _, row := range mapToPrint {
		fmt.Println(string(row))
	}
	fmt.Println()
}

func printMapWithScores(raceTrackMap [][]rune, noCheatScores map[utils.Tuple[int, int]]int, cheats []utils.Tuple[Position, Position]) {
	fmt.Println()
	mapToPrint := make([][]string, len(raceTrackMap))
	for i, row := range raceTrackMap {
		mapToPrint[i] = make([]string, len(row))
		for j, cell := range row {
			mapToPrint[i][j] = fmt.Sprintf("[%-11s]", string(cell))
		}
	}

	for _, cheat := range cheats {
		fromValue := raceTrackMap[cheat.Left.Coordinates.Left][cheat.Left.Coordinates.Right]
		fromScore := cheat.Left.Score
		destValue := raceTrackMap[cheat.Right.Coordinates.Left][cheat.Right.Coordinates.Right]
		destScore := cheat.Right.Score

		mapToPrint[cheat.Left.Coordinates.Left][cheat.Left.Coordinates.Right] = fmt.Sprintf("!%-3s:%-3s:%-3s!", string(fromValue), strconv.Itoa(noCheatScores[cheat.Left.Coordinates]), strconv.Itoa(fromScore))
		mapToPrint[cheat.Right.Coordinates.Left][cheat.Right.Coordinates.Right] = fmt.Sprintf("[%-3s:%-3s:%-3s]", string(destValue), strconv.Itoa(noCheatScores[cheat.Right.Coordinates]), strconv.Itoa(destScore))
	}

	for _, row := range mapToPrint {
		fmt.Println(row)
	}

	fmt.Println()
}

func printDetailedCheat(raceTrackMap [][]rune, noCheatScores map[utils.Tuple[int, int]]int, cheat utils.Tuple[Position, Position], totalTime, cheatTimeToDest, noCheatTimeToDest int) {
	fmt.Println()
	mapToPrint := make([][]string, len(raceTrackMap))
	for i, row := range raceTrackMap {
		mapToPrint[i] = make([]string, len(row))
		for j, cell := range row {
			mapToPrint[i][j] = fmt.Sprintf("[%-11s]", string(cell))
		}
	}

	fromValue := raceTrackMap[cheat.Left.Coordinates.Left][cheat.Left.Coordinates.Right]
	fromScore := cheat.Left.Score
	destValue := raceTrackMap[cheat.Right.Coordinates.Left][cheat.Right.Coordinates.Right]
	destScore := cheat.Right.Score

	mapToPrint[cheat.Left.Coordinates.Left][cheat.Left.Coordinates.Right] = fmt.Sprintf("!%-3s:%-3s:%-3s!", string(fromValue), strconv.Itoa(noCheatScores[cheat.Left.Coordinates]), strconv.Itoa(fromScore))
	mapToPrint[cheat.Right.Coordinates.Left][cheat.Right.Coordinates.Right] = fmt.Sprintf("$%-3s:%-3s:%-3s$", string(destValue), strconv.Itoa(noCheatScores[cheat.Right.Coordinates]), strconv.Itoa(destScore))

	for _, row := range mapToPrint {
		fmt.Println(row)
	}
	fmt.Println("Total:", totalTime)
	fmt.Println("No cheat time:", noCheatTimeToDest)
	fmt.Println("Cheat time:", cheatTimeToDest)
	fmt.Println(fmt.Sprintf("Saved time | %d - (%d + (%d - %d)): %d", totalTime, cheatTimeToDest, totalTime, noCheatTimeToDest, totalTime-(cheatTimeToDest+(totalTime-noCheatTimeToDest))))

	fmt.Println()
}
