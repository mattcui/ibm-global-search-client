# SearchBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **[]string** | The query string | 
**Fields** | Pointer to **[]string** | The list of the fields that are returned by the search. The CRN is always returned. | [optional] 
**Token** | Pointer to **string** | An opaque token. Initially set to null or undefined, then pass the value returned by the previous call | [optional] 

## Methods

### NewSearchBody

`func NewSearchBody(query []string, ) *SearchBody`

NewSearchBody instantiates a new SearchBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchBodyWithDefaults

`func NewSearchBodyWithDefaults() *SearchBody`

NewSearchBodyWithDefaults instantiates a new SearchBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *SearchBody) GetQuery() []string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SearchBody) GetQueryOk() (*[]string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SearchBody) SetQuery(v []string)`

SetQuery sets Query field to given value.


### GetFields

`func (o *SearchBody) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *SearchBody) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *SearchBody) SetFields(v []string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *SearchBody) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetToken

`func (o *SearchBody) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *SearchBody) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *SearchBody) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *SearchBody) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


