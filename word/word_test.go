package word

import (
	"testing"

	"github.com/holiman/uint256"
)

func TestToWords(t *testing.T) {
	tests := []struct {
		name   string
		hex    string // Hexadecimal representation of the uint256 value
		expect [8]uint32
	}{
		{
			name: "HalfMaxUint256",
			hex:  "0x7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
			expect: [8]uint32{
				0x7fffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
			},
		},
		{
			name: "RandomNumber",
			hex:  "0x1a2b3c4d5e6f708192a3b4c5d6e7f8091a2b3c4d5e6f708192a3b4c5d6e7f809",
			expect: [8]uint32{
				0x1a2b3c4d, 0x5e6f7081, 0x92a3b4c5, 0xd6e7f809,
				0x1a2b3c4d, 0x5e6f7081, 0x92a3b4c5, 0xd6e7f809,
			},
		},
		{
			name: "MaxUint256",
			hex:  "0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
			expect: [8]uint32{
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
			},
		},
		{
			name: "ArbitraryNumber",
			hex:  "0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef",
			expect: [8]uint32{
				0x12345678, 0x90abcdef, 0x12345678, 0x90abcdef,
				0x12345678, 0x90abcdef, 0x12345678, 0x90abcdef,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			num, _ := uint256.FromHex(tc.hex)

			result := ToWords(num)

			for i, word := range result.Words {
				if word != tc.expect[i] {
					t.Errorf("Test %s failed: word %d is 0x%x, want 0x%x", tc.name, i, word, tc.expect[i])
				}
			}
		})
	}
}
