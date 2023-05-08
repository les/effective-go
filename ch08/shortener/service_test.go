package shortener

import (
	"context"

	"github.com/inancgumus/effective-go/ch08/short"
)

// fakeLinkStore is a fake implementation of the short.LinkStore
// interface. It provides programmable behavior for the Create
// and Retrieve methods.
type fakeLinkStore struct {
	create   func(context.Context, short.Link) error
	retrieve func(context.Context, string) (short.Link, error)
}

// Create calls the fake create method if it is not nil. Otherwise,
// it returns nil for convenience.
func (f *fakeLinkStore) Create(ctx context.Context, ln short.Link) error {
	if f.create == nil {
		return nil
	}
	return f.create(ctx, ln)
}

// Retrieve calls the fake retrieve method if it is not nil. Otherwise,
// it returns an empty link for convenience.
func (f *fakeLinkStore) Retrieve(ctx context.Context, key string) (short.Link, error) {
	if f.retrieve == nil {
		return short.Link{}, nil
	}
	return f.retrieve(ctx, key)
}
