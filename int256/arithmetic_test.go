package int256

import (
	"os"
	"runtime/pprof"
	"testing"

	"github.com/gnoswap-labs/uint256"
	base "github.com/linhbkhn95/int256" // for benchmark
)

func TestAdd(t *testing.T) {
	tests := []struct {
		x, y, want string
	}{
		{"0", "1", "1"},
		{"1", "0", "1"},
		{"1", "1", "2"},
		{"1", "2", "3"},
		// NEGATIVE
		{"-1", "1", "0"},
		{"1", "-1", "0"},
		{"3", "-3", "0"},
		{"-1", "-1", "-2"},
		{"-1", "-2", "-3"},
		{"-1", "3", "2"},
		{"3", "-1", "2"},
		// OVERFLOW
		{"115792089237316195423570985008687907853269984665640564039457584007913129639935", "1", "0"},
	}

	for _, tc := range tests {
		x, err := FromDecimal(tc.x)
		if err != nil {
			t.Error(err)
			continue
		}

		y, err := FromDecimal(tc.y)
		if err != nil {
			t.Error(err)
			continue
		}

		want, err := FromDecimal(tc.want)
		if err != nil {
			t.Error(err)
			continue
		}

		got := New()
		got.Add(x, y)

		if got.Neq(want) {
			t.Errorf("Add(%s, %s) = %v, want %v", tc.x, tc.y, got.ToString(), want.ToString())
		}
	}
}

func TestAddUint256(t *testing.T) {
	tests := []struct {
		x, y, want string
	}{
		{"0", "1", "1"},
		{"1", "0", "1"},
		{"1", "1", "2"},
		{"1", "2", "3"},
		{"-1", "1", "0"},
		{"-1", "3", "2"},
		{"-115792089237316195423570985008687907853269984665640564039457584007913129639934", "115792089237316195423570985008687907853269984665640564039457584007913129639935", "1"},
		{"-115792089237316195423570985008687907853269984665640564039457584007913129639935", "115792089237316195423570985008687907853269984665640564039457584007913129639934", "-1"},
		// OVERFLOW
		{"-115792089237316195423570985008687907853269984665640564039457584007913129639935", "115792089237316195423570985008687907853269984665640564039457584007913129639935", "0"},
	}

	for _, tc := range tests {
		x, err := FromDecimal(tc.x)
		if err != nil {
			t.Error(err)
			continue
		}

		y, err := uint256.FromDecimal(tc.y)
		if err != nil {
			t.Error(err)
			continue
		}

		want, err := FromDecimal(tc.want)
		if err != nil {
			t.Error(err)
			continue
		}

		got := New()
		got.AddUint256(x, y)

		if got.Neq(want) {
			t.Errorf("AddUint256(%s, %s) = %v, want %v", tc.x, tc.y, got.ToString(), want.ToString())
		}
	}
}

func TestAddDelta(t *testing.T) {
	tests := []struct {
		z, x, y, want string
	}{
		{"0", "0", "0", "0"},
		{"0", "0", "1", "1"},
		{"0", "1", "0", "1"},
		{"0", "1", "1", "2"},
		{"1", "2", "3", "5"},
		{"5", "10", "-3", "7"},
		// underflow
		{"1", "2", "-3", "115792089237316195423570985008687907853269984665640564039457584007913129639935"},
	}

	for _, tc := range tests {
		z, err := uint256.FromDecimal(tc.z)
		if err != nil {
			t.Error(err)
			continue
		}

		x, err := uint256.FromDecimal(tc.x)
		if err != nil {
			t.Error(err)
			continue
		}

		y, err := FromDecimal(tc.y)
		if err != nil {
			t.Error(err)
			continue
		}

		want, err := uint256.FromDecimal(tc.want)
		if err != nil {
			t.Error(err)
			continue
		}

		AddDelta(z, x, y)

		if z.Neq(want) {
			t.Errorf("AddDelta(%s, %s, %s) = %v, want %v", tc.z, tc.x, tc.y, z.ToString(), want.ToString())
		}
	}
}

