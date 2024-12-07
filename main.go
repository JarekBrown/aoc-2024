package main

import (
	"flag"
	"aoc24/solutions"
)

func main() {
	var day int
	flag.IntVar(&day, "d", 0, "which day of advent of code")

	var part int
	flag.IntVar(&part, "p", 0, "which part of the day's challenges")

	flag.Parse()

	solutions.Driver(day, part)
}
