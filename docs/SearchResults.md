# SearchResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]ResultItem**](ResultItem.md) | The array of results. Each item will represent a resource and will contains all visible properties associated with it. | 
**MoreData** | **bool** | If false, there is no more data to retrieve on the next page. If true, it&#39;s possible that there is more data to retrieve on the next page. | 
**Token** | Pointer to **string** | The search token to use on the next call. | [optional] 
**FilterError** | **bool** | It is set to true if the result is partial of an IAM error when validating user authorization on one or more resources. This field is DEPRECATED and will be removed in future versions of this API | 
**PartialData** | **int32** | Indicates if the result that is set might be partial or not. Value 0 means the result set is complete. A value greater than 0 means the result set might be incomplete. Its single bits identify the cause. The first bit means the error is in the IAM filter. The second bit means errors are in elasticsearch shards. | 
**Offset** | **int32** | Offset parameter specified by the user | 
**Limit** | **int32** | Limit parameter specified by the user | 

## Methods

### NewSearchResults

`func NewSearchResults(items []ResultItem, moreData bool, filterError bool, partialData int32, offset int32, limit int32, ) *SearchResults`

NewSearchResults instantiates a new SearchResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResultsWithDefaults

`func NewSearchResultsWithDefaults() *SearchResults`

NewSearchResultsWithDefaults instantiates a new SearchResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *SearchResults) GetItems() []ResultItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SearchResults) GetItemsOk() (*[]ResultItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SearchResults) SetItems(v []ResultItem)`

SetItems sets Items field to given value.


### GetMoreData

`func (o *SearchResults) GetMoreData() bool`

GetMoreData returns the MoreData field if non-nil, zero value otherwise.

### GetMoreDataOk

`func (o *SearchResults) GetMoreDataOk() (*bool, bool)`

GetMoreDataOk returns a tuple with the MoreData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoreData

`func (o *SearchResults) SetMoreData(v bool)`

SetMoreData sets MoreData field to given value.


### GetToken

`func (o *SearchResults) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *SearchResults) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *SearchResults) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *SearchResults) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetFilterError

`func (o *SearchResults) GetFilterError() bool`

GetFilterError returns the FilterError field if non-nil, zero value otherwise.

### GetFilterErrorOk

`func (o *SearchResults) GetFilterErrorOk() (*bool, bool)`

GetFilterErrorOk returns a tuple with the FilterError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterError

`func (o *SearchResults) SetFilterError(v bool)`

SetFilterError sets FilterError field to given value.


### GetPartialData

`func (o *SearchResults) GetPartialData() int32`

GetPartialData returns the PartialData field if non-nil, zero value otherwise.

### GetPartialDataOk

`func (o *SearchResults) GetPartialDataOk() (*int32, bool)`

GetPartialDataOk returns a tuple with the PartialData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialData

`func (o *SearchResults) SetPartialData(v int32)`

SetPartialData sets PartialData field to given value.


### GetOffset

`func (o *SearchResults) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *SearchResults) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *SearchResults) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetLimit

`func (o *SearchResults) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SearchResults) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SearchResults) SetLimit(v int32)`

SetLimit sets Limit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


