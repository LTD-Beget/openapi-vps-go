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

// MarketplaceGetSoftwareInfoResponse struct for MarketplaceGetSoftwareInfoResponse
type MarketplaceGetSoftwareInfoResponse struct {
	Software *MarketplaceSoftwareInfo `json:"software,omitempty"`
}

// NewMarketplaceGetSoftwareInfoResponse instantiates a new MarketplaceGetSoftwareInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketplaceGetSoftwareInfoResponse() *MarketplaceGetSoftwareInfoResponse {
	this := MarketplaceGetSoftwareInfoResponse{}
	return &this
}

// NewMarketplaceGetSoftwareInfoResponseWithDefaults instantiates a new MarketplaceGetSoftwareInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketplaceGetSoftwareInfoResponseWithDefaults() *MarketplaceGetSoftwareInfoResponse {
	this := MarketplaceGetSoftwareInfoResponse{}
	return &this
}

// GetSoftware returns the Software field value if set, zero value otherwise.
func (o *MarketplaceGetSoftwareInfoResponse) GetSoftware() MarketplaceSoftwareInfo {
	if o == nil || isNil(o.Software) {
		var ret MarketplaceSoftwareInfo
		return ret
	}
	return *o.Software
}

// GetSoftwareOk returns a tuple with the Software field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceGetSoftwareInfoResponse) GetSoftwareOk() (*MarketplaceSoftwareInfo, bool) {
	if o == nil || isNil(o.Software) {
    return nil, false
	}
	return o.Software, true
}

// HasSoftware returns a boolean if a field has been set.
func (o *MarketplaceGetSoftwareInfoResponse) HasSoftware() bool {
	if o != nil && !isNil(o.Software) {
		return true
	}

	return false
}

// SetSoftware gets a reference to the given MarketplaceSoftwareInfo and assigns it to the Software field.
func (o *MarketplaceGetSoftwareInfoResponse) SetSoftware(v MarketplaceSoftwareInfo) {
	o.Software = &v
}

func (o MarketplaceGetSoftwareInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Software) {
		toSerialize["software"] = o.Software
	}
	return json.Marshal(toSerialize)
}

type NullableMarketplaceGetSoftwareInfoResponse struct {
	value *MarketplaceGetSoftwareInfoResponse
	isSet bool
}

func (v NullableMarketplaceGetSoftwareInfoResponse) Get() *MarketplaceGetSoftwareInfoResponse {
	return v.value
}

func (v *NullableMarketplaceGetSoftwareInfoResponse) Set(val *MarketplaceGetSoftwareInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketplaceGetSoftwareInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketplaceGetSoftwareInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketplaceGetSoftwareInfoResponse(val *MarketplaceGetSoftwareInfoResponse) *NullableMarketplaceGetSoftwareInfoResponse {
	return &NullableMarketplaceGetSoftwareInfoResponse{value: val, isSet: true}
}

func (v NullableMarketplaceGetSoftwareInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketplaceGetSoftwareInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


