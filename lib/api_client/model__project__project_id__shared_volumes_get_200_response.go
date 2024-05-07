/*
Podinate API

The API for the simple containerisation solution Podinate.

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_client

import (
	"encoding/json"
)

// checks if the ProjectProjectIdSharedVolumesGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectProjectIdSharedVolumesGet200Response{}

// ProjectProjectIdSharedVolumesGet200Response struct for ProjectProjectIdSharedVolumesGet200Response
type ProjectProjectIdSharedVolumesGet200Response struct {
	Items []SharedVolume `json:"items,omitempty"`
	// The total number of shared volumes
	Total int32 `json:"total"`
	// The current page number
	Page int32 `json:"page"`
	// The number of items per page
	Limit int32 `json:"limit"`
}

// NewProjectProjectIdSharedVolumesGet200Response instantiates a new ProjectProjectIdSharedVolumesGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectProjectIdSharedVolumesGet200Response(total int32, page int32, limit int32) *ProjectProjectIdSharedVolumesGet200Response {
	this := ProjectProjectIdSharedVolumesGet200Response{}
	this.Total = total
	this.Page = page
	this.Limit = limit
	return &this
}

// NewProjectProjectIdSharedVolumesGet200ResponseWithDefaults instantiates a new ProjectProjectIdSharedVolumesGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectProjectIdSharedVolumesGet200ResponseWithDefaults() *ProjectProjectIdSharedVolumesGet200Response {
	this := ProjectProjectIdSharedVolumesGet200Response{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ProjectProjectIdSharedVolumesGet200Response) GetItems() []SharedVolume {
	if o == nil || IsNil(o.Items) {
		var ret []SharedVolume
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectProjectIdSharedVolumesGet200Response) GetItemsOk() ([]SharedVolume, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ProjectProjectIdSharedVolumesGet200Response) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []SharedVolume and assigns it to the Items field.
func (o *ProjectProjectIdSharedVolumesGet200Response) SetItems(v []SharedVolume) {
	o.Items = v
}

// GetTotal returns the Total field value
func (o *ProjectProjectIdSharedVolumesGet200Response) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *ProjectProjectIdSharedVolumesGet200Response) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *ProjectProjectIdSharedVolumesGet200Response) SetTotal(v int32) {
	o.Total = v
}

// GetPage returns the Page field value
func (o *ProjectProjectIdSharedVolumesGet200Response) GetPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *ProjectProjectIdSharedVolumesGet200Response) GetPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *ProjectProjectIdSharedVolumesGet200Response) SetPage(v int32) {
	o.Page = v
}

// GetLimit returns the Limit field value
func (o *ProjectProjectIdSharedVolumesGet200Response) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *ProjectProjectIdSharedVolumesGet200Response) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *ProjectProjectIdSharedVolumesGet200Response) SetLimit(v int32) {
	o.Limit = v
}

func (o ProjectProjectIdSharedVolumesGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectProjectIdSharedVolumesGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	toSerialize["total"] = o.Total
	toSerialize["page"] = o.Page
	toSerialize["limit"] = o.Limit
	return toSerialize, nil
}

type NullableProjectProjectIdSharedVolumesGet200Response struct {
	value *ProjectProjectIdSharedVolumesGet200Response
	isSet bool
}

func (v NullableProjectProjectIdSharedVolumesGet200Response) Get() *ProjectProjectIdSharedVolumesGet200Response {
	return v.value
}

func (v *NullableProjectProjectIdSharedVolumesGet200Response) Set(val *ProjectProjectIdSharedVolumesGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectProjectIdSharedVolumesGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectProjectIdSharedVolumesGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectProjectIdSharedVolumesGet200Response(val *ProjectProjectIdSharedVolumesGet200Response) *NullableProjectProjectIdSharedVolumesGet200Response {
	return &NullableProjectProjectIdSharedVolumesGet200Response{value: val, isSet: true}
}

func (v NullableProjectProjectIdSharedVolumesGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectProjectIdSharedVolumesGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}