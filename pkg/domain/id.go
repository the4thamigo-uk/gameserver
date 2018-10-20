package domain

import (
	"github.com/rs/xid"
)

// NewID generates a new unique id for an object in the domain
var NewID = DefaultNewID

// DefaultNewID is the default implementation of NewID, and creates an id string similar to MongoDB.
func DefaultNewID() string {
	id := xid.New()
	return id.String()
}
