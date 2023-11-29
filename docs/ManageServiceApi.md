# \ManageServiceApi

All URIs are relative to *https://api.beget.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ManageServiceAttachIpAddress**](ManageServiceApi.md#ManageServiceAttachIpAddress) | **Post** /v1/vps/{id}/network/{ip_address} | 
[**ManageServiceAttachSshKey**](ManageServiceApi.md#ManageServiceAttachSshKey) | **Post** /v1/vps/{id}/sshKey/{ssh_key_id} | 
[**ManageServiceAttachToPrivateNetwork**](ManageServiceApi.md#ManageServiceAttachToPrivateNetwork) | **Post** /v1/vps/{id}/private-network/{network_id} | 
[**ManageServiceChangeConfiguration**](ManageServiceApi.md#ManageServiceChangeConfiguration) | **Put** /v1/vps/server/{id}/configuration | 
[**ManageServiceChangeSshAccess**](ManageServiceApi.md#ManageServiceChangeSshAccess) | **Put** /v1/vps/{id}/ssh/access | 
[**ManageServiceCheckSoftwareRequirements**](ManageServiceApi.md#ManageServiceCheckSoftwareRequirements) | **Post** /v1/vps/software/requirements | 
[**ManageServiceCreateVps**](ManageServiceApi.md#ManageServiceCreateVps) | **Post** /v1/vps/server | 
[**ManageServiceDetachFromPrivateNetwork**](ManageServiceApi.md#ManageServiceDetachFromPrivateNetwork) | **Delete** /v1/vps/{id}/private-network/{network_id} | 
[**ManageServiceDetachIpAddress**](ManageServiceApi.md#ManageServiceDetachIpAddress) | **Delete** /v1/vps/network/detach/{ip_address} | 
[**ManageServiceDetachSshKey**](ManageServiceApi.md#ManageServiceDetachSshKey) | **Delete** /v1/vps/{id}/sshKey/{ssh_key_id} | 
[**ManageServiceDisablePostInstallAlert**](ManageServiceApi.md#ManageServiceDisablePostInstallAlert) | **Delete** /v1/vps/{id}/software/post-install-alert | 
[**ManageServiceGetAvailableConfiguration**](ManageServiceApi.md#ManageServiceGetAvailableConfiguration) | **Get** /v1/vps/configuration | 
[**ManageServiceGetFileManagerSettings**](ManageServiceApi.md#ManageServiceGetFileManagerSettings) | **Post** /v1/vps/{id}/fm | 
[**ManageServiceGetHistory**](ManageServiceApi.md#ManageServiceGetHistory) | **Get** /v1/vps/{id}/history | 
[**ManageServiceGetInfo**](ManageServiceApi.md#ManageServiceGetInfo) | **Get** /v1/vps/server/{id} | 
[**ManageServiceGetInstalledSoftware**](ManageServiceApi.md#ManageServiceGetInstalledSoftware) | **Get** /v1/vps/{id}/software | 
[**ManageServiceGetList**](ManageServiceApi.md#ManageServiceGetList) | **Get** /v1/vps/server/list | 
[**ManageServiceGetRegionList**](ManageServiceApi.md#ManageServiceGetRegionList) | **Get** /v1/vps/region | 
[**ManageServiceGetStatuses**](ManageServiceApi.md#ManageServiceGetStatuses) | **Get** /v1/vps/server/statuses | 
[**ManageServiceRebootVps**](ManageServiceApi.md#ManageServiceRebootVps) | **Post** /v1/vps/server/{id}/reboot | 
[**ManageServiceReinstall**](ManageServiceApi.md#ManageServiceReinstall) | **Post** /v1/vps/server/{id}/reinstall | 
[**ManageServiceRemoveVps**](ManageServiceApi.md#ManageServiceRemoveVps) | **Post** /v1/vps/server/{id}/remove | 
[**ManageServiceReserveVpsSubdomain**](ManageServiceApi.md#ManageServiceReserveVpsSubdomain) | **Get** /v1/vps/subdomain/reserve | 
[**ManageServiceResetPassword**](ManageServiceApi.md#ManageServiceResetPassword) | **Put** /v1/vps/{id}/password | 
[**ManageServiceResetVps**](ManageServiceApi.md#ManageServiceResetVps) | **Post** /v1/vps/server/{id}/reset | 
[**ManageServiceStartRescue**](ManageServiceApi.md#ManageServiceStartRescue) | **Post** /v1/vps/server/{id}/rescue | 
[**ManageServiceStartVps**](ManageServiceApi.md#ManageServiceStartVps) | **Post** /v1/vps/server/{id}/start | 
[**ManageServiceStopRescue**](ManageServiceApi.md#ManageServiceStopRescue) | **Delete** /v1/vps/server/{id}/rescue | 
[**ManageServiceStopVps**](ManageServiceApi.md#ManageServiceStopVps) | **Post** /v1/vps/server/{id}/stop | 
[**ManageServiceUnarchive**](ManageServiceApi.md#ManageServiceUnarchive) | **Delete** /v1/vps/archive/{id} | 
[**ManageServiceUpdateInfo**](ManageServiceApi.md#ManageServiceUpdateInfo) | **Put** /v1/vps/server/{id}/info | 



## ManageServiceAttachIpAddress

> ManageAttachIpAddressResponse ManageServiceAttachIpAddress(ctx, id, ipAddress).ManageAttachIpAddressRequest(manageAttachIpAddressRequest).Execute()



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
    ipAddress := "ipAddress_example" // string | 
    manageAttachIpAddressRequest := *openapiclient.NewManageAttachIpAddressRequest() // ManageAttachIpAddressRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManageServiceApi.ManageServiceAttachIpAddress(context.Background(), id, ipAddress).ManageAttachIpAddressRequest(manageAttachIpAddressRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManageServiceApi.ManageServiceAttachIpAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ManageServiceAttachIpAddress`: ManageAttachIpAddressResponse
    fmt.Fprintf(os.Stdout, "Response from `ManageServiceApi.ManageServiceAttachIpAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**ipAddress** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManageServiceAttachIpAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **manageAttachIpAddressRequest** | [**ManageAttachIpAddressRequest**](ManageAttachIpAddressRequest.md) |  | 

### Return type

[**ManageAttachIpAddressResponse**](ManageAttachIpAddressResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManageServiceAttachSshKey

> ManageAttachSshKeyResponse ManageServiceAttachSshKey(ctx, id, sshKeyId).Execute()



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
    sshKeyId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManageServiceApi.ManageServiceAttachSshKey(context.Background(), id, sshKeyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManageServiceApi.ManageServiceAttachSshKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ManageServiceAttachSshKey`: ManageAttachSshKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `ManageServiceApi.ManageServiceAttachSshKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**sshKeyId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManageServiceAttachSshKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ManageAttachSshKeyResponse**](ManageAttachSshKeyResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManageServiceAttachToPrivateNetwork

> ManageAttachToPrivateNetworkResponse ManageServiceAttachToPrivateNetwork(ctx, id, networkId).ManageAttachToPrivateNetworkRequest(manageAttachToPrivateNetworkRequest).Execute()



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
    networkId := "networkId_example" // string | 
    manageAttachToPrivateNetworkRequest := *openapiclient.NewManageAttachToPrivateNetworkRequest() // ManageAttachToPrivateNetworkRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManageServiceApi.ManageServiceAttachToPrivateNetwork(context.Background(), id, networkId).ManageAttachToPrivateNetworkRequest(manageAttachToPrivateNetworkRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManageServiceApi.ManageServiceAttachToPrivateNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ManageServiceAttachToPrivateNetwork`: ManageAttachToPrivateNetworkResponse
    fmt.Fprintf(os.Stdout, "Response from `ManageServiceApi.ManageServiceAttachToPrivateNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManageServiceAttachToPrivateNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **manageAttachToPrivateNetworkRequest** | [**ManageAttachToPrivateNetworkRequest**](ManageAttachToPrivateNetworkRequest.md) |  | 

### Return type

[**ManageAttachToPrivateNetworkResponse**](ManageAttachToPrivateNetworkResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManageServiceChangeConfiguration

> ManageChangeConfigurationResponse ManageServiceChangeConfiguration(ctx, id).ManageChangeConfigurationRequest(manageChangeConfigurationRequest).Execute()



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
    manageChangeConfigurationRequest := *openapiclient.NewManageChangeConfigurationRequest() // ManageChangeConfigurationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManageServiceApi.ManageServiceChangeConfiguration(context.Background(), id).ManageChangeConfigurationRequest(manageChangeConfigurationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManageServiceApi.ManageServiceChangeConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ManageServiceChangeConfiguration`: ManageChangeConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `ManageServiceApi.ManageServiceChangeConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManageServiceChangeConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **manageChangeConfigurationRequest** | [**ManageChangeConfigurationRequest**](ManageChangeConfigurationRequest.md) |  | 

### Return type

[**ManageChangeConfigurationResponse**](ManageChangeConfigurationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManageServiceChangeSshAccess

> ManageChangeSshAccessResponse ManageServiceChangeSshAccess(ctx, id).ManageChangeSshAccessRequest(manageChangeSshAccessRequest).Execute()



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
    manageChangeSshAccessRequest := *openapiclient.NewManageChangeSshAccessRequest() // ManageChangeSshAccessRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManageServiceApi.ManageServiceChangeSshAccess(context.Background(), id).ManageChangeSshAccessRequest(manageChangeSshAccessRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManageServiceApi.ManageServiceChangeSshAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ManageServiceChangeSshAccess`: ManageChangeSshAccessResponse
    fmt.Fprintf(os.Stdout, "Response from `ManageServiceApi.ManageServiceChangeSshAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManageServiceChangeSshAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **manageChangeSshAccessRequest** | [**ManageChangeSshAccessRequest**](ManageChangeSshAccessRequest.md) |  | 

### Return type

[**ManageChangeSshAccessResponse**](ManageChangeSshAccessResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManageServiceCheckSoftwareRequirements

> ManageCheckSoftwareRequirementsResponse ManageServiceCheckSoftwareRequirements(ctx).ManageCheckSoftwareRequirementsRequest(manageCheckSoftwareRequirementsRequest).Execute()



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
    manageCheckSoftwareRequirementsRequest := *openapiclient.NewManageCheckSoftwareRequirementsRequest() // ManageCheckSoftwareRequirementsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManageServiceApi.ManageServiceCheckSoftwareRequirements(context.Background()).ManageCheckSoftwareRequirementsRequest(manageCheckSoftwareRequirementsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManageServiceApi.ManageServiceCheckSoftwareRequirements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ManageServiceCheckSoftwareRequirements`: ManageCheckSoftwareRequirementsResponse
    fmt.Fprintf(os.Stdout, "Response from `ManageServiceApi.ManageServiceCheckSoftwareRequirements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiManageServiceCheckSoftwareRequirementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **manageCheckSoftwareRequirementsRequest** | [**ManageCheckSoftwareRequirementsRequest**](ManageCheckSoftwareRequirementsRequest.md) |  | 

### Return type

[**ManageCheckSoftwareRequirementsResponse**](ManageCheckSoftwareRequirementsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManageServiceCreateVps

> ManageCreateVpsResponse ManageServiceCreateVps(ctx).ManageCreateVpsRequest(manageCreateVpsRequest).Execute()



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
    manageCreateVpsRequest := *openapiclient.NewManageCreateVpsRequest() // ManageCreateVpsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManageServiceApi.ManageServiceCreateVps(context.Background()).ManageCreateVpsRequest(manageCreateVpsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManageServiceApi.ManageServiceCreateVps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ManageServiceCreateVps`: ManageCreateVpsResponse
    fmt.Fprintf(os.Stdout, "Response from `ManageServiceApi.ManageServiceCreateVps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiManageServiceCreateVpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **manageCreateVpsRequest** | [**ManageCreateVpsRequest**](ManageCreateVpsRequest.md) |  | 

### Return type

[**ManageCreateVpsResponse**](ManageCreateVpsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManageServiceDetachFromPrivateNetwork

> ManageDetachFromPrivateNetworkResponse ManageServiceDetachFromPrivateNetwork(ctx, id, networkId).Execute()



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
    networkId := "networkId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManageServiceApi.ManageServiceDetachFromPrivateNetwork(context.Background(), id, networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManageServiceApi.ManageServiceDetachFromPrivateNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ManageServiceDetachFromPrivateNetwork`: ManageDetachFromPrivateNetworkResponse
    fmt.Fprintf(os.Stdout, "Response from `ManageServiceApi.ManageServiceDetachFromPrivateNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManageServiceDetachFromPrivateNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ManageDetachFromPrivateNetworkResponse**](ManageDetachFromPrivateNetworkResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManageServiceDetachIpAddress

> ManageDetachIpAddressResponse ManageServiceDetachIpAddress(ctx, ipAddress).Execute()



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
    resp, r, err := apiClient.ManageServiceApi.ManageServiceDetachIpAddress(context.Background(), ipAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManageServiceApi.ManageServiceDetachIpAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ManageServiceDetachIpAddress`: ManageDetachIpAddressResponse
    fmt.Fprintf(os.Stdout, "Response from `ManageServiceApi.ManageServiceDetachIpAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ipAddress** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManageServiceDetachIpAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ManageDetachIpAddressResponse**](ManageDetachIpAddressResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManageServiceDetachSshKey

> ManageDetachSshKeyResponse ManageServiceDetachSshKey(ctx, id, sshKeyId).Execute()



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
    sshKeyId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManageServiceApi.ManageServiceDetachSshKey(context.Background(), id, sshKeyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManageServiceApi.ManageServiceDetachSshKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ManageServiceDetachSshKey`: ManageDetachSshKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `ManageServiceApi.ManageServiceDetachSshKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**sshKeyId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManageServiceDetachSshKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ManageDetachSshKeyResponse**](ManageDetachSshKeyResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManageServiceDisablePostInstallAlert

> ManageDisablePostInstallAlertResponse ManageServiceDisablePostInstallAlert(ctx, id).Execute()



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
    resp, r, err := apiClient.ManageServiceApi.ManageServiceDisablePostInstallAlert(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManageServiceApi.ManageServiceDisablePostInstallAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ManageServiceDisablePostInstallAlert`: ManageDisablePostInstallAlertResponse
    fmt.Fprintf(os.Stdout, "Response from `ManageServiceApi.ManageServiceDisablePostInstallAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManageServiceDisablePostInstallAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ManageDisablePostInstallAlertResponse**](ManageDisablePostInstallAlertResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManageServiceGetAvailableConfiguration

> ManageGetAvailableConfigurationResponse ManageServiceGetAvailableConfiguration(ctx).Execute()



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
    resp, r, err := apiClient.ManageServiceApi.ManageServiceGetAvailableConfiguration(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManageServiceApi.ManageServiceGetAvailableConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ManageServiceGetAvailableConfiguration`: ManageGetAvailableConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `ManageServiceApi.ManageServiceGetAvailableConfiguration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiManageServiceGetAvailableConfigurationRequest struct via the builder pattern


### Return type

[**ManageGetAvailableConfigurationResponse**](ManageGetAvailableConfigurationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManageServiceGetFileManagerSettings

> ManageGetFileManagerSettingsResponse ManageServiceGetFileManagerSettings(ctx, id).Execute()



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
    resp, r, err := apiClient.ManageServiceApi.ManageServiceGetFileManagerSettings(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManageServiceApi.ManageServiceGetFileManagerSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ManageServiceGetFileManagerSettings`: ManageGetFileManagerSettingsResponse
    fmt.Fprintf(os.Stdout, "Response from `ManageServiceApi.ManageServiceGetFileManagerSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManageServiceGetFileManagerSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ManageGetFileManagerSettingsResponse**](ManageGetFileManagerSettingsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManageServiceGetHistory

> ManageGetHistoryResponse ManageServiceGetHistory(ctx, id).Execute()



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
    resp, r, err := apiClient.ManageServiceApi.ManageServiceGetHistory(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManageServiceApi.ManageServiceGetHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ManageServiceGetHistory`: ManageGetHistoryResponse
    fmt.Fprintf(os.Stdout, "Response from `ManageServiceApi.ManageServiceGetHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManageServiceGetHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ManageGetHistoryResponse**](ManageGetHistoryResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManageServiceGetInfo

> ManageGetInfoResponse ManageServiceGetInfo(ctx, id).Execute()



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
    resp, r, err := apiClient.ManageServiceApi.ManageServiceGetInfo(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManageServiceApi.ManageServiceGetInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ManageServiceGetInfo`: ManageGetInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `ManageServiceApi.ManageServiceGetInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManageServiceGetInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ManageGetInfoResponse**](ManageGetInfoResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManageServiceGetInstalledSoftware

> ManageGetInstalledSoftwareResponse ManageServiceGetInstalledSoftware(ctx, id).Execute()



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
    resp, r, err := apiClient.ManageServiceApi.ManageServiceGetInstalledSoftware(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManageServiceApi.ManageServiceGetInstalledSoftware``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ManageServiceGetInstalledSoftware`: ManageGetInstalledSoftwareResponse
    fmt.Fprintf(os.Stdout, "Response from `ManageServiceApi.ManageServiceGetInstalledSoftware`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManageServiceGetInstalledSoftwareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ManageGetInstalledSoftwareResponse**](ManageGetInstalledSoftwareResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManageServiceGetList

> ManageGetListResponse ManageServiceGetList(ctx).Execute()



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
    resp, r, err := apiClient.ManageServiceApi.ManageServiceGetList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManageServiceApi.ManageServiceGetList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ManageServiceGetList`: ManageGetListResponse
    fmt.Fprintf(os.Stdout, "Response from `ManageServiceApi.ManageServiceGetList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiManageServiceGetListRequest struct via the builder pattern


### Return type

[**ManageGetListResponse**](ManageGetListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManageServiceGetRegionList

> ManageGetRegionListResponse ManageServiceGetRegionList(ctx).Execute()



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
    resp, r, err := apiClient.ManageServiceApi.ManageServiceGetRegionList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManageServiceApi.ManageServiceGetRegionList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ManageServiceGetRegionList`: ManageGetRegionListResponse
    fmt.Fprintf(os.Stdout, "Response from `ManageServiceApi.ManageServiceGetRegionList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiManageServiceGetRegionListRequest struct via the builder pattern


### Return type

[**ManageGetRegionListResponse**](ManageGetRegionListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManageServiceGetStatuses

> ManageGetStatusesResponse ManageServiceGetStatuses(ctx).Execute()



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
    resp, r, err := apiClient.ManageServiceApi.ManageServiceGetStatuses(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManageServiceApi.ManageServiceGetStatuses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ManageServiceGetStatuses`: ManageGetStatusesResponse
    fmt.Fprintf(os.Stdout, "Response from `ManageServiceApi.ManageServiceGetStatuses`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiManageServiceGetStatusesRequest struct via the builder pattern


### Return type

[**ManageGetStatusesResponse**](ManageGetStatusesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManageServiceRebootVps

> ManageRebootVpsResponse ManageServiceRebootVps(ctx, id).Execute()



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
    resp, r, err := apiClient.ManageServiceApi.ManageServiceRebootVps(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManageServiceApi.ManageServiceRebootVps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ManageServiceRebootVps`: ManageRebootVpsResponse
    fmt.Fprintf(os.Stdout, "Response from `ManageServiceApi.ManageServiceRebootVps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManageServiceRebootVpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ManageRebootVpsResponse**](ManageRebootVpsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManageServiceReinstall

> ManageReinstallResponse ManageServiceReinstall(ctx, id).ManageReinstallRequest(manageReinstallRequest).Execute()



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
    manageReinstallRequest := *openapiclient.NewManageReinstallRequest() // ManageReinstallRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManageServiceApi.ManageServiceReinstall(context.Background(), id).ManageReinstallRequest(manageReinstallRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManageServiceApi.ManageServiceReinstall``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ManageServiceReinstall`: ManageReinstallResponse
    fmt.Fprintf(os.Stdout, "Response from `ManageServiceApi.ManageServiceReinstall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManageServiceReinstallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **manageReinstallRequest** | [**ManageReinstallRequest**](ManageReinstallRequest.md) |  | 

### Return type

[**ManageReinstallResponse**](ManageReinstallResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManageServiceRemoveVps

> ManageRemoveVpsResponse ManageServiceRemoveVps(ctx, id).Execute()



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
    resp, r, err := apiClient.ManageServiceApi.ManageServiceRemoveVps(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManageServiceApi.ManageServiceRemoveVps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ManageServiceRemoveVps`: ManageRemoveVpsResponse
    fmt.Fprintf(os.Stdout, "Response from `ManageServiceApi.ManageServiceRemoveVps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManageServiceRemoveVpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ManageRemoveVpsResponse**](ManageRemoveVpsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManageServiceReserveVpsSubdomain

> ManageReserveVpsSubdomainResponse ManageServiceReserveVpsSubdomain(ctx).Execute()



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
    resp, r, err := apiClient.ManageServiceApi.ManageServiceReserveVpsSubdomain(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManageServiceApi.ManageServiceReserveVpsSubdomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ManageServiceReserveVpsSubdomain`: ManageReserveVpsSubdomainResponse
    fmt.Fprintf(os.Stdout, "Response from `ManageServiceApi.ManageServiceReserveVpsSubdomain`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiManageServiceReserveVpsSubdomainRequest struct via the builder pattern


### Return type

[**ManageReserveVpsSubdomainResponse**](ManageReserveVpsSubdomainResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManageServiceResetPassword

> ManageResetPasswordResponse ManageServiceResetPassword(ctx, id).Execute()



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
    resp, r, err := apiClient.ManageServiceApi.ManageServiceResetPassword(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManageServiceApi.ManageServiceResetPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ManageServiceResetPassword`: ManageResetPasswordResponse
    fmt.Fprintf(os.Stdout, "Response from `ManageServiceApi.ManageServiceResetPassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManageServiceResetPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ManageResetPasswordResponse**](ManageResetPasswordResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManageServiceResetVps

> ManageResetVpsResponse ManageServiceResetVps(ctx, id).Execute()



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
    resp, r, err := apiClient.ManageServiceApi.ManageServiceResetVps(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManageServiceApi.ManageServiceResetVps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ManageServiceResetVps`: ManageResetVpsResponse
    fmt.Fprintf(os.Stdout, "Response from `ManageServiceApi.ManageServiceResetVps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManageServiceResetVpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ManageResetVpsResponse**](ManageResetVpsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManageServiceStartRescue

> ManageStartRescueResponse ManageServiceStartRescue(ctx, id).Execute()



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
    resp, r, err := apiClient.ManageServiceApi.ManageServiceStartRescue(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManageServiceApi.ManageServiceStartRescue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ManageServiceStartRescue`: ManageStartRescueResponse
    fmt.Fprintf(os.Stdout, "Response from `ManageServiceApi.ManageServiceStartRescue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManageServiceStartRescueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ManageStartRescueResponse**](ManageStartRescueResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManageServiceStartVps

> ManageStartVpsResponse ManageServiceStartVps(ctx, id).Execute()



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
    resp, r, err := apiClient.ManageServiceApi.ManageServiceStartVps(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManageServiceApi.ManageServiceStartVps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ManageServiceStartVps`: ManageStartVpsResponse
    fmt.Fprintf(os.Stdout, "Response from `ManageServiceApi.ManageServiceStartVps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManageServiceStartVpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ManageStartVpsResponse**](ManageStartVpsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManageServiceStopRescue

> ManageStopRescueResponse ManageServiceStopRescue(ctx, id).Execute()



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
    resp, r, err := apiClient.ManageServiceApi.ManageServiceStopRescue(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManageServiceApi.ManageServiceStopRescue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ManageServiceStopRescue`: ManageStopRescueResponse
    fmt.Fprintf(os.Stdout, "Response from `ManageServiceApi.ManageServiceStopRescue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManageServiceStopRescueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ManageStopRescueResponse**](ManageStopRescueResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManageServiceStopVps

> ManageStopVpsResponse ManageServiceStopVps(ctx, id).Execute()



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
    resp, r, err := apiClient.ManageServiceApi.ManageServiceStopVps(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManageServiceApi.ManageServiceStopVps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ManageServiceStopVps`: ManageStopVpsResponse
    fmt.Fprintf(os.Stdout, "Response from `ManageServiceApi.ManageServiceStopVps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManageServiceStopVpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ManageStopVpsResponse**](ManageStopVpsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManageServiceUnarchive

> ManageUnarchiveResponse ManageServiceUnarchive(ctx, id).Execute()



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
    resp, r, err := apiClient.ManageServiceApi.ManageServiceUnarchive(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManageServiceApi.ManageServiceUnarchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ManageServiceUnarchive`: ManageUnarchiveResponse
    fmt.Fprintf(os.Stdout, "Response from `ManageServiceApi.ManageServiceUnarchive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManageServiceUnarchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ManageUnarchiveResponse**](ManageUnarchiveResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManageServiceUpdateInfo

> ManageUpdateInfoResponse ManageServiceUpdateInfo(ctx, id).ManageUpdateInfoRequest(manageUpdateInfoRequest).Execute()



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
    manageUpdateInfoRequest := *openapiclient.NewManageUpdateInfoRequest() // ManageUpdateInfoRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManageServiceApi.ManageServiceUpdateInfo(context.Background(), id).ManageUpdateInfoRequest(manageUpdateInfoRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManageServiceApi.ManageServiceUpdateInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ManageServiceUpdateInfo`: ManageUpdateInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `ManageServiceApi.ManageServiceUpdateInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManageServiceUpdateInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **manageUpdateInfoRequest** | [**ManageUpdateInfoRequest**](ManageUpdateInfoRequest.md) |  | 

### Return type

[**ManageUpdateInfoResponse**](ManageUpdateInfoResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

