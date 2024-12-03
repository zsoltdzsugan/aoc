package day2

import (
	"fmt"

	"github.com/zsoltdzsugan/aoc/year2024/day2/helper"
)

func Main() {
	data := helper.GetRows()

	safeCount := countSafeReportsWithDampener(data)
	fmt.Printf("Number of safe reports: %d\n", safeCount)
}

// Checks if a report is safe using the original rules.
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

// Checks if removing a single level makes the report safe.
func isSafeWithDampener(report []int) bool {
	if isSafe(report) {
		return true
	}

	for i := 0; i < len(report); i++ {
		// Create a new slice excluding the i-th level
		modified := append([]int{}, report[:i]...)
		modified = append(modified, report[i+1:]...)
		if isSafe(modified) {
			return true
		}
	}

	return false
}

// Counts the number of reports that are safe, considering the Problem Dampener.
func countSafeReportsWithDampener(data [][]int) int {
	safeCount := 0

	for _, report := range data {
		if isSafeWithDampener(report) {
			safeCount++
		}
	}

	return safeCount
}
