package middleware

import (
	"ad-library-middleware/facebook"
	"ad-library-middleware/google"
	"ad-library-middleware/schemas"
	"strconv"

	"github.com/google/uuid"
)

const (
	project = "saguaro-outside-spends"
	dataset = "fb_outside_spend"
)

// GetAdLibraryData uses parameters from a Request object to return a Response struct
// containing information from the Facebook ad library.
func GetAdLibraryData(req *facebook.Request) (*facebook.Response, error) {
	client := facebook.NewClient()

	resp, err := client.GetAdLibraryData(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// InsertAdLibraryData inserts data from a Facebook client response into the relevant
// Bigquery table.
func InsertAdLibraryData(resp *facebook.Response) error {
	conn, err := google.NewConnection(project)
	if err != nil {
		return err
	}

	rows := unpackResponseContent(resp)
	if err != nil {
		return err
	}

	err = conn.Insert(dataset, "tblAdLibrary", rows)
	if err != nil {
		return err
	}

	return nil
}

func unpackResponseContent(resp *facebook.Response) []*schemas.AdLibraryData {
	var rows []*schemas.AdLibraryData

	for _, item := range resp.Content {
		impressionsLower, _ := strconv.Atoi(item.Impressions.LowerBound)
		impressionsUpper, _ := strconv.Atoi(item.Impressions.UpperBound)
		potentialReachLower, _ := strconv.Atoi(item.PotentialReach.LowerBound)
		potentialReachUpper, _ := strconv.Atoi(item.PotentialReach.UpperBound)
		spendLower, _ := strconv.Atoi(item.Spend.LowerBound)
		spendUpper, _ := strconv.Atoi(item.Spend.UpperBound)

		row := &schemas.AdLibraryData{
			ID:                        item.ID,
			AdCreationDate:            item.AdCreationDate,
			AdCreativeBody:            item.AdCreativeBody,
			AdCreativeLinkCaption:     item.AdCreativeLinkCaption,
			AdCreativeLinkDescription: item.AdCreativeLinkDescription,
			AdCreativeLinkTitle:       item.AdCreativeLinkTitle,
			AdDeliveryStartDate:       item.AdDeliveryStartDate,
			AdDeliveryStopDate:        item.AdDeliveryStopDate,
			AdSnapshotURL:             item.AdSnapshotURL,
			CurrencyID:                uuid.NewString(),
			FundingEntityID:           uuid.NewString(),
			ImpressionsLower:          impressionsLower,
			ImpressionsUpper:          impressionsUpper,
			PageID:                    uuid.NewString(),
			PotentialReachLower:       potentialReachLower,
			PotentialReachUpper:       potentialReachUpper,
			SpendLower:                spendLower,
			SpendUpper:                spendUpper,
		}

		rows = append(rows, row)
	}

	return rows
}
