// Package catalog totals a list of priced items.
package catalog

import (
	"github.com/evanmschultz/hylla-fixture-go-v/discount"
	"github.com/evanmschultz/hylla-fixture-go-v/pricing"
	"github.com/evanmschultz/hylla-fixture-go-v/tax"
)

type Catalog struct {
	items []pricing.Priced
}

func New() *Catalog {
	return &Catalog{}
}

func (c *Catalog) Add(item pricing.Priced) {
	c.items = append(c.items, item)
}

func (c *Catalog) Count() int {
	return len(c.items)
}

func (c *Catalog) Subtotal() int64 {
	var sum int64
	for _, item := range c.items {
		sum += item.PriceCents()
	}
	return sum
}

func (c *Catalog) Total(ratePercent int64) int64 {
	var bulk int64
	if c.Count() >= 3 {
		bulk = 10
	}
	return tax.ApplyTax(discount.ApplyDiscount(c.Subtotal(), bulk), ratePercent)
}
