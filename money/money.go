// Package money rounds and prints amounts in cents.
package money

import (
	"fmt"
	"math"
)

func RoundHalfUp(x float64) int64 {
	return int64(math.Floor(x + 0.5))
}

func FormatCents(cents int64) string {
	sign := ""
	if cents < 0 {
		sign = "-"
		cents = -cents
	}
	return fmt.Sprintf("%s$%d.%02d", sign, cents/100, cents%100)
}
