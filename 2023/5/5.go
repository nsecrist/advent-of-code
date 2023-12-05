package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Almanac struct {
	seeds   []int64
	seedLoc []int64
	sts     RangeMap
	stf     RangeMap
	ftw     RangeMap
	wtl     RangeMap
	ltt     RangeMap
	tth     RangeMap
	htl     RangeMap
}

type RangeMap struct {
	dest []int64
	src  []int64
	l    []int64
}

func main() {

	input, _ := ReadFileInput("input.txt")

	fmt.Println("Part 1:", Part1(input)) //3374647
	fmt.Println("Part 2:", Part2(input))
}

func Part1(input string) string {
	data := strings.Split(input, "\r\n\r\n")
	alamnac := &Almanac{}
	alamnac.seeds = ParseSeeds(data[0])
	alamnac.sts = ParseMap(data[1])
	alamnac.stf = ParseMap(data[2])
	alamnac.ftw = ParseMap(data[3])
	alamnac.wtl = ParseMap(data[4])
	alamnac.ltt = ParseMap(data[5])
	alamnac.tth = ParseMap(data[6])
	alamnac.htl = ParseMap(data[7])
	alamnac.SeedsToLocs()
	min := slices.Min(alamnac.seedLoc)
	return strconv.FormatInt(min, 10)
}

func Part2(input string) string {
	data := strings.Split(input, "\r\n\r\n")
	alamnac := &Almanac{}
	alamnac.seeds = ParseSeeds(data[0])
	alamnac.sts = ParseMap(data[1])
	alamnac.stf = ParseMap(data[2])
	alamnac.ftw = ParseMap(data[3])
	alamnac.wtl = ParseMap(data[4])
	alamnac.ltt = ParseMap(data[5])
	alamnac.tth = ParseMap(data[6])
	alamnac.htl = ParseMap(data[7])
	alamnac.SeedRangeToLocs()
	min := slices.Min(alamnac.seedLoc)
	return strconv.FormatInt(min, 10)
}

func ParseSeeds(seeds string) []int64 {
	seeds = strings.Split(seeds, ": ")[1]
	seeds_split := strings.Split(seeds, " ")
	var seedSlice []int64
	for _, seed := range seeds_split {
		s, _ := strconv.ParseInt(seed, 10, 64)
		seedSlice = append(seedSlice, s)
	}
	return seedSlice
}

func ParseMap(m string) RangeMap {
	lines := strings.Split(m, "\r\n")
	pm := &RangeMap{}
	for i := 1; i < len(lines); i++ {
		vals := strings.Split(lines[i], " ")
		dest, _ := strconv.ParseInt(vals[0], 10, 64)
		src, _ := strconv.ParseInt(vals[1], 10, 64)
		l, _ := strconv.ParseInt(vals[2], 10, 64)
		pm.dest = append(pm.dest, dest)
		pm.src = append(pm.src, src)
		pm.l = append(pm.l, l)
	}
	return *pm
}

func (a *Almanac) SeedsToLocs() {
	for _, seed := range a.seeds {
		a.seedLoc = append(a.seedLoc, a.seedToLoc(seed))
	}
}

func (a *Almanac) seedToLoc(c int64) int64 {
	c = a.sts.NextPos(c)
	c = a.stf.NextPos(c)
	c = a.ftw.NextPos(c)
	c = a.wtl.NextPos(c)
	c = a.ltt.NextPos(c)
	c = a.tth.NextPos(c)
	c = a.htl.NextPos(c)
	return c
}

func (p *RangeMap) NextPos(c int64) int64 {
	for i, src := range p.src {
		l := p.l[i]
		dest := p.dest[i]
		if c >= src && c <= src+l-1 {
			return dest + (c - src)
		}
	}
	return c
}

func (a *Almanac) SeedRangeToLocs() {
	for i := 0; i < len(a.seeds); i = i + 2 {
		seed := a.seeds[i]
		totalSeeds := a.seeds[i+1]
		minLoc := int64(^uint64(0) >> 1)
		for j := seed; j < seed+totalSeeds; j++ {
			loc := a.seedToLoc(j)
			if loc < minLoc {
				minLoc = loc
			}
		}
		a.seedLoc = append(a.seedLoc, minLoc)
	}
}

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
