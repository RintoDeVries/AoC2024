package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func parseInput(filePath string) (map[int]int, error) {
	b, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	result := make(map[int]int)
	for _, strNum := range strings.Fields(string(b)) {
		num, err := strconv.Atoi(strNum)
		if err != nil {
			return result, nil
		}
		result[num] += 1
	}

	return result, nil
}

func calculateStoneProgression(prevStones map[int]int) map[int]int {
	nextStones := make(map[int]int)
	for oldValue, numStones := range prevStones {
		if oldValue == 0 {
			newValue := 1
			nextStones[newValue] += numStones
		} else if oldValuStr := strconv.Itoa(oldValue); len(oldValuStr)%2 == 0 {
			num1, _ := strconv.Atoi(oldValuStr[:len(oldValuStr)/2])
			num2, _ := strconv.Atoi(oldValuStr[len(oldValuStr)/2:])
			nextStones[num1] += numStones
			nextStones[num2] += numStones
		} else {
			newValue := oldValue * 2024
			nextStones[newValue] += numStones
		}
	}

	return nextStones
}

func numStones(stones map[int]int) int {
	result := 0
	for _, v := range stones {
		result += v
	}
	return result
}

func part1(input map[int]int, numRounds int) int {
	next := input
	for range numRounds {
		next = calculateStoneProgression(next)
	}
	return numStones(next)
}

func main() {
	input, err := parseInput("input.txt")
	if err != nil {
		panic(err)
	}

	start := time.Now()
	result := part1(input, 25)
	elapsed := time.Since(start)
	fmt.Printf("Part 1 = %v, and it took %v\n", result, elapsed)

	start = time.Now()
	result = part1(input, 75)
	elapsed = time.Since(start)
	fmt.Printf("Part 2 = %v, and it took %v\n", result, elapsed)
}
