package util

import "fmt"

func AsBool(v string) bool {
	return v == "1" || v == "true" || v == "yes"
}

func AsString(v any) string {
	return fmt.Sprintf("%+v", v)
}

func BoolAsString(v bool) string {
	if v == true {
		return "1"
	}

	return "0"
}
