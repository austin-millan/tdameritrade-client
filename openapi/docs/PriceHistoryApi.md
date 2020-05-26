# \PriceHistoryApi

All URIs are relative to *https://api.tdameritrade.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPriceHistory**](PriceHistoryApi.md#GetPriceHistory) | **Get** /marketdata/{symbol}/pricehistory | Historical price data for charts



## GetPriceHistory

> CandleList GetPriceHistory(ctx, apikey, symbol, periodType, period, frequencyType, frequency, endDate, startDate, needExtendedHoursData, x)

Historical price data for charts

Get price history for a symbol

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apikey** | **string**| API Key | 
**symbol** | **string**| Enter one symbol | 
**periodType** | **string**| The type of period to show. Valid values are day, month, year, or ytd (year to date). Default is day. | 
**period** | **int32**| The number of periods to show.  Example: For a 2 day / 1 min chart, the values would be: period: 2 periodType: day frequency: 1 frequencyType: min Valid periods by periodType (defaults marked with an asterisk): day: 1, 2, 3, 4, 5, 10* month: 1*, 2, 3, 6 year: 1*, 2, 3, 5, 10, 15, 20 ytd: 1* | 
**frequencyType** | **string**| The type of frequency with which a new candle is formed.  Valid frequencyTypes by periodType (defaults marked with an asterisk):  day: minute* month: daily, weekly* year: daily, weekly, monthly* ytd: daily, weekly*  | 
**frequency** | **int32**| The number of the frequencyType to be included in each candle.  Valid frequencies by frequencyType (defaults marked with an asterisk):  minute: 1*, 5, 10, 15, 30 daily: 1* weekly: 1* monthly: 1*  | 
**endDate** | **string**| End date as milliseconds since epoch. If startDate and endDate are provided, period should not be provided. Default is previous trading day. | 
**startDate** | **string**| Start date as milliseconds since epoch. If startDate and endDate are provided, period should not be provided. | 
**needExtendedHoursData** | **bool**| true to return extended hours data, false for regular market hours only. Default is true | 
**x** | **string**| x | 

### Return type

[**CandleList**](CandleList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

