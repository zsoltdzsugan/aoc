package part2

import (
	"fmt"

	"github.com/zsoltdzsugan/aoc/year2024/day1/helper"
)

func Main() {

	leftList, rightList := helper.SortData()

	freq := make(map[int]int)

	for _, num := range rightList {
		freq[num]++
	}

	similarityScore := 0
	for _, num := range leftList {
		if count, exists := freq[num]; exists {
			similarityScore += num * count
		}
	}

	fmt.Println(int(similarityScore))
}
