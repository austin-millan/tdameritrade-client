# CashAccountOrderStrategies

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
**OrderActivityCollection** | Pointer to [**[]Execution**](Execution.md) |  | [optional] 
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

### NewCashAccountOrderStrategies

`func NewCashAccountOrderStrategies() *CashAccountOrderStrategies`

NewCashAccountOrderStrategies instantiates a new CashAccountOrderStrategies object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCashAccountOrderStrategiesWithDefaults

`func NewCashAccountOrderStrategiesWithDefaults() *CashAccountOrderStrategies`

NewCashAccountOrderStrategiesWithDefaults instantiates a new CashAccountOrderStrategies object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *CashAccountOrderStrategies) GetAccountId() float32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CashAccountOrderStrategies) GetAccountIdOk() (*float32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CashAccountOrderStrategies) SetAccountId(v float32)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *CashAccountOrderStrategies) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetActivationPrice

`func (o *CashAccountOrderStrategies) GetActivationPrice() float32`

GetActivationPrice returns the ActivationPrice field if non-nil, zero value otherwise.

### GetActivationPriceOk

`func (o *CashAccountOrderStrategies) GetActivationPriceOk() (*float32, bool)`

GetActivationPriceOk returns a tuple with the ActivationPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationPrice

`func (o *CashAccountOrderStrategies) SetActivationPrice(v float32)`

SetActivationPrice sets ActivationPrice field to given value.

### HasActivationPrice

`func (o *CashAccountOrderStrategies) HasActivationPrice() bool`

HasActivationPrice returns a boolean if a field has been set.

### GetCancelTime

`func (o *CashAccountOrderStrategies) GetCancelTime() CashAccountCancelTime`

GetCancelTime returns the CancelTime field if non-nil, zero value otherwise.

### GetCancelTimeOk

`func (o *CashAccountOrderStrategies) GetCancelTimeOk() (*CashAccountCancelTime, bool)`

GetCancelTimeOk returns a tuple with the CancelTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelTime

`func (o *CashAccountOrderStrategies) SetCancelTime(v CashAccountCancelTime)`

SetCancelTime sets CancelTime field to given value.

### HasCancelTime

`func (o *CashAccountOrderStrategies) HasCancelTime() bool`

HasCancelTime returns a boolean if a field has been set.

### GetCancelable

`func (o *CashAccountOrderStrategies) GetCancelable() bool`

GetCancelable returns the Cancelable field if non-nil, zero value otherwise.

### GetCancelableOk

`func (o *CashAccountOrderStrategies) GetCancelableOk() (*bool, bool)`

GetCancelableOk returns a tuple with the Cancelable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelable

`func (o *CashAccountOrderStrategies) SetCancelable(v bool)`

SetCancelable sets Cancelable field to given value.

### HasCancelable

`func (o *CashAccountOrderStrategies) HasCancelable() bool`

HasCancelable returns a boolean if a field has been set.

### GetChildOrderStrategies

`func (o *CashAccountOrderStrategies) GetChildOrderStrategies() []map[string]interface{}`

GetChildOrderStrategies returns the ChildOrderStrategies field if non-nil, zero value otherwise.

### GetChildOrderStrategiesOk

`func (o *CashAccountOrderStrategies) GetChildOrderStrategiesOk() (*[]map[string]interface{}, bool)`

GetChildOrderStrategiesOk returns a tuple with the ChildOrderStrategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildOrderStrategies

`func (o *CashAccountOrderStrategies) SetChildOrderStrategies(v []map[string]interface{})`

SetChildOrderStrategies sets ChildOrderStrategies field to given value.

### HasChildOrderStrategies

`func (o *CashAccountOrderStrategies) HasChildOrderStrategies() bool`

HasChildOrderStrategies returns a boolean if a field has been set.

### GetCloseTime

`func (o *CashAccountOrderStrategies) GetCloseTime() string`

GetCloseTime returns the CloseTime field if non-nil, zero value otherwise.

### GetCloseTimeOk

`func (o *CashAccountOrderStrategies) GetCloseTimeOk() (*string, bool)`

GetCloseTimeOk returns a tuple with the CloseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseTime

`func (o *CashAccountOrderStrategies) SetCloseTime(v string)`

SetCloseTime sets CloseTime field to given value.

### HasCloseTime

`func (o *CashAccountOrderStrategies) HasCloseTime() bool`

HasCloseTime returns a boolean if a field has been set.

### GetComplexOrderStrategyType

