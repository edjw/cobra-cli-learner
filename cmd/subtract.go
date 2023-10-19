/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// Flags
var roundResultFlag bool = false

// subtractCmd represents the subtract command
var subtractCmd = &cobra.Command{
	Use:     "subtract",
	Short:   "Subtract numbers",
	Long:    `Takes a list of numbers and subtracts one from the previous one, one after another`,
	Args:    cobra.MatchAll(cobra.MinimumNArgs(2)),
	Example: "subtract 5 4 3 2 1\nsubtract 5.2 3.3 3 1\n",
	RunE: func(cmd *cobra.Command, args []string) error {

		var sum float64

		for index, arg := range args {
			num, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				return fmt.Errorf("error converting argument to float: %v", err)
			}

			if index == 0 {
				sum = num // Set sum to the first item in the args slice
			} else {
				sum -= num
			}

		}

		if roundResultFlag {
			fmt.Printf("\nThe answer is: %.2f\n\n", sum)
		} else {
			fmt.Printf("\nThe answer is: %v\n\n", sum)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(subtractCmd)

	subtractCmd.Flags().BoolVarP(&roundResultFlag, "round", "r", false, "Round the result")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// subtractCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// subtractCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
