# SshKeyUpdateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to [**StructuresSshKeyInfo**](StructuresSshKeyInfo.md) |  | [optional] 
**Error** | Pointer to [**SshKeyUpdateResponseError**](SshKeyUpdateResponseError.md) |  | [optional] 

## Methods

### NewSshKeyUpdateResponse

`func NewSshKeyUpdateResponse() *SshKeyUpdateResponse`

NewSshKeyUpdateResponse instantiates a new SshKeyUpdateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshKeyUpdateResponseWithDefaults

`func NewSshKeyUpdateResponseWithDefaults() *SshKeyUpdateResponse`

NewSshKeyUpdateResponseWithDefaults instantiates a new SshKeyUpdateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *SshKeyUpdateResponse) GetKey() StructuresSshKeyInfo`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SshKeyUpdateResponse) GetKeyOk() (*StructuresSshKeyInfo, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SshKeyUpdateResponse) SetKey(v StructuresSshKeyInfo)`

SetKey sets Key field to given value.

### HasKey

`func (o *SshKeyUpdateResponse) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetError

`func (o *SshKeyUpdateResponse) GetError() SshKeyUpdateResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SshKeyUpdateResponse) GetErrorOk() (*SshKeyUpdateResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SshKeyUpdateResponse) SetError(v SshKeyUpdateResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *SshKeyUpdateResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


