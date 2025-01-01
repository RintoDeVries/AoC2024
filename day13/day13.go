package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type position struct {
	x int
	y int
}

type machine struct {
	buttonA      position
	buttonB      position
	goalPosition position
}

func parseMachine(lines string) (machine, error) {
	pattern := `Button A: X\+(\d*), Y\+(\d*)\nButton B: X\+(\d*), Y\+(\d*)\nPrize: X=(\d*), Y=(\d*)`
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(lines)
	if len(matches) != 7 {
		return machine{}, fmt.Errorf("Something went wrong parsing")
	}

	intMatches := make([]int, len(matches)-1)
	for i, match := range matches[1:] {
		num, err := strconv.Atoi(match)
		if err != nil {
			return machine{}, err
		}
		intMatches[i] = num
	}

	return machine{buttonA: position{intMatches[0], intMatches[1]},
			buttonB:      position{intMatches[2], intMatches[3]},
			goalPosition: position{intMatches[4], intMatches[5]}},
		nil

}

func parseInput(filePath string) ([]machine, error) {
	machines := []machine{}

	file, err := os.Open(filePath)
	if err != nil {
		return machines, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var builder strings.Builder
	lineCounter := 0

	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		if lineCounter%3 != 0 {
			builder.WriteString("\n")
		}

		builder.WriteString(scanner.Text())

		lineCounter++
		if lineCounter%3 == 0 {
			machine, err := parseMachine(builder.String())
			builder.Reset()
			if err != nil {
				return machines, err
			}
			machines = append(machines, machine)
		}
	}
	return machines, nil
}

type state struct {
	currPosition position
	pressedA     int
	pressedB     int
}

func (s state) score() int {
	return s.pressedA*3 + s.pressedB
}

func (s state) step() state {
	return state{s.currPosition, s.pressedA, s.pressedB}
}

func (s state) buttonA(m machine) state {
	newPosition := position{x: s.currPosition.x + m.buttonA.x, y: s.currPosition.y + m.buttonA.y}
	return state{newPosition, s.pressedA + 1, s.pressedB}
}

func (s state) buttonB(m machine) state {
	newPosition := position{x: s.currPosition.x + m.buttonB.x, y: s.currPosition.y + m.buttonB.y}
	return state{newPosition, s.pressedA, s.pressedB + 1}
}

func bestSolutionDFS(machine machine, s state, bestScore int, mem map[state]int) int {
	if score, found := mem[s]; found {
		return score
	}

	if s.pressedA > 100 || s.pressedB > 100 || s.currPosition.x > machine.goalPosition.x || s.currPosition.y > machine.goalPosition.y {
		return 99999
	}

	if s.currPosition == machine.goalPosition {
		mem[s] = s.score()
		return s.score()
	}

	bestScore = min(bestScore, bestSolutionDFS(machine, s.step().buttonA(machine), bestScore, mem))
	bestScore = min(bestScore, bestSolutionDFS(machine, s.step().buttonB(machine), bestScore, mem))

	mem[s] = bestScore
	return bestScore
}

func part1(machines []machine) int {
	sum := 0
	for _, machine := range machines {
		currState := state{currPosition: position{0, 0}, pressedA: 0, pressedB: 0}
		mem := make(map[state]int)
		if result := bestSolutionDFS(machine, currState, 99999, mem); result != 99999 {
			sum += result
		}
	}
	return sum

}

func main() {
	machines, err := parseInput("input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(part1(machines))
}
