# {{classname}}

All URIs are relative to *https://api.global-search-tagging.cloud.ibm.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Search**](SearchApi.md#Search) | **Post** /v3/resources/search | Find instances of resources (v3)
[**V2Search**](SearchApi.md#V2Search) | **Post** /v2/resources/search | Find instances of resources (v2)

# **Search**
> ScanResult Search(ctx, body, optional)
Find instances of resources (v3)

Find Cloud Foundry resources, IAM-enabled resources, or  storage and network resources running on classic infrastructure in a  specific account ID. You can apply query strings if necessary.   To filter results, you can insert a string using the Lucene syntax and the  query string is parsed into a series of terms and operators. A term can be  a single word or a phrase, in which case the search is performed for all  the words, in the same order. To filter for a specific value regardless of  the property that contains it, type the search term without specifying a  field. Only resources that belong to the account ID and that are accessible  by the client are returned.   You must use `/v3/resources/search` when you need to fetch more than `10000`  resource items. The `/v2/resources/search` prohibits paginating through such  a big number. On the first call, the operation returns a live cursor on the  data that you must use on all the subsequent calls to get the next batch of  results until you get the empty result set. By default, the fields returned  for every resource are \"crn\", \"name\", \"family\", \"type\", and \"account_id\". You  can specify the subset of the fields you want in your request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ScanBody**](ScanBody.md)| It contains the query filters on the first operation call, or the search_cursor on next calls. On subsequent calls, set the &#x60;search_cursor&#x60; to the value returned by the previous call. After the first, you must set only the &#x60;search_cursor&#x60;. Any other parameter but the &#x60;search_cursor&#x60; will be ignored. The &#x60;search_cursor&#x60; encodes all the information needed to get the next batch of &#x60;limit&#x60; data. | 
 **optional** | ***SearchApiSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiSearchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionId** | **optional.**| An aplhanumeric string that can be used to trace a request across services. If not specified it will be automatically generated with the prefix \&quot;gst-\&quot; | 
 **accountId** | **optional.**| The account ID to filter resources. | 
 **limit** | **optional.**| The maximum number of hits to return. Defaults to 10. | [default to 10]
 **timeout** | **optional.**| A search timeout, bounding the search request to be executed within the specified time value and bail with the hits accumulated up to that point when expired. Defaults to the system defined timeout. | [default to 0]
 **sort** | [**optional.Interface of []string**](string.md)| Comma separated properties names used for sorting | 

### Return type

[**ScanResult**](ScanResult.md)

### Authorization

[IAM](../README.md#IAM)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V2Search**
> SearchResults V2Search(ctx, body, optional)
Find instances of resources (v2)

Find resources in the internal search platform for a specific account ID. You can optionally apply query strings. To filter results, you can insert a string using the Lucene syntax and the query string is parsed into a series of terms and operators. A term can be a single word or a phrase, in which case the search is performed for all the words, in the same order. To filter for a specific value regardless of the property that contains it, use an asterisk as the key name. Results are filtered by the account ID and IAM subsystem. Only resources that belong to the account ID and that are accessible by the client are returned. The usage of this API is deprecated for cases when the provider is not specified or  `provider=ghost`.  `/v3/resources/search` replaces this API for the deprecated cases.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SearchBody**](SearchBody.md)| It contains the query filters, the list of fields to be returned, and the search token (initally set to null or undefined) | 
 **optional** | ***SearchApiV2SearchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiV2SearchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionId** | **optional.**| An aplhanumeric string that can be used to trace a request across services. If not specified it will be automatically generated with the prefix \&quot;gst-\&quot; | 
 **accountId** | **optional.**| The account ID to filter resources. | 
 **limit** | **optional.**| The maximum number of hits to return. Defaults to 10. | [default to 10]
 **timeout** | **optional.**| A search timeout, bounding the search request to be executed within the specified time value and bail with the hits accumulated up to that point when expired. Defaults to the system defined timeout. | [default to 0]
 **offset** | **optional.**| The offset is the index of the item you want to start returning data from. Default is 0. | [default to 0]
 **sort** | **optional.**| Comma separated properties name used for sorting | 
 **provider** | **optional.**| Search providers. Supported values are &#x60;ghost&#x60; and &#x60;ims&#x60;. If not specified, the defaults to &#x60;ghost&#x60;. Specify a single provider. | [default to ghost]

### Return type

[**SearchResults**](SearchResults.md)

### Authorization

[IAM](../README.md#IAM)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

