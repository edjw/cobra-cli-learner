/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/edjw/cobra-cli-learner/numberfunctions"
	"github.com/spf13/cobra"
)

// createLargeNumbersCmd represents the createLargeNumbers command
var createLargeNumbersCmd = &cobra.Command{
	Use:   "createLargeNumbers",
	Short: "Makes a very long list of very large numbers and copies them to your clipboard",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		number_list := numberfunctions.CreateListThousandLargeNumbers()

		// Copy to clipboard
		err := clipboard.WriteAll(number_list)
		if err != nil {
			fmt.Printf("Failed to copy to clipboard: %s", err)
			return
		}

		// No error, print success message
		fmt.Println("Random numbers successfully copied to clipboard!")
	},
}

func init() {
	rootCmd.AddCommand(createLargeNumbersCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createLargeNumbersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createLargeNumbersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
