package day6

import (
	"regexp"
	"strconv"
	"strings"
)

func Part1(input *string) string {
	data := strings.Split(*input, "\n")
	times, distances := ParseData(data)

	result := 0
	for i := 0; i < len(times); i++ {
		waysToWin := NumWaysToWin(times[i], distances[i])
		if result == 0 {
			result += waysToWin
		} else {
			result *= waysToWin
		}
	}

	return strconv.Itoa(result)
}

func Part2(input *string) string {
	data := strings.Split(*input, "\n")
	time, distance := ParseData2(data)

	result := NumWaysToWin(time, distance)
	return strconv.Itoa(result)
}

func ParseData(data []string) (times []int, distances []int) {
	times = ParseIntegers(strings.Split(data[0], ":")[1])
	distances = ParseIntegers(strings.Split(data[1], ":")[1])

	return
}

func ParseData2(data []string) (int, int) {
	time, _ := CombineIntegers(strings.Split(data[0], ":")[1])
	distance, _ := CombineIntegers(strings.Split(data[1], ":")[1])

	return time, distance
}

func CombineIntegers(in string) (result int, err error) {
	return strconv.Atoi(strings.ReplaceAll(in, " ", ""))
}

func ParseIntegers(in string) (result []int) {
	r := regexp.MustCompile(`\b\d+\b`)

	matches := r.FindAllString(in, -1)
	result = make([]int, len(matches))

	for i := 0; i < len(matches); i++ {
		result[i], _ = strconv.Atoi(matches[i])
	}

	return
}

func NumWaysToWin(totalRaceTime int, totalRaceDistance int) (result int) {
	result = 0
	for holdTime := 1; holdTime < totalRaceTime; holdTime++ {
		timeLeft := totalRaceTime - holdTime

		if (timeLeft * holdTime) > totalRaceDistance {
			result++
		}
	}

	return
}
