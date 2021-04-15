package facebook

// Response is no longer needed now that an SDK is used to consume the paginated API
type Response struct {
	Content []Item `json:"data"`
	Paging  Paging `json:"paging"`
}

// Item represents a political advertisement coming down from the Ad Library. Notice that fields
// like AdCreationDate actually map to JSON fields such as "ad_creation_time". This is because
// full UTC values never seem to come down from the Ad Library, so the database columns have
// been named using "Date" instead of "Time".
type Item struct {
	ID                        string        `json:"id"`
	AdCreationDate            string        `json:"ad_creation_time,omitempty"`
	AdCreativeBody            string        `json:"ad_creative_body,omitempty"`
	AdCreativeLinkCaption     string        `json:"ad_creative_link_caption,omitempty"`
	AdCreativeLinkDescription string        `json:"ad_creative_link_description,omitempty"`
	AdCreativeLinkTitle       string        `json:"ad_creative_link_title,omitempty"`
	AdDeliveryStartDate       string        `json:"ad_delivery_start_time,omitempty"`
	AdDeliveryStopDate        string        `json:"ad_delivery_stop_time,omitempty"`
	AdSnapshotURL             string        `json:"ad_snapshot_url,omitempty"`
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

type Paging struct {
	Cursors Cursors `json:"cursors"`
}

type Cursors struct {
	Before string `json:"before"`
	After  string `json:"after"`
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
