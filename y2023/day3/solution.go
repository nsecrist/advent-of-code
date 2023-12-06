package day3

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Part1(input *string) string {
	data := strings.Split(*input, "\n")
	result := 0
	partNums := make([]int, 0)
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			char := rune(data[i][j])

			// Skip periods
			if char == '.' {
				continue
			}

			// Check if the character is a symbol
			if isSymbol(char) {
				// Check adjacent numbers (including diagonals)
				for x := -1; x <= 1; x++ {
					for y := -1; y <= 1; y++ {
						if i+x >= 0 && i+x < len(data) && j+y >= 0 && j+y < len(data[i+x]) {
							adjacentChar := rune(data[i+x][j+y])
							if isDigit(adjacentChar) {
								number, indexes, _ := GetNumber(data[i+x], j+y)
								partNums = append(partNums, number)
								data[i+x] = ClearIndexes(data[i+x], indexes)
							}
						}
					}
				}
			}
		}
	}
	result = SumValuesInSlice(partNums)
	return strconv.Itoa(result)
}

func Part2(input *string) string {
	data := strings.Split(*input, "\n")
	result := 0
	partNums := make([]int, 0)
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			char := rune(data[i][j])

			// Skip periods
			if char == '.' {
				continue
			}

			// Check if the character is a gear symbol
			if isGearSymbol(char) {
				gears := make([]int, 0)
				// Check adjacent numbers (including diagonals)
				for x := -1; x <= 1; x++ {
					for y := -1; y <= 1; y++ {
						if i+x >= 0 && i+x < len(data) && j+y >= 0 && j+y < len(data[i+x]) {
							adjacentChar := rune(data[i+x][j+y])
							if isDigit(adjacentChar) {
								number, indexes, _ := GetNumber(data[i+x], j+y)
								gears = append(gears, number)
								data[i+x] = ClearIndexes(data[i+x], indexes)
							}
						}
					}
				}
				if len(gears) == 2 {
					partNums = append(partNums, gears[0]*gears[1])
				}
			}
		}
	}
	result = SumValuesInSlice(partNums)
	return strconv.Itoa(result)
}

func ClearIndexes(input string, indexes []int) string {
	for _, idx := range indexes {
		input = ReplaceAtIndex(input, rune('.'), idx)
	}

	return input
}

func ReplaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func SumValuesInSlice(slice []int) (result int) {
	for _, value := range slice {
		result += value
	}

	return
}

func RemoveAdjacentDuplicates(input []int) []int {
	var result []int

	for i := 0; i < len(input); i++ {
		// Skip if the current element is equal to the next element
		if i < len(input)-1 && input[i] == input[i+1] {
			continue
		}

		// Otherwise, add the current element to the result slice
		result = append(result, input[i])
	}

	return result
}

func GetNumber(line string, index int) (int, []int, error) {
	var numberString string = ""
	indexes := make([]int, 0)
	runeLine := []rune(line)
	for i := index; i < len(line) && isDigit(runeLine[i]); i-- {
		numberString += string(line[i])
		indexes = append(indexes, i)
		if i == 0 {
			break
		}
	}

	reversed := ReverseString(numberString)

	for i := index + 1; i < len(line) && isDigit(runeLine[i]); i++ {
		reversed += string(line[i])
		indexes = append(indexes, i)
	}
	result, err := strconv.Atoi(reversed)
	return result, indexes, err
}

func isSymbol(char rune) bool {
	symbols := "@#$_&-+/*=%"
	return strings.ContainsRune(symbols, char)
}

func isGearSymbol(char rune) bool {
	symbols := "*"
	return strings.ContainsRune(symbols, char)
}

func isDigit(char rune) bool {
	return char >= '0' && char <= '9'
}

// Reverses a string
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
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
