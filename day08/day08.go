package main

import (
	"bufio"
	"fmt"
	"os"
)

type position struct {
	row int
	col int
}

type bounds struct {
	maxRow int
	maxCol int
}

func parseInput(filePath string) (map[byte][]position, bounds, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, bounds{}, err
	}
	defer file.Close()

	allAntennas := make(map[byte][]position)
	scanner := bufio.NewScanner(file)
	row, col := 0, 0

	for scanner.Scan() {
		col = 0
		for _, value := range scanner.Bytes() {
			if value != '.' && value != '#' {
				allAntennas[value] = append(allAntennas[value], position{row, col})
			}
			col++
		}
		row++
	}
	bounds := bounds{maxRow: row - 1, maxCol: col - 1}
	return allAntennas, bounds, nil
}

func computeAntiNodesFromPair(pos1 position, pos2 position) (position, position) {
	posDiff := position{row: pos2.row - pos1.row, col: pos2.col - pos1.col}
	antiNode1 := position{row: pos2.row + posDiff.row, col: pos2.col + posDiff.col}
	antiNode2 := position{row: pos1.row - posDiff.row, col: pos1.col - posDiff.col}
	return antiNode1, antiNode2
}

func inBounds(pos position, bounds bounds) bool {
	return pos.row >= 0 && pos.row <= bounds.maxRow && pos.col >= 0 && pos.col <= bounds.maxCol
}

func computeAllAntiNodesWithinBounds(antennas []position, bounds bounds) []position {
	result := make([]position, 0)
	for i := 0; i < len(antennas); i++ {
		for j := i; j < len(antennas); j++ {
			if i != j {
				antiNode1, antiNode2 := computeAntiNodesFromPair(antennas[i], antennas[j])
				for _, antiNode := range [...]position{antiNode1, antiNode2} {
					if inBounds(antiNode, bounds) {
						result = append(result, antiNode)
					}
				}
			}
		}
	}
	return result
}

func part1(allAntennas map[byte][]position, bounds bounds) int {
	allAntiNodes := make(map[position]struct{})
	for _, antennas := range allAntennas {
		for _, antiNode := range computeAllAntiNodesWithinBounds(antennas, bounds) {
			allAntiNodes[antiNode] = struct{}{}
		}
	}
	return len(allAntiNodes)
}

func computeAntiNodesFromPairPt2(pos1 position, pos2 position, bounds bounds) []position {
	result := []position{pos1, pos2}
	posDiff := position{row: pos2.row - pos1.row, col: pos2.col - pos1.col}

	// Strategy as follows: first discover in one direction, then in the other. In this setup we need to manually add pos1 and pos2, which we did
	currPos := pos2
	for {
		currPos = position{row: currPos.row + posDiff.row, col: currPos.col + posDiff.col}
		if inBounds(currPos, bounds) {
			result = append(result, currPos)
		} else {
			break
		}
	}

	currPos = pos1
	for {
		currPos = position{row: currPos.row - posDiff.row, col: currPos.col - posDiff.col}
		if inBounds(currPos, bounds) {
			result = append(result, currPos)
		} else {
			break
		}
	}

	return result
}

func computeAllAntiNodesWithinBoundsPt2(antennas []position, bounds bounds) []position {
	result := make([]position, 0)
	for i := 0; i < len(antennas); i++ {
		for j := i; j < len(antennas); j++ {
			if i != j {
				antiNodes := computeAntiNodesFromPairPt2(antennas[i], antennas[j], bounds)
				for _, antiNode := range antiNodes {
					if inBounds(antiNode, bounds) {
						result = append(result, antiNode)
					}
				}
			}
		}
	}
	return result
}

func part2(allAntennas map[byte][]position, bounds bounds) int {
	allAntiNodes := make(map[position]struct{})
	for _, antennas := range allAntennas {
		for _, antiNode := range computeAllAntiNodesWithinBoundsPt2(antennas, bounds) {
			allAntiNodes[antiNode] = struct{}{}
		}
	}
	return len(allAntiNodes)
}

func main() {
	allAntennas, bounds, err := parseInput("input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("part1 = ", part1(allAntennas, bounds))
	fmt.Println("part2 = ", part2(allAntennas, bounds))
}
