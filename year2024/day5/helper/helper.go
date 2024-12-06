package helper

import (
	"fmt"
	"os"
	"strings"
)

func ReadFileToArrays() ([][]string, [][]string) {
	input, err := os.ReadFile("../puzzle.input")
	if err != nil {
		fmt.Println("error reading the file")
	}
	strInput := strings.Split(string(input), "\n")
	isSplit := false
	var order [][]string
	var updates [][]string
	for _, lines := range strInput {
		if len(lines) == 0 {
			isSplit = true
		}
		if isSplit && len(lines) != 0 {
			updates = append(updates, strings.Split(lines, ","))
		}
		if !isSplit && len(lines) != 0 {
			order = append(order, strings.Split(lines, "|"))
		}
	}

	return order, updates
}
