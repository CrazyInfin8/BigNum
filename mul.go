package bignum

func (l BigInt) Mul(r BigInt) (n BigInt){
	if l.IsZero() || r.IsZero() {
		return nil
	}
	ll, lr := l.Len(), r.Len()
	if !(l[ll-1] == 0xFF || l[ll-1] == 0x00) {
		ll += 3
	} else {
		ll += 2
	}
	if !(r[lr-1] == 0xFF || r[lr-1] == 0x00) {
		lr += 3
	} else {
		lr += 2
	}
	maxLen := ll + lr
	m := make(BigInt, maxLen)
	n = make(BigInt, 0, maxLen)
	buf, i, j := 0, 0, 0
	for ; i < ll; i++ {
		buf = 0
		j = 0
		ld := int(l.Digit(i))
		for ; j < lr; j++ {
			b := ld * int(r.Digit(j))
			buf += b
			m[i+j] = byte(buf & 0xFF)
			buf >>= 8
		}
		m[i+j] = byte(buf & 0xFF)
		n.Add(m[:j])
		m.ZeroOut()
	}
	// Negate number if value should be negative (could probably move this to another function later)
	if l.IsNeg() != r.IsNeg() {
		add1 := true
		for i := range n {
			if add1 {
				if n[i] != 0x00 {
					n[i] = ^(n[i] - 1)
					add1 = false
				}
			} else {
				n[i] = ^n[i]
			}
		}
		n = append(n, 0xFF)
	} else {
		n = append(n, 0x00)
	}
	n.Clean()
	return n
}