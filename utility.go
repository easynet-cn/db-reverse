package main

import "strings"

func UpperCamelCase(txt string) string {
	sb := new(strings.Builder)

	strs := strings.Split(txt, "_")

	for _, str := range strs {
		sb.WriteString(strings.ToUpper(string(str[0])))
		sb.WriteString(str[1:])
	}

	return sb.String()
}

func LowerCamelCase(txt string) string {
	sb := new(strings.Builder)

	strs := strings.Split(txt, "_")

	for i, str := range strs {
		if i == 0 {
			sb.WriteString(strings.ToLower(string(str[0])))
		} else {
			sb.WriteString(strings.ToUpper(string(str[0])))
		}

		sb.WriteString(str[1:])
	}

	return sb.String()
}
