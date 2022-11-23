# StructuresFileInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsLink** | Pointer to **bool** |  | [optional] 
**IsDir** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Ext** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Group** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **string** |  | [optional] 
**Mtime** | Pointer to **float64** |  | [optional] 
**MtimeStr** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to **string** |  | [optional] 

## Methods

### NewStructuresFileInfo

`func NewStructuresFileInfo() *StructuresFileInfo`

NewStructuresFileInfo instantiates a new StructuresFileInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStructuresFileInfoWithDefaults

`func NewStructuresFileInfoWithDefaults() *StructuresFileInfo`

NewStructuresFileInfoWithDefaults instantiates a new StructuresFileInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsLink

`func (o *StructuresFileInfo) GetIsLink() bool`

GetIsLink returns the IsLink field if non-nil, zero value otherwise.

### GetIsLinkOk

`func (o *StructuresFileInfo) GetIsLinkOk() (*bool, bool)`

GetIsLinkOk returns a tuple with the IsLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLink

`func (o *StructuresFileInfo) SetIsLink(v bool)`

SetIsLink sets IsLink field to given value.

### HasIsLink

`func (o *StructuresFileInfo) HasIsLink() bool`

HasIsLink returns a boolean if a field has been set.

### GetIsDir

`func (o *StructuresFileInfo) GetIsDir() bool`

GetIsDir returns the IsDir field if non-nil, zero value otherwise.

### GetIsDirOk

`func (o *StructuresFileInfo) GetIsDirOk() (*bool, bool)`

GetIsDirOk returns a tuple with the IsDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDir

`func (o *StructuresFileInfo) SetIsDir(v bool)`

SetIsDir sets IsDir field to given value.

### HasIsDir

`func (o *StructuresFileInfo) HasIsDir() bool`

HasIsDir returns a boolean if a field has been set.

### GetName

`func (o *StructuresFileInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StructuresFileInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StructuresFileInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StructuresFileInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExt

`func (o *StructuresFileInfo) GetExt() string`

GetExt returns the Ext field if non-nil, zero value otherwise.

### GetExtOk

`func (o *StructuresFileInfo) GetExtOk() (*string, bool)`

GetExtOk returns a tuple with the Ext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExt

`func (o *StructuresFileInfo) SetExt(v string)`

SetExt sets Ext field to given value.

### HasExt

`func (o *StructuresFileInfo) HasExt() bool`

HasExt returns a boolean if a field has been set.

### GetPath

`func (o *StructuresFileInfo) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *StructuresFileInfo) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *StructuresFileInfo) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *StructuresFileInfo) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetOwner

`func (o *StructuresFileInfo) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *StructuresFileInfo) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *StructuresFileInfo) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *StructuresFileInfo) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetGroup

`func (o *StructuresFileInfo) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *StructuresFileInfo) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *StructuresFileInfo) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *StructuresFileInfo) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetSize

`func (o *StructuresFileInfo) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *StructuresFileInfo) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *StructuresFileInfo) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *StructuresFileInfo) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetMtime

`func (o *StructuresFileInfo) GetMtime() float64`

GetMtime returns the Mtime field if non-nil, zero value otherwise.

### GetMtimeOk

`func (o *StructuresFileInfo) GetMtimeOk() (*float64, bool)`

GetMtimeOk returns a tuple with the Mtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtime

`func (o *StructuresFileInfo) SetMtime(v float64)`

SetMtime sets Mtime field to given value.

### HasMtime

`func (o *StructuresFileInfo) HasMtime() bool`

HasMtime returns a boolean if a field has been set.

### GetMtimeStr

`func (o *StructuresFileInfo) GetMtimeStr() string`

GetMtimeStr returns the MtimeStr field if non-nil, zero value otherwise.

### GetMtimeStrOk

`func (o *StructuresFileInfo) GetMtimeStrOk() (*string, bool)`

GetMtimeStrOk returns a tuple with the MtimeStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtimeStr

`func (o *StructuresFileInfo) SetMtimeStr(v string)`

SetMtimeStr sets MtimeStr field to given value.

### HasMtimeStr

`func (o *StructuresFileInfo) HasMtimeStr() bool`

HasMtimeStr returns a boolean if a field has been set.

### GetMode

`func (o *StructuresFileInfo) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *StructuresFileInfo) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *StructuresFileInfo) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *StructuresFileInfo) HasMode() bool`

HasMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


