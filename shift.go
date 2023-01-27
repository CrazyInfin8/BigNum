package bignum

func (big BigInt) Shift(digits int) (new BigInt) {
	big.Clean()
	if big.IsZero() {
		return nil
	}
	newLen := big.Len() + digits
	if newLen < 0 {
		goto ensureNeg
	}
	new = make(BigInt, newLen)
	if digits > 0 {
		copy(new[digits:], big)
		goto ensureNeg
	}
	if digits < 0 {
		copy(new, (big)[digits*-1:])
		goto ensureNeg
	}
	copy(new, big)
ensureNeg:
	if big.IsNeg() != new.IsNeg() {
		if len(new) == 0 {
			return BigInt{0xFF}
		}
		new[len(new)-1] |= 0x80
	}
	return new
}

func (big BigInt) BitShift(bits int) BigInt {
	big.Clean()
	if big.IsZero() {
		return nil // (zero)
	}
	digitsShift := bits / 8
	bitsShift := int8(bits % 8)
	var new BigInt
	if bitsShift == 0 {
		return big.Shift(digitsShift)
	}
	if bits > 0 {
		newLen := big.Len()
		newLen += (bits + 7) / 8
		new = make(BigInt, newLen)
		for i := 0; i < len(big)+1; i++ {
			if i < len(big) {
				new[i+digitsShift+1] |= big.DigitRaw(i) >> (8 - bitsShift)
			}
			new[i+digitsShift] |= big.DigitRaw(i) << (bitsShift)
		}
	}
	if bits < 0 {
		newLen := big.Len() + digitsShift
		if newLen <= 0 {
			goto ensureNeg
		}
		new = make(BigInt, newLen+1)

		for i := 0; i < len(new); i++ {
			new[i] |= big.DigitRaw(i-digitsShift) >> (bitsShift * -1)
			new[i] |= big.DigitRaw(i-digitsShift+1) << (8 + bitsShift)
		}

		goto ensureNeg
	}

ensureNeg:
	if big.IsNeg() != new.IsNeg() {
		if len(new) == 0 {
			return BigInt{0xFF}
		}
		new[len(new)-1] |= 0x80
	}
	return new
}
