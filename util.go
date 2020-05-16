package rsdic

import (
	"fmt"
)

func floor(num uint64, div uint64) uint64 {
	return (num + div - 1) / div
}

func decompose(x uint64, y uint64) (uint64, uint64) {
	return x / y, x % y
}

func setSlice(bits []uint64, pos uint64, codeLen uint8, val uint64) {
	if codeLen == 0 {
		return
	}
	block, offset := decompose(pos, smallBlockSize)
	bits[block] |= val << offset
	if offset+uint64(codeLen) > smallBlockSize {
		bits[block+1] |= (val >> (smallBlockSize - offset))
	}
}

func getBit(x uint64, pos uint8) bool {
	return ((x >> pos) & 1) == 1
}

func getSlice(bits []uint64, pos uint64, codeLen uint8) uint64 {
	if codeLen == 0 {
		return 0
	}
	block, offset := decompose(pos, smallBlockSize)
	ret := (bits[block] >> offset)
	if offset+uint64(codeLen) > smallBlockSize {
		ret |= (bits[block+1] << (smallBlockSize - offset))
	}
	if codeLen == 64 {
		return ret
	}
	return ret & ((1 << codeLen) - 1)
}

func bitNum(x uint64, n uint64, b bool) uint64 {
	if b {
		return x
	}
	return n - x
}

func printBit(x uint64) {
	for i := 0; i < 64; i++ {
		fmt.Printf("%d", i%10)
	}
	fmt.Printf("\n")
	for i := uint8(0); i < 64; i++ {
		if getBit(x, i) {
			fmt.Printf("1")
		} else {
			fmt.Printf("0")
		}
	}
	fmt.Printf("\n")
}
