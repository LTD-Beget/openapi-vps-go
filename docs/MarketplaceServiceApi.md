# \MarketplaceServiceApi

All URIs are relative to *https://api.beget.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MarketplaceServiceGetSoftwareList**](MarketplaceServiceApi.md#MarketplaceServiceGetSoftwareList) | **Get** /v1/vps/marketplace/software/list | 



## MarketplaceServiceGetSoftwareList

> MarketplaceGetSoftwareListResponse MarketplaceServiceGetSoftwareList(ctx).CategoryName(categoryName).DisplayName(displayName).IsPinned(isPinned).Execute()



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
    categoryName := "categoryName_example" // string |  (optional)
    displayName := "displayName_example" // string |  (optional)
    isPinned := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketplaceServiceApi.MarketplaceServiceGetSoftwareList(context.Background()).CategoryName(categoryName).DisplayName(displayName).IsPinned(isPinned).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceServiceApi.MarketplaceServiceGetSoftwareList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarketplaceServiceGetSoftwareList`: MarketplaceGetSoftwareListResponse
    fmt.Fprintf(os.Stdout, "Response from `MarketplaceServiceApi.MarketplaceServiceGetSoftwareList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceServiceGetSoftwareListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **categoryName** | **string** |  | 
 **displayName** | **string** |  | 
 **isPinned** | **bool** |  | 

### Return type

[**MarketplaceGetSoftwareListResponse**](MarketplaceGetSoftwareListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

