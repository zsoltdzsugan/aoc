package txtreader

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func OpenFile(filepath string) [][]interface{} {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("error opening file:", err)
	}
	defer file.Close()

	var data [][]interface{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line) // split by spaces or tabs

		var lineData []interface{}
		for _, part := range parts {
			if num, err := strconv.Atoi(part); err == nil {
				lineData = append(lineData, num)
			} else {
				lineData = append(lineData, part)
			}
		}

		data = append(data, lineData)

		if err := scanner.Err(); err != nil {
			fmt.Println("error reading file:", err)
		}

	}

	return data
}
