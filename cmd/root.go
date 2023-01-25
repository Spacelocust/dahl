package cmd

import (
	"os"

	initial "github.com/Spacelocust/dahl/cmd/init"
	"github.com/Spacelocust/dahl/cmd/run"
	"github.com/Spacelocust/dahl/template"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const versionApp = "v0.1.0"

var rootCmd = &cobra.Command{
	Use:     "dahl",
	Long:   "Dahl is a CLI app for quickly generating files from a template,\nproviding users with an efficient way to create the same file structure.",
	Version: versionApp,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(run.RunCmd)
	rootCmd.AddCommand(initial.InitCmd)
	rootCmd.SetHelpTemplate(template.Help)
}

func initConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigName(".dahl-config")
}
