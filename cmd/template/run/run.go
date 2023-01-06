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

	// Models
	flagsAction model.FlagsAction
	template    model.Template

	// Regex to check extension set
	regexExt = regexp.MustCompile(`\.[a-z]{2,6}$`)
)

// runCmd represents the run command
var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "run template",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pathYml := "templates." + args[0]

		// Get access to first level values of a template
		configYml := viper.GetStringMap(pathYml)

		if len(configYml) > 0 {

			// Set if the flag yes is used
			if isChanged := cmd.Flag("yes").Changed; isChanged {
				flagsAction.IsYes = true
			}

			// Set if the flag force is used
			if isChanged := cmd.Flag("force").Changed; isChanged {
				flagsAction.IsForce = true
			}

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

			template.FlagsAction = flagsAction

			template.Props = viper.Sub(pathYml + ".props")

			if err := template.CreateFile(); err != nil {
				logger.LogError(err.Error())
			}

			logger.LogSuccess("your file has been created successfully")

		} else {
			logger.LogError("template not found")
		}
	},
}

func init() {
	RunCmd.Flags().StringVarP(&filename, "filename", "n", "", "name of file")
	RunCmd.Flags().StringVarP(&to, "to", "", "", "destination of file")
	RunCmd.Flags().StringVarP(&prefix, "prefix", "p", "", "prefix of filename")
	RunCmd.Flags().StringVarP(&suffix, "suffix", "s", "", "suffix of filename")
	RunCmd.Flags().StringVarP(&extension, "extension", "e", "", "extension of filename")

	RunCmd.Flags().BoolP("force", "f", true, "force overwritten on a existing file")
	RunCmd.Flags().BoolP("yes", "y", true, "shorcut for yes answer for the question for creating a the new path")
}
