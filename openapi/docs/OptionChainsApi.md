# \OptionChainsApi

All URIs are relative to *https://api.tdameritrade.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOptionChains**](OptionChainsApi.md#GetOptionChains) | **Get** /marketdata/chains | Get Option Chains for optionable symbols



## GetOptionChains

> []map[string]interface{} GetOptionChains(ctx, apikey, symbol, contractType, strikeCount, includeQuotes, strategy)

Get Option Chains for optionable symbols

Get option chain for an optionable Symbol

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apikey** | **string**| API Key | 
**symbol** | **string**| Enter one symbol | 
**contractType** | [**PutOrCall**](.md)| Type of contracts to return in the chain. Can be CALL, PUT, or ALL. Default is ALL. | 
**strikeCount** | **string**| The number of strikes to return above and below the at-the-money price. | 
**includeQuotes** | **string**| Include quotes for options in the option chain. Can be TRUE or FALSE. Default is FALSE. | 
**strategy** | [**Strategy**](.md)| Passing a value returns a Strategy Chain. Possible values are SINGLE, ANALYTICAL (allows use of the volatility, underlyingPrice, interestRate, and daysToExpiration params to calculate theoretical values), COVERED, VERTICAL, CALENDAR, STRANGLE, STRADDLE, BUTTERFLY, CONDOR, DIAGONAL, COLLAR, or ROLL. Default is SINGLE. | 

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

