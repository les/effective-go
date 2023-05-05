// Package short provides business rules and services for URL shortening.
package short

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

// MaxKeyLen is the maximum length of a key.
const MaxKeyLen = 16

// Link represents a link.
type Link struct {
	Key string
	URL string
}

// validateNewLink checks a new link's validity.
func validateNewLink(ln Link) error {
	if err := validateLinkKey(ln.Key); err != nil {
		return err
	}
	u, err := url.ParseRequestURI(ln.URL)
	if err != nil {
		return err
	}
	if u.Host == "" {
		return errors.New("empty host")
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return errors.New("scheme must be http or https")
	}
	return nil
}

// validateLinkKey checks the key's validity.
func validateLinkKey(k string) error {
	if strings.TrimSpace(k) == "" {
		return errors.New("empty key")
	}
	if len(k) > MaxKeyLen {
		return fmt.Errorf("key too long (max %d)", MaxKeyLen)
	}
	return nil
}
