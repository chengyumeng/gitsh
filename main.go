package main

import (
	"github.com/chengyumeng/gitsh/cmd"
)

var Version string

func main() {
	cmd.Version = Version

	cmd.RootCmd.Execute()
}