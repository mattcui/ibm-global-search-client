package main

import (
	"context"
	"fmt"
	"os"

	igsc "github.com/mattcui/ibm-global-search-client"
)

func main() {
	token := "xxxxxxxxxx"
	ctx := context.Background()
	ctx = context.WithValue(ctx, igsc.ContextAccessToken, token)

	body := *igsc.NewScanBody() // ScanBody | It contains the query filters on the first operation call, or the search_cursor on next calls. On subsequent calls, set the `search_cursor` to the value returned by the previous call. After the first, you must set only the `search_cursor`. Any other parameter but the `search_cursor` will be ignored. The `search_cursor` encodes all the information needed to get the next batch of `limit` data.
	limit := int32(10)          // int32 | The maximum number of hits to return. Defaults to 10. (optional) (default to 10)
	timeout := int32(300)       // int32 | A search timeout, bounding the search request to be executed within the specified time value and bail with the hits accumulated up to that point when expired. Defaults to the system defined timeout. (optional) (default to 0)
	//sort := []string{"Inner_example"} // []string | Comma separated properties names used for sorting (optional)

	cfg := igsc.NewConfiguration()
	client := igsc.NewAPIClient(cfg)
	body.SetQuery("*")

	resp, r, err := client.SearchApi.Search(ctx).Body(body).Limit(limit).Timeout(timeout).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.Search``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Search`: ScanResult
	fmt.Fprintf(os.Stdout, "Response from `SearchApi.Search`: %v\n", resp)

}
