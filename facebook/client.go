package facebook

import (
	"encoding/json"
	"io/ioutil"
	"os"

	fb "github.com/huandu/facebook/v2"
)

const (
	adLibraryEndpoint = "/ads_archive"
	defaultFields     = "['id','ad_creation_time','ad_creative_body','ad_creative_link_caption','ad_creative_link_description','ad_creative_link_title','ad_delivery_start_time','ad_delivery_stop_time','ad_snapshot_url','demographic_distribution','funding_entity','impressions','page_id','page_name','potential_reach','publisher_platforms','region_distribution','spend']"
)

type Credentials struct {
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

func GetCredentials(file string) (*Credentials, error) {
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

	return credentials, nil
}

type Sdk struct {
	Session *fb.Session
}

func NewSdk(cred *Credentials, access_token string) *Sdk {
	var globalApp = fb.New(cred.AppID, cred.AppSecret)
	s := globalApp.Session(access_token)

	return &Sdk{
		Session: s,
	}
}

func (sdk *Sdk) GetAdLibraryData(req *Request) ([]*Item, error) {
	result, err := sdk.Session.Get(adLibraryEndpoint, fb.Params{
		"access_token": req.AccessToken,
		"fields":       defaultFields,

		"ad_delivery_date_max": req.AdDeliveryDateMax,
		"ad_delivery_date_min": req.AdDeliveryDateMin,
		"ad_reached_countries": req.AdReachedCountries,
		"ad_type":              req.AdType,
		"search_terms":         req.SearchTerms,
	})
	if err != nil {
		return nil, err
	}

	paging, _ := result.Paging(sdk.Session)

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
