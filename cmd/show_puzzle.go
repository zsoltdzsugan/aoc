package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var showPuzzleCmd = &cobra.Command{
	Use:   "puzzle",
	Short: "Show challenge puzzle",
	Long:  "Show challenge puzzle",
	Run:   showPuzzle,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if year == 0 || day == 0 {
			return fmt.Errorf("Specify which puzzle do you wish to view. Example: aoc puzzle -y 2024 -d 1")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(showPuzzleCmd)
}

func showPuzzle(cmd *cobra.Command, args []string) {
	fileName := fmt.Sprintf("puzzle.txt")
	filePath := filepath.Join(fmt.Sprintf("year%d", year), fmt.Sprintf("day%d", day), fileName)

	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("error reading file:", err)
		return
	}

	fmt.Println(string(content))
}
