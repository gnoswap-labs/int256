package word

import (
	"encoding/binary"

	"github.com/holiman/uint256"
)

type Uint32Words struct {
	Words [8]uint32
}

func ToWords(num *uint256.Int) *Uint32Words {
	// convert uint256 to bytes
	bytes := num.Bytes()

	// convert byte array to 32bit size words
	words := new(Uint32Words)
	for i := 0; i < 8; i++ {
		// Extract each 32-bit word.
		// The binary.BigEndian.Uint32 function reads 4 bytes from the byte slice and converts it to uint32.
		words.Words[i] = binary.BigEndian.Uint32(bytes[i*4 : (i+1)*4])
	}

	return words
}