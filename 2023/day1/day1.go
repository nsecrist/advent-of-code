package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	input, err := ReadLines("input.txt")
	if err != nil {
		log.Fatalf("ReadLines: %s", err)
	}

	calibrationValue := 0
	for _, line := range input {
		fd, ld := FindFirstAndLastDigits(line)
		calibrationValue += fd*10 + ld
	}

	fmt.Println("Calibration Value: ", calibrationValue)
}

func FindFirstAndLastDigits(s string) (fD, lD int) {
	firstIndex := -1
	lastIndex := -1

	// Iterate through the string to find the first and last integers
	for i, char := range s {
		// Check if the character is a digit
		if unicode.IsDigit(char) {
			// If it's the first integer, set firstIndex
			if firstIndex == -1 {
				firstIndex, lastIndex = i, i
			}

			// Update lastIndex with each occurrence
			lastIndex = i
		}
	}

	fD = int(s[firstIndex] - '0')
	lD = int(s[lastIndex] - '0')

	return
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
