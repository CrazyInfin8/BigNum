package bignum

import (
)




func (l *BigInt)DivMod(r *BigInt) (quot BigInt, rem BigInt) {
	if r.IsZero() {
		panic("divide by zero")
	}
	return
}

