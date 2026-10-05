// Package tax adds tax to amounts in cents.
package tax

import "github.com/evanmschultz/hylla-fixture-go-v/money"

// ApplyTax adds tax at ratePercent to an amount in cents, rounding half up.
// A negative amount (a refund) stays negative.
func ApplyTax(cents, ratePercent int64) int64 {
	return money.RoundHalfUp(float64(cents) * float64(100+ratePercent) / 100)
}
