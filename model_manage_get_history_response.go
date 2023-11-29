/*
API Облачных серверов

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package begetOpenapiVps

import (
	"encoding/json"
)

// ManageGetHistoryResponse struct for ManageGetHistoryResponse
type ManageGetHistoryResponse struct {
	History []ManageHistoryItem `json:"history,omitempty"`
}

// NewManageGetHistoryResponse instantiates a new ManageGetHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageGetHistoryResponse() *ManageGetHistoryResponse {
	this := ManageGetHistoryResponse{}
	return &this
}

// NewManageGetHistoryResponseWithDefaults instantiates a new ManageGetHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageGetHistoryResponseWithDefaults() *ManageGetHistoryResponse {
	this := ManageGetHistoryResponse{}
	return &this
}

// GetHistory returns the History field value if set, zero value otherwise.
func (o *ManageGetHistoryResponse) GetHistory() []ManageHistoryItem {
	if o == nil || isNil(o.History) {
		var ret []ManageHistoryItem
		return ret
	}
	return o.History
}

// GetHistoryOk returns a tuple with the History field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageGetHistoryResponse) GetHistoryOk() ([]ManageHistoryItem, bool) {
	if o == nil || isNil(o.History) {
    return nil, false
	}
	return o.History, true
}

// HasHistory returns a boolean if a field has been set.
func (o *ManageGetHistoryResponse) HasHistory() bool {
	if o != nil && !isNil(o.History) {
		return true
	}

	return false
}

// SetHistory gets a reference to the given []ManageHistoryItem and assigns it to the History field.
func (o *ManageGetHistoryResponse) SetHistory(v []ManageHistoryItem) {
	o.History = v
}

func (o ManageGetHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.History) {
		toSerialize["history"] = o.History
	}
	return json.Marshal(toSerialize)
}

type NullableManageGetHistoryResponse struct {
	value *ManageGetHistoryResponse
	isSet bool
}

func (v NullableManageGetHistoryResponse) Get() *ManageGetHistoryResponse {
	return v.value
}

func (v *NullableManageGetHistoryResponse) Set(val *ManageGetHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableManageGetHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableManageGetHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageGetHistoryResponse(val *ManageGetHistoryResponse) *NullableManageGetHistoryResponse {
	return &NullableManageGetHistoryResponse{value: val, isSet: true}
}

func (v NullableManageGetHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageGetHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