func TestAddDeltaOverflow(t *testing.T) {
	tests := []struct {
		z, x, y string
		want    bool
	}{
		{"0", "0", "0", false},
		// underflow
		{"1", "2", "-3", true},
	}

	for _, tc := range tests {
		z, err := uint256.FromDecimal(tc.z)
		if err != nil {
			t.Error(err)
			continue
		}

		x, err := uint256.FromDecimal(tc.x)
		if err != nil {
			t.Error(err)
			continue
		}

		y, err := FromDecimal(tc.y)
		if err != nil {
			t.Error(err)
			continue
		}

		result := AddDeltaOverflow(z, x, y)
		if result != tc.want {
			t.Errorf("AddDeltaOverflow(%s, %s, %s) = %v, want %v", tc.z, tc.x, tc.y, result, tc.want)
		}
	}
}

func TestSub(t *testing.T) {
	tests := []struct {
		x, y, want string
	}{
		{"1", "0", "1"},
		{"1", "1", "0"},
		{"-1", "1", "-2"},
		{"1", "-1", "2"},
		{"-1", "-1", "0"},
		{"-115792089237316195423570985008687907853269984665640564039457584007913129639935", "-115792089237316195423570985008687907853269984665640564039457584007913129639935", "0"},
		{"-115792089237316195423570985008687907853269984665640564039457584007913129639935", "0", "-115792089237316195423570985008687907853269984665640564039457584007913129639935"},
		{x: "-115792089237316195423570985008687907853269984665640564039457584007913129639935", y: "1", want: "0"},
	}

	for _, tc := range tests {
		x, err := FromDecimal(tc.x)
		if err != nil {
			t.Error(err)
			continue
		}

		y, err := FromDecimal(tc.y)
		if err != nil {
			t.Error(err)
			continue
		}

		want, err := FromDecimal(tc.want)
		if err != nil {
			t.Error(err)
			continue
		}

		got := New()
		got.Sub(x, y)

		if got.Neq(want) {
			t.Errorf("Sub(%s, %s) = %v, want %v", tc.x, tc.y, got.ToString(), want.ToString())
		}
	}
}

func TestSubUint256(t *testing.T) {
	tests := []struct {
		x, y, want string
	}{
		{"0", "1", "-1"},
		{"1", "0", "1"},
		{"1", "1", "0"},
		{"1", "2", "-1"},
		{"-1", "1", "-2"},
		{"-1", "3", "-4"},
		// underflow
		{"-115792089237316195423570985008687907853269984665640564039457584007913129639935", "1", "-0"},
		{"-115792089237316195423570985008687907853269984665640564039457584007913129639935", "2", "-1"},
		{"-115792089237316195423570985008687907853269984665640564039457584007913129639935", "3", "-2"},
	}

	for _, tc := range tests {
		x, err := FromDecimal(tc.x)
		if err != nil {
			t.Error(err)
			continue
		}

		y, err := uint256.FromDecimal(tc.y)
		if err != nil {
			t.Error(err)
			continue
		}

		want, err := FromDecimal(tc.want)
		if err != nil {
			t.Error(err)
			continue
		}

		got := New()
		got.SubUint256(x, y)

		if got.Neq(want) {
			t.Errorf("SubUint256(%s, %s) = %v, want %v", tc.x, tc.y, got.ToString(), want.ToString())
		}
	}
}

