package facebook

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	fb "github.com/huandu/facebook/v2"
)

const (
	adLibraryEndpoint    = "/ads_archive"
	tokenRefreshEndpoint = "/oauth/access_token"
	defaultFields        = "['id','ad_creation_time','ad_creative_body','ad_creative_link_caption','ad_creative_link_description','ad_creative_link_title','ad_delivery_start_time','ad_delivery_stop_time','ad_snapshot_url','demographic_distribution','funding_entity','impressions','page_id','page_name','potential_reach','publisher_platforms','region_distribution','spend']"
)

// Credentials contain the information necessary for pulling data from the Facebook API, as well
// as refreshing access tokens.
type Credentials struct {
	File        string
	AppID       string `json:"app_id"`
	AppSecret   string `json:"app_secret"`
	AccessToken string `json:"access_token"`
}

// NewCredentials unmarshals the given JSON file into the returned Credentials struct.
func NewCredentials(file string) (*Credentials, error) {
	jsonFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var credentials *Credentials
	json.Unmarshal(bytes, &credentials)
	credentials.File = file

	return credentials, nil
}

// Sdk is used to interface with the Facebook API.
type Sdk struct {
	Session     *fb.Session
	Credentials *Credentials
}

// NewSdk creates an Sdk instance using the given Credentials.
func NewSdk(c *Credentials) *Sdk {
	var globalApp = fb.New(c.AppID, c.AppSecret)
	s := globalApp.Session(c.AccessToken)

	return &Sdk{
		Session:     s,
		Credentials: c,
	}
}

// GetAdLibraryData makes a call to the Facebook Ad Library and returns the Items retrieved
// using the given request.
func (sdk *Sdk) GetAdLibraryData(req *Request) ([]*Item, error) {
	result, err := sdk.Session.Get(adLibraryEndpoint, fb.Params{
		"fields": defaultFields,

		"ad_delivery_date_max": req.AdDeliveryDateMax,
		"ad_delivery_date_min": req.AdDeliveryDateMin,
		"ad_reached_countries": req.AdReachedCountries,
		"ad_type":              req.AdType,
		"search_terms":         req.SearchTerms,
	})
	if err != nil {
		return nil, err
	}

	paging, err := result.Paging(sdk.Session)
	if err != nil {
		return nil, err
	}

	var results []fb.Result
	var items []*Item

	for {
		results = append(results, paging.Data()...)

		done, err := paging.Next()
		if err != nil {
			return nil, err
		}
		if done {
			break
		}
	}

	for _, result := range results {
		bytes, err := json.Marshal(result)
		if err != nil {
			return nil, err
		}

		var item *Item
		json.Unmarshal(bytes, &item)
		items = append(items, item)
	}

	return items, nil
}

// StoreRefreshToken retrieves a long-lived (60 day) token using the access token stored in the
// Sdk's Credentials. It then replaces the original access token with the new one.
func (sdk *Sdk) StoreRefreshToken() error {
	result, err := sdk.Session.Get(tokenRefreshEndpoint, fb.Params{
		"grant_type":        "fb_exchange_token",
		"client_id":         sdk.Credentials.AppID,
		"client_secret":     sdk.Credentials.AppSecret,
		"fb_exchange_token": sdk.Credentials.AccessToken,
	})
	if err != nil {
		return err
	}

	sdk.Credentials.AccessToken = fmt.Sprintf("%v", result["access_token"])

	bytes, err := json.MarshalIndent(sdk.Credentials, "", "    ")
	if err != nil {
		return err
	}

	err = os.WriteFile(sdk.Credentials.File, bytes, 0644)
	if err != nil {
		return err
	}

	return nil
}
