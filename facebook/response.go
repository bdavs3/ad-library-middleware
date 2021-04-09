package facebook

type Response struct {
	Content []Entry `json:"data"`
}

type Entry struct {
	ID                        string        `json:"id"`
	AdCreationTime            string        `json:"ad_creation_time,omitempty"`
	AdCreativeBody            string        `json:"ad_creative_body,omitempty"`
	AdCreativeLinkCaption     string        `json:"ad_creative_link_caption,omitempty"`
	AdCreativeLinkDescription string        `json:"ad_creative_link_description,omitempty"`
	AdCreativeLinkTitle       string        `json:"ad_creative_link_title,omitempty"`
	AdDeliveryStartTime       string        `json:"ad_delivery_start_time,omitempty"`
	AdDeliveryStopTime        string        `json:"ad_delivery_stop_time,omitempty"`
	AdSnapshotURL             string        `json:"ad_snapshot_url,omitempty"`
	Currency                  string        `json:"currency,omitempty"`
	DemographicDistribution   []Demographic `json:"demographic_distribution,omitempty"`
	FundingEntity             string        `json:"funding_entity,omitempty"`
	Impressions               InsightsRange `json:"impressions,omitempty"`
	PageID                    string        `json:"page_id,omitempty"`
	PageName                  string        `json:"page_name,omitempty"`
	PotentialReach            InsightsRange `json:"potential_reach,omitempty"`
	PublisherPlatforms        []string      `json:"publisher_platforms,omitempty"`
	RegionDistribution        []Region      `json:"region_distribution,omitempty"`
	Spend                     InsightsRange `json:"spend,omitempty"`
}

type Demographic struct {
	Age        string `json:"age,omitempty"`
	Gender     string `json:"gender,omitempty"`
	Percentage string `json:"percentage,omitempty"`
}

type Region struct {
	Region     string `json:"region,omitempty"`
	Percentage string `json:"percentage,omitempty"`
}

type InsightsRange struct {
	LowerBound string `json:"lower_bound,omitempty"`
	UpperBound string `json:"upper_bound,omitempty"`
}
