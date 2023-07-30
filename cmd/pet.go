/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// petCmd represents the pet command
var petCmd = &cobra.Command{
	Use:   "pet",
	Short: "Get pets based on status or details about a single pet by id.",
	Long:  `This command helps in getting pets based on`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pet called")
	},
}

func init() {
	getCmd.AddCommand(petCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// petCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// petCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
