# MarketplaceSoftwareInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DescriptionEn** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to [**StructuresSoftwareMetadata**](StructuresSoftwareMetadata.md) |  | [optional] 
**FieldData** | Pointer to [**[]MarketplaceFieldDesc**](MarketplaceFieldDesc.md) |  | [optional] 
**Requirements** | Pointer to [**MarketplaceSoftwareInfoRequirements**](MarketplaceSoftwareInfoRequirements.md) |  | [optional] 
**Category** | Pointer to [**[]StructuresSoftwareCategory**](StructuresSoftwareCategory.md) |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**DocumentationSlug** | Pointer to **string** |  | [optional] 
**UnattendedInstallAvailable** | Pointer to **bool** |  | [optional] 
**License** | Pointer to [**[]StructuresSoftwareLicense**](StructuresSoftwareLicense.md) |  | [optional] 

## Methods

### NewMarketplaceSoftwareInfo

`func NewMarketplaceSoftwareInfo() *MarketplaceSoftwareInfo`

NewMarketplaceSoftwareInfo instantiates a new MarketplaceSoftwareInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceSoftwareInfoWithDefaults

`func NewMarketplaceSoftwareInfoWithDefaults() *MarketplaceSoftwareInfo`

NewMarketplaceSoftwareInfoWithDefaults instantiates a new MarketplaceSoftwareInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MarketplaceSoftwareInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketplaceSoftwareInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketplaceSoftwareInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MarketplaceSoftwareInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MarketplaceSoftwareInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MarketplaceSoftwareInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MarketplaceSoftwareInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MarketplaceSoftwareInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *MarketplaceSoftwareInfo) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MarketplaceSoftwareInfo) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MarketplaceSoftwareInfo) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MarketplaceSoftwareInfo) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetVersion

`func (o *MarketplaceSoftwareInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MarketplaceSoftwareInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MarketplaceSoftwareInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MarketplaceSoftwareInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDescription

`func (o *MarketplaceSoftwareInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MarketplaceSoftwareInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MarketplaceSoftwareInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MarketplaceSoftwareInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDescriptionEn

`func (o *MarketplaceSoftwareInfo) GetDescriptionEn() string`

GetDescriptionEn returns the DescriptionEn field if non-nil, zero value otherwise.

### GetDescriptionEnOk

`func (o *MarketplaceSoftwareInfo) GetDescriptionEnOk() (*string, bool)`

GetDescriptionEnOk returns a tuple with the DescriptionEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionEn

`func (o *MarketplaceSoftwareInfo) SetDescriptionEn(v string)`

SetDescriptionEn sets DescriptionEn field to given value.

### HasDescriptionEn

`func (o *MarketplaceSoftwareInfo) HasDescriptionEn() bool`

HasDescriptionEn returns a boolean if a field has been set.

### GetMetadata

`func (o *MarketplaceSoftwareInfo) GetMetadata() StructuresSoftwareMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MarketplaceSoftwareInfo) GetMetadataOk() (*StructuresSoftwareMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MarketplaceSoftwareInfo) SetMetadata(v StructuresSoftwareMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *MarketplaceSoftwareInfo) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetFieldData

`func (o *MarketplaceSoftwareInfo) GetFieldData() []MarketplaceFieldDesc`

GetFieldData returns the FieldData field if non-nil, zero value otherwise.

### GetFieldDataOk

`func (o *MarketplaceSoftwareInfo) GetFieldDataOk() (*[]MarketplaceFieldDesc, bool)`

GetFieldDataOk returns a tuple with the FieldData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldData

`func (o *MarketplaceSoftwareInfo) SetFieldData(v []MarketplaceFieldDesc)`

SetFieldData sets FieldData field to given value.

### HasFieldData

`func (o *MarketplaceSoftwareInfo) HasFieldData() bool`

HasFieldData returns a boolean if a field has been set.

### GetRequirements

`func (o *MarketplaceSoftwareInfo) GetRequirements() MarketplaceSoftwareInfoRequirements`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *MarketplaceSoftwareInfo) GetRequirementsOk() (*MarketplaceSoftwareInfoRequirements, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *MarketplaceSoftwareInfo) SetRequirements(v MarketplaceSoftwareInfoRequirements)`

SetRequirements sets Requirements field to given value.

### HasRequirements

`func (o *MarketplaceSoftwareInfo) HasRequirements() bool`

HasRequirements returns a boolean if a field has been set.

### GetCategory

`func (o *MarketplaceSoftwareInfo) GetCategory() []StructuresSoftwareCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *MarketplaceSoftwareInfo) GetCategoryOk() (*[]StructuresSoftwareCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *MarketplaceSoftwareInfo) SetCategory(v []StructuresSoftwareCategory)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *MarketplaceSoftwareInfo) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetSlug

`func (o *MarketplaceSoftwareInfo) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *MarketplaceSoftwareInfo) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *MarketplaceSoftwareInfo) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *MarketplaceSoftwareInfo) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetDocumentationSlug

`func (o *MarketplaceSoftwareInfo) GetDocumentationSlug() string`

GetDocumentationSlug returns the DocumentationSlug field if non-nil, zero value otherwise.

### GetDocumentationSlugOk

`func (o *MarketplaceSoftwareInfo) GetDocumentationSlugOk() (*string, bool)`

GetDocumentationSlugOk returns a tuple with the DocumentationSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationSlug

`func (o *MarketplaceSoftwareInfo) SetDocumentationSlug(v string)`

SetDocumentationSlug sets DocumentationSlug field to given value.

### HasDocumentationSlug

`func (o *MarketplaceSoftwareInfo) HasDocumentationSlug() bool`

HasDocumentationSlug returns a boolean if a field has been set.

### GetUnattendedInstallAvailable

`func (o *MarketplaceSoftwareInfo) GetUnattendedInstallAvailable() bool`

GetUnattendedInstallAvailable returns the UnattendedInstallAvailable field if non-nil, zero value otherwise.

### GetUnattendedInstallAvailableOk

`func (o *MarketplaceSoftwareInfo) GetUnattendedInstallAvailableOk() (*bool, bool)`

GetUnattendedInstallAvailableOk returns a tuple with the UnattendedInstallAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnattendedInstallAvailable

`func (o *MarketplaceSoftwareInfo) SetUnattendedInstallAvailable(v bool)`

SetUnattendedInstallAvailable sets UnattendedInstallAvailable field to given value.

### HasUnattendedInstallAvailable

`func (o *MarketplaceSoftwareInfo) HasUnattendedInstallAvailable() bool`

HasUnattendedInstallAvailable returns a boolean if a field has been set.

### GetLicense

`func (o *MarketplaceSoftwareInfo) GetLicense() []StructuresSoftwareLicense`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *MarketplaceSoftwareInfo) GetLicenseOk() (*[]StructuresSoftwareLicense, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *MarketplaceSoftwareInfo) SetLicense(v []StructuresSoftwareLicense)`

SetLicense sets License field to given value.

### HasLicense

`func (o *MarketplaceSoftwareInfo) HasLicense() bool`

HasLicense returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


