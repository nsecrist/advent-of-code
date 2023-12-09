package common

import (
	"strconv"
	"strings"
)

func StringToLinesSlice(input *string, delimeter string) (lines []string) {
	return strings.Split(*input, delimeter)
}

func StringSliceToIntsSlice(in []string) (result []int) {
	result = make([]int, len(in))

	for i, s := range in {
		result[i], _ = strconv.Atoi(s)
	}
	return
}
