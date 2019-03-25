package cmd

import (
	"fmt"
	"github.com/kooksee/kweb/version"
	"github.com/spf13/cobra"
)

// VersionCmd ...
var VersionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "Show version info",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("kweb version", version.Version)
		fmt.Println("kweb commit version", version.CommitVersion)
		fmt.Println("kweb build version", version.BuildVersion)
	},
}
