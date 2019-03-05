package guid

// GUID generates guid as c #
type GUID struct {
	_a int32
	_b int16
	_c int16
	_d byte
	_e byte
	_f byte
	_g byte
	_h byte
	_i byte
	_j byte
	_k byte
}

//NewGUID create a new instance of GUID
func NewGUID(b []byte) *GUID {
	return &GUID{
		_a: (int32(b[3]) << 24) | (int32(b[2]) << 16) | (int32(b[1]) << 8) | int32(b[0]),
		_b: int16(((int32(b[5]) << 8) | int32(b[4]))),
		_c: int16(((int32(b[7]) << 8) | int32(b[6]))),
		_d: b[8],
		_e: b[9],
		_f: b[10],
		_g: b[11],
		_h: b[12],
		_i: b[13],
		_j: b[14],
		_k: b[15],
	}
}

func hexToChar(a int32) rune {
	a = a & 0xf
	if a > 9 {
		return rune(a - 10 + 0x61)
	}
	return rune(a + 0x30)
}

func hexsToChars(guidChars []rune, offset, a, b int32) int32 {
	guidChars[offset] = hexToChar(a >> 4)
	offset++
	guidChars[offset] = hexToChar(a)
	offset++
	guidChars[offset] = hexToChar(b >> 4)
	offset++
	guidChars[offset] = hexToChar(b)
	offset++
	return offset
}

//ToString conver GUID to string format
func (u GUID) ToString() string {
	var offset int32
	guidChars := make([]rune, 36)

	offset = hexsToChars(guidChars, offset, u._a>>24, u._a>>16)
	offset = hexsToChars(guidChars, offset, u._a>>8, u._a)
	guidChars[offset] = '-'
	offset++
	offset = hexsToChars(guidChars, offset, int32(u._b>>8), int32(u._b))
	guidChars[offset] = '-'
	offset++
	offset = hexsToChars(guidChars, offset, int32(u._c>>8), int32(u._c))
	guidChars[offset] = '-'
	offset++
	offset = hexsToChars(guidChars, offset, int32(u._d), int32(u._e))
	guidChars[offset] = '-'
	offset++
	offset = hexsToChars(guidChars, offset, int32(u._f), int32(u._g))
	offset = hexsToChars(guidChars, offset, int32(u._h), int32(u._i))
	offset = hexsToChars(guidChars, offset, int32(u._j), int32(u._k))
	return string(guidChars)
}
