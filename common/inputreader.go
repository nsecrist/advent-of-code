package common

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Read entire file from inout path
func ReadFileInput(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s, error: %s", path, err.Error())
	}
	data, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s, error: %s", path, err.Error())
	}
	return string(data), nil
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

func StringToLinesSlice(input *string, delimeter string) (lines []string) {
	return strings.Split(*input, delimeter)
}
