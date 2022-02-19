package main

import (
	"errors"
	"fmt"
)

func main() {
	data := loader("input.txt")
	Part1(data)
	Part2(data)
}

//lint:file-ignore SA9003 Ignore empty branches

func Part1(data []int) int {
	increased_count := 0

	for i, v := range data {
		// assessment := "N/A - no previous measurement"
		if i != 0 {
			if prev := data[i-1]; v > prev {
				// assessment = "increased"
				increased_count++
			} else if v == prev {
				// assessment = "no change"
			} else {
				// assessment = "decreased"
			}
		}
		// fmt.Printf("i=%d v=%d %s \n", i, v, assessment)
	}
	fmt.Printf("Part 1 - Total Increased Count: %d\n", increased_count)
	return increased_count
}

func Part2(data []int) (int, error) {
	increased_count := 0

	// Cancel if less than 3
	if len(data) < 3 {
		return 0, errors.New("less than 3 measurements in input")
	}

	prev := 0

	for i := 2; i < len(data); i++ {
		v := data[i] + data[i-1] + data[i-2]
		// assessment := "N/A - no previous measurement"
		if i != 2 {
			if v > prev {
				// assessment = "increased"
				increased_count++
			} else if v == prev {
				// assessment = "no change"
			} else {
				// assessment = "decreased"
			}
		}
		prev = v
		// fmt.Printf("i=%d v=%d %s \n", i, v, assessment)
	}

	fmt.Printf("Part 2 - Total Increased Count: %d\n", increased_count)
	return increased_count, nil
}
