# SavedOrderChildOrderStrategies

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **float32** |  | [optional] 
**ActivationPrice** | Pointer to **float32** |  | [optional] 
**CancelTime** | Pointer to [**CashAccountCancelTime**](CashAccount_cancelTime.md) |  | [optional] 
**Cancelable** | Pointer to **bool** |  | [optional] 
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
**StopPrice** | Pointer to **float32** |  | [optional] 
**StopPriceLinkBasis** | Pointer to [**PriceLinkBasis**](PriceLinkBasis.md) |  | [optional] 
**StopPriceLinkType** | Pointer to [**PriceLinkType**](PriceLinkType.md) |  | [optional] 
**StopPriceOffset** | Pointer to **float32** |  | [optional] 
**StopType** | Pointer to [**StopType**](StopType.md) |  | [optional] 
**Tag** | Pointer to **string** |  | [optional] 
**TaxLotMethod** | Pointer to [**TaxLotMethod**](TaxLotMethod.md) |  | [optional] 

## Methods

### NewSavedOrderChildOrderStrategies

`func NewSavedOrderChildOrderStrategies() *SavedOrderChildOrderStrategies`

NewSavedOrderChildOrderStrategies instantiates a new SavedOrderChildOrderStrategies object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSavedOrderChildOrderStrategiesWithDefaults

`func NewSavedOrderChildOrderStrategiesWithDefaults() *SavedOrderChildOrderStrategies`

NewSavedOrderChildOrderStrategiesWithDefaults instantiates a new SavedOrderChildOrderStrategies object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *SavedOrderChildOrderStrategies) GetAccountId() float32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SavedOrderChildOrderStrategies) GetAccountIdOk() (*float32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SavedOrderChildOrderStrategies) SetAccountId(v float32)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *SavedOrderChildOrderStrategies) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetActivationPrice

`func (o *SavedOrderChildOrderStrategies) GetActivationPrice() float32`

GetActivationPrice returns the ActivationPrice field if non-nil, zero value otherwise.

### GetActivationPriceOk

`func (o *SavedOrderChildOrderStrategies) GetActivationPriceOk() (*float32, bool)`

GetActivationPriceOk returns a tuple with the ActivationPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationPrice

`func (o *SavedOrderChildOrderStrategies) SetActivationPrice(v float32)`

SetActivationPrice sets ActivationPrice field to given value.

### HasActivationPrice

`func (o *SavedOrderChildOrderStrategies) HasActivationPrice() bool`

HasActivationPrice returns a boolean if a field has been set.

### GetCancelTime

`func (o *SavedOrderChildOrderStrategies) GetCancelTime() CashAccountCancelTime`

GetCancelTime returns the CancelTime field if non-nil, zero value otherwise.

### GetCancelTimeOk

`func (o *SavedOrderChildOrderStrategies) GetCancelTimeOk() (*CashAccountCancelTime, bool)`

GetCancelTimeOk returns a tuple with the CancelTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelTime

`func (o *SavedOrderChildOrderStrategies) SetCancelTime(v CashAccountCancelTime)`

SetCancelTime sets CancelTime field to given value.

### HasCancelTime

`func (o *SavedOrderChildOrderStrategies) HasCancelTime() bool`

HasCancelTime returns a boolean if a field has been set.

### GetCancelable

`func (o *SavedOrderChildOrderStrategies) GetCancelable() bool`

GetCancelable returns the Cancelable field if non-nil, zero value otherwise.

### GetCancelableOk

`func (o *SavedOrderChildOrderStrategies) GetCancelableOk() (*bool, bool)`

GetCancelableOk returns a tuple with the Cancelable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelable

`func (o *SavedOrderChildOrderStrategies) SetCancelable(v bool)`

SetCancelable sets Cancelable field to given value.

### HasCancelable

`func (o *SavedOrderChildOrderStrategies) HasCancelable() bool`

HasCancelable returns a boolean if a field has been set.

### GetCloseTime

`func (o *SavedOrderChildOrderStrategies) GetCloseTime() string`

GetCloseTime returns the CloseTime field if non-nil, zero value otherwise.

### GetCloseTimeOk

`func (o *SavedOrderChildOrderStrategies) GetCloseTimeOk() (*string, bool)`

GetCloseTimeOk returns a tuple with the CloseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseTime

`func (o *SavedOrderChildOrderStrategies) SetCloseTime(v string)`

SetCloseTime sets CloseTime field to given value.

### HasCloseTime

`func (o *SavedOrderChildOrderStrategies) HasCloseTime() bool`

HasCloseTime returns a boolean if a field has been set.

