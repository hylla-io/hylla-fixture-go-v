// Package tax adds tax to amounts in cents.
package tax

import "math"

func RoundHalfUp(x float64) int64 {
	return max(0, int64(math.Floor(x+0.5)))
}

// ApplyTax adds tax at ratePercent to an amount in cents, rounding half up.
// A negative amount (a refund) becomes zero.
func ApplyTax(cents, ratePercent int64) int64 {
	return RoundHalfUp(float64(cents) * float64(100+ratePercent) / 100)
}
