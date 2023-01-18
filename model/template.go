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

type Props struct {
	JsonData string
	PathYml  *viper.Viper
}

type Template struct {
	From        string
	To          string
	Filename    string
	Prefix      string
	Suffix      string
	Extension   string
	Props       Props
	FlagsAction FlagsAction
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


	// TODO: refactor this part to be more clean
	var jsonMap map[string]interface{}
	if template.Props.JsonData != "" {
		// Convert json object to a map
		val, err := tool.JsonToMap(template.Props.JsonData)
		if err != nil {
			return err
		}

		jsonMap = val
	}

	// Get a Map of pair: [["_{ value }_"]: value] from a template
	matchs := tool.SSToMap(regexPattern.FindAllStringSubmatch(string(data), -1))

	// k: _{ value }_ , v: value
	var foundVal interface{}
	for k, v := range matchs {
		if value, err := GetValue(template, v); err == nil {
			properties = append(properties, k, value)
		} else {
			// Check if the pattern contain a "@" like _{ @value }_
			if strings.Contains(v[0:1], "@") {
				v = strings.Replace(v, "@", "", 1)
			}

			if len(jsonMap) > 0 {
				// Get value from jsonMap with the value pattern from the template
				val, _ := tool.FindInMap(jsonMap, strings.Split(v, "."))
				foundVal = val
			}


			// TODO: refactor this part to be more clean
			if template.Props.PathYml != nil {
				if prop := template.Props.PathYml.GetString(v); prop != "" {
					// Replace the value from the configYml with the value from the jsonMap
					if foundVal != nil && foundVal.(string) != "" {
						properties = append(properties, k, foundVal.(string))
					} else {
						properties = append(properties, k, prop)
					}
				} else {
					errString = fmt.Sprint(errString, k+", ")
				}
			} else {
				if foundVal != nil && foundVal.(string) != "" {
					properties = append(properties, k, foundVal.(string))
				} else {
					errString = fmt.Sprint(errString, k+", ")
				}
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

	// Check if the file already exist and ask if we overwritten or not
	if !template.FlagsAction.IsForce {
		if isExist, err := tool.Exists(pathTo + template.Extension); isExist && err == nil {
			if err := prompt.YesNo("File already exist, do you want to overwritten it ?", "your file creation has failed", false); err != nil {
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