### GetComplexOrderStrategyType

`func (o *SavedOrderChildOrderStrategies) GetComplexOrderStrategyType() ComplexOrderStrategyType`

GetComplexOrderStrategyType returns the ComplexOrderStrategyType field if non-nil, zero value otherwise.

### GetComplexOrderStrategyTypeOk

`func (o *SavedOrderChildOrderStrategies) GetComplexOrderStrategyTypeOk() (*ComplexOrderStrategyType, bool)`

GetComplexOrderStrategyTypeOk returns a tuple with the ComplexOrderStrategyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplexOrderStrategyType

`func (o *SavedOrderChildOrderStrategies) SetComplexOrderStrategyType(v ComplexOrderStrategyType)`

SetComplexOrderStrategyType sets ComplexOrderStrategyType field to given value.

### HasComplexOrderStrategyType

`func (o *SavedOrderChildOrderStrategies) HasComplexOrderStrategyType() bool`

HasComplexOrderStrategyType returns a boolean if a field has been set.

### GetDestinationLinkName

`func (o *SavedOrderChildOrderStrategies) GetDestinationLinkName() string`

GetDestinationLinkName returns the DestinationLinkName field if non-nil, zero value otherwise.

### GetDestinationLinkNameOk

`func (o *SavedOrderChildOrderStrategies) GetDestinationLinkNameOk() (*string, bool)`

GetDestinationLinkNameOk returns a tuple with the DestinationLinkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationLinkName

`func (o *SavedOrderChildOrderStrategies) SetDestinationLinkName(v string)`

SetDestinationLinkName sets DestinationLinkName field to given value.

### HasDestinationLinkName

`func (o *SavedOrderChildOrderStrategies) HasDestinationLinkName() bool`

HasDestinationLinkName returns a boolean if a field has been set.

### GetDuration

`func (o *SavedOrderChildOrderStrategies) GetDuration() Duration`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *SavedOrderChildOrderStrategies) GetDurationOk() (*Duration, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *SavedOrderChildOrderStrategies) SetDuration(v Duration)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *SavedOrderChildOrderStrategies) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEditable

`func (o *SavedOrderChildOrderStrategies) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *SavedOrderChildOrderStrategies) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *SavedOrderChildOrderStrategies) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *SavedOrderChildOrderStrategies) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetEnteredTime

`func (o *SavedOrderChildOrderStrategies) GetEnteredTime() string`

GetEnteredTime returns the EnteredTime field if non-nil, zero value otherwise.

### GetEnteredTimeOk

`func (o *SavedOrderChildOrderStrategies) GetEnteredTimeOk() (*string, bool)`

GetEnteredTimeOk returns a tuple with the EnteredTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnteredTime

`func (o *SavedOrderChildOrderStrategies) SetEnteredTime(v string)`

SetEnteredTime sets EnteredTime field to given value.

### HasEnteredTime

`func (o *SavedOrderChildOrderStrategies) HasEnteredTime() bool`

HasEnteredTime returns a boolean if a field has been set.

### GetFilledQuantity

`func (o *SavedOrderChildOrderStrategies) GetFilledQuantity() float32`

GetFilledQuantity returns the FilledQuantity field if non-nil, zero value otherwise.

### GetFilledQuantityOk

`func (o *SavedOrderChildOrderStrategies) GetFilledQuantityOk() (*float32, bool)`

GetFilledQuantityOk returns a tuple with the FilledQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilledQuantity

`func (o *SavedOrderChildOrderStrategies) SetFilledQuantity(v float32)`

SetFilledQuantity sets FilledQuantity field to given value.

### HasFilledQuantity

`func (o *SavedOrderChildOrderStrategies) HasFilledQuantity() bool`

HasFilledQuantity returns a boolean if a field has been set.

### GetOrderActivityCollection

`func (o *SavedOrderChildOrderStrategies) GetOrderActivityCollection() []Execution`

GetOrderActivityCollection returns the OrderActivityCollection field if non-nil, zero value otherwise.

### GetOrderActivityCollectionOk

`func (o *SavedOrderChildOrderStrategies) GetOrderActivityCollectionOk() (*[]Execution, bool)`

GetOrderActivityCollectionOk returns a tuple with the OrderActivityCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderActivityCollection

`func (o *SavedOrderChildOrderStrategies) SetOrderActivityCollection(v []Execution)`

SetOrderActivityCollection sets OrderActivityCollection field to given value.

### HasOrderActivityCollection

`func (o *SavedOrderChildOrderStrategies) HasOrderActivityCollection() bool`

HasOrderActivityCollection returns a boolean if a field has been set.

### GetOrderId

`func (o *SavedOrderChildOrderStrategies) GetOrderId() float32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *SavedOrderChildOrderStrategies) GetOrderIdOk() (*float32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *SavedOrderChildOrderStrategies) SetOrderId(v float32)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *SavedOrderChildOrderStrategies) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrderLegCollection

