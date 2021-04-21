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

// UploadResponseData uses the given items from the Facebook Ad Library to update BigQuery.
func UploadResponseData(items []*facebook.Item, conn *database.Connection) error {
	var rows []*schemas.TblAdLibrary

	for _, item := range items {
		// Update tlkpFundingEntity
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

		// Update tlkpPage
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

		// Update tlkpAgeRange, tlkpGender, and tblDemographicDistribution
		for _, demographic := range item.DemographicDistribution {
			iter := conn.Select("tlkpAgeRange")
			ageRangeID, err := findValue(iter, demographic.Age, "AgeRange", "AgeRangeID")
			if err != nil {
				return err
			}
			if ageRangeID == "" {
				ageRangeID = uuid.NewString()
				row := &schemas.TlkpAgeRange{
					ID:       ageRangeID,
					AgeRange: demographic.Age,
				}
				conn.Insert("tlkpAgeRange", row)
			}

			iter = conn.Select("tlkpGender")
			genderID, err := findValue(iter, demographic.Gender, "Gender", "GenderID")
			if err != nil {
				return err
			}
			if genderID == "" {
				genderID = uuid.NewString()
				row := &schemas.TlkpGender{
					ID:     genderID,
					Gender: demographic.Gender,
				}
				conn.Insert("tlkpGender", row)
			}

			percentage, _ := strconv.ParseFloat(demographic.Percentage, 32)
			row := &schemas.TblDemographicDistribution{
				ID:          uuid.NewString(),
				AdLibraryID: item.ID,
				AgeRangeID:  ageRangeID,
				GenderID:    genderID,
				Percentage:  float32(percentage),
			}
			conn.Insert("tblDemographicDistribution", row)
		}

		// Update tlkpPublisherPlatform and tblPublisherPlatform
		for _, platform := range item.PublisherPlatforms {
			iter := conn.Select("tlkpPublisherPlatform")
			publisherPlatformID, err := findValue(iter, platform, "PublisherPlatform", "PublisherPlatformID")
			if err != nil {
				return err
			}
			if publisherPlatformID == "" {
				publisherPlatformID = uuid.NewString()
				row := &schemas.TlkpPublisherPlatform{
					ID:                publisherPlatformID,
					PublisherPlatform: platform,
				}
				conn.Insert("tlkpPublisherPlatform", row)
			}
			row := &schemas.TblPublisherPlatform{
				ID:          publisherPlatformID,
				AdLibraryID: item.ID,
			}
			conn.Insert("tblPublisherPlatform", row)
		}

		// Update tlkpRegion and tblRegionDistribution
		for _, region := range item.RegionDistribution {
			iter = conn.Select("tlkpRegion")
			regionID, err := findValue(iter, region.Region, "Region", "RegionID")
			if err != nil {
				return err
			}
			if regionID == "" {
				regionID = uuid.NewString()
				row := &schemas.TlkpRegion{
					ID:     regionID,
					Region: region.Region,
				}
				conn.Insert("tlkpRegion", row)
			}

			percentage, _ := strconv.ParseFloat(region.Percentage, 32)
			row := &schemas.TblRegionDistribution{
				ID:          uuid.NewString(),
				AdLibraryID: item.ID,
				RegionID:    regionID,
				Percentage:  float32(percentage),
			}
			conn.Insert("tblRegionDistribution", row)
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

	// Update tblAdLibrary
	err := conn.Insert("tblAdLibrary", rows)
	if err != nil {
		return err
	}

	return nil
}

// UploadBasic is a simplified version of UploadResponseData, which uses the given
// Items from the Facebook Ad Library to update BigQuery. This function only updates
// tblAdLibrary in order to reduce runtime. It is intended to be used for testing until
// a caching layer is used in conjunction with UploadResponseData.
func UploadBasic(items []*facebook.Item, conn *database.Connection) error {
	var rows []*schemas.TblAdLibrary

	for _, item := range items {
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

	// Update tblAdLibrary
	err := conn.Insert("tblAdLibrary", rows)
	if err != nil {
		return err
	}

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
