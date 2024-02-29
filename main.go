package main

import (
	"fmt"

	"github.com/gnoswap-labs/int256/word"
	"github.com/holiman/uint256"
)

func main() {
	num, _ := uint256.FromHex("0x1234567890abcdef1234567890abcdef1234567890abcdef1234544440abcdef")

	w := word.ToWords(num)

	for i := 0; i < 8; i++ {
		fmt.Printf("%d: %x\n", i, w.Words[i])
	}
}
