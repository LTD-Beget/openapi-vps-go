# ManageCreateVpsResponseError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**VariableError** | Pointer to [**ManageCreateVpsResponseErrorSoftwareVariableError**](ManageCreateVpsResponseErrorSoftwareVariableError.md) |  | [optional] 
**InsufficientFundsError** | Pointer to [**ManageCreateVpsResponseErrorInsufficientFundsError**](ManageCreateVpsResponseErrorInsufficientFundsError.md) |  | [optional] 

## Methods

### NewManageCreateVpsResponseError

`func NewManageCreateVpsResponseError() *ManageCreateVpsResponseError`

NewManageCreateVpsResponseError instantiates a new ManageCreateVpsResponseError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManageCreateVpsResponseErrorWithDefaults

`func NewManageCreateVpsResponseErrorWithDefaults() *ManageCreateVpsResponseError`

NewManageCreateVpsResponseErrorWithDefaults instantiates a new ManageCreateVpsResponseError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ManageCreateVpsResponseError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ManageCreateVpsResponseError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ManageCreateVpsResponseError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ManageCreateVpsResponseError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *ManageCreateVpsResponseError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ManageCreateVpsResponseError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ManageCreateVpsResponseError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ManageCreateVpsResponseError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetVariableError

`func (o *ManageCreateVpsResponseError) GetVariableError() ManageCreateVpsResponseErrorSoftwareVariableError`

GetVariableError returns the VariableError field if non-nil, zero value otherwise.

### GetVariableErrorOk

`func (o *ManageCreateVpsResponseError) GetVariableErrorOk() (*ManageCreateVpsResponseErrorSoftwareVariableError, bool)`

GetVariableErrorOk returns a tuple with the VariableError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableError

`func (o *ManageCreateVpsResponseError) SetVariableError(v ManageCreateVpsResponseErrorSoftwareVariableError)`

SetVariableError sets VariableError field to given value.

### HasVariableError

`func (o *ManageCreateVpsResponseError) HasVariableError() bool`

HasVariableError returns a boolean if a field has been set.

### GetInsufficientFundsError

`func (o *ManageCreateVpsResponseError) GetInsufficientFundsError() ManageCreateVpsResponseErrorInsufficientFundsError`

GetInsufficientFundsError returns the InsufficientFundsError field if non-nil, zero value otherwise.

### GetInsufficientFundsErrorOk

`func (o *ManageCreateVpsResponseError) GetInsufficientFundsErrorOk() (*ManageCreateVpsResponseErrorInsufficientFundsError, bool)`

GetInsufficientFundsErrorOk returns a tuple with the InsufficientFundsError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsufficientFundsError

`func (o *ManageCreateVpsResponseError) SetInsufficientFundsError(v ManageCreateVpsResponseErrorInsufficientFundsError)`

SetInsufficientFundsError sets InsufficientFundsError field to given value.

### HasInsufficientFundsError

`func (o *ManageCreateVpsResponseError) HasInsufficientFundsError() bool`

HasInsufficientFundsError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


