package main

import (
	"encoding/hex"
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	decode("D2FE28")
}

func decode(in string) {

	data, err := hex.DecodeString(in)
	if err != nil {
		panic(err)
	}
	fmt.Printf("% x\n", data)

	getBit := func(loc int, sz int) (result uint64) {
		byteNumStart := loc / 8
		byteNumEnd := (loc + sz) / 8
		bitOffsetStart := 8 - loc%8

		byteNum := byteNumStart
		for byteNum <= byteNumEnd {
			var c uint64
			if byteNum == byteNumStart {
				var mask uint8 = (1 << bitOffsetStart) - 1
				c = uint64(data[byteNum] & mask)
			} else {
				c = uint64(data[byteNum])
			}
			result = result << 8
			result |= c
			fmt.Printf("i=%d c=%b res=%b\n", byteNum, c, result)
			// i += min(sz-i, 8)
			byteNum += 1
		}
		// Trim the end
		fmt.Printf("trim amount=%d\n", byteNum*8-(8-bitOffsetStart)-sz)
		// var mask uint64 = (1 ^ (1 << (i - sz)) - 1)
		result = result >> ((byteNum-byteNumStart)*8 - (8 - bitOffsetStart) - sz)

		return result
	}

	fmt.Printf("VVV=%b\n", getBit(0, 3))
	fmt.Printf("TTT=%b\n", getBit(3, 3))
	fmt.Printf("AAAAA=%b\n", getBit(6, 5))
	fmt.Printf("BBBBB=%b\n", getBit(6+5, 5))
	fmt.Printf("CCCCC=%b\n", getBit(6+5+5, 5))
}
