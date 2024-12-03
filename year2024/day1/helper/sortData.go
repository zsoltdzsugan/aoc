package helper

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func SortData() (x, y []int) {
	data, err := os.ReadFile("../puzzle.input")
	if err != nil {
		fmt.Println("Couldn't open file:", err)
		return nil, nil
	}

	var leftList []int
	var rightList []int
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) < 2 {
			fmt.Println("Invalid line format:", line)
			continue
		}

		left, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error parsing left value:", err)
			continue
		}
		right, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("Error parsing right value:", err)
			continue
		}

		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	return leftList, rightList
}
