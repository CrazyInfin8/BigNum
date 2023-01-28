package bignum

// Add adds the value of "r" into the value of "l".
func (l BigInt) Add(r BigInt) (new BigInt){
	maxLen := l.Len()
	if rLen := r.Len(); rLen > maxLen {
		maxLen = rLen
	}
	maxLen++
	new = make(BigInt, maxLen)
	var buf int = 0
	for i := 0; i < maxLen; i++ {
		buf += int(l.DigitRaw(i)) + int(r.DigitRaw(i))
		new[i] = byte(buf & 0xFF)
		buf >>= 8
	}
	new.Clean()
	return new
}

// Add adds the value of "r" into the value of "l".
func (l *BigInt) AddInPlace(r BigInt) {
	maxLen := l.Len()
	if rLen := r.Len(); rLen > maxLen {
		maxLen = rLen
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