func TestMul(t *testing.T) {
	tests := []struct {
		x, y, want string
	}{
		{"5", "3", "15"},
		{"-5", "3", "-15"},
		{"5", "-3", "-15"},
		{"0", "3", "0"},
		{"3", "0", "0"},
	}

	for _, tc := range tests {
		x, err := FromDecimal(tc.x)
		if err != nil {
			t.Error(err)
			continue
		}

		y, err := FromDecimal(tc.y)
		if err != nil {
			t.Error(err)
			continue
		}

		want, err := FromDecimal(tc.want)
		if err != nil {
			t.Error(err)
			continue
		}

		got := New()
		got.Mul(x, y)

		if got.Neq(want) {
			t.Errorf("Mul(%s, %s) = %v, want %v", tc.x, tc.y, got.ToString(), want.ToString())
		}
	}
}

func TestDiv(t *testing.T) {
	tests := []struct {
		x, y, expected string
	}{
		{"1", "1", "1"},
		{"0", "1", "0"},
		{"-1", "1", "-1"},
		{"1", "-1", "-1"},
		{"-1", "-1", "1"},
		{"-6", "3", "-2"},
		{"10", "-2", "-5"},
		{"-10", "3", "-3"},
		{"7", "3", "2"},
		{"-7", "3", "-2"},
		// the maximum value of a positive number in int256 is less than the maximum value of a uint256
		{"57896044618658097711785492504343953926634992332820282019728792003956564819967", "2", "28948022309329048855892746252171976963317496166410141009864396001978282409983"}, // (Max int256 - 1) / 2
		{"-57896044618658097711785492504343953926634992332820282019728792003956564819967", "2", "-28948022309329048855892746252171976963317496166410141009864396001978282409983"}, // (Min int256 + 1) / 2
	}

	for _, tt := range tests {
		t.Run(tt.x+"/"+tt.y, func(t *testing.T) {
			x := MustFromDecimal(tt.x)
			y := MustFromDecimal(tt.y)
			result := Zero().Div(x, y)
			if result.ToString() != tt.expected {
				t.Errorf("Div(%s, %s) = %s, want %s", tt.x, tt.y, result.ToString(), tt.expected)
			}
		})
	}

	t.Run("Division by zero", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Div(1, 0) did not panic")
			}
		}()
		x := MustFromDecimal("1")
		y := MustFromDecimal("0")
		Zero().Div(x, y)
	})
}

func TestQuo(t *testing.T) {
	tests := []struct {
		x, y, want string
	}{
		{"0", "1", "0"},
		{"0", "-1", "0"},
		{"10", "1", "10"},
		{"10", "-1", "-10"},
		{"-10", "1", "-10"},
		{"-10", "-1", "10"},
		// {"10", "-3", "-3"},
		{"10", "3", "3"},
	}

	for _, tc := range tests {
		x, err := FromDecimal(tc.x)
		if err != nil {
			t.Error(err)
			continue
		}

		y, err := FromDecimal(tc.y)
		if err != nil {
			t.Error(err)
			continue
		}

		want, err := FromDecimal(tc.want)
		if err != nil {
			t.Error(err)
			continue
		}

		got := New()
		got.Quo(x, y)

		if got.Neq(want) {
			t.Errorf("Quo(%s, %s) = %v, want %v", tc.x, tc.y, got.ToString(), want.ToString())
		}
	}
}

func TestRem(t *testing.T) {
	tests := []struct {
		x, y, want string
	}{
		{"0", "1", "0"},
		{"0", "-1", "0"},
		{"10", "1", "0"},
		{"10", "-1", "0"},
		{"-10", "1", "0"},
		{"-10", "-1", "0"},
		{"10", "3", "1"},
		{"10", "-3", "1"},
		{"-10", "3", "-1"},
		{"-10", "-3", "-1"},
	}

	for _, tc := range tests {
		x, err := FromDecimal(tc.x)
		if err != nil {
			t.Error(err)
			continue
		}

		y, err := FromDecimal(tc.y)
		if err != nil {
			t.Error(err)
			continue
		}

		want, err := FromDecimal(tc.want)
		if err != nil {
			t.Error(err)
			continue
		}

		got := New()
		got.Rem(x, y)

		if got.Neq(want) {
			t.Errorf("Rem(%s, %s) = %v, want %v", tc.x, tc.y, got.ToString(), want.ToString())
		}
	}
}

