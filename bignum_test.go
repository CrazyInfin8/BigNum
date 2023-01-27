package bignum

import (
	"fmt"
	"math"
	"unsafe"

	// "math/rand"
	"strings"
	"testing"
)

func TestBitShift(t *testing.T) {
	println("Hello world")
	num := -12
	shift := -48
	big := FromInt(num)
	new := big.BitShift(shift)

	t.Logf("\nOld: %s\nNew: %s (%v)", big.toHex(), new.toHex(), new)
	if shift < 0 {
		shift *= -1
		t.Logf("\nOld: %s (%d)\nNew: %s (%d)", intToHex(num), num, intToHex(num>>shift), num>>shift)
	} else {
		t.Logf("\nOld: %s (%d)\nNew: %s (%d)", intToHex(num), num, intToHex(num<<shift), num<<shift)
	}
}

func addFuzzs(f *testing.F) {
	values := [][2]int{{0, 0}, {1, 0}, {0, 1}, {1, 1}, {0, -0}, {1, -0}, {0, -1}, {1, -1}, {-0, 0}, {-1, 0}, {-0, 1}, {-1, 1}, {-0, -0}, {-1, -0}, {-0, -1}, {-1, -1}, {0xFF, 1}, {1, 0xFF}, {-0xFF, -1}, {-1, -0xFF}, {0xFF, -1}, {1, -0xFF}, {-0xFF, 1}, {-1, 0xFF}, {0xABCDEF, 0xFEDCBA}, {0x1ABCDEF, 0x1FEDCBA}, {-0x1ABCDEF, 0x1FEDCBA}, {0x1ABCDEF, -0x1FEDCBA}, {-0x1ABCDEF, -0x1FEDCBA}, {math.MaxInt, math.MinInt}, {math.MinInt, math.MaxInt}}
	for _, v := range values {
		f.Add(v[0], v[1])
	}
}

func FuzzBitShift(f *testing.F) {
	f.Add(0x123456789ABCDEF0, int8(-127))
	f.Add(0x123456789ABCDEF0, int8(-5))
	f.Add(0x123456789ABCDEF0, int8(-16))
	f.Add(0x123456789ABCDEF0, int8(-9))
	f.Add(0x123456789ABCDEF0, int8(0))
	f.Add(0x123456789ABCDEF0, int8(5))
	f.Add(0x123456789ABCDEF0, int8(9))
	f.Add(0x123456789ABCDEF0, int8(16))
	f.Add(0x123456789ABCDEF0, int8(127))
	f.Add(0, int8(8))
	f.Fuzz(func(t *testing.T, a int, b int8) {
		var iStr string
		if b < 0 {
			iStr = intToHex(a >> int(b*-1))
		} else {
			iStr = intToHex(a << int(b))
		}
		aBig := FromInt(a)
		nBig := aBig.BitShift(int(b))
		bStr := nBig.toHex()
		if iStr != bStr {
			t.Errorf("Values don't match: ( %s | %s )\nL: %s (%d)\nR: %s (%d)", iStr, bStr, intToHex(a), a, intToHex(int(b)), b)
		}
	})
}

func FuzzAdd(f *testing.F) {
	addFuzzs(f)
	f.Fuzz(func(t *testing.T, a int, b int) {
		aBig := FromInt(int(a))
		bBig := FromInt(int(b))
		aBig.Add(bBig)
		iStr := intToHex(a + b)
		bStr := aBig.toHex()
		if iStr != bStr {
			t.Errorf("Values don't match: ( %s | %s )\nL: %s (%d)\nR: %s (%d)", iStr, bStr, intToHex(a), a, intToHex(b), b)
		}
	})
}

func FuzzSub(f *testing.F) {
	addFuzzs(f)
	f.Fuzz(func(t *testing.T, a int, b int) {
		aBig := FromInt(int(a))
		bBig := FromInt(int(b))
		aBig.Sub(bBig)
		iStr := intToHex(a - b)
		bStr := aBig.toHex()
		if iStr != bStr {
			t.Errorf("Values don't match: ( %s | %s )\nL: %s (%d)\nR: %s (%d)", iStr, bStr, intToHex(a), a, intToHex(b), b)
		}
	})
}

func FuzzMul(f *testing.F) {
	addFuzzs(f)
	f.Fuzz(func(t *testing.T, a int, b int) {
		aBig := FromInt(int(a))
		bBig := FromInt(int(b))
		nBig := aBig.Mul(bBig)
		iStr := intToHex(a * b)
		bStr := nBig.toHex()
		if iStr != bStr {
			t.Errorf("Values don't match: ( %s | %s )\nL: %s (%d)\nR: %s (%d)", iStr, bStr, intToHex(a), a, intToHex(b), b)

		}
	})
}

const intSize = int(unsafe.Sizeof(int(0)))

func (big BigInt) toHex() string {
	buf := strings.Builder{}
	// We can't really compare more bytes than allowed in an int, so just compare the lower 32/64-bits
	for i := intSize - 1; i > -1; i-- {
		fmt.Fprintf(&buf, "%02X", big.DigitRaw(i))
	}
	return buf.String()
}

func intToHex(num int) string {
	switch intSize {
	case 4:
		return fmt.Sprintf("%08X", uint(num))
	case 8:
		return fmt.Sprintf("%016X", uint(num))
	default:
		panic("Unsupported byte-length of int")
	}
}
