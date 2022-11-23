# StatisticGetDiskResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataRead** | Pointer to [**StatisticSeriesData**](StatisticSeriesData.md) |  | [optional] 
**DataWrite** | Pointer to [**StatisticSeriesData**](StatisticSeriesData.md) |  | [optional] 

## Methods

### NewStatisticGetDiskResponse

`func NewStatisticGetDiskResponse() *StatisticGetDiskResponse`

NewStatisticGetDiskResponse instantiates a new StatisticGetDiskResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatisticGetDiskResponseWithDefaults

`func NewStatisticGetDiskResponseWithDefaults() *StatisticGetDiskResponse`

NewStatisticGetDiskResponseWithDefaults instantiates a new StatisticGetDiskResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataRead

`func (o *StatisticGetDiskResponse) GetDataRead() StatisticSeriesData`

GetDataRead returns the DataRead field if non-nil, zero value otherwise.

### GetDataReadOk

`func (o *StatisticGetDiskResponse) GetDataReadOk() (*StatisticSeriesData, bool)`

GetDataReadOk returns a tuple with the DataRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRead

`func (o *StatisticGetDiskResponse) SetDataRead(v StatisticSeriesData)`

SetDataRead sets DataRead field to given value.

### HasDataRead

`func (o *StatisticGetDiskResponse) HasDataRead() bool`

HasDataRead returns a boolean if a field has been set.

### GetDataWrite

`func (o *StatisticGetDiskResponse) GetDataWrite() StatisticSeriesData`

GetDataWrite returns the DataWrite field if non-nil, zero value otherwise.

### GetDataWriteOk

`func (o *StatisticGetDiskResponse) GetDataWriteOk() (*StatisticSeriesData, bool)`

GetDataWriteOk returns a tuple with the DataWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataWrite

`func (o *StatisticGetDiskResponse) SetDataWrite(v StatisticSeriesData)`

SetDataWrite sets DataWrite field to given value.

### HasDataWrite

`func (o *StatisticGetDiskResponse) HasDataWrite() bool`

HasDataWrite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


