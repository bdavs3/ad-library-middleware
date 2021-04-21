package facebook

// Request is a struct that represents the possible querystring parameters that can be used
// in a call to the Facebook Ad Library. For more information, see the section on search
// parameters here: https://www.facebook.com/ads/library/api/?source=archive-landing-page
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
