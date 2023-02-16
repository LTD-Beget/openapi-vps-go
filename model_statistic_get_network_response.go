/*
API Облачных серверов

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.4.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package begetOpenapiVps

import (
	"encoding/json"
)

// StatisticGetNetworkResponse struct for StatisticGetNetworkResponse
type StatisticGetNetworkResponse struct {
	DataRx *StatisticSeriesData `json:"data_rx,omitempty"`
	DataTx *StatisticSeriesData `json:"data_tx,omitempty"`
	DataRxPrivate *StatisticSeriesData `json:"data_rx_private,omitempty"`
	DataTxPrivate *StatisticSeriesData `json:"data_tx_private,omitempty"`
}

// NewStatisticGetNetworkResponse instantiates a new StatisticGetNetworkResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatisticGetNetworkResponse() *StatisticGetNetworkResponse {
	this := StatisticGetNetworkResponse{}
	return &this
}

// NewStatisticGetNetworkResponseWithDefaults instantiates a new StatisticGetNetworkResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatisticGetNetworkResponseWithDefaults() *StatisticGetNetworkResponse {
	this := StatisticGetNetworkResponse{}
	return &this
}

// GetDataRx returns the DataRx field value if set, zero value otherwise.
func (o *StatisticGetNetworkResponse) GetDataRx() StatisticSeriesData {
	if o == nil || isNil(o.DataRx) {
		var ret StatisticSeriesData
		return ret
	}
	return *o.DataRx
}

// GetDataRxOk returns a tuple with the DataRx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticGetNetworkResponse) GetDataRxOk() (*StatisticSeriesData, bool) {
	if o == nil || isNil(o.DataRx) {
    return nil, false
	}
	return o.DataRx, true
}

// HasDataRx returns a boolean if a field has been set.
func (o *StatisticGetNetworkResponse) HasDataRx() bool {
	if o != nil && !isNil(o.DataRx) {
		return true
	}

	return false
}

// SetDataRx gets a reference to the given StatisticSeriesData and assigns it to the DataRx field.
func (o *StatisticGetNetworkResponse) SetDataRx(v StatisticSeriesData) {
	o.DataRx = &v
}

// GetDataTx returns the DataTx field value if set, zero value otherwise.
func (o *StatisticGetNetworkResponse) GetDataTx() StatisticSeriesData {
	if o == nil || isNil(o.DataTx) {
		var ret StatisticSeriesData
		return ret
	}
	return *o.DataTx
}

// GetDataTxOk returns a tuple with the DataTx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticGetNetworkResponse) GetDataTxOk() (*StatisticSeriesData, bool) {
	if o == nil || isNil(o.DataTx) {
    return nil, false
	}
	return o.DataTx, true
}

// HasDataTx returns a boolean if a field has been set.
func (o *StatisticGetNetworkResponse) HasDataTx() bool {
	if o != nil && !isNil(o.DataTx) {
		return true
	}

	return false
}

// SetDataTx gets a reference to the given StatisticSeriesData and assigns it to the DataTx field.
func (o *StatisticGetNetworkResponse) SetDataTx(v StatisticSeriesData) {
	o.DataTx = &v
}

// GetDataRxPrivate returns the DataRxPrivate field value if set, zero value otherwise.
func (o *StatisticGetNetworkResponse) GetDataRxPrivate() StatisticSeriesData {
	if o == nil || isNil(o.DataRxPrivate) {
		var ret StatisticSeriesData
		return ret
	}
	return *o.DataRxPrivate
}

// GetDataRxPrivateOk returns a tuple with the DataRxPrivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticGetNetworkResponse) GetDataRxPrivateOk() (*StatisticSeriesData, bool) {
	if o == nil || isNil(o.DataRxPrivate) {
    return nil, false
	}
	return o.DataRxPrivate, true
}

// HasDataRxPrivate returns a boolean if a field has been set.
func (o *StatisticGetNetworkResponse) HasDataRxPrivate() bool {
	if o != nil && !isNil(o.DataRxPrivate) {
		return true
	}

	return false
}

// SetDataRxPrivate gets a reference to the given StatisticSeriesData and assigns it to the DataRxPrivate field.
func (o *StatisticGetNetworkResponse) SetDataRxPrivate(v StatisticSeriesData) {
	o.DataRxPrivate = &v
}

// GetDataTxPrivate returns the DataTxPrivate field value if set, zero value otherwise.
func (o *StatisticGetNetworkResponse) GetDataTxPrivate() StatisticSeriesData {
	if o == nil || isNil(o.DataTxPrivate) {
		var ret StatisticSeriesData
		return ret
	}
	return *o.DataTxPrivate
}

// GetDataTxPrivateOk returns a tuple with the DataTxPrivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticGetNetworkResponse) GetDataTxPrivateOk() (*StatisticSeriesData, bool) {
	if o == nil || isNil(o.DataTxPrivate) {
    return nil, false
	}
	return o.DataTxPrivate, true
}

// HasDataTxPrivate returns a boolean if a field has been set.
func (o *StatisticGetNetworkResponse) HasDataTxPrivate() bool {
	if o != nil && !isNil(o.DataTxPrivate) {
		return true
	}

	return false
}

// SetDataTxPrivate gets a reference to the given StatisticSeriesData and assigns it to the DataTxPrivate field.
func (o *StatisticGetNetworkResponse) SetDataTxPrivate(v StatisticSeriesData) {
	o.DataTxPrivate = &v
}

func (o StatisticGetNetworkResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DataRx) {
		toSerialize["data_rx"] = o.DataRx
	}
	if !isNil(o.DataTx) {
		toSerialize["data_tx"] = o.DataTx
	}
	if !isNil(o.DataRxPrivate) {
		toSerialize["data_rx_private"] = o.DataRxPrivate
	}
	if !isNil(o.DataTxPrivate) {
		toSerialize["data_tx_private"] = o.DataTxPrivate
	}
	return json.Marshal(toSerialize)
}

type NullableStatisticGetNetworkResponse struct {
	value *StatisticGetNetworkResponse
	isSet bool
}

func (v NullableStatisticGetNetworkResponse) Get() *StatisticGetNetworkResponse {
	return v.value
}

func (v *NullableStatisticGetNetworkResponse) Set(val *StatisticGetNetworkResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStatisticGetNetworkResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStatisticGetNetworkResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatisticGetNetworkResponse(val *StatisticGetNetworkResponse) *NullableStatisticGetNetworkResponse {
	return &NullableStatisticGetNetworkResponse{value: val, isSet: true}
}

func (v NullableStatisticGetNetworkResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatisticGetNetworkResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


