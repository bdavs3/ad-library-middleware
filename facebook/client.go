package facebook

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	adLibraryURL = "https://graph.facebook.com/v10.0/ads_archive"
	timeout      = 5 * time.Second
	accessToken  = "EAACFomGEN60BABU0dlxZA4ZChhlx4GfSGXRAIW6S2ZBXD0J6kjZACZB0ZCj2adNcyqthZAly5nn8W1UjR9EZBzHYgGZCrfCuW5b6q1fxJ6GZC6jA1cDk26UdswvK5lz6QCo0x2XcMHrXToeQZBgzGQmc9CiZCUBYaLFskV5kTzsUmZABZCOjqL35fn6beZAGvW0JfoECw5YdZAA1snlPNXWZA3o7RyP48ic4l6oLFMzEaeBG9iWAS51GCoIfEXiOF5WfMzCQmIHUZD"
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

func (c *Client) MakeRequest(method string, requestBody io.Reader) (*Response, error) {
	req, err := http.NewRequest(
		method,
		c.BaseURL,
		requestBody,
	)
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s\n%s", http.StatusText(resp.StatusCode), body)
	}

	var response *Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
