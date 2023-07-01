# MarginAccountOrderStrategies

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **float32** |  | [optional] 
**ActivationPrice** | Pointer to **float32** |  | [optional] 
**CancelTime** | Pointer to [**CashAccountCancelTime**](CashAccount_cancelTime.md) |  | [optional] 
**Cancelable** | Pointer to **bool** |  | [optional] 
**ChildOrderStrategies** | Pointer to **[]map[string]interface{}** |  | [optional] 
**CloseTime** | Pointer to **string** |  | [optional] 
**ComplexOrderStrategyType** | Pointer to [**ComplexOrderStrategyType**](ComplexOrderStrategyType.md) |  | [optional] 
**DestinationLinkName** | Pointer to **string** |  | [optional] 
**Duration** | Pointer to [**Duration**](Duration.md) |  | [optional] 
**Editable** | Pointer to **bool** |  | [optional] 
**EnteredTime** | Pointer to **string** |  | [optional] 
**FilledQuantity** | Pointer to **float32** |  | [optional] 
**OrderActivityCollection** | Pointer to [**Execution**](Execution.md) |  | [optional] 
**OrderId** | Pointer to **float32** |  | [optional] 
**OrderLegCollection** | Pointer to [**[]CashAccountOrderLegCollection**](CashAccountOrderLegCollection.md) |  | [optional] 
**OrderStrategyType** | Pointer to [**OrderStrategyType**](OrderStrategyType.md) |  | [optional] 
**OrderType** | Pointer to [**OrderType**](OrderType.md) |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**PriceLinkBasis** | Pointer to [**PriceLinkBasis**](PriceLinkBasis.md) |  | [optional] 
**PriceLinkType** | Pointer to [**PriceLinkType**](PriceLinkType.md) |  | [optional] 
**Quantity** | Pointer to **float32** |  | [optional] 
**ReleaseTime** | Pointer to **string** |  | [optional] 
**RemainingQuantity** | Pointer to **float32** |  | [optional] 
**ReplacingOrderCollection** | Pointer to **[]map[string]interface{}** |  | [optional] 
**RequestedDestination** | Pointer to [**DestinationExchange**](DestinationExchange.md) |  | [optional] 
**Session** | Pointer to [**Session**](Session.md) |  | [optional] 
**SpecialInstruction** | Pointer to [**SpecialInstruction**](SpecialInstruction.md) |  | [optional] 
**Status** | Pointer to [**OrderStatus**](OrderStatus.md) |  | [optional] 
**StatusDescription** | Pointer to **string** |  | [optional] 
**StopPrice** | Pointer to **float32** |  | [optional] 
**StopPriceLinkBasis** | Pointer to [**PriceLinkBasis**](PriceLinkBasis.md) |  | [optional] 
**StopPriceLinkType** | Pointer to [**PriceLinkType**](PriceLinkType.md) |  | [optional] 
**StopPriceOffset** | Pointer to **float32** |  | [optional] 
**StopType** | Pointer to [**StopType**](StopType.md) |  | [optional] 
**Tag** | Pointer to **string** |  | [optional] 
**TaxLotMethod** | Pointer to [**TaxLotMethod**](TaxLotMethod.md) |  | [optional] 

## Methods

### NewMarginAccountOrderStrategies

`func NewMarginAccountOrderStrategies() *MarginAccountOrderStrategies`

NewMarginAccountOrderStrategies instantiates a new MarginAccountOrderStrategies object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarginAccountOrderStrategiesWithDefaults

`func NewMarginAccountOrderStrategiesWithDefaults() *MarginAccountOrderStrategies`

NewMarginAccountOrderStrategiesWithDefaults instantiates a new MarginAccountOrderStrategies object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *MarginAccountOrderStrategies) GetAccountId() float32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *MarginAccountOrderStrategies) GetAccountIdOk() (*float32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *MarginAccountOrderStrategies) SetAccountId(v float32)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *MarginAccountOrderStrategies) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetActivationPrice

`func (o *MarginAccountOrderStrategies) GetActivationPrice() float32`

GetActivationPrice returns the ActivationPrice field if non-nil, zero value otherwise.

### GetActivationPriceOk

`func (o *MarginAccountOrderStrategies) GetActivationPriceOk() (*float32, bool)`

GetActivationPriceOk returns a tuple with the ActivationPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationPrice

