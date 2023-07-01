# OrderChildOrderStrategies

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CancelTime** | Pointer to [**CashAccountCancelTime**](CashAccount_cancelTime.md) |  | [optional] 
**ComplexOrderStrategyType** | Pointer to [**ComplexOrderStrategyType**](ComplexOrderStrategyType.md) |  | [optional] 
**DestinationLinkName** | Pointer to **string** |  | [optional] 
**Duration** | Pointer to [**Duration**](Duration.md) |  | [optional] 
**FilledQuantity** | Pointer to **float32** |  | [optional] 
**OrderType** | Pointer to [**OrderType**](OrderType.md) |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**PriceLinkBasis** | Pointer to [**PriceLinkBasis**](PriceLinkBasis.md) |  | [optional] 
**PriceLinkType** | Pointer to [**PriceLinkType**](PriceLinkType.md) |  | [optional] 
**Quantity** | Pointer to **float32** |  | [optional] 
**ReleaseTime** | Pointer to **string** |  | [optional] 
**RemainingQuantity** | Pointer to **float32** |  | [optional] 
**RequestedDestination** | Pointer to [**DestinationExchange**](DestinationExchange.md) |  | [optional] 
**Session** | Pointer to [**Session**](Session.md) |  | [optional] 
**StopPrice** | Pointer to **float32** |  | [optional] 
**StopPriceLinkBasis** | Pointer to [**PriceLinkBasis**](PriceLinkBasis.md) |  | [optional] 
**StopPriceLinkType** | Pointer to [**PriceLinkType**](PriceLinkType.md) |  | [optional] 
**StopPriceOffset** | Pointer to **float32** |  | [optional] 
**StopType** | Pointer to [**StopType**](StopType.md) |  | [optional] 
**TaxLotMethod** | Pointer to [**TaxLotMethod**](TaxLotMethod.md) |  | [optional] 

## Methods

### NewOrderChildOrderStrategies

`func NewOrderChildOrderStrategies() *OrderChildOrderStrategies`

NewOrderChildOrderStrategies instantiates a new OrderChildOrderStrategies object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderChildOrderStrategiesWithDefaults

`func NewOrderChildOrderStrategiesWithDefaults() *OrderChildOrderStrategies`

NewOrderChildOrderStrategiesWithDefaults instantiates a new OrderChildOrderStrategies object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancelTime

`func (o *OrderChildOrderStrategies) GetCancelTime() CashAccountCancelTime`

GetCancelTime returns the CancelTime field if non-nil, zero value otherwise.

### GetCancelTimeOk

`func (o *OrderChildOrderStrategies) GetCancelTimeOk() (*CashAccountCancelTime, bool)`

GetCancelTimeOk returns a tuple with the CancelTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelTime

`func (o *OrderChildOrderStrategies) SetCancelTime(v CashAccountCancelTime)`

SetCancelTime sets CancelTime field to given value.

### HasCancelTime

`func (o *OrderChildOrderStrategies) HasCancelTime() bool`

HasCancelTime returns a boolean if a field has been set.

### GetComplexOrderStrategyType

`func (o *OrderChildOrderStrategies) GetComplexOrderStrategyType() ComplexOrderStrategyType`

GetComplexOrderStrategyType returns the ComplexOrderStrategyType field if non-nil, zero value otherwise.

### GetComplexOrderStrategyTypeOk

`func (o *OrderChildOrderStrategies) GetComplexOrderStrategyTypeOk() (*ComplexOrderStrategyType, bool)`

GetComplexOrderStrategyTypeOk returns a tuple with the ComplexOrderStrategyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplexOrderStrategyType

`func (o *OrderChildOrderStrategies) SetComplexOrderStrategyType(v ComplexOrderStrategyType)`

SetComplexOrderStrategyType sets ComplexOrderStrategyType field to given value.

### HasComplexOrderStrategyType

`func (o *OrderChildOrderStrategies) HasComplexOrderStrategyType() bool`

HasComplexOrderStrategyType returns a boolean if a field has been set.

### GetDestinationLinkName

`func (o *OrderChildOrderStrategies) GetDestinationLinkName() string`

GetDestinationLinkName returns the DestinationLinkName field if non-nil, zero value otherwise.

### GetDestinationLinkNameOk

`func (o *OrderChildOrderStrategies) GetDestinationLinkNameOk() (*string, bool)`

GetDestinationLinkNameOk returns a tuple with the DestinationLinkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationLinkName

`func (o *OrderChildOrderStrategies) SetDestinationLinkName(v string)`

SetDestinationLinkName sets DestinationLinkName field to given value.

### HasDestinationLinkName

`func (o *OrderChildOrderStrategies) HasDestinationLinkName() bool`

HasDestinationLinkName returns a boolean if a field has been set.

### GetDuration

`func (o *OrderChildOrderStrategies) GetDuration() Duration`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *OrderChildOrderStrategies) GetDurationOk() (*Duration, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *OrderChildOrderStrategies) SetDuration(v Duration)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *OrderChildOrderStrategies) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetFilledQuantity

