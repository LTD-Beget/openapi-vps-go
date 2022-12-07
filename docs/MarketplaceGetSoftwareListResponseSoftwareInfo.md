# MarketplaceGetSoftwareListResponseSoftwareInfo

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
**Requirements** | Pointer to [**MarketplaceGetSoftwareListResponseSoftwareInfoRequirements**](MarketplaceGetSoftwareListResponseSoftwareInfoRequirements.md) |  | [optional] 
**Category** | Pointer to [**[]StructuresSoftwareCategory**](StructuresSoftwareCategory.md) |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**DocumentationSlug** | Pointer to **string** |  | [optional] 

## Methods

### NewMarketplaceGetSoftwareListResponseSoftwareInfo

`func NewMarketplaceGetSoftwareListResponseSoftwareInfo() *MarketplaceGetSoftwareListResponseSoftwareInfo`

NewMarketplaceGetSoftwareListResponseSoftwareInfo instantiates a new MarketplaceGetSoftwareListResponseSoftwareInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceGetSoftwareListResponseSoftwareInfoWithDefaults

`func NewMarketplaceGetSoftwareListResponseSoftwareInfoWithDefaults() *MarketplaceGetSoftwareListResponseSoftwareInfo`

NewMarketplaceGetSoftwareListResponseSoftwareInfoWithDefaults instantiates a new MarketplaceGetSoftwareListResponseSoftwareInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MarketplaceGetSoftwareListResponseSoftwareInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketplaceGetSoftwareListResponseSoftwareInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketplaceGetSoftwareListResponseSoftwareInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MarketplaceGetSoftwareListResponseSoftwareInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MarketplaceGetSoftwareListResponseSoftwareInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MarketplaceGetSoftwareListResponseSoftwareInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MarketplaceGetSoftwareListResponseSoftwareInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MarketplaceGetSoftwareListResponseSoftwareInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *MarketplaceGetSoftwareListResponseSoftwareInfo) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MarketplaceGetSoftwareListResponseSoftwareInfo) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MarketplaceGetSoftwareListResponseSoftwareInfo) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MarketplaceGetSoftwareListResponseSoftwareInfo) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetVersion

`func (o *MarketplaceGetSoftwareListResponseSoftwareInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MarketplaceGetSoftwareListResponseSoftwareInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MarketplaceGetSoftwareListResponseSoftwareInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MarketplaceGetSoftwareListResponseSoftwareInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDescription

`func (o *MarketplaceGetSoftwareListResponseSoftwareInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MarketplaceGetSoftwareListResponseSoftwareInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MarketplaceGetSoftwareListResponseSoftwareInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MarketplaceGetSoftwareListResponseSoftwareInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDescriptionEn

`func (o *MarketplaceGetSoftwareListResponseSoftwareInfo) GetDescriptionEn() string`

GetDescriptionEn returns the DescriptionEn field if non-nil, zero value otherwise.

### GetDescriptionEnOk

`func (o *MarketplaceGetSoftwareListResponseSoftwareInfo) GetDescriptionEnOk() (*string, bool)`

GetDescriptionEnOk returns a tuple with the DescriptionEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionEn

`func (o *MarketplaceGetSoftwareListResponseSoftwareInfo) SetDescriptionEn(v string)`

SetDescriptionEn sets DescriptionEn field to given value.

### HasDescriptionEn

`func (o *MarketplaceGetSoftwareListResponseSoftwareInfo) HasDescriptionEn() bool`

HasDescriptionEn returns a boolean if a field has been set.

### GetMetadata

`func (o *MarketplaceGetSoftwareListResponseSoftwareInfo) GetMetadata() StructuresSoftwareMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MarketplaceGetSoftwareListResponseSoftwareInfo) GetMetadataOk() (*StructuresSoftwareMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MarketplaceGetSoftwareListResponseSoftwareInfo) SetMetadata(v StructuresSoftwareMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *MarketplaceGetSoftwareListResponseSoftwareInfo) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetFieldData

`func (o *MarketplaceGetSoftwareListResponseSoftwareInfo) GetFieldData() []MarketplaceFieldDesc`

GetFieldData returns the FieldData field if non-nil, zero value otherwise.

### GetFieldDataOk

`func (o *MarketplaceGetSoftwareListResponseSoftwareInfo) GetFieldDataOk() (*[]MarketplaceFieldDesc, bool)`

GetFieldDataOk returns a tuple with the FieldData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldData

`func (o *MarketplaceGetSoftwareListResponseSoftwareInfo) SetFieldData(v []MarketplaceFieldDesc)`

SetFieldData sets FieldData field to given value.

### HasFieldData

`func (o *MarketplaceGetSoftwareListResponseSoftwareInfo) HasFieldData() bool`

HasFieldData returns a boolean if a field has been set.

### GetRequirements

`func (o *MarketplaceGetSoftwareListResponseSoftwareInfo) GetRequirements() MarketplaceGetSoftwareListResponseSoftwareInfoRequirements`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *MarketplaceGetSoftwareListResponseSoftwareInfo) GetRequirementsOk() (*MarketplaceGetSoftwareListResponseSoftwareInfoRequirements, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *MarketplaceGetSoftwareListResponseSoftwareInfo) SetRequirements(v MarketplaceGetSoftwareListResponseSoftwareInfoRequirements)`

SetRequirements sets Requirements field to given value.

### HasRequirements

`func (o *MarketplaceGetSoftwareListResponseSoftwareInfo) HasRequirements() bool`

HasRequirements returns a boolean if a field has been set.

### GetCategory

`func (o *MarketplaceGetSoftwareListResponseSoftwareInfo) GetCategory() []StructuresSoftwareCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *MarketplaceGetSoftwareListResponseSoftwareInfo) GetCategoryOk() (*[]StructuresSoftwareCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *MarketplaceGetSoftwareListResponseSoftwareInfo) SetCategory(v []StructuresSoftwareCategory)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *MarketplaceGetSoftwareListResponseSoftwareInfo) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetSlug

`func (o *MarketplaceGetSoftwareListResponseSoftwareInfo) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *MarketplaceGetSoftwareListResponseSoftwareInfo) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *MarketplaceGetSoftwareListResponseSoftwareInfo) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *MarketplaceGetSoftwareListResponseSoftwareInfo) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetDocumentationSlug

`func (o *MarketplaceGetSoftwareListResponseSoftwareInfo) GetDocumentationSlug() string`

GetDocumentationSlug returns the DocumentationSlug field if non-nil, zero value otherwise.

### GetDocumentationSlugOk

`func (o *MarketplaceGetSoftwareListResponseSoftwareInfo) GetDocumentationSlugOk() (*string, bool)`

GetDocumentationSlugOk returns a tuple with the DocumentationSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationSlug

`func (o *MarketplaceGetSoftwareListResponseSoftwareInfo) SetDocumentationSlug(v string)`

SetDocumentationSlug sets DocumentationSlug field to given value.

### HasDocumentationSlug

`func (o *MarketplaceGetSoftwareListResponseSoftwareInfo) HasDocumentationSlug() bool`

HasDocumentationSlug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


