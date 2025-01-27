package main

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"
)

func main() {
	computers, computerGraph := parseInput()
	part1(computers, computerGraph)
	part2(computers, computerGraph)
}

func parseInput() ([]string, map[string][]string) {
	fileName := "./input"
	bytes, err := os.ReadFile(fileName) // Read the entire file
	if err != nil {
		panic(fmt.Sprintf("Failure reading file: %s", fileName))
	}
	lines := strings.Split(string(bytes), "\n")

	graphSet := make(map[string]map[string]bool)
	for _, line := range lines {
		computerPair := strings.Split(line, "-")
		if _, exists := graphSet[computerPair[0]]; !exists {
			graphSet[computerPair[0]] = make(map[string]bool)
		}
		graphSet[computerPair[0]][computerPair[1]] = true
		if _, exists := graphSet[computerPair[1]]; !exists {
			graphSet[computerPair[1]] = make(map[string]bool)
		}
		graphSet[computerPair[1]][computerPair[0]] = true
	}

	computers := make([]string, 0)
	computerGraph := make(map[string][]string)
	for computer, connectedComputers := range graphSet {
		computers = append(computers, computer)
		for connectedComputer := range connectedComputers {
			computerGraph[computer] = append(computerGraph[computer], connectedComputer)
		}
	}

	return computers, computerGraph
}

func part1(computers []string, computerGraph map[string][]string) {
	computerTrios := findCliques(computers, computerGraph, 3)

	result := 0
	for _, computerTrio := range computerTrios {
		hasComputerStartingWithT := false
		for _, computer := range computerTrio {
			if strings.HasPrefix(computer, "t") {
				hasComputerStartingWithT = true
				break
			}
		}
		if hasComputerStartingWithT {
			result++
		}
	}
	fmt.Println(fmt.Sprintf("Day 23 Part 1 Result: %d", result))
}

func part2(computers []string, computerGraph map[string][]string) {
	maxClique := findMaxClique(computers, computerGraph)
	sort.Slice(maxClique, func(i, j int) bool {
		return maxClique[i] < maxClique[j]
	})
	fmt.Println(fmt.Sprintf("Day 23 Part 2 Result: %s", strings.Join(maxClique, ",")))
}

func findCliques(nodes []string, graph map[string][]string, k int) [][]string {
	stack := utils.NewDeque[[]string]()
	for _, node := range nodes {
		stack.PushBack([]string{node})
	}

	result := make([][]string, 0)
	for stack.Len() > 0 {
		currentClique, _ := stack.PopBack()
		currentNode := currentClique[len(currentClique)-1]

		if len(currentClique) == k && isClique(graph, currentClique) {
			resultClique := make([]string, len(currentClique))
			copy(resultClique, currentClique)
			result = append(result, resultClique)
			continue
		}

		for _, connectedNode := range graph[currentNode] {
			if connectedNode > currentNode && !slices.Contains(currentClique, connectedNode) {
				newClique := make([]string, len(currentClique)+1)
				copy(newClique, currentClique)
				newClique[len(newClique)-1] = connectedNode
				stack.PushBack(newClique)
			}
		}
	}
	return result
}

func findMaxClique(nodes []string, graph map[string][]string) []string {
	stack := utils.NewDeque[[]string]()
	for _, node := range nodes {
		stack.PushBack([]string{node})
	}

	result := make([]string, 0)
	for stack.Len() > 0 {
		currentClique, _ := stack.PopBack()
		currentNode := currentClique[len(currentClique)-1]

		if len(currentClique) > len(result) && isClique(graph, currentClique) {
			resultClique := make([]string, len(currentClique))
			copy(resultClique, currentClique)
			result = resultClique
			continue
		}

		for _, connectedNode := range graph[currentNode] {
			if connectedNode > currentNode && !slices.Contains(currentClique, connectedNode) {
				newClique := make([]string, len(currentClique)+1)
				copy(newClique, currentClique)
				newClique[len(newClique)-1] = connectedNode
				stack.PushBack(newClique)
			}
		}
	}
	return result
}

func isClique(graph map[string][]string, nodes []string) bool {
	size := len(nodes)
	for i := 0; i < size; i++ {
		node := nodes[i]
		for j := i + 1; j < i+size; j++ {
			nextNode := nodes[(j%size+size)%size]
			if !slices.Contains(graph[node], nextNode) {
				return false
			}
		}
	}
	return true
}
