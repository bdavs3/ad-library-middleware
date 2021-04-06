package facebook

type Request struct {
	AccessToken        string   `url:"access_token,omitempty"`
	AdActiveStatus     []string `url:"ad_active_status,omitempty"`
	AdDeliveryDateMax  string   `url:"ad_delivery_date_max,omitempty"`
	AdDeliveryDateMin  string   `url:"ad_delivery_date_min,omitempty"`
	AdReachedCountries string   `url:"ad_reached_countries"`
	AdType             string   `url:"ad_type,omitempty"`
	Bylines            []string `url:"bylines,omitempty"`
	DeliveryByRegion   []string `url:"delivery_by_region,omitempty"`
	PotentialReachMax  int      `url:"potential_reach_max,omitempty"`
	PotentialReachMin  int      `url:"potential_reach_min,omitempty"`
	PublisherPlatforms []string `url:"publisher_platforms,omitempty"`
	SearchPageIDs      []int    `url:"search_page_ids,omitempty"`
	SearchTerms        string   `url:"search_terms,omitempty"`
}