`func (o *MarginAccountOrderStrategies) SetActivationPrice(v float32)`

SetActivationPrice sets ActivationPrice field to given value.

### HasActivationPrice

`func (o *MarginAccountOrderStrategies) HasActivationPrice() bool`

HasActivationPrice returns a boolean if a field has been set.

### GetCancelTime

`func (o *MarginAccountOrderStrategies) GetCancelTime() CashAccountCancelTime`

GetCancelTime returns the CancelTime field if non-nil, zero value otherwise.

### GetCancelTimeOk

`func (o *MarginAccountOrderStrategies) GetCancelTimeOk() (*CashAccountCancelTime, bool)`

GetCancelTimeOk returns a tuple with the CancelTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelTime

`func (o *MarginAccountOrderStrategies) SetCancelTime(v CashAccountCancelTime)`

SetCancelTime sets CancelTime field to given value.

### HasCancelTime

`func (o *MarginAccountOrderStrategies) HasCancelTime() bool`

HasCancelTime returns a boolean if a field has been set.

### GetCancelable

`func (o *MarginAccountOrderStrategies) GetCancelable() bool`

GetCancelable returns the Cancelable field if non-nil, zero value otherwise.

### GetCancelableOk

`func (o *MarginAccountOrderStrategies) GetCancelableOk() (*bool, bool)`

GetCancelableOk returns a tuple with the Cancelable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelable

`func (o *MarginAccountOrderStrategies) SetCancelable(v bool)`

SetCancelable sets Cancelable field to given value.

### HasCancelable

`func (o *MarginAccountOrderStrategies) HasCancelable() bool`

HasCancelable returns a boolean if a field has been set.

### GetChildOrderStrategies

`func (o *MarginAccountOrderStrategies) GetChildOrderStrategies() []map[string]interface{}`

GetChildOrderStrategies returns the ChildOrderStrategies field if non-nil, zero value otherwise.

### GetChildOrderStrategiesOk

`func (o *MarginAccountOrderStrategies) GetChildOrderStrategiesOk() (*[]map[string]interface{}, bool)`

GetChildOrderStrategiesOk returns a tuple with the ChildOrderStrategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildOrderStrategies

`func (o *MarginAccountOrderStrategies) SetChildOrderStrategies(v []map[string]interface{})`

SetChildOrderStrategies sets ChildOrderStrategies field to given value.

### HasChildOrderStrategies

`func (o *MarginAccountOrderStrategies) HasChildOrderStrategies() bool`

HasChildOrderStrategies returns a boolean if a field has been set.

### GetCloseTime

`func (o *MarginAccountOrderStrategies) GetCloseTime() string`

GetCloseTime returns the CloseTime field if non-nil, zero value otherwise.

### GetCloseTimeOk

`func (o *MarginAccountOrderStrategies) GetCloseTimeOk() (*string, bool)`

GetCloseTimeOk returns a tuple with the CloseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseTime

`func (o *MarginAccountOrderStrategies) SetCloseTime(v string)`

SetCloseTime sets CloseTime field to given value.

### HasCloseTime

`func (o *MarginAccountOrderStrategies) HasCloseTime() bool`

HasCloseTime returns a boolean if a field has been set.

### GetComplexOrderStrategyType

`func (o *MarginAccountOrderStrategies) GetComplexOrderStrategyType() ComplexOrderStrategyType`

GetComplexOrderStrategyType returns the ComplexOrderStrategyType field if non-nil, zero value otherwise.

### GetComplexOrderStrategyTypeOk

`func (o *MarginAccountOrderStrategies) GetComplexOrderStrategyTypeOk() (*ComplexOrderStrategyType, bool)`

GetComplexOrderStrategyTypeOk returns a tuple with the ComplexOrderStrategyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplexOrderStrategyType

`func (o *MarginAccountOrderStrategies) SetComplexOrderStrategyType(v ComplexOrderStrategyType)`

SetComplexOrderStrategyType sets ComplexOrderStrategyType field to given value.

### HasComplexOrderStrategyType

`func (o *MarginAccountOrderStrategies) HasComplexOrderStrategyType() bool`

HasComplexOrderStrategyType returns a boolean if a field has been set.

### GetDestinationLinkName

`func (o *MarginAccountOrderStrategies) GetDestinationLinkName() string`

