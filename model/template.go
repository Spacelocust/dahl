package model

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/Spacelocust/dahl/utils/prompt"
	"github.com/Spacelocust/dahl/utils/tool"
	"github.com/spf13/viper"
)

type Template struct {
	From        string
	To          string
	Filename    string
	Prefix      string
	Suffix      string
	Extension   string
	FlagsAction FlagsAction
	Props       *viper.Viper
}

const (
	TEMPLATE_DIR = "./.dahl"
	REGEX        = `\_{\s*(\S+)\s*\}_` // "_\\{\\s*(.*?)\\s*\\}\\_"
)

// Regex to get _{ ... }_ pattern
var regexPattern = regexp.MustCompile(REGEX)

func (template Template) CreateFile() error {
	// Properties from template going to be replace in file
	var properties []string

	// Error string contains all of pattern declared in a template but with no values
	var errString string

	pathTo := template.To + "/" + template.Filename

	// Read a template file
	data, err := os.ReadFile(TEMPLATE_DIR + template.From)
	if err != nil {
		return err
	}

	// Get a Map of ["_{ ... }_"]: ... from a template
	arrMatch := tool.SSToMap(regexPattern.FindAllStringSubmatch(string(data), -1))

	for k, v := range arrMatch {
		value, err := GetValue(template, v)
		if err == nil {
			properties = append(properties, k, value)
		} else {
			// Check if the pattern contain a "@" like _{ @value }_
			if strings.Contains(v[0:1], "@") {
				v = strings.Replace(v, "@", "", 1)
			}

			if prop := template.Props.GetString(v); prop != "" {
				properties = append(properties, k, prop)
			} else {
				errString = fmt.Sprint(errString, k+", ")
			}
		}
	}

	if errString != "" {
		return errors.New(errString + "declared in template, but no values set")
	}

	// Replace patterns in the file's content
	replacer := strings.NewReplacer(properties...)

	// Content with replaced values
	newContent := replacer.Replace(string(data))

	// Check if the path directories exist and ask if we make it or not
	if !template.FlagsAction.IsYes {
		if isExist, err := tool.Exists(template.To + "/"); !isExist && err == nil {
			if err := prompt.YesNo("Path directories doesn't exist, do you want to make it ?", "your file creation has failed", true); err != nil {
				return err
			}
		}
		if err != nil {
			return err
		}
	}

	// Create directory(ies) if not exist
	err = os.MkdirAll(template.To+"/", os.ModePerm)
	if err != nil {
		return err
	}

	if template.Prefix != "" {
		pathTo = template.Prefix + pathTo
	}

	if template.Suffix != "" {
		pathTo = pathTo + template.Suffix
	}

	// Check if the file all ready exist and ask if we overwritten or not
	if !template.FlagsAction.IsForce {
		if isExist, err := tool.Exists(pathTo + template.Extension); isExist && err == nil {
			if err := prompt.YesNo("File all ready exist, do you want to overwritten it ?", "your file creation has failed", false); err != nil {
				return err
			}
		}
		if err != nil {
			return err
		}
	}

	// Write the updated content back to the file
	err = os.WriteFile(pathTo+template.Extension, []byte(newContent), 0666)
	if err != nil {
		return err
	}

	return nil
}
