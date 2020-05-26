# \AccountsAndTradingApi

All URIs are relative to *https://api.tdameritrade.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelOrder**](AccountsAndTradingApi.md#CancelOrder) | **Delete** /accounts/{accountId}/orders/{orderId} | Cancel Order
[**DeleteSavedOrder**](AccountsAndTradingApi.md#DeleteSavedOrder) | **Delete** /accounts/{accountId}/savedorders/{savedOrderId} | Delete Saved Order
[**GetAccount**](AccountsAndTradingApi.md#GetAccount) | **Get** /accounts/{accountId} | Get Account
[**GetAccounts**](AccountsAndTradingApi.md#GetAccounts) | **Get** /accounts | Get Accounts
[**GetOrder**](AccountsAndTradingApi.md#GetOrder) | **Get** /accounts/{accountId}/orders/{orderId} | Get Order
[**GetOrdersByPath**](AccountsAndTradingApi.md#GetOrdersByPath) | **Get** /accounts/{accountId}/orders | Get Orders By Path
[**GetOrdersByQuery**](AccountsAndTradingApi.md#GetOrdersByQuery) | **Get** /orders | Get Orders By Query
[**GetSavedOrder**](AccountsAndTradingApi.md#GetSavedOrder) | **Get** /accounts/{accountId}/savedorders/{savedOrderId} | Get Saved Order
[**GetSavedOrdersByPath**](AccountsAndTradingApi.md#GetSavedOrdersByPath) | **Get** /accounts/{accountId}/savedorders | Get Saved Orders by Path
[**PlaceOrder**](AccountsAndTradingApi.md#PlaceOrder) | **Post** /accounts/{accountId}/orders | Place Order
[**ReplaceOrder**](AccountsAndTradingApi.md#ReplaceOrder) | **Put** /accounts/{accountId}/orders/{orderId} | Replace Order
[**ReplaceSavedOrder**](AccountsAndTradingApi.md#ReplaceSavedOrder) | **Put** /accounts/{accountId}/savedorders/{savedOrderId} | Replace Saved Order
[**SaveOrder**](AccountsAndTradingApi.md#SaveOrder) | **Post** /accounts/{accountId}/savedorders | Create Saved Order



## CancelOrder

> CancelOrder(ctx, accountId, orderId)

Cancel Order

Cancel a specific order for a specific account. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32**| Account Number | 
**orderId** | **int64**| Order Number | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSavedOrder

> DeleteSavedOrder(ctx, accountId, savedOrderId)

Delete Saved Order

Delete a specific saved order for a specific account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32**| Account Number | 
**savedOrderId** | **int64**| Order Number | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccount

> Account GetAccount(ctx, accountId)

Get Account

Account balances, positions, and orders for a specific account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32**| Account Number | 

### Return type

[**Account**](Account.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccounts

> Account GetAccounts(ctx, )

Get Accounts

Account balances, positions, and orders for a specific account.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**Account**](Account.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrder

> Order GetOrder(ctx, accountId, orderId)

Get Order

Get a specific order for a specific account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32**| Account Number | 
**orderId** | **int64**| Order Number | 

### Return type

[**Order**](Order.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrdersByPath

> Order GetOrdersByPath(ctx, accountId, maxResults, fromEnteredTime, toEnteredTime, status)

Get Orders By Path

Orders for a specific account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32**| Account Number | 
**maxResults** | **int32**| The maximum number of orders to retrieve. | 
**fromEnteredTime** | **string**| Specifies that no orders entered before this time should be returned. Valid ISO-8601 formats are  yyyy-MM-dd and yyyy-MM-dd&#39;T&#39;HH:mm:ssz Date must be within 60 days from today&#39;s date. &#39;toEnteredTime&#39; must also be set. | 
**toEnteredTime** | **string**| Specifies that no orders entered after this time should be returned.Valid ISO-8601 formats are :yyyy-MM-dd and yyyy-MM-dd&#39;T&#39;HH:mm:ssz. &#39;fromEnteredTime&#39; must also be set. | 
**status** | **int64**| The maximum number of orders to retrieve. | 

### Return type

[**Order**](Order.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrdersByQuery

> GetOrdersByQuery(ctx, accountId, maxResults, fromEnteredTime, toEnteredTime, status)

Get Orders By Query

All orders for a specific account or, if account ID isn't specified, orders will be returned for all linked accounts

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32**| Account Number | 
**maxResults** | **int32**| The maximum number of orders to retrieve. | 
**fromEnteredTime** | **string**| Specifies that no orders entered before this time should be returned. Valid ISO-8601 formats are  yyyy-MM-dd and yyyy-MM-dd&#39;T&#39;HH:mm:ssz Date must be within 60 days from today&#39;s date. &#39;toEnteredTime&#39; must also be set. | 
**toEnteredTime** | **string**| Specifies that no orders entered after this time should be returned.Valid ISO-8601 formats are :yyyy-MM-dd and yyyy-MM-dd&#39;T&#39;HH:mm:ssz. &#39;fromEnteredTime&#39; must also be set. | 
**status** | **int64**| The maximum number of orders to retrieve. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSavedOrder

> Order GetSavedOrder(ctx, accountId, savedOrderId)

Get Saved Order

Specific saved order by its ID, for a specific account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32**| Account Number | 
**savedOrderId** | **int64**| Order Number | 

### Return type

[**Order**](Order.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSavedOrdersByPath

> []SavedOrder GetSavedOrdersByPath(ctx, accountId)

Get Saved Orders by Path

Saved orders for a specific account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32**| Account Number | 

### Return type

[**[]SavedOrder**](SavedOrder.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlaceOrder

> Order PlaceOrder(ctx, accountId, optional)

Place Order

Place an order for a specific account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32**| Account Number | 
 **optional** | ***PlaceOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PlaceOrderOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **order** | [**optional.Interface of Order**](Order.md)| The order to place. | 

### Return type

[**Order**](Order.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceOrder

> ReplaceOrder(ctx, accountId, orderId, optional)

Replace Order

Replace an existing order for an account. The existing order will be replaced by the new order. Once replaced, the old order will be canceled and a new order will be created.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32**| Account Number | 
**orderId** | **int64**| Order Number | 
 **optional** | ***ReplaceOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReplaceOrderOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **order** | [**optional.Interface of Order**](Order.md)| The order to place. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceSavedOrder

> ReplaceSavedOrder(ctx, accountId, savedOrderId, optional)

Replace Saved Order

Replace an existing saved order for an account. The existing saved order will be replaced by the new order.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32**| Account Number | 
**savedOrderId** | **int64**| Order Number | 
 **optional** | ***ReplaceSavedOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReplaceSavedOrderOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **savedOrder** | [**optional.Interface of SavedOrder**](SavedOrder.md)| The order to place. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveOrder

> SaveOrder(ctx, accountId, orderId, optional)

Create Saved Order

Save an order for a specific account. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32**| Account Number | 
**orderId** | **int64**| Order Number | 
 **optional** | ***SaveOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SaveOrderOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **savedOrder** | [**optional.Interface of SavedOrder**](SavedOrder.md)| The order to save. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

