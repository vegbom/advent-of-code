package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type stack []string

func (s *stack) Push(v string) {
	*s = append(*s, v)
}

func (s *stack) Pop() string {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

func main() {
	err := Loader("input.txt")
	if err != nil {
		log.Fatal(err)
	}
}

func parse(line string) (int, int) {
	s := stack{}
	part1Score := 0

	for _, c := range line {
		currentToken := string(c)
		if isOpen(currentToken) {
			s.Push(currentToken)
		} else {
			actualOpen := s.Pop()
			expected := getMatchingClose(actualOpen)
			if currentToken != expected {
				// fmt.Printf("%s - At:%d Expected %s Got %s\n", line, i, expected, currentToken)
				part1Score += getPart1Score(currentToken)
				break
			}
		}
	}

	part2Score := 0
	for len(s) > 0 && part1Score == 0 {
		actualOpen := s.Pop()
		close := getMatchingClose(actualOpen)
		// fmt.Printf("%s", close)
		part2Score = part2Score*5 + getPart2Score(close)
	}

	if part2Score > 0 {
		fmt.Printf("%s - INCOMPLETE, SCORE: %d \n", line, part2Score)
	}

	return part1Score, part2Score
}

func getPart1Score(closeChar string) int {
	switch closeChar {
	case ")":
		return 3
	case "]":
		return 57
	case "}":
		return 1197
	case ">":
		return 25137
	}
	return 0
}

func getPart2Score(closeChar string) int {
	switch closeChar {
	case ")":
		return 1
	case "]":
		return 2
	case "}":
		return 3
	case ">":
		return 4
	}
	return 0
}

func getMatchingClose(openChar string) string {
	switch openChar {
	case "(":
		return ")"
	case "[":
		return "]"
	case "{":
		return "}"
	case "<":
		return ">"
	}
	return ""
}

func isOpen(char string) bool {
	if char == "(" || char == "[" || char == "{" || char == "<" {
		return true
	}
	return false
}

func Loader(filename string) (err error) {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	part1 := 0
	part2 := make([]int, 0)
	for scanner.Scan() {
		p1, p2 := parse(scanner.Text())
		part1 += p1
		if p2 > 0 {
			part2 = append(part2, p2)
		}
	}

	sort.Ints(part2)
	part2Median := part2[(len(part2))/2]

	fmt.Printf("PART 1: %d PART 2: %d\n", part1, part2Median)

	return nil
}
