package okapi

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Client provides the Okapi API client which should be embedded in your client
// objects.
type Client struct {
	httpClient *http.Client
	config     Config
}

// NewClient returns a new Okapi client with a default httpClient.
func NewClient(config *Config) (*Client, error) {
	httpClient := &http.Client{}

	client := &Client{
		config:     *config,
		httpClient: httpClient,
	}

	return client, nil
}

// newRequest creates a new request object.
func (c *Client) newRequest(method, path string) (*request, error) {
	base, err := url.Parse(c.config.Address)
	if err != nil {
		return nil, err
	}

	u, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	r := &request{
		config: &c.config,
		method: method,
		url: &url.URL{
			Scheme:  base.Scheme,
			User:    base.User,
			Host:    base.Host,
			Path:    u.Path,
			RawPath: u.RawPath,
		},
		params: make(map[string][]string),
	}

	if c.config.WaitTime != 0 {
		r.params.Set("wait", r.config.WaitTime.String())
	}

	for key, values := range u.Query() {
		for _, value := range values {
			r.params.Add(key, value)
		}
	}
	return r, nil
}

// do builds an http.Request, converts the request object to a standard
// http.Request, and performs the get query itself. The response body is decoded
// into the QueryOptions Out interface. It raises an error if the reponse status
// code is anything but 200.
func (c *Client) do(method, endpoint string, q *QueryOptions) error {
	req, err := c.newRequest("GET", endpoint)
	if err != nil {
		return err
	}

	http, err := req.toHTTP()
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Do(http)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var buf bytes.Buffer
		io.Copy(&buf, resp.Body)
		return fmt.Errorf("Unexpected response code: %d (%s)", resp.StatusCode, buf.String())
	}

	if q.Out != nil {
		if err := decodeBody(resp, &q.Out); err != nil {
			return err
		}
	}
	return nil
}
