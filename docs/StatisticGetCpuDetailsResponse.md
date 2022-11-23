# StatisticGetCpuDetailsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **[]string** |  | [optional] 
**User** | Pointer to **[]float64** |  | [optional] 
**Nice** | Pointer to **[]float64** |  | [optional] 
**System** | Pointer to **[]float64** |  | [optional] 
**Idle** | Pointer to **[]float64** |  | [optional] 
**Iowait** | Pointer to **[]float64** |  | [optional] 
**Irq** | Pointer to **[]float64** |  | [optional] 
**Softirq** | Pointer to **[]float64** |  | [optional] 

## Methods

### NewStatisticGetCpuDetailsResponse

`func NewStatisticGetCpuDetailsResponse() *StatisticGetCpuDetailsResponse`

NewStatisticGetCpuDetailsResponse instantiates a new StatisticGetCpuDetailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatisticGetCpuDetailsResponseWithDefaults

`func NewStatisticGetCpuDetailsResponseWithDefaults() *StatisticGetCpuDetailsResponse`

NewStatisticGetCpuDetailsResponseWithDefaults instantiates a new StatisticGetCpuDetailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *StatisticGetCpuDetailsResponse) GetDate() []string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *StatisticGetCpuDetailsResponse) GetDateOk() (*[]string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *StatisticGetCpuDetailsResponse) SetDate(v []string)`

SetDate sets Date field to given value.

### HasDate

`func (o *StatisticGetCpuDetailsResponse) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetUser

`func (o *StatisticGetCpuDetailsResponse) GetUser() []float64`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *StatisticGetCpuDetailsResponse) GetUserOk() (*[]float64, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *StatisticGetCpuDetailsResponse) SetUser(v []float64)`

SetUser sets User field to given value.

### HasUser

`func (o *StatisticGetCpuDetailsResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetNice

`func (o *StatisticGetCpuDetailsResponse) GetNice() []float64`

GetNice returns the Nice field if non-nil, zero value otherwise.

### GetNiceOk

`func (o *StatisticGetCpuDetailsResponse) GetNiceOk() (*[]float64, bool)`

GetNiceOk returns a tuple with the Nice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNice

`func (o *StatisticGetCpuDetailsResponse) SetNice(v []float64)`

SetNice sets Nice field to given value.

### HasNice

`func (o *StatisticGetCpuDetailsResponse) HasNice() bool`

HasNice returns a boolean if a field has been set.

### GetSystem

`func (o *StatisticGetCpuDetailsResponse) GetSystem() []float64`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *StatisticGetCpuDetailsResponse) GetSystemOk() (*[]float64, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *StatisticGetCpuDetailsResponse) SetSystem(v []float64)`

SetSystem sets System field to given value.

### HasSystem

`func (o *StatisticGetCpuDetailsResponse) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetIdle

`func (o *StatisticGetCpuDetailsResponse) GetIdle() []float64`

GetIdle returns the Idle field if non-nil, zero value otherwise.

### GetIdleOk

`func (o *StatisticGetCpuDetailsResponse) GetIdleOk() (*[]float64, bool)`

GetIdleOk returns a tuple with the Idle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdle

`func (o *StatisticGetCpuDetailsResponse) SetIdle(v []float64)`

SetIdle sets Idle field to given value.

### HasIdle

`func (o *StatisticGetCpuDetailsResponse) HasIdle() bool`

HasIdle returns a boolean if a field has been set.

### GetIowait

`func (o *StatisticGetCpuDetailsResponse) GetIowait() []float64`

GetIowait returns the Iowait field if non-nil, zero value otherwise.

### GetIowaitOk

`func (o *StatisticGetCpuDetailsResponse) GetIowaitOk() (*[]float64, bool)`

GetIowaitOk returns a tuple with the Iowait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIowait

`func (o *StatisticGetCpuDetailsResponse) SetIowait(v []float64)`

SetIowait sets Iowait field to given value.

### HasIowait

`func (o *StatisticGetCpuDetailsResponse) HasIowait() bool`

HasIowait returns a boolean if a field has been set.

### GetIrq

`func (o *StatisticGetCpuDetailsResponse) GetIrq() []float64`

GetIrq returns the Irq field if non-nil, zero value otherwise.

### GetIrqOk

`func (o *StatisticGetCpuDetailsResponse) GetIrqOk() (*[]float64, bool)`

GetIrqOk returns a tuple with the Irq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIrq

`func (o *StatisticGetCpuDetailsResponse) SetIrq(v []float64)`

SetIrq sets Irq field to given value.

### HasIrq

`func (o *StatisticGetCpuDetailsResponse) HasIrq() bool`

HasIrq returns a boolean if a field has been set.

### GetSoftirq

`func (o *StatisticGetCpuDetailsResponse) GetSoftirq() []float64`

GetSoftirq returns the Softirq field if non-nil, zero value otherwise.

### GetSoftirqOk

`func (o *StatisticGetCpuDetailsResponse) GetSoftirqOk() (*[]float64, bool)`

GetSoftirqOk returns a tuple with the Softirq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftirq

`func (o *StatisticGetCpuDetailsResponse) SetSoftirq(v []float64)`

SetSoftirq sets Softirq field to given value.

### HasSoftirq

`func (o *StatisticGetCpuDetailsResponse) HasSoftirq() bool`

HasSoftirq returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


