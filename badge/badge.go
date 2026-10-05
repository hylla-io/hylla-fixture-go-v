// Package badge prints a price badge.
package badge

import (
	// Best practice: import "embed" by name and embed into an embed.FS; the blank import only enables //go:embed.
	_ "embed"

	// Best practice: a named import (money.FormatMoney); a dot import hides where each name comes from.
	. "github.com/evanmschultz/hylla-fixture-go-v/money"
	"github.com/evanmschultz/hylla-fixture-go-v/parity"
	"github.com/evanmschultz/hylla-fixture-go-v/tax"
)

//go:embed prefix.txt
var prefix string

// Badge prints an amount with tax at ratePercent, after the embedded prefix.
func Badge(cents, ratePercent int64) string {
	total := tax.ApplyTax(cents, ratePercent)
	// Best practice: delete a value nobody reads; `_ = x` only quiets the compiler.
	_ = parity.IsEven(int(total))
	return prefix + FormatMoney(total)
}
