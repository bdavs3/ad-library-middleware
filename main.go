package main

import (
	"ad-library-middleware/facebook"
	"ad-library-middleware/middleware"
	"fmt"
	"log"
	"os"
)

func main() {
	req := &facebook.Request{
		AccessToken:        os.Getenv("access_token"),
		SearchTerms:        "arizona",
		AdReachedCountries: "US",
		// This is currently the only ad_type supported. But want to keep this here
		// in case Facebook adds others.
		AdType: "POLITICAL_AND_ISSUE_ADS",
	}

	var after string

	for {
		resp, err := middleware.GetAdLibraryData(req, after)
		if err != nil {
			log.Fatal(fmt.Sprintf("Err retrieving FB data:\n%v", err))
		}

		err = middleware.InsertAdLibraryData(resp)
		if err != nil {
			log.Fatal(fmt.Sprintf("Err inserting to Bigquery:\n%v", err))
		}

		after = resp.Paging.Cursors.After
		if after == "" {
			break
		}
	}
}
