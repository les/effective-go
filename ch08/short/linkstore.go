package short

import (
	"context"
	"errors"
	"fmt"

	"github.com/inancgumus/effective-go/ch08/bite"
)

// Create persists the given link. It returns bite.ErrInvalidRequest
// if the link is invalid. Or it returns an error if the link cannot
// be created.
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
