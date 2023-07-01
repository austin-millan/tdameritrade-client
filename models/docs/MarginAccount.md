# MarginAccount

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

### NewMarginAccount

`func NewMarginAccount() *MarginAccount`

NewMarginAccount instantiates a new MarginAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarginAccountWithDefaults

`func NewMarginAccountWithDefaults() *MarginAccount`

NewMarginAccountWithDefaults instantiates a new MarginAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *MarginAccount) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *MarginAccount) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *MarginAccount) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *MarginAccount) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetCurrentBalances

`func (o *MarginAccount) GetCurrentBalances() MarginAccountCurrentBalances`

GetCurrentBalances returns the CurrentBalances field if non-nil, zero value otherwise.

### GetCurrentBalancesOk

`func (o *MarginAccount) GetCurrentBalancesOk() (*MarginAccountCurrentBalances, bool)`

GetCurrentBalancesOk returns a tuple with the CurrentBalances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentBalances

`func (o *MarginAccount) SetCurrentBalances(v MarginAccountCurrentBalances)`

SetCurrentBalances sets CurrentBalances field to given value.

### HasCurrentBalances

`func (o *MarginAccount) HasCurrentBalances() bool`

HasCurrentBalances returns a boolean if a field has been set.

### GetInitialBalances

`func (o *MarginAccount) GetInitialBalances() MarginAccountInitialBalances`

GetInitialBalances returns the InitialBalances field if non-nil, zero value otherwise.

### GetInitialBalancesOk

`func (o *MarginAccount) GetInitialBalancesOk() (*MarginAccountInitialBalances, bool)`

GetInitialBalancesOk returns a tuple with the InitialBalances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialBalances

`func (o *MarginAccount) SetInitialBalances(v MarginAccountInitialBalances)`

SetInitialBalances sets InitialBalances field to given value.

### HasInitialBalances

`func (o *MarginAccount) HasInitialBalances() bool`

HasInitialBalances returns a boolean if a field has been set.

### GetIsClosingOnlyRestricted

`func (o *MarginAccount) GetIsClosingOnlyRestricted() bool`

GetIsClosingOnlyRestricted returns the IsClosingOnlyRestricted field if non-nil, zero value otherwise.

### GetIsClosingOnlyRestrictedOk

`func (o *MarginAccount) GetIsClosingOnlyRestrictedOk() (*bool, bool)`

GetIsClosingOnlyRestrictedOk returns a tuple with the IsClosingOnlyRestricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClosingOnlyRestricted

`func (o *MarginAccount) SetIsClosingOnlyRestricted(v bool)`

SetIsClosingOnlyRestricted sets IsClosingOnlyRestricted field to given value.

### HasIsClosingOnlyRestricted

`func (o *MarginAccount) HasIsClosingOnlyRestricted() bool`

HasIsClosingOnlyRestricted returns a boolean if a field has been set.

### GetIsDayTrader

`func (o *MarginAccount) GetIsDayTrader() bool`

GetIsDayTrader returns the IsDayTrader field if non-nil, zero value otherwise.

### GetIsDayTraderOk

`func (o *MarginAccount) GetIsDayTraderOk() (*bool, bool)`

GetIsDayTraderOk returns a tuple with the IsDayTrader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDayTrader

`func (o *MarginAccount) SetIsDayTrader(v bool)`

SetIsDayTrader sets IsDayTrader field to given value.

### HasIsDayTrader

`func (o *MarginAccount) HasIsDayTrader() bool`

HasIsDayTrader returns a boolean if a field has been set.

### GetOrderStrategies

`func (o *MarginAccount) GetOrderStrategies() []MarginAccountOrderStrategies`

GetOrderStrategies returns the OrderStrategies field if non-nil, zero value otherwise.

### GetOrderStrategiesOk

`func (o *MarginAccount) GetOrderStrategiesOk() (*[]MarginAccountOrderStrategies, bool)`

GetOrderStrategiesOk returns a tuple with the OrderStrategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStrategies

`func (o *MarginAccount) SetOrderStrategies(v []MarginAccountOrderStrategies)`

SetOrderStrategies sets OrderStrategies field to given value.

### HasOrderStrategies

`func (o *MarginAccount) HasOrderStrategies() bool`

HasOrderStrategies returns a boolean if a field has been set.

### GetPositions

`func (o *MarginAccount) GetPositions() []CashAccountPositions`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *MarginAccount) GetPositionsOk() (*[]CashAccountPositions, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *MarginAccount) SetPositions(v []CashAccountPositions)`

SetPositions sets Positions field to given value.

### HasPositions

`func (o *MarginAccount) HasPositions() bool`

HasPositions returns a boolean if a field has been set.

### GetProjectedBalances

`func (o *MarginAccount) GetProjectedBalances() MarginAccountCurrentBalances`

GetProjectedBalances returns the ProjectedBalances field if non-nil, zero value otherwise.

### GetProjectedBalancesOk

`func (o *MarginAccount) GetProjectedBalancesOk() (*MarginAccountCurrentBalances, bool)`

GetProjectedBalancesOk returns a tuple with the ProjectedBalances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectedBalances

`func (o *MarginAccount) SetProjectedBalances(v MarginAccountCurrentBalances)`

SetProjectedBalances sets ProjectedBalances field to given value.

### HasProjectedBalances

`func (o *MarginAccount) HasProjectedBalances() bool`

HasProjectedBalances returns a boolean if a field has been set.

### GetRoundTrips

`func (o *MarginAccount) GetRoundTrips() float32`

GetRoundTrips returns the RoundTrips field if non-nil, zero value otherwise.

### GetRoundTripsOk

`func (o *MarginAccount) GetRoundTripsOk() (*float32, bool)`

GetRoundTripsOk returns a tuple with the RoundTrips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundTrips

`func (o *MarginAccount) SetRoundTrips(v float32)`

SetRoundTrips sets RoundTrips field to given value.

### HasRoundTrips

`func (o *MarginAccount) HasRoundTrips() bool`

HasRoundTrips returns a boolean if a field has been set.

### GetType

`func (o *MarginAccount) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MarginAccount) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MarginAccount) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MarginAccount) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


