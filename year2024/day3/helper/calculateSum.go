package helper

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func ReadFileToString(filePath string) string {
	if input, err := os.ReadFile(filePath); err != nil {
		fmt.Println("error reading file:", err)
		return ""
	} else {
		return string(input)
	}
}

func CalNumsOfMul(array []string) int {
	pattern := `([0-9]{1,3})`
	r := regexp.MustCompile(pattern)
	sum := 0
	for _, row := range array {
		found := r.FindAllString(row, -1)
		x, err := strconv.Atoi(found[0])
		if err != nil {
			continue
		}
		y, err := strconv.Atoi(found[1])
		if err != nil {
			continue
		}
		sum += x * y
	}
	return sum
}
