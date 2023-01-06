package tool

import (
	"errors"
	"os"
)

// Replace a config value by a flag value if the flag is set
func OverrideConfig(flagValue *string, configValue string) error {
	if *flagValue == "" {
		if configValue != "" {
			*flagValue = configValue
		} else {
			return errors.New("undefined property")
		}
	}

	return nil
}

// Convert an Interface to a String value
func InterfaceToString(inter interface{}) string {
	if inter == nil {
		return ""
	}

	return inter.(string)
}

// Convert a Slice to a Map
func SliceToMap(arr []string) map[string]string {
	newMap := make(map[string]string)
	for i := 0; i < len(arr); i += 2 {
		newMap[arr[i]] = arr[i+1]
	}

	return newMap
}

// Convert a Slice of slice to a Map
func SSToMap(arr [][]string) map[string]string {
	newMap := make(map[string]string)
	for _, v := range arr {
		newMap[v[0]] = v[1]
	}

	return newMap
}

// Check if a path exist
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
