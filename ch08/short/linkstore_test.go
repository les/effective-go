package short

import (
	"context"
	"errors"
	"testing"

	"github.com/inancgumus/effective-go/ch08/bite"
	"github.com/inancgumus/effective-go/ch08/sqlx/sqlxtest"
)

func TestLinkStore(t *testing.T) {
	var (
		ctx   = context.Background()
		store = &LinkStore{
			DB: sqlxtest.Dial(t),
		}
		link = Link{
			Key: "go", URL: "https://go.dev",
		}
	)
	t.Run("create", func(t *testing.T) {
		err := store.Create(ctx, link)
		if err != nil {
			t.Errorf("Create() err = %q, want <nil>", err)
		}
	})
	t.Run("create/exists", func(t *testing.T) {
		err := store.Create(ctx, link)
		if !errors.Is(err, bite.ErrExists) {
			t.Errorf("Create(%q) err = %q, want %q", link.Key, err, bite.ErrExists)
		}
	})
	t.Run("retrieve", func(t *testing.T) {
		got, err := store.Retrieve(ctx, link.Key)
		if err != nil {
			t.Errorf("Retrieve(%q) err = %q, want <nil>", link.Key, err)
		}
		if got != link {
			t.Errorf("Retrieve(%q) = %#v, want %#v", link.Key, got, link)
		}
	})
	t.Run("retrieve/not_found", func(t *testing.T) {
		_, err := store.Retrieve(ctx, "not-found")
		if !errors.Is(err, bite.ErrNotExist) {
			t.Errorf("Retrieve(%q) err = %q, want %q", link.Key, err, bite.ErrNotExist)
		}
	})
}