`func (o *SavedOrderChildOrderStrategies) GetOrderLegCollection() []CashAccountOrderLegCollection`

GetOrderLegCollection returns the OrderLegCollection field if non-nil, zero value otherwise.

### GetOrderLegCollectionOk

`func (o *SavedOrderChildOrderStrategies) GetOrderLegCollectionOk() (*[]CashAccountOrderLegCollection, bool)`

GetOrderLegCollectionOk returns a tuple with the OrderLegCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderLegCollection

`func (o *SavedOrderChildOrderStrategies) SetOrderLegCollection(v []CashAccountOrderLegCollection)`

SetOrderLegCollection sets OrderLegCollection field to given value.

### HasOrderLegCollection

`func (o *SavedOrderChildOrderStrategies) HasOrderLegCollection() bool`

HasOrderLegCollection returns a boolean if a field has been set.

### GetOrderStrategyType

`func (o *SavedOrderChildOrderStrategies) GetOrderStrategyType() OrderStrategyType`

GetOrderStrategyType returns the OrderStrategyType field if non-nil, zero value otherwise.

### GetOrderStrategyTypeOk

`func (o *SavedOrderChildOrderStrategies) GetOrderStrategyTypeOk() (*OrderStrategyType, bool)`

GetOrderStrategyTypeOk returns a tuple with the OrderStrategyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStrategyType

`func (o *SavedOrderChildOrderStrategies) SetOrderStrategyType(v OrderStrategyType)`

SetOrderStrategyType sets OrderStrategyType field to given value.

### HasOrderStrategyType

`func (o *SavedOrderChildOrderStrategies) HasOrderStrategyType() bool`

HasOrderStrategyType returns a boolean if a field has been set.

### GetOrderType

`func (o *SavedOrderChildOrderStrategies) GetOrderType() OrderType`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *SavedOrderChildOrderStrategies) GetOrderTypeOk() (*OrderType, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *SavedOrderChildOrderStrategies) SetOrderType(v OrderType)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *SavedOrderChildOrderStrategies) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetPrice

`func (o *SavedOrderChildOrderStrategies) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *SavedOrderChildOrderStrategies) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *SavedOrderChildOrderStrategies) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *SavedOrderChildOrderStrategies) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceLinkBasis

`func (o *SavedOrderChildOrderStrategies) GetPriceLinkBasis() PriceLinkBasis`

GetPriceLinkBasis returns the PriceLinkBasis field if non-nil, zero value otherwise.

### GetPriceLinkBasisOk

`func (o *SavedOrderChildOrderStrategies) GetPriceLinkBasisOk() (*PriceLinkBasis, bool)`

GetPriceLinkBasisOk returns a tuple with the PriceLinkBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceLinkBasis

`func (o *SavedOrderChildOrderStrategies) SetPriceLinkBasis(v PriceLinkBasis)`

SetPriceLinkBasis sets PriceLinkBasis field to given value.

### HasPriceLinkBasis

`func (o *SavedOrderChildOrderStrategies) HasPriceLinkBasis() bool`

HasPriceLinkBasis returns a boolean if a field has been set.

### GetPriceLinkType

`func (o *SavedOrderChildOrderStrategies) GetPriceLinkType() PriceLinkType`

GetPriceLinkType returns the PriceLinkType field if non-nil, zero value otherwise.

### GetPriceLinkTypeOk

`func (o *SavedOrderChildOrderStrategies) GetPriceLinkTypeOk() (*PriceLinkType, bool)`

GetPriceLinkTypeOk returns a tuple with the PriceLinkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceLinkType

`func (o *SavedOrderChildOrderStrategies) SetPriceLinkType(v PriceLinkType)`

SetPriceLinkType sets PriceLinkType field to given value.

### HasPriceLinkType

`func (o *SavedOrderChildOrderStrategies) HasPriceLinkType() bool`

HasPriceLinkType returns a boolean if a field has been set.

### GetQuantity

`func (o *SavedOrderChildOrderStrategies) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *SavedOrderChildOrderStrategies) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *SavedOrderChildOrderStrategies) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *SavedOrderChildOrderStrategies) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetReleaseTime

`func (o *SavedOrderChildOrderStrategies) GetReleaseTime() string`

