package part1

import (
	"fmt"
	"math"

	"github.com/zsoltdzsugan/aoc/year2024/day1/helper"
)

func Main() {
	leftList, rightList := helper.SortData()

	totalDistance := 0.0
	for i := 0; i < len(leftList); i++ {
		totalDistance += math.Abs(float64(leftList[i]) - float64(rightList[i]))
	}

	//return int(totalDistance)
	fmt.Println(int(totalDistance))
}
