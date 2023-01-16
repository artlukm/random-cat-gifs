// Код структуры клиента и функции NewClient
package lib

import (
	"net/http"
	"strings"
)

const (
	defaultBaseURL = "https://randomcatgifs.com"
	defaultTempDir = "temp"
)

type Client struct {
	HTTPClient *http.Client
	BaseURL    string
	TempDir    string
	UserAgent  string
	Debug      bool
}

type ClientOption func(*Client)

// HTTPClient sets HTTPClient in Client
func HTTPClient(cl *http.Client) ClientOption {
	return ClientOption(
		func(c *Client) {
			c.HTTPClient = cl
		},
	)
}

// BaseURL sets BaseURL in Client
func BaseURL(url string) ClientOption {
	return ClientOption(
		func(c *Client) {
			c.BaseURL = url
		},
	)
}

// TempDir sets TempDir in Client
func TempDir(path string) ClientOption {
	return ClientOption(
		func(c *Client) {
			c.TempDir = strings.TrimSuffix(path, "/")
		},
	)
}

// UserAgent sets UserAgent in Client
func UserAgent(ua string) ClientOption {
	return ClientOption(
		func(c *Client) {
			c.UserAgent = ua
		},
	)
}

// NewClient returns *Client
func NewClient(opts ...ClientOption) *Client {
	cl := &Client{
		HTTPClient: http.DefaultClient,
		BaseURL:    defaultBaseURL,
		TempDir:    defaultTempDir,
	}

	for _, opt := range opts {
		opt(cl)
	}

	return cl
}
