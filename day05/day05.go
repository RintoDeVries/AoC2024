package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseRule(line string) ([2]int, error) {
	trimmed_line := strings.TrimSuffix(line, "\n")
	parts := strings.Split(trimmed_line, "|")
	if len(parts) < 2 {
		return [2]int{}, errors.New("parsing rule failed")
	}
	firstNum, firstErr := strconv.Atoi(parts[0])
	if firstErr != nil {
		return [2]int{}, firstErr
	}
	secondNum, secondErr := strconv.Atoi(parts[1])
	if secondErr != nil {
		return [2]int{}, secondErr
	}
	return [2]int{firstNum, secondNum}, nil
}

func csvLineToIntSlice(line string) ([]int, error) {
	trimmed_line := strings.TrimSuffix(line, "\n")
	parts := strings.Split(trimmed_line, ",")
	result := make([]int, len(parts))
	if len(parts) < 1 {
		return nil, errors.New("parsing failed")
	}
	for index, num := range parts {
		num, err := strconv.Atoi(num)
		if err != nil {
			return nil, err
		}
		result[index] = num
	}
	return result, nil
}

func parseInput(filepath string) ([][2]int, [][]int, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var finishedParsingRules = false
	var rules [][2]int
	var updates [][]int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			finishedParsingRules = true
			continue
		}
		if finishedParsingRules {
			update, err := csvLineToIntSlice(line)
			if err != nil {
				return nil, nil, err
			}
			updates = append(updates, update)
		} else {
			rule, err := parseRule(line)
			if err != nil {
				return nil, nil, err
			}
			rules = append(rules, rule)
		}
	}
	return rules, updates, nil
}

func findIndex(slice []int, target int) int {
	for i, num := range slice {
		if num == target {
			return i
		}
	}
	return -1
}

func isInCorrectOrder(rules [][2]int, update []int) bool {
	for _, rule := range rules {
		idx1 := findIndex(update, rule[0])
		idx2 := findIndex(update, rule[1])
		if idx1 != -1 && idx2 != -1 && idx1 > idx2 {
			return false
		}
	}
	return true
}

func part1(rules [][2]int, updates [][]int) int {
	sum := 0
	for _, update := range updates {
		if isInCorrectOrder(rules, update) {
			sum += update[len(update)/2]
		}
	}
	return sum
}

func orderUpdate(rules [][2]int, update []int) []int {
	// deep copy to not modify input
	result := make([]int, len(update))
	copy(result, update)
	for {
		numSwapsThisIteration := 0
		for _, rule := range rules {
			idx1 := findIndex(result, rule[0])
			idx2 := findIndex(result, rule[1])
			if idx1 != -1 && idx2 != -1 && idx1 > idx2 {
				result[idx1], result[idx2] = result[idx2], result[idx1]
				numSwapsThisIteration++
			}
		}
		if numSwapsThisIteration == 0 {
			return result
		}
	}
}

func part2(rules [][2]int, updates [][]int) int {
	sum := 0
	for _, update := range updates {
		if !isInCorrectOrder(rules, update) {
			sorted_update := orderUpdate(rules, update)
			sum += sorted_update[len(sorted_update)/2]
		}
	}
	return sum
}

func main() {
	rules, updates, err := parseInput("input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(part1(rules, updates))
	fmt.Println(part2(rules, updates))
}
