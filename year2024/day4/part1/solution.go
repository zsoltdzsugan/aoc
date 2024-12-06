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
	directions := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
		{1, 1},
		{-1, -1},
		{-1, 1},
		{1, -1},
	}

	word := "XMAS"
	wordLength := len(word)
	count := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			for _, dir := range directions {
				match := true
				for k := 0; k < wordLength; k++ {
					newX := i + k*dir[0]
					newY := j + k*dir[1]
					if !isInBounds(grid, newX, newY) || grid[newX][newY] != rune(word[k]) {
						match = false
						break
					}
				}
				if match {
					count++
				}
			}
		}
	}
	return count
}

func isInBounds(grid [][]rune, x, y int) bool {
	return x >= 0 && y >= 0 && x < len(grid) && y < len(grid[0])
}
