package middleware

import (
	"ad-library-middleware/database"
	"ad-library-middleware/database/schemas"
	"ad-library-middleware/facebook"
	"fmt"
	"strconv"

	"cloud.google.com/go/bigquery"
	"github.com/google/uuid"
	"google.golang.org/api/iterator"
)

// UploadResponseData uses the given Response from the Facebook Ad Library to return a slice of
// structs that conform to the schema of the given table.
func UploadResponseData(resp *facebook.Response, conn *database.Connection) error {
	var rows []*schemas.TblAdLibrary

	for _, item := range resp.Content {
		iter := conn.Select("tlkpFundingEntity")
		fundingEntityID, err := findValue(iter, item.FundingEntity, "FundingEntity", "FundingEntityID")
		if err != nil {
			return err
		}
		if fundingEntityID == "" {
			fundingEntityID = uuid.NewString()
			row := &schemas.TlkpFundingEntity{
				ID:            fundingEntityID,
				FundingEntity: item.FundingEntity,
			}
			conn.Insert("tlkpFundingEntity", row)
		}

		iter = conn.Select("tlkpPage")
		pageID, err := findValue(iter, item.PageName, "Page", "PageID")
		if err != nil {
			return err
		}
		if pageID == "" {
			pageID = item.PageID // Facebook supplies PageID, so no need for UUID
			row := &schemas.TlkpPage{
				ID:   pageID,
				Page: item.PageName,
			}
			conn.Insert("tlkpPage", row)
		}

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
			FundingEntityID:           fundingEntityID,
			ImpressionsLower:          impressionsLower,
			ImpressionsUpper:          impressionsUpper,
			PageID:                    pageID,
			PotentialReachLower:       potentialReachLower,
			PotentialReachUpper:       potentialReachUpper,
			SpendLower:                spendLower,
			SpendUpper:                spendUpper,
		}

		rows = append(rows, row)
	}

	conn.Insert("tblAdLibrary", rows)

	return nil
}

// findValue looks for the given value using a RowIterator and column name. If the value is
// found, the value from the specified idCol is returned. Otherwise, the empty string is returned.
func findValue(iter *bigquery.RowIterator, value, searchCol, idCol string) (string, error) {
	for {
		var values map[string]bigquery.Value
		err := iter.Next(&values)
		if values[searchCol] == value {
			return fmt.Sprintf("%v", values[idCol]), nil
		}
		if err == iterator.Done {
			break
		}
		if err != nil {
			return "", err
		}
	}

	return "", nil
}
