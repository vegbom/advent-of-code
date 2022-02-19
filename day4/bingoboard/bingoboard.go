package bingoboard

import (
	"fmt"
	"log"
)

const BOARD_SIZE int = 5

type Bingoboard struct {
	is_called [][]bool
	numbers   [][]int
}

func New() Bingoboard {
	is_called := make([][]bool, BOARD_SIZE)
	numbers := make([][]int, BOARD_SIZE)
	for i := range is_called {
		is_called[i] = make([]bool, BOARD_SIZE)
		numbers[i] = make([]int, BOARD_SIZE)
	}

	b := Bingoboard{is_called, numbers}
	return b
}

func (b Bingoboard) SetRow(data []int, row_num int) {
	if len(data) != BOARD_SIZE {
		log.Fatal("Expected setRow to be called with data size of 5")
	}
	b.numbers[row_num] = data
}

func (b Bingoboard) Call(number_called int) bool {
	for i := range b.numbers {
		for j := range b.numbers[i] {
			if b.numbers[i][j] == number_called {
				b.is_called[i][j] = true
			}
		}
	}

	// Check if bingo
	for i := range b.is_called {
		h := 0 // Horizontal
		v := 0 // Vertical
		for j := 0; j < BOARD_SIZE; j++ {
			if b.is_called[i][j] {
				h++
			}
			if b.is_called[j][i] {
				v++
			}
		}
		if h == BOARD_SIZE || v == BOARD_SIZE {
			fmt.Printf("BINGO!\n")
			return true
		}
	}
	return false
}

func (b Bingoboard) GetUnmarkedSum() int {
	sum := 0
	for i := range b.is_called {
		for j := range b.is_called[i] {
			if !b.is_called[i][j] {
				sum += b.numbers[i][j]
			}
		}
	}
	return sum
}

func (b Bingoboard) PrintBoard() {
	for i := range b.numbers {
		fmt.Printf("%v\n", b.numbers[i])
	}
}
