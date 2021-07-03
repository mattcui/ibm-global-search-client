# ScanBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | Pointer to **string** | The Lucene-formatted query string. Default to &#39;*&#39; if not set | [optional] 
**Fields** | Pointer to **[]string** | The list of the fields returned by the search. Defaults to all. &#x60;crn&#x60; is always returned. | [optional] 
**SearchCursor** | Pointer to **string** | An opaque search cursor that is returned on each operation call and that must be set on next call. | [optional] 

## Methods

### NewScanBody

`func NewScanBody() *ScanBody`

NewScanBody instantiates a new ScanBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScanBodyWithDefaults

`func NewScanBodyWithDefaults() *ScanBody`

NewScanBodyWithDefaults instantiates a new ScanBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *ScanBody) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ScanBody) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ScanBody) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *ScanBody) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetFields

`func (o *ScanBody) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ScanBody) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ScanBody) SetFields(v []string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *ScanBody) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetSearchCursor

`func (o *ScanBody) GetSearchCursor() string`

GetSearchCursor returns the SearchCursor field if non-nil, zero value otherwise.

### GetSearchCursorOk

`func (o *ScanBody) GetSearchCursorOk() (*string, bool)`

GetSearchCursorOk returns a tuple with the SearchCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchCursor

`func (o *ScanBody) SetSearchCursor(v string)`

SetSearchCursor sets SearchCursor field to given value.

### HasSearchCursor

`func (o *ScanBody) HasSearchCursor() bool`

HasSearchCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