`func (o *OrderChildOrderStrategies) GetFilledQuantity() float32`

GetFilledQuantity returns the FilledQuantity field if non-nil, zero value otherwise.

### GetFilledQuantityOk

`func (o *OrderChildOrderStrategies) GetFilledQuantityOk() (*float32, bool)`

GetFilledQuantityOk returns a tuple with the FilledQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilledQuantity

`func (o *OrderChildOrderStrategies) SetFilledQuantity(v float32)`

SetFilledQuantity sets FilledQuantity field to given value.

### HasFilledQuantity

`func (o *OrderChildOrderStrategies) HasFilledQuantity() bool`

HasFilledQuantity returns a boolean if a field has been set.

### GetOrderType

`func (o *OrderChildOrderStrategies) GetOrderType() OrderType`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *OrderChildOrderStrategies) GetOrderTypeOk() (*OrderType, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *OrderChildOrderStrategies) SetOrderType(v OrderType)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *OrderChildOrderStrategies) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetPrice

`func (o *OrderChildOrderStrategies) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *OrderChildOrderStrategies) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *OrderChildOrderStrategies) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *OrderChildOrderStrategies) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceLinkBasis

`func (o *OrderChildOrderStrategies) GetPriceLinkBasis() PriceLinkBasis`

GetPriceLinkBasis returns the PriceLinkBasis field if non-nil, zero value otherwise.

### GetPriceLinkBasisOk

`func (o *OrderChildOrderStrategies) GetPriceLinkBasisOk() (*PriceLinkBasis, bool)`

GetPriceLinkBasisOk returns a tuple with the PriceLinkBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceLinkBasis

`func (o *OrderChildOrderStrategies) SetPriceLinkBasis(v PriceLinkBasis)`

SetPriceLinkBasis sets PriceLinkBasis field to given value.

### HasPriceLinkBasis

`func (o *OrderChildOrderStrategies) HasPriceLinkBasis() bool`

HasPriceLinkBasis returns a boolean if a field has been set.

### GetPriceLinkType

`func (o *OrderChildOrderStrategies) GetPriceLinkType() PriceLinkType`

GetPriceLinkType returns the PriceLinkType field if non-nil, zero value otherwise.

### GetPriceLinkTypeOk

`func (o *OrderChildOrderStrategies) GetPriceLinkTypeOk() (*PriceLinkType, bool)`

GetPriceLinkTypeOk returns a tuple with the PriceLinkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceLinkType

`func (o *OrderChildOrderStrategies) SetPriceLinkType(v PriceLinkType)`

SetPriceLinkType sets PriceLinkType field to given value.

### HasPriceLinkType

`func (o *OrderChildOrderStrategies) HasPriceLinkType() bool`

HasPriceLinkType returns a boolean if a field has been set.

### GetQuantity

`func (o *OrderChildOrderStrategies) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OrderChildOrderStrategies) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OrderChildOrderStrategies) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *OrderChildOrderStrategies) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetReleaseTime

`func (o *OrderChildOrderStrategies) GetReleaseTime() string`

GetReleaseTime returns the ReleaseTime field if non-nil, zero value otherwise.

### GetReleaseTimeOk

`func (o *OrderChildOrderStrategies) GetReleaseTimeOk() (*string, bool)`

GetReleaseTimeOk returns a tuple with the ReleaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseTime

`func (o *OrderChildOrderStrategies) SetReleaseTime(v string)`

SetReleaseTime sets ReleaseTime field to given value.

### HasReleaseTime

`func (o *OrderChildOrderStrategies) HasReleaseTime() bool`

HasReleaseTime returns a boolean if a field has been set.

### GetRemainingQuantity

`func (o *OrderChildOrderStrategies) GetRemainingQuantity() float32`

GetRemainingQuantity returns the RemainingQuantity field if non-nil, zero value otherwise.

### GetRemainingQuantityOk

`func (o *OrderChildOrderStrategies) GetRemainingQuantityOk() (*float32, bool)`

GetRemainingQuantityOk returns a tuple with the RemainingQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingQuantity

`func (o *OrderChildOrderStrategies) SetRemainingQuantity(v float32)`

SetRemainingQuantity sets RemainingQuantity field to given value.

### HasRemainingQuantity

`func (o *OrderChildOrderStrategies) HasRemainingQuantity() bool`

HasRemainingQuantity returns a boolean if a field has been set.

### GetRequestedDestination

`func (o *OrderChildOrderStrategies) GetRequestedDestination() DestinationExchange`

GetRequestedDestination returns the RequestedDestination field if non-nil, zero value otherwise.

### GetRequestedDestinationOk

`func (o *OrderChildOrderStrategies) GetRequestedDestinationOk() (*DestinationExchange, bool)`

GetRequestedDestinationOk returns a tuple with the RequestedDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedDestination

