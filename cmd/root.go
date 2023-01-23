package cmd

import (
	"os"

	initial "github.com/Spacelocust/dahl/cmd/init"
	"github.com/Spacelocust/dahl/cmd/run"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "dahl",
	Short: "A brief description of your application",
	Long:  `Root`,
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
}

func initConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigName(".dahl-config")
}
