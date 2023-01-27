package bignum

import "unsafe"

type BigInt []byte

func FromInt(i int) (big BigInt) {
	big = make(BigInt, 0, unsafe.Sizeof(i))
	if i < 0 {
		big = append(big, byte(i))
		i >>= 8
		for i < -1 {
			big = append(big, byte(i))
			i >>= 8
		}
		if !big.IsNeg() { // ensure BigNum is negative
			big = append(big, 0xFF)
		}
	} else {
		for i > 0 {
			u := uint(i)
			big = append(big, byte(u))
			i >>= 8
		}
		if big.IsNeg() { // ensure BigNum is positive
			big = append(big, 0)
		}
	}
	return big
}

func (big BigInt) IsZero() bool {
	if len(big) == 0 {
		return true
	}
	i := len(big) - 1
	for i > -1 {
		if big[i] != 0 {
			return false
		}
		i--
	}
	return true
}

func (big BigInt) IsNeg() bool {
	if big.IsZero() {
		return false
	}
	return big[len(big)-1] > 0x7F
}

func (big BigInt) Len() int {
	l := len(big)
	if big.IsNeg() {
		for l > 1 && big[l-1] == 0xFF {
			if big[l-2] < 0x80 {
				break
			}
			l--
		}
	} else {
		for l > 0 && big[l-1] == 0 {
			if l > 1 && big[l-2] > 0x7F {
				break
			}
			l--
		}
	}
	return l
}

func (big BigInt) LastDigit() (d byte) {
	if big.IsZero() {
		return 0
	}
	d = big.Digit(big.Len())
	return d
}

// Digit returns the natural Base256 digit at index "i".
func (big BigInt) Digit(i int) (d byte) {
	if i >= len(big) {
		return 0
	} else if i < 0 {
		panic("Index cannot be lower than 0")
	}
	if big.IsNeg() {
		if i == 0 {
			return ^big[i] + 1
		}
		if len(big) > 1 && big[0] == 0 {
			var j int
			for j = 1; j < i; j++ {
				if big[j] != 0 {
					break
				}
			}
			if j == i {
				return ^big[i] + 1
			}
		}
		return ^big[i]
	}
	return big[i]
}

// DigitRaw returns the actual Base256 digit as it is stored in index "i".
func (big BigInt) DigitRaw(i int) (d byte) {
	if i >= len(big) {
		if big.IsNeg() {
			return 0xFF
		} else {
			return 0x00
		}
	} else if i < 0 {
		panic("Index cannot be lower than 0")
	}
	return big[i]

}

// Grow extends the size of the underlying array without changing it's value.
func (l *BigInt) Grow(capacity int) {
	if cap(*l) < capacity {
		temp := make(BigInt, len(*l), capacity)
		copy(temp, *l)
		*l = temp
	}
	digit := l.DigitRaw(len(*l))
	for len(*l) < capacity {
		*l = append(*l, digit)
	}
}

// Clean shortens the length of the underlying array as small as it can be without changing it's value.
func (l *BigInt) Clean() {
	*l = (*l)[:l.Len()]
}

func (big *BigInt) ZeroOut() {
	for i := range *big {
		(*big)[i] = 0
	}
}