GetDestinationLinkName returns the DestinationLinkName field if non-nil, zero value otherwise.

### GetDestinationLinkNameOk

`func (o *MarginAccountOrderStrategies) GetDestinationLinkNameOk() (*string, bool)`

GetDestinationLinkNameOk returns a tuple with the DestinationLinkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationLinkName

`func (o *MarginAccountOrderStrategies) SetDestinationLinkName(v string)`

SetDestinationLinkName sets DestinationLinkName field to given value.

### HasDestinationLinkName

`func (o *MarginAccountOrderStrategies) HasDestinationLinkName() bool`

HasDestinationLinkName returns a boolean if a field has been set.

### GetDuration

`func (o *MarginAccountOrderStrategies) GetDuration() Duration`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *MarginAccountOrderStrategies) GetDurationOk() (*Duration, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *MarginAccountOrderStrategies) SetDuration(v Duration)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *MarginAccountOrderStrategies) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEditable

`func (o *MarginAccountOrderStrategies) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *MarginAccountOrderStrategies) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *MarginAccountOrderStrategies) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *MarginAccountOrderStrategies) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetEnteredTime

`func (o *MarginAccountOrderStrategies) GetEnteredTime() string`

GetEnteredTime returns the EnteredTime field if non-nil, zero value otherwise.

### GetEnteredTimeOk

`func (o *MarginAccountOrderStrategies) GetEnteredTimeOk() (*string, bool)`

GetEnteredTimeOk returns a tuple with the EnteredTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnteredTime

`func (o *MarginAccountOrderStrategies) SetEnteredTime(v string)`

SetEnteredTime sets EnteredTime field to given value.

### HasEnteredTime

`func (o *MarginAccountOrderStrategies) HasEnteredTime() bool`

HasEnteredTime returns a boolean if a field has been set.

### GetFilledQuantity

`func (o *MarginAccountOrderStrategies) GetFilledQuantity() float32`

GetFilledQuantity returns the FilledQuantity field if non-nil, zero value otherwise.

### GetFilledQuantityOk

`func (o *MarginAccountOrderStrategies) GetFilledQuantityOk() (*float32, bool)`

GetFilledQuantityOk returns a tuple with the FilledQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilledQuantity

`func (o *MarginAccountOrderStrategies) SetFilledQuantity(v float32)`

SetFilledQuantity sets FilledQuantity field to given value.

### HasFilledQuantity

`func (o *MarginAccountOrderStrategies) HasFilledQuantity() bool`

HasFilledQuantity returns a boolean if a field has been set.

### GetOrderActivityCollection

`func (o *MarginAccountOrderStrategies) GetOrderActivityCollection() Execution`

GetOrderActivityCollection returns the OrderActivityCollection field if non-nil, zero value otherwise.

### GetOrderActivityCollectionOk

`func (o *MarginAccountOrderStrategies) GetOrderActivityCollectionOk() (*Execution, bool)`

GetOrderActivityCollectionOk returns a tuple with the OrderActivityCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderActivityCollection

`func (o *MarginAccountOrderStrategies) SetOrderActivityCollection(v Execution)`

SetOrderActivityCollection sets OrderActivityCollection field to given value.

### HasOrderActivityCollection

`func (o *MarginAccountOrderStrategies) HasOrderActivityCollection() bool`

HasOrderActivityCollection returns a boolean if a field has been set.

### GetOrderId

`func (o *MarginAccountOrderStrategies) GetOrderId() float32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *MarginAccountOrderStrategies) GetOrderIdOk() (*float32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *MarginAccountOrderStrategies) SetOrderId(v float32)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *MarginAccountOrderStrategies) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrderLegCollection

`func (o *MarginAccountOrderStrategies) GetOrderLegCollection() []CashAccountOrderLegCollection`

GetOrderLegCollection returns the OrderLegCollection field if non-nil, zero value otherwise.

### GetOrderLegCollectionOk

`func (o *MarginAccountOrderStrategies) GetOrderLegCollectionOk() (*[]CashAccountOrderLegCollection, bool)`

GetOrderLegCollectionOk returns a tuple with the OrderLegCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderLegCollection

`func (o *MarginAccountOrderStrategies) SetOrderLegCollection(v []CashAccountOrderLegCollection)`

SetOrderLegCollection sets OrderLegCollection field to given value.

### HasOrderLegCollection

