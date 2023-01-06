package model

import (
	"errors"
	"reflect"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type FlagsAction struct {
	IsForce bool
	IsYes   bool
}

// Get a field value from a struct dynamicly
func GetValue[T Template|FlagsAction](s T, field string) (string, error) {
	r := reflect.ValueOf(s)
	f := reflect.Indirect(r).FieldByName(cases.Title(language.English).String(field))
	if !f.IsValid() {
		return "", errors.New("field invalid or doesn't exist")
	}

	return f.String(), nil
}