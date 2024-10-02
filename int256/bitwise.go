package int256

// Not returns the bitwise NOT of x, setting z to the result and returning z.
func (z *Int) Not(x *Int) *Int {
	z.value.Not(&x.value)
	return z
}

// And returns the bitwise AND of x and y, setting z to the result and returning z.
func (z *Int) And(x, y *Int) *Int {
	z.value.And(&x.value, &y.value)
	return z
}

// Or returns the bitwise OR of x and y, setting z to the result and returning z.
func (z *Int) Or(x, y *Int) *Int {
	z.value.Or(&x.value, &y.value)
	return z
}

// Xor returns the bitwise XOR of x and y, setting z to the result and returning z.
func (z *Int) Xor(x, y *Int) *Int {
	z.value.Xor(&x.value, &y.value)
	return z
}

// Rsh returns the result of shifting x right by n bits, setting z to the result and returning z.
func (z *Int) Rsh(x *Int, n uint) *Int {
	z.value.Rsh(&x.value, n)
	return z
}

// Lsh returns the result of shifting x left by n bits, setting z to the result and returning z.
func (z *Int) Lsh(x *Int, n uint) *Int {
	z.value.Lsh(&x.value, n)
	return z
}
