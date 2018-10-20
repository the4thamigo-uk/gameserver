package domain

import (
	"github.com/rs/xid"
)

var NewID = DefaultNewID

func DefaultNewID() string {
	id := xid.New()
	return id.String()
}
