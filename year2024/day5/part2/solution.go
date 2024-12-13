package day5

import (
	"fmt"
	"strconv"

	"github.com/zsoltdzsugan/aoc/year2024/day5/helper"
)

func Main() {
	orders, updates := helper.ReadFileToArrays()

	result := 0
	for i := range updates {
		if !isOrdered(orders, updates[i]) {
			middleIdx := len(updates[i]) / 2
			if n, err := strconv.Atoi(updates[i][middleIdx]); err == nil {
				result += n
			}
		}
	}
	fmt.Println(result)
}

func isOrdered(orders [][]string, update []string) bool {
	for _, order := range orders {
		l, r := order[0], order[1]
		lIdx, rIdx := -1, -1

		for i, u := range update {
			if u == l {
				lIdx = i
			}
			if u == r {
				rIdx = i
			}
		}

		if lIdx != -1 && rIdx != -1 && lIdx > rIdx {
			orderIncorrect(orders, update)
			return false
		}
	}
	return true
}

func orderIncorrect(orders [][]string, update []string) {

	for {
		swap := false
		for _, order := range orders {
			l, r := order[0], order[1]
			for i := 0; i < len(update); i++ {
				for j := i + 1; j < len(update); j++ {
					if update[i] == r && update[j] == l {
						update[i], update[j] = update[j], update[i]
						swap = true
					}
				}
			}
		}
		if !swap {
			break
		}
	}
}
