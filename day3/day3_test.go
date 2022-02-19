package main

import (
	"fmt"
	"testing"
)

func TestLoader_BadInputs(t *testing.T) {
	_, _, err0 := Loader("testbadinput0.txt")
	if err0 == nil {
		t.Errorf("Incorrect Result. Expected Error.")
	}
	fmt.Println(err0)
	_, _, err1 := Loader("testbadinput1.txt")
	if err1 == nil {
		t.Errorf("Incorrect Result. Expected Error.")
	}
	fmt.Println(err1)
	_, _, err2 := Loader("testbadinput2.txt")
	if err2 == nil {
		t.Errorf("Incorrect Result. Expected Error.")
	}
	fmt.Println(err2)
}

func TestLoader_CorrectInput0(t *testing.T) {
	_, result, err := Loader("input0.txt")
	expected := 5
	if result != expected || err != nil {
		t.Errorf("Incorrect Result. Expected: %d, got: %d.", expected, result)
	}
}

func TestLoader_CorrectInput1(t *testing.T) {
	_, result, err := Loader("input.txt")
	expected := 12
	if result != expected || err != nil {
		t.Errorf("Incorrect Result. Expected: %d, got: %d.", expected, result)
	}
}

func TestPart1(t *testing.T) {
	data, word_sz, err := Loader("input.txt")
	if err != nil {
		t.Errorf("Unrelated Error")
	}

	result, err1 := Part1(data, word_sz)
	var expected uint64 = 3277364
	if result != expected || err1 != nil {
		t.Errorf("Incorrect Result. Expected: %d, got: %d.", expected, result)
	}
}

func TestPart1_IndeterminateInput(t *testing.T) {
	data, word_sz, err := Loader("testintederminateinput.txt")
	if err != nil {
		t.Errorf("Unrelated Error")
	}

	result, err1 := Part1(data, word_sz)
	var expected uint64 = 0
	if err1 == nil || result != expected {
		t.Errorf("Incorrect Result. Expected Error.")
	}
	fmt.Println(err1)
}

func TestPart2_O2(t *testing.T) {
	data, word_sz, err := Loader("input.txt")
	if err != nil {
		t.Errorf("Unrelated Error")
	}

	result := Part2(data, word_sz, MostCommon)
	var expected uint64 = 3583
	if result != expected {
		t.Errorf("Incorrect Result. Expected: %d, got: %d.", expected, result)
	}
}

func TestPart2_CO2(t *testing.T) {
	data, word_sz, err := Loader("input.txt")
	if err != nil {
		t.Errorf("Unrelated Error")
	}

	result := Part2(data, word_sz, LeastCommon)
	var expected uint64 = 1601
	if result != expected {
		t.Errorf("Incorrect Result. Expected: %d, got: %d.", expected, result)
	}
}