`func (o *CashAccountOrderStrategies) GetComplexOrderStrategyType() ComplexOrderStrategyType`

GetComplexOrderStrategyType returns the ComplexOrderStrategyType field if non-nil, zero value otherwise.

### GetComplexOrderStrategyTypeOk

`func (o *CashAccountOrderStrategies) GetComplexOrderStrategyTypeOk() (*ComplexOrderStrategyType, bool)`

GetComplexOrderStrategyTypeOk returns a tuple with the ComplexOrderStrategyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplexOrderStrategyType

`func (o *CashAccountOrderStrategies) SetComplexOrderStrategyType(v ComplexOrderStrategyType)`

SetComplexOrderStrategyType sets ComplexOrderStrategyType field to given value.

### HasComplexOrderStrategyType

`func (o *CashAccountOrderStrategies) HasComplexOrderStrategyType() bool`

HasComplexOrderStrategyType returns a boolean if a field has been set.

### GetDestinationLinkName

`func (o *CashAccountOrderStrategies) GetDestinationLinkName() string`

GetDestinationLinkName returns the DestinationLinkName field if non-nil, zero value otherwise.

### GetDestinationLinkNameOk

`func (o *CashAccountOrderStrategies) GetDestinationLinkNameOk() (*string, bool)`

GetDestinationLinkNameOk returns a tuple with the DestinationLinkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationLinkName

`func (o *CashAccountOrderStrategies) SetDestinationLinkName(v string)`

SetDestinationLinkName sets DestinationLinkName field to given value.

### HasDestinationLinkName

`func (o *CashAccountOrderStrategies) HasDestinationLinkName() bool`

HasDestinationLinkName returns a boolean if a field has been set.

### GetDuration

`func (o *CashAccountOrderStrategies) GetDuration() Duration`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *CashAccountOrderStrategies) GetDurationOk() (*Duration, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *CashAccountOrderStrategies) SetDuration(v Duration)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *CashAccountOrderStrategies) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEditable

`func (o *CashAccountOrderStrategies) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *CashAccountOrderStrategies) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *CashAccountOrderStrategies) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *CashAccountOrderStrategies) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetEnteredTime

`func (o *CashAccountOrderStrategies) GetEnteredTime() string`

GetEnteredTime returns the EnteredTime field if non-nil, zero value otherwise.

### GetEnteredTimeOk

`func (o *CashAccountOrderStrategies) GetEnteredTimeOk() (*string, bool)`

GetEnteredTimeOk returns a tuple with the EnteredTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnteredTime

`func (o *CashAccountOrderStrategies) SetEnteredTime(v string)`

SetEnteredTime sets EnteredTime field to given value.

### HasEnteredTime

`func (o *CashAccountOrderStrategies) HasEnteredTime() bool`

HasEnteredTime returns a boolean if a field has been set.

### GetFilledQuantity

`func (o *CashAccountOrderStrategies) GetFilledQuantity() float32`

GetFilledQuantity returns the FilledQuantity field if non-nil, zero value otherwise.

### GetFilledQuantityOk

`func (o *CashAccountOrderStrategies) GetFilledQuantityOk() (*float32, bool)`

GetFilledQuantityOk returns a tuple with the FilledQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilledQuantity

`func (o *CashAccountOrderStrategies) SetFilledQuantity(v float32)`

SetFilledQuantity sets FilledQuantity field to given value.

### HasFilledQuantity

`func (o *CashAccountOrderStrategies) HasFilledQuantity() bool`

HasFilledQuantity returns a boolean if a field has been set.

### GetOrderActivityCollection

`func (o *CashAccountOrderStrategies) GetOrderActivityCollection() []Execution`

GetOrderActivityCollection returns the OrderActivityCollection field if non-nil, zero value otherwise.

### GetOrderActivityCollectionOk

`func (o *CashAccountOrderStrategies) GetOrderActivityCollectionOk() (*[]Execution, bool)`

GetOrderActivityCollectionOk returns a tuple with the OrderActivityCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderActivityCollection

`func (o *CashAccountOrderStrategies) SetOrderActivityCollection(v []Execution)`

SetOrderActivityCollection sets OrderActivityCollection field to given value.

### HasOrderActivityCollection

`func (o *CashAccountOrderStrategies) HasOrderActivityCollection() bool`

HasOrderActivityCollection returns a boolean if a field has been set.

### GetOrderId

`func (o *CashAccountOrderStrategies) GetOrderId() float32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CashAccountOrderStrategies) GetOrderIdOk() (*float32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CashAccountOrderStrategies) SetOrderId(v float32)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *CashAccountOrderStrategies) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrderLegCollection

