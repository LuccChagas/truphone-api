package utils

import "strings"

func Sanitizer(str string) string {

	str = strings.TrimPrefix(str, " ")
	str = strings.TrimSuffix(str, " ")
	str = strings.Title(strings.ToLower(str))

	return str
}
