/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// blogDeleteCmd represents the blogDelete command
var BlogDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Elimina un post del blog",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("blogDelete called")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// blogDeleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// blogDeleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
