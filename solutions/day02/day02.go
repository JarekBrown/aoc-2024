package day02

// https://adventofcode.com/2024/day/1

import (
	"strconv"
	"strings"
)

func PartOne(data []byte) int {
	output := 0

	content := string(data)
	lines := strings.Split(content, "\n")

	// go over each line
	for i := range len(lines) - 1 {
		values, _ := StringsToInts(strings.Fields(lines[i])) // convert to integers
		change := ""                                         // used to track if line increases or decreases
		safe := true

		// check each pair of values in the line
		for j := range len(values) - 1 {
			diff := values[j+1] - values[j]
			switch {
			case diff == 0: // when there is no change
				safe = false
			case diff < 0: // when the values are decreasing
				if change == "inc" {
					safe = false
					break
				} else if diff < -3 {
					safe = false
					break
				}
				change = "dec"
			case diff > 0: // when the values are increasing
				if change == "dec" {
					safe = false
					break
				} else if diff > 3 {
					safe = false
					break
				}
				change = "inc"
			}
			if !safe {
				break
			}
		}
		if safe {
			output++
		}
	}
	return output
}

func PartTwo(data []byte) int {
	output := 0

	content := string(data)
	lines := strings.Split(content, "\n")

	// go over each line
	for i := range len(lines) - 1 {
		values, _ := StringsToInts(strings.Fields(lines[i])) // convert to integers
		output += CheckLine(values, 0)
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

func CheckLine(input []int, counter int) int {
	safe := true
	change := ""

	// check each pair of values in the line
	for j := range len(input) - 1 {
		diff := input[j+1] - input[j]
		switch {
		case diff == 0: // when there is no change
			safe = false
		case diff < 0: // when the values are decreasing
			if change == "inc" {
				safe = false
				break
			} else if diff < -3 {
				safe = false
				break
			}
			change = "dec"
		case diff > 0: // when the values are increasing
			if change == "dec" {
				safe = false
				break
			} else if diff > 3 {
				safe = false
				break
			}
			change = "inc"
		}
		if !safe {
			break
		}
	}
	if safe {
		return 1
	} else {
		counter++
		if counter > 1 {
			return 0
		}
		// check if removing one value makes the report safe
		for index := range input {
			reduced := RemoveElement(input, index)
			val := CheckLine(reduced, counter)
			if val == 1 {
				return 1
			}
		}
		return 0
	}
}

func RemoveElement(input []int, index int) []int {
	length := len(input)
	if index == length-1 {
		return input[:index]
	} else {
		output := make([]int, 0)
		output = append(output, input[:index]...)
		return append(output, input[index+1:]...)
	}
}
