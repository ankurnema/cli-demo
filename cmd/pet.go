/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"github.com/ankurnema/cli-demo/pkg/petstore"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"strings"
)

var status string

var (
	colPetId   = "Pet Id"
	colPetName = "Pet Name"
	rowHeader  = table.Row{colPetId, colPetName}
)

// petCmd represents the pet command
var petCmd = &cobra.Command{
	Use:   "pet",
	Short: "Get pets based on status or details about a single pet by id.",
	Long:  `This command helps in getting pets based on`,
	Run: func(cmd *cobra.Command, args []string) {
		// Business logic

		params := petstore.FindPetsByStatusParams{
			Status: []petstore.FindPetsByStatusParamsStatus{petstore.FindPetsByStatusParamsStatus(status)},
		}

		petResponse, err := petStoreClient.FindPetsByStatusWithResponse(context.Background(), &params)
		if err != nil {
			panic(err)
		}
		pets := *petResponse.JSON200

		// We got pets. Lets iterate over and print value using table

		tw := table.NewWriter()
		tw.AppendHeader(rowHeader)
		var rows []table.Row
		for _, pet := range pets {
			rows = append(rows, table.Row{*pet.Id, pet.Name})
		}

		tw.AppendRows(rows)
		tw.SetIndexColumn(1)
		tw.SetTitle(strings.ToTitle(status + " Pet Details"))

		fmt.Println(tw.Render())
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

	petCmd.Flags().StringVarP(&status, "status", "s", "",
		"Provide status value. It can be available, pending, sold")

}
