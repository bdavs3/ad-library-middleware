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
		SearchTerms:        "california",
		AdReachedCountries: "US",
		Limit:              25,
	}

	resp, err := middleware.GetAdLibraryData(req)
	if err != nil {
		log.Fatal(fmt.Sprintf("Err retrieving FB data:\n%v", err))
	}

	err = middleware.InsertAdLibraryData(resp)
	if err != nil {
		log.Fatal(fmt.Sprintf("Err inserting to Bigquery:\n%v", err))
	}
}
