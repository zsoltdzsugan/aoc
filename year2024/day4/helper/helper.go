package helper

import (
	"fmt"
	"os"
	"strings"
)

func ReadFileTo2D() [][]rune {
	input, err := os.ReadFile("../puzzle.input")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return [][]rune{}
	}

	strInput := strings.Split(strings.TrimSpace(string(input)), "\n")

	var array [][]rune

	for _, row := range strInput {
		var rowRunes []rune
		for _, col := range row {
			rowRunes = append(rowRunes, col)
		}
		array = append(array, rowRunes)
	}

	return array
}
