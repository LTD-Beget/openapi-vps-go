/*
API Облачных серверов

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.4.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package begetOpenapiVps

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)


// MarketplaceServiceApiService MarketplaceServiceApi service
type MarketplaceServiceApiService service

type ApiMarketplaceServiceGetSoftwareListRequest struct {
	ctx context.Context
	ApiService *MarketplaceServiceApiService
	categoryName *string
	displayName *string
	isPinned *bool
}

func (r ApiMarketplaceServiceGetSoftwareListRequest) CategoryName(categoryName string) ApiMarketplaceServiceGetSoftwareListRequest {
	r.categoryName = &categoryName
	return r
}

func (r ApiMarketplaceServiceGetSoftwareListRequest) DisplayName(displayName string) ApiMarketplaceServiceGetSoftwareListRequest {
	r.displayName = &displayName
	return r
}

func (r ApiMarketplaceServiceGetSoftwareListRequest) IsPinned(isPinned bool) ApiMarketplaceServiceGetSoftwareListRequest {
	r.isPinned = &isPinned
	return r
}

func (r ApiMarketplaceServiceGetSoftwareListRequest) Execute() (*MarketplaceGetSoftwareListResponse, *http.Response, error) {
	return r.ApiService.MarketplaceServiceGetSoftwareListExecute(r)
}

/*
MarketplaceServiceGetSoftwareList Method for MarketplaceServiceGetSoftwareList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiMarketplaceServiceGetSoftwareListRequest
*/
func (a *MarketplaceServiceApiService) MarketplaceServiceGetSoftwareList(ctx context.Context) ApiMarketplaceServiceGetSoftwareListRequest {
	return ApiMarketplaceServiceGetSoftwareListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MarketplaceGetSoftwareListResponse
func (a *MarketplaceServiceApiService) MarketplaceServiceGetSoftwareListExecute(r ApiMarketplaceServiceGetSoftwareListRequest) (*MarketplaceGetSoftwareListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MarketplaceGetSoftwareListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MarketplaceServiceApiService.MarketplaceServiceGetSoftwareList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/vps/marketplace/software/list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.categoryName != nil {
		localVarQueryParams.Add("category_name", parameterToString(*r.categoryName, ""))
	}
	if r.displayName != nil {
		localVarQueryParams.Add("display_name", parameterToString(*r.displayName, ""))
	}
	if r.isPinned != nil {
		localVarQueryParams.Add("is_pinned", parameterToString(*r.isPinned, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
