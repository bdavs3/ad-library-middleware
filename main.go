package main

import (
	"ad-library-middleware/facebook"
	"fmt"
	"log"
)

func main() {
	req := &facebook.Request{
		AccessToken:        "EAACFomGEN60BAEOHY15xXGxZBg69YtAvyzbM5Qj8MIv3nPPP7vZB70Ij0pNnQalwZAZAp9tlDnDavWHeEF79kUSFndlxTZCPT64LQlpU4WBk4UOiV7DRLwXW0KHecLHb68JyvqFRUBC9Ix4Ka4yeFiUhlYf2TI6q6OWnaZB4dxT0mle1KieYH1WZB8YYyUZB8fZCbybC1pZA2MLhGAUrJWtDUrpRuIdQBdZAZBtMlX1CBieZCI8ZCEKQxeCdpxdzzmCxlvgbgZD",
		SearchTerms:        "california",
		AdReachedCountries: "US",
	}

	// v, err := query.Values(req)
	// if err != nil {
	// 	log.Fatal(fmt.Sprintf("Err in querystring converstion:\n%v", err))
	// }
	// fmt.Println(v)

	// params := v.Encode()
	// fmt.Println(params)

	client := facebook.NewClient()

	resp, err := client.GetAdLibraryData(req)
	if err != nil {
		log.Fatal(fmt.Sprintf("Err retrieving data:\n%v", err))
	}

	for _, entry := range resp.Content {
		fmt.Println(entry)
	}
}
