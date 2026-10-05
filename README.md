# hylla-fixture-go-v

A tiny catalog used as a Hylla ingest fixture, released with version tags. Every file is small and
its contents are known exactly.

## Money

`money.RoundHalfUp` rounds half up. `money.FormatCents` prints cents as dollars. Both live in
[money/money.go](money/money.go).

## Tax

`tax.ApplyTax` in [tax/tax.go](tax/tax.go) adds tax and rounds half up. Refunds pass through:
`ApplyTax(-1000, 10)` is `-1100`.

## Totals

`catalog.Catalog.Total` taxes the subtotal. `report.RenderReport` prints a report; rounding is
described under [Tax](#tax).
