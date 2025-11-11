# \BackupServiceApi

All URIs are relative to *https://api.beget.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BackupServiceGetAvailableCopies**](BackupServiceApi.md#BackupServiceGetAvailableCopies) | **Get** /v1/vps/backup | 
[**BackupServiceGetBackupFileList**](BackupServiceApi.md#BackupServiceGetBackupFileList) | **Get** /v1/vps/{id}/backup/{copy_id} | 
[**BackupServiceGetOrders**](BackupServiceApi.md#BackupServiceGetOrders) | **Get** /v1/vps/backup/orders | 
[**BackupServiceRestoreFile**](BackupServiceApi.md#BackupServiceRestoreFile) | **Post** /v1/vps/{id}/backup/{copy_id}/file | 
[**BackupServiceRestoreServer**](BackupServiceApi.md#BackupServiceRestoreServer) | **Post** /v1/vps/{id}/backup/{copy_id}/server | 



## BackupServiceGetAvailableCopies

> BackupGetAvailableCopiesResponse BackupServiceGetAvailableCopies(ctx).Filter(filter).Execute()



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
    filter := "filter_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupServiceApi.BackupServiceGetAvailableCopies(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupServiceApi.BackupServiceGetAvailableCopies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupServiceGetAvailableCopies`: BackupGetAvailableCopiesResponse
    fmt.Fprintf(os.Stdout, "Response from `BackupServiceApi.BackupServiceGetAvailableCopies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackupServiceGetAvailableCopiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** |  | 

### Return type

[**BackupGetAvailableCopiesResponse**](BackupGetAvailableCopiesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupServiceGetBackupFileList

> BackupGetBackupFileListResponse BackupServiceGetBackupFileList(ctx, id, copyId).Path(path).Execute()



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
    copyId := "copyId_example" // string | 
    path := "path_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupServiceApi.BackupServiceGetBackupFileList(context.Background(), id, copyId).Path(path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupServiceApi.BackupServiceGetBackupFileList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupServiceGetBackupFileList`: BackupGetBackupFileListResponse
    fmt.Fprintf(os.Stdout, "Response from `BackupServiceApi.BackupServiceGetBackupFileList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**copyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupServiceGetBackupFileListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **path** | **string** |  | 

### Return type

[**BackupGetBackupFileListResponse**](BackupGetBackupFileListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupServiceGetOrders

> BackupGetOrdersResponse BackupServiceGetOrders(ctx).Limit(limit).Offset(offset).Execute()



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
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupServiceApi.BackupServiceGetOrders(context.Background()).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupServiceApi.BackupServiceGetOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupServiceGetOrders`: BackupGetOrdersResponse
    fmt.Fprintf(os.Stdout, "Response from `BackupServiceApi.BackupServiceGetOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackupServiceGetOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 

### Return type

[**BackupGetOrdersResponse**](BackupGetOrdersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupServiceRestoreFile

> BackupRestoreFileResponse BackupServiceRestoreFile(ctx, id, copyId).BackupRestoreFileRequest(backupRestoreFileRequest).Execute()



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
    copyId := "copyId_example" // string | 
    backupRestoreFileRequest := *openapiclient.NewBackupRestoreFileRequest() // BackupRestoreFileRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupServiceApi.BackupServiceRestoreFile(context.Background(), id, copyId).BackupRestoreFileRequest(backupRestoreFileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupServiceApi.BackupServiceRestoreFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupServiceRestoreFile`: BackupRestoreFileResponse
    fmt.Fprintf(os.Stdout, "Response from `BackupServiceApi.BackupServiceRestoreFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**copyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupServiceRestoreFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **backupRestoreFileRequest** | [**BackupRestoreFileRequest**](BackupRestoreFileRequest.md) |  | 

### Return type

[**BackupRestoreFileResponse**](BackupRestoreFileResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupServiceRestoreServer

> BackupRestoreServerResponse BackupServiceRestoreServer(ctx, id, copyId).BackupRestoreServerRequest(backupRestoreServerRequest).Execute()



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
    copyId := "copyId_example" // string | 
    backupRestoreServerRequest := *openapiclient.NewBackupRestoreServerRequest() // BackupRestoreServerRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupServiceApi.BackupServiceRestoreServer(context.Background(), id, copyId).BackupRestoreServerRequest(backupRestoreServerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupServiceApi.BackupServiceRestoreServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupServiceRestoreServer`: BackupRestoreServerResponse
    fmt.Fprintf(os.Stdout, "Response from `BackupServiceApi.BackupServiceRestoreServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**copyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackupServiceRestoreServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **backupRestoreServerRequest** | [**BackupRestoreServerRequest**](BackupRestoreServerRequest.md) |  | 

### Return type

[**BackupRestoreServerResponse**](BackupRestoreServerResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

