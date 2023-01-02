/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package run

import (
	"regexp"

	"github.com/Spacelocust/dahl/helper/file"
	"github.com/Spacelocust/dahl/model"

	"github.com/Spacelocust/dahl/helper/logger"
	"github.com/Spacelocust/dahl/helper/tool"

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

	properties  = map[string]string{}
	flagsAction = model.FlagsAction{}

	// Regex to check extension set
	regexExt = regexp.MustCompile(`\.[a-z]{2,6}$`)
)

// runCmd represents the run command
var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "run template",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		template := "templates." + args[0]

		// Get access to first level values of a template
		configTemplate := viper.GetStringMap(template)

		if len(configTemplate) > 0 {

			if isChanged := cmd.Flag("yes").Changed; isChanged {
				flagsAction.IsYes = true
			}

			if isChanged := cmd.Flag("force").Changed; isChanged {
				flagsAction.IsForce = true
			}

			if err := tool.OverrideConfig(&filename, tool.InterfaceToString(configTemplate["filename"])); err != nil {
				logger.LogError(err.Error(), "filename")
			} else {
				properties["filename"] = filename
			}

			if err := tool.OverrideConfig(&to, tool.InterfaceToString(configTemplate["to"])); err != nil {
				logger.LogError(err.Error(), "to")
			} else {
				properties["to"] = to
			}

			if err := tool.OverrideConfig(&from, tool.InterfaceToString(configTemplate["from"])); err != nil {
				logger.LogError(err.Error(), "from")
			} else {
				properties["from"] = from
			}

			if err := tool.OverrideConfig(&extension, tool.InterfaceToString(configTemplate["extension"])); err == nil {
				isValid := regexExt.MatchString(extension)

				if !isValid {
					logger.LogError("wrong format extension, please use '.your-extension' format")
				}

				properties["extension"] = extension
			}

			if err := tool.OverrideConfig(&suffix, tool.InterfaceToString(configTemplate["suffix"])); err == nil {
				properties["suffix"] = suffix
			}

			if err := tool.OverrideConfig(&prefix, tool.InterfaceToString(configTemplate["prefix"])); err == nil {
				properties["prefix"] = prefix
			}

			if err := file.CreateFile(properties, template, flagsAction); err != nil {
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
