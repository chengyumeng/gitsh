package cmd

import (
	"github.com/spf13/cobra"
	"github.com/chengyumeng/gitsh/cmd/author"
	"fmt"
)

var Version string

var RootCmd = &cobra.Command{
	Use: "gitsh",
}

var VersionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("gitsh %s \n", Version)
	},
}

func init() {
	cobra.EnableCommandSorting = false
	RootCmd.AddCommand(VersionCmd)
	RootCmd.AddCommand(author.AuthorCmd)
}

