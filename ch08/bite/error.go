// Package bite provides cross-cutting business rules and services.
//
// For now, it has only errors.
//
// For example, other shared types—structs, interfaces, etc., may be in different
// files/subpackages.
//
// There can be a type called User type in this package to represent a Bite company
// user. It can be in a file called user.go.
//
// Note:
// This could be in the root of the project, but I kept it here because the root
// directory's name is ch0x and it's not a good name for a package. For example,
// ch0x.ErrExists! :-]
package bite

import "errors"

var (
	ErrExists         = errors.New("already exists")
	ErrNotExist       = errors.New("does not exist")
	ErrInvalidRequest = errors.New("invalid request")
	ErrInternal       = errors.New("internal error: please try again later or contact support")
)

// other shared types—structs, interfaces, etc., may be in different files.
