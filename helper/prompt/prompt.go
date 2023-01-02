package prompt

import (
	"errors"
	"fmt"
	"strings"
)

// Show a Yes No prompt
func YesNo(s string, errMsg string, resDef ...string) error {
	var (
		response string
		question = fmt.Sprintf("%s [y/n]", s)
	)

	if len(resDef) > 0 {
		response = resDef[0]
		question = question + " " + "(default: " + response + ")"
	}

	fmt.Printf(question)

	_, err := fmt.Scanln(&response)
	if err != nil {
		if err.Error() == "unexpected newline" {
			if response != "" {
				return nil
			}

			fmt.Println("Invalid response. Please enter 'y' or 'n'.")
			return YesNo(s, errMsg)
		}

		return err
	}

	response = strings.ToLower(response)
	if response == "y" || response == "yes" {
		return nil
	} else if response == "n" || response == "no" {
		return errors.New(errMsg)
	} else {
		fmt.Println("Invalid response. Please enter 'y' or 'n'.")
		return YesNo(s, errMsg)
	}
}
