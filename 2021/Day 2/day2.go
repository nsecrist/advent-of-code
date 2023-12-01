package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	input, err := ReadLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	horizontalPosition, depth := CalculatePositionAndDepthWithAim(input)
	result := horizontalPosition * depth

	fmt.Printf("Final Horizontal Position: %d\n", horizontalPosition)
	fmt.Printf("Final Depth: %d\n", depth)
	fmt.Printf("Result (Horizontal Position * Depth): %d\n", result)
}

// Caluclates the Position and Depth using Aim
func CalculatePositionAndDepthWithAim(commands []string) (int, int) {
	aim := 0
	hPos := 0
	depth := 0

	for _, command := range commands {
		parts := strings.Fields(command)
		direction := parts[0]
		value := ParseValue(parts[1])

		switch direction {
		case "forward":
			hPos += value
			depth += (aim * value)
		case "down":
			aim += value
		case "up":
			aim -= value
		}
	}

	return hPos, depth
}

func ParseValue(s string) int {
	var value int
	fmt.Sscanf(s, "%d", &value)
	return value
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
