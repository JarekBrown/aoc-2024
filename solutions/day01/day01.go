package day01

// https://adventofcode.com/2024/day/1

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func PartOne(data []byte) int {

	content := string(data)
	lines := strings.Fields(content)
	values, _ := StringsToInts(lines)

	a := make([]int, len(values)/2)
	b := make([]int, len(values)/2)
	j := 0
	for i := 0; i < len(values)-1; i += 2 {
		a[j] = values[i]
		b[j] = values[i+1]
		j++
	}
	sort.Ints(a)
	sort.Ints(b)

	output := 0
	for i := range a {
		output += AbsInt(a[i], b[i])
	}
	return output
}

func PartTwo(data []byte) int {
	content := string(data)
	lines := strings.Fields(content)
	values, _ := StringsToInts(lines)

	a := make([]int, len(values)/2)
	b := make([]int, len(values)/2)
	j := 0
	compare := make(map[int]int)
	for i := 0; i < len(values)-1; i += 2 {
		a[j] = values[i]
		b[j] = values[i+1]
		compare[b[j]]++
		j++
	}

	score := 0
	for _, val := range a {
		score += val * compare[val]
	}

	return score
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

func AbsInt(x int, y int) int {
	return int(math.Abs(float64(x) - float64(y)))
}