`func (o *MarginAccountOrderStrategies) HasOrderLegCollection() bool`

HasOrderLegCollection returns a boolean if a field has been set.

### GetOrderStrategyType

`func (o *MarginAccountOrderStrategies) GetOrderStrategyType() OrderStrategyType`

GetOrderStrategyType returns the OrderStrategyType field if non-nil, zero value otherwise.

### GetOrderStrategyTypeOk

`func (o *MarginAccountOrderStrategies) GetOrderStrategyTypeOk() (*OrderStrategyType, bool)`

GetOrderStrategyTypeOk returns a tuple with the OrderStrategyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStrategyType

`func (o *MarginAccountOrderStrategies) SetOrderStrategyType(v OrderStrategyType)`

SetOrderStrategyType sets OrderStrategyType field to given value.

### HasOrderStrategyType

`func (o *MarginAccountOrderStrategies) HasOrderStrategyType() bool`

HasOrderStrategyType returns a boolean if a field has been set.

### GetOrderType

`func (o *MarginAccountOrderStrategies) GetOrderType() OrderType`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *MarginAccountOrderStrategies) GetOrderTypeOk() (*OrderType, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *MarginAccountOrderStrategies) SetOrderType(v OrderType)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *MarginAccountOrderStrategies) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetPrice

`func (o *MarginAccountOrderStrategies) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *MarginAccountOrderStrategies) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *MarginAccountOrderStrategies) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *MarginAccountOrderStrategies) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceLinkBasis

`func (o *MarginAccountOrderStrategies) GetPriceLinkBasis() PriceLinkBasis`

GetPriceLinkBasis returns the PriceLinkBasis field if non-nil, zero value otherwise.

### GetPriceLinkBasisOk

`func (o *MarginAccountOrderStrategies) GetPriceLinkBasisOk() (*PriceLinkBasis, bool)`

GetPriceLinkBasisOk returns a tuple with the PriceLinkBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceLinkBasis

`func (o *MarginAccountOrderStrategies) SetPriceLinkBasis(v PriceLinkBasis)`

SetPriceLinkBasis sets PriceLinkBasis field to given value.

### HasPriceLinkBasis

`func (o *MarginAccountOrderStrategies) HasPriceLinkBasis() bool`

HasPriceLinkBasis returns a boolean if a field has been set.

### GetPriceLinkType

`func (o *MarginAccountOrderStrategies) GetPriceLinkType() PriceLinkType`

GetPriceLinkType returns the PriceLinkType field if non-nil, zero value otherwise.

### GetPriceLinkTypeOk

`func (o *MarginAccountOrderStrategies) GetPriceLinkTypeOk() (*PriceLinkType, bool)`

GetPriceLinkTypeOk returns a tuple with the PriceLinkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceLinkType

`func (o *MarginAccountOrderStrategies) SetPriceLinkType(v PriceLinkType)`

SetPriceLinkType sets PriceLinkType field to given value.

### HasPriceLinkType

`func (o *MarginAccountOrderStrategies) HasPriceLinkType() bool`

HasPriceLinkType returns a boolean if a field has been set.

### GetQuantity

`func (o *MarginAccountOrderStrategies) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *MarginAccountOrderStrategies) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *MarginAccountOrderStrategies) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *MarginAccountOrderStrategies) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetReleaseTime

`func (o *MarginAccountOrderStrategies) GetReleaseTime() string`

GetReleaseTime returns the ReleaseTime field if non-nil, zero value otherwise.

### GetReleaseTimeOk

`func (o *MarginAccountOrderStrategies) GetReleaseTimeOk() (*string, bool)`

GetReleaseTimeOk returns a tuple with the ReleaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseTime

`func (o *MarginAccountOrderStrategies) SetReleaseTime(v string)`

SetReleaseTime sets ReleaseTime field to given value.

### HasReleaseTime

`func (o *MarginAccountOrderStrategies) HasReleaseTime() bool`

HasReleaseTime returns a boolean if a field has been set.

### GetRemainingQuantity

`func (o *MarginAccountOrderStrategies) GetRemainingQuantity() float32`

GetRemainingQuantity returns the RemainingQuantity field if non-nil, zero value otherwise.

### GetRemainingQuantityOk

`func (o *MarginAccountOrderStrategies) GetRemainingQuantityOk() (*float32, bool)`

