package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "web-cli",
	Short: "Read Atom Feeds",
	Long: `Just a small CLI app.
	       Read Atom feeds`,
}

var cmdLs = &cobra.Command{
	Use:   "ls",
	Short: "list news",
	Long:  "list first 5 news",
	Run: func(cmd *cobra.Command,
		args []string) {
		list()
	},
}

var cmdDescribe = &cobra.Command{
	Use:   "describe [id]",
	Short: "show details of feed with id",
	Long:  "Details for an article with id",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		describe(args[0])
	},
}

func Exec() {
	rootCmd.AddCommand(cmdLs)
	rootCmd.AddCommand(cmdDescribe)
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
