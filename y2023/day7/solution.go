package day7

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Hand struct {
	Cards string
	Bid   int
}

func Part1(input *string) string {
	data := strings.Split(*input, "\n")

	hands := []Hand{}
	for _, s := range data {
		hands = append(hands, ParseHand(s))
	}

	return strconv.Itoa((CalcTotal(hands, false)))
}

func Part2(input *string) string {
	data := strings.Split(*input, "\n")

	hands := []Hand{}
	for _, s := range data {
		hands = append(hands, ParseHand(s))
	}

	return strconv.Itoa((CalcTotal(hands, true)))
}

func CalcTotal(hands []Hand, wilds bool) int {
	t := 0
	slices.SortFunc(hands, func(a, b Hand) int {
		return cmp(a.Cards, b.Cards, wilds)
	})
	for i, h := range hands {
		t += (i + 1) * h.Bid
	}
	return t
}

func ParseHand(line string) Hand {
	h := Hand{}
	fmt.Sscanf(line, "%s %d", &h.Cards, &h.Bid)
	return h
}

func cmp(a, b string, jokers bool) int {
	j, r := "J", "TAJBQCKDAE"
	if jokers {
		j, r = "23456789TQKA", "TAJ0QCKDAE"
	}

	typ := func(cards string) string {
		k := 0
		for _, j := range strings.Split(j, "") {
			n, t := strings.ReplaceAll(cards, "J", j), 0
			for _, s := range n {
				t += strings.Count(n, string(s))
			}
			k = slices.Max([]int{k, t})
		}
		return map[int]string{5: "0", 7: "1", 9: "2", 11: "3", 13: "4", 17: "5", 25: "6"}[k]
	}

	return strings.Compare(
		typ(a)+strings.NewReplacer(strings.Split(r, "")...).Replace(a),
		typ(b)+strings.NewReplacer(strings.Split(r, "")...).Replace(b),
	)
}
