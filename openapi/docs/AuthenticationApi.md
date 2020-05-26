# \AuthenticationApi

All URIs are relative to *https://api.tdameritrade.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostAccessToken**](AuthenticationApi.md#PostAccessToken) | **Post** /oath2/token | Post Access Token



## PostAccessToken

> EasObject PostAccessToken(ctx, optional)

Post Access Token

The token endpoint returns an access token along with an optional refresh token.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PostAccessTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostAccessTokenOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **easObject** | [**optional.Interface of EasObject**](EasObject.md)| The access token. | 

### Return type

[**EasObject**](EASObject.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

