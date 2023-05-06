package shortener

import "github.com/inancgumus/effective-go/ch08/short"

// Service represents the services provided to the shortener.
//
// At the moment, the only service is the short link service
// (short.LinkStore) that provides the ability to create and
// retrieve short links from the data store.
//
// But, in the future, there could be other services such as
// a stats service that provides the ability to save and
// retrieve stats. Or a logger service that provides the
// ability to log messages.
type Service struct {
	LinkStore *short.LinkStore
}
