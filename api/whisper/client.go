package whisper

import (
	"net/http"
	"os"
)

const (
	DefaultBaseURL = "https://api.openai.com/v1"
	DefaultModel   = "whisper-1"
)

type Client struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
}

type ClientOption func(*Client)

func WithKey(key string) ClientOption {
	return func(c *Client) {
		c.apiKey = key
	}
}

func WithBaseURL(url string) ClientOption {
	return func(c *Client) {
		c.baseURL = url
	}
}

func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

func NewClient(opts ...ClientOption) *Client {
	c := &Client{}
	for _, opt := range opts {
		opt(c)
	}

	if c.apiKey == "" {
		c.apiKey = os.Getenv("OPEN_API_KEY")
	}
	if c.baseURL == "" {
		c.baseURL = os.Getenv("OPENAI_BASE_URL")
	}
	if c.httpClient == nil {
		c.httpClient = http.DefaultClient
	}

	return c
}