func TestMod(t *testing.T) {
	tests := []struct {
		x, y, want string
	}{
		{"0", "1", "0"},
		{"0", "-1", "0"},
		{"10", "1", "0"},
		{"10", "-1", "0"},
		{"-10", "1", "0"},
		{"-10", "-1", "0"},
		{"10", "3", "1"},
		{"10", "-3", "1"},
		{"-10", "3", "2"},
		{"-10", "-3", "2"},
	}

	for _, tc := range tests {
		x, err := FromDecimal(tc.x)
		if err != nil {
			t.Error(err)
			continue
		}

		y, err := FromDecimal(tc.y)
		if err != nil {
			t.Error(err)
			continue
		}

		want, err := FromDecimal(tc.want)
		if err != nil {
			t.Error(err)
			continue
		}

		got := New()
		got.Mod(x, y)

		if got.Neq(want) {
			t.Errorf("Mod(%s, %s) = %v, want %v", tc.x, tc.y, got.ToString(), want.ToString())
		}
	}
}

func TestModPanic(t *testing.T) {
	tests := []struct {
		x, y string
	}{
		{"10", "0"},
		{"10", "-0"},
		{"-10", "0"},
		{"-10", "-0"},
	}

	for _, tc := range tests {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Mod(%s, %s) did not panic", tc.x, tc.y)
			}
		}()
		x, err := FromDecimal(tc.x)
		if err != nil {
			t.Error(err)
			continue
		}

		y, err := FromDecimal(tc.y)
		if err != nil {
			t.Error(err)
			continue
		}

		result := New().Mod(x, y)
		t.Errorf("Mod(%s, %s) = %v, want %v", tc.x, tc.y, result.ToString(), "0")
	}
}

func TestDivE(t *testing.T) {
	testCases := []struct {
		x, y int64
		want int64
	}{
		{8, 3, 2},
		{8, -3, -2},
		{-8, 3, -3},
		{-8, -3, 3},
		{1, 2, 0},
		{1, -2, 0},
		{-1, 2, -1},
		{-1, -2, 1},
		{0, 1, 0},
		{0, -1, 0},
	}

	for _, tc := range testCases {
		x := NewInt(tc.x)
		y := NewInt(tc.y)
		want := NewInt(tc.want)
		got := new(Int).DivE(x, y)
		if got.Cmp(want) != 0 {
			t.Errorf("DivE(%v, %v) = %v, want %v", tc.x, tc.y, got, want)
		}
	}
}

func TestDivEByZero(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("DivE did not panic on division by zero")
		}
	}()

	x := NewInt(1)
	y := NewInt(0)
	new(Int).DivE(x, y)
}

func TestModEByZero(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("ModE did not panic on division by zero")
		}
	}()

	x := NewInt(1)
	y := NewInt(0)
	new(Int).ModE(x, y)
}

func TestLargeNumbers(t *testing.T) {
	x, _ := new(Int).SetString("123456789012345678901234567890")
	y, _ := new(Int).SetString("987654321098765432109876543210")

	// Expected results (calculated separately)
	expectedQ, _ := new(Int).SetString("0")
	expectedR, _ := new(Int).SetString("123456789012345678901234567890")

	gotQ := new(Int).DivE(x, y)
	gotR := new(Int).ModE(x, y)

	if gotQ.Cmp(expectedQ) != 0 {
		t.Errorf("DivE with large numbers: got %v, want %v", gotQ, expectedQ)
	}

	if gotR.Cmp(expectedR) != 0 {
		t.Errorf("ModE with large numbers: got %v, want %v", gotR, expectedR)
	}
}

