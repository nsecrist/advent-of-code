package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	engineSchematic := string(input)

	Part1(engineSchematic)
	Part2(engineSchematic)
}

func Part1(input string) {
	lines := strings.Split(input, "\n")
	result := 0

	fixed := lines

	// fixed := make([]string, len(lines))
	// for i, fix := range lines {
	// 	fixed[i] = fix + "\n"
	// }

	for i := 0; i < len(fixed); i++ {
		for j := 0; j < len(fixed[i]); j++ {
			char := rune(fixed[i][j])

			// Skip periods
			if char == '.' {
				continue
			}

			// Check if the character is a symbol
			if isSymbol(char) {
				// Check adjacent numbers (including diagonals)
				for x := -1; x <= 1; x++ {
					for y := -1; y <= 1; y++ {
						if i+x >= 0 && i+x < len(fixed) && j+y >= 0 && j+y < len(fixed[i+x]) {
							adjacentChar := rune(fixed[i+x][j+y])
							if isDigit(adjacentChar) {
								number, _ := GetNumber(fixed[i+x], j+y)
								result += number
							}
						}
					}
				}
			}
		}
	}

	fmt.Println("Part 1 Result: ", result)
}

func Part2(input string) {

	result := 0

	fmt.Println("Part 2 Result: ", result)
}

func GetNumber(line string, index int) (int, error) {
	var numberString string = ""
	runeLine := []rune(line)
	for i := index; i < len(line) && isDigit(runeLine[i]); i-- {
		numberString += string(line[i])
		if i == 0 {
			break
		}
	}

	reversed := ReverseString(numberString)

	for i := index + 1; i < len(line) && isDigit(runeLine[i]); i++ {
		reversed += string(line[i])
	}

	return strconv.Atoi(reversed)
}

func isSymbol(char rune) bool {
	// Define symbols that indicate a part number
	// symbols := "*+#$-%=/&"
	// symbols := "!@#$%^&*()+-=~`_?><,[]{}/|;:"
	symbols := "@#$_&-+/*=\r"
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
