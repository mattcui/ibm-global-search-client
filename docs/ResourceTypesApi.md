# \ResourceTypesApi

All URIs are relative to *https://api.global-search-tagging.cloud.ibm.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSupportedTypes**](ResourceTypesApi.md#GetSupportedTypes) | **Get** /v2/resources/supported_types | DEPRECATED. Get all GhoST indices.



## GetSupportedTypes

> SupportedTypesList GetSupportedTypes(ctx).Execute()

DEPRECATED. Get all GhoST indices.



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ResourceTypesApi.GetSupportedTypes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceTypesApi.GetSupportedTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSupportedTypes`: SupportedTypesList
    fmt.Fprintf(os.Stdout, "Response from `ResourceTypesApi.GetSupportedTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportedTypesRequest struct via the builder pattern


### Return type

[**SupportedTypesList**](SupportedTypesList.md)

### Authorization

[IAM](../README.md#IAM)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

