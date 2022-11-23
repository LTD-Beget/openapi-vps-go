# ManageGetFileManagerSettingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**ManageGetFileManagerSettingsResponseCredentials**](ManageGetFileManagerSettingsResponseCredentials.md) |  | [optional] 
**Error** | Pointer to [**ManageGetFileManagerSettingsResponseError**](ManageGetFileManagerSettingsResponseError.md) |  | [optional] 

## Methods

### NewManageGetFileManagerSettingsResponse

`func NewManageGetFileManagerSettingsResponse() *ManageGetFileManagerSettingsResponse`

NewManageGetFileManagerSettingsResponse instantiates a new ManageGetFileManagerSettingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManageGetFileManagerSettingsResponseWithDefaults

`func NewManageGetFileManagerSettingsResponseWithDefaults() *ManageGetFileManagerSettingsResponse`

NewManageGetFileManagerSettingsResponseWithDefaults instantiates a new ManageGetFileManagerSettingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *ManageGetFileManagerSettingsResponse) GetCredentials() ManageGetFileManagerSettingsResponseCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *ManageGetFileManagerSettingsResponse) GetCredentialsOk() (*ManageGetFileManagerSettingsResponseCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *ManageGetFileManagerSettingsResponse) SetCredentials(v ManageGetFileManagerSettingsResponseCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *ManageGetFileManagerSettingsResponse) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetError

`func (o *ManageGetFileManagerSettingsResponse) GetError() ManageGetFileManagerSettingsResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ManageGetFileManagerSettingsResponse) GetErrorOk() (*ManageGetFileManagerSettingsResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ManageGetFileManagerSettingsResponse) SetError(v ManageGetFileManagerSettingsResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *ManageGetFileManagerSettingsResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


