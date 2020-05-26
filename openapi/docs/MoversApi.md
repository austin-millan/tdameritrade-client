# \MoversApi

All URIs are relative to *https://api.tdameritrade.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMovers**](MoversApi.md#GetMovers) | **Get** /marketdata/{index}/movers | Retrieve mover information by index symbol, direction type and change



## GetMovers

> Mover GetMovers(ctx, apikey, index, direction, change)

Retrieve mover information by index symbol, direction type and change

Top 10 (up or down) movers by value or percent for a particular market

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apikey** | **string**| API Key | 
**index** | **string**| The index symbol to get movers from. | 
**direction** | **string**| To return movers with the specified directions of up or down. | 
**change** | **string**| To return movers with the specified change types of percent or value. | 

### Return type

[**Mover**](Mover.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

