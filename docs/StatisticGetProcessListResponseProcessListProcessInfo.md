# StatisticGetProcessListResponseProcessListProcessInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pid** | Pointer to **int32** |  | [optional] 
**Nice** | Pointer to **int32** |  | [optional] 
**Virt** | Pointer to **string** |  | [optional] 
**Rss** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**CpuPercent** | Pointer to **float32** |  | [optional] 
**MemPercent** | Pointer to **float32** |  | [optional] 
**TimeRunning** | Pointer to **float32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 

## Methods

### NewStatisticGetProcessListResponseProcessListProcessInfo

`func NewStatisticGetProcessListResponseProcessListProcessInfo() *StatisticGetProcessListResponseProcessListProcessInfo`

NewStatisticGetProcessListResponseProcessListProcessInfo instantiates a new StatisticGetProcessListResponseProcessListProcessInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatisticGetProcessListResponseProcessListProcessInfoWithDefaults

`func NewStatisticGetProcessListResponseProcessListProcessInfoWithDefaults() *StatisticGetProcessListResponseProcessListProcessInfo`

NewStatisticGetProcessListResponseProcessListProcessInfoWithDefaults instantiates a new StatisticGetProcessListResponseProcessListProcessInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPid

`func (o *StatisticGetProcessListResponseProcessListProcessInfo) GetPid() int32`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *StatisticGetProcessListResponseProcessListProcessInfo) GetPidOk() (*int32, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *StatisticGetProcessListResponseProcessListProcessInfo) SetPid(v int32)`

SetPid sets Pid field to given value.

### HasPid

`func (o *StatisticGetProcessListResponseProcessListProcessInfo) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetNice

`func (o *StatisticGetProcessListResponseProcessListProcessInfo) GetNice() int32`

GetNice returns the Nice field if non-nil, zero value otherwise.

### GetNiceOk

`func (o *StatisticGetProcessListResponseProcessListProcessInfo) GetNiceOk() (*int32, bool)`

GetNiceOk returns a tuple with the Nice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNice

`func (o *StatisticGetProcessListResponseProcessListProcessInfo) SetNice(v int32)`

SetNice sets Nice field to given value.

### HasNice

`func (o *StatisticGetProcessListResponseProcessListProcessInfo) HasNice() bool`

HasNice returns a boolean if a field has been set.

### GetVirt

`func (o *StatisticGetProcessListResponseProcessListProcessInfo) GetVirt() string`

GetVirt returns the Virt field if non-nil, zero value otherwise.

### GetVirtOk

`func (o *StatisticGetProcessListResponseProcessListProcessInfo) GetVirtOk() (*string, bool)`

GetVirtOk returns a tuple with the Virt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirt

`func (o *StatisticGetProcessListResponseProcessListProcessInfo) SetVirt(v string)`

SetVirt sets Virt field to given value.

### HasVirt

`func (o *StatisticGetProcessListResponseProcessListProcessInfo) HasVirt() bool`

HasVirt returns a boolean if a field has been set.

### GetRss

`func (o *StatisticGetProcessListResponseProcessListProcessInfo) GetRss() string`

GetRss returns the Rss field if non-nil, zero value otherwise.

### GetRssOk

`func (o *StatisticGetProcessListResponseProcessListProcessInfo) GetRssOk() (*string, bool)`

GetRssOk returns a tuple with the Rss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRss

`func (o *StatisticGetProcessListResponseProcessListProcessInfo) SetRss(v string)`

SetRss sets Rss field to given value.

### HasRss

`func (o *StatisticGetProcessListResponseProcessListProcessInfo) HasRss() bool`

HasRss returns a boolean if a field has been set.

### GetState

`func (o *StatisticGetProcessListResponseProcessListProcessInfo) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StatisticGetProcessListResponseProcessListProcessInfo) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StatisticGetProcessListResponseProcessListProcessInfo) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *StatisticGetProcessListResponseProcessListProcessInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCpuPercent

`func (o *StatisticGetProcessListResponseProcessListProcessInfo) GetCpuPercent() float32`

GetCpuPercent returns the CpuPercent field if non-nil, zero value otherwise.

### GetCpuPercentOk

`func (o *StatisticGetProcessListResponseProcessListProcessInfo) GetCpuPercentOk() (*float32, bool)`

GetCpuPercentOk returns a tuple with the CpuPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuPercent

`func (o *StatisticGetProcessListResponseProcessListProcessInfo) SetCpuPercent(v float32)`

SetCpuPercent sets CpuPercent field to given value.

### HasCpuPercent

`func (o *StatisticGetProcessListResponseProcessListProcessInfo) HasCpuPercent() bool`

HasCpuPercent returns a boolean if a field has been set.

### GetMemPercent

`func (o *StatisticGetProcessListResponseProcessListProcessInfo) GetMemPercent() float32`

GetMemPercent returns the MemPercent field if non-nil, zero value otherwise.

### GetMemPercentOk

`func (o *StatisticGetProcessListResponseProcessListProcessInfo) GetMemPercentOk() (*float32, bool)`

GetMemPercentOk returns a tuple with the MemPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemPercent

`func (o *StatisticGetProcessListResponseProcessListProcessInfo) SetMemPercent(v float32)`

SetMemPercent sets MemPercent field to given value.

### HasMemPercent

`func (o *StatisticGetProcessListResponseProcessListProcessInfo) HasMemPercent() bool`

HasMemPercent returns a boolean if a field has been set.

### GetTimeRunning

`func (o *StatisticGetProcessListResponseProcessListProcessInfo) GetTimeRunning() float32`

GetTimeRunning returns the TimeRunning field if non-nil, zero value otherwise.

### GetTimeRunningOk

`func (o *StatisticGetProcessListResponseProcessListProcessInfo) GetTimeRunningOk() (*float32, bool)`

GetTimeRunningOk returns a tuple with the TimeRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRunning

`func (o *StatisticGetProcessListResponseProcessListProcessInfo) SetTimeRunning(v float32)`

SetTimeRunning sets TimeRunning field to given value.

### HasTimeRunning

`func (o *StatisticGetProcessListResponseProcessListProcessInfo) HasTimeRunning() bool`

HasTimeRunning returns a boolean if a field has been set.

### GetName

`func (o *StatisticGetProcessListResponseProcessListProcessInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StatisticGetProcessListResponseProcessListProcessInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StatisticGetProcessListResponseProcessListProcessInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StatisticGetProcessListResponseProcessListProcessInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUser

`func (o *StatisticGetProcessListResponseProcessListProcessInfo) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *StatisticGetProcessListResponseProcessListProcessInfo) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *StatisticGetProcessListResponseProcessListProcessInfo) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *StatisticGetProcessListResponseProcessListProcessInfo) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


