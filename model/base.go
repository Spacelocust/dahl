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

// Get a field value from a struct
func GetValue[T Template | FlagsAction](s T, field string) (string, error) {
	r := reflect.ValueOf(s)
	f := reflect.Indirect(r).FieldByName(cases.Title(language.English).String(field))
	if !f.IsValid() {
		return "", errors.New("field invalid or doesn't exist")
	}

	return f.String(), nil
}

/*
	func (t *Template) SetValue(field string, newValue string) error {
		if f := reflect.ValueOf(&t).Indirect().FieldByName(field); f.IsValid() {
			f.SetString(newValue)
		} else {
			return errors.New("field invalid or doesn't exist")
		}

		return nil
	}
*/

// Set a field value for a struct
func SetValue(s *Template, field string, value string) {
	v := reflect.ValueOf(s).Elem().FieldByName(field)
	if v.IsValid() {
		v.SetString(value)
	}
}
