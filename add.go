package bignum

// Add adds the value of "r" into the value of "l".
func (l *BigInt) Add(r BigInt) {
	maxLen := l.Len()
	if len(r) > maxLen {
		maxLen = r.Len()
	}
	maxLen++
	l.Grow(maxLen)
	var buf int = 0
	for i := 0; i < maxLen; i++ {
		buf += int(l.DigitRaw(i)) + int(r.DigitRaw(i))
		(*l)[i] = byte(buf & 0xFF)
		buf >>= 8
	}
	l.Clean()
}