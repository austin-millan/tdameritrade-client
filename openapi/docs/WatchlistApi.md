# \WatchlistApi

All URIs are relative to *https://api.tdameritrade.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWatchlist**](WatchlistApi.md#CreateWatchlist) | **Post** /accounts/{accountId}/watchlists | APIs to perform CRUD operations on Account Watchlist
[**DeleteWatchlist**](WatchlistApi.md#DeleteWatchlist) | **Delete** /accounts/{accountId}/watchlists/{watchlistId} | APIs to perform CRUD operations on Account Watchlist
[**GetWatchlist**](WatchlistApi.md#GetWatchlist) | **Get** /accounts/{accountId}/watchlists/{watchlistId} | APIs to perform CRUD operations on Account Watchlist
[**GetWatchlistMultipleAccounts**](WatchlistApi.md#GetWatchlistMultipleAccounts) | **Get** /accounts/watchlists | APIs to perform CRUD operations on Account Watchlist
[**GetWatchlistSingleAccount**](WatchlistApi.md#GetWatchlistSingleAccount) | **Get** /accounts/{accountId}/watchlists | APIs to perform CRUD operations on Account Watchlist
[**ReplaceWatchlist**](WatchlistApi.md#ReplaceWatchlist) | **Put** /accounts/{accountId}/watchlists/{watchlistId} | APIs to perform CRUD operations on Account Watchlist
[**UpdateWatchlist**](WatchlistApi.md#UpdateWatchlist) | **Patch** /accounts/{accountId}/watchlists/{watchlistId} | APIs to perform CRUD operations on Account Watchlist



## CreateWatchlist

> CreateWatchlist(ctx, accountId, optional)

APIs to perform CRUD operations on Account Watchlist

Create watchlist for specific account.This method does not verify that the symbol or asset type are valid.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
 **optional** | ***CreateWatchlistOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateWatchlistOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createWatchlist** | [**optional.Interface of CreateWatchlist**](CreateWatchlist.md)|  | 

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


## DeleteWatchlist

> DeleteWatchlist(ctx, accountId, watchlistId)

APIs to perform CRUD operations on Account Watchlist

Delete watchlist for a specific account. This method does not verify that the symbol or asset type are valid.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**watchlistId** | **string**| Watchlist ID | 

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


## GetWatchlist

> []Watchlist GetWatchlist(ctx, accountId, watchlistId)

APIs to perform CRUD operations on Account Watchlist

Specific watchlist for a specific account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**watchlistId** | **string**| Watchlist ID | 

### Return type

[**[]Watchlist**](Watchlist.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWatchlistMultipleAccounts

> []Watchlist GetWatchlistMultipleAccounts(ctx, )

APIs to perform CRUD operations on Account Watchlist

All watchlists for all of the user's linked accounts.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Watchlist**](Watchlist.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWatchlistSingleAccount

> []Watchlist GetWatchlistSingleAccount(ctx, accountId)

APIs to perform CRUD operations on Account Watchlist

All watchlists of an account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 

### Return type

[**[]Watchlist**](Watchlist.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceWatchlist

> ReplaceWatchlist(ctx, accountId, watchlistId, optional)

APIs to perform CRUD operations on Account Watchlist

Replace watchlist for specific account. This method does not verify that the symbol or asset type are valid.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**watchlistId** | **string**| Watchlist ID | 
 **optional** | ***ReplaceWatchlistOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReplaceWatchlistOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateWatchlist** | [**optional.Interface of UpdateWatchlist**](UpdateWatchlist.md)|  | 

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


## UpdateWatchlist

> UpdateWatchlist(ctx, accountId, watchlistId, optional)

APIs to perform CRUD operations on Account Watchlist

Partially update watchlist for a specific account: change watchlist name, add to the beginning/end of a watchlist, update or delete items in a watchlist. This method does not verify that the symbol or asset type are valid.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**watchlistId** | **string**| Watchlist ID | 
 **optional** | ***UpdateWatchlistOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateWatchlistOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateWatchlist** | [**optional.Interface of UpdateWatchlist**](UpdateWatchlist.md)|  | 

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

