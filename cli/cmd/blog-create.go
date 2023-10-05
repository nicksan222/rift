/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	blog_create "rift/cmd/blog-create"

	"github.com/spf13/cobra"
)

// blog/CreateCmd represents the blog/create command
var BlogCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new blog post",

	Run: func(cmd *cobra.Command, args []string) {
		// Getting title
		title, err := blog_create.InputTitle()

		if err != nil {
			panic(err)
		}

		description, err := blog_create.InputDescription()

		if err != nil {
			panic(err)
		}

		image, err := blog_create.InputImage()

		if err != nil {
			panic(err)
		}

		tags, err := blog_create.InputTags()

		if err != nil {
			panic(err)
		}

		fmt.Println(title)
		fmt.Println(description)
		fmt.Println(image)
		fmt.Println(tags)
	},
}

func init() {
}
