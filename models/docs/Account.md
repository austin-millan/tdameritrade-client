# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecuritiesAccount** | Pointer to [**AccountType**](AccountType.md) |  | [optional] 

## Methods

### NewAccount

`func NewAccount() *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecuritiesAccount

`func (o *Account) GetSecuritiesAccount() AccountType`

GetSecuritiesAccount returns the SecuritiesAccount field if non-nil, zero value otherwise.

### GetSecuritiesAccountOk

`func (o *Account) GetSecuritiesAccountOk() (*AccountType, bool)`

GetSecuritiesAccountOk returns a tuple with the SecuritiesAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecuritiesAccount

`func (o *Account) SetSecuritiesAccount(v AccountType)`

SetSecuritiesAccount sets SecuritiesAccount field to given value.

### HasSecuritiesAccount

`func (o *Account) HasSecuritiesAccount() bool`

HasSecuritiesAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


