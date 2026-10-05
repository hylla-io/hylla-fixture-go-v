package parity

func IsOdd(n int) bool {
	if n == 0 {
		return false
	}
	return IsEven(n - 1)
}
