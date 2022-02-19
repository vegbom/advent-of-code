package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Direction uint8

const (
	MostCommon  Direction = 0
	LeastCommon Direction = 1
)

func main() {
	data, word_sz, err := Loader("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Part1(data, word_sz)
	o2_rating := Part2(data, word_sz, MostCommon)
	co2_rating := Part2(data, word_sz, LeastCommon)
	fmt.Printf("Oxygen Rating: %d\nCO2 Rating: %d\nCombined Life Rating: %d", o2_rating, co2_rating, o2_rating*co2_rating)
}

func Loader(filename string) (data []uint64, line_len int, err error) {
	f, err := os.Open(filename)

	if err != nil {
		return nil, 0, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	data = make([]uint64, 0)

	for scanner.Scan() {
		// Input Validation
		if line_len == 0 {
			line_len = len(scanner.Text())
			if line_len > 64 {
				return nil, 0, errors.New("input line too long")
			}
			fmt.Printf("Line Length is %d\n", line_len)
		} else if line_len != len(scanner.Text()) {
			return nil, 0, errors.New("input line lengths inconsistent")
		}

		i, err := strconv.ParseUint(scanner.Text(), 2, line_len)
		if err != nil {
			fmt.Println(err)
		}
		data = append(data, i)
		// fmt.Printf("Parsed %s as 0b%b\n", scanner.Text(), i)
		// fmt.Println("scanner.Text(): " + scanner.Text() + " int: " + strconv.FormatInt(int64(i), 10))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data, line_len, nil
}

func Part1(data []uint64, word_sz int) (uint64, error) {

	bit_count := make([]int, word_sz)
	for _, v := range data {
		for j := 0; j < word_sz; j++ {
			if ((v >> j) & 1) == 1 {
				bit_count[j]++
			} else {
				bit_count[j]--
			}
		}
	}

	fmt.Printf("bit_count: %v\n", bit_count)

	// Generate the answer
	var gamma uint64 = 0
	for i := 0; i < word_sz; i++ {
		if bit_count[i] > 0 {
			gamma = gamma | (1 << i)
		} else if bit_count[i] == 0 {
			return 0, errors.New("bit is indeterminate")
		}
	}

	bitmask := (uint64)(1<<64-1) >> (64 - word_sz)
	var epsilon uint64 = ^gamma
	epsilon = epsilon & bitmask
	fmt.Printf("gamma: %d 0b%b epsilon: %d 0b%b\n", gamma, gamma, epsilon, epsilon)
	fmt.Printf("Power Consumption: %d\n", gamma*epsilon)

	return (gamma * epsilon), nil
}

func Part2(data []uint64, word_sz int, dir Direction) (answer uint64) {
	// Starting from left-most bit to the right-most bit
	for j := word_sz - 1; j >= 0; j-- {
		bit_count := 0 // bit_count is positive when there are more ones than zeroes
		accepted_sz := 0
		current_bit := false

		for _, v := range data {
			// Skip numbers that don't match known answer bits
			if (v >> (j + 1)) != (answer >> (j + 1)) {
				continue
			}
			accepted_sz++
			// fmt.Printf("[j=%d i=%d] Accepted %b\n", j, i, v)

			if ((v >> j) & 1) == 1 {
				bit_count++
			} else {
				bit_count--
			}
		}

		// Assume dir == MostCommon
		if bit_count > 0 {
			// More ones than zeroes
			current_bit = true
		} else if bit_count == 0 {
			current_bit = true
		} else {
			// More zeroes than ones
			current_bit = false
		}

		// NOT the whole thing except when we only have one left
		// in which case it's the same as MostCommon
		if dir == LeastCommon && accepted_sz > 1 {
			current_bit = !current_bit
		}

		if current_bit {
			answer |= 1 << j
		}
		// fmt.Printf("End of j=%d bit_count = %d answer = %b \n", j, bit_count, answer)
	}

	return answer
}
