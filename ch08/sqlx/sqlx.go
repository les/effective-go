// Package sqlx provides a wrapper around database/sql.DB.
//
// To take control of the database driver dependencies, sqlite
// driver import happens in a single package. Also, this is a
// good place to change the driver or support multiple ones.
//
// Another benefit of this package is that it provides an
// interception point for the database operations. For example,
// to log/trace the queries; provide a simpler API, etc.
package sqlx

import (
	_ "modernc.org/sqlite"
)
