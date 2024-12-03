package helper

import (
	txtreader "github.com/zsoltdzsugan/aoc/txt_reader"
)

func GetRows() [][]int {
	data := txtreader.OpenFile("../puzzle.input")

	var result [][]int

	for _, row := range data {
		var intRow []int
		for _, value := range row {
			intRow = append(intRow, value.(int))
		}
		result = append(result, intRow)
	}
	return result
}
