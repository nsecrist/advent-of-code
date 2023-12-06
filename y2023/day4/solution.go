package day4

import (
	"bufio"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Part1(input *string) string {
	data := strings.Split(*input, "\n")
	result := 0

	for _, card := range data {
		score, _ := GetCardPointValue(card)
		result += score
	}

	return strconv.Itoa(result)
}

func Part2(input *string) string {
	data := strings.Split(*input, "\n")
	result := 0

	cardCount := make([]int, len(data))

	for i := range cardCount {
		cardCount[i] = 1
	}

	for i, card := range data {
		_, matches := GetCardPointValue(card)

		cardCount = AddCardCopies(i, cardCount[i], matches, cardCount)
	}

	for _, val := range cardCount {
		result += val
	}

	return strconv.Itoa(result)
}

func AddCardCopies(cI int, numCards int, matches int, cards []int) []int {
	for i := 1; i <= matches; i++ {
		cards[cI+i] += numCards
	}

	return cards
}

func GetCardPointValue(card string) (score, matches int) {
	cardRegex := regexp.MustCompile(`^(?P<g>.*?):\s*(?P<wn>.*?)\s*\|\s*(?P<my>.*)$`)
	groups := SubMatchMap(cardRegex, card)

	winningNumbers := NumbersStringToArray(groups["wn"])
	myNumbers := NumbersStringToArray(groups["my"])

	matches = 0

	for _, myNum := range myNumbers {
		for _, wNum := range winningNumbers {
			if myNum == wNum {
				matches++
				break
			}
		}
	}

	if matches > 0 {
		return powInt(2, matches-1), matches
	} else {
		return 0, matches
	}
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func SubMatchMap(r *regexp.Regexp, str string) map[string]string {
	match := r.FindStringSubmatch(str)
	subMatchMap := make(map[string]string)
	for i, name := range r.SubexpNames() {
		if i != 0 {
			subMatchMap[name] = match[i]
		}
	}

	return subMatchMap
}

func NumbersStringToArray(numbers string) []int {
	numreg := regexp.MustCompile(`\b\d+\b`)
	matches := numreg.FindAllString(numbers, -1)

	ints := make([]int, len(matches))
	for i, match := range matches {
		ints[i], _ = strconv.Atoi(match)
	}

	return ints
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
