package okapi

import (
	"context"
	"time"
)

// QueryOptions is used to pass additional parameters to a request object.
type QueryOptions struct {
	// WaitTime overrides the global WaitTime set in the client Config on a
	// per-request basis.
	WaitTime time.Duration

	// Params are additional key value pairs that will be included in the request
	// values.
	Params map[string]string

	// obj is a body alternative. If provided and body is left empty, this object
	// will be json encoded and used as the request.Body
	obj interface{}

	// ctx is passed through to http.NewRequestWithContext. Defaults to the
	// background context when building the http.Request
	ctx context.Context
}

// Context memoizes the QueryOption's ctx field.
func (o *QueryOptions) Context() context.Context {
	if o != nil && o.ctx != nil {
		return o.ctx
	}

	return context.Background()
}