`func (o *CashAccountOrderStrategies) GetOrderLegCollection() []CashAccountOrderLegCollection`

GetOrderLegCollection returns the OrderLegCollection field if non-nil, zero value otherwise.

### GetOrderLegCollectionOk

`func (o *CashAccountOrderStrategies) GetOrderLegCollectionOk() (*[]CashAccountOrderLegCollection, bool)`

GetOrderLegCollectionOk returns a tuple with the OrderLegCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderLegCollection

`func (o *CashAccountOrderStrategies) SetOrderLegCollection(v []CashAccountOrderLegCollection)`

SetOrderLegCollection sets OrderLegCollection field to given value.

### HasOrderLegCollection

`func (o *CashAccountOrderStrategies) HasOrderLegCollection() bool`

HasOrderLegCollection returns a boolean if a field has been set.

### GetOrderStrategyType

`func (o *CashAccountOrderStrategies) GetOrderStrategyType() OrderStrategyType`

GetOrderStrategyType returns the OrderStrategyType field if non-nil, zero value otherwise.

### GetOrderStrategyTypeOk

`func (o *CashAccountOrderStrategies) GetOrderStrategyTypeOk() (*OrderStrategyType, bool)`

GetOrderStrategyTypeOk returns a tuple with the OrderStrategyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStrategyType

`func (o *CashAccountOrderStrategies) SetOrderStrategyType(v OrderStrategyType)`

SetOrderStrategyType sets OrderStrategyType field to given value.

### HasOrderStrategyType

`func (o *CashAccountOrderStrategies) HasOrderStrategyType() bool`

HasOrderStrategyType returns a boolean if a field has been set.

### GetOrderType

`func (o *CashAccountOrderStrategies) GetOrderType() OrderType`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *CashAccountOrderStrategies) GetOrderTypeOk() (*OrderType, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *CashAccountOrderStrategies) SetOrderType(v OrderType)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *CashAccountOrderStrategies) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetPrice

`func (o *CashAccountOrderStrategies) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CashAccountOrderStrategies) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CashAccountOrderStrategies) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CashAccountOrderStrategies) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceLinkBasis

`func (o *CashAccountOrderStrategies) GetPriceLinkBasis() PriceLinkBasis`

GetPriceLinkBasis returns the PriceLinkBasis field if non-nil, zero value otherwise.

### GetPriceLinkBasisOk

`func (o *CashAccountOrderStrategies) GetPriceLinkBasisOk() (*PriceLinkBasis, bool)`

GetPriceLinkBasisOk returns a tuple with the PriceLinkBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceLinkBasis

`func (o *CashAccountOrderStrategies) SetPriceLinkBasis(v PriceLinkBasis)`

SetPriceLinkBasis sets PriceLinkBasis field to given value.

### HasPriceLinkBasis

`func (o *CashAccountOrderStrategies) HasPriceLinkBasis() bool`

HasPriceLinkBasis returns a boolean if a field has been set.

### GetPriceLinkType

`func (o *CashAccountOrderStrategies) GetPriceLinkType() PriceLinkType`

GetPriceLinkType returns the PriceLinkType field if non-nil, zero value otherwise.

### GetPriceLinkTypeOk

`func (o *CashAccountOrderStrategies) GetPriceLinkTypeOk() (*PriceLinkType, bool)`

GetPriceLinkTypeOk returns a tuple with the PriceLinkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceLinkType

`func (o *CashAccountOrderStrategies) SetPriceLinkType(v PriceLinkType)`

SetPriceLinkType sets PriceLinkType field to given value.

### HasPriceLinkType

`func (o *CashAccountOrderStrategies) HasPriceLinkType() bool`

HasPriceLinkType returns a boolean if a field has been set.

### GetQuantity

`func (o *CashAccountOrderStrategies) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CashAccountOrderStrategies) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CashAccountOrderStrategies) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *CashAccountOrderStrategies) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetReleaseTime

`func (o *CashAccountOrderStrategies) GetReleaseTime() string`

GetReleaseTime returns the ReleaseTime field if non-nil, zero value otherwise.

### GetReleaseTimeOk

`func (o *CashAccountOrderStrategies) GetReleaseTimeOk() (*string, bool)`

GetReleaseTimeOk returns a tuple with the ReleaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseTime

`func (o *CashAccountOrderStrategies) SetReleaseTime(v string)`

SetReleaseTime sets ReleaseTime field to given value.

### HasReleaseTime

