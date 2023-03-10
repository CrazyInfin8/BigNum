package bignum

func (l BigInt) Sub(r BigInt) (new BigInt){
	maxLen := l.Len()
	if len(r) > maxLen {
		maxLen = r.Len()
	}
	maxLen++
	new = make(BigInt, maxLen)
	var buf = 1
	for i := 0; i < maxLen; i++ {
		buf += int(l.DigitRaw(i)) + int(^r.DigitRaw(i))
		new[i] = byte(buf & 0xFF)
		buf >>= 8
	}
	new.Clean()
	return new
}

func (l *BigInt) SubInPlace(r BigInt) {
	maxLen := l.Len()
	if len(r) > maxLen {
		maxLen = r.Len()
	}
	maxLen++
	l.Grow(maxLen)
	var buf = 1
	for i := 0; i < maxLen; i++ {
		buf += int(l.DigitRaw(i)) + int(^r.DigitRaw(i))
		(*l)[i] = byte(buf & 0xFF)
		buf >>= 8
	}
	l.Clean()
}