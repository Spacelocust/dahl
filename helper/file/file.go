package file

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/Spacelocust/dahl/helper/tool"
	"github.com/spf13/viper"
)

const TEMPLATE_DIR = "./.dahl"

// Regex to get _{ ... }_ pattern
var regexPattern = regexp.MustCompile("_\\{\\s*(.*?)\\s*\\}\\_")

func CreateFile(args map[string]string, template string) error {

	// Get access to props key from the config template
	props := viper.Sub(template + ".props")

	// Properties from args going to be replace in file
	var properties []string

	// Error string contains all of pattern declared in a template but with no values
	var errString string

	to := args["to"]
	from := args["from"]
	filename := args["filename"]
	extension := args["extension"]
	pathTo := to + "/" + filename

	// Read template file
	data, err := os.ReadFile(TEMPLATE_DIR + from)
	if err != nil {
		return err
	}

	// Get a Map of ["_{ ... }_"]: ... from a template
	arrMatch := tool.SSToMap(regexPattern.FindAllStringSubmatch(string(data), -1))

	for k, v := range arrMatch {
		value, ok := args[v]
		if ok {
			properties = append(properties, k, value)
		} else {

			// Check if the pattern contains a "@" like _{ @value }_
			if strings.Contains(v[0:1], "@") {
				v = strings.Replace(v, "@", "", 1)
			}

			if prop := props.GetString(v); prop != "" {
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

	// Create all directory(ies) if not presente
	err = os.MkdirAll(to+"/", os.ModePerm)
	if err != nil {
		return err
	}

	if args["prefix"] != "" {
		pathTo = args["prefix"] + pathTo
	}

	if args["suffix"] != "" {
		pathTo = pathTo + args["suffix"]
	}

	// Write the updated content back to the file
	err = os.WriteFile(pathTo+extension, []byte(newContent), 0666)
	if err != nil {
		return err
	}

	return nil
}
