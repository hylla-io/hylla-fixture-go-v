// Package report prints a catalog.
package report

import (
	"fmt"
	"strings"

	"github.com/evanmschultz/hylla-fixture-go-v/catalog"
	"github.com/evanmschultz/hylla-fixture-go-v/money"
	"github.com/evanmschultz/hylla-fixture-go-v/parity"
)

func RenderReport(c *catalog.Catalog, ratePercent int64) string {
	kind := "odd"
	if parity.IsEven(c.Count()) {
		kind = "even"
	}
	return strings.Join([]string{
		fmt.Sprintf("items: %d (%s)", c.Count(), kind),
		"subtotal: " + money.FormatMoney(c.Subtotal()),
		"total: " + money.FormatMoney(c.Total(ratePercent)),
	}, "\n")
}
