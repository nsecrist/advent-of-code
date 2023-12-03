package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

const MaxRed = 12
const MaxGreen = 13
const MaxBlue = 14

func main() {
	input, err := ReadLines("input.txt")
	if err != nil {
		log.Fatalf("ReadLines: %s", err)
	}

	Part1(input)
	Part2(input)
}

func Part1(input []string) {

	result := 0

	redRegex := regexp.MustCompile(`(\d+)\s+red`)
	greenRegex := regexp.MustCompile(`(\d+)\s+green`)
	blueRegex := regexp.MustCompile(`(\d+)\s+blue`)
	gameRegex := regexp.MustCompile(`(\d+):`)

	for _, line := range input {
		if (IsGamePossible(redRegex.FindAllStringSubmatch(line, -1), MaxRed)) &&
			IsGamePossible(greenRegex.FindAllStringSubmatch(line, -1), MaxGreen) &&
			IsGamePossible(blueRegex.FindAllStringSubmatch(line, -1), MaxBlue) {
			gameId, _ := strconv.Atoi(gameRegex.FindStringSubmatch(line)[1])
			result += gameId
		}
	}

	fmt.Println("Part 1 Result: ", result)
}

func Part2(input []string) {

	result := 0

	redRegex := regexp.MustCompile(`(\d+)\s+red`)
	greenRegex := regexp.MustCompile(`(\d+)\s+green`)
	blueRegex := regexp.MustCompile(`(\d+)\s+blue`)

	for _, line := range input {
		result += GetMaxCube(redRegex.FindAllStringSubmatch(line, -1)) *
			GetMaxCube(greenRegex.FindAllStringSubmatch(line, -1)) *
			GetMaxCube(blueRegex.FindAllStringSubmatch(line, -1))
	}

	fmt.Println("Part 2 Result: ", result)
}

func IsGamePossible(input [][]string, max int) bool {
	for _, match := range input {
		for _, capture := range match {
			gameCompare, _ := strconv.Atoi(capture)
			if gameCompare > max {
				return false
			}
		}
	}

	return true
}

func GetMaxCube(input [][]string) int {
	maxCubes := 0
	for _, match := range input {
		for _, capture := range match {
			gameCompare, _ := strconv.Atoi(capture)
			if gameCompare > maxCubes {
				maxCubes = gameCompare
			}
		}
	}

	return maxCubes
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
