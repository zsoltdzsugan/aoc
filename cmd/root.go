package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var year int
var day int

var rootCmd = &cobra.Command{
	Use:   "aoc",
	Short: "Advent of Code CLI",
	Long: `A CLI for Advent of Code challenges.
	It contains the challenges, the puzzle inputs and my solutions (if i managed to solve it :))`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&year, "year", "y", 0, "Specify the year of the challenge [2024]")
	rootCmd.PersistentFlags().IntVarP(&day, "day", "d", 0, "Specify the day of the challenge [1-24]")

	rootCmd.MarkFlagRequired("year")
	rootCmd.MarkFlagRequired("day")
}