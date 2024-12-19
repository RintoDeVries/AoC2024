package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type position struct {
	row int
	col int
}

type bounds struct {
	maxRow int
	maxCol int
}

type direction struct {
	dRow int
	dCol int
}

type state struct {
	pos position
	dir direction
}

func parseInput(filepath string) (map[position]struct{}, bounds, state, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, bounds{}, state{}, err
	}
	defer file.Close()

	obstacles := make(map[position]struct{})
	var startState state

	scanner := bufio.NewScanner(file)
	row, col := 0, 0

	for scanner.Scan() {
		col = 0
		for _, value := range scanner.Bytes() {
			if value == '#' {
				obstacles[position{row, col}] = struct{}{}
			}
			if value == '^' {
				startState = state{position{row, col}, direction{-1, 0}}
			}

			col++
		}
		row++
	}
	bounds := bounds{maxRow: row - 1, maxCol: col - 1}
	return obstacles, bounds, startState, nil
}

func neighbor(curr_pos position, direction direction) position {
	return position{
		col: curr_pos.col + direction.dCol,
		row: curr_pos.row + direction.dRow,
	}
}

func inBounds(pos position, bounds bounds) bool {
	return pos.row >= 0 && pos.row <= bounds.maxRow && pos.col >= 0 && pos.col <= bounds.maxCol
}

func rotateRight(dir direction) direction {
	return direction{dRow: dir.dCol, dCol: -dir.dRow}
}

func simulate(curr_state state, obstacles map[position]struct{}, bounds bounds, visited map[state]struct{}) (state, map[state]struct{}, bool, bool) {
	next_pos := neighbor(curr_state.pos, curr_state.dir)
	leavesMap, containsLoop := false, false

	if !inBounds(next_pos, bounds) {
		leavesMap = true
		return curr_state, visited, leavesMap, containsLoop
	}

	if _, found := obstacles[next_pos]; found {
		return state{pos: curr_state.pos, dir: rotateRight(curr_state.dir)}, visited, leavesMap, containsLoop
	}

	next_state := state{pos: next_pos, dir: curr_state.dir}
	if _, found := visited[next_state]; found {
		containsLoop = true
		return curr_state, visited, leavesMap, containsLoop
	}
	visited[next_state] = struct{}{}
	return next_state, visited, leavesMap, containsLoop

}

func uniquePositions(visited map[state]struct{}) map[position]struct{} {
	unique := make(map[position]struct{})
	for s := range visited {
		unique[s.pos] = struct{}{}
	}
	return unique
}

func part1(startState state, obstacles map[position]struct{}, bounds bounds) map[position]struct{} {
	visited := make(map[state]struct{})
	currState := startState
	isDone := false
	for !isDone {
		currState, visited, isDone, _ = simulate(currState, obstacles, bounds, visited)
	}

	return uniquePositions(visited)
}

func part2(startState state, obstacles map[position]struct{}, bounds bounds, uniquePositions map[position]struct{}) int {
	result := 0

	for pos := range uniquePositions {
		currState := startState
		obstacles[pos] = struct{}{}
		visited := make(map[state]struct{})
		isDone, containsLoop := false, false
		for !(isDone || containsLoop) {
			currState, visited, isDone, containsLoop = simulate(currState, obstacles, bounds, visited)
			if isDone {
				continue
			}
			if containsLoop {
				result++
			}
		}
		delete(obstacles, pos)
	}

	return result

}

func main() {
	obstacles, bounds, startState, err := parseInput("input.txt")
	if err != nil {
		panic(err)
	}

	start := time.Now()
	uniquePositions := part1(startState, obstacles, bounds)
	elapsed := time.Since(start)
	fmt.Println("Result for pt1 =", len(uniquePositions), "and it took", elapsed, "bounds are", bounds)

	start = time.Now()
	pt2 := part2(startState, obstacles, bounds, uniquePositions)
	elapsed = time.Since(start)
	fmt.Println("Result for pt2 =", pt2, "and it took", elapsed, "bounds are", bounds)

}
