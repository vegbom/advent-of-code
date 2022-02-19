package main

import (
	"encoding/hex"
	"fmt"
	"math"
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
	ExpectEndOfPacket        ParseState = 8
)

type Op int

const (
	Sum         Op = 0
	Product     Op = 1
	Minimum     Op = 2
	Maximum     Op = 3
	Literal     Op = 4
	GreaterThan Op = 5
	LessThan    Op = 6
	EqualTo     Op = 7
	Undefined   Op = -1
)

var verSum int = 0

func main() {
	ans := decode(puzzleInput)
	fmt.Printf("Part 1: Version Sum=%d\n", verSum) //852
	fmt.Printf("Part 2: Answer=%d\n", ans)         //19348959966392
}

func mask(offset int) uint64 {
	return (1 << offset) - 1
}

func GetBits(loc int, sz int, data []byte) (result uint64) {
	byteNumStart := loc / 8
	byteNumEnd := (loc + (sz - 1)) / 8
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

func decode(in string) (ans int) {
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

	_, ans = ParsePack(0, data)
	return ans
}

func ParsePack(bit int, data []byte) (int, int) {
	fmt.Printf("BEGIN PACK from bit %d\n", bit)
	state := ExpectPacketVersion
	subpackByBitLen := 0
	subpackByBitStart := 0
	subpackByNumLen := 0
	subpackByNumCurrent := 0
	operands := make([]int, 0)
	var operation Op = Undefined
	result := 0
	var litValue uint64 = 0
	for {
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
			operation = Op(packetType)
			if operation == Literal {
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
				operands = []int{int(litValue)}
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
			bit, result = ParsePack(bit, data)
			operands = append(operands, result)
			if bit-subpackByBitStart >= subpackByBitLen {
				state = ExpectEndOfPacket
			}
		case ExpectSubpacksByNumPacks:
			bit, result = ParsePack(bit, data)
			operands = append(operands, result)
			subpackByNumCurrent++
			if subpackByNumCurrent >= subpackByNumLen {
				state = ExpectEndOfPacket
			}
		case ExpectEndOfPacket:
			fmt.Printf("END PACK at bit %d, result %d \n", bit, Calculate(operation, operands))
			return bit, Calculate(operation, operands)
		}
	}
}

func Calculate(operation Op, operands []int) int {
	switch operation {
	case Sum:
		s := 0
		for _, v := range operands {
			s += v
		}
		return s
	case Product:
		s := 0
		for i, v := range operands {
			if i == 0 {
				s = v
			} else {
				s *= v
			}
		}
		return s
	case Minimum:
		s := math.MaxInt
		for _, v := range operands {
			if v < s {
				s = v
			}
		}
		return s
	case Maximum:
		s := 0
		for _, v := range operands {
			if v > s {
				s = v
			}
		}
		return s
	case GreaterThan:
		if operands[0] > operands[1] {
			return 1
		}
		return 0
	case LessThan:
		if operands[0] < operands[1] {
			return 1
		}
		return 0
	case EqualTo:
		if operands[0] == operands[1] {
			return 1
		}
		return 0
	default:
		// incl Literal Value
		return operands[0]
	}
}
