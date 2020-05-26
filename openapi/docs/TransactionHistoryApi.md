# \TransactionHistoryApi

All URIs are relative to *https://api.tdameritrade.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTransaction**](TransactionHistoryApi.md#GetTransaction) | **Get** /accounts/{accountId}/transactions/{transactionId} | APIs to access transaction history on the account
[**GetTransactions**](TransactionHistoryApi.md#GetTransactions) | **Get** /accounts/{accountId}/transactions | APIs to access transaction history on the account



## GetTransaction

> Transaction GetTransaction(ctx, accountId, transactionId)

APIs to access transaction history on the account

Transaction for a specific account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**transactionId** | **string**| Transaction ID | 

### Return type

[**Transaction**](Transaction.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactions

> Transaction GetTransactions(ctx, accountId, type_, symbol, startDate, endDate)

APIs to access transaction history on the account

Transactions for a specific account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
**type_** | [**TransactionType**](.md)| Only transactions with the specified type will be returned. | 
**symbol** | **string**| Only transactions with the specified symbol will be returned. | 
**startDate** | **string**| Only transactions after the Start Date will be returned. Note: The maximum date range is one year. Valid ISO-8601 formats are: yyyy-MM-dd. | 
**endDate** | **string**| Only transactions before the End Date will be returned. Note: The maximum date range is one year. Valid ISO-8601 formats are: yyyy-MM-dd. | 

### Return type

[**Transaction**](Transaction.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

