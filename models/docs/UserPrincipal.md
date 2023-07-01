# UserPrincipal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessLevel** | Pointer to **string** |  | [optional] 
**Accounts** | Pointer to [**[]UserPrincipalAccounts**](UserPrincipalAccounts.md) |  | [optional] 
**AuthToken** | Pointer to **string** |  | [optional] 
**LastLoginTime** | Pointer to **string** |  | [optional] 
**LoginTime** | Pointer to **string** |  | [optional] 
**PrimaryAccountId** | Pointer to **string** |  | [optional] 
**ProfessionalStatus** | Pointer to **string** |  | [optional] 
**Quotes** | Pointer to [**UserPrincipalQuotes**](UserPrincipal_quotes.md) |  | [optional] 
**StalePassword** | Pointer to **bool** |  | [optional] 
**StreamerInfo** | Pointer to [**UserPrincipalStreamerInfo**](UserPrincipal_streamerInfo.md) |  | [optional] 
**StreamerSubscriptionKeys** | Pointer to [**UserPrincipalStreamerSubscriptionKeys**](UserPrincipal_streamerSubscriptionKeys.md) |  | [optional] 
**TokenExpirationTime** | Pointer to **string** |  | [optional] 
**UserCdDomainId** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 

## Methods

### NewUserPrincipal

`func NewUserPrincipal() *UserPrincipal`

NewUserPrincipal instantiates a new UserPrincipal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPrincipalWithDefaults

`func NewUserPrincipalWithDefaults() *UserPrincipal`

NewUserPrincipalWithDefaults instantiates a new UserPrincipal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessLevel

`func (o *UserPrincipal) GetAccessLevel() string`

GetAccessLevel returns the AccessLevel field if non-nil, zero value otherwise.

### GetAccessLevelOk

`func (o *UserPrincipal) GetAccessLevelOk() (*string, bool)`

GetAccessLevelOk returns a tuple with the AccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevel

`func (o *UserPrincipal) SetAccessLevel(v string)`

SetAccessLevel sets AccessLevel field to given value.

### HasAccessLevel

`func (o *UserPrincipal) HasAccessLevel() bool`

HasAccessLevel returns a boolean if a field has been set.

### GetAccounts

`func (o *UserPrincipal) GetAccounts() []UserPrincipalAccounts`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *UserPrincipal) GetAccountsOk() (*[]UserPrincipalAccounts, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *UserPrincipal) SetAccounts(v []UserPrincipalAccounts)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *UserPrincipal) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetAuthToken

`func (o *UserPrincipal) GetAuthToken() string`

GetAuthToken returns the AuthToken field if non-nil, zero value otherwise.

### GetAuthTokenOk

`func (o *UserPrincipal) GetAuthTokenOk() (*string, bool)`

GetAuthTokenOk returns a tuple with the AuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthToken

`func (o *UserPrincipal) SetAuthToken(v string)`

SetAuthToken sets AuthToken field to given value.

### HasAuthToken

`func (o *UserPrincipal) HasAuthToken() bool`

HasAuthToken returns a boolean if a field has been set.

### GetLastLoginTime

`func (o *UserPrincipal) GetLastLoginTime() string`

GetLastLoginTime returns the LastLoginTime field if non-nil, zero value otherwise.

### GetLastLoginTimeOk

`func (o *UserPrincipal) GetLastLoginTimeOk() (*string, bool)`

GetLastLoginTimeOk returns a tuple with the LastLoginTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginTime

`func (o *UserPrincipal) SetLastLoginTime(v string)`

SetLastLoginTime sets LastLoginTime field to given value.

### HasLastLoginTime

`func (o *UserPrincipal) HasLastLoginTime() bool`

HasLastLoginTime returns a boolean if a field has been set.

### GetLoginTime

`func (o *UserPrincipal) GetLoginTime() string`

GetLoginTime returns the LoginTime field if non-nil, zero value otherwise.

### GetLoginTimeOk

`func (o *UserPrincipal) GetLoginTimeOk() (*string, bool)`

GetLoginTimeOk returns a tuple with the LoginTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginTime

`func (o *UserPrincipal) SetLoginTime(v string)`

SetLoginTime sets LoginTime field to given value.

### HasLoginTime

`func (o *UserPrincipal) HasLoginTime() bool`

HasLoginTime returns a boolean if a field has been set.

### GetPrimaryAccountId

`func (o *UserPrincipal) GetPrimaryAccountId() string`

GetPrimaryAccountId returns the PrimaryAccountId field if non-nil, zero value otherwise.

### GetPrimaryAccountIdOk

`func (o *UserPrincipal) GetPrimaryAccountIdOk() (*string, bool)`

GetPrimaryAccountIdOk returns a tuple with the PrimaryAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryAccountId

`func (o *UserPrincipal) SetPrimaryAccountId(v string)`

SetPrimaryAccountId sets PrimaryAccountId field to given value.

### HasPrimaryAccountId

`func (o *UserPrincipal) HasPrimaryAccountId() bool`

HasPrimaryAccountId returns a boolean if a field has been set.

### GetProfessionalStatus

`func (o *UserPrincipal) GetProfessionalStatus() string`

