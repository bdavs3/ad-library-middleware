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

	var lookupTables = map[string]struct {
		id  string
		val string
	}{
		"tlkpFundingEntity": {
			id:  "FundingEntityID",
			val: "FundingEntity",
		},
		"tlkpPage": {
			id:  "PageID",
			val: "Page",
		},
		"tlkpAgeRange": {
			id:  "AgeRangeID",
			val: "AgeRange",
		},
		"tlkpGender": {
			id:  "GenderID",
			val: "Gender",
		},
		"tlkpPublisherPlatform": {
			id:  "PublisherPlatformID",
			val: "PublisherPlatform",
		},
		"tlkpRegion": {
			id:  "RegionID",
			val: "Region",
		},
	}

	cacheLayer := database.NewCacheLayer()
	for name := range lookupTables {
		cacheLayer.AddTable(name)
	}

	for i, item := range items {
		fmt.Println(i)

		var lookupData interface{}

		// Update tlkpFundingEntity
		table := lookupTables["tlkpFundingEntity"]
		fundingEntityID, new, err := lookupId("tlkpFundingEntity", item.FundingEntity, table.id, table.val, cacheLayer, conn)
		if err != nil {
			return err
		}

		if new {
			lookupData = &schemas.TlkpFundingEntity{
				ID:            fundingEntityID,
				FundingEntity: item.FundingEntity,
			}
			err = conn.Insert("tlkpFundingEntity", lookupData)
			if err != nil {
				return err
			}
		}

		// Update tlkpPage
		table = lookupTables["tlkpPage"]
		pageID, new, err := lookupId("tlkpPage", item.PageName, table.id, table.val, cacheLayer, conn)
		if err != nil {
			return err
		}

		if new {
			lookupData = &schemas.TlkpPage{
				ID:   pageID,
				Page: item.PageName,
			}
			err = conn.Insert("tlkpPage", lookupData)
			if err != nil {
				return err
			}
		}

		for _, demographic := range item.DemographicDistribution {
			// Update tlkpAgeRange
			table = lookupTables["tlkpAgeRange"]
			ageRangeID, new, err := lookupId("tlkpAgeRange", demographic.Age, table.id, table.val, cacheLayer, conn)
			if err != nil {
				return err
			}

			if new {
				lookupData = &schemas.TlkpAgeRange{
					ID:       ageRangeID,
					AgeRange: demographic.Age,
				}
				err = conn.Insert("tlkpAgeRange", lookupData)
				if err != nil {
					return err
				}
			}

			// Update tlkpGender
			table = lookupTables["tlkpGender"]
			genderID, new, err := lookupId("tlkpGender", demographic.Gender, table.id, table.val, cacheLayer, conn)
			if err != nil {
				return err
			}

			if new {
				lookupData = &schemas.TlkpGender{
					ID:     genderID,
					Gender: demographic.Gender,
				}
				err = conn.Insert("tlkpGender", lookupData)
				if err != nil {
					return err
				}
			}

			// Update tblDemographicDistribution
			percentage, _ := strconv.ParseFloat(demographic.Percentage, 32)
			row := &schemas.TblDemographicDistribution{
				ID:          uuid.NewString(),
				AdLibraryID: item.ID,
				AgeRangeID:  ageRangeID,
				GenderID:    genderID,
				Percentage:  float32(percentage),
			}
			err = conn.Insert("tblDemographicDistribution", row)
			if err != nil {
				return err
			}
		}

		for _, platform := range item.PublisherPlatforms {
			// Update tlkpPublisherPlatform
			table = lookupTables["tlkpPublisherPlatform"]
			publisherPlatformID, new, err := lookupId("tlkpPublisherPlatform", platform, table.id, table.val, cacheLayer, conn)
			if err != nil {
				return err
			}

			if new {
				lookupData = &schemas.TlkpPublisherPlatform{
					ID:                publisherPlatformID,
					PublisherPlatform: platform,
				}
				err = conn.Insert("tlkpPublisherPlatform", lookupData)
				if err != nil {
					return err
				}
			}

			// Update tblPublisherPlatform
			row := &schemas.TblPublisherPlatform{
				ID:          publisherPlatformID,
				AdLibraryID: item.ID,
			}
			err = conn.Insert("tblPublisherPlatform", row)
			if err != nil {
				return err
			}
		}

		for _, region := range item.RegionDistribution {
			// Update tlkpRegion
			table = lookupTables["tlkpRegion"]
			regionID, new, err := lookupId("tlkpRegion", region.Region, table.id, table.val, cacheLayer, conn)
			if err != nil {
				return err
			}

			if new {
				lookupData = &schemas.TlkpRegion{
					ID:     regionID,
					Region: region.Region,
				}
				err = conn.Insert("tlkpRegion", lookupData)
				if err != nil {
					return err
				}
			}

			// Update tblRegionDistribution
			percentage, _ := strconv.ParseFloat(region.Percentage, 32)
			row := &schemas.TblRegionDistribution{
				ID:          uuid.NewString(),
				AdLibraryID: item.ID,
				RegionID:    regionID,
				Percentage:  float32(percentage),
			}
			err = conn.Insert("tblRegionDistribution", row)
			if err != nil {
				return err
			}
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

// lookupId attempts to retrieve the id associated with a given value inside in a lookup table.
// First, it checks the in-memory cache. Then, it makes a network request to look for the value
// in BigQuery (if found here, the cache is updated accordingly). If the value is still not found,
// a new id is created and associated with the value in both BigQuery and the cache. The boolean
// return value indicates whether a new id was created (e.g. a row needs to be inserted to the
// lookup table).
func lookupId(table, value, idCol, valCol string, cl *database.CacheLayer, conn *database.Connection) (string, bool, error) {
	new := false

	id, err := cl.FindValue(table, value)
	if err != nil {
		return "", false, err
	}

	if id == nil {
		iter := conn.Select(table)
		id, err = findValue(iter, value, idCol, valCol)
		if err != nil {
			return "", false, err
		}
		if id == "" {
			new = true
			id = uuid.NewString()
		}
		cl.AddKVPair(table, value, id)
	}

	return id.(string), new, nil
}
