/*
 * Global Search
 *
 * Search for resources with the global and shared resource properties repository integrated in the IBM Cloud platform. The search repository stores and searches cloud resources attributes, which categorize or classify resources. A resource is a physical or logical component that can be created or reserved for an application or service instance and is owned by resource providers, such as Cloud Foundry, IBM Kubernetes Service, or resource controller in IBM Cloud. Resources are uniquely identified by a Cloud Resource Name (CRN)  or by an IMS ID. The properties of a resource include tags and system properties. Both properties are defined in an IBM Cloud billing account, and span across many regions.
 *
 * API version: 2.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ScanResult The search scan response
type ScanResult struct {
	// The search cursor to use on all calls after the first one.
	SearchCursor string `json:"search_cursor"`
	// Value of the limit parameter specified by the user
	Limit *int32 `json:"limit,omitempty"`
	// The array of results. Each item represents a resource. An empty array signals the end of the result set, there are no more hits to fetch
	Items []ResultItem `json:"items"`
}

// NewScanResult instantiates a new ScanResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScanResult(searchCursor string, items []ResultItem) *ScanResult {
	this := ScanResult{}
	this.SearchCursor = searchCursor
	this.Items = items
	return &this
}

// NewScanResultWithDefaults instantiates a new ScanResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScanResultWithDefaults() *ScanResult {
	this := ScanResult{}
	return &this
}

// GetSearchCursor returns the SearchCursor field value
func (o *ScanResult) GetSearchCursor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SearchCursor
}

// GetSearchCursorOk returns a tuple with the SearchCursor field value
// and a boolean to check if the value has been set.
func (o *ScanResult) GetSearchCursorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SearchCursor, true
}

// SetSearchCursor sets field value
func (o *ScanResult) SetSearchCursor(v string) {
	o.SearchCursor = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ScanResult) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScanResult) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ScanResult) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *ScanResult) SetLimit(v int32) {
	o.Limit = &v
}

// GetItems returns the Items field value
func (o *ScanResult) GetItems() []ResultItem {
	if o == nil {
		var ret []ResultItem
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ScanResult) GetItemsOk() (*[]ResultItem, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value
func (o *ScanResult) SetItems(v []ResultItem) {
	o.Items = v
}

func (o ScanResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["search_cursor"] = o.SearchCursor
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if true {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableScanResult struct {
	value *ScanResult
	isSet bool
}

func (v NullableScanResult) Get() *ScanResult {
	return v.value
}

func (v *NullableScanResult) Set(val *ScanResult) {
	v.value = val
	v.isSet = true
}

func (v NullableScanResult) IsSet() bool {
	return v.isSet
}

func (v *NullableScanResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScanResult(val *ScanResult) *NullableScanResult {
	return &NullableScanResult{value: val, isSet: true}
}

func (v NullableScanResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScanResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


