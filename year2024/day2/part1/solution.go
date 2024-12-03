package main

import (
	"fmt"

	"github.com/zsoltdzsugan/aoc/year2024/day2/helper"
)

func main() {
	data := helper.GetRows()

	safeCount := countSafeReports(data)
	fmt.Printf("Number of safe reports: %d\n", safeCount)
}

func isSafe(report []int) bool {
	if len(report) < 2 {
		return false
	}

	increasing := true
	decreasing := true

	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]
		if diff < -3 || diff > 3 || diff == 0 {
			return false
		}
		if diff < 0 {
			increasing = false
		} else if diff > 0 {
			decreasing = false
		}
	}

	return increasing || decreasing
}

func countSafeReports(data [][]int) int {
	safeCount := 0

	for _, report := range data {
		if isSafe(report) {
			safeCount++
		}
	}

	return safeCount
}