GetReleaseTime returns the ReleaseTime field if non-nil, zero value otherwise.

### GetReleaseTimeOk

`func (o *SavedOrderChildOrderStrategies) GetReleaseTimeOk() (*string, bool)`

GetReleaseTimeOk returns a tuple with the ReleaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseTime

`func (o *SavedOrderChildOrderStrategies) SetReleaseTime(v string)`

SetReleaseTime sets ReleaseTime field to given value.

### HasReleaseTime

`func (o *SavedOrderChildOrderStrategies) HasReleaseTime() bool`

HasReleaseTime returns a boolean if a field has been set.

### GetRemainingQuantity

`func (o *SavedOrderChildOrderStrategies) GetRemainingQuantity() float32`

GetRemainingQuantity returns the RemainingQuantity field if non-nil, zero value otherwise.

### GetRemainingQuantityOk

`func (o *SavedOrderChildOrderStrategies) GetRemainingQuantityOk() (*float32, bool)`

GetRemainingQuantityOk returns a tuple with the RemainingQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingQuantity

`func (o *SavedOrderChildOrderStrategies) SetRemainingQuantity(v float32)`

SetRemainingQuantity sets RemainingQuantity field to given value.

### HasRemainingQuantity

`func (o *SavedOrderChildOrderStrategies) HasRemainingQuantity() bool`

HasRemainingQuantity returns a boolean if a field has been set.

### GetReplacingOrderCollection

`func (o *SavedOrderChildOrderStrategies) GetReplacingOrderCollection() []map[string]interface{}`

GetReplacingOrderCollection returns the ReplacingOrderCollection field if non-nil, zero value otherwise.

### GetReplacingOrderCollectionOk

`func (o *SavedOrderChildOrderStrategies) GetReplacingOrderCollectionOk() (*[]map[string]interface{}, bool)`

GetReplacingOrderCollectionOk returns a tuple with the ReplacingOrderCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacingOrderCollection

`func (o *SavedOrderChildOrderStrategies) SetReplacingOrderCollection(v []map[string]interface{})`

SetReplacingOrderCollection sets ReplacingOrderCollection field to given value.

### HasReplacingOrderCollection

`func (o *SavedOrderChildOrderStrategies) HasReplacingOrderCollection() bool`

HasReplacingOrderCollection returns a boolean if a field has been set.

### GetRequestedDestination

`func (o *SavedOrderChildOrderStrategies) GetRequestedDestination() DestinationExchange`

GetRequestedDestination returns the RequestedDestination field if non-nil, zero value otherwise.

### GetRequestedDestinationOk

`func (o *SavedOrderChildOrderStrategies) GetRequestedDestinationOk() (*DestinationExchange, bool)`

GetRequestedDestinationOk returns a tuple with the RequestedDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedDestination

`func (o *SavedOrderChildOrderStrategies) SetRequestedDestination(v DestinationExchange)`

SetRequestedDestination sets RequestedDestination field to given value.

### HasRequestedDestination

`func (o *SavedOrderChildOrderStrategies) HasRequestedDestination() bool`

HasRequestedDestination returns a boolean if a field has been set.

### GetSession

`func (o *SavedOrderChildOrderStrategies) GetSession() Session`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *SavedOrderChildOrderStrategies) GetSessionOk() (*Session, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *SavedOrderChildOrderStrategies) SetSession(v Session)`

SetSession sets Session field to given value.

### HasSession

`func (o *SavedOrderChildOrderStrategies) HasSession() bool`

HasSession returns a boolean if a field has been set.

### GetSpecialInstruction

`func (o *SavedOrderChildOrderStrategies) GetSpecialInstruction() SpecialInstruction`

GetSpecialInstruction returns the SpecialInstruction field if non-nil, zero value otherwise.

### GetSpecialInstructionOk

`func (o *SavedOrderChildOrderStrategies) GetSpecialInstructionOk() (*SpecialInstruction, bool)`

GetSpecialInstructionOk returns a tuple with the SpecialInstruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialInstruction

`func (o *SavedOrderChildOrderStrategies) SetSpecialInstruction(v SpecialInstruction)`

SetSpecialInstruction sets SpecialInstruction field to given value.

### HasSpecialInstruction

`func (o *SavedOrderChildOrderStrategies) HasSpecialInstruction() bool`

HasSpecialInstruction returns a boolean if a field has been set.

### GetStatus

`func (o *SavedOrderChildOrderStrategies) GetStatus() OrderStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SavedOrderChildOrderStrategies) GetStatusOk() (*OrderStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SavedOrderChildOrderStrategies) SetStatus(v OrderStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SavedOrderChildOrderStrategies) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStopPrice

