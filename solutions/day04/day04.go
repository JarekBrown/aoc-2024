package day04

import (
	"fmt"
	"slices"
	"strings"
)

func PartOne(data []byte) int {
	output := 0

	content := string(data)
	lines := strings.Split(content, "\n")
	output += Horizontal(lines)
	output += Vertical(lines)
	output += Diagonal(lines)

	return output
}

func PartTwo(data []byte) int {
	output := 0

	content := string(data)
	lines := strings.Split(content, "\n")

	output = Cross(lines)
	return output
}

func Horizontal(lines []string) int {
	count := 0
	for _, line := range lines {
		count += strings.Count(line, "XMAS")
		count += strings.Count(line, "SAMX")
	}

	return count
}

func Vertical(lines []string) int {
	count := 0

	for i := range len(lines[0]) - 1 {
		column := ""
		for j := range len(lines) - 1 {
			column += string(lines[j][i])
		}
		count += strings.Count(column, "XMAS")
		count += strings.Count(column, "SAMX")
	}

	return count
}

func Diagonal(lines []string) int {
	count := 0
	x_max := len(lines[0]) - 1
	y_max := len(lines) - 1

	for i := range y_max {
		for j := range x_max {
			if j >= 3 && string(lines[i][j]) == "X" {
				if i >= 3 && string(lines[i-1][j-1]) == "M" {
					word := ""
					word += string(lines[i-1][j-1])
					word += string(lines[i-2][j-2])
					word += string(lines[i-3][j-3])
					if word == "MAS" {
						count++
					}
				}
				if i <= x_max-4 && string(lines[i+1][j-1]) == "M" {
					word := ""
					word += string(lines[i+1][j-1])
					word += string(lines[i+2][j-2])
					word += string(lines[i+3][j-3])
					if word == "MAS" {
						count++
					}
				}
			}
			if j <= y_max-3 && string(lines[i][j]) == "X" {
				if i >= 3 && string(lines[i-1][j+1]) == "M" {
					word := ""
					word += string(lines[i-1][j+1])
					word += string(lines[i-2][j+2])
					word += string(lines[i-3][j+3])
					if word == "MAS" {
						count++
					}
				}
				if i <= x_max-4 && string(lines[i+1][j+1]) == "M" {
					word := ""
					word += string(lines[i+1][j+1])
					word += string(lines[i+2][j+2])
					word += string(lines[i+3][j+3])
					if word == "MAS" {
						count++
					}
				}
			}
		}
	}
	return count
}

func Cross(lines []string) int {
	count := 0
	x_max := len(lines[0]) - 1
	y_max := len(lines) - 1
	var cross_points []string

	for i := range y_max {
		for j := range x_max {
			if j <= x_max-2 && (string(lines[i][j]) == "M" || string(lines[i][j]) == "S") {
				if i <= y_max-3 {
					word := string(lines[i][j])
					word += string(lines[i+1][j+1])
					word += string(lines[i+2][j+2])
					if word == "MAS" || word == "SAM" {
						location := fmt.Sprintf("(%d,%d)", i+1, j+1)
						cross_points = append(cross_points, location)
					}
				}
			}
		}
	}
	for i := range y_max {
		for j := range x_max {
			if j <= x_max-2 && (string(lines[i][j]) == "M" || string(lines[i][j]) == "S") {
				if i >= 2 {
					word := string(lines[i][j])
					word += string(lines[i-1][j+1])
					word += string(lines[i-2][j+2])
					if word == "MAS" || word == "SAM" {
						location := fmt.Sprintf("(%d,%d)", i-1, j+1)
						if slices.Contains(cross_points, location) {
							count++
						}
					}
				}
			}
		}
	}

	return count
}
