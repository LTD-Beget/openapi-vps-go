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
	"strings"
)


// SnapshotServiceApiService SnapshotServiceApi service
type SnapshotServiceApiService service

type ApiSnapshotServiceCreateRequest struct {
	ctx context.Context
	ApiService *SnapshotServiceApiService
	snapshotCreateRequest *SnapshotCreateRequest
}

func (r ApiSnapshotServiceCreateRequest) SnapshotCreateRequest(snapshotCreateRequest SnapshotCreateRequest) ApiSnapshotServiceCreateRequest {
	r.snapshotCreateRequest = &snapshotCreateRequest
	return r
}

func (r ApiSnapshotServiceCreateRequest) Execute() (*SnapshotCreateResponse, *http.Response, error) {
	return r.ApiService.SnapshotServiceCreateExecute(r)
}

/*
SnapshotServiceCreate Method for SnapshotServiceCreate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSnapshotServiceCreateRequest
*/
func (a *SnapshotServiceApiService) SnapshotServiceCreate(ctx context.Context) ApiSnapshotServiceCreateRequest {
	return ApiSnapshotServiceCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SnapshotCreateResponse
func (a *SnapshotServiceApiService) SnapshotServiceCreateExecute(r ApiSnapshotServiceCreateRequest) (*SnapshotCreateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SnapshotCreateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SnapshotServiceApiService.SnapshotServiceCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/vps/snapshot"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.snapshotCreateRequest == nil {
		return localVarReturnValue, nil, reportError("snapshotCreateRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.snapshotCreateRequest
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

type ApiSnapshotServiceCreateCalculatorRequest struct {
	ctx context.Context
	ApiService *SnapshotServiceApiService
	snapshotCreateCalculatorRequest *SnapshotCreateCalculatorRequest
}

func (r ApiSnapshotServiceCreateCalculatorRequest) SnapshotCreateCalculatorRequest(snapshotCreateCalculatorRequest SnapshotCreateCalculatorRequest) ApiSnapshotServiceCreateCalculatorRequest {
	r.snapshotCreateCalculatorRequest = &snapshotCreateCalculatorRequest
	return r
}

func (r ApiSnapshotServiceCreateCalculatorRequest) Execute() (*SnapshotCreateCalculatorResponse, *http.Response, error) {
	return r.ApiService.SnapshotServiceCreateCalculatorExecute(r)
}

/*
SnapshotServiceCreateCalculator Method for SnapshotServiceCreateCalculator

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSnapshotServiceCreateCalculatorRequest
*/
func (a *SnapshotServiceApiService) SnapshotServiceCreateCalculator(ctx context.Context) ApiSnapshotServiceCreateCalculatorRequest {
	return ApiSnapshotServiceCreateCalculatorRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SnapshotCreateCalculatorResponse
func (a *SnapshotServiceApiService) SnapshotServiceCreateCalculatorExecute(r ApiSnapshotServiceCreateCalculatorRequest) (*SnapshotCreateCalculatorResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SnapshotCreateCalculatorResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SnapshotServiceApiService.SnapshotServiceCreateCalculator")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/vps/snapshot/calculator"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.snapshotCreateCalculatorRequest == nil {
		return localVarReturnValue, nil, reportError("snapshotCreateCalculatorRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.snapshotCreateCalculatorRequest
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

type ApiSnapshotServiceEditRequest struct {
	ctx context.Context
	ApiService *SnapshotServiceApiService
	id string
	snapshotEditRequest *SnapshotEditRequest
}

func (r ApiSnapshotServiceEditRequest) SnapshotEditRequest(snapshotEditRequest SnapshotEditRequest) ApiSnapshotServiceEditRequest {
	r.snapshotEditRequest = &snapshotEditRequest
	return r
}

func (r ApiSnapshotServiceEditRequest) Execute() (*SnapshotEditResponse, *http.Response, error) {
	return r.ApiService.SnapshotServiceEditExecute(r)
}

/*
SnapshotServiceEdit Method for SnapshotServiceEdit

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiSnapshotServiceEditRequest
*/
func (a *SnapshotServiceApiService) SnapshotServiceEdit(ctx context.Context, id string) ApiSnapshotServiceEditRequest {
	return ApiSnapshotServiceEditRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return SnapshotEditResponse
func (a *SnapshotServiceApiService) SnapshotServiceEditExecute(r ApiSnapshotServiceEditRequest) (*SnapshotEditResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SnapshotEditResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SnapshotServiceApiService.SnapshotServiceEdit")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/vps/snapshot/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.snapshotEditRequest == nil {
		return localVarReturnValue, nil, reportError("snapshotEditRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.snapshotEditRequest
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

type ApiSnapshotServiceGetAllRequest struct {
	ctx context.Context
	ApiService *SnapshotServiceApiService
}

func (r ApiSnapshotServiceGetAllRequest) Execute() (*SnapshotGetAllResponse, *http.Response, error) {
	return r.ApiService.SnapshotServiceGetAllExecute(r)
}

/*
SnapshotServiceGetAll Method for SnapshotServiceGetAll

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSnapshotServiceGetAllRequest
*/
func (a *SnapshotServiceApiService) SnapshotServiceGetAll(ctx context.Context) ApiSnapshotServiceGetAllRequest {
	return ApiSnapshotServiceGetAllRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SnapshotGetAllResponse
func (a *SnapshotServiceApiService) SnapshotServiceGetAllExecute(r ApiSnapshotServiceGetAllRequest) (*SnapshotGetAllResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SnapshotGetAllResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SnapshotServiceApiService.SnapshotServiceGetAll")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/vps/snapshot"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiSnapshotServiceGetAllRestoresRequest struct {
	ctx context.Context
	ApiService *SnapshotServiceApiService
	id *string
}

func (r ApiSnapshotServiceGetAllRestoresRequest) Id(id string) ApiSnapshotServiceGetAllRestoresRequest {
	r.id = &id
	return r
}

func (r ApiSnapshotServiceGetAllRestoresRequest) Execute() (*SnapshotGetAllRestoresResponse, *http.Response, error) {
	return r.ApiService.SnapshotServiceGetAllRestoresExecute(r)
}

/*
SnapshotServiceGetAllRestores Method for SnapshotServiceGetAllRestores

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSnapshotServiceGetAllRestoresRequest
*/
func (a *SnapshotServiceApiService) SnapshotServiceGetAllRestores(ctx context.Context) ApiSnapshotServiceGetAllRestoresRequest {
	return ApiSnapshotServiceGetAllRestoresRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SnapshotGetAllRestoresResponse
func (a *SnapshotServiceApiService) SnapshotServiceGetAllRestoresExecute(r ApiSnapshotServiceGetAllRestoresRequest) (*SnapshotGetAllRestoresResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SnapshotGetAllRestoresResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SnapshotServiceApiService.SnapshotServiceGetAllRestores")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/vps/snapshot/restore"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.id != nil {
		localVarQueryParams.Add("id", parameterToString(*r.id, ""))
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

type ApiSnapshotServiceRemoveRequest struct {
	ctx context.Context
	ApiService *SnapshotServiceApiService
	id string
}

func (r ApiSnapshotServiceRemoveRequest) Execute() (*SnapshotRemoveResponse, *http.Response, error) {
	return r.ApiService.SnapshotServiceRemoveExecute(r)
}

/*
SnapshotServiceRemove Method for SnapshotServiceRemove

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiSnapshotServiceRemoveRequest
*/
func (a *SnapshotServiceApiService) SnapshotServiceRemove(ctx context.Context, id string) ApiSnapshotServiceRemoveRequest {
	return ApiSnapshotServiceRemoveRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return SnapshotRemoveResponse
func (a *SnapshotServiceApiService) SnapshotServiceRemoveExecute(r ApiSnapshotServiceRemoveRequest) (*SnapshotRemoveResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SnapshotRemoveResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SnapshotServiceApiService.SnapshotServiceRemove")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/vps/snapshot/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiSnapshotServiceRestoreRequest struct {
	ctx context.Context
	ApiService *SnapshotServiceApiService
	id string
	snapshotRestoreRequest *SnapshotRestoreRequest
}

func (r ApiSnapshotServiceRestoreRequest) SnapshotRestoreRequest(snapshotRestoreRequest SnapshotRestoreRequest) ApiSnapshotServiceRestoreRequest {
	r.snapshotRestoreRequest = &snapshotRestoreRequest
	return r
}

func (r ApiSnapshotServiceRestoreRequest) Execute() (*SnapshotRestoreResponse, *http.Response, error) {
	return r.ApiService.SnapshotServiceRestoreExecute(r)
}

/*
SnapshotServiceRestore Method for SnapshotServiceRestore

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiSnapshotServiceRestoreRequest
*/
func (a *SnapshotServiceApiService) SnapshotServiceRestore(ctx context.Context, id string) ApiSnapshotServiceRestoreRequest {
	return ApiSnapshotServiceRestoreRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return SnapshotRestoreResponse
func (a *SnapshotServiceApiService) SnapshotServiceRestoreExecute(r ApiSnapshotServiceRestoreRequest) (*SnapshotRestoreResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SnapshotRestoreResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SnapshotServiceApiService.SnapshotServiceRestore")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/vps/snapshot/{id}/restore"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.snapshotRestoreRequest == nil {
		return localVarReturnValue, nil, reportError("snapshotRestoreRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.snapshotRestoreRequest
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
