# \ConfiguratorServiceApi

All URIs are relative to *https://api.beget.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfiguratorServiceGetCalculation**](ConfiguratorServiceApi.md#ConfiguratorServiceGetCalculation) | **Get** /v1/vps/configurator/calculation | 
[**ConfiguratorServiceGetConfiguratorInfo**](ConfiguratorServiceApi.md#ConfiguratorServiceGetConfiguratorInfo) | **Get** /v1/vps/configurator/info | 



## ConfiguratorServiceGetCalculation

> ConfiguratorGetCalculationResponse ConfiguratorServiceGetCalculation(ctx).ParamsCpuCount(paramsCpuCount).ParamsDiskSize(paramsDiskSize).ParamsMemory(paramsMemory).VpsId(vpsId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    paramsCpuCount := int32(56) // int32 |  (optional)
    paramsDiskSize := int32(56) // int32 |  (optional)
    paramsMemory := int32(56) // int32 |  (optional)
    vpsId := "vpsId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfiguratorServiceApi.ConfiguratorServiceGetCalculation(context.Background()).ParamsCpuCount(paramsCpuCount).ParamsDiskSize(paramsDiskSize).ParamsMemory(paramsMemory).VpsId(vpsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfiguratorServiceApi.ConfiguratorServiceGetCalculation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfiguratorServiceGetCalculation`: ConfiguratorGetCalculationResponse
    fmt.Fprintf(os.Stdout, "Response from `ConfiguratorServiceApi.ConfiguratorServiceGetCalculation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConfiguratorServiceGetCalculationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paramsCpuCount** | **int32** |  | 
 **paramsDiskSize** | **int32** |  | 
 **paramsMemory** | **int32** |  | 
 **vpsId** | **string** |  | 

### Return type

[**ConfiguratorGetCalculationResponse**](ConfiguratorGetCalculationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfiguratorServiceGetConfiguratorInfo

> ConfiguratorGetConfiguratorInfoResponse ConfiguratorServiceGetConfiguratorInfo(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfiguratorServiceApi.ConfiguratorServiceGetConfiguratorInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfiguratorServiceApi.ConfiguratorServiceGetConfiguratorInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfiguratorServiceGetConfiguratorInfo`: ConfiguratorGetConfiguratorInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `ConfiguratorServiceApi.ConfiguratorServiceGetConfiguratorInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiConfiguratorServiceGetConfiguratorInfoRequest struct via the builder pattern


### Return type

[**ConfiguratorGetConfiguratorInfoResponse**](ConfiguratorGetConfiguratorInfoResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

