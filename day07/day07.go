package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type equation struct {
	testValue int64
	nums      []int
}

func parseLine(line string) (equation, error) {
	testValueAndNumbers := strings.SplitN(line, ":", 2)
	if len(testValueAndNumbers) != 2 {
		return equation{}, fmt.Errorf("invalid line format, expected 'testValue: nums'")
	}

	testValue, err := strconv.ParseInt(strings.TrimSpace(testValueAndNumbers[0]), 10, 64)
	if err != nil {
		return equation{}, fmt.Errorf("invalid test value: %v", err)
	}

	numbersString := strings.TrimSpace(testValueAndNumbers[1])
	numberStrings := strings.Fields(numbersString)
	nums := make([]int, len(numberStrings))

	for i, numStr := range numberStrings {
		nums[i], err = strconv.Atoi(numStr)
		if err != nil {
			return equation{}, fmt.Errorf("invalid number '%s' at index %d: %v", numStr, i, err)
		}
	}

	return equation{
		testValue: int64(testValue),
		nums:      nums,
	}, nil
}

func parseInput(filepath string) ([]equation, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var equations []equation

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		equation, err := parseLine(scanner.Text())
		if err != nil {
			return nil, err
		}
		equations = append(equations, equation)
	}
	return equations, nil
}

func productOfSlice(nums []int) int64 {
	product := int64(1)
	for _, num := range nums {
		product *= int64(num)
	}
	return product
}

func isPossiblePt1(eq equation) bool {
	numPlaces := len(eq.nums) - 1
	numProductCombinations := 1 << numPlaces
	for attempt := range numProductCombinations {
		result := int64(-1)
		for place := range numPlaces {
			currBit := (attempt >> place) & 1
			if currBit == 0 {
				if result == -1 {
					result = int64(eq.nums[place]) + int64(eq.nums[place+1])
				} else {
					result += int64(eq.nums[place+1])
				}
			} else {
				if result == -1 {
					result = int64(eq.nums[place]) * int64(eq.nums[place+1])
				} else {
					result *= int64(eq.nums[place+1])
				}
			}
		}
		if result == eq.testValue {
			return true
		}
	}
	return false
}

func part1(eqs []equation) int64 {
	result := int64(0)
	for _, eq := range eqs {
		if isPossiblePt1(eq) {
			result += eq.testValue
		}
	}
	return result
}

func intPow(n, m int) int {
	if m == 0 {
		return 1
	}

	if m == 1 {
		return n
	}

	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

func concatInt64(a, b int64) int64 {
	strA := strconv.FormatInt(a, 10)
	strB := strconv.FormatInt(b, 10)
	concatStr := strA + strB
	result, err := strconv.ParseInt(concatStr, 10, 64)
	if err != nil {
		// Handle error, e.g., log or return a default value
		fmt.Println("Error converting concatenated string to int64:", err)
		return 0
	}
	return result
}

func isPossiblePt2(eq equation) bool {
	numPlaces := len(eq.nums) - 1
	numProductCombinations := intPow(3, numPlaces)
	for attempt := range numProductCombinations {
		result := int64(-1)
		for place := range numPlaces {
			currBit := (attempt / intPow(3, place)) % 3
			if currBit == 0 {
				if result == -1 {
					result = int64(eq.nums[place]) + int64(eq.nums[place+1])
				} else {
					result += int64(eq.nums[place+1])
				}
			} else if currBit == 1 {
				if result == -1 {
					result = int64(eq.nums[place]) * int64(eq.nums[place+1])
				} else {
					result *= int64(eq.nums[place+1])
				}
			} else {
				if result == -1 {
					result = concatInt64(int64(eq.nums[place]), int64(eq.nums[place+1]))
				} else {
					result = concatInt64(result, int64(eq.nums[place+1]))
				}
			}
		}
		if result == eq.testValue {
			return true
		}
	}
	return false
}

func part2(eqs []equation) int64 {
	result := int64(0)
	for _, eq := range eqs {
		if isPossiblePt2(eq) {
			result += eq.testValue
		}
	}
	return result
}

func main() {
	eqs, _ := parseInput("input.txt")
	fmt.Println(part1(eqs))
	fmt.Println(part2(eqs))
}
