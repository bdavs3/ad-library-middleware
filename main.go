package main

import (
	"ad-library-middleware/facebook"
	"fmt"
	"log"
	"os"
)

func main() {
	req := &facebook.Request{
		AccessToken:        os.Getenv("access_token"),
		SearchTerms:        "california",
		AdReachedCountries: "US",
		Limit:              10,
	}

	client := facebook.NewClient()

	resp, err := client.GetAdLibraryData(req)
	if err != nil {
		log.Fatal(fmt.Sprintf("Err retrieving data:\n%v", err))
	}

	for _, entry := range resp.Content {
		fmt.Printf("%+v\n", entry)
		// This is where BigQuery will be invoked
	}
}
