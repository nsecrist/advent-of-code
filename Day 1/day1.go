package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	lines, err := ReadLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	fmt.Println("Number of Measurements Larger than Previous: ", ScanMeasurements(lines))
}

func ScanMeasurements(lines []string) int64 {
	var prevInt int64
	var numLarger int64
	for i, line := range lines {
		currentInt, _ := strconv.ParseInt(line, 10, 32)

		if i > 0 {
			if currentInt > prevInt {
				numLarger++
			}
		}

		prevInt = currentInt
	}
	return numLarger
}

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
