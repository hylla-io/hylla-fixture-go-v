# hylla-fixture-go-v

A tiny catalog used as a Hylla ingest fixture, released with version tags. Every file is small and
its contents are known exactly.

## Money

`money.FormatMoney` prints cents as dollars. It lives in [money/money.go](money/money.go).

## Tax

`tax.ApplyTax` in [tax/tax.go](tax/tax.go) adds tax and rounds with `tax.RoundHalfUp`, which lives
beside it. Refunds clamp to zero: `ApplyTax(-1000, 10)` is `0`.

## Totals

`catalog.Catalog.Total` takes a bulk discount with `discount.ApplyDiscount` at three or more
items, then taxes. `report.RenderReport` prints a report; rounding is described under [Tax](#tax).
