package main

import (
	"ad-library-middleware/database"
	"ad-library-middleware/facebook"
	"ad-library-middleware/middleware"
	"fmt"
	"log"
)

const (
	appCredentials = "fb-credentials.json"
	project        = "saguaro-outside-spends"
	dataset        = "fb_outside_spend"
)

func main() {
	req := &facebook.Request{
		AdReachedCountries: "US",
		// This is currently the only ad_type supported. But should keep this here
		// in case Facebook decides to support others.
		AdType:      "POLITICAL_AND_ISSUE_ADS",
		SearchTerms: "california",
	}

	credentials, err := facebook.NewCredentials(appCredentials)
	if err != nil {
		log.Fatal(fmt.Sprintf("Err creating Facebook credentials:\n%v", err))
	}

	sdk := facebook.NewSdk(credentials)

	items, err := sdk.GetAdLibraryData(req)
	if err != nil {
		log.Fatal(fmt.Sprintf("Err retrieving Facebook Ad Library data:\n%v", err))
	}

	conn, err := database.NewConnection(project, dataset)
	if err != nil {
		log.Fatal(fmt.Sprintf("Err connecting to BigQuery:\n%v", err))
	}

	// err = middleware.UploadResponseData(items, conn)
	err = middleware.UploadBasic(items, conn)
	if err != nil {
		log.Fatal(fmt.Sprintf("Err in UploadResponseData:\n%v", err))
	}

	err = sdk.StoreRefreshToken()
	if err != nil {
		log.Fatal(fmt.Sprintf("Err in StoreRefreshToken:\n%v", err))
	}
}
