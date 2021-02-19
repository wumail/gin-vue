package utility

import "strings"

func ReplaceI(str string) string {
	str = strings.ReplaceAll(str, "\\n\\t", "<br\\>")
	str = strings.ReplaceAll(str, "\\t", "\\    ")
	return str
}

func ReplaceO(str string) string {
	strings.ReplaceAll(str, "<br\\>", "\\n\\t")
	strings.ReplaceAll(str, "\\    ", "\\t")
	return str
}