GetProfessionalStatus returns the ProfessionalStatus field if non-nil, zero value otherwise.

### GetProfessionalStatusOk

`func (o *UserPrincipal) GetProfessionalStatusOk() (*string, bool)`

GetProfessionalStatusOk returns a tuple with the ProfessionalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfessionalStatus

`func (o *UserPrincipal) SetProfessionalStatus(v string)`

SetProfessionalStatus sets ProfessionalStatus field to given value.

### HasProfessionalStatus

`func (o *UserPrincipal) HasProfessionalStatus() bool`

HasProfessionalStatus returns a boolean if a field has been set.

### GetQuotes

`func (o *UserPrincipal) GetQuotes() UserPrincipalQuotes`

GetQuotes returns the Quotes field if non-nil, zero value otherwise.

### GetQuotesOk

`func (o *UserPrincipal) GetQuotesOk() (*UserPrincipalQuotes, bool)`

GetQuotesOk returns a tuple with the Quotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotes

`func (o *UserPrincipal) SetQuotes(v UserPrincipalQuotes)`

SetQuotes sets Quotes field to given value.

### HasQuotes

`func (o *UserPrincipal) HasQuotes() bool`

HasQuotes returns a boolean if a field has been set.

### GetStalePassword

`func (o *UserPrincipal) GetStalePassword() bool`

GetStalePassword returns the StalePassword field if non-nil, zero value otherwise.

### GetStalePasswordOk

`func (o *UserPrincipal) GetStalePasswordOk() (*bool, bool)`

GetStalePasswordOk returns a tuple with the StalePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStalePassword

`func (o *UserPrincipal) SetStalePassword(v bool)`

SetStalePassword sets StalePassword field to given value.

### HasStalePassword

`func (o *UserPrincipal) HasStalePassword() bool`

HasStalePassword returns a boolean if a field has been set.

### GetStreamerInfo

`func (o *UserPrincipal) GetStreamerInfo() UserPrincipalStreamerInfo`

GetStreamerInfo returns the StreamerInfo field if non-nil, zero value otherwise.

### GetStreamerInfoOk

`func (o *UserPrincipal) GetStreamerInfoOk() (*UserPrincipalStreamerInfo, bool)`

GetStreamerInfoOk returns a tuple with the StreamerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamerInfo

`func (o *UserPrincipal) SetStreamerInfo(v UserPrincipalStreamerInfo)`

SetStreamerInfo sets StreamerInfo field to given value.

### HasStreamerInfo

`func (o *UserPrincipal) HasStreamerInfo() bool`

HasStreamerInfo returns a boolean if a field has been set.

### GetStreamerSubscriptionKeys

`func (o *UserPrincipal) GetStreamerSubscriptionKeys() UserPrincipalStreamerSubscriptionKeys`

GetStreamerSubscriptionKeys returns the StreamerSubscriptionKeys field if non-nil, zero value otherwise.

### GetStreamerSubscriptionKeysOk

`func (o *UserPrincipal) GetStreamerSubscriptionKeysOk() (*UserPrincipalStreamerSubscriptionKeys, bool)`

GetStreamerSubscriptionKeysOk returns a tuple with the StreamerSubscriptionKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamerSubscriptionKeys

`func (o *UserPrincipal) SetStreamerSubscriptionKeys(v UserPrincipalStreamerSubscriptionKeys)`

SetStreamerSubscriptionKeys sets StreamerSubscriptionKeys field to given value.

### HasStreamerSubscriptionKeys

`func (o *UserPrincipal) HasStreamerSubscriptionKeys() bool`

HasStreamerSubscriptionKeys returns a boolean if a field has been set.

### GetTokenExpirationTime

`func (o *UserPrincipal) GetTokenExpirationTime() string`

GetTokenExpirationTime returns the TokenExpirationTime field if non-nil, zero value otherwise.

### GetTokenExpirationTimeOk

`func (o *UserPrincipal) GetTokenExpirationTimeOk() (*string, bool)`

GetTokenExpirationTimeOk returns a tuple with the TokenExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExpirationTime

`func (o *UserPrincipal) SetTokenExpirationTime(v string)`

SetTokenExpirationTime sets TokenExpirationTime field to given value.

### HasTokenExpirationTime

`func (o *UserPrincipal) HasTokenExpirationTime() bool`

HasTokenExpirationTime returns a boolean if a field has been set.

### GetUserCdDomainId

`func (o *UserPrincipal) GetUserCdDomainId() string`

GetUserCdDomainId returns the UserCdDomainId field if non-nil, zero value otherwise.

### GetUserCdDomainIdOk

`func (o *UserPrincipal) GetUserCdDomainIdOk() (*string, bool)`

GetUserCdDomainIdOk returns a tuple with the UserCdDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCdDomainId

`func (o *UserPrincipal) SetUserCdDomainId(v string)`

SetUserCdDomainId sets UserCdDomainId field to given value.

### HasUserCdDomainId

`func (o *UserPrincipal) HasUserCdDomainId() bool`

HasUserCdDomainId returns a boolean if a field has been set.

### GetUserId

`func (o *UserPrincipal) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserPrincipal) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserPrincipal) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UserPrincipal) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


