package day3

import (
	"fmt"
	"regexp"

	"github.com/zsoltdzsugan/aoc/year2024/day3/helper"
)

func Main() {
	puzzle := helper.ReadFileToString("../puzzle.input")

	found := findNums(puzzle)
	sum := helper.CalNumsOfMul(found)

	fmt.Println(sum)
}

func findNums(puzzle string) []string {
	pattern := `(do\(\))|(don't\(\))|mul\([0-9]{1,3}\,[0-9]{1,3}\)`
	r := regexp.MustCompile(pattern)
	array := r.FindAllStringIndex(puzzle, -1)
	enabled := true
	var found []string
	for _, row := range array {
		substr := puzzle[row[0]:row[1]]
		if substr == "don't()" {
			enabled = false
		} else if substr == "do()" {
			enabled = true
		} else {
			if enabled {
				found = append(found, substr)
			}
		}
	}
	return found
}
