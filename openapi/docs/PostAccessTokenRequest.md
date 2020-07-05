# PostAccessTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrantType** | **string** | The grant type of the oAuth scheme. Possible values are authorization_code, refresh_token | [optional] [default to GRANT_TYPE_AUTHORIZATION_CODE]
**RefreshToken** | **string** | Required if using refresh token grant | [optional] 
**AccessType** | **string** | Set to offline to receive a refresh token on an authorization_code grant type request. Do not set to offline on a refresh_token grant type request. | [optional] 
**Code** | **string** | Required if trying to use authorization code grant | [optional] 
**ClientId** | **string** | OAuth User ID of your application | [optional] 
**RedirectUri** | **string** | Required if trying to use authorization code grant | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


