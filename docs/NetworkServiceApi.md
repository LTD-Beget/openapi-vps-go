# \NetworkServiceApi

All URIs are relative to *https://api.beget.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NetworkServiceCreatePrivateNetwork**](NetworkServiceApi.md#NetworkServiceCreatePrivateNetwork) | **Post** /v1/vps/private-network | 
[**NetworkServiceGetNetworkInfo**](NetworkServiceApi.md#NetworkServiceGetNetworkInfo) | **Get** /v1/vps/network | 
[**NetworkServiceOrderIpAddress**](NetworkServiceApi.md#NetworkServiceOrderIpAddress) | **Post** /v1/vps/network | 
[**NetworkServiceRemoveIpAddress**](NetworkServiceApi.md#NetworkServiceRemoveIpAddress) | **Delete** /v1/vps/network/{ip_address} | 
[**NetworkServiceSuggestPrivateAddress**](NetworkServiceApi.md#NetworkServiceSuggestPrivateAddress) | **Post** /v1/vps/private-network/{network_id}/suggested-address | 



## NetworkServiceCreatePrivateNetwork

> NetworkCreatePrivateNetworkResponse NetworkServiceCreatePrivateNetwork(ctx).NetworkCreatePrivateNetworkRequest(networkCreatePrivateNetworkRequest).Execute()



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
    networkCreatePrivateNetworkRequest := *openapiclient.NewNetworkCreatePrivateNetworkRequest() // NetworkCreatePrivateNetworkRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkServiceApi.NetworkServiceCreatePrivateNetwork(context.Background()).NetworkCreatePrivateNetworkRequest(networkCreatePrivateNetworkRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkServiceApi.NetworkServiceCreatePrivateNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkServiceCreatePrivateNetwork`: NetworkCreatePrivateNetworkResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkServiceApi.NetworkServiceCreatePrivateNetwork`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNetworkServiceCreatePrivateNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkCreatePrivateNetworkRequest** | [**NetworkCreatePrivateNetworkRequest**](NetworkCreatePrivateNetworkRequest.md) |  | 

### Return type

[**NetworkCreatePrivateNetworkResponse**](NetworkCreatePrivateNetworkResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkServiceGetNetworkInfo

> NetworkGetNetworkInfoResponse NetworkServiceGetNetworkInfo(ctx).Execute()



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
    resp, r, err := apiClient.NetworkServiceApi.NetworkServiceGetNetworkInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkServiceApi.NetworkServiceGetNetworkInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkServiceGetNetworkInfo`: NetworkGetNetworkInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkServiceApi.NetworkServiceGetNetworkInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkServiceGetNetworkInfoRequest struct via the builder pattern


### Return type

[**NetworkGetNetworkInfoResponse**](NetworkGetNetworkInfoResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkServiceOrderIpAddress

> NetworkOrderIpAddressResponse NetworkServiceOrderIpAddress(ctx).NetworkOrderIpAddressRequest(networkOrderIpAddressRequest).Execute()



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
    networkOrderIpAddressRequest := *openapiclient.NewNetworkOrderIpAddressRequest() // NetworkOrderIpAddressRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkServiceApi.NetworkServiceOrderIpAddress(context.Background()).NetworkOrderIpAddressRequest(networkOrderIpAddressRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkServiceApi.NetworkServiceOrderIpAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkServiceOrderIpAddress`: NetworkOrderIpAddressResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkServiceApi.NetworkServiceOrderIpAddress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNetworkServiceOrderIpAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkOrderIpAddressRequest** | [**NetworkOrderIpAddressRequest**](NetworkOrderIpAddressRequest.md) |  | 

### Return type

[**NetworkOrderIpAddressResponse**](NetworkOrderIpAddressResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkServiceRemoveIpAddress

> NetworkRemoveIpAddressResponse NetworkServiceRemoveIpAddress(ctx, ipAddress).Execute()



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
    ipAddress := "ipAddress_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkServiceApi.NetworkServiceRemoveIpAddress(context.Background(), ipAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkServiceApi.NetworkServiceRemoveIpAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkServiceRemoveIpAddress`: NetworkRemoveIpAddressResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkServiceApi.NetworkServiceRemoveIpAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ipAddress** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkServiceRemoveIpAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkRemoveIpAddressResponse**](NetworkRemoveIpAddressResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkServiceSuggestPrivateAddress

> NetworkSuggestPrivateAddressResponse NetworkServiceSuggestPrivateAddress(ctx, networkId).NetworkSuggestPrivateAddressRequest(networkSuggestPrivateAddressRequest).Execute()



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
    networkId := "networkId_example" // string | 
    networkSuggestPrivateAddressRequest := *openapiclient.NewNetworkSuggestPrivateAddressRequest() // NetworkSuggestPrivateAddressRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkServiceApi.NetworkServiceSuggestPrivateAddress(context.Background(), networkId).NetworkSuggestPrivateAddressRequest(networkSuggestPrivateAddressRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkServiceApi.NetworkServiceSuggestPrivateAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkServiceSuggestPrivateAddress`: NetworkSuggestPrivateAddressResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkServiceApi.NetworkServiceSuggestPrivateAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkServiceSuggestPrivateAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkSuggestPrivateAddressRequest** | [**NetworkSuggestPrivateAddressRequest**](NetworkSuggestPrivateAddressRequest.md) |  | 

### Return type

[**NetworkSuggestPrivateAddressResponse**](NetworkSuggestPrivateAddressResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

