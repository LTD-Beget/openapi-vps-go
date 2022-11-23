# ManageReinstallResponseError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**VariableError** | Pointer to [**ManageReinstallResponseErrorSoftwareVariableError**](ManageReinstallResponseErrorSoftwareVariableError.md) |  | [optional] 

## Methods

### NewManageReinstallResponseError

`func NewManageReinstallResponseError() *ManageReinstallResponseError`

NewManageReinstallResponseError instantiates a new ManageReinstallResponseError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManageReinstallResponseErrorWithDefaults

`func NewManageReinstallResponseErrorWithDefaults() *ManageReinstallResponseError`

NewManageReinstallResponseErrorWithDefaults instantiates a new ManageReinstallResponseError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ManageReinstallResponseError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ManageReinstallResponseError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ManageReinstallResponseError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ManageReinstallResponseError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *ManageReinstallResponseError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ManageReinstallResponseError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ManageReinstallResponseError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ManageReinstallResponseError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetVariableError

`func (o *ManageReinstallResponseError) GetVariableError() ManageReinstallResponseErrorSoftwareVariableError`

GetVariableError returns the VariableError field if non-nil, zero value otherwise.

### GetVariableErrorOk

`func (o *ManageReinstallResponseError) GetVariableErrorOk() (*ManageReinstallResponseErrorSoftwareVariableError, bool)`

GetVariableErrorOk returns a tuple with the VariableError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableError

`func (o *ManageReinstallResponseError) SetVariableError(v ManageReinstallResponseErrorSoftwareVariableError)`

SetVariableError sets VariableError field to given value.

### HasVariableError

`func (o *ManageReinstallResponseError) HasVariableError() bool`

HasVariableError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


