# Error

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | A code that identifies the error. | [optional] 
**Description** | Pointer to **string** | A description of the solution to the error. | [optional] 
**Level** | Pointer to **string** | The severity level of the error. | [optional] 
**Message** | Pointer to **string** | A description of the error detected in the request. | [optional] 
**MoreInfo** | Pointer to **string** | The location of additional information about the error. | [optional] 
**Response** | Pointer to **string** | The response received for the request. | [optional] 

## Methods

### NewError

`func NewError() *Error`

NewError instantiates a new Error object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorWithDefaults

`func NewErrorWithDefaults() *Error`

NewErrorWithDefaults instantiates a new Error object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *Error) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Error) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Error) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Error) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDescription

`func (o *Error) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Error) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Error) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Error) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLevel

`func (o *Error) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *Error) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *Error) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *Error) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetMessage

`func (o *Error) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Error) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Error) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Error) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMoreInfo

`func (o *Error) GetMoreInfo() string`

GetMoreInfo returns the MoreInfo field if non-nil, zero value otherwise.

### GetMoreInfoOk

`func (o *Error) GetMoreInfoOk() (*string, bool)`

GetMoreInfoOk returns a tuple with the MoreInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoreInfo

`func (o *Error) SetMoreInfo(v string)`

SetMoreInfo sets MoreInfo field to given value.

### HasMoreInfo

`func (o *Error) HasMoreInfo() bool`

HasMoreInfo returns a boolean if a field has been set.

### GetResponse

`func (o *Error) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *Error) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *Error) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *Error) HasResponse() bool`

HasResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


