package ctapi

import (
	"net/http"
	"time"
)

// Config define CTAPI client config.
type Config struct {
	HttpTransport *http.Transport
	Transport     http.RoundTripper
	Timeout       time.Duration
}

// NewConfig create config with default values.
func NewConfig() *Config {
	return &Config{}
}
