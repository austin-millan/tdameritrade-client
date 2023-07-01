# CashAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** |  | [optional] 
**CurrentBalances** | Pointer to [**CashAccountCurrentBalances**](CashAccount_currentBalances.md) |  | [optional] 
**InitialBalances** | Pointer to [**CashAccountInitialBalances**](CashAccount_initialBalances.md) |  | [optional] 
**IsClosingOnlyRestricted** | Pointer to **bool** |  | [optional] 
**IsDayTrader** | Pointer to **bool** |  | [optional] 
**OrderStrategies** | Pointer to [**[]CashAccountOrderStrategies**](CashAccountOrderStrategies.md) |  | [optional] 
**Positions** | Pointer to [**[]CashAccountPositions**](CashAccountPositions.md) |  | [optional] 
**ProjectedBalances** | Pointer to [**CashAccountCurrentBalances**](CashAccount_currentBalances.md) |  | [optional] 
**RoundTrips** | Pointer to **float32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewCashAccount

`func NewCashAccount() *CashAccount`

NewCashAccount instantiates a new CashAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCashAccountWithDefaults

`func NewCashAccountWithDefaults() *CashAccount`

NewCashAccountWithDefaults instantiates a new CashAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *CashAccount) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CashAccount) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CashAccount) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *CashAccount) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetCurrentBalances

`func (o *CashAccount) GetCurrentBalances() CashAccountCurrentBalances`

GetCurrentBalances returns the CurrentBalances field if non-nil, zero value otherwise.

### GetCurrentBalancesOk

`func (o *CashAccount) GetCurrentBalancesOk() (*CashAccountCurrentBalances, bool)`

GetCurrentBalancesOk returns a tuple with the CurrentBalances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentBalances

`func (o *CashAccount) SetCurrentBalances(v CashAccountCurrentBalances)`

SetCurrentBalances sets CurrentBalances field to given value.

### HasCurrentBalances

`func (o *CashAccount) HasCurrentBalances() bool`

HasCurrentBalances returns a boolean if a field has been set.

### GetInitialBalances

`func (o *CashAccount) GetInitialBalances() CashAccountInitialBalances`

GetInitialBalances returns the InitialBalances field if non-nil, zero value otherwise.

### GetInitialBalancesOk

`func (o *CashAccount) GetInitialBalancesOk() (*CashAccountInitialBalances, bool)`

GetInitialBalancesOk returns a tuple with the InitialBalances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialBalances

`func (o *CashAccount) SetInitialBalances(v CashAccountInitialBalances)`

SetInitialBalances sets InitialBalances field to given value.

### HasInitialBalances

`func (o *CashAccount) HasInitialBalances() bool`

HasInitialBalances returns a boolean if a field has been set.

### GetIsClosingOnlyRestricted

`func (o *CashAccount) GetIsClosingOnlyRestricted() bool`

GetIsClosingOnlyRestricted returns the IsClosingOnlyRestricted field if non-nil, zero value otherwise.

### GetIsClosingOnlyRestrictedOk

`func (o *CashAccount) GetIsClosingOnlyRestrictedOk() (*bool, bool)`

GetIsClosingOnlyRestrictedOk returns a tuple with the IsClosingOnlyRestricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClosingOnlyRestricted

`func (o *CashAccount) SetIsClosingOnlyRestricted(v bool)`

SetIsClosingOnlyRestricted sets IsClosingOnlyRestricted field to given value.

### HasIsClosingOnlyRestricted

`func (o *CashAccount) HasIsClosingOnlyRestricted() bool`

HasIsClosingOnlyRestricted returns a boolean if a field has been set.

### GetIsDayTrader

`func (o *CashAccount) GetIsDayTrader() bool`

GetIsDayTrader returns the IsDayTrader field if non-nil, zero value otherwise.

### GetIsDayTraderOk

`func (o *CashAccount) GetIsDayTraderOk() (*bool, bool)`

GetIsDayTraderOk returns a tuple with the IsDayTrader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDayTrader

`func (o *CashAccount) SetIsDayTrader(v bool)`

SetIsDayTrader sets IsDayTrader field to given value.

### HasIsDayTrader

`func (o *CashAccount) HasIsDayTrader() bool`

HasIsDayTrader returns a boolean if a field has been set.

### GetOrderStrategies

`func (o *CashAccount) GetOrderStrategies() []CashAccountOrderStrategies`

GetOrderStrategies returns the OrderStrategies field if non-nil, zero value otherwise.

### GetOrderStrategiesOk

`func (o *CashAccount) GetOrderStrategiesOk() (*[]CashAccountOrderStrategies, bool)`

GetOrderStrategiesOk returns a tuple with the OrderStrategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStrategies

`func (o *CashAccount) SetOrderStrategies(v []CashAccountOrderStrategies)`

SetOrderStrategies sets OrderStrategies field to given value.

### HasOrderStrategies

`func (o *CashAccount) HasOrderStrategies() bool`

HasOrderStrategies returns a boolean if a field has been set.

### GetPositions

`func (o *CashAccount) GetPositions() []CashAccountPositions`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *CashAccount) GetPositionsOk() (*[]CashAccountPositions, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *CashAccount) SetPositions(v []CashAccountPositions)`

SetPositions sets Positions field to given value.

### HasPositions

`func (o *CashAccount) HasPositions() bool`

HasPositions returns a boolean if a field has been set.

### GetProjectedBalances

`func (o *CashAccount) GetProjectedBalances() CashAccountCurrentBalances`

GetProjectedBalances returns the ProjectedBalances field if non-nil, zero value otherwise.

### GetProjectedBalancesOk

`func (o *CashAccount) GetProjectedBalancesOk() (*CashAccountCurrentBalances, bool)`

GetProjectedBalancesOk returns a tuple with the ProjectedBalances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectedBalances

`func (o *CashAccount) SetProjectedBalances(v CashAccountCurrentBalances)`

SetProjectedBalances sets ProjectedBalances field to given value.

### HasProjectedBalances

`func (o *CashAccount) HasProjectedBalances() bool`

HasProjectedBalances returns a boolean if a field has been set.

### GetRoundTrips

`func (o *CashAccount) GetRoundTrips() float32`

GetRoundTrips returns the RoundTrips field if non-nil, zero value otherwise.

### GetRoundTripsOk

`func (o *CashAccount) GetRoundTripsOk() (*float32, bool)`

GetRoundTripsOk returns a tuple with the RoundTrips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundTrips

`func (o *CashAccount) SetRoundTrips(v float32)`

SetRoundTrips sets RoundTrips field to given value.

### HasRoundTrips

`func (o *CashAccount) HasRoundTrips() bool`

HasRoundTrips returns a boolean if a field has been set.

### GetType

`func (o *CashAccount) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CashAccount) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CashAccount) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CashAccount) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


