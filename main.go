package main

import (
	"github.com/kooksee/go-assert"
	"github.com/kooksee/kweb/cmd"
	"github.com/kooksee/kweb/internal/cnst"
	"os"
)

func main() {
	rootCmd := cmd.RootCmd
	rootCmd.AddCommand(
		cmd.InitFilesCmd,
		cmd.VersionCmd,
	)

	assert.MustNotError(cmd.PrepareBaseCmd(rootCmd, cnst.EnvPrefix,
		os.ExpandEnv(cnst.CurPath)).Execute())
}
