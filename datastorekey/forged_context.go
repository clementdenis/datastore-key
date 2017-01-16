package datastorekey

import (
	"context"
)

// ForgedContext embeds an appengine.Context, and has a custom appId.
// It overrides method FullyQualifiedAppID()
type ForgedContext struct {
	context.Context
	appId string
}

func (c *ForgedContext) FullyQualifiedAppID() string {
	return c.appId
}
