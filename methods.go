package okapi

// Get performs a GET request to the given endpoint with additional query
// options and decodes the response body into the optional, provided interface.
// From Mozilla: the Get method requests a representation of the specified
//               resource. Requests using Get should only retrieve data.
func (c *Client) Get(endpoint string, q *QueryOptions) error {
	return c.do("GET", endpoint, q)
}

// Head performs a HEAD request to the given endpoint with additional query
// options and decodes the response body into the optional, provided interface.
// From Mozilla: The Head method asks for a response identical to a Get request,
//               but without the response body.
func (c *Client) Head(endpoint string, q *QueryOptions) error {
	return c.do("HEAD", endpoint, q)
}

// Post performs a POST request to the given endpoint with additional query
// options and decodes the response body into the optional, provided interface.
// From Mozilla: The POST method submits an entity to the specified resource,
//               often causing a change in state or side effects on the server.
func (c *Client) Post(endpoint string, q *QueryOptions) error {
	return c.do("POST", endpoint, q)
}

// Put performs a PUT request to the given endpoint with additional query
// options and decodes the response body into the optional, provided interface.
// From Mozilla: The PUT method replaces all current representations of the
//               target resource with the request payload.
func (c *Client) Put(endpoint string, q *QueryOptions) error {
	return c.do("PUT", endpoint, q)
}

// Delete performs a DELETE request to the given endpoint with additional query
// options and decodes the response body into the optional, provided interface.
// From Mozilla: The DELETE method deletes the specified resource.
func (c *Client) Delete(endpoint string, q *QueryOptions) error {
	return c.do("DELETE", endpoint, q)
}

// Connect performs a CONNECT request to the given endpoint with additional
// query options and decodes the response body into the optional, provided
// interface.
// From Mozilla: The CONNECT method establishes a tunnel to the server
//               identified by the target resource.
func (c *Client) Connect(endpoint string, q *QueryOptions) error {
	return c.do("CONNECT", endpoint, q)
}

// Options performs a OPTIONS request to the given endpoint with additional query
// options and decodes the response body into the optional, provided interface.
// From Mozilla: The OPTIONS method describes the communication options for the
//               target resource.
func (c *Client) Options(endpoint string, q *QueryOptions) error {
	return c.do("OPTIONS", endpoint, q)
}

// Trace performs a TRACE request to the given endpoint with additional query
// options and decodes the response body into the optional, provided interface.
// From Mozilla: The TRACE method performs a message loop-back test along the
//               path to the target resource.
func (c *Client) Trace(endpoint string, q *QueryOptions) error {
	return c.do("TRACE", endpoint, q)
}

// Patch performs a PATCH request to the given endpoint with additional query
// options and decodes the response body into the optional, provided interface.
// From Mozilla: The PATCH method applies partial modifications to a resource.
func (c *Client) Patch(endpoint string, q *QueryOptions) error {
	return c.do("PATCH", endpoint, q)
}
