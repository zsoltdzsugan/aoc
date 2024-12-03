package day3

import (
	"fmt"
	"regexp"

	"github.com/zsoltdzsugan/aoc/year2024/day3/helper"
)

func Main() {
	puzzle := helper.ReadFileToString("../puzzle.input")

	pattern := `mul\([0-9]{1,3}\,[0-9]{1,3}\)`
	r := regexp.MustCompile(pattern)
	array := r.FindAllString(puzzle, -1)

	sum := helper.CalNumsOfMul(array)

	fmt.Println(sum)
}
