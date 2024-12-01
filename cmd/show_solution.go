package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/zsoltdzsugan/aoc/cmd/helper"
)

var showSolutionCmd = &cobra.Command{
	Use:   "solution",
	Short: "Shows solution",
	Long:  "Shows specified solution",
	Run:   showSolution,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		part, _ := cmd.Flags().GetInt("part")
		if year == 0 || day == 0 || part == 0 {
			return fmt.Errorf("Specify which solution you wish to view. Example: aoc solution -y 2024 -d 1 -p 1")
		}
		return nil
	},
}

func init() {
	showSolutionCmd.Flags().IntP("part", "p", 0, "part of the solution [1-2]")
	showSolutionCmd.MarkFlagRequired("part")

	rootCmd.AddCommand(showSolutionCmd)
}

func showSolution(cmd *cobra.Command, args []string) {
	part, _ := cmd.Flags().GetInt("part")

	// Build the solution file path
	yearDir := fmt.Sprintf("year%d", year)
	dayDir := fmt.Sprintf("day%d", day)
	partDir := fmt.Sprintf("part%d", part)
	fileName := fmt.Sprintf("solution.go")
	filePath := filepath.Join(yearDir, dayDir, partDir, fileName)

	// Check if the solution file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Printf("Solution file does not exist: %s\n", filePath)
		return
	}

	// Prepare the package import path
	modulePath := helper.GetModulePath() // Retrieves module name from go.mod
	importPath := fmt.Sprintf("%s/%s/%s/%s", modulePath, yearDir, dayDir, partDir)

	// Create a temporary main.go file to execute the solution
	tempMain := fmt.Sprintf(`
package main

import (
	solution "%s"
)

func main() {
	// Assuming the solution file has a Main() function
	solution.Main()
}`, importPath) // Strip .go extension for import

	tempFilePath := filepath.Join(os.TempDir(), "temp_main.go")
	err := os.WriteFile(tempFilePath, []byte(tempMain), 0644)
	if err != nil {
		fmt.Println("Error creating temporary main file:", err)
		return
	}
	defer os.Remove(tempFilePath) // Clean up after execution

	// Run the temporary main.go
	cmdExec := exec.Command("go", "run", tempFilePath)
	cmdExec.Dir = filepath.Dir(filePath) // Ensure proper context for relative paths

	// Capture the output
	output, err := cmdExec.CombinedOutput()
	if err != nil {
		fmt.Println("Error running the Go file:", err)
		fmt.Println(string(output)) // Show any output generated before the error
		return
	}

	// Print the output of the executed Go file
	fmt.Println(string(output))
}
