# \MarketHoursApi

All URIs are relative to *https://api.tdameritrade.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHours**](MarketHoursApi.md#GetHours) | **Get** /marketdata/hours | Operating hours of markets
[**GetMarketHours**](MarketHoursApi.md#GetMarketHours) | **Get** /marketdata/{market}/hours | Operating hours of markets



## GetHours

> Hours GetHours(ctx, apikey, markets, date)

Operating hours of markets

Retrieve market hours for specified markets

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apikey** | **string**| API Key | 
**markets** | **string**| The markets for which you&#39;re requesting market hours, comma-separated. Valid markets are EQUITY, OPTION, FUTURE, BOND, or FOREX. | 
**date** | **string**| The date for which market hours information is requested. Valid ISO-8601 formats are : yyyy-MM-dd and yyyy-MM-dd&#39;T&#39;HH:mm:ssz. | 

### Return type

[**Hours**](Hours.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketHours

> Hours GetMarketHours(ctx, apikey, market, date)

Operating hours of markets

Retrieve market hours for specified single market

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apikey** | **string**| API Key | 
**market** | **string**| The markets for which you&#39;re requesting market hours, comma-separated. Valid markets are EQUITY, OPTION, FUTURE, BOND, or FOREX. | 
**date** | **string**| The date for which market hours information is requested. Valid ISO-8601 formats are : yyyy-MM-dd and yyyy-MM-dd&#39;T&#39;HH:mm:ssz. | 

### Return type

[**Hours**](Hours.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

