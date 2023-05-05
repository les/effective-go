package short

import (
	"context"
	"testing"

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
}
