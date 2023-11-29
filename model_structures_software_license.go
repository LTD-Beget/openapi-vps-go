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

// StructuresSoftwareLicense struct for StructuresSoftwareLicense
type StructuresSoftwareLicense struct {
	Id *int32 `json:"id,omitempty"`
	Type *string `json:"type,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	Plan *string `json:"plan,omitempty"`
	AdditionalInfo *string `json:"additional_info,omitempty"`
	AdditionalInfoEn *string `json:"additional_info_en,omitempty"`
	Primary *bool `json:"primary,omitempty"`
	Billing *StructuresSoftwareLicenseBillingType `json:"billing,omitempty"`
}

// NewStructuresSoftwareLicense instantiates a new StructuresSoftwareLicense object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStructuresSoftwareLicense() *StructuresSoftwareLicense {
	this := StructuresSoftwareLicense{}
	return &this
}

// NewStructuresSoftwareLicenseWithDefaults instantiates a new StructuresSoftwareLicense object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStructuresSoftwareLicenseWithDefaults() *StructuresSoftwareLicense {
	this := StructuresSoftwareLicense{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *StructuresSoftwareLicense) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresSoftwareLicense) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StructuresSoftwareLicense) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *StructuresSoftwareLicense) SetId(v int32) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *StructuresSoftwareLicense) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresSoftwareLicense) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *StructuresSoftwareLicense) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *StructuresSoftwareLicense) SetType(v string) {
	o.Type = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *StructuresSoftwareLicense) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresSoftwareLicense) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
    return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *StructuresSoftwareLicense) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *StructuresSoftwareLicense) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *StructuresSoftwareLicense) GetPlan() string {
	if o == nil || isNil(o.Plan) {
		var ret string
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresSoftwareLicense) GetPlanOk() (*string, bool) {
	if o == nil || isNil(o.Plan) {
    return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *StructuresSoftwareLicense) HasPlan() bool {
	if o != nil && !isNil(o.Plan) {
		return true
	}

	return false
}

// SetPlan gets a reference to the given string and assigns it to the Plan field.
func (o *StructuresSoftwareLicense) SetPlan(v string) {
	o.Plan = &v
}

// GetAdditionalInfo returns the AdditionalInfo field value if set, zero value otherwise.
func (o *StructuresSoftwareLicense) GetAdditionalInfo() string {
	if o == nil || isNil(o.AdditionalInfo) {
		var ret string
		return ret
	}
	return *o.AdditionalInfo
}

// GetAdditionalInfoOk returns a tuple with the AdditionalInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresSoftwareLicense) GetAdditionalInfoOk() (*string, bool) {
	if o == nil || isNil(o.AdditionalInfo) {
    return nil, false
	}
	return o.AdditionalInfo, true
}

// HasAdditionalInfo returns a boolean if a field has been set.
func (o *StructuresSoftwareLicense) HasAdditionalInfo() bool {
	if o != nil && !isNil(o.AdditionalInfo) {
		return true
	}

	return false
}

// SetAdditionalInfo gets a reference to the given string and assigns it to the AdditionalInfo field.
func (o *StructuresSoftwareLicense) SetAdditionalInfo(v string) {
	o.AdditionalInfo = &v
}

// GetAdditionalInfoEn returns the AdditionalInfoEn field value if set, zero value otherwise.
func (o *StructuresSoftwareLicense) GetAdditionalInfoEn() string {
	if o == nil || isNil(o.AdditionalInfoEn) {
		var ret string
		return ret
	}
	return *o.AdditionalInfoEn
}

// GetAdditionalInfoEnOk returns a tuple with the AdditionalInfoEn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresSoftwareLicense) GetAdditionalInfoEnOk() (*string, bool) {
	if o == nil || isNil(o.AdditionalInfoEn) {
    return nil, false
	}
	return o.AdditionalInfoEn, true
}

// HasAdditionalInfoEn returns a boolean if a field has been set.
func (o *StructuresSoftwareLicense) HasAdditionalInfoEn() bool {
	if o != nil && !isNil(o.AdditionalInfoEn) {
		return true
	}

	return false
}

// SetAdditionalInfoEn gets a reference to the given string and assigns it to the AdditionalInfoEn field.
func (o *StructuresSoftwareLicense) SetAdditionalInfoEn(v string) {
	o.AdditionalInfoEn = &v
}

// GetPrimary returns the Primary field value if set, zero value otherwise.
func (o *StructuresSoftwareLicense) GetPrimary() bool {
	if o == nil || isNil(o.Primary) {
		var ret bool
		return ret
	}
	return *o.Primary
}

// GetPrimaryOk returns a tuple with the Primary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresSoftwareLicense) GetPrimaryOk() (*bool, bool) {
	if o == nil || isNil(o.Primary) {
    return nil, false
	}
	return o.Primary, true
}

// HasPrimary returns a boolean if a field has been set.
func (o *StructuresSoftwareLicense) HasPrimary() bool {
	if o != nil && !isNil(o.Primary) {
		return true
	}

	return false
}

// SetPrimary gets a reference to the given bool and assigns it to the Primary field.
func (o *StructuresSoftwareLicense) SetPrimary(v bool) {
	o.Primary = &v
}

// GetBilling returns the Billing field value if set, zero value otherwise.
func (o *StructuresSoftwareLicense) GetBilling() StructuresSoftwareLicenseBillingType {
	if o == nil || isNil(o.Billing) {
		var ret StructuresSoftwareLicenseBillingType
		return ret
	}
	return *o.Billing
}

// GetBillingOk returns a tuple with the Billing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresSoftwareLicense) GetBillingOk() (*StructuresSoftwareLicenseBillingType, bool) {
	if o == nil || isNil(o.Billing) {
    return nil, false
	}
	return o.Billing, true
}

// HasBilling returns a boolean if a field has been set.
func (o *StructuresSoftwareLicense) HasBilling() bool {
	if o != nil && !isNil(o.Billing) {
		return true
	}

	return false
}

// SetBilling gets a reference to the given StructuresSoftwareLicenseBillingType and assigns it to the Billing field.
func (o *StructuresSoftwareLicense) SetBilling(v StructuresSoftwareLicenseBillingType) {
	o.Billing = &v
}

func (o StructuresSoftwareLicense) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !isNil(o.Plan) {
		toSerialize["plan"] = o.Plan
	}
	if !isNil(o.AdditionalInfo) {
		toSerialize["additional_info"] = o.AdditionalInfo
	}
	if !isNil(o.AdditionalInfoEn) {
		toSerialize["additional_info_en"] = o.AdditionalInfoEn
	}
	if !isNil(o.Primary) {
		toSerialize["primary"] = o.Primary
	}
	if !isNil(o.Billing) {
		toSerialize["billing"] = o.Billing
	}
	return json.Marshal(toSerialize)
}

type NullableStructuresSoftwareLicense struct {
	value *StructuresSoftwareLicense
	isSet bool
}

func (v NullableStructuresSoftwareLicense) Get() *StructuresSoftwareLicense {
	return v.value
}

func (v *NullableStructuresSoftwareLicense) Set(val *StructuresSoftwareLicense) {
	v.value = val
	v.isSet = true
}

func (v NullableStructuresSoftwareLicense) IsSet() bool {
	return v.isSet
}

func (v *NullableStructuresSoftwareLicense) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStructuresSoftwareLicense(val *StructuresSoftwareLicense) *NullableStructuresSoftwareLicense {
	return &NullableStructuresSoftwareLicense{value: val, isSet: true}
}

func (v NullableStructuresSoftwareLicense) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStructuresSoftwareLicense) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


