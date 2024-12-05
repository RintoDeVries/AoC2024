package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parseInput(file string) (string, error) {
	buf, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

func textToAnswer(text string) int {
	pattern := `mul\((\d{1,3}),(\d{1,3})\)`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(text, -1)
	result := 0
	for _, match := range matches {
		if len(match) > 2 {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
			result += num1 * num2
		}
	}
	return result
}

func filterPt2(text string) string {
	enabled := true
	var result strings.Builder
	i := 0
	for i < len(text) {
		if i+6 < len(text) && text[i:i+7] == "don't()" {
			if enabled {
				enabled = false
			}
			i += 6
			continue
		}

		if i+4 < len(text) && text[i:i+4] == "do()" {
			if !enabled {
				enabled = true
			}
			i += 4
			continue
		}
		if enabled {
			result.WriteByte(text[i])
		}
		i++
	}
	return result.String()
}

func main() {
	text, err := parseInput("input.txt")
	if err != nil {
		panic(err)
	}
	// part 1
	fmt.Println("Result for pt1 = ", textToAnswer(text))

	// part 2
	fmt.Println("Result for pt2 = ", textToAnswer(filterPt2(text)))
}
