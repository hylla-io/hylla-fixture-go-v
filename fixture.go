// Package fixture re-exports the catalog's public surface.
package fixture

import (
	"github.com/evanmschultz/hylla-fixture-go-v/badge"
	"github.com/evanmschultz/hylla-fixture-go-v/catalog"
	"github.com/evanmschultz/hylla-fixture-go-v/discount"
	"github.com/evanmschultz/hylla-fixture-go-v/money"
	"github.com/evanmschultz/hylla-fixture-go-v/pricing"
	"github.com/evanmschultz/hylla-fixture-go-v/report"
	"github.com/evanmschultz/hylla-fixture-go-v/tax"
)

type (
	Book    = pricing.Book
	Catalog = catalog.Catalog
	Priced  = pricing.Priced
)

var (
	ApplyDiscount = discount.ApplyDiscount
	ApplyTax      = tax.ApplyTax
	Badge         = badge.Badge
	FormatMoney   = money.FormatMoney
	NewCatalog    = catalog.New
	RenderReport  = report.RenderReport
)
