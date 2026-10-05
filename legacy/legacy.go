// Package legacy keeps the pre-catalog total for old callers.
package legacy

import "github.com/evanmschultz/hylla-fixture-go-v/pricing"

// LegacyTotal sums prices with no tax.
func LegacyTotal(items []pricing.Priced) int64 {
	var sum int64
	for _, item := range items {
		sum += item.PriceCents()
	}
	return sum
}
