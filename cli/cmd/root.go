/*
Copyright © 2023 Nicholas Santi <nicksan222@icloud.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "rift",
	Short: "This is a CLI for the Rift CMS.",
	Long: `Ever wonder why all existing CMS's are so bloated? Well, so did I.
Rift is a CMS that is built to be lightweight and fast.

Rift is built with Go and uses YAML for data storage.
For most use cases that is more than enough.

Rift is also built with a CLI in mind. This means that you can
create, edit, and delete content from the command line, even
on a remote server.

Thanks to Astro's content management system, you can also
have typesafe content.
	`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.rift.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
