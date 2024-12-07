package day01

import (
	"os"
	"testing"
)

func TestPartOne(t *testing.T) {
	data1, err := os.ReadFile("test1.txt")
	if err != nil {
		panic(err)
	}
	answer := PartOne(data1)
	want := 11
	if want != answer {
		t.Error("Got:", answer, "\nWanted:", want)
	}
}

func TestPartTwo(t *testing.T) {
	data1, err := os.ReadFile("test2.txt")
	if err != nil {
		panic(err)
	}
	answer := PartTwo(data1)
	want := 31
	if want != answer {
		t.Error("Got:", answer, "\nWanted:", want)
	}
}
