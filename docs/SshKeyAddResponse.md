# SshKeyAddResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to [**StructuresSshKeyInfo**](StructuresSshKeyInfo.md) |  | [optional] 
**Error** | Pointer to [**SshKeyAddResponseError**](SshKeyAddResponseError.md) |  | [optional] 

## Methods

### NewSshKeyAddResponse

`func NewSshKeyAddResponse() *SshKeyAddResponse`

NewSshKeyAddResponse instantiates a new SshKeyAddResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshKeyAddResponseWithDefaults

`func NewSshKeyAddResponseWithDefaults() *SshKeyAddResponse`

NewSshKeyAddResponseWithDefaults instantiates a new SshKeyAddResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *SshKeyAddResponse) GetKey() StructuresSshKeyInfo`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SshKeyAddResponse) GetKeyOk() (*StructuresSshKeyInfo, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SshKeyAddResponse) SetKey(v StructuresSshKeyInfo)`

SetKey sets Key field to given value.

### HasKey

`func (o *SshKeyAddResponse) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetError

`func (o *SshKeyAddResponse) GetError() SshKeyAddResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SshKeyAddResponse) GetErrorOk() (*SshKeyAddResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SshKeyAddResponse) SetError(v SshKeyAddResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *SshKeyAddResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


