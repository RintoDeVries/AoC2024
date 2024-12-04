package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func lineToSlice(line string) ([]int, error) {
	parts := strings.Split(line, " ")

	if len(parts) == 0 {
		return []int{}, errors.New("couldn't parse line")
	}
	ints := make([]int, len(parts))
	for i, s := range parts {
		num, err := strconv.Atoi(s)
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

// yes, golang does not have abs for ints...
func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func isSafePt1(report []int) bool {
	isIncreasing, isDecreasing, isDistOk := true, true, true
	prevIndex := 0
	for index, _ := range report {
		if index != 0 {
			diff := report[index] - report[prevIndex]
			if diff > 0 {
				isIncreasing = false
			}
			if diff < 0 {
				isDecreasing = false
			}
			if absInt(diff) > 3 || absInt(diff) < 1 {
				isDistOk = false
			}
		}
		if !isIncreasing && !isDecreasing || !isDistOk {
			return false
		}
		prevIndex = index
	}
	return true
}

func isSafePt2(report []int) bool {
	if isSafePt1(report) {
		return true
	} else { // ... 100% can be done in a single pass, but iT wOrKs
		for i := range report {
			report_without_level_i := make([]int, len(report))
			copy(report_without_level_i, report)
			report_without_level_i = append(report_without_level_i[:i], report_without_level_i[i+1:]...)
			if isSafePt1(report_without_level_i) {
				return true
			}
		}
	}
	return false
}

func main() {
	reports, err := parseInput("input.txt")
	if err != nil {
		panic(err)
	}
	result_pt1, result_pt2 := 0, 0
	for _, report := range reports {
		if isSafePt1(report) {
			result_pt1 += 1
		}
		if isSafePt2(report) {
			result_pt2 += 1
		}
	}
	fmt.Println("result pt1 = ", result_pt1)
	fmt.Println("result pt2 = ", result_pt2)
}
