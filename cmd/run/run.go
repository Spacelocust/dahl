/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package run

import (
	"regexp"

	"github.com/Spacelocust/dahl/model"

	"github.com/Spacelocust/dahl/utils/logger"
	"github.com/Spacelocust/dahl/utils/tool"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	from      string
	to        string
	filename  string
	prefix    string
	suffix    string
	extension string
	props     string

	// Models
	template model.Template

	// Regex to check extension set
	regexExt = regexp.MustCompile(`\.[a-z]{2,6}$`)
)

// runCmd represents the run command
var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "run template",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		template.Props = model.Props{}

		if len(args) > 0 {
			if err := viper.ReadInConfig(); err != nil {
				logger.LogError(".dahl-config.yml file not found", viper.ConfigFileUsed())
			}

			pathYml := "templates." + args[0]
			configYml := viper.GetStringMap(pathYml)
			template.Props.PathYml = viper.Sub(pathYml + ".props")

			if len(configYml) > 0 {
				withConfig(configYml)
			} else {
				logger.LogError("template not found")
			}
		} else {
			withoutConfig()
		}

		// Set if the flag yes is used
		if isChanged := cmd.Flag("yes").Changed; isChanged {
			template.FlagsAction.IsYes = true
		}

		// Set if the flag force is used
		if isChanged := cmd.Flag("force").Changed; isChanged {
			template.FlagsAction.IsForce = true
		}

		if props != "" {
			template.Props.JsonData = props
		}

		if err := template.CreateFile(); err != nil {
			logger.LogError(err.Error())
		}

		logger.LogSuccess(template.Filename+template.Extension, "has been saved successfully")
	},
}

func withoutConfig() {
	if filename != "" {
		template.Filename = filename
	} else {
		logger.LogError("undefined property", "filename")
	}

	if to != "" {
		template.To = to
	} else {
		logger.LogError("undefined property", "to")
	}

	if from != "" {
		template.From = from
	} else {
		logger.LogError("undefined property", "from")
	}

	if extension != "" {
		if isValid := regexExt.MatchString(extension); !isValid {
			logger.LogError("wrong format extension, please use '.your-extension' format")
		}

		template.Extension = extension
	}

	if suffix != "" {
		template.Suffix = suffix
	}

	if prefix != "" {
		template.Prefix = prefix
	}
}

func withConfig(configYml map[string]interface{}) {
	if err := tool.OverrideConfig(&filename, tool.InterfaceToString(configYml["filename"])); err == nil {
		template.Filename = filename
	} else {
		logger.LogError(err.Error(), "filename")
	}

	if err := tool.OverrideConfig(&to, tool.InterfaceToString(configYml["to"])); err == nil {
		template.To = to
	} else {
		logger.LogError(err.Error(), "to")
	}

	if err := tool.OverrideConfig(&from, tool.InterfaceToString(configYml["from"])); err == nil {
		template.From = from
	} else {
		logger.LogError(err.Error(), "from")
	}

	if err := tool.OverrideConfig(&extension, tool.InterfaceToString(configYml["extension"])); err == nil {
		if isValid := regexExt.MatchString(extension); !isValid {
			logger.LogError("wrong format extension, please use '.your-extension' format")
		}

		template.Extension = extension
	}

	if err := tool.OverrideConfig(&suffix, tool.InterfaceToString(configYml["suffix"])); err == nil {
		template.Suffix = suffix
	}

	if err := tool.OverrideConfig(&prefix, tool.InterfaceToString(configYml["prefix"])); err == nil {
		template.Prefix = prefix
	}
}

/*func flagSetting(template *model.Template, flagValue string, key string, config map[string]interface{}, optional bool) {
	if err := tool.OverrideConfig(&flagValue, tool.InterfaceToString(config[key])); err == nil {
		model.SetValue(&template, key, flagValue)
	} else {
		if !optional {
			logger.LogError(err.Error(), key)
		}
	}
}*/

func init() {
	RunCmd.Flags().StringVarP(&filename, "filename", "n", "", "")
	RunCmd.Flags().StringVar(&from, "from", "", "file localtion")
	RunCmd.Flags().StringVar(&to, "to", "", "file destination")
	RunCmd.Flags().StringVarP(&prefix, "prefix", "p", "", "filename prefix")
	RunCmd.Flags().StringVarP(&suffix, "suffix", "s", "", "filename suffix")
	RunCmd.Flags().StringVarP(&extension, "extension", "e", "", "extension filename")
	RunCmd.Flags().StringVar(&props, "props", "", "props file template")

	RunCmd.Flags().BoolP("force", "f", true, "force overwritten on a existing file")
	RunCmd.Flags().BoolP("yes", "y", true, "shorcut for yes answer for the question for creating a the new path")
}
