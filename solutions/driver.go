package solutions

import (
	"aoc24/solutions/day01"
	"aoc24/solutions/day02"
	"aoc24/solutions/day03"
	"fmt"
	"os"
)

func Driver(day int, part int) {
	switch {
	case day == 1:
		Day1(part)
	case day == 2:
		Day2(part)
	case day == 3:
		Day3(part)
	default:
		fmt.Println("No solution for that day and/or part")

	}
}

func Day1(part int) {
	var answer int
	switch part {
	case 1:
		data1, err := os.ReadFile("solutions/day01/part1.txt")
		if err != nil {
			panic(err)
		}
		answer = day01.PartOne(data1)
	case 2:
		data2, err := os.ReadFile("solutions/day01/part2.txt")
		if err != nil {
			panic(err)
		}
		answer = day01.PartTwo(data2)

	}
	fmt.Println(answer)
}

func Day2(part int) {
	var answer int
	switch part {
	case 1:
		data1, err := os.ReadFile("solutions/day02/part1.txt")
		if err != nil {
			panic(err)
		}
		answer = day02.PartOne(data1)
	case 2:
		data2, err := os.ReadFile("solutions/day02/part2.txt")
		if err != nil {
			panic(err)
		}
		answer = day02.PartTwo(data2)

	}
	fmt.Println(answer)
}

func Day3(part int) {
	var answer int
	switch part {
	case 1:
		data1, err := os.ReadFile("solutions/day03/part1.txt")
		if err != nil {
			panic(err)
		}
		answer = day03.PartOne(data1)
	case 2:
		data2, err := os.ReadFile("solutions/day03/part2.txt")
		if err != nil {
			panic(err)
		}
		answer = day03.PartTwo(data2)

	}
	fmt.Println(answer)
}
