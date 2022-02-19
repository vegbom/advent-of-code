package main

import (
	"testing"
)

func TestPart1Sample(t *testing.T) {
	result := Part1(loader("input0.txt"))
	expected := 7
	if result != expected {
		t.Errorf("Incorrect Result. Expected: %d, got: %d.", expected, result)
	}
}

func TestPart1Actual(t *testing.T) {
	result := Part1(loader("input.txt"))
	expected := 1387
	if result != expected {
		t.Errorf("Incorrect Result. Expected: %d, got: %d.", expected, result)
	}
}

func TestPart2Sample(t *testing.T) {
	result, err := Part2(loader("input0.txt"))
	expected := 5
	if result != expected || err != nil {
		t.Errorf("Incorrect Result. Expected: %d, got: %d.", expected, result)
	}
}

func TestPart2Actual(t *testing.T) {
	result, err := Part2(loader("input.txt"))
	expected := 1362
	if result != expected || err != nil {
		t.Errorf("Incorrect Result. Expected: %d, got: %d.", expected, result)
	}
}

func TestPart2_NotEnoughData(t *testing.T) {
	data := []int{1, 2}
	result, err := Part2(data)
	if result == 0 && err == nil {
		t.Errorf("Incorrect Result. Should have failed with error.")
	}
}
