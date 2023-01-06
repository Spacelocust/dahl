package prompt

import (
	"errors"
	"fmt"
	"strings"
)

// Show a Yes No prompt
func YesNo(s string, errMsg string, def ...bool) error {
	var (
		response   string
		defaultRes string
		question   = fmt.Sprintf("%s [y/n]", s)
	)

	// Check default set
	if len(def) > 0 {
		if def[0] {
			defaultRes = "yes"
			question = fmt.Sprintf("%s (default: yes) ", question)
		} else {
			defaultRes = "no"
			question = fmt.Sprintf("%s (default: no) ", question)
		}
	}

	fmt.Printf("%s", question)

	_, err := fmt.Scanln(&response)
	response = strings.ToLower(response)

	if err != nil {
		if err.Error() == "unexpected newline" {
			if defaultRes == "yes" {
				return nil
			} else if defaultRes == "no" {
				return errors.New(errMsg)
			}

			fmt.Println("Invalid response. Please enter 'y' or 'n'.")
			return YesNo(s, errMsg)
		}

		return err
	}

	if response == "y" || response == "yes" {
		return nil
	} else if response == "n" || response == "no" {
		return errors.New(errMsg)
	} else {
		fmt.Println("Invalid response. Please enter 'y' or 'n'.")
		return YesNo(s, errMsg)
	}
}