`func (o *OrderChildOrderStrategies) SetRequestedDestination(v DestinationExchange)`

SetRequestedDestination sets RequestedDestination field to given value.

### HasRequestedDestination

`func (o *OrderChildOrderStrategies) HasRequestedDestination() bool`

HasRequestedDestination returns a boolean if a field has been set.

### GetSession

`func (o *OrderChildOrderStrategies) GetSession() Session`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *OrderChildOrderStrategies) GetSessionOk() (*Session, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *OrderChildOrderStrategies) SetSession(v Session)`

SetSession sets Session field to given value.

### HasSession

`func (o *OrderChildOrderStrategies) HasSession() bool`

HasSession returns a boolean if a field has been set.

### GetStopPrice

`func (o *OrderChildOrderStrategies) GetStopPrice() float32`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *OrderChildOrderStrategies) GetStopPriceOk() (*float32, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *OrderChildOrderStrategies) SetStopPrice(v float32)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *OrderChildOrderStrategies) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.

### GetStopPriceLinkBasis

`func (o *OrderChildOrderStrategies) GetStopPriceLinkBasis() PriceLinkBasis`

GetStopPriceLinkBasis returns the StopPriceLinkBasis field if non-nil, zero value otherwise.

### GetStopPriceLinkBasisOk

`func (o *OrderChildOrderStrategies) GetStopPriceLinkBasisOk() (*PriceLinkBasis, bool)`

GetStopPriceLinkBasisOk returns a tuple with the StopPriceLinkBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPriceLinkBasis

`func (o *OrderChildOrderStrategies) SetStopPriceLinkBasis(v PriceLinkBasis)`

SetStopPriceLinkBasis sets StopPriceLinkBasis field to given value.

### HasStopPriceLinkBasis

`func (o *OrderChildOrderStrategies) HasStopPriceLinkBasis() bool`

HasStopPriceLinkBasis returns a boolean if a field has been set.

### GetStopPriceLinkType

`func (o *OrderChildOrderStrategies) GetStopPriceLinkType() PriceLinkType`

GetStopPriceLinkType returns the StopPriceLinkType field if non-nil, zero value otherwise.

### GetStopPriceLinkTypeOk

`func (o *OrderChildOrderStrategies) GetStopPriceLinkTypeOk() (*PriceLinkType, bool)`

GetStopPriceLinkTypeOk returns a tuple with the StopPriceLinkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPriceLinkType

`func (o *OrderChildOrderStrategies) SetStopPriceLinkType(v PriceLinkType)`

SetStopPriceLinkType sets StopPriceLinkType field to given value.

### HasStopPriceLinkType

`func (o *OrderChildOrderStrategies) HasStopPriceLinkType() bool`

HasStopPriceLinkType returns a boolean if a field has been set.

### GetStopPriceOffset

`func (o *OrderChildOrderStrategies) GetStopPriceOffset() float32`

GetStopPriceOffset returns the StopPriceOffset field if non-nil, zero value otherwise.

### GetStopPriceOffsetOk

`func (o *OrderChildOrderStrategies) GetStopPriceOffsetOk() (*float32, bool)`

GetStopPriceOffsetOk returns a tuple with the StopPriceOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPriceOffset

`func (o *OrderChildOrderStrategies) SetStopPriceOffset(v float32)`

SetStopPriceOffset sets StopPriceOffset field to given value.

### HasStopPriceOffset

`func (o *OrderChildOrderStrategies) HasStopPriceOffset() bool`

HasStopPriceOffset returns a boolean if a field has been set.

### GetStopType

`func (o *OrderChildOrderStrategies) GetStopType() StopType`

GetStopType returns the StopType field if non-nil, zero value otherwise.

### GetStopTypeOk

`func (o *OrderChildOrderStrategies) GetStopTypeOk() (*StopType, bool)`

GetStopTypeOk returns a tuple with the StopType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopType

`func (o *OrderChildOrderStrategies) SetStopType(v StopType)`

SetStopType sets StopType field to given value.

### HasStopType

`func (o *OrderChildOrderStrategies) HasStopType() bool`

HasStopType returns a boolean if a field has been set.

### GetTaxLotMethod

`func (o *OrderChildOrderStrategies) GetTaxLotMethod() TaxLotMethod`

GetTaxLotMethod returns the TaxLotMethod field if non-nil, zero value otherwise.

### GetTaxLotMethodOk

`func (o *OrderChildOrderStrategies) GetTaxLotMethodOk() (*TaxLotMethod, bool)`

GetTaxLotMethodOk returns a tuple with the TaxLotMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxLotMethod

`func (o *OrderChildOrderStrategies) SetTaxLotMethod(v TaxLotMethod)`

SetTaxLotMethod sets TaxLotMethod field to given value.

### HasTaxLotMethod

`func (o *OrderChildOrderStrategies) HasTaxLotMethod() bool`

HasTaxLotMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


