package day4

import (
	"fmt"

	"github.com/zsoltdzsugan/aoc/year2024/day4/helper"
)

func Main() {
	puzzle := helper.ReadFileTo2D()
	totalCount := countXMAS(puzzle)

	fmt.Println(totalCount)
}

func countXMAS(grid [][]rune) int {
	word := "MAS"
	wordReverse := "SAM"
	count := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 'A' {
				// Check top-left to bottom-right diagonal
				left := checkDiagonal(grid, i, j, -1, -1, 1, 1)
				// Check top-right to bottom-left diagonal
				right := checkDiagonal(grid, i, j, -1, 1, 1, -1)

				if (left == word || left == wordReverse) && (right == word || right == wordReverse) {
					count++
				}
			}
		}
	}
	return count
}

func checkDiagonal(grid [][]rune, i, j, dx1, dy1, dx2, dy2 int) string {
	isInBounds := func(x, y int) bool {
		return x >= 0 && y >= 0 && x < len(grid) && y < len(grid[0])
	}

	if isInBounds(i+dx1, j+dy1) && isInBounds(i+dx2, j+dy2) {
		top := grid[i+dx1][j+dy1]
		bottom := grid[i+dx2][j+dy2]
		return string(top) + string(grid[i][j]) + string(bottom)
	}
	return ""
}
