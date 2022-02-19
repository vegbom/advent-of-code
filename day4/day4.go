package main

// Learning objectives: Classes

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"vegbom/day4/bingoboard"
)

type ParseState int

const (
	ExpectNumbersCalled   ParseState = 0
	ExpectBlankLine       ParseState = 1
	ExpectFirstBoard      ParseState = 2
	ExpectSubsequentBoard ParseState = 3
)

func main() {
	numbers_called, boards, err := Loader("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	const is_part_1 bool = false

	// Is it possible that multiple boards could win at the same time?
	// Assuming that only one board can win per number called.
	for _, number := range numbers_called {
		for board_id, board := range boards {
			if board.Call(number) {
				board.PrintBoard()
				sum := board.GetUnmarkedSum()
				fmt.Printf("Called %d , Board #%d won!\n", number, board_id)
				fmt.Printf("sum_unmarked * number_just_called = score\n")
				fmt.Printf("%d * %d = %d\n", sum, number, sum*number)
				if is_part_1 {
					// PART 1 Answer: 741 * 14 = 10374
					return
				}
				// Reset this board
				boards[board_id] = bingoboard.New()
			}
		}
	}
	// PART 2 Answer: 278 * 89 = 24742
}

func Loader(filename string) (numbers_called []int, boards []bingoboard.Bingoboard, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	boards = make([]bingoboard.Bingoboard, 0)

	numbers_called = nil
	current_row := -1
	var current_board bingoboard.Bingoboard
	state := ExpectNumbersCalled

	for scanner.Scan() {
		switch state {
		case ExpectNumbersCalled:
			for _, n := range strings.Split(scanner.Text(), ",") {
				i, err := strconv.Atoi(n)
				if err != nil {
					return nil, nil, err
				}
				numbers_called = append(numbers_called, i)
			}
			fmt.Printf("numbers_called: %v\n", numbers_called)
			state = ExpectBlankLine
		case ExpectBlankLine:
			if strings.TrimSpace(scanner.Text()) == "" {
				state = ExpectFirstBoard
			} else {
				log.Fatalf("Expected Blank Line, got %s", scanner.Text())
			}
		case ExpectFirstBoard:
			current_row = 0
			current_board = bingoboard.New()
			state = ExpectSubsequentBoard
			fallthrough
		case ExpectSubsequentBoard:
			row_data := make([]int, 0)
			for _, n := range strings.Fields(scanner.Text()) {
				i, err := strconv.Atoi(n)
				if err != nil {
					return nil, nil, err
				}
				row_data = append(row_data, i)
			}
			current_board.SetRow(row_data, current_row)
			current_row++
			if current_row == bingoboard.BOARD_SIZE {
				state = ExpectBlankLine
				boards = append(boards, current_board)
			}
		}
	}

	return numbers_called, boards, err
}
