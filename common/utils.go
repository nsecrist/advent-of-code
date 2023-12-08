package common

import "strings"

func StringToLinesSlice(input *string, delimeter string) (lines []string) {
	return strings.Split(*input, delimeter)
}
