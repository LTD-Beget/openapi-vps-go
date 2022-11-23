# \StatisticServiceApi

All URIs are relative to *https://api.beget.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StatisticServiceGetCpu**](StatisticServiceApi.md#StatisticServiceGetCpu) | **Get** /v1/vps/statistic/cpu/{id} | 
[**StatisticServiceGetCpuDetails**](StatisticServiceApi.md#StatisticServiceGetCpuDetails) | **Get** /v1/vps/statistic/cpu-details/{id} | 
[**StatisticServiceGetDisk**](StatisticServiceApi.md#StatisticServiceGetDisk) | **Get** /v1/vps/statistic/disk/{id} | 
[**StatisticServiceGetDiskUsage**](StatisticServiceApi.md#StatisticServiceGetDiskUsage) | **Get** /v1/vps/statistic/disk-usage/{id} | 
[**StatisticServiceGetLoadAverage**](StatisticServiceApi.md#StatisticServiceGetLoadAverage) | **Get** /v1/vps/statistic/load-average/{id} | 
[**StatisticServiceGetMemory**](StatisticServiceApi.md#StatisticServiceGetMemory) | **Get** /v1/vps/statistic/memory/{id} | 
[**StatisticServiceGetNetwork**](StatisticServiceApi.md#StatisticServiceGetNetwork) | **Get** /v1/vps/statistic/network/{id} | 
[**StatisticServiceGetProcessList**](StatisticServiceApi.md#StatisticServiceGetProcessList) | **Get** /v1/vps/statistic/processes/{id} | 



## StatisticServiceGetCpu

> StatisticGetCpuResponse StatisticServiceGetCpu(ctx, id).Period(period).Execute()



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
    id := "id_example" // string | 
    period := "period_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StatisticServiceApi.StatisticServiceGetCpu(context.Background(), id).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatisticServiceApi.StatisticServiceGetCpu``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StatisticServiceGetCpu`: StatisticGetCpuResponse
    fmt.Fprintf(os.Stdout, "Response from `StatisticServiceApi.StatisticServiceGetCpu`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStatisticServiceGetCpuRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** |  | 

### Return type

[**StatisticGetCpuResponse**](StatisticGetCpuResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StatisticServiceGetCpuDetails

> StatisticGetCpuDetailsResponse StatisticServiceGetCpuDetails(ctx, id).Period(period).Execute()



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
    id := "id_example" // string | 
    period := "period_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StatisticServiceApi.StatisticServiceGetCpuDetails(context.Background(), id).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatisticServiceApi.StatisticServiceGetCpuDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StatisticServiceGetCpuDetails`: StatisticGetCpuDetailsResponse
    fmt.Fprintf(os.Stdout, "Response from `StatisticServiceApi.StatisticServiceGetCpuDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStatisticServiceGetCpuDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** |  | 

### Return type

[**StatisticGetCpuDetailsResponse**](StatisticGetCpuDetailsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StatisticServiceGetDisk

> StatisticGetDiskResponse StatisticServiceGetDisk(ctx, id).Period(period).Execute()



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
    id := "id_example" // string | 
    period := "period_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StatisticServiceApi.StatisticServiceGetDisk(context.Background(), id).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatisticServiceApi.StatisticServiceGetDisk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StatisticServiceGetDisk`: StatisticGetDiskResponse
    fmt.Fprintf(os.Stdout, "Response from `StatisticServiceApi.StatisticServiceGetDisk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStatisticServiceGetDiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** |  | 

### Return type

[**StatisticGetDiskResponse**](StatisticGetDiskResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StatisticServiceGetDiskUsage

> StatisticGetDiskUsageResponse StatisticServiceGetDiskUsage(ctx, id).Period(period).Execute()



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
    id := "id_example" // string | 
    period := "period_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StatisticServiceApi.StatisticServiceGetDiskUsage(context.Background(), id).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatisticServiceApi.StatisticServiceGetDiskUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StatisticServiceGetDiskUsage`: StatisticGetDiskUsageResponse
    fmt.Fprintf(os.Stdout, "Response from `StatisticServiceApi.StatisticServiceGetDiskUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStatisticServiceGetDiskUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** |  | 

### Return type

[**StatisticGetDiskUsageResponse**](StatisticGetDiskUsageResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StatisticServiceGetLoadAverage

> StatisticGetLoadAverageResponse StatisticServiceGetLoadAverage(ctx, id).Period(period).Execute()



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
    id := "id_example" // string | 
    period := "period_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StatisticServiceApi.StatisticServiceGetLoadAverage(context.Background(), id).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatisticServiceApi.StatisticServiceGetLoadAverage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StatisticServiceGetLoadAverage`: StatisticGetLoadAverageResponse
    fmt.Fprintf(os.Stdout, "Response from `StatisticServiceApi.StatisticServiceGetLoadAverage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStatisticServiceGetLoadAverageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** |  | 

### Return type

[**StatisticGetLoadAverageResponse**](StatisticGetLoadAverageResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StatisticServiceGetMemory

> StatisticGetMemoryResponse StatisticServiceGetMemory(ctx, id).Period(period).Execute()



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
    id := "id_example" // string | 
    period := "period_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StatisticServiceApi.StatisticServiceGetMemory(context.Background(), id).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatisticServiceApi.StatisticServiceGetMemory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StatisticServiceGetMemory`: StatisticGetMemoryResponse
    fmt.Fprintf(os.Stdout, "Response from `StatisticServiceApi.StatisticServiceGetMemory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStatisticServiceGetMemoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** |  | 

### Return type

[**StatisticGetMemoryResponse**](StatisticGetMemoryResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StatisticServiceGetNetwork

> StatisticGetNetworkResponse StatisticServiceGetNetwork(ctx, id).Period(period).Execute()



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
    id := "id_example" // string | 
    period := "period_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StatisticServiceApi.StatisticServiceGetNetwork(context.Background(), id).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatisticServiceApi.StatisticServiceGetNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StatisticServiceGetNetwork`: StatisticGetNetworkResponse
    fmt.Fprintf(os.Stdout, "Response from `StatisticServiceApi.StatisticServiceGetNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStatisticServiceGetNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** |  | 

### Return type

[**StatisticGetNetworkResponse**](StatisticGetNetworkResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StatisticServiceGetProcessList

> StatisticGetProcessListResponse StatisticServiceGetProcessList(ctx, id).Execute()



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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StatisticServiceApi.StatisticServiceGetProcessList(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatisticServiceApi.StatisticServiceGetProcessList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StatisticServiceGetProcessList`: StatisticGetProcessListResponse
    fmt.Fprintf(os.Stdout, "Response from `StatisticServiceApi.StatisticServiceGetProcessList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStatisticServiceGetProcessListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StatisticGetProcessListResponse**](StatisticGetProcessListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

