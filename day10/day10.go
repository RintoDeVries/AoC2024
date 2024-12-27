package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func lineToSlice(line string) ([]int, error) {
	ints := make([]int, len(line))
	for i, s := range line {
		if s == '.' {
			ints[i] = -9
			continue
		}
		num, err := strconv.Atoi(string(s))
		if err != nil {
			return []int{}, err
		}
		ints[i] = num
	}
	return ints, nil
}

func parseInput(path string) ([][]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineSlice, err := lineToSlice(scanner.Text())
		if err != nil {
			return [][]int{}, err
		}
		result = append(result, lineSlice)
	}
	return result, scanner.Err()
}

func inBounds(m [][]int, row int, col int) bool {
	return row >= 0 && col >= 0 && row < len(m) && col < len(m[0])
}

func canMoveDown(m [][]int, row int, col int) bool {
	newRow, newCol := row-1, col
	if inBounds(m, newRow, newCol) {
		if m[newRow][newCol]-m[row][col] == 1 {
			return true
		}
	}
	return false
}

func canMoveUp(m [][]int, row int, col int) bool {
	newRow, newCol := row+1, col
	if inBounds(m, newRow, newCol) {
		if m[newRow][newCol]-m[row][col] == 1 {
			return true
		}
	}
	return false
}

func canMoveLeft(m [][]int, row int, col int) bool {
	newRow, newCol := row, col-1
	if inBounds(m, newRow, newCol) {
		if m[newRow][newCol]-m[row][col] == 1 {
			return true
		}
	}
	return false
}

func canMoveRight(m [][]int, row int, col int) bool {
	newRow, newCol := row, col+1
	if inBounds(m, newRow, newCol) {
		if m[newRow][newCol]-m[row][col] == 1 {
			return true
		}
	}
	return false
}

type position struct {
	row int
	col int
}

func uniqueEndPointsFromPosition(m [][]int, row int, col int, uniqueEndPoints map[position]struct{}) int {
	if m[row][col] == 9 {
		uniqueEndPoints[position{row, col}] = struct{}{}
	}
	if canMoveDown(m, row, col) {
		uniqueEndPointsFromPosition(m, row-1, col, uniqueEndPoints)
	}
	if canMoveUp(m, row, col) {
		uniqueEndPointsFromPosition(m, row+1, col, uniqueEndPoints)
	}
	if canMoveRight(m, row, col) {
		uniqueEndPointsFromPosition(m, row, col+1, uniqueEndPoints)
	}
	if canMoveLeft(m, row, col) {
		uniqueEndPointsFromPosition(m, row, col-1, uniqueEndPoints)
	}
	return len(uniqueEndPoints)
}

func uniqueRoutesFromPosition(m [][]int, row int, col int, numRoutes int) int {
	if m[row][col] == 9 {
		numRoutes++
	}

	if canMoveDown(m, row, col) {
		numRoutes = uniqueRoutesFromPosition(m, row-1, col, numRoutes)
	}
	if canMoveUp(m, row, col) {
		numRoutes = uniqueRoutesFromPosition(m, row+1, col, numRoutes)
	}
	if canMoveRight(m, row, col) {
		numRoutes = uniqueRoutesFromPosition(m, row, col+1, numRoutes)
	}
	if canMoveLeft(m, row, col) {
		numRoutes = uniqueRoutesFromPosition(m, row, col-1, numRoutes)
	}

	return numRoutes
}

func part1(m [][]int) int {
	result := 0
	for row := 0; row < len(m); row++ {
		for col := 0; col < len(m[0]); col++ {
			if m[row][col] == 0 {
				uniqueEndPoints := make(map[position]struct{})
				result += uniqueEndPointsFromPosition(m, row, col, uniqueEndPoints)
			}
		}
	}
	return result
}

func part2(m [][]int) int {
	result := 0
	for row := 0; row < len(m); row++ {
		for col := 0; col < len(m[0]); col++ {
			if m[row][col] == 0 {
				result += uniqueRoutesFromPosition(m, row, col, 0)
			}
		}
	}
	return result
}

func main() {
	m, err := parseInput("input.txt")
	if err != nil {
		panic(err)
	}
	start := time.Now()
	result := part1(m)
	elapsed := time.Since(start)
	fmt.Printf("Part 1 = %v, and it took %v\n", result, elapsed)

	start = time.Now()
	result = part2(m)
	elapsed = time.Since(start)
	fmt.Printf("Part 2 = %v, and it took %v\n", result, elapsed)

}
