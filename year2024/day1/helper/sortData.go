package helper

import (
	"slices"

	txtreader "github.com/zsoltdzsugan/aoc/txt_reader"
)

func SortData() (x, y []int) {
	data := txtreader.OpenFile("../puzzle.txt")

	var leftList []int
	var rightList []int
	for _, row := range data {
		leftList = append(leftList, row[0].(int))
		rightList = append(rightList, row[1].(int))
	}

	slices.Sort(leftList)
	slices.Sort(rightList)

	return leftList, rightList
}
