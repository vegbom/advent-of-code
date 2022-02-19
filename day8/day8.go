package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strings"
)

const Reset = "\033[0m"
const Red = "\033[31m"
const Green = "\033[32m"
const Yellow = "\033[33m"

func main() {
	patterns, outputs, err := Loader("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	_ = patterns
	_ = outputs

	fmt.Printf("Part 1: digits 1, 4, 7, or 8 appeared %d times\n", Part1(outputs))
	fmt.Printf("Part 2: Sum: %d\n", Part2(patterns, outputs))
}

func Part2(patterns [][]string, outputs [][]string) (count int) {

	for i, pattern := range patterns {
		segments := SecondPass(FirstPass(pattern))

		output := outputs[i]
		numericalValue := 0
		for pow, digit := range output {
			outInt := render(digit, segments)
			fmt.Printf("%v ", outInt)
			numericalValue += outInt * int(math.Pow(10, float64(3-pow)))
		}
		fmt.Printf("N=%d\n", numericalValue)
		count += numericalValue
	}
	return count
}

func FirstPass(patternsOberved []string) (knownDigits map[int]string, unkLength5 []string) {
	knownDigits = make(map[int]string)
	for _, pattern := range patternsOberved {
		switch len(pattern) {
		case 2:
			knownDigits[1] = pattern
		case 4:
			knownDigits[4] = pattern
		case 3:
			knownDigits[7] = pattern
		case 7:
			knownDigits[8] = pattern
		case 5:
			unkLength5 = append(unkLength5, pattern)
		}
	}
	return knownDigits, unkLength5
}

func SecondPass(knownDigits map[int]string, unkLength5 []string) map[string]string {
	segment := make(map[string]string)

	regionI := UncertaintyRegion{}
	regionI.add(string(knownDigits[1][0]))
	regionI.add(string(knownDigits[1][1]))
	// fmt.Printf("regionI: %v ", regionI)

	regionII := UncertaintyRegion{}
	for _, letter := range knownDigits[4] {
		if !regionI.contains(string(letter)) {
			regionII.add(string(letter))
		}
	}
	// fmt.Printf("regionII: %v ", regionII)

	for _, letter := range knownDigits[7] {
		if !regionI.contains(string(letter)) && !regionII.contains(string(letter)) {
			segment["a"] = string(letter)
			break
		}
	}
	// fmt.Printf("Mapped Segment A: %v ", segment["a"])

	regionIII := UncertaintyRegion{}
	for _, letter := range knownDigits[8] {
		if !regionI.contains(string(letter)) && !regionII.contains(string(letter)) && string(letter) != segment["a"] {
			regionIII.add(string(letter))
		}
	}
	// fmt.Printf("regionIII: %v ", regionIII)

	// Find who is 3
	for _, pattern := range unkLength5 {
		regionCount := make([]int, 3)
		for _, letter := range pattern {
			if regionI.contains(string(letter)) {
				regionCount[0]++
			} else if regionII.contains(string(letter)) {
				regionCount[1]++
			} else if regionIII.contains(string(letter)) {
				regionCount[2]++
			}
		}
		if regionCount[0] == 2 {
			knownDigits[3] = pattern
		} else if regionCount[2] == 2 {
			knownDigits[2] = pattern
		} else {
			knownDigits[5] = pattern
		}
	}
	// Disambiguate Regions II and III using 3
	for _, letter := range knownDigits[3] {
		if regionII.contains(string(letter)) {
			segment["d"] = string(letter)
			segment["b"] = regionII.getOther(string(letter))
		}
		if regionIII.contains(string(letter)) {
			segment["g"] = string(letter)
			segment["e"] = regionIII.getOther(string(letter))
		}
	}
	// Disambiguate Regions I using 2
	for _, letter := range knownDigits[2] {
		if regionI.contains(string(letter)) {
			segment["c"] = string(letter)
			segment["f"] = regionI.getOther(string(letter))
		}
	}

	// Invert it
	segmentInverted := make(map[string]string)
	for k, v := range segment {
		segmentInverted[v] = k
	}
	// fmt.Printf("\nsegment: %v \n inverted: %v", segment, segmentInverted)

	return segmentInverted
}

func render(observedPattern string, segmentMap map[string]string) int {
	corrected := make([]string, 0)
	for _, letter := range observedPattern {
		corrected = append(corrected, segmentMap[string(letter)])
	}
	sort.Strings(corrected)

	switch strings.Join(corrected, "") {
	case "abcefg":
		return 0
	case "cf":
		return 1
	case "acdeg":
		return 2
	case "acdfg":
		return 3
	case "bcdf":
		return 4
	case "abdfg":
		return 5
	case "abdefg":
		return 6
	case "acf":
		return 7
	case "abcdefg":
		return 8
	case "abcdfg":
		return 9
	}
	return -1
}

func Part1(outputs [][]string) (count int) {
	for _, display := range outputs {
		for _, digit := range display {
			if isUnique(digit) > 0 {
				fmt.Printf("%s%v%s ", Green, digit, Reset)
				count++
			} else {
				fmt.Printf("%v ", digit)
			}
		}
		fmt.Printf("\n")
	}
	return count
}

func isUnique(s string) int {
	switch len(s) {
	case 2:
		return 1
	case 4:
		return 4
	case 3:
		return 7
	case 7:
		return 8
	default:
		return -1
	}
}

func Loader(filename string) (patterns [][]string, outputs [][]string, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	patterns = make([][]string, 0)
	outputs = make([][]string, 0)

	for scanner.Scan() {
		patterns = append(patterns, strings.Fields(strings.Split(scanner.Text(), "|")[0]))
		outputs = append(outputs, strings.Fields(strings.Split(scanner.Text(), "|")[1]))
	}

	return patterns, outputs, err
}

type UncertaintyRegion struct {
	item1 string
	item2 string
}

func (r *UncertaintyRegion) add(s string) {
	if len(r.item1) == 0 {
		r.item1 = s
	} else {
		r.item2 = s
	}
}

func (r UncertaintyRegion) contains(s string) bool {
	if r.item1 == s || r.item2 == s {
		return true
	}
	return false
}

func (r UncertaintyRegion) getOther(s string) string {
	if r.item1 == s {
		return r.item2
	}
	return r.item1
}
