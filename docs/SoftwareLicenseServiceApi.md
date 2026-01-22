# \SoftwareLicenseServiceApi

All URIs are relative to *https://api.beget.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SoftwareLicenseServiceChangeLicensePlan**](SoftwareLicenseServiceApi.md#SoftwareLicenseServiceChangeLicensePlan) | **Patch** /v1/vps/software/license/{vps_id} | 
[**SoftwareLicenseServiceGetLicenseInfo**](SoftwareLicenseServiceApi.md#SoftwareLicenseServiceGetLicenseInfo) | **Get** /v1/vps/software/license | 



## SoftwareLicenseServiceChangeLicensePlan

> SoftwareLicenseChangeLicensePlanResponse SoftwareLicenseServiceChangeLicensePlan(ctx, vpsId).SoftwareLicenseChangeLicensePlanRequest(softwareLicenseChangeLicensePlanRequest).Execute()



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
    vpsId := "vpsId_example" // string | 
    softwareLicenseChangeLicensePlanRequest := *openapiclient.NewSoftwareLicenseChangeLicensePlanRequest() // SoftwareLicenseChangeLicensePlanRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SoftwareLicenseServiceApi.SoftwareLicenseServiceChangeLicensePlan(context.Background(), vpsId).SoftwareLicenseChangeLicensePlanRequest(softwareLicenseChangeLicensePlanRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoftwareLicenseServiceApi.SoftwareLicenseServiceChangeLicensePlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SoftwareLicenseServiceChangeLicensePlan`: SoftwareLicenseChangeLicensePlanResponse
    fmt.Fprintf(os.Stdout, "Response from `SoftwareLicenseServiceApi.SoftwareLicenseServiceChangeLicensePlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSoftwareLicenseServiceChangeLicensePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **softwareLicenseChangeLicensePlanRequest** | [**SoftwareLicenseChangeLicensePlanRequest**](SoftwareLicenseChangeLicensePlanRequest.md) |  | 

### Return type

[**SoftwareLicenseChangeLicensePlanResponse**](SoftwareLicenseChangeLicensePlanResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SoftwareLicenseServiceGetLicenseInfo

> SoftwareLicenseGetLicenseInfoResponse SoftwareLicenseServiceGetLicenseInfo(ctx).LicenseId(licenseId).Execute()



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
    licenseId := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SoftwareLicenseServiceApi.SoftwareLicenseServiceGetLicenseInfo(context.Background()).LicenseId(licenseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoftwareLicenseServiceApi.SoftwareLicenseServiceGetLicenseInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SoftwareLicenseServiceGetLicenseInfo`: SoftwareLicenseGetLicenseInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `SoftwareLicenseServiceApi.SoftwareLicenseServiceGetLicenseInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSoftwareLicenseServiceGetLicenseInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **licenseId** | **int32** |  | 

### Return type

[**SoftwareLicenseGetLicenseInfoResponse**](SoftwareLicenseGetLicenseInfoResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

