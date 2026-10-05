// Package pricing holds things that carry a price.
package pricing

// Priced is anything with a price in whole cents.
type Priced interface {
	PriceCents() int64
}

type Book struct {
	Title     string
	UnitCents int64
	Quantity  int64
}

var _ Priced = Book{}

func (b Book) PriceCents() int64 {
	return b.UnitCents * b.Quantity
}
