package main

import (
	"ad-library-middleware/bigquery"
	"ad-library-middleware/facebook"
	"ad-library-middleware/middleware"
	"fmt"
	"log"
	"os"
)

const (
	project = "saguaro-outside-spends"
	dataset = "fb_outside_spend"
)

func main() {
	req := &facebook.Request{
		AccessToken:        os.Getenv("access_token"),
		SearchTerms:        "alaska",
		AdReachedCountries: "US",
		// This is currently the only ad_type supported. But should keep this here
		// in case Facebook decides to support others.
		AdType: "POLITICAL_AND_ISSUE_ADS",
	}

	var after string
	tables := []string{"tblAdLibrary"}

	for {
		resp, err := facebook.NewClient().GetAdLibraryData(req, after)
		if err != nil {
			log.Fatal(fmt.Sprintf("Err retrieving Facebook Ad Library data:\n%v", err))
		}

		conn, err := bigquery.NewConnection(project, dataset)
		if err != nil {
			log.Fatal(fmt.Sprintf("Err connecting to BigQuery:\n%v", err))
		}

		var tableData interface{}

		for _, tableName := range tables {
			tableData, err = middleware.InterpretResponse(resp, tableName)
			if err != nil {
				log.Fatal(fmt.Sprintf("Could not interpet Facebook response into %v:\n%v", tableName, err))
			}
			conn.Insert(tableName, tableData)
			if err != nil {
				log.Fatal(fmt.Sprintf("Err inserting to %v:\n%v", tableName, err))
			}
		}

		after = resp.Paging.Cursors.After
		if after == "" {
			break
		}
	}
}
