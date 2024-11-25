/*
API Облачных серверов

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.6.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package begetOpenapiVps

import (
	"encoding/json"
)

// MarketplaceSoftwareInfo struct for MarketplaceSoftwareInfo
type MarketplaceSoftwareInfo struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	Version *string `json:"version,omitempty"`
	Description *string `json:"description,omitempty"`
	DescriptionEn *string `json:"description_en,omitempty"`
	Metadata *StructuresSoftwareMetadata `json:"metadata,omitempty"`
	FieldData []MarketplaceFieldDesc `json:"field_data,omitempty"`
	Requirements *MarketplaceSoftwareInfoRequirements `json:"requirements,omitempty"`
	Category []StructuresSoftwareCategory `json:"category,omitempty"`
	Slug *string `json:"slug,omitempty"`
	DocumentationSlug *string `json:"documentation_slug,omitempty"`
	UnattendedInstallAvailable *bool `json:"unattended_install_available,omitempty"`
	License []StructuresSoftwareLicense `json:"license,omitempty"`
}

// NewMarketplaceSoftwareInfo instantiates a new MarketplaceSoftwareInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketplaceSoftwareInfo() *MarketplaceSoftwareInfo {
	this := MarketplaceSoftwareInfo{}
	return &this
}

// NewMarketplaceSoftwareInfoWithDefaults instantiates a new MarketplaceSoftwareInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketplaceSoftwareInfoWithDefaults() *MarketplaceSoftwareInfo {
	this := MarketplaceSoftwareInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MarketplaceSoftwareInfo) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceSoftwareInfo) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MarketplaceSoftwareInfo) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *MarketplaceSoftwareInfo) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MarketplaceSoftwareInfo) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceSoftwareInfo) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MarketplaceSoftwareInfo) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MarketplaceSoftwareInfo) SetName(v string) {
	o.Name = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *MarketplaceSoftwareInfo) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceSoftwareInfo) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
    return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MarketplaceSoftwareInfo) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *MarketplaceSoftwareInfo) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *MarketplaceSoftwareInfo) GetVersion() string {
	if o == nil || isNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceSoftwareInfo) GetVersionOk() (*string, bool) {
	if o == nil || isNil(o.Version) {
    return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *MarketplaceSoftwareInfo) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *MarketplaceSoftwareInfo) SetVersion(v string) {
	o.Version = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MarketplaceSoftwareInfo) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceSoftwareInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MarketplaceSoftwareInfo) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MarketplaceSoftwareInfo) SetDescription(v string) {
	o.Description = &v
}

// GetDescriptionEn returns the DescriptionEn field value if set, zero value otherwise.
func (o *MarketplaceSoftwareInfo) GetDescriptionEn() string {
	if o == nil || isNil(o.DescriptionEn) {
		var ret string
		return ret
	}
	return *o.DescriptionEn
}

// GetDescriptionEnOk returns a tuple with the DescriptionEn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceSoftwareInfo) GetDescriptionEnOk() (*string, bool) {
	if o == nil || isNil(o.DescriptionEn) {
    return nil, false
	}
	return o.DescriptionEn, true
}

// HasDescriptionEn returns a boolean if a field has been set.
func (o *MarketplaceSoftwareInfo) HasDescriptionEn() bool {
	if o != nil && !isNil(o.DescriptionEn) {
		return true
	}

	return false
}

// SetDescriptionEn gets a reference to the given string and assigns it to the DescriptionEn field.
func (o *MarketplaceSoftwareInfo) SetDescriptionEn(v string) {
	o.DescriptionEn = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *MarketplaceSoftwareInfo) GetMetadata() StructuresSoftwareMetadata {
	if o == nil || isNil(o.Metadata) {
		var ret StructuresSoftwareMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceSoftwareInfo) GetMetadataOk() (*StructuresSoftwareMetadata, bool) {
	if o == nil || isNil(o.Metadata) {
    return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *MarketplaceSoftwareInfo) HasMetadata() bool {
	if o != nil && !isNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given StructuresSoftwareMetadata and assigns it to the Metadata field.
func (o *MarketplaceSoftwareInfo) SetMetadata(v StructuresSoftwareMetadata) {
	o.Metadata = &v
}

// GetFieldData returns the FieldData field value if set, zero value otherwise.
func (o *MarketplaceSoftwareInfo) GetFieldData() []MarketplaceFieldDesc {
	if o == nil || isNil(o.FieldData) {
		var ret []MarketplaceFieldDesc
		return ret
	}
	return o.FieldData
}

// GetFieldDataOk returns a tuple with the FieldData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceSoftwareInfo) GetFieldDataOk() ([]MarketplaceFieldDesc, bool) {
	if o == nil || isNil(o.FieldData) {
    return nil, false
	}
	return o.FieldData, true
}

// HasFieldData returns a boolean if a field has been set.
func (o *MarketplaceSoftwareInfo) HasFieldData() bool {
	if o != nil && !isNil(o.FieldData) {
		return true
	}

	return false
}

// SetFieldData gets a reference to the given []MarketplaceFieldDesc and assigns it to the FieldData field.
func (o *MarketplaceSoftwareInfo) SetFieldData(v []MarketplaceFieldDesc) {
	o.FieldData = v
}

// GetRequirements returns the Requirements field value if set, zero value otherwise.
func (o *MarketplaceSoftwareInfo) GetRequirements() MarketplaceSoftwareInfoRequirements {
	if o == nil || isNil(o.Requirements) {
		var ret MarketplaceSoftwareInfoRequirements
		return ret
	}
	return *o.Requirements
}

// GetRequirementsOk returns a tuple with the Requirements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceSoftwareInfo) GetRequirementsOk() (*MarketplaceSoftwareInfoRequirements, bool) {
	if o == nil || isNil(o.Requirements) {
    return nil, false
	}
	return o.Requirements, true
}

// HasRequirements returns a boolean if a field has been set.
func (o *MarketplaceSoftwareInfo) HasRequirements() bool {
	if o != nil && !isNil(o.Requirements) {
		return true
	}

	return false
}

// SetRequirements gets a reference to the given MarketplaceSoftwareInfoRequirements and assigns it to the Requirements field.
func (o *MarketplaceSoftwareInfo) SetRequirements(v MarketplaceSoftwareInfoRequirements) {
	o.Requirements = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *MarketplaceSoftwareInfo) GetCategory() []StructuresSoftwareCategory {
	if o == nil || isNil(o.Category) {
		var ret []StructuresSoftwareCategory
		return ret
	}
	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceSoftwareInfo) GetCategoryOk() ([]StructuresSoftwareCategory, bool) {
	if o == nil || isNil(o.Category) {
    return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *MarketplaceSoftwareInfo) HasCategory() bool {
	if o != nil && !isNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given []StructuresSoftwareCategory and assigns it to the Category field.
func (o *MarketplaceSoftwareInfo) SetCategory(v []StructuresSoftwareCategory) {
	o.Category = v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *MarketplaceSoftwareInfo) GetSlug() string {
	if o == nil || isNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceSoftwareInfo) GetSlugOk() (*string, bool) {
	if o == nil || isNil(o.Slug) {
    return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *MarketplaceSoftwareInfo) HasSlug() bool {
	if o != nil && !isNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *MarketplaceSoftwareInfo) SetSlug(v string) {
	o.Slug = &v
}

// GetDocumentationSlug returns the DocumentationSlug field value if set, zero value otherwise.
func (o *MarketplaceSoftwareInfo) GetDocumentationSlug() string {
	if o == nil || isNil(o.DocumentationSlug) {
		var ret string
		return ret
	}
	return *o.DocumentationSlug
}

// GetDocumentationSlugOk returns a tuple with the DocumentationSlug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceSoftwareInfo) GetDocumentationSlugOk() (*string, bool) {
	if o == nil || isNil(o.DocumentationSlug) {
    return nil, false
	}
	return o.DocumentationSlug, true
}

// HasDocumentationSlug returns a boolean if a field has been set.
func (o *MarketplaceSoftwareInfo) HasDocumentationSlug() bool {
	if o != nil && !isNil(o.DocumentationSlug) {
		return true
	}

	return false
}

// SetDocumentationSlug gets a reference to the given string and assigns it to the DocumentationSlug field.
func (o *MarketplaceSoftwareInfo) SetDocumentationSlug(v string) {
	o.DocumentationSlug = &v
}

// GetUnattendedInstallAvailable returns the UnattendedInstallAvailable field value if set, zero value otherwise.
func (o *MarketplaceSoftwareInfo) GetUnattendedInstallAvailable() bool {
	if o == nil || isNil(o.UnattendedInstallAvailable) {
		var ret bool
		return ret
	}
	return *o.UnattendedInstallAvailable
}

// GetUnattendedInstallAvailableOk returns a tuple with the UnattendedInstallAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceSoftwareInfo) GetUnattendedInstallAvailableOk() (*bool, bool) {
	if o == nil || isNil(o.UnattendedInstallAvailable) {
    return nil, false
	}
	return o.UnattendedInstallAvailable, true
}

// HasUnattendedInstallAvailable returns a boolean if a field has been set.
func (o *MarketplaceSoftwareInfo) HasUnattendedInstallAvailable() bool {
	if o != nil && !isNil(o.UnattendedInstallAvailable) {
		return true
	}

	return false
}

// SetUnattendedInstallAvailable gets a reference to the given bool and assigns it to the UnattendedInstallAvailable field.
func (o *MarketplaceSoftwareInfo) SetUnattendedInstallAvailable(v bool) {
	o.UnattendedInstallAvailable = &v
}

// GetLicense returns the License field value if set, zero value otherwise.
func (o *MarketplaceSoftwareInfo) GetLicense() []StructuresSoftwareLicense {
	if o == nil || isNil(o.License) {
		var ret []StructuresSoftwareLicense
		return ret
	}
	return o.License
}

// GetLicenseOk returns a tuple with the License field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceSoftwareInfo) GetLicenseOk() ([]StructuresSoftwareLicense, bool) {
	if o == nil || isNil(o.License) {
    return nil, false
	}
	return o.License, true
}

// HasLicense returns a boolean if a field has been set.
func (o *MarketplaceSoftwareInfo) HasLicense() bool {
	if o != nil && !isNil(o.License) {
		return true
	}

	return false
}

// SetLicense gets a reference to the given []StructuresSoftwareLicense and assigns it to the License field.
func (o *MarketplaceSoftwareInfo) SetLicense(v []StructuresSoftwareLicense) {
	o.License = v
}

func (o MarketplaceSoftwareInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !isNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.DescriptionEn) {
		toSerialize["description_en"] = o.DescriptionEn
	}
	if !isNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !isNil(o.FieldData) {
		toSerialize["field_data"] = o.FieldData
	}
	if !isNil(o.Requirements) {
		toSerialize["requirements"] = o.Requirements
	}
	if !isNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !isNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if !isNil(o.DocumentationSlug) {
		toSerialize["documentation_slug"] = o.DocumentationSlug
	}
	if !isNil(o.UnattendedInstallAvailable) {
		toSerialize["unattended_install_available"] = o.UnattendedInstallAvailable
	}
	if !isNil(o.License) {
		toSerialize["license"] = o.License
	}
	return json.Marshal(toSerialize)
}

type NullableMarketplaceSoftwareInfo struct {
	value *MarketplaceSoftwareInfo
	isSet bool
}

func (v NullableMarketplaceSoftwareInfo) Get() *MarketplaceSoftwareInfo {
	return v.value
}

func (v *NullableMarketplaceSoftwareInfo) Set(val *MarketplaceSoftwareInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketplaceSoftwareInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketplaceSoftwareInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketplaceSoftwareInfo(val *MarketplaceSoftwareInfo) *NullableMarketplaceSoftwareInfo {
	return &NullableMarketplaceSoftwareInfo{value: val, isSet: true}
}

func (v NullableMarketplaceSoftwareInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketplaceSoftwareInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


