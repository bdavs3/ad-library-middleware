package schemas

type TblAdLibrary struct {
	ID                        string `bigquery:"AdLibraryID"`
	AdCreationDate            string `bigquery:"AdCreationDate"`
	AdCreativeBody            string `bigquery:"AdCreativeBody"`
	AdCreativeLinkCaption     string `bigquery:"AdCreativeLinkCaption"`
	AdCreativeLinkDescription string `bigquery:"AdCreativeLinkDescription"`
	AdCreativeLinkTitle       string `bigquery:"AdCreativeLinkTitle"`
	AdDeliveryStartDate       string `bigquery:"AdDeliveryStartDate"`
	AdDeliveryStopDate        string `bigquery:"AdDeliveryStopDate"`
	AdSnapshotURL             string `bigquery:"AdSnapshotURL"`
	CurrencyID                string `bigquery:"CurrencyID"`
	FundingEntityID           string `bigquery:"FundingEntityID"`
	ImpressionsLower          int    `bigquery:"ImpressionsLower"`
	ImpressionsUpper          int    `bigquery:"ImpressionsUpper"`
	PageID                    string `bigquery:"PageID"`
	PotentialReachLower       int    `bigquery:"PotentialReachLower"`
	PotentialReachUpper       int    `bigquery:"PotentialReachUpper"`
	SpendLower                int    `bigquery:"SpendLower"`
	SpendUpper                int    `bigquery:"SpendUpper"`
}

type TblDemographicDistribution struct {
	ID          string  `bigquery:"DemographicDistributionID"`
	AdLibraryID string  `bigquery:"AdLibraryID"`
	AgeRangeID  string  `bigquery:"AgeRangeID"`
	GenderID    string  `bigquery:"GenderID"`
	Percentage  float32 `bigquery:"Percentage"`
}

type TblPublisherPlatform struct {
	ID          string `bigquery:"PublisherPlatformID"`
	AdLibraryID string `bigquery:"AdLibraryID"`
}

type TblRegionDistribution struct {
	ID          string  `bigquery:"RegionDistributionID"`
	AdLibraryID string  `bigquery:"AdLibraryID"`
	RegionID    string  `bigquery:"RegionID"`
	Percentage  float32 `bigquery:"Percentage"`
}

type TlkpAgeRange struct {
	ID       string `bigquery:"AgeRangeID"`
	AgeRange string `bigquery:"AgeRange"`
}

type TlkpCurrency struct {
	ID       string `bigquery:"CurrencyID"`
	Currency string `bigquery:"Currency"`
}

type TlkpFundingEntity struct {
	ID            string `bigquery:"FundingEntityID"`
	FundingEntity string `bigquery:"FundingEntity"`
}

type TlkpGender struct {
	ID     string `bigquery:"GenderID"`
	Gender string `bigquery:"Gender"`
}

type TlkpPage struct {
	ID   string `bigquery:"PageID"`
	Page string `bigquery:"Page"`
}

type TlkpPublisherPlatform struct {
	ID                string `bigquery:"PublisherPlatformID"`
	PublisherPlatform string `bigquery:"PublisherPlatform"`
}

type TlkpRegion struct {
	ID     string `bigquery:"RegionID"`
	Region string `bigquery:"Region"`
}