`func (o *CashAccountOrderStrategies) HasReleaseTime() bool`

HasReleaseTime returns a boolean if a field has been set.

### GetRemainingQuantity

`func (o *CashAccountOrderStrategies) GetRemainingQuantity() float32`

GetRemainingQuantity returns the RemainingQuantity field if non-nil, zero value otherwise.

### GetRemainingQuantityOk

`func (o *CashAccountOrderStrategies) GetRemainingQuantityOk() (*float32, bool)`

GetRemainingQuantityOk returns a tuple with the RemainingQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingQuantity

`func (o *CashAccountOrderStrategies) SetRemainingQuantity(v float32)`

SetRemainingQuantity sets RemainingQuantity field to given value.

### HasRemainingQuantity

`func (o *CashAccountOrderStrategies) HasRemainingQuantity() bool`

HasRemainingQuantity returns a boolean if a field has been set.

### GetReplacingOrderCollection

`func (o *CashAccountOrderStrategies) GetReplacingOrderCollection() []map[string]interface{}`

GetReplacingOrderCollection returns the ReplacingOrderCollection field if non-nil, zero value otherwise.

### GetReplacingOrderCollectionOk

`func (o *CashAccountOrderStrategies) GetReplacingOrderCollectionOk() (*[]map[string]interface{}, bool)`

GetReplacingOrderCollectionOk returns a tuple with the ReplacingOrderCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacingOrderCollection

`func (o *CashAccountOrderStrategies) SetReplacingOrderCollection(v []map[string]interface{})`

SetReplacingOrderCollection sets ReplacingOrderCollection field to given value.

### HasReplacingOrderCollection

`func (o *CashAccountOrderStrategies) HasReplacingOrderCollection() bool`

HasReplacingOrderCollection returns a boolean if a field has been set.

### GetRequestedDestination

`func (o *CashAccountOrderStrategies) GetRequestedDestination() DestinationExchange`

GetRequestedDestination returns the RequestedDestination field if non-nil, zero value otherwise.

### GetRequestedDestinationOk

`func (o *CashAccountOrderStrategies) GetRequestedDestinationOk() (*DestinationExchange, bool)`

GetRequestedDestinationOk returns a tuple with the RequestedDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedDestination

`func (o *CashAccountOrderStrategies) SetRequestedDestination(v DestinationExchange)`

SetRequestedDestination sets RequestedDestination field to given value.

### HasRequestedDestination

`func (o *CashAccountOrderStrategies) HasRequestedDestination() bool`

HasRequestedDestination returns a boolean if a field has been set.

### GetSession

`func (o *CashAccountOrderStrategies) GetSession() Session`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *CashAccountOrderStrategies) GetSessionOk() (*Session, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *CashAccountOrderStrategies) SetSession(v Session)`

SetSession sets Session field to given value.

### HasSession

`func (o *CashAccountOrderStrategies) HasSession() bool`

HasSession returns a boolean if a field has been set.

### GetSpecialInstruction

`func (o *CashAccountOrderStrategies) GetSpecialInstruction() SpecialInstruction`

GetSpecialInstruction returns the SpecialInstruction field if non-nil, zero value otherwise.

### GetSpecialInstructionOk

`func (o *CashAccountOrderStrategies) GetSpecialInstructionOk() (*SpecialInstruction, bool)`

GetSpecialInstructionOk returns a tuple with the SpecialInstruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialInstruction

`func (o *CashAccountOrderStrategies) SetSpecialInstruction(v SpecialInstruction)`

SetSpecialInstruction sets SpecialInstruction field to given value.

### HasSpecialInstruction

`func (o *CashAccountOrderStrategies) HasSpecialInstruction() bool`

HasSpecialInstruction returns a boolean if a field has been set.

### GetStatus

`func (o *CashAccountOrderStrategies) GetStatus() OrderStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CashAccountOrderStrategies) GetStatusOk() (*OrderStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CashAccountOrderStrategies) SetStatus(v OrderStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CashAccountOrderStrategies) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDescription

`func (o *CashAccountOrderStrategies) GetStatusDescription() string`

GetStatusDescription returns the StatusDescription field if non-nil, zero value otherwise.

### GetStatusDescriptionOk

`func (o *CashAccountOrderStrategies) GetStatusDescriptionOk() (*string, bool)`

GetStatusDescriptionOk returns a tuple with the StatusDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDescription

`func (o *CashAccountOrderStrategies) SetStatusDescription(v string)`

SetStatusDescription sets StatusDescription field to given value.

