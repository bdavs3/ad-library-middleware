package main

import (
	"ad-library-middleware/database"
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
		AdReachedCountries: "US",
		// This is currently the only ad_type supported. But should keep this here
		// in case Facebook decides to support others.
		AdType:      "POLITICAL_AND_ISSUE_ADS",
		SearchTerms: "california",
	}

	credentials, err := facebook.GetCredentials("fb-credentials.json")
	if err != nil {
		log.Fatal(fmt.Sprintf("Err creating Facebook credentials:\n%v", err))
	}

	items, err := facebook.NewSdk(credentials, req.AccessToken).GetAdLibraryData(req)
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
}
