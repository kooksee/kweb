package cmd

import (
	"github.com/kooksee/kweb/internal/kweb"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	RootCmd.PersistentFlags().String("log_level", "debug", "Log level")
}

var RootCmd = &cobra.Command{
	Use:   "kweb",
	Short: "kweb server",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		return viper.Unmarshal(kweb.DefaultApp().Cfg)
	},
}
