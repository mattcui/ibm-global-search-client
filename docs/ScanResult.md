# ScanResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchCursor** | **string** | The search cursor to use on all calls after the first one. | 
**Limit** | Pointer to **int32** | Value of the limit parameter specified by the user | [optional] 
**Items** | [**[]ResultItem**](ResultItem.md) | The array of results. Each item represents a resource. An empty array signals the end of the result set, there are no more hits to fetch | 

## Methods

### NewScanResult

`func NewScanResult(searchCursor string, items []ResultItem, ) *ScanResult`

NewScanResult instantiates a new ScanResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScanResultWithDefaults

`func NewScanResultWithDefaults() *ScanResult`

NewScanResultWithDefaults instantiates a new ScanResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchCursor

`func (o *ScanResult) GetSearchCursor() string`

GetSearchCursor returns the SearchCursor field if non-nil, zero value otherwise.

### GetSearchCursorOk

`func (o *ScanResult) GetSearchCursorOk() (*string, bool)`

GetSearchCursorOk returns a tuple with the SearchCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchCursor

`func (o *ScanResult) SetSearchCursor(v string)`

SetSearchCursor sets SearchCursor field to given value.


### GetLimit

`func (o *ScanResult) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ScanResult) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ScanResult) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ScanResult) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetItems

`func (o *ScanResult) GetItems() []ResultItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ScanResult) GetItemsOk() (*[]ResultItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ScanResult) SetItems(v []ResultItem)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


