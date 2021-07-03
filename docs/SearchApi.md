# \SearchApi

All URIs are relative to *https://api.global-search-tagging.cloud.ibm.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Search**](SearchApi.md#Search) | **Post** /v3/resources/search | Find instances of resources (v3)
[**V2Search**](SearchApi.md#V2Search) | **Post** /v2/resources/search | Find instances of resources (v2)



## Search

> ScanResult Search(ctx).Body(body).TransactionId(transactionId).AccountId(accountId).Limit(limit).Timeout(timeout).Sort(sort).Execute()

Find instances of resources (v3)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewScanBody() // ScanBody | It contains the query filters on the first operation call, or the search_cursor on next calls. On subsequent calls, set the `search_cursor` to the value returned by the previous call. After the first, you must set only the `search_cursor`. Any other parameter but the `search_cursor` will be ignored. The `search_cursor` encodes all the information needed to get the next batch of `limit` data.
    transactionId := "transactionId_example" // string | An aplhanumeric string that can be used to trace a request across services. If not specified it will be automatically generated with the prefix \"gst-\" (optional)
    accountId := "accountId_example" // string | The account ID to filter resources. (optional)
    limit := int32(56) // int32 | The maximum number of hits to return. Defaults to 10. (optional) (default to 10)
    timeout := int32(56) // int32 | A search timeout, bounding the search request to be executed within the specified time value and bail with the hits accumulated up to that point when expired. Defaults to the system defined timeout. (optional) (default to 0)
    sort := []string{"Inner_example"} // []string | Comma separated properties names used for sorting (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SearchApi.Search(context.Background()).Body(body).TransactionId(transactionId).AccountId(accountId).Limit(limit).Timeout(timeout).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.Search``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Search`: ScanResult
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.Search`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ScanBody**](ScanBody.md) | It contains the query filters on the first operation call, or the search_cursor on next calls. On subsequent calls, set the &#x60;search_cursor&#x60; to the value returned by the previous call. After the first, you must set only the &#x60;search_cursor&#x60;. Any other parameter but the &#x60;search_cursor&#x60; will be ignored. The &#x60;search_cursor&#x60; encodes all the information needed to get the next batch of &#x60;limit&#x60; data. | 
 **transactionId** | **string** | An aplhanumeric string that can be used to trace a request across services. If not specified it will be automatically generated with the prefix \&quot;gst-\&quot; | 
 **accountId** | **string** | The account ID to filter resources. | 
 **limit** | **int32** | The maximum number of hits to return. Defaults to 10. | [default to 10]
 **timeout** | **int32** | A search timeout, bounding the search request to be executed within the specified time value and bail with the hits accumulated up to that point when expired. Defaults to the system defined timeout. | [default to 0]
 **sort** | **[]string** | Comma separated properties names used for sorting | 

### Return type

[**ScanResult**](ScanResult.md)

### Authorization

[IAM](../README.md#IAM)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2Search

> SearchResults V2Search(ctx).Body(body).TransactionId(transactionId).AccountId(accountId).Limit(limit).Timeout(timeout).Offset(offset).Sort(sort).Provider(provider).Execute()

Find instances of resources (v2)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewSearchBody([]string{"Query_example"}) // SearchBody | It contains the query filters, the list of fields to be returned, and the search token (initally set to null or undefined)
    transactionId := "transactionId_example" // string | An aplhanumeric string that can be used to trace a request across services. If not specified it will be automatically generated with the prefix \"gst-\" (optional)
    accountId := "accountId_example" // string | The account ID to filter resources. (optional)
    limit := int32(56) // int32 | The maximum number of hits to return. Defaults to 10. (optional) (default to 10)
    timeout := int32(56) // int32 | A search timeout, bounding the search request to be executed within the specified time value and bail with the hits accumulated up to that point when expired. Defaults to the system defined timeout. (optional) (default to 0)
    offset := int32(56) // int32 | The offset is the index of the item you want to start returning data from. Default is 0. (optional) (default to 0)
    sort := "sort_example" // string | Comma separated properties name used for sorting (optional)
    provider := "provider_example" // string | Search providers. Supported values are `ghost` and `ims`. If not specified, the defaults to `ghost`. Specify a single provider. (optional) (default to "ghost")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SearchApi.V2Search(context.Background()).Body(body).TransactionId(transactionId).AccountId(accountId).Limit(limit).Timeout(timeout).Offset(offset).Sort(sort).Provider(provider).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.V2Search``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2Search`: SearchResults
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.V2Search`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2SearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SearchBody**](SearchBody.md) | It contains the query filters, the list of fields to be returned, and the search token (initally set to null or undefined) | 
 **transactionId** | **string** | An aplhanumeric string that can be used to trace a request across services. If not specified it will be automatically generated with the prefix \&quot;gst-\&quot; | 
 **accountId** | **string** | The account ID to filter resources. | 
 **limit** | **int32** | The maximum number of hits to return. Defaults to 10. | [default to 10]
 **timeout** | **int32** | A search timeout, bounding the search request to be executed within the specified time value and bail with the hits accumulated up to that point when expired. Defaults to the system defined timeout. | [default to 0]
 **offset** | **int32** | The offset is the index of the item you want to start returning data from. Default is 0. | [default to 0]
 **sort** | **string** | Comma separated properties name used for sorting | 
 **provider** | **string** | Search providers. Supported values are &#x60;ghost&#x60; and &#x60;ims&#x60;. If not specified, the defaults to &#x60;ghost&#x60;. Specify a single provider. | [default to &quot;ghost&quot;]

### Return type

[**SearchResults**](SearchResults.md)

### Authorization

[IAM](../README.md#IAM)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

