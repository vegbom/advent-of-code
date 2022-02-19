package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	template, rules := Loader("input.txt")

	// Part 1
	polymer1 := template
	for step := 1; step <= 10; step++ {
		polymer1 = Expand(polymer1, rules)
		// fmt.Printf("After Step %d: %s\n", step, polymer1)
	}
	PolymerStats(polymer1)

	// Part 2
	Expand2(template, rules, 40)
}

func PolymerStats(polymer string) {
	elements := make(map[string]int)
	for _, char := range polymer {
		elements[string(char)]++
	}

	max := 0
	min := -1
	for key, value := range elements {
		fmt.Printf("Count of %s: %d\n", key, value)
		if value < min || min < 0 {
			min = value
		}
		if value > max {
			max = value
		}
	}

	fmt.Printf("PART 1: Most Common - Least Common: %d - %d = %d\n", max, min, max-min)
}

func Expand(before string, rules map[string]map[string]string) (after string) {

	for i := 0; i < len(before)-1; i++ {
		h := string(before[i])
		t := string(before[i+1])

		if val, ok := rules[h][t]; ok {
			fmt.Printf("Found Rule: %s%s -> %s\n", h, t, val)
			after += h
			after += val
		} else {
			fmt.Printf("WARN: No Rule for %s%s\n", h, t)
			after += h
		}
	}

	after += string(before[len(before)-1])

	return after
}

func Expand2(init string, rules map[string]map[string]string, steps int) {
	statLetter := make(map[string]int)
	statPair := make(map[string]int)
	for _, char := range init {
		statLetter[string(char)]++
	}

	for i := 0; i < len(init)-1; i++ {
		h := string(init[i])
		t := string(init[i+1])

		k := h + t
		statPair[k]++
	}

	for step := 0; step < steps; step++ {
		statPair = Grow(rules, statLetter, statPair)
	}

	max := 0
	min := -1
	for key, value := range statLetter {
		fmt.Printf("Count of %s: %d\n", key, value)
		if value < min || min < 0 {
			min = value
		}
		if value > max {
			max = value
		}
	}

	fmt.Printf("PART 2: Most Common - Least Common: %d - %d = %d\n", max, min, max-min)
}

func Grow(rules map[string]map[string]string, statLetter map[string]int, statPair map[string]int) map[string]int {

	statPairCurrent := make(map[string]int)

	for key, count := range statPair {
		h := string(key[0])
		t := string(key[1])
		if insert, ok := rules[h][t]; ok {
			spawnPair1 := h + insert
			spawnPair2 := insert + t
			statLetter[insert] += count
			statPairCurrent[spawnPair1] += count
			statPairCurrent[spawnPair2] += count
		} else {
			log.Fatalf("Wrong: No Rule for %s%s\n", h, t)
		}
	}

	return statPairCurrent

}

func Loader(filename string) (string, map[string]map[string]string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	rules := make(map[string]map[string]string)
	template := ""

	scanTemplate := true
	r := regexp.MustCompile(`(?P<head>[a-zA-Z])(?P<tail>[a-zA-Z])\s->\s(?P<insert>[a-zA-Z])`)

	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "" {
			scanTemplate = false
			continue
		}

		if scanTemplate {
			template = scanner.Text()
		} else {
			fmt.Printf("INSTR: %v\n", scanner.Text())
			matches := r.FindStringSubmatch(scanner.Text())
			head := matches[r.SubexpIndex("head")]
			tail := matches[r.SubexpIndex("tail")]
			insert := matches[r.SubexpIndex("insert")]

			if rules[head] == nil {
				rules[head] = make(map[string]string)
			}
			rules[head][tail] = insert
		}
	}

	return template, rules
}
