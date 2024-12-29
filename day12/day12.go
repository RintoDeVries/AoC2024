package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
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
		result = append(result, append([]byte{}, scanner.Bytes()...))
	}
	return result, scanner.Err()
}

type position struct {
	row int
	col int
}

type region struct {
	plots map[position]struct{}
	id    byte
}

func (r region) computeArea() int {
	return len(r.plots)
}

func inBounds(p position, g [][]byte) bool {
	return p.row >= 0 && p.col >= 0 && p.row < len(g) && p.col < len(g[0])
}

func neighboringPlots(p position, garden [][]byte) []position {
	candidateNeighbors := []position{
		{p.row - 1, p.col}, // Up
		{p.row + 1, p.col}, // Down
		{p.row, p.col - 1}, // Left
		{p.row, p.col + 1}, // Right
	}

	result := []position{}
	for _, cn := range candidateNeighbors {
		if inBounds(cn, garden) && garden[p.row][p.col] == garden[cn.row][cn.col] {
			result = append(result, cn)
		}
	}
	return result
}

func (r region) computePerimeter(garden [][]byte) int {
	perimeter := 0
	for p := range r.plots {
		perimeter += 4 - len(neighboringPlots(p, garden))
	}
	return perimeter
}

func computeCurrentRegion(p position, garden [][]byte) region {
	posQueue := []position{p}
	result := region{plots: make(map[position]struct{})}
	result.plots[p] = struct{}{}

	for len(posQueue) > 0 {
		neighboringPlots := neighboringPlots(posQueue[0], garden)
		for _, n := range neighboringPlots {
			_, alreadyAdded := result.plots[n]
			if !alreadyAdded {
				posQueue = append(posQueue, n)
				result.plots[n] = struct{}{}
			}
		}
		posQueue = posQueue[1:]
	}
	return result
}

func computeRegions(garden [][]byte) []region {
	regions := []region{}
	visited := make(map[position]struct{})

	for row := range garden {
		for col := range garden[0] {
			currPos := position{row, col}
			if _, exists := visited[currPos]; !exists {
				currPosRegion := computeCurrentRegion(currPos, garden)
				regions = append(regions, currPosRegion)

				for p := range currPosRegion.plots {
					visited[p] = struct{}{}
				}
			}
		}
	}
	return regions
}

func part1(garden [][]byte) int {
	result := 0
	for _, r := range computeRegions(garden) {
		result += r.computePerimeter(garden) * r.computeArea()
	}
	return result
}

func main() {
	garden, err := parseInput("input.txt")
	if err != nil {
		panic(err)
	}

	start := time.Now()
	result := part1(garden)
	elapsed := time.Since(start)
	fmt.Printf("Part 1 = %v, and it took %v\n", result, elapsed)
}
