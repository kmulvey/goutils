package goutils

const separator int64 = 0xFFFFFFFF

// Pack stores two ints in one. Go stores ints as 8 bytes so we store
// the first int in the bottom four and the second in the top four.
// This has a limitation of only being able to store a max value of 4294967295.
func Pack(a, b int64) int64 {
	return a | (b << 32)
}

// Unpack extracts two ints stored in one. 0xFFFFFFFF represents 32 bits of
// 1s so we AND them to x to get the first number, then we shift and AND to get
// the second number.
func Unpack(x int64) (int64, int64) {
	return x & separator, (x >> 32) & separator
}