func TestAbs(t *testing.T) {
	tests := []struct {
		x, want string
	}{
		{"0", "0"},
		{"1", "1"},
		{"-1", "1"},
		{"-2", "2"},
		{"-100000000000", "100000000000"},
	}

	for _, tc := range tests {
		x, err := FromDecimal(tc.x)
		if err != nil {
			t.Error(err)
			continue
		}

		got := x.Abs()

		if got.ToString() != tc.want {
			t.Errorf("Abs(%s) = %v, want %v", tc.x, got.ToString(), tc.want)
		}
	}
}

// Benchmarks

func BenchmarkAdd(b *testing.B) {
	x := NewInt(1234567890)
	y := NewInt(9876543210)
	z := NewInt(0)

	xx := base.NewInt(1234567890)
	yy := base.NewInt(9876543210)
	zz := base.NewInt(0)

	// prevent compiler optimizations
	var r Int
	var rr base.Int

	b.Run("gno int256 Add", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			z.Add(x, y)
		}
		z = &r
	})

	b.Run("base int256 package Add", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			zz.Add(xx, yy)
		}
		zz = &rr
	})
}

func BenchmarkSub(b *testing.B) {
	x := NewInt(9876543210)
	y := NewInt(1234567890)
	z := NewInt(0)

	xx := base.NewInt(9876543210)
	yy := base.NewInt(1234567890)
	zz := base.NewInt(0)

	// prevent compiler optimizations
	var r Int
	var rr base.Int

	b.Run("gno int256 Sub", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			z.Sub(x, y)
		}
		z = &r
	})

	b.Run("base int256 package Sub", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			zz.Sub(xx, yy)
		}
		zz = &rr
	})
}

func BenchmarkMul(b *testing.B) {
	x := NewInt(12345)
	y := NewInt(67890)
	z := NewInt(0)

	xx := base.NewInt(12345)
	yy := base.NewInt(67890)
	zz := base.NewInt(0)

	// prevent compiler optimizations
	var r Int
	var rr base.Int

	b.Run("gno int256 Mul", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			z.Mul(x, y)
		}
		z = &r
	})

	b.Run("base int256 package Mul", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			zz.Mul(xx, yy)
		}
		zz = &rr
	})
}

func BenchmarkDiv(b *testing.B) {
	x := NewInt(9876543210)
	y := NewInt(1234)
	z := NewInt(0)

	xx := base.NewInt(9876543210)
	yy := base.NewInt(1234)
	zz := base.NewInt(0)

	// prevent compiler optimizations
	var r Int
	var rr base.Int

	b.Run("gno int256 Div", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			z.Div(x, y)
		}
		z = &r
	})

	b.Run("base int256 package Div", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			zz.Div(xx, yy)
		}
		zz = &rr
	})
}

func BenchmarkDiv_profile(b *testing.B) {
	x := NewInt(9876543210)
	y := NewInt(1234)
	z := NewInt(0)

    f, _ := os.Create("cpu_profile_div.prof")
    pprof.StartCPUProfile(f)
    defer pprof.StopCPUProfile()

    for i := 0; i < b.N; i++ {
        z.Div(x, y)
    }
}

func BenchmarkQuo(b *testing.B) {
	x := NewInt(9876543210)
	y := NewInt(1234)
	z := NewInt(0)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		z.Quo(x, y)
	}
}

func BenchmarkRem(b *testing.B) {
	x := NewInt(9876543210)
	y := NewInt(1234)
	z := NewInt(0)

	xx := base.NewInt(9876543210)
	yy := base.NewInt(1234)
	zz := base.NewInt(0)

	// prevent compiler optimizations
	var r Int
	var rr base.Int

	b.Run("gno int256 Rem", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			z.Rem(x, y)
		}
		z = &r
	})

	b.Run("base int256 package Rem", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			zz.Rem(xx, yy)
		}
		zz = &rr
	})
}

func BenchmarkMod(b *testing.B) {
	x := NewInt(9876543210)
	y := NewInt(1234)
	z := NewInt(0)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		z.Mod(x, y)
	}
}
