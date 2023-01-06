package cmd

import (
	"os"

	"github.com/Spacelocust/dahl/cmd/stack"
	"github.com/Spacelocust/dahl/cmd/template"
	"github.com/Spacelocust/dahl/utils/logger"

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

	rootCmd.AddCommand(stack.StackCmd)
	rootCmd.AddCommand(template.TemplateCmd)
}

func initConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigName("dahl-config")

	if err := viper.ReadInConfig(); err != nil {
		logger.LogError("dahl-config file not found", viper.ConfigFileUsed())
	}
}
