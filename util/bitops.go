package util

// SetBit sets the bit at pos in the integer n
func SetBit(n int, pos uint) int {
	n |= (1 << pos)
	return n
}

// ClearBit clears the bit at pos in n.
func ClearBit(n int, pos uint) int {
	return n &^ (1 << pos)
}

// HasBit checks whether bit pos in int n is set.
func HasBit(n int, pos uint) bool {
	val := n & (1 << pos)
	return (val > 0)
}
