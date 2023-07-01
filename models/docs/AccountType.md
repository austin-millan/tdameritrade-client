# AccountType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** |  | [optional] 
**CurrentBalances** | Pointer to [**MarginAccountCurrentBalances**](MarginAccount_currentBalances.md) |  | [optional] 
**InitialBalances** | Pointer to [**MarginAccountInitialBalances**](MarginAccount_initialBalances.md) |  | [optional] 
**IsClosingOnlyRestricted** | Pointer to **bool** |  | [optional] 
**IsDayTrader** | Pointer to **bool** |  | [optional] 
**OrderStrategies** | Pointer to [**[]MarginAccountOrderStrategies**](MarginAccountOrderStrategies.md) |  | [optional] 
**Positions** | Pointer to [**[]CashAccountPositions**](CashAccountPositions.md) |  | [optional] 
**ProjectedBalances** | Pointer to [**MarginAccountCurrentBalances**](MarginAccount_currentBalances.md) |  | [optional] 
**RoundTrips** | Pointer to **float32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewAccountType

`func NewAccountType() *AccountType`

NewAccountType instantiates a new AccountType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountTypeWithDefaults

`func NewAccountTypeWithDefaults() *AccountType`

NewAccountTypeWithDefaults instantiates a new AccountType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AccountType) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AccountType) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AccountType) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AccountType) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetCurrentBalances

`func (o *AccountType) GetCurrentBalances() MarginAccountCurrentBalances`

GetCurrentBalances returns the CurrentBalances field if non-nil, zero value otherwise.

### GetCurrentBalancesOk

`func (o *AccountType) GetCurrentBalancesOk() (*MarginAccountCurrentBalances, bool)`

GetCurrentBalancesOk returns a tuple with the CurrentBalances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentBalances

`func (o *AccountType) SetCurrentBalances(v MarginAccountCurrentBalances)`

SetCurrentBalances sets CurrentBalances field to given value.

### HasCurrentBalances

`func (o *AccountType) HasCurrentBalances() bool`

HasCurrentBalances returns a boolean if a field has been set.

### GetInitialBalances

`func (o *AccountType) GetInitialBalances() MarginAccountInitialBalances`

GetInitialBalances returns the InitialBalances field if non-nil, zero value otherwise.

### GetInitialBalancesOk

`func (o *AccountType) GetInitialBalancesOk() (*MarginAccountInitialBalances, bool)`

GetInitialBalancesOk returns a tuple with the InitialBalances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialBalances

`func (o *AccountType) SetInitialBalances(v MarginAccountInitialBalances)`

SetInitialBalances sets InitialBalances field to given value.

### HasInitialBalances

`func (o *AccountType) HasInitialBalances() bool`

HasInitialBalances returns a boolean if a field has been set.

### GetIsClosingOnlyRestricted

`func (o *AccountType) GetIsClosingOnlyRestricted() bool`

GetIsClosingOnlyRestricted returns the IsClosingOnlyRestricted field if non-nil, zero value otherwise.

### GetIsClosingOnlyRestrictedOk

`func (o *AccountType) GetIsClosingOnlyRestrictedOk() (*bool, bool)`

GetIsClosingOnlyRestrictedOk returns a tuple with the IsClosingOnlyRestricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClosingOnlyRestricted

`func (o *AccountType) SetIsClosingOnlyRestricted(v bool)`

SetIsClosingOnlyRestricted sets IsClosingOnlyRestricted field to given value.

### HasIsClosingOnlyRestricted

`func (o *AccountType) HasIsClosingOnlyRestricted() bool`

HasIsClosingOnlyRestricted returns a boolean if a field has been set.

### GetIsDayTrader

`func (o *AccountType) GetIsDayTrader() bool`

GetIsDayTrader returns the IsDayTrader field if non-nil, zero value otherwise.

### GetIsDayTraderOk

`func (o *AccountType) GetIsDayTraderOk() (*bool, bool)`

GetIsDayTraderOk returns a tuple with the IsDayTrader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDayTrader

`func (o *AccountType) SetIsDayTrader(v bool)`

SetIsDayTrader sets IsDayTrader field to given value.

### HasIsDayTrader

`func (o *AccountType) HasIsDayTrader() bool`

HasIsDayTrader returns a boolean if a field has been set.

### GetOrderStrategies

`func (o *AccountType) GetOrderStrategies() []MarginAccountOrderStrategies`

GetOrderStrategies returns the OrderStrategies field if non-nil, zero value otherwise.

### GetOrderStrategiesOk

`func (o *AccountType) GetOrderStrategiesOk() (*[]MarginAccountOrderStrategies, bool)`

GetOrderStrategiesOk returns a tuple with the OrderStrategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStrategies

`func (o *AccountType) SetOrderStrategies(v []MarginAccountOrderStrategies)`

SetOrderStrategies sets OrderStrategies field to given value.

### HasOrderStrategies

`func (o *AccountType) HasOrderStrategies() bool`

HasOrderStrategies returns a boolean if a field has been set.

### GetPositions

`func (o *AccountType) GetPositions() []CashAccountPositions`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *AccountType) GetPositionsOk() (*[]CashAccountPositions, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *AccountType) SetPositions(v []CashAccountPositions)`

SetPositions sets Positions field to given value.

### HasPositions

`func (o *AccountType) HasPositions() bool`

HasPositions returns a boolean if a field has been set.

### GetProjectedBalances

`func (o *AccountType) GetProjectedBalances() MarginAccountCurrentBalances`

GetProjectedBalances returns the ProjectedBalances field if non-nil, zero value otherwise.

### GetProjectedBalancesOk

`func (o *AccountType) GetProjectedBalancesOk() (*MarginAccountCurrentBalances, bool)`

GetProjectedBalancesOk returns a tuple with the ProjectedBalances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectedBalances

`func (o *AccountType) SetProjectedBalances(v MarginAccountCurrentBalances)`

SetProjectedBalances sets ProjectedBalances field to given value.

### HasProjectedBalances

`func (o *AccountType) HasProjectedBalances() bool`

HasProjectedBalances returns a boolean if a field has been set.

### GetRoundTrips

`func (o *AccountType) GetRoundTrips() float32`

GetRoundTrips returns the RoundTrips field if non-nil, zero value otherwise.

### GetRoundTripsOk

`func (o *AccountType) GetRoundTripsOk() (*float32, bool)`

GetRoundTripsOk returns a tuple with the RoundTrips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundTrips

`func (o *AccountType) SetRoundTrips(v float32)`

SetRoundTrips sets RoundTrips field to given value.

### HasRoundTrips

`func (o *AccountType) HasRoundTrips() bool`

HasRoundTrips returns a boolean if a field has been set.

### GetType

`func (o *AccountType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountType) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AccountType) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


