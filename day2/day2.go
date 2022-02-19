package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Part 1:")
	Part1("input.txt")
	fmt.Println("Part 2:")
	Part2("input.txt")
}

func Part1(filename string) int {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	position_v := 0
	position_h := 0

	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		instruction := words[0]
		argument, err := strconv.Atoi(words[1])
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println("scanner.Text(): " + scanner.Text() + " instruction: " + words[0] + " value: " + strconv.FormatInt(int64(movement), 10))
		if instruction == "forward" {
			position_h += argument
		} else if instruction == "down" {
			position_v += argument
		} else if instruction == "up" {
			position_v -= argument
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("H: %d\tV: %d\t Multiply: %d\n", position_h, position_v, position_h*position_v)
	// H: 1925	V: 879	 Multiply: 1692075
	return position_h * position_v
}

func Part2(filename string) int {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	position_v := 0
	position_h := 0
	aim := 0

	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		i, err := strconv.Atoi(words[1])
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println("scanner.Text(): " + scanner.Text() + " action: " + words[0] + " value: " + strconv.FormatInt(int64(i), 10))
		if words[0] == "forward" {
			position_h += i
			position_v += aim * i
		} else if words[0] == "down" {
			aim += i
		} else if words[0] == "up" {
			aim -= i
		}

		// fmt.Printf("%-8s %d -> H: %d\tV: %d\tAIM: %d\n", words[0], i, position_h, position_v, aim)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("H: %d\tV: %d\t Multiply: %d\n", position_h, position_v, position_h*position_v)
	// H: 1925	V: 908844	 Multiply: 1749524700
	return position_h * position_v
}
