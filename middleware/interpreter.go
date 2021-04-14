package middleware

import (
	"ad-library-middleware/bigquery/schemas"
	"ad-library-middleware/facebook"
	"errors"
	"strconv"

	"github.com/google/uuid"
)

// InterpretResponse uses the given Response from the Facebook Ad Library to return a slice of
// structs that conform to the schema of the given table.
func InterpretResponse(resp *facebook.Response, tableName string) (interface{}, error) {
	switch tableName {
	case "tblAdLibrary":
		var rows []*schemas.TblAdLibrary

		for _, item := range resp.Content {
			impressionsLower, _ := strconv.Atoi(item.Impressions.LowerBound)
			impressionsUpper, _ := strconv.Atoi(item.Impressions.UpperBound)
			potentialReachLower, _ := strconv.Atoi(item.PotentialReach.LowerBound)
			potentialReachUpper, _ := strconv.Atoi(item.PotentialReach.UpperBound)
			spendLower, _ := strconv.Atoi(item.Spend.LowerBound)
			spendUpper, _ := strconv.Atoi(item.Spend.UpperBound)

			row := &schemas.TblAdLibrary{
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

		return rows, nil
	default:
		return nil, errors.New("incorrect table name supplied")
	}
}
