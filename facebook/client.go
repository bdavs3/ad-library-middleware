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
	adLibraryURL  = "https://graph.facebook.com/v10.0/ads_archive?"
	searchPattern = "fields=%v&access_token=%v&search_terms=%v&ad_reached_countries=%v&ad_type=%v"
	defaultFields = "['id','ad_creation_time','ad_creative_body','ad_creative_link_caption','ad_creative_link_description','ad_creative_link_title','ad_delivery_start_time','ad_delivery_stop_time','ad_snapshot_url','currency','demographic_distribution','funding_entity','impressions','page_id','page_name','potential_reach','publisher_platforms','region_distribution','spend']"
	timeout       = 5 * time.Second
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

func (c *Client) GetAdLibraryData(req *Request, after string) (*Response, error) {
	params := fmt.Sprintf(
		searchPattern,
		defaultFields,
		req.AccessToken,
		req.SearchTerms,
		req.AdReachedCountries,
		req.AdType,
	)

	// Handles pagination
	if after != "" {
		params += fmt.Sprintf("&after=%v", after)
	}

	response, err := c.MakeRequest(
		http.MethodGet,
		params,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) MakeRequest(method string, parameters string, requestBody io.Reader) (*Response, error) {
	req, err := http.NewRequest(
		method,
		c.BaseURL+parameters,
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
