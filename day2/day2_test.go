package main

import (
	"testing"
)

func TestPart1Sample(t *testing.T) {
	result := Part1("input0.txt")
	expected := 150
	if result != expected {
		t.Errorf("Incorrect Result. Expected: %d, got: %d.", expected, result)
	}
}

func TestPart1Actual(t *testing.T) {
	result := Part1("input.txt")
	expected := 1692075
	if result != expected {
		t.Errorf("Incorrect Result. Expected: %d, got: %d.", expected, result)
	}
}

func TestPart2Sample(t *testing.T) {
	result := Part2("input0.txt")
	expected := 900
	if result != expected {
		t.Errorf("Incorrect Result. Expected: %d, got: %d.", expected, result)
	}
}

func TestPart2Actual(t *testing.T) {
	result := Part2("input.txt")
	expected := 1749524700
	if result != expected {
		t.Errorf("Incorrect Result. Expected: %d, got: %d.", expected, result)
	}
}
