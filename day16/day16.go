package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	decode("D2FE28")
}

func decode(in string) {

	data, err := hex.DecodeString(in)
	if err != nil {
		panic(err)
	}
	fmt.Printf("% x", data)

	getBit := func(loc int, sz int) (result uint64) {
		byteNum := loc / 8
	}

}
