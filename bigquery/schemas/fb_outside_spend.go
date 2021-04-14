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
