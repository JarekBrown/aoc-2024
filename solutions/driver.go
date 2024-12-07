package solutions

import (
	"aoc24/solutions/day01"
	"fmt"
	"os"
)

func Driver(day int, part int) {
	switch {
	case day == 1:
		Day1(part)
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
			data2, err := os.ReadFile("solutions/day01/part1.txt")
			if err != nil {
				panic(err)
			}
			answer = day01.PartTwo(data2)
			
	}
	fmt.Println(answer)
	// data1, err := os.ReadFile("solutions/day01/part1.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// data2, err := os.ReadFile("solutions/day01/part1.txt")
	// if err != nil {
	// 	panic(err)
	// }
}
