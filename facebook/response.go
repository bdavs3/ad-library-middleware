package facebook

type Response struct {
	Content []Entry `json:"data"`
}

type Entry struct {
	ID                        string                 `json:"id"`
	AdCreationTime            string                 `json:"ad_creation_time,omitempty"`
	AdCreativeBody            string                 `json:"ad_creative_body,omitempty"`
	AdCreativeLinkCaption     string                 `json:"ad_creative_link_caption,omitempty"`
	AdCreativeLinkDescription string                 `json:"ad_creative_link_description,omitempty"`
	AdCreativeLinkTitle       string                 `json:"ad_creative_link_title,omitempty"`
	AdDeliveryStartTime       string                 `json:"ad_delivery_start_time,omitempty"`
	AdDeliveryStopTime        string                 `json:"ad_delivery_stop_time,omitempty"`
	AdSnapshotURL             string                 `json:"ad_snapshot_url,omitempty"`
	Currency                  string                 `json:"currency,omitempty"`
	DemographicDistribution   []AudienceDistribution `json:"demographic_distribution,omitempty"`
	FundingEntity             string                 `json:"funding_entity,omitempty"`
	Impressions               InsightsRange          `json:"impressions,omitempty"`
	PageID                    string                 `json:"page_id,omitempty"`
	PageName                  string                 `json:"page_name,omitempty"`
	PotentialReach            InsightsRange          `json:"potential_reach,omitempty"`
	PublisherPlatforms        []string               `json:"publisher_platforms,omitempty"`
	RegionDistribution        []AudienceDistribution `json:"region_distribution,omitempty"`
	Spend                     InsightsRange          `json:"spend,omitempty"`
}

type AudienceDistribution struct {
	Age        string
	Gender     string
	Percentage string
	Region     string
}

type InsightsRange struct {
	LowerBound string
	UpperBound string
}
