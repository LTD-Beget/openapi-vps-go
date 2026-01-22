# \SnapshotServiceApi

All URIs are relative to *https://api.beget.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SnapshotServiceCreate**](SnapshotServiceApi.md#SnapshotServiceCreate) | **Post** /v1/vps/snapshot | 
[**SnapshotServiceCreateCalculator**](SnapshotServiceApi.md#SnapshotServiceCreateCalculator) | **Post** /v1/vps/snapshot/calculator | 
[**SnapshotServiceEdit**](SnapshotServiceApi.md#SnapshotServiceEdit) | **Put** /v1/vps/snapshot/{id} | 
[**SnapshotServiceGetAll**](SnapshotServiceApi.md#SnapshotServiceGetAll) | **Get** /v1/vps/snapshot | 
[**SnapshotServiceGetAllRestores**](SnapshotServiceApi.md#SnapshotServiceGetAllRestores) | **Get** /v1/vps/snapshot/restore | 
[**SnapshotServiceRemove**](SnapshotServiceApi.md#SnapshotServiceRemove) | **Delete** /v1/vps/snapshot/{id} | 
[**SnapshotServiceRestore**](SnapshotServiceApi.md#SnapshotServiceRestore) | **Post** /v1/vps/snapshot/{id}/restore | 



## SnapshotServiceCreate

> SnapshotCreateResponse SnapshotServiceCreate(ctx).SnapshotCreateRequest(snapshotCreateRequest).Execute()



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
    snapshotCreateRequest := *openapiclient.NewSnapshotCreateRequest() // SnapshotCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SnapshotServiceApi.SnapshotServiceCreate(context.Background()).SnapshotCreateRequest(snapshotCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotServiceApi.SnapshotServiceCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SnapshotServiceCreate`: SnapshotCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `SnapshotServiceApi.SnapshotServiceCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSnapshotServiceCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **snapshotCreateRequest** | [**SnapshotCreateRequest**](SnapshotCreateRequest.md) |  | 

### Return type

[**SnapshotCreateResponse**](SnapshotCreateResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnapshotServiceCreateCalculator

> SnapshotCreateCalculatorResponse SnapshotServiceCreateCalculator(ctx).SnapshotCreateCalculatorRequest(snapshotCreateCalculatorRequest).Execute()



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
    snapshotCreateCalculatorRequest := *openapiclient.NewSnapshotCreateCalculatorRequest() // SnapshotCreateCalculatorRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SnapshotServiceApi.SnapshotServiceCreateCalculator(context.Background()).SnapshotCreateCalculatorRequest(snapshotCreateCalculatorRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotServiceApi.SnapshotServiceCreateCalculator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SnapshotServiceCreateCalculator`: SnapshotCreateCalculatorResponse
    fmt.Fprintf(os.Stdout, "Response from `SnapshotServiceApi.SnapshotServiceCreateCalculator`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSnapshotServiceCreateCalculatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **snapshotCreateCalculatorRequest** | [**SnapshotCreateCalculatorRequest**](SnapshotCreateCalculatorRequest.md) |  | 

### Return type

[**SnapshotCreateCalculatorResponse**](SnapshotCreateCalculatorResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnapshotServiceEdit

> SnapshotEditResponse SnapshotServiceEdit(ctx, id).SnapshotEditRequest(snapshotEditRequest).Execute()



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
    id := "id_example" // string | 
    snapshotEditRequest := *openapiclient.NewSnapshotEditRequest() // SnapshotEditRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SnapshotServiceApi.SnapshotServiceEdit(context.Background(), id).SnapshotEditRequest(snapshotEditRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotServiceApi.SnapshotServiceEdit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SnapshotServiceEdit`: SnapshotEditResponse
    fmt.Fprintf(os.Stdout, "Response from `SnapshotServiceApi.SnapshotServiceEdit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSnapshotServiceEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **snapshotEditRequest** | [**SnapshotEditRequest**](SnapshotEditRequest.md) |  | 

### Return type

[**SnapshotEditResponse**](SnapshotEditResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnapshotServiceGetAll

> SnapshotGetAllResponse SnapshotServiceGetAll(ctx).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SnapshotServiceApi.SnapshotServiceGetAll(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotServiceApi.SnapshotServiceGetAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SnapshotServiceGetAll`: SnapshotGetAllResponse
    fmt.Fprintf(os.Stdout, "Response from `SnapshotServiceApi.SnapshotServiceGetAll`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSnapshotServiceGetAllRequest struct via the builder pattern


### Return type

[**SnapshotGetAllResponse**](SnapshotGetAllResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnapshotServiceGetAllRestores

> SnapshotGetAllRestoresResponse SnapshotServiceGetAllRestores(ctx).Id(id).Execute()



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
    id := "id_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SnapshotServiceApi.SnapshotServiceGetAllRestores(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotServiceApi.SnapshotServiceGetAllRestores``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SnapshotServiceGetAllRestores`: SnapshotGetAllRestoresResponse
    fmt.Fprintf(os.Stdout, "Response from `SnapshotServiceApi.SnapshotServiceGetAllRestores`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSnapshotServiceGetAllRestoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 

### Return type

[**SnapshotGetAllRestoresResponse**](SnapshotGetAllRestoresResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnapshotServiceRemove

> SnapshotRemoveResponse SnapshotServiceRemove(ctx, id).Execute()



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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SnapshotServiceApi.SnapshotServiceRemove(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotServiceApi.SnapshotServiceRemove``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SnapshotServiceRemove`: SnapshotRemoveResponse
    fmt.Fprintf(os.Stdout, "Response from `SnapshotServiceApi.SnapshotServiceRemove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSnapshotServiceRemoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SnapshotRemoveResponse**](SnapshotRemoveResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnapshotServiceRestore

> SnapshotRestoreResponse SnapshotServiceRestore(ctx, id).SnapshotRestoreRequest(snapshotRestoreRequest).Execute()



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
    id := "id_example" // string | 
    snapshotRestoreRequest := *openapiclient.NewSnapshotRestoreRequest() // SnapshotRestoreRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SnapshotServiceApi.SnapshotServiceRestore(context.Background(), id).SnapshotRestoreRequest(snapshotRestoreRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotServiceApi.SnapshotServiceRestore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SnapshotServiceRestore`: SnapshotRestoreResponse
    fmt.Fprintf(os.Stdout, "Response from `SnapshotServiceApi.SnapshotServiceRestore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSnapshotServiceRestoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **snapshotRestoreRequest** | [**SnapshotRestoreRequest**](SnapshotRestoreRequest.md) |  | 

### Return type

[**SnapshotRestoreResponse**](SnapshotRestoreResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

