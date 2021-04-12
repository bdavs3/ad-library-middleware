package facebook

type Request struct {
	AccessToken        string
	AdActiveStatus     []string
	AdDeliveryDateMax  string
	AdDeliveryDateMin  string
	AdReachedCountries string
	AdType             string
	Bylines            []string
	DeliveryByRegion   []string
	PotentialReachMax  int
	PotentialReachMin  int
	PublisherPlatforms []string
	SearchPageIDs      []int
	SearchTerms        string
	Limit              int
}
