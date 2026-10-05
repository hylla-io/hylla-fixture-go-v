// Package discount takes a percentage off amounts in cents.
package discount

import "github.com/evanmschultz/hylla-fixture-go-v/tax"

// ApplyDiscount takes percent off an amount in cents, rounding the discount half up.
func ApplyDiscount(cents, percent int64) int64 {
	return cents - tax.RoundHalfUp(float64(cents)*float64(percent)/100)
}
