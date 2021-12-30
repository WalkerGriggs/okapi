package okapi

import (
	"time"
)

type Config struct {
	// Address is the advertised HTTP address of the target endpoints. The address
	// must include the scheme, host (or IP), and port.
	Address string

	// WaitTime is the length of time which the client will block for requests.
	WaitTime time.Duration
}

// DefaultConfig returns a default configuration for the client
func DefaultConfig() *Config {
	return &Config{
		Address:  "http://127.0.0.1:8080",
		WaitTime: 30 * time.Second,
	}
}
