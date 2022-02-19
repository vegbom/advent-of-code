package main

import (
	"encoding/hex"
	"fmt"
)

type ParseState int

const (
	ExpectPacketVersion      ParseState = 0
	ExpectPacketType         ParseState = 1
	ExpectLenType            ParseState = 2
	ExpectLenBits            ParseState = 3
	ExpectLenNumPacks        ParseState = 4
	ExpectLiteralValue       ParseState = 5
	ExpectSubpacksByBit      ParseState = 6
	ExpectSubpacksByNumPacks ParseState = 7
	ExpectEndOfPacket        ParseState = 99
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

var verSum int = 0

func main() {
	// decode("D2FE28")
	// decode("38006F45291200")
	// decode("EE00D40C823060")

	decode(puzzleInput)
	fmt.Printf("verSum=%d\n", verSum) //852
}

func mask(offset int) uint64 {
	return (1 << offset) - 1
}

func GetBits(loc int, sz int, data []byte) (result uint64) {
	byteNumStart := loc / 8
	byteNumEnd := (loc + sz) / 8
	bitOffsetStart := 8 - loc%8

	byteNum := byteNumStart
	for byteNum <= byteNumEnd {
		var c uint64
		if byteNum == byteNumStart {
			c = uint64(data[byteNum]) & mask(bitOffsetStart)
		} else {
			c = uint64(data[byteNum])
		}
		result = result << 8
		result |= c
		// fmt.Printf("i=%d c=%b res=%b\n", byteNum, c, result)
		byteNum += 1
	}
	// Trim the end
	result = result >> ((byteNum-byteNumStart)*8 - (8 - bitOffsetStart) - sz)

	return result
}

func decode(in string) {
	data, err := hex.DecodeString(in)
	if err != nil {
		panic(err)
	}
	fmt.Printf("% x\n", data)

	// fmt.Printf("VVV=%b\n", GetBits(0, 3))
	// fmt.Printf("TTT=%b\n", GetBits(3, 3))
	// fmt.Printf("AAAAA=%b\n", GetBits(6, 5))
	// fmt.Printf("BBBBB=%b\n", GetBits(6+5, 5))
	// fmt.Printf("CCCCC=%b\n", GetBits(6+5+5, 5))

	// actual decode part
	ParsePack(0, data)
}

func ParsePack(bit int, data []byte) int {
	fmt.Printf("BEGIN SUBPACK from bit %d\n", bit)
	state := ExpectPacketVersion
	subpackByBitLen := 0
	subpackByBitStart := 0
	subpackByNumLen := 0
	subpackByNumCurrent := 0
	var litValue uint64 = 0
	for bit < len(data)*8 {
		switch state {
		case ExpectPacketVersion:
			packetVer := GetBits(bit, 3, data)
			fmt.Printf("Packet Ver=%d ", packetVer)
			verSum += int(packetVer)
			bit += 3
			state = ExpectPacketType
		case ExpectPacketType:
			packetType := GetBits(bit, 3, data)
			fmt.Printf("Type=%d ", packetType)
			bit += 3
			if packetType == 4 {
				// Literal Value
				litValue = 0
				state = ExpectLiteralValue
			} else {
				// Operator
				state = ExpectLenType
			}
		case ExpectLiteralValue:
			v := GetBits(bit, 5, data)
			litValue |= v & mask(4)
			bit += 5
			if v>>4 == 0 {
				fmt.Printf("Literal Value=%d\n", litValue)
				state = ExpectEndOfPacket
			} else {
				litValue = litValue << 4
			}
		case ExpectLenType:
			v := GetBits(bit, 1, data)
			bit++
			if v == 0 {
				state = ExpectLenBits
			} else {
				state = ExpectLenNumPacks
			}
		case ExpectLenBits:
			subpackByBitLen = int(GetBits(bit, 15, data))
			fmt.Printf("Subpacket Len Bits=%d\n", subpackByBitLen)
			bit += 15
			subpackByBitStart = bit
			state = ExpectSubpacksByBit
		case ExpectLenNumPacks:
			subpackByNumLen = int(GetBits(bit, 11, data))
			fmt.Printf("Subpacket Len Num=%d\n", subpackByNumLen)
			bit += 11
			subpackByNumCurrent = 0
			state = ExpectSubpacksByNumPacks
		case ExpectSubpacksByBit:
			bit = ParsePack(bit, data)
			if bit-subpackByBitStart >= subpackByBitLen {
				state = ExpectEndOfPacket
			}
		case ExpectSubpacksByNumPacks:
			bit = ParsePack(bit, data)
			subpackByNumCurrent++
			if subpackByNumCurrent >= subpackByNumLen {
				state = ExpectEndOfPacket
			}
		case ExpectEndOfPacket:
			fmt.Printf("END SUBPACK at bit %d\n", bit)
			return bit
		}
	}
	return bit
}

func SelectPacketType(t int) ParseState {
	switch t {
	case 4:
		// Literal Value
		return ExpectLiteralValue
	case 6:
		return ExpectLenType
	case 3:
		return ExpectLenType
	}
	return ExpectLiteralValue
	//TODO
}
