# \UserInfoPreferencesApi

All URIs are relative to *https://api.tdameritrade.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPreferences**](UserInfoPreferencesApi.md#GetPreferences) | **Get** /accounts/{accountId}/preferences | APIs to access user-authorized accounts and their preferences
[**GetStreamerSubscriptionKeys**](UserInfoPreferencesApi.md#GetStreamerSubscriptionKeys) | **Get** /userprincipals/streamersubscriptionkeys | APIs to access user-authorized accounts and their preferences
[**GetUserPrincipals**](UserInfoPreferencesApi.md#GetUserPrincipals) | **Get** /userprincipals | APIs to access user-authorized accounts and their preferences
[**UpdatePreferences**](UserInfoPreferencesApi.md#UpdatePreferences) | **Put** /accounts/{accountId}/preferences | APIs to access user-authorized accounts and their preferences



## GetPreferences

> Preferences GetPreferences(ctx, accountId)

APIs to access user-authorized accounts and their preferences

Preferences for a specific account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 

### Return type

[**Preferences**](Preferences.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStreamerSubscriptionKeys

> SubscriptionKey GetStreamerSubscriptionKeys(ctx, accountIds)

APIs to access user-authorized accounts and their preferences

SubscriptionKey for provided accounts or default accounts.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountIds** | [**[]string**](string.md)| Account IDs | 

### Return type

[**SubscriptionKey**](SubscriptionKey.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserPrincipals

> UserPrincipal GetUserPrincipals(ctx, fields)

APIs to access user-authorized accounts and their preferences

User Principal details.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fields** | [**[]string**](string.md)| A comma separated String which allows one to specify additional fields to return. None of these fields are returned by default. Possible values in this String can be: streamerSubscriptionKeys, streamerConnectionInfo, preferences, surrogateIds | 

### Return type

[**UserPrincipal**](UserPrincipal.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePreferences

> Preferences UpdatePreferences(ctx, accountId, optional)

APIs to access user-authorized accounts and their preferences

Update preferences for a specific account. Please note that the directOptionsRouting and directEquityRouting values cannot be modified via this operation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Account ID | 
 **optional** | ***UpdatePreferencesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdatePreferencesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePreferences** | [**optional.Interface of UpdatePreferences**](UpdatePreferences.md)|  | 

### Return type

[**Preferences**](Preferences.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

