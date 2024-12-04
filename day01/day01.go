package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func lineToNums(line string) (int, int, error) {
	parts := strings.Split(line, "   ")
	if len(parts) < 2 {
		return 0, 0, errors.New("input string must contain at least two characters")
	}
	leftNum, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, err
	}
	rightNum, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, err
	}
	return leftNum, rightNum, nil
}

// yes, golang does not have abs for ints...
func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func findDistancePart1(leftList []int, rightList []int) int {
	sort.Ints(leftList)
	sort.Ints(rightList)

	var result int = 0
	for index := range len(leftList) {
		result += absInt(leftList[index] - rightList[index])
	}
	return result
}

func countOccurrences(slice []int, target int) int {
	count := 0
	for _, value := range slice {
		if value == target {
			count++
		}
	}
	return count
}

func findDistancePart2(leftList []int, rightList []int) int {
	var result int = 0

	for index := range len(leftList) {
		leftNum := leftList[index]
		occurences := countOccurrences(rightList, leftNum)
		result += leftNum * occurences
	}
	return result

}

func run(path string) {
	lines, err := readLines(path)
	if err != nil {
		panic(err)
	}

	leftList, rightList := make([]int, len(lines)), make([]int, len(lines))
	for index, line := range lines {
		leftNum, rightNum, err := lineToNums(line)
		if err != nil {
			panic(err)
		}
		leftList[index] = leftNum
		rightList[index] = rightNum
	}

	fmt.Println("Answer pt1= ", findDistancePart1(leftList, rightList))
	fmt.Println("Answer pt2= ", findDistancePart2(leftList, rightList))
}

func main() {
	start := time.Now()
	run("input.txt")
	fmt.Println("Execution took ", time.Now().Sub(start))
}
