package main

import (
	"bufio"
	"fmt"
	"os"
)

func searchLeft(input [][]byte, target []byte, row int, col int) bool {
	for i := 0; i < len(target); i++ {
		if input[row][col-i] != target[i] {
			return false
		}
	}
	return true
}

func searchRight(input [][]byte, target []byte, row int, col int) bool {
	for i := 0; i < len(target); i++ {
		if input[row][col+i] != target[i] {
			return false
		}
	}
	return true
}

func searchUp(input [][]byte, target []byte, row int, col int) bool {
	for i := 0; i < len(target); i++ {
		if input[row-i][col] != target[i] {
			return false
		}
	}
	return true
}

func searchDown(input [][]byte, target []byte, row int, col int) bool {
	for i := 0; i < len(target); i++ {
		if input[row+i][col] != target[i] {
			return false
		}
	}
	return true
}

func searchTopLeft(input [][]byte, target []byte, row int, col int) bool {
	for i := 0; i < len(target); i++ {
		if input[row-i][col-i] != target[i] {
			return false
		}
	}
	return true
}

func searchTopRight(input [][]byte, target []byte, row int, col int) bool {
	for i := 0; i < len(target); i++ {
		if input[row-i][col+i] != target[i] {
			return false
		}
	}
	return true
}

func searchBottomLeft(input [][]byte, target []byte, row int, col int) bool {
	for i := 0; i < len(target); i++ {
		if input[row+i][col-i] != target[i] {
			return false
		}
	}
	return true
}

func searchBottomRight(input [][]byte, target []byte, row int, col int) bool {
	for i := 0; i < len(target); i++ {
		if input[row+i][col+i] != target[i] {
			return false
		}
	}
	return true
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func search(input [][]byte, target []byte, row int, col int) int {
	return boolToInt(searchLeft(input, target, row, col)) +
		boolToInt(searchRight(input, target, row, col)) +
		boolToInt(searchUp(input, target, row, col)) +
		boolToInt(searchDown(input, target, row, col)) +
		boolToInt(searchTopLeft(input, target, row, col)) +
		boolToInt(searchTopRight(input, target, row, col)) +
		boolToInt(searchBottomLeft(input, target, row, col)) +
		boolToInt(searchBottomRight(input, target, row, col))
}

func parseInput(path string) ([][]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result [][]byte
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := append([]byte(nil), scanner.Bytes()...) // What the hell go... this is a very nasty bug...
		result = append(result, line)
	}

	return result, nil
}

func addPaddingToMatrix(unpadded [][]byte, padding int) [][]byte {
	numRows := len(unpadded) + 2*padding
	numCols := len(unpadded[0]) + 2*padding

	result := make([][]byte, numRows)
	for i := range result {
		result[i] = make([]byte, numCols)
	}

	for i := range unpadded {
		copy(result[i+padding][padding:], unpadded[i])
	}

	return result
}

func part1(path string) int {
	unpadded, err := parseInput(path)
	if err != nil {
		panic(err)
	}

	target := []byte("XMAS")
	padding := len(target)
	padded := addPaddingToMatrix(unpadded, padding)
	sum := 0
	for row := padding; row < len(unpadded)+padding; row++ {
		for col := padding; col < len(unpadded[0])+padding; col++ {
			sum += search(padded, target, row, col)
		}
	}
	return sum
}

func main() {
	fmt.Println(part1("input.txt"))
}
