package day9

import (
	"slices"
	"strconv"
	"strings"

	"github.com/nsecrist/advent-of-code/common"
)

func Part1(input *string) string {
	result := 0
	data := strings.Split(*input, "\n")
	for _, line := range data {
		lineInts := common.StringSliceToIntsSlice(strings.Split(line, " "))
		result += ExtrapolateNum(lineInts)
	}

	return strconv.Itoa(result)
}

func Part2(input *string) string {
	result := 0
	data := strings.Split(*input, "\n")
	for _, line := range data {
		lineInts := common.StringSliceToIntsSlice(strings.Split(line, " "))
		slices.Reverse(lineInts)
		result += ExtrapolateNum(lineInts)
	}

	return strconv.Itoa(result)
}

func ExtrapolateNum(nums []int) int {
	var res int
	extr := make([][]int, 0)
	extr = append(extr, nums)
	for {
		next := []int{}
		for i := 0; i < len(extr[len(extr)-1])-1; i++ {
			next = append(next, extr[len(extr)-1][i+1]-extr[len(extr)-1][i])
		}
		extr = append(extr, next)
		allzero := true
		for _, v := range next {
			if v != 0 {
				allzero = false
			}
		}
		if allzero {
			break
		}
	}

	for i := len(extr) - 2; i >= 0; i-- {
		extr[i] = append(extr[i], extr[i][len(extr[i])-1]+extr[i+1][len(extr[i+1])-1])
	}
	res = extr[0][len(extr[0])-1]
	return res
}