`func (o *SavedOrderChildOrderStrategies) GetStopPrice() float32`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *SavedOrderChildOrderStrategies) GetStopPriceOk() (*float32, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *SavedOrderChildOrderStrategies) SetStopPrice(v float32)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *SavedOrderChildOrderStrategies) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.

### GetStopPriceLinkBasis

`func (o *SavedOrderChildOrderStrategies) GetStopPriceLinkBasis() PriceLinkBasis`

GetStopPriceLinkBasis returns the StopPriceLinkBasis field if non-nil, zero value otherwise.

### GetStopPriceLinkBasisOk

`func (o *SavedOrderChildOrderStrategies) GetStopPriceLinkBasisOk() (*PriceLinkBasis, bool)`

GetStopPriceLinkBasisOk returns a tuple with the StopPriceLinkBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPriceLinkBasis

`func (o *SavedOrderChildOrderStrategies) SetStopPriceLinkBasis(v PriceLinkBasis)`

SetStopPriceLinkBasis sets StopPriceLinkBasis field to given value.

### HasStopPriceLinkBasis

`func (o *SavedOrderChildOrderStrategies) HasStopPriceLinkBasis() bool`

HasStopPriceLinkBasis returns a boolean if a field has been set.

### GetStopPriceLinkType

`func (o *SavedOrderChildOrderStrategies) GetStopPriceLinkType() PriceLinkType`

GetStopPriceLinkType returns the StopPriceLinkType field if non-nil, zero value otherwise.

### GetStopPriceLinkTypeOk

`func (o *SavedOrderChildOrderStrategies) GetStopPriceLinkTypeOk() (*PriceLinkType, bool)`

GetStopPriceLinkTypeOk returns a tuple with the StopPriceLinkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPriceLinkType

`func (o *SavedOrderChildOrderStrategies) SetStopPriceLinkType(v PriceLinkType)`

SetStopPriceLinkType sets StopPriceLinkType field to given value.

### HasStopPriceLinkType

`func (o *SavedOrderChildOrderStrategies) HasStopPriceLinkType() bool`

HasStopPriceLinkType returns a boolean if a field has been set.

### GetStopPriceOffset

`func (o *SavedOrderChildOrderStrategies) GetStopPriceOffset() float32`

GetStopPriceOffset returns the StopPriceOffset field if non-nil, zero value otherwise.

### GetStopPriceOffsetOk

`func (o *SavedOrderChildOrderStrategies) GetStopPriceOffsetOk() (*float32, bool)`

GetStopPriceOffsetOk returns a tuple with the StopPriceOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPriceOffset

`func (o *SavedOrderChildOrderStrategies) SetStopPriceOffset(v float32)`

SetStopPriceOffset sets StopPriceOffset field to given value.

### HasStopPriceOffset

`func (o *SavedOrderChildOrderStrategies) HasStopPriceOffset() bool`

HasStopPriceOffset returns a boolean if a field has been set.

### GetStopType

`func (o *SavedOrderChildOrderStrategies) GetStopType() StopType`

GetStopType returns the StopType field if non-nil, zero value otherwise.

### GetStopTypeOk

`func (o *SavedOrderChildOrderStrategies) GetStopTypeOk() (*StopType, bool)`

GetStopTypeOk returns a tuple with the StopType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopType

`func (o *SavedOrderChildOrderStrategies) SetStopType(v StopType)`

SetStopType sets StopType field to given value.

### HasStopType

`func (o *SavedOrderChildOrderStrategies) HasStopType() bool`

HasStopType returns a boolean if a field has been set.

### GetTag

`func (o *SavedOrderChildOrderStrategies) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *SavedOrderChildOrderStrategies) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *SavedOrderChildOrderStrategies) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *SavedOrderChildOrderStrategies) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTaxLotMethod

`func (o *SavedOrderChildOrderStrategies) GetTaxLotMethod() TaxLotMethod`

GetTaxLotMethod returns the TaxLotMethod field if non-nil, zero value otherwise.

### GetTaxLotMethodOk

`func (o *SavedOrderChildOrderStrategies) GetTaxLotMethodOk() (*TaxLotMethod, bool)`

GetTaxLotMethodOk returns a tuple with the TaxLotMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxLotMethod

`func (o *SavedOrderChildOrderStrategies) SetTaxLotMethod(v TaxLotMethod)`

SetTaxLotMethod sets TaxLotMethod field to given value.

### HasTaxLotMethod

`func (o *SavedOrderChildOrderStrategies) HasTaxLotMethod() bool`

HasTaxLotMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


