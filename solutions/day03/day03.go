package day03

import (
	"regexp"
	"strconv"
	"strings"
)

func PartOne(data []byte) int {
	output := 0

	content := string(data)

	re := regexp.MustCompile(`mul\(\d+\,\d+\)`)
	matches := re.FindAllString(content, -1)
	for _, match := range matches {
		output += ParseMatch(match)
	}
	return output
}

func PartTwo(data []byte) int {
	output := 0

	content := string(data)

	re := regexp.MustCompile(`mul\(\d+\,\d+\)|do\(\)|don't\(\)`)
	matches := re.FindAllString(content, -1)
	use := true

	for _, match := range matches {
		if strings.Contains(match, "don't()") {
			use = false
		} else if strings.Contains(match, "do()") {
			use = true
		} else if use {
			output += ParseMatch(match)
		}
	}

	return output
}

func StringsToInts(input []string) ([]int, error) {
	output := make([]int, len(input))
	for i, str := range input {
		val, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		output[i] = val
	}
	return output, nil
}

func ParseMatch(input string) int {
	preRem, _ := strings.CutPrefix(input, "mul(")
	suffRem, _ := strings.CutSuffix(preRem, ")")
	values := strings.Split(suffRem, ",")
	x, _ := strconv.Atoi(values[0])
	y, _ := strconv.Atoi(values[1])
	return x * y
}
