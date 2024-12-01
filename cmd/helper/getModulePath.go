package helper

import (
	"fmt"
	"os"
	"strings"
)

func GetModulePath() string {
	goModPath := "go.mod"
	content, err := os.ReadFile(goModPath)
	if err != nil {
		fmt.Println("Error reading go.mod:", err)
		return ""
	}

	// Find the first line starting with "module "
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module "))
		}
	}
	return ""
}
