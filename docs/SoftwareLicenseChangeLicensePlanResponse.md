# SoftwareLicenseChangeLicensePlanResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**License** | Pointer to [**StructuresIssuedSoftwareLicense**](StructuresIssuedSoftwareLicense.md) |  | [optional] 
**Error** | Pointer to [**SoftwareLicenseChangeLicensePlanResponseError**](SoftwareLicenseChangeLicensePlanResponseError.md) |  | [optional] 

## Methods

### NewSoftwareLicenseChangeLicensePlanResponse

`func NewSoftwareLicenseChangeLicensePlanResponse() *SoftwareLicenseChangeLicensePlanResponse`

NewSoftwareLicenseChangeLicensePlanResponse instantiates a new SoftwareLicenseChangeLicensePlanResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwareLicenseChangeLicensePlanResponseWithDefaults

`func NewSoftwareLicenseChangeLicensePlanResponseWithDefaults() *SoftwareLicenseChangeLicensePlanResponse`

NewSoftwareLicenseChangeLicensePlanResponseWithDefaults instantiates a new SoftwareLicenseChangeLicensePlanResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicense

`func (o *SoftwareLicenseChangeLicensePlanResponse) GetLicense() StructuresIssuedSoftwareLicense`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *SoftwareLicenseChangeLicensePlanResponse) GetLicenseOk() (*StructuresIssuedSoftwareLicense, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *SoftwareLicenseChangeLicensePlanResponse) SetLicense(v StructuresIssuedSoftwareLicense)`

SetLicense sets License field to given value.

### HasLicense

`func (o *SoftwareLicenseChangeLicensePlanResponse) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetError

`func (o *SoftwareLicenseChangeLicensePlanResponse) GetError() SoftwareLicenseChangeLicensePlanResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SoftwareLicenseChangeLicensePlanResponse) GetErrorOk() (*SoftwareLicenseChangeLicensePlanResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SoftwareLicenseChangeLicensePlanResponse) SetError(v SoftwareLicenseChangeLicensePlanResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *SoftwareLicenseChangeLicensePlanResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