GetRemainingQuantityOk returns a tuple with the RemainingQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingQuantity

`func (o *MarginAccountOrderStrategies) SetRemainingQuantity(v float32)`

SetRemainingQuantity sets RemainingQuantity field to given value.

### HasRemainingQuantity

`func (o *MarginAccountOrderStrategies) HasRemainingQuantity() bool`

HasRemainingQuantity returns a boolean if a field has been set.

### GetReplacingOrderCollection

`func (o *MarginAccountOrderStrategies) GetReplacingOrderCollection() []map[string]interface{}`

GetReplacingOrderCollection returns the ReplacingOrderCollection field if non-nil, zero value otherwise.

### GetReplacingOrderCollectionOk

`func (o *MarginAccountOrderStrategies) GetReplacingOrderCollectionOk() (*[]map[string]interface{}, bool)`

GetReplacingOrderCollectionOk returns a tuple with the ReplacingOrderCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacingOrderCollection

`func (o *MarginAccountOrderStrategies) SetReplacingOrderCollection(v []map[string]interface{})`

SetReplacingOrderCollection sets ReplacingOrderCollection field to given value.

### HasReplacingOrderCollection

`func (o *MarginAccountOrderStrategies) HasReplacingOrderCollection() bool`

HasReplacingOrderCollection returns a boolean if a field has been set.

### GetRequestedDestination

`func (o *MarginAccountOrderStrategies) GetRequestedDestination() DestinationExchange`

GetRequestedDestination returns the RequestedDestination field if non-nil, zero value otherwise.

### GetRequestedDestinationOk

`func (o *MarginAccountOrderStrategies) GetRequestedDestinationOk() (*DestinationExchange, bool)`

GetRequestedDestinationOk returns a tuple with the RequestedDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedDestination

`func (o *MarginAccountOrderStrategies) SetRequestedDestination(v DestinationExchange)`

SetRequestedDestination sets RequestedDestination field to given value.

### HasRequestedDestination

`func (o *MarginAccountOrderStrategies) HasRequestedDestination() bool`

HasRequestedDestination returns a boolean if a field has been set.

### GetSession

`func (o *MarginAccountOrderStrategies) GetSession() Session`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *MarginAccountOrderStrategies) GetSessionOk() (*Session, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *MarginAccountOrderStrategies) SetSession(v Session)`

SetSession sets Session field to given value.

### HasSession

`func (o *MarginAccountOrderStrategies) HasSession() bool`

HasSession returns a boolean if a field has been set.

### GetSpecialInstruction

`func (o *MarginAccountOrderStrategies) GetSpecialInstruction() SpecialInstruction`

GetSpecialInstruction returns the SpecialInstruction field if non-nil, zero value otherwise.

### GetSpecialInstructionOk

`func (o *MarginAccountOrderStrategies) GetSpecialInstructionOk() (*SpecialInstruction, bool)`

GetSpecialInstructionOk returns a tuple with the SpecialInstruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialInstruction

`func (o *MarginAccountOrderStrategies) SetSpecialInstruction(v SpecialInstruction)`

SetSpecialInstruction sets SpecialInstruction field to given value.

### HasSpecialInstruction

`func (o *MarginAccountOrderStrategies) HasSpecialInstruction() bool`

HasSpecialInstruction returns a boolean if a field has been set.

### GetStatus

`func (o *MarginAccountOrderStrategies) GetStatus() OrderStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MarginAccountOrderStrategies) GetStatusOk() (*OrderStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MarginAccountOrderStrategies) SetStatus(v OrderStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MarginAccountOrderStrategies) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDescription

`func (o *MarginAccountOrderStrategies) GetStatusDescription() string`

GetStatusDescription returns the StatusDescription field if non-nil, zero value otherwise.

### GetStatusDescriptionOk

`func (o *MarginAccountOrderStrategies) GetStatusDescriptionOk() (*string, bool)`

GetStatusDescriptionOk returns a tuple with the StatusDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDescription

`func (o *MarginAccountOrderStrategies) SetStatusDescription(v string)`

SetStatusDescription sets StatusDescription field to given value.

### HasStatusDescription

`func (o *MarginAccountOrderStrategies) HasStatusDescription() bool`

HasStatusDescription returns a boolean if a field has been set.

### GetStopPrice

