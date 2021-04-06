package facebook

import (
	"net/http"
	"time"
)

const (
	adLibraryURL = "https://graph.facebook.com/v10.0/ads_archive"
	timeout      = 5 * time.Second
)

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewClient() *Client {
	return &Client{
		BaseURL:    adLibraryURL,
		HTTPClient: &http.Client{Timeout: timeout},
	}
}

func (c *Client) String() string {
	return c.BaseURL
}
