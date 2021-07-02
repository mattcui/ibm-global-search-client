# SearchResults

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]ModelMap**](map.md) | The array of results. Each item will represent a resource and will contains all visible properties associated with it. | [default to null]
**MoreData** | **bool** | If false, there is no more data to retrieve on the next page. If true, it&#x27;s possible that there is more data to retrieve on the next page. | [default to null]
**Token** | **string** | The search token to use on the next call. | [optional] [default to null]
**FilterError** | **bool** | It is set to true if the result is partial of an IAM error when validating user authorization on one or more resources. This field is DEPRECATED and will be removed in future versions of this API | [default to null]
**PartialData** | **int32** | Indicates if the result that is set might be partial or not. Value 0 means the result set is complete. A value greater than 0 means the result set might be incomplete. Its single bits identify the cause. The first bit means the error is in the IAM filter. The second bit means errors are in elasticsearch shards. | [default to null]
**Offset** | **int32** | Offset parameter specified by the user | [default to null]
**Limit** | **int32** | Limit parameter specified by the user | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

