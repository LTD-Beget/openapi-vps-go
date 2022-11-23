# StructuresOrderInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**VpsId** | Pointer to **string** |  | [optional] 
**VpsName** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**DateCreate** | Pointer to **string** |  | [optional] 
**DateComplete** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **[]string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**CopyInfo** | Pointer to [**StructuresCopyInfo**](StructuresCopyInfo.md) |  | [optional] 
**AffectLive** | Pointer to **bool** |  | [optional] 
**Target** | Pointer to **string** |  | [optional] 
**ErrorDetails** | Pointer to [**StructuresOrderInfoErrorDetails**](StructuresOrderInfoErrorDetails.md) |  | [optional] 

## Methods

### NewStructuresOrderInfo

`func NewStructuresOrderInfo() *StructuresOrderInfo`

NewStructuresOrderInfo instantiates a new StructuresOrderInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStructuresOrderInfoWithDefaults

`func NewStructuresOrderInfoWithDefaults() *StructuresOrderInfo`

NewStructuresOrderInfoWithDefaults instantiates a new StructuresOrderInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StructuresOrderInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StructuresOrderInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StructuresOrderInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StructuresOrderInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVpsId

`func (o *StructuresOrderInfo) GetVpsId() string`

GetVpsId returns the VpsId field if non-nil, zero value otherwise.

### GetVpsIdOk

`func (o *StructuresOrderInfo) GetVpsIdOk() (*string, bool)`

GetVpsIdOk returns a tuple with the VpsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpsId

`func (o *StructuresOrderInfo) SetVpsId(v string)`

SetVpsId sets VpsId field to given value.

### HasVpsId

`func (o *StructuresOrderInfo) HasVpsId() bool`

HasVpsId returns a boolean if a field has been set.

### GetVpsName

`func (o *StructuresOrderInfo) GetVpsName() string`

GetVpsName returns the VpsName field if non-nil, zero value otherwise.

### GetVpsNameOk

`func (o *StructuresOrderInfo) GetVpsNameOk() (*string, bool)`

GetVpsNameOk returns a tuple with the VpsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpsName

`func (o *StructuresOrderInfo) SetVpsName(v string)`

SetVpsName sets VpsName field to given value.

### HasVpsName

`func (o *StructuresOrderInfo) HasVpsName() bool`

HasVpsName returns a boolean if a field has been set.

### GetType

`func (o *StructuresOrderInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StructuresOrderInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StructuresOrderInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StructuresOrderInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDateCreate

`func (o *StructuresOrderInfo) GetDateCreate() string`

GetDateCreate returns the DateCreate field if non-nil, zero value otherwise.

### GetDateCreateOk

`func (o *StructuresOrderInfo) GetDateCreateOk() (*string, bool)`

GetDateCreateOk returns a tuple with the DateCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreate

`func (o *StructuresOrderInfo) SetDateCreate(v string)`

SetDateCreate sets DateCreate field to given value.

### HasDateCreate

`func (o *StructuresOrderInfo) HasDateCreate() bool`

HasDateCreate returns a boolean if a field has been set.

### GetDateComplete

`func (o *StructuresOrderInfo) GetDateComplete() string`

GetDateComplete returns the DateComplete field if non-nil, zero value otherwise.

### GetDateCompleteOk

`func (o *StructuresOrderInfo) GetDateCompleteOk() (*string, bool)`

GetDateCompleteOk returns a tuple with the DateComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateComplete

`func (o *StructuresOrderInfo) SetDateComplete(v string)`

SetDateComplete sets DateComplete field to given value.

### HasDateComplete

`func (o *StructuresOrderInfo) HasDateComplete() bool`

HasDateComplete returns a boolean if a field has been set.

### GetPath

`func (o *StructuresOrderInfo) GetPath() []string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *StructuresOrderInfo) GetPathOk() (*[]string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *StructuresOrderInfo) SetPath(v []string)`

SetPath sets Path field to given value.

### HasPath

`func (o *StructuresOrderInfo) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetStatus

`func (o *StructuresOrderInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StructuresOrderInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StructuresOrderInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StructuresOrderInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCopyInfo

`func (o *StructuresOrderInfo) GetCopyInfo() StructuresCopyInfo`

GetCopyInfo returns the CopyInfo field if non-nil, zero value otherwise.

### GetCopyInfoOk

`func (o *StructuresOrderInfo) GetCopyInfoOk() (*StructuresCopyInfo, bool)`

GetCopyInfoOk returns a tuple with the CopyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyInfo

`func (o *StructuresOrderInfo) SetCopyInfo(v StructuresCopyInfo)`

SetCopyInfo sets CopyInfo field to given value.

### HasCopyInfo

`func (o *StructuresOrderInfo) HasCopyInfo() bool`

HasCopyInfo returns a boolean if a field has been set.

### GetAffectLive

`func (o *StructuresOrderInfo) GetAffectLive() bool`

GetAffectLive returns the AffectLive field if non-nil, zero value otherwise.

### GetAffectLiveOk

`func (o *StructuresOrderInfo) GetAffectLiveOk() (*bool, bool)`

GetAffectLiveOk returns a tuple with the AffectLive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectLive

`func (o *StructuresOrderInfo) SetAffectLive(v bool)`

SetAffectLive sets AffectLive field to given value.

### HasAffectLive

`func (o *StructuresOrderInfo) HasAffectLive() bool`

HasAffectLive returns a boolean if a field has been set.

### GetTarget

`func (o *StructuresOrderInfo) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *StructuresOrderInfo) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *StructuresOrderInfo) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *StructuresOrderInfo) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetErrorDetails

`func (o *StructuresOrderInfo) GetErrorDetails() StructuresOrderInfoErrorDetails`

GetErrorDetails returns the ErrorDetails field if non-nil, zero value otherwise.

### GetErrorDetailsOk

`func (o *StructuresOrderInfo) GetErrorDetailsOk() (*StructuresOrderInfoErrorDetails, bool)`

GetErrorDetailsOk returns a tuple with the ErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetails

`func (o *StructuresOrderInfo) SetErrorDetails(v StructuresOrderInfoErrorDetails)`

SetErrorDetails sets ErrorDetails field to given value.

### HasErrorDetails

`func (o *StructuresOrderInfo) HasErrorDetails() bool`

HasErrorDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


