package main

import (
	"ad-library-middleware/facebook"
	"fmt"
	"log"
)

func main() {
	req := &facebook.Request{
		AccessToken:        "EAACFomGEN60BAIwSbPB8W07U9xXh8TfaCZAW5tiUu9jfXKQBbe4vutM7RQegfxl8qA1ib67WEHdZAqqQXqAH1GMqb0XwauTftLBwxiKGIauSEDWpQkQTE8X79CGfnrHL3ZCYoWAECXy2LFPlQA5F9ZCvr7LUBMinJXtTLlh3boodrArVPPJ4i387atoKt0ZADA6Jfn1ZBHrgi6UlimJTPgWkls6ZC0qzu34CzCwXceJKJulRSbVR7IPSeqfpXTehVoZD",
		SearchTerms:        "california",
		AdReachedCountries: "US",
	}

	client := facebook.NewClient()

	resp, err := client.GetAdLibraryData(req)
	if err != nil {
		log.Fatal(fmt.Sprintf("Err retrieving data:\n%v", err))
	}

	for _, entry := range resp.Content {
		fmt.Println(entry) // This is where BigQuery will be invoked
	}
}
