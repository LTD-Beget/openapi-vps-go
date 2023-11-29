# StructuresIssuedSoftwareLicense

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Plan** | Pointer to **string** |  | [optional] 
**AdditionalInfo** | Pointer to **string** |  | [optional] 
**AdditionalInfoEn** | Pointer to **string** |  | [optional] 
**Primary** | Pointer to **bool** |  | [optional] 
**Billing** | Pointer to [**StructuresSoftwareLicenseBillingType**](StructuresSoftwareLicenseBillingType.md) |  | [optional] 

## Methods

### NewStructuresIssuedSoftwareLicense

`func NewStructuresIssuedSoftwareLicense() *StructuresIssuedSoftwareLicense`

NewStructuresIssuedSoftwareLicense instantiates a new StructuresIssuedSoftwareLicense object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStructuresIssuedSoftwareLicenseWithDefaults

`func NewStructuresIssuedSoftwareLicenseWithDefaults() *StructuresIssuedSoftwareLicense`

NewStructuresIssuedSoftwareLicenseWithDefaults instantiates a new StructuresIssuedSoftwareLicense object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StructuresIssuedSoftwareLicense) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StructuresIssuedSoftwareLicense) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StructuresIssuedSoftwareLicense) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *StructuresIssuedSoftwareLicense) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *StructuresIssuedSoftwareLicense) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StructuresIssuedSoftwareLicense) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StructuresIssuedSoftwareLicense) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StructuresIssuedSoftwareLicense) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDisplayName

`func (o *StructuresIssuedSoftwareLicense) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *StructuresIssuedSoftwareLicense) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *StructuresIssuedSoftwareLicense) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *StructuresIssuedSoftwareLicense) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetPlan

`func (o *StructuresIssuedSoftwareLicense) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *StructuresIssuedSoftwareLicense) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *StructuresIssuedSoftwareLicense) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *StructuresIssuedSoftwareLicense) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *StructuresIssuedSoftwareLicense) GetAdditionalInfo() string`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *StructuresIssuedSoftwareLicense) GetAdditionalInfoOk() (*string, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *StructuresIssuedSoftwareLicense) SetAdditionalInfo(v string)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *StructuresIssuedSoftwareLicense) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.

### GetAdditionalInfoEn

`func (o *StructuresIssuedSoftwareLicense) GetAdditionalInfoEn() string`

GetAdditionalInfoEn returns the AdditionalInfoEn field if non-nil, zero value otherwise.

### GetAdditionalInfoEnOk

`func (o *StructuresIssuedSoftwareLicense) GetAdditionalInfoEnOk() (*string, bool)`

GetAdditionalInfoEnOk returns a tuple with the AdditionalInfoEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfoEn

`func (o *StructuresIssuedSoftwareLicense) SetAdditionalInfoEn(v string)`

SetAdditionalInfoEn sets AdditionalInfoEn field to given value.

### HasAdditionalInfoEn

`func (o *StructuresIssuedSoftwareLicense) HasAdditionalInfoEn() bool`

HasAdditionalInfoEn returns a boolean if a field has been set.

### GetPrimary

`func (o *StructuresIssuedSoftwareLicense) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *StructuresIssuedSoftwareLicense) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *StructuresIssuedSoftwareLicense) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *StructuresIssuedSoftwareLicense) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetBilling

`func (o *StructuresIssuedSoftwareLicense) GetBilling() StructuresSoftwareLicenseBillingType`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *StructuresIssuedSoftwareLicense) GetBillingOk() (*StructuresSoftwareLicenseBillingType, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilling

`func (o *StructuresIssuedSoftwareLicense) SetBilling(v StructuresSoftwareLicenseBillingType)`

SetBilling sets Billing field to given value.

### HasBilling

`func (o *StructuresIssuedSoftwareLicense) HasBilling() bool`

HasBilling returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


