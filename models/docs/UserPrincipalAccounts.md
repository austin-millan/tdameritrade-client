# UserPrincipalAccounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountCdDomainId** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**Acl** | Pointer to **string** |  | [optional] 
**Authorizations** | Pointer to [**UserPrincipalAuthorizations**](UserPrincipal_authorizations.md) |  | [optional] 
**Company** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Preferences** | Pointer to [**UserPrincipalPreferences**](UserPrincipal_preferences.md) |  | [optional] 
**Segment** | Pointer to **string** |  | [optional] 
**SurrogateIds** | Pointer to **string** |  | [optional] 

## Methods

### NewUserPrincipalAccounts

`func NewUserPrincipalAccounts() *UserPrincipalAccounts`

NewUserPrincipalAccounts instantiates a new UserPrincipalAccounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPrincipalAccountsWithDefaults

`func NewUserPrincipalAccountsWithDefaults() *UserPrincipalAccounts`

NewUserPrincipalAccountsWithDefaults instantiates a new UserPrincipalAccounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountCdDomainId

`func (o *UserPrincipalAccounts) GetAccountCdDomainId() string`

GetAccountCdDomainId returns the AccountCdDomainId field if non-nil, zero value otherwise.

### GetAccountCdDomainIdOk

`func (o *UserPrincipalAccounts) GetAccountCdDomainIdOk() (*string, bool)`

GetAccountCdDomainIdOk returns a tuple with the AccountCdDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCdDomainId

`func (o *UserPrincipalAccounts) SetAccountCdDomainId(v string)`

SetAccountCdDomainId sets AccountCdDomainId field to given value.

### HasAccountCdDomainId

`func (o *UserPrincipalAccounts) HasAccountCdDomainId() bool`

HasAccountCdDomainId returns a boolean if a field has been set.

### GetAccountId

`func (o *UserPrincipalAccounts) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UserPrincipalAccounts) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UserPrincipalAccounts) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *UserPrincipalAccounts) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAcl

`func (o *UserPrincipalAccounts) GetAcl() string`

GetAcl returns the Acl field if non-nil, zero value otherwise.

### GetAclOk

`func (o *UserPrincipalAccounts) GetAclOk() (*string, bool)`

GetAclOk returns a tuple with the Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcl

`func (o *UserPrincipalAccounts) SetAcl(v string)`

SetAcl sets Acl field to given value.

### HasAcl

`func (o *UserPrincipalAccounts) HasAcl() bool`

HasAcl returns a boolean if a field has been set.

### GetAuthorizations

`func (o *UserPrincipalAccounts) GetAuthorizations() UserPrincipalAuthorizations`

GetAuthorizations returns the Authorizations field if non-nil, zero value otherwise.

### GetAuthorizationsOk

`func (o *UserPrincipalAccounts) GetAuthorizationsOk() (*UserPrincipalAuthorizations, bool)`

GetAuthorizationsOk returns a tuple with the Authorizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizations

`func (o *UserPrincipalAccounts) SetAuthorizations(v UserPrincipalAuthorizations)`

SetAuthorizations sets Authorizations field to given value.

### HasAuthorizations

`func (o *UserPrincipalAccounts) HasAuthorizations() bool`

HasAuthorizations returns a boolean if a field has been set.

### GetCompany

`func (o *UserPrincipalAccounts) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *UserPrincipalAccounts) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *UserPrincipalAccounts) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *UserPrincipalAccounts) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetDescription

`func (o *UserPrincipalAccounts) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserPrincipalAccounts) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserPrincipalAccounts) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UserPrincipalAccounts) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *UserPrincipalAccounts) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UserPrincipalAccounts) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UserPrincipalAccounts) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UserPrincipalAccounts) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetPreferences

`func (o *UserPrincipalAccounts) GetPreferences() UserPrincipalPreferences`

GetPreferences returns the Preferences field if non-nil, zero value otherwise.

### GetPreferencesOk

`func (o *UserPrincipalAccounts) GetPreferencesOk() (*UserPrincipalPreferences, bool)`

GetPreferencesOk returns a tuple with the Preferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferences

`func (o *UserPrincipalAccounts) SetPreferences(v UserPrincipalPreferences)`

SetPreferences sets Preferences field to given value.

### HasPreferences

`func (o *UserPrincipalAccounts) HasPreferences() bool`

HasPreferences returns a boolean if a field has been set.

### GetSegment

`func (o *UserPrincipalAccounts) GetSegment() string`

GetSegment returns the Segment field if non-nil, zero value otherwise.

### GetSegmentOk

`func (o *UserPrincipalAccounts) GetSegmentOk() (*string, bool)`

GetSegmentOk returns a tuple with the Segment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment

`func (o *UserPrincipalAccounts) SetSegment(v string)`

SetSegment sets Segment field to given value.

### HasSegment

`func (o *UserPrincipalAccounts) HasSegment() bool`

HasSegment returns a boolean if a field has been set.

### GetSurrogateIds

`func (o *UserPrincipalAccounts) GetSurrogateIds() string`

GetSurrogateIds returns the SurrogateIds field if non-nil, zero value otherwise.

### GetSurrogateIdsOk

`func (o *UserPrincipalAccounts) GetSurrogateIdsOk() (*string, bool)`

GetSurrogateIdsOk returns a tuple with the SurrogateIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurrogateIds

`func (o *UserPrincipalAccounts) SetSurrogateIds(v string)`

SetSurrogateIds sets SurrogateIds field to given value.

### HasSurrogateIds

`func (o *UserPrincipalAccounts) HasSurrogateIds() bool`

HasSurrogateIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


