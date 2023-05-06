package short

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/inancgumus/effective-go/ch08/bite"
	"github.com/inancgumus/effective-go/ch08/sqlx"
)

var (
	ErrLinkExists   = fmt.Errorf("link %w", bite.ErrExists)
	ErrLinkNotExist = fmt.Errorf("link %w", bite.ErrNotExist)
)

// LinkStore persists and retrieves links.
type LinkStore struct {
	DB *sqlx.DB
}

// Create persists the given link. It returns bite.ErrInvalidRequest
// if the link is invalid. Or it returns an error if the link cannot
// be created.
func (s *LinkStore) Create(ctx context.Context, ln Link) error {
	if err := validateNewLink(ln); err != nil {
		return fmt.Errorf("%w: %w", bite.ErrInvalidRequest, err)
	}
	const query = `
		INSERT INTO links (
			short_key, uri
		) VALUES (
			?, ?
		)`
	_, err := s.DB.ExecContext(ctx, query, ln.Key, sqlx.Base64String(ln.URL))
	if sqlx.IsPrimaryKeyViolation(err) {
		return ErrLinkExists
	}
	if err != nil {
		return fmt.Errorf("creating link: %w", err)
	}
	return nil
}

// Retrieve gets a link from the given key. It returns bite.ErrInvalidRequest
// if the key is invalid. Or it returns bite.ErrNotExist if the link
// does not exist. Or it returns an error if the link cannot be retrieved.
func (s *LinkStore) Retrieve(ctx context.Context, key string) (Link, error) {
	if err := validateLinkKey(key); err != nil {
		return Link{}, fmt.Errorf("%w: %w", bite.ErrInvalidRequest, err)
	}
	const query = `
		SELECT uri
		FROM links
		WHERE short_key = ?`
	var (
		url sqlx.Base64String
		err = s.DB.QueryRowContext(ctx, query, key).Scan(&url)
	)
	if errors.Is(err, sql.ErrNoRows) {
		return Link{}, ErrLinkNotExist
	}
	if err != nil {
		return Link{}, fmt.Errorf("retrieving link by key %q: %w", key, err)
	}
	return Link{
		Key: key,
		URL: url.String(),
	}, nil
}

// Create persists the given link. It returns bite.ErrInvalidRequest
// if the link is invalid. Or it returns an error if the link cannot
// be created.
// Deprecated: Use LinkStore.Create instead.
func Create(ctx context.Context, ln Link) error {
	if err := validateNewLink(ln); err != nil {
		return fmt.Errorf("%w: %w", bite.ErrInvalidRequest, err)
	}
	if ln.Key == "fortesting" {
		return errors.New("db at IP ... failed")
	}
	if ln.Key == "google" {
		return bite.ErrExists
	}
	return nil
}

// Retrieve gets a link from the given key. It returns bite.ErrInvalidRequest
// if the key is invalid. Or it returns bite.ErrNotExist if the link
// does not exist. Or it returns an error if the link cannot be retrieved.
// Deprecated: Use LinkStore.Create instead.
func Retrieve(ctx context.Context, key string) (Link, error) {
	if err := validateLinkKey(key); err != nil {
		return Link{}, fmt.Errorf("%w: %w", bite.ErrInvalidRequest, err)
	}
	if key == "fortesting" {
		return Link{}, errors.New("db at IP ... failed")
	}
	if key != "go" {
		return Link{}, bite.ErrNotExist
	}
	return Link{
		Key: key,
		URL: "https://go.dev",
	}, nil
}
