# \QuotesApi

All URIs are relative to *https://api.tdameritrade.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetQuote**](QuotesApi.md#GetQuote) | **Get** /marketdata/{symbol}/quotes | Search for instrument and fundamental data
[**GetQuotes**](QuotesApi.md#GetQuotes) | **Get** /marketdata/quotes | Search for instrument and fundamental data



## GetQuote

> []map[string]interface{} GetQuote(ctx, symbol, apiKey)

Search for instrument and fundamental data

Get quote for a symbol.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbol** | **string**| Account ID | 
**apiKey** | **string**| API Key | 

### Return type

**[]map[string]interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuotes

> []map[string]interface{} GetQuotes(ctx, apiKey, symbol)

Search for instrument and fundamental data

Get quote for one or more symbols.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiKey** | **string**| API Key | 
**symbol** | **string**| Symbol to search for | 

### Return type

**[]map[string]interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

