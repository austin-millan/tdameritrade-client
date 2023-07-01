# OptionChain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallExpDateMap** | Pointer to **string** |  | [optional] 
**DaysToExpiration** | Pointer to **float32** |  | [optional] 
**InterestRate** | Pointer to **float32** |  | [optional] 
**Interval** | Pointer to **float32** |  | [optional] 
**IsDelayed** | Pointer to **bool** |  | [optional] 
**IsIndex** | Pointer to **bool** |  | [optional] 
**PutExpDateMap** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**OrderStatus**](OrderStatus.md) |  | [optional] 
**Strategy** | Pointer to [**Strategy**](Strategy.md) |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Underlying** | Pointer to [**OptionChainUnderlying**](OptionChain_underlying.md) |  | [optional] 
**UnderlyingPrice** | Pointer to **float32** |  | [optional] 
**Volatility** | Pointer to **float32** |  | [optional] 

## Methods

### NewOptionChain

`func NewOptionChain() *OptionChain`

NewOptionChain instantiates a new OptionChain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionChainWithDefaults

`func NewOptionChainWithDefaults() *OptionChain`

NewOptionChainWithDefaults instantiates a new OptionChain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallExpDateMap

`func (o *OptionChain) GetCallExpDateMap() string`

GetCallExpDateMap returns the CallExpDateMap field if non-nil, zero value otherwise.

### GetCallExpDateMapOk

`func (o *OptionChain) GetCallExpDateMapOk() (*string, bool)`

GetCallExpDateMapOk returns a tuple with the CallExpDateMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallExpDateMap

`func (o *OptionChain) SetCallExpDateMap(v string)`

SetCallExpDateMap sets CallExpDateMap field to given value.

### HasCallExpDateMap

`func (o *OptionChain) HasCallExpDateMap() bool`

HasCallExpDateMap returns a boolean if a field has been set.

### GetDaysToExpiration

`func (o *OptionChain) GetDaysToExpiration() float32`

GetDaysToExpiration returns the DaysToExpiration field if non-nil, zero value otherwise.

### GetDaysToExpirationOk

`func (o *OptionChain) GetDaysToExpirationOk() (*float32, bool)`

GetDaysToExpirationOk returns a tuple with the DaysToExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysToExpiration

`func (o *OptionChain) SetDaysToExpiration(v float32)`

SetDaysToExpiration sets DaysToExpiration field to given value.

### HasDaysToExpiration

`func (o *OptionChain) HasDaysToExpiration() bool`

HasDaysToExpiration returns a boolean if a field has been set.

### GetInterestRate

`func (o *OptionChain) GetInterestRate() float32`

GetInterestRate returns the InterestRate field if non-nil, zero value otherwise.

### GetInterestRateOk

`func (o *OptionChain) GetInterestRateOk() (*float32, bool)`

GetInterestRateOk returns a tuple with the InterestRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestRate

`func (o *OptionChain) SetInterestRate(v float32)`

SetInterestRate sets InterestRate field to given value.

### HasInterestRate

`func (o *OptionChain) HasInterestRate() bool`

HasInterestRate returns a boolean if a field has been set.

### GetInterval

`func (o *OptionChain) GetInterval() float32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *OptionChain) GetIntervalOk() (*float32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *OptionChain) SetInterval(v float32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *OptionChain) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetIsDelayed

`func (o *OptionChain) GetIsDelayed() bool`

GetIsDelayed returns the IsDelayed field if non-nil, zero value otherwise.

### GetIsDelayedOk

`func (o *OptionChain) GetIsDelayedOk() (*bool, bool)`

GetIsDelayedOk returns a tuple with the IsDelayed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDelayed

`func (o *OptionChain) SetIsDelayed(v bool)`

SetIsDelayed sets IsDelayed field to given value.

### HasIsDelayed

`func (o *OptionChain) HasIsDelayed() bool`

HasIsDelayed returns a boolean if a field has been set.

### GetIsIndex

`func (o *OptionChain) GetIsIndex() bool`

GetIsIndex returns the IsIndex field if non-nil, zero value otherwise.

### GetIsIndexOk

`func (o *OptionChain) GetIsIndexOk() (*bool, bool)`

GetIsIndexOk returns a tuple with the IsIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIndex

`func (o *OptionChain) SetIsIndex(v bool)`

SetIsIndex sets IsIndex field to given value.

### HasIsIndex

`func (o *OptionChain) HasIsIndex() bool`

HasIsIndex returns a boolean if a field has been set.

### GetPutExpDateMap

`func (o *OptionChain) GetPutExpDateMap() string`

GetPutExpDateMap returns the PutExpDateMap field if non-nil, zero value otherwise.

### GetPutExpDateMapOk

`func (o *OptionChain) GetPutExpDateMapOk() (*string, bool)`

GetPutExpDateMapOk returns a tuple with the PutExpDateMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPutExpDateMap

`func (o *OptionChain) SetPutExpDateMap(v string)`

SetPutExpDateMap sets PutExpDateMap field to given value.

### HasPutExpDateMap

`func (o *OptionChain) HasPutExpDateMap() bool`

HasPutExpDateMap returns a boolean if a field has been set.

### GetStatus

`func (o *OptionChain) GetStatus() OrderStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OptionChain) GetStatusOk() (*OrderStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OptionChain) SetStatus(v OrderStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OptionChain) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStrategy

`func (o *OptionChain) GetStrategy() Strategy`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *OptionChain) GetStrategyOk() (*Strategy, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *OptionChain) SetStrategy(v Strategy)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *OptionChain) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### GetSymbol

`func (o *OptionChain) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *OptionChain) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *OptionChain) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *OptionChain) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetUnderlying

`func (o *OptionChain) GetUnderlying() OptionChainUnderlying`

GetUnderlying returns the Underlying field if non-nil, zero value otherwise.

### GetUnderlyingOk

`func (o *OptionChain) GetUnderlyingOk() (*OptionChainUnderlying, bool)`

GetUnderlyingOk returns a tuple with the Underlying field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderlying

`func (o *OptionChain) SetUnderlying(v OptionChainUnderlying)`

SetUnderlying sets Underlying field to given value.

### HasUnderlying

`func (o *OptionChain) HasUnderlying() bool`

HasUnderlying returns a boolean if a field has been set.

### GetUnderlyingPrice

`func (o *OptionChain) GetUnderlyingPrice() float32`

GetUnderlyingPrice returns the UnderlyingPrice field if non-nil, zero value otherwise.

### GetUnderlyingPriceOk

`func (o *OptionChain) GetUnderlyingPriceOk() (*float32, bool)`

GetUnderlyingPriceOk returns a tuple with the UnderlyingPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderlyingPrice

`func (o *OptionChain) SetUnderlyingPrice(v float32)`

SetUnderlyingPrice sets UnderlyingPrice field to given value.

### HasUnderlyingPrice

`func (o *OptionChain) HasUnderlyingPrice() bool`

HasUnderlyingPrice returns a boolean if a field has been set.

### GetVolatility

`func (o *OptionChain) GetVolatility() float32`

GetVolatility returns the Volatility field if non-nil, zero value otherwise.

### GetVolatilityOk

`func (o *OptionChain) GetVolatilityOk() (*float32, bool)`

GetVolatilityOk returns a tuple with the Volatility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolatility

`func (o *OptionChain) SetVolatility(v float32)`

SetVolatility sets Volatility field to given value.

### HasVolatility

`func (o *OptionChain) HasVolatility() bool`

HasVolatility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


