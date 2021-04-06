package facebook

type Request struct {
	AdActiveStatus     []string `json:"ad_active_status"`
	AdDeliveryDateMax  string   `json:"ad_delivery_date_max"`
	AdDeliveryDateMin  string   `json:"ad_delivery_date_min"`
	AdReachedCountries []string `json:"ad_reached_countries"`
	AdType             string   `json:"ad_type"`
	Bylines            []string `json:"bylines"`
	DeliveryByRegion   []string `json:"delivery_by_region"`
	PotentialReachMax  int      `json:"potential_reach_max"`
	PotentialReachMin  int      `json:"potential_reach_min"`
	PublisherPlatforms []string `json:"publisher_platforms"`
	SearchPageIDs      []int    `json:"search_page_ids"`
	SearchTerms        []string `json:"search_terms"`
}
