# StatisticGetProcessListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Processes** | Pointer to [**StatisticGetProcessListResponseProcessList**](StatisticGetProcessListResponseProcessList.md) |  | [optional] 
**Error** | Pointer to [**StatisticGetProcessListResponseError**](StatisticGetProcessListResponseError.md) |  | [optional] 

## Methods

### NewStatisticGetProcessListResponse

`func NewStatisticGetProcessListResponse() *StatisticGetProcessListResponse`

NewStatisticGetProcessListResponse instantiates a new StatisticGetProcessListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatisticGetProcessListResponseWithDefaults

`func NewStatisticGetProcessListResponseWithDefaults() *StatisticGetProcessListResponse`

NewStatisticGetProcessListResponseWithDefaults instantiates a new StatisticGetProcessListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcesses

`func (o *StatisticGetProcessListResponse) GetProcesses() StatisticGetProcessListResponseProcessList`

GetProcesses returns the Processes field if non-nil, zero value otherwise.

### GetProcessesOk

`func (o *StatisticGetProcessListResponse) GetProcessesOk() (*StatisticGetProcessListResponseProcessList, bool)`

GetProcessesOk returns a tuple with the Processes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcesses

`func (o *StatisticGetProcessListResponse) SetProcesses(v StatisticGetProcessListResponseProcessList)`

SetProcesses sets Processes field to given value.

### HasProcesses

`func (o *StatisticGetProcessListResponse) HasProcesses() bool`

HasProcesses returns a boolean if a field has been set.

### GetError

`func (o *StatisticGetProcessListResponse) GetError() StatisticGetProcessListResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *StatisticGetProcessListResponse) GetErrorOk() (*StatisticGetProcessListResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *StatisticGetProcessListResponse) SetError(v StatisticGetProcessListResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *StatisticGetProcessListResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


