# SshKeyRemoveResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to [**StructuresSshKeyInfo**](StructuresSshKeyInfo.md) |  | [optional] 
**Error** | Pointer to [**SshKeyRemoveResponseError**](SshKeyRemoveResponseError.md) |  | [optional] 

## Methods

### NewSshKeyRemoveResponse

`func NewSshKeyRemoveResponse() *SshKeyRemoveResponse`

NewSshKeyRemoveResponse instantiates a new SshKeyRemoveResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshKeyRemoveResponseWithDefaults

`func NewSshKeyRemoveResponseWithDefaults() *SshKeyRemoveResponse`

NewSshKeyRemoveResponseWithDefaults instantiates a new SshKeyRemoveResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *SshKeyRemoveResponse) GetKey() StructuresSshKeyInfo`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SshKeyRemoveResponse) GetKeyOk() (*StructuresSshKeyInfo, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SshKeyRemoveResponse) SetKey(v StructuresSshKeyInfo)`

SetKey sets Key field to given value.

### HasKey

`func (o *SshKeyRemoveResponse) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetError

`func (o *SshKeyRemoveResponse) GetError() SshKeyRemoveResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SshKeyRemoveResponse) GetErrorOk() (*SshKeyRemoveResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SshKeyRemoveResponse) SetError(v SshKeyRemoveResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *SshKeyRemoveResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


