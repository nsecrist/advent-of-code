package day1

import (
	"bufio"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func GetDigitWords() []string {
	return []string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}
}

func Part1(input *string) string {
	data := strings.Split(*input, "\n")
	calibrationValue := 0
	for _, line := range data {
		fd := FindFirstOrLastDigit(line, false, false)
		ld := FindFirstOrLastDigit(line, true, false)
		calibrationValue += fd*10 + ld
	}

	return strconv.Itoa(calibrationValue)
}

func Part2(input *string) string {
	data := strings.Split(*input, "\n")
	calibrationValue := 0
	for _, line := range data {
		fd := FindFirstOrLastDigit(line, false, true)
		ld := FindFirstOrLastDigit(line, true, true)
		calibrationValue += fd*10 + ld
	}

	return strconv.Itoa(calibrationValue)
}

// Reverses a string
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

// Allows you to get the first or last digit from a calibration line.
func FindFirstOrLastDigit(s string, lastDigit bool, includeWords bool) (digit int) {

	if lastDigit {
		s = ReverseString(s)
	}

	r := regexp.MustCompile(`(\d)?([a-z]+)?(\d)?`)

	match := r.FindStringSubmatch(s)

	if match[1] != "" { // First Group
		digit, _ = strconv.Atoi(string(match[1][0]))
		return
	} else if match[2] != "" { // Second Group
		digit = ConvertToDigit(match[2], lastDigit)
		if digit > -1 {
			return
		}
	}

	// Last group
	digit, _ = strconv.Atoi(string(match[3][0]))
	return
}

// Converts calibration value to digit
func ConvertToDigit(s string, isReversed bool) (result int) {
	resultMap := make(map[int]int)

	if !isReversed {
		for _, word := range GetDigitWords() {
			if strings.Contains(s, word) {
				resultMap[strings.Index(s, word)] = GetIntFromWord(word)
			}
		}
	} else {
		for _, word := range GetDigitWords() {
			if strings.Contains(s, ReverseString(word)) {
				resultMap[strings.Index(s, ReverseString(word))] = GetIntFromWord(word)
			}
		}
	}

	if len(resultMap) < 1 {
		return -1
	}

	keys := make([]int, 0, len(resultMap))
	for k := range resultMap {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	return resultMap[keys[0]]
}

// Switch statement to convert from lower-case digit words to integers
func GetIntFromWord(w string) int {
	switch w {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "zero":
		return 0
	default:
		return -1
	}
}

// Read Lines from file with input path
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}
