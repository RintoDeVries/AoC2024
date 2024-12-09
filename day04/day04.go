package main

import (
	"bufio"
	"fmt"
	"os"
)

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

func searhPt1(input [][]byte, target []byte, row int, col int) int {
	return boolToInt(searchLeft(input, target, row, col)) +
		boolToInt(searchRight(input, target, row, col)) +
		boolToInt(searchUp(input, target, row, col)) +
		boolToInt(searchDown(input, target, row, col)) +
		boolToInt(searchTopLeft(input, target, row, col)) +
		boolToInt(searchTopRight(input, target, row, col)) +
		boolToInt(searchBottomLeft(input, target, row, col)) +
		boolToInt(searchBottomRight(input, target, row, col))
}

func part1(unpadded_input [][]byte) int {
	target := []byte("XMAS")
	padded_input := addPaddingToMatrix(unpadded_input, len(target))
	sum := 0
	for row := len(target); row < len(padded_input)-len(target); row++ {
		for col := len(target); col < len(padded_input[0])-len(target); col++ {
			sum += searhPt1(padded_input, target, row, col)
		}
	}
	return sum
}

func searchPt2(input [][]byte, target []byte, row int, col int) int {
	if (searchTopRight(input, target, row+1, col-1) || searchBottomLeft(input, target, row-1, col+1)) &&
		(searchTopLeft(input, target, row+1, col+1) || searchBottomRight(input, target, row-1, col-1)) {
		return 1
	} else {
		return 0
	}
}

func part2(unpadded_input [][]byte) int {
	target := []byte("SAM")
	padded_input := addPaddingToMatrix(unpadded_input, len(target))

	sum := 0
	for row := len(target); row < len(padded_input)-len(target); row++ {
		for col := len(target); col < len(padded_input[0])-len(target); col++ {
			sum += searchPt2(padded_input, target, row, col)
		}
	}
	return sum
}

func main() {
	unpadded_input, err := parseInput("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Answer to part1 = ", part1(unpadded_input))
	fmt.Println("Answer to part2 = ", part2(unpadded_input))
}
