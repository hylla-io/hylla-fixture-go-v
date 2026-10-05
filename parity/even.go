// Package parity decides whether a count is even or odd.
package parity

func IsEven(n int) bool {
	if n == 0 {
		return true
	}
	return IsOdd(n - 1)
}