`func (o *MarginAccountOrderStrategies) GetStopPrice() float32`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *MarginAccountOrderStrategies) GetStopPriceOk() (*float32, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *MarginAccountOrderStrategies) SetStopPrice(v float32)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *MarginAccountOrderStrategies) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.

### GetStopPriceLinkBasis

`func (o *MarginAccountOrderStrategies) GetStopPriceLinkBasis() PriceLinkBasis`

GetStopPriceLinkBasis returns the StopPriceLinkBasis field if non-nil, zero value otherwise.

### GetStopPriceLinkBasisOk

`func (o *MarginAccountOrderStrategies) GetStopPriceLinkBasisOk() (*PriceLinkBasis, bool)`

GetStopPriceLinkBasisOk returns a tuple with the StopPriceLinkBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPriceLinkBasis

`func (o *MarginAccountOrderStrategies) SetStopPriceLinkBasis(v PriceLinkBasis)`

SetStopPriceLinkBasis sets StopPriceLinkBasis field to given value.

### HasStopPriceLinkBasis

`func (o *MarginAccountOrderStrategies) HasStopPriceLinkBasis() bool`

HasStopPriceLinkBasis returns a boolean if a field has been set.

### GetStopPriceLinkType

`func (o *MarginAccountOrderStrategies) GetStopPriceLinkType() PriceLinkType`

GetStopPriceLinkType returns the StopPriceLinkType field if non-nil, zero value otherwise.

### GetStopPriceLinkTypeOk

`func (o *MarginAccountOrderStrategies) GetStopPriceLinkTypeOk() (*PriceLinkType, bool)`

GetStopPriceLinkTypeOk returns a tuple with the StopPriceLinkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPriceLinkType

`func (o *MarginAccountOrderStrategies) SetStopPriceLinkType(v PriceLinkType)`

SetStopPriceLinkType sets StopPriceLinkType field to given value.

### HasStopPriceLinkType

`func (o *MarginAccountOrderStrategies) HasStopPriceLinkType() bool`

HasStopPriceLinkType returns a boolean if a field has been set.

### GetStopPriceOffset

`func (o *MarginAccountOrderStrategies) GetStopPriceOffset() float32`

GetStopPriceOffset returns the StopPriceOffset field if non-nil, zero value otherwise.

### GetStopPriceOffsetOk

`func (o *MarginAccountOrderStrategies) GetStopPriceOffsetOk() (*float32, bool)`

GetStopPriceOffsetOk returns a tuple with the StopPriceOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPriceOffset

`func (o *MarginAccountOrderStrategies) SetStopPriceOffset(v float32)`

SetStopPriceOffset sets StopPriceOffset field to given value.

### HasStopPriceOffset

`func (o *MarginAccountOrderStrategies) HasStopPriceOffset() bool`

HasStopPriceOffset returns a boolean if a field has been set.

### GetStopType

`func (o *MarginAccountOrderStrategies) GetStopType() StopType`

GetStopType returns the StopType field if non-nil, zero value otherwise.

### GetStopTypeOk

`func (o *MarginAccountOrderStrategies) GetStopTypeOk() (*StopType, bool)`

GetStopTypeOk returns a tuple with the StopType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopType

`func (o *MarginAccountOrderStrategies) SetStopType(v StopType)`

SetStopType sets StopType field to given value.

### HasStopType

`func (o *MarginAccountOrderStrategies) HasStopType() bool`

HasStopType returns a boolean if a field has been set.

### GetTag

`func (o *MarginAccountOrderStrategies) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *MarginAccountOrderStrategies) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *MarginAccountOrderStrategies) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *MarginAccountOrderStrategies) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTaxLotMethod

`func (o *MarginAccountOrderStrategies) GetTaxLotMethod() TaxLotMethod`

GetTaxLotMethod returns the TaxLotMethod field if non-nil, zero value otherwise.

### GetTaxLotMethodOk

`func (o *MarginAccountOrderStrategies) GetTaxLotMethodOk() (*TaxLotMethod, bool)`

GetTaxLotMethodOk returns a tuple with the TaxLotMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxLotMethod

`func (o *MarginAccountOrderStrategies) SetTaxLotMethod(v TaxLotMethod)`

SetTaxLotMethod sets TaxLotMethod field to given value.

### HasTaxLotMethod

`func (o *MarginAccountOrderStrategies) HasTaxLotMethod() bool`

HasTaxLotMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


