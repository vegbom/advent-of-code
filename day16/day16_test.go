package main

import (
	"testing"
)

func TestExampleOperations(t *testing.T) {
	ans := 0
	ans = decode("C200B40A82")
	if ans != 3 {
		t.Errorf("Incorrect Result.")
	}
	ans = decode("04005AC33890")
	if ans != 54 {
		t.Errorf("Incorrect Result.")
	}
	ans = decode("880086C3E88112")
	if ans != 7 {
		t.Errorf("Incorrect Result.")
	}
	ans = decode("CE00C43D881120")
	if ans != 9 {
		t.Errorf("Incorrect Result.")
	}
	ans = decode("D8005AC2A8F0")
	if ans != 1 {
		t.Errorf("Incorrect Result.")
	}
	ans = decode("F600BC2D8F")
	if ans != 0 {
		t.Errorf("Incorrect Result.")
	}
	ans = decode("9C005AC2F8F0")
	if ans != 0 {
		t.Errorf("Incorrect Result.")
	}
	ans = decode("9C0141080250320F1802104A08")
	if ans != 1 {
		t.Errorf("Incorrect Result.")
	}
}
