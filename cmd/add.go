/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/edjw/cobra-cli-learner/numberfunctions"
	"github.com/spf13/cobra"
)

// Flags
var doubleResultFlag bool = false

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add numbers together",
	Long:    `Add numbers together`,
	Args:    cobra.MatchAll(cobra.MinimumNArgs(2)),
	Example: "add 1 2 3 4 5\nadd 1.2 2.3 3.4 4.5\nadd 1 2 3 4 5 --double\nadd 1.2 2.3 3.4 4.5 -d",
	RunE: func(cmd *cobra.Command, args []string) error {

		var sum float64

		for _, arg := range args {
			num, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				return fmt.Errorf("error converting argument to float: %v", err)
			}
			sum += num
		}

		if doubleResultFlag {
			sum *= 2
		}

		fmt.Printf("\nThe answer is: %v\n\n", numberfunctions.FormatLargeNumber(sum))

		// fmt.Printf("\nThe answer is: %v\n\n", sum)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().BoolVarP(&doubleResultFlag, "double", "d", false, "Double the result")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
