package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var showChallengeCmd = &cobra.Command{
	Use:   "challenge",
	Short: "Shows challenge",
	Long:  "Shows specified challenge's description with an example",
	Run:   showChallenge,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if year == 0 || day == 0 || part == 0 {
			return fmt.Errorf("Specify which challenge you wish to view. Example: aoc challenge -y 2024 -d 1 -p 1")
		}
		return nil
	},
}

func init() {
	rootCmd.MarkFlagRequired("part")
	rootCmd.AddCommand(showChallengeCmd)
}

func showChallenge(cmd *cobra.Command, args []string) {
	fileName := fmt.Sprintf("challenge.md")
	filePath := filepath.Join(fmt.Sprintf("year%d", year), fmt.Sprintf("day%d", day), fmt.Sprintf("part%d", part), fileName)

	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println(string(content))
}
