package init

import (
	"os"

	"github.com/Spacelocust/dahl/utils/logger"
	"github.com/Spacelocust/dahl/utils/tool"
	"github.com/Spacelocust/dahl/template"
	"github.com/spf13/cobra"
)

const configFile = ".dahl-config.yml"

var (
	dirTemplate = "./.dahl/"
	configFileTemp = template.ConfigFile
)

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Generate directory template and configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		isHello := false

		// Set if the flag hello-world is used
		if isChanged := cmd.Flag("hello-world").Changed; isChanged {
			isHello = true
		}

		if isHello {
			dirTemplate = dirTemplate + "example/"
		} else {
			configFileTemp = ""
		}

		// Create a dahl-config.yml file if mot exist
		if isExist, err := tool.Exists(configFile); !isExist && err == nil {
			err = os.WriteFile(configFile, []byte(configFileTemp), 0666)
			if err != nil {
				logger.LogError(err.Error())
			}
		}

		// Create path ./.dahl/example/
		if isExist, err := tool.Exists(dirTemplate); !isExist && err == nil {
			err = os.MkdirAll(dirTemplate, os.ModePerm)
			if err != nil {
				logger.LogError(err.Error())
			}

			if isHello {
				// Create a hello-world.dahl at ./.dahl/example/
				err = os.WriteFile(dirTemplate+"hello-world.dahl", []byte(template.HelloWorld), 0666)
				if err != nil {
					logger.LogError(err.Error())
				}
			}
		}
	},
}

func init() {
	InitCmd.Flags().Bool("hello-world", true, "initialise with an example template and configuration")
}
