package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Replace this with your actual file path or provide the input directly in the code
	filePath := "input.txt"

	// Read the diagnostic report from the file
	diagnosticReport, err := readDiagnosticReport(filePath)
	if err != nil {
		fmt.Println("Error reading diagnostic report:", err)
		return
	}

	// Calculate the gamma and epsilon rates
	gammaRate := calculateRate(diagnosticReport, true)
	epsilonRate := calculateRate(diagnosticReport, false)

	// Calculate the power consumption
	powerConsumption := gammaRate * epsilonRate

	fmt.Printf("Gamma Rate: %d\n", gammaRate)
	fmt.Printf("Epsilon Rate: %d\n", epsilonRate)
	fmt.Printf("Power Consumption: %d\n", powerConsumption)
}

func readDiagnosticReport(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var diagnosticReport []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		diagnosticReport = append(diagnosticReport, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return diagnosticReport, nil
}

func calculateRate(diagnosticReport []string, mostCommon bool) int {
	bitCounts := make([]map[byte]int, len(diagnosticReport[0]))

	for i := range diagnosticReport[0] {
		bitCounts[i] = make(map[byte]int)
	}

	for _, binaryNumber := range diagnosticReport {
		for i, bit := range binaryNumber {
			bitCounts[i][byte(bit)]++
		}
	}

	var rate int
	for _, counts := range bitCounts {
		var selectedBit byte
		if mostCommon {
			selectedBit = findMostCommonBit(counts)
		} else {
			selectedBit = findLeastCommonBit(counts)
		}

		rate = rate<<1 | int(selectedBit-'0')
	}

	return rate
}

func findMostCommonBit(counts map[byte]int) byte {
	var mostCommonBit byte
	maxCount := 0

	for bit, count := range counts {
		if count > maxCount {
			maxCount = count
			mostCommonBit = bit
		}
	}

	return mostCommonBit
}

func findLeastCommonBit(counts map[byte]int) byte {
	var leastCommonBit byte
	minCount := len(counts)

	for bit, count := range counts {
		if count < minCount {
			minCount = count
			leastCommonBit = bit
		}
	}

	return leastCommonBit
}
