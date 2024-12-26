package main

import (
	"fmt"
	"os"
	"time"
)

func parseInput(input string) []int {
	var result []int
	isEmptyDiskSpace := false
	for stringIndex, char := range input {
		if !isEmptyDiskSpace {
			num := int(char - '0')
			for i := 0; i < num; i++ {
				result = append(result, stringIndex/2) // the id
			}
			isEmptyDiskSpace = true
		} else {
			num := int(char - '0')
			for i := 0; i < num; i++ {
				result = append(result, -1)
			}
			isEmptyDiskSpace = false
		}
	}
	return result
}

func findLeftMostEmpty(fileSystem []int, start int) int {
	for i := start; i < len(fileSystem); i++ {
		if fileSystem[i] == -1 {
			return i
		}
	}
	panic("this may never happen)")
}

func findRightMostFile(fileSystem []int, start int) int {
	for i := start; i > 0; i-- {
		if fileSystem[i] != -1 {
			return i
		}
	}
	panic("this may never happen")
}

func compactFilePt1(fs []int) []int {
	fileSystem := make([]int, len(fs))
	copy(fileSystem, fs)

	leftMostEmpty := 0
	rightMostFile := len(fileSystem) - 1

	for leftMostEmpty < rightMostFile {
		leftMostEmpty = findLeftMostEmpty(fileSystem, leftMostEmpty)
		rightMostFile = findRightMostFile(fileSystem, rightMostFile)
		if leftMostEmpty > rightMostFile {
			break
		}
		fileSystem[leftMostEmpty], fileSystem[rightMostFile] = fileSystem[rightMostFile], fileSystem[leftMostEmpty]
	}
	return fileSystem
}

func endOfCurrentFileRight(fileSystem []int, index int) int {
	value := fileSystem[index]
	if value != -1 {
		panic("this should never happen")
	}
	for i := index; i < len(fileSystem); i++ {
		if fileSystem[i] != value {
			return i - 1
		}
	}
	return -1
}

func endOfCurrentFileLeft(fileSystem []int, index int) int {
	value := fileSystem[index]
	for i := index; i >= 0; i-- {
		if fileSystem[i] != value {
			return i + 1
		}
	}
	return -1
}

func compactFilePt2(fs []int) []int {
	// deep copy to prevent modifying input
	fileSystem := make([]int, len(fs))
	copy(fileSystem, fs)

	// loop from left to right
	rightIndexCurrFile := findRightMostFile(fileSystem, len(fileSystem)-1)
	leftIndexCurrFile := endOfCurrentFileLeft(fileSystem, rightIndexCurrFile)

	for leftIndexCurrFile > 0 {
		for i := 0; i < leftIndexCurrFile; i++ {
			if fileSystem[i] == -1 {
				sizeAvailable := endOfCurrentFileRight(fileSystem, i) - i + 1
				currFileSize := rightIndexCurrFile - leftIndexCurrFile + 1
				if sizeAvailable >= (rightIndexCurrFile - leftIndexCurrFile + 1) {
					for j := i; j < i+currFileSize; j++ {
						fileSystem[j] = fileSystem[rightIndexCurrFile]
					}
					for j := leftIndexCurrFile; j <= rightIndexCurrFile; j++ {
						fileSystem[j] = -1
					}
					break
				}
			}
		}

		rightIndexCurrFile = findRightMostFile(fileSystem, leftIndexCurrFile-1)
		leftIndexCurrFile = endOfCurrentFileLeft(fileSystem, rightIndexCurrFile)
	}
	return fileSystem

}

func computeChecksum(fileSystem []int) int {
	result := 0
	for i := 0; i < len(fileSystem); i++ {
		if fileSystem[i] == -1 {
			continue
		}
		result += fileSystem[i] * i
	}
	return result
}

func part1(input []int) int {
	compactInput := compactFilePt1(input)
	return computeChecksum(compactInput)
}

func part2(input []int) int {
	compactInput := compactFilePt2(input)
	return computeChecksum(compactInput)
}

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := parseInput(string(b))

	start := time.Now()
	result := part1(input)
	elapsed := time.Since(start)
	fmt.Printf("Part 1 = %v, and it took %v\n", result, elapsed)

	start = time.Now()
	result = part2(input)
	elapsed = time.Since(start)
	fmt.Printf("Part 2 = %v, and it took %v\n", result, elapsed)
}
