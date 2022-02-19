package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func loader(filename string) []int {
	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	// Make new Slice to store the data
	data := make([]int, 0)

	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}
		// fmt.Println("scanner.Text(): " + scanner.Text() + " int: " + strconv.FormatInt(int64(i), 10))
		data = append(data, i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}
