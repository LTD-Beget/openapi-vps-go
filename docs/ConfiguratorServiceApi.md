# \ConfiguratorServiceApi

All URIs are relative to *https://api.beget.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfiguratorServiceGetCalculation**](ConfiguratorServiceApi.md#ConfiguratorServiceGetCalculation) | **Get** /v1/vps/configurator/calculation | 
[**ConfiguratorServiceGetConfiguratorInfo**](ConfiguratorServiceApi.md#ConfiguratorServiceGetConfiguratorInfo) | **Get** /v1/vps/configurator/info | 



## ConfiguratorServiceGetCalculation

> ConfiguratorGetCalculationResponse ConfiguratorServiceGetCalculation(ctx).ParamsCpuCount(paramsCpuCount).ParamsDiskSize(paramsDiskSize).ParamsMemory(paramsMemory).Region(region).VpsId(vpsId).SoftwareId(softwareId).SnapshotId(snapshotId).ImageId(imageId).ConfigurationGroup(configurationGroup).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/LTD-Beget/openapi-vps-go"
)

func main() {
    paramsCpuCount := int32(56) // int32 |  (optional)
    paramsDiskSize := int32(56) // int32 |  (optional)
    paramsMemory := int32(56) // int32 |  (optional)
    region := "region_example" // string |  (optional)
    vpsId := "vpsId_example" // string |  (optional)
    softwareId := int32(56) // int32 |  (optional)
    snapshotId := "snapshotId_example" // string |  (optional)
    imageId := "imageId_example" // string |  (optional)
    configurationGroup := "configurationGroup_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfiguratorServiceApi.ConfiguratorServiceGetCalculation(context.Background()).ParamsCpuCount(paramsCpuCount).ParamsDiskSize(paramsDiskSize).ParamsMemory(paramsMemory).Region(region).VpsId(vpsId).SoftwareId(softwareId).SnapshotId(snapshotId).ImageId(imageId).ConfigurationGroup(configurationGroup).Execute()
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
 **region** | **string** |  | 
 **vpsId** | **string** |  | 
 **softwareId** | **int32** |  | 
 **snapshotId** | **string** |  | 
 **imageId** | **string** |  | 
 **configurationGroup** | **string** |  | 

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

> ConfiguratorGetConfiguratorInfoResponse ConfiguratorServiceGetConfiguratorInfo(ctx).Region(region).ConfigurationGroup(configurationGroup).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/LTD-Beget/openapi-vps-go"
)

func main() {
    region := "region_example" // string |  (optional)
    configurationGroup := "configurationGroup_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfiguratorServiceApi.ConfiguratorServiceGetConfiguratorInfo(context.Background()).Region(region).ConfigurationGroup(configurationGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfiguratorServiceApi.ConfiguratorServiceGetConfiguratorInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfiguratorServiceGetConfiguratorInfo`: ConfiguratorGetConfiguratorInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `ConfiguratorServiceApi.ConfiguratorServiceGetConfiguratorInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConfiguratorServiceGetConfiguratorInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **region** | **string** |  | 
 **configurationGroup** | **string** |  | 

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

