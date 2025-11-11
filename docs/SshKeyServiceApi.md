# \SshKeyServiceApi

All URIs are relative to *https://api.beget.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SshKeyServiceAdd**](SshKeyServiceApi.md#SshKeyServiceAdd) | **Post** /v1/vps/sshKey | 
[**SshKeyServiceGetAll**](SshKeyServiceApi.md#SshKeyServiceGetAll) | **Get** /v1/vps/sshKey | 
[**SshKeyServiceRemove**](SshKeyServiceApi.md#SshKeyServiceRemove) | **Delete** /v1/vps/sshKey/{id} | 
[**SshKeyServiceUpdate**](SshKeyServiceApi.md#SshKeyServiceUpdate) | **Put** /v1/vps/sshKey/{id} | 



## SshKeyServiceAdd

> SshKeyAddResponse SshKeyServiceAdd(ctx).SshKeyAddRequest(sshKeyAddRequest).Execute()



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
    sshKeyAddRequest := *openapiclient.NewSshKeyAddRequest() // SshKeyAddRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SshKeyServiceApi.SshKeyServiceAdd(context.Background()).SshKeyAddRequest(sshKeyAddRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SshKeyServiceApi.SshKeyServiceAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SshKeyServiceAdd`: SshKeyAddResponse
    fmt.Fprintf(os.Stdout, "Response from `SshKeyServiceApi.SshKeyServiceAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSshKeyServiceAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sshKeyAddRequest** | [**SshKeyAddRequest**](SshKeyAddRequest.md) |  | 

### Return type

[**SshKeyAddResponse**](SshKeyAddResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SshKeyServiceGetAll

> SshKeyGetAllResponse SshKeyServiceGetAll(ctx).Execute()



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
    resp, r, err := apiClient.SshKeyServiceApi.SshKeyServiceGetAll(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SshKeyServiceApi.SshKeyServiceGetAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SshKeyServiceGetAll`: SshKeyGetAllResponse
    fmt.Fprintf(os.Stdout, "Response from `SshKeyServiceApi.SshKeyServiceGetAll`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSshKeyServiceGetAllRequest struct via the builder pattern


### Return type

[**SshKeyGetAllResponse**](SshKeyGetAllResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SshKeyServiceRemove

> SshKeyRemoveResponse SshKeyServiceRemove(ctx, id).Force(force).Execute()



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
    id := int32(56) // int32 | 
    force := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SshKeyServiceApi.SshKeyServiceRemove(context.Background(), id).Force(force).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SshKeyServiceApi.SshKeyServiceRemove``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SshKeyServiceRemove`: SshKeyRemoveResponse
    fmt.Fprintf(os.Stdout, "Response from `SshKeyServiceApi.SshKeyServiceRemove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSshKeyServiceRemoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** |  | 

### Return type

[**SshKeyRemoveResponse**](SshKeyRemoveResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SshKeyServiceUpdate

> SshKeyUpdateResponse SshKeyServiceUpdate(ctx, id).SshKeyUpdateRequest(sshKeyUpdateRequest).Execute()



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
    id := int32(56) // int32 | 
    sshKeyUpdateRequest := *openapiclient.NewSshKeyUpdateRequest() // SshKeyUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SshKeyServiceApi.SshKeyServiceUpdate(context.Background(), id).SshKeyUpdateRequest(sshKeyUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SshKeyServiceApi.SshKeyServiceUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SshKeyServiceUpdate`: SshKeyUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `SshKeyServiceApi.SshKeyServiceUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSshKeyServiceUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sshKeyUpdateRequest** | [**SshKeyUpdateRequest**](SshKeyUpdateRequest.md) |  | 

### Return type

[**SshKeyUpdateResponse**](SshKeyUpdateResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