### HasStatusDescription

`func (o *CashAccountOrderStrategies) HasStatusDescription() bool`

HasStatusDescription returns a boolean if a field has been set.

### GetStopPrice

`func (o *CashAccountOrderStrategies) GetStopPrice() float32`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *CashAccountOrderStrategies) GetStopPriceOk() (*float32, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *CashAccountOrderStrategies) SetStopPrice(v float32)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *CashAccountOrderStrategies) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.

### GetStopPriceLinkBasis

`func (o *CashAccountOrderStrategies) GetStopPriceLinkBasis() PriceLinkBasis`

GetStopPriceLinkBasis returns the StopPriceLinkBasis field if non-nil, zero value otherwise.

### GetStopPriceLinkBasisOk

`func (o *CashAccountOrderStrategies) GetStopPriceLinkBasisOk() (*PriceLinkBasis, bool)`

GetStopPriceLinkBasisOk returns a tuple with the StopPriceLinkBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPriceLinkBasis

`func (o *CashAccountOrderStrategies) SetStopPriceLinkBasis(v PriceLinkBasis)`

SetStopPriceLinkBasis sets StopPriceLinkBasis field to given value.

### HasStopPriceLinkBasis

`func (o *CashAccountOrderStrategies) HasStopPriceLinkBasis() bool`

HasStopPriceLinkBasis returns a boolean if a field has been set.

### GetStopPriceLinkType

`func (o *CashAccountOrderStrategies) GetStopPriceLinkType() PriceLinkType`

GetStopPriceLinkType returns the StopPriceLinkType field if non-nil, zero value otherwise.

### GetStopPriceLinkTypeOk

`func (o *CashAccountOrderStrategies) GetStopPriceLinkTypeOk() (*PriceLinkType, bool)`

GetStopPriceLinkTypeOk returns a tuple with the StopPriceLinkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPriceLinkType

`func (o *CashAccountOrderStrategies) SetStopPriceLinkType(v PriceLinkType)`

SetStopPriceLinkType sets StopPriceLinkType field to given value.

### HasStopPriceLinkType

`func (o *CashAccountOrderStrategies) HasStopPriceLinkType() bool`

HasStopPriceLinkType returns a boolean if a field has been set.

### GetStopPriceOffset

`func (o *CashAccountOrderStrategies) GetStopPriceOffset() float32`

GetStopPriceOffset returns the StopPriceOffset field if non-nil, zero value otherwise.

### GetStopPriceOffsetOk

`func (o *CashAccountOrderStrategies) GetStopPriceOffsetOk() (*float32, bool)`

GetStopPriceOffsetOk returns a tuple with the StopPriceOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPriceOffset

`func (o *CashAccountOrderStrategies) SetStopPriceOffset(v float32)`

SetStopPriceOffset sets StopPriceOffset field to given value.

### HasStopPriceOffset

`func (o *CashAccountOrderStrategies) HasStopPriceOffset() bool`

HasStopPriceOffset returns a boolean if a field has been set.

### GetStopType

`func (o *CashAccountOrderStrategies) GetStopType() StopType`

GetStopType returns the StopType field if non-nil, zero value otherwise.

### GetStopTypeOk

`func (o *CashAccountOrderStrategies) GetStopTypeOk() (*StopType, bool)`

GetStopTypeOk returns a tuple with the StopType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopType

`func (o *CashAccountOrderStrategies) SetStopType(v StopType)`

SetStopType sets StopType field to given value.

### HasStopType

`func (o *CashAccountOrderStrategies) HasStopType() bool`

HasStopType returns a boolean if a field has been set.

### GetTag

`func (o *CashAccountOrderStrategies) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CashAccountOrderStrategies) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CashAccountOrderStrategies) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *CashAccountOrderStrategies) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTaxLotMethod

`func (o *CashAccountOrderStrategies) GetTaxLotMethod() TaxLotMethod`

GetTaxLotMethod returns the TaxLotMethod field if non-nil, zero value otherwise.

### GetTaxLotMethodOk

`func (o *CashAccountOrderStrategies) GetTaxLotMethodOk() (*TaxLotMethod, bool)`

GetTaxLotMethodOk returns a tuple with the TaxLotMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxLotMethod

`func (o *CashAccountOrderStrategies) SetTaxLotMethod(v TaxLotMethod)`

SetTaxLotMethod sets TaxLotMethod field to given value.

### HasTaxLotMethod

`func (o *CashAccountOrderStrategies) HasTaxLotMethod() bool`

HasTaxLotMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


