# SavedOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **float32** |  | [optional] 
**ActivationPrice** | Pointer to **float32** |  | [optional] 
**CancelTime** | Pointer to [**CashAccountCancelTime**](CashAccount_cancelTime.md) |  | [optional] 
**Cancelable** | Pointer to **bool** |  | [optional] 
**ChildOrderStrategies** | Pointer to [**[]SavedOrderChildOrderStrategies**](SavedOrderChildOrderStrategies.md) |  | [optional] 
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
**RequestedDestination** | Pointer to [**DestinationExchange**](DestinationExchange.md) |  | [optional] 
**SavedOrderId** | Pointer to **float32** |  | [optional] 
**SavedTime** | Pointer to **string** |  | [optional] 
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

### NewSavedOrder

`func NewSavedOrder() *SavedOrder`

NewSavedOrder instantiates a new SavedOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSavedOrderWithDefaults

`func NewSavedOrderWithDefaults() *SavedOrder`

NewSavedOrderWithDefaults instantiates a new SavedOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *SavedOrder) GetAccountId() float32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SavedOrder) GetAccountIdOk() (*float32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SavedOrder) SetAccountId(v float32)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *SavedOrder) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetActivationPrice

`func (o *SavedOrder) GetActivationPrice() float32`

GetActivationPrice returns the ActivationPrice field if non-nil, zero value otherwise.

### GetActivationPriceOk

`func (o *SavedOrder) GetActivationPriceOk() (*float32, bool)`

GetActivationPriceOk returns a tuple with the ActivationPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationPrice

`func (o *SavedOrder) SetActivationPrice(v float32)`

SetActivationPrice sets ActivationPrice field to given value.

### HasActivationPrice

`func (o *SavedOrder) HasActivationPrice() bool`

HasActivationPrice returns a boolean if a field has been set.

### GetCancelTime

`func (o *SavedOrder) GetCancelTime() CashAccountCancelTime`

GetCancelTime returns the CancelTime field if non-nil, zero value otherwise.

### GetCancelTimeOk

`func (o *SavedOrder) GetCancelTimeOk() (*CashAccountCancelTime, bool)`

GetCancelTimeOk returns a tuple with the CancelTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelTime

`func (o *SavedOrder) SetCancelTime(v CashAccountCancelTime)`

SetCancelTime sets CancelTime field to given value.

### HasCancelTime

`func (o *SavedOrder) HasCancelTime() bool`

HasCancelTime returns a boolean if a field has been set.

### GetCancelable

`func (o *SavedOrder) GetCancelable() bool`

GetCancelable returns the Cancelable field if non-nil, zero value otherwise.

### GetCancelableOk

`func (o *SavedOrder) GetCancelableOk() (*bool, bool)`

GetCancelableOk returns a tuple with the Cancelable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelable

`func (o *SavedOrder) SetCancelable(v bool)`

SetCancelable sets Cancelable field to given value.

### HasCancelable

`func (o *SavedOrder) HasCancelable() bool`

HasCancelable returns a boolean if a field has been set.

### GetChildOrderStrategies

`func (o *SavedOrder) GetChildOrderStrategies() []SavedOrderChildOrderStrategies`

GetChildOrderStrategies returns the ChildOrderStrategies field if non-nil, zero value otherwise.

### GetChildOrderStrategiesOk

`func (o *SavedOrder) GetChildOrderStrategiesOk() (*[]SavedOrderChildOrderStrategies, bool)`

GetChildOrderStrategiesOk returns a tuple with the ChildOrderStrategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildOrderStrategies

`func (o *SavedOrder) SetChildOrderStrategies(v []SavedOrderChildOrderStrategies)`

SetChildOrderStrategies sets ChildOrderStrategies field to given value.

### HasChildOrderStrategies

`func (o *SavedOrder) HasChildOrderStrategies() bool`

HasChildOrderStrategies returns a boolean if a field has been set.

### GetCloseTime

`func (o *SavedOrder) GetCloseTime() string`

GetCloseTime returns the CloseTime field if non-nil, zero value otherwise.

### GetCloseTimeOk

`func (o *SavedOrder) GetCloseTimeOk() (*string, bool)`

GetCloseTimeOk returns a tuple with the CloseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseTime

`func (o *SavedOrder) SetCloseTime(v string)`

SetCloseTime sets CloseTime field to given value.

### HasCloseTime

`func (o *SavedOrder) HasCloseTime() bool`

HasCloseTime returns a boolean if a field has been set.

### GetComplexOrderStrategyType

`func (o *SavedOrder) GetComplexOrderStrategyType() ComplexOrderStrategyType`

GetComplexOrderStrategyType returns the ComplexOrderStrategyType field if non-nil, zero value otherwise.

### GetComplexOrderStrategyTypeOk

`func (o *SavedOrder) GetComplexOrderStrategyTypeOk() (*ComplexOrderStrategyType, bool)`

GetComplexOrderStrategyTypeOk returns a tuple with the ComplexOrderStrategyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplexOrderStrategyType

`func (o *SavedOrder) SetComplexOrderStrategyType(v ComplexOrderStrategyType)`

SetComplexOrderStrategyType sets ComplexOrderStrategyType field to given value.

### HasComplexOrderStrategyType

`func (o *SavedOrder) HasComplexOrderStrategyType() bool`

HasComplexOrderStrategyType returns a boolean if a field has been set.

### GetDestinationLinkName

`func (o *SavedOrder) GetDestinationLinkName() string`

GetDestinationLinkName returns the DestinationLinkName field if non-nil, zero value otherwise.

### GetDestinationLinkNameOk

`func (o *SavedOrder) GetDestinationLinkNameOk() (*string, bool)`

GetDestinationLinkNameOk returns a tuple with the DestinationLinkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationLinkName

`func (o *SavedOrder) SetDestinationLinkName(v string)`

SetDestinationLinkName sets DestinationLinkName field to given value.

### HasDestinationLinkName

`func (o *SavedOrder) HasDestinationLinkName() bool`

HasDestinationLinkName returns a boolean if a field has been set.

### GetDuration

`func (o *SavedOrder) GetDuration() Duration`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *SavedOrder) GetDurationOk() (*Duration, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *SavedOrder) SetDuration(v Duration)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *SavedOrder) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEditable

`func (o *SavedOrder) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *SavedOrder) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *SavedOrder) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *SavedOrder) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetEnteredTime

`func (o *SavedOrder) GetEnteredTime() string`

GetEnteredTime returns the EnteredTime field if non-nil, zero value otherwise.

### GetEnteredTimeOk

`func (o *SavedOrder) GetEnteredTimeOk() (*string, bool)`

GetEnteredTimeOk returns a tuple with the EnteredTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnteredTime

`func (o *SavedOrder) SetEnteredTime(v string)`

SetEnteredTime sets EnteredTime field to given value.

### HasEnteredTime

`func (o *SavedOrder) HasEnteredTime() bool`

HasEnteredTime returns a boolean if a field has been set.

### GetFilledQuantity

`func (o *SavedOrder) GetFilledQuantity() float32`

GetFilledQuantity returns the FilledQuantity field if non-nil, zero value otherwise.

### GetFilledQuantityOk

`func (o *SavedOrder) GetFilledQuantityOk() (*float32, bool)`

GetFilledQuantityOk returns a tuple with the FilledQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilledQuantity

`func (o *SavedOrder) SetFilledQuantity(v float32)`

SetFilledQuantity sets FilledQuantity field to given value.

### HasFilledQuantity

`func (o *SavedOrder) HasFilledQuantity() bool`

HasFilledQuantity returns a boolean if a field has been set.

### GetOrderActivityCollection

`func (o *SavedOrder) GetOrderActivityCollection() []Execution`

GetOrderActivityCollection returns the OrderActivityCollection field if non-nil, zero value otherwise.

### GetOrderActivityCollectionOk

`func (o *SavedOrder) GetOrderActivityCollectionOk() (*[]Execution, bool)`

GetOrderActivityCollectionOk returns a tuple with the OrderActivityCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderActivityCollection

`func (o *SavedOrder) SetOrderActivityCollection(v []Execution)`

SetOrderActivityCollection sets OrderActivityCollection field to given value.

### HasOrderActivityCollection

`func (o *SavedOrder) HasOrderActivityCollection() bool`

HasOrderActivityCollection returns a boolean if a field has been set.

### GetOrderId

`func (o *SavedOrder) GetOrderId() float32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *SavedOrder) GetOrderIdOk() (*float32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *SavedOrder) SetOrderId(v float32)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *SavedOrder) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrderLegCollection

`func (o *SavedOrder) GetOrderLegCollection() []CashAccountOrderLegCollection`

GetOrderLegCollection returns the OrderLegCollection field if non-nil, zero value otherwise.

### GetOrderLegCollectionOk

`func (o *SavedOrder) GetOrderLegCollectionOk() (*[]CashAccountOrderLegCollection, bool)`

GetOrderLegCollectionOk returns a tuple with the OrderLegCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderLegCollection

`func (o *SavedOrder) SetOrderLegCollection(v []CashAccountOrderLegCollection)`

SetOrderLegCollection sets OrderLegCollection field to given value.

### HasOrderLegCollection

`func (o *SavedOrder) HasOrderLegCollection() bool`

HasOrderLegCollection returns a boolean if a field has been set.

### GetOrderStrategyType

`func (o *SavedOrder) GetOrderStrategyType() OrderStrategyType`

GetOrderStrategyType returns the OrderStrategyType field if non-nil, zero value otherwise.

### GetOrderStrategyTypeOk

`func (o *SavedOrder) GetOrderStrategyTypeOk() (*OrderStrategyType, bool)`

GetOrderStrategyTypeOk returns a tuple with the OrderStrategyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStrategyType

`func (o *SavedOrder) SetOrderStrategyType(v OrderStrategyType)`

SetOrderStrategyType sets OrderStrategyType field to given value.

### HasOrderStrategyType

`func (o *SavedOrder) HasOrderStrategyType() bool`

HasOrderStrategyType returns a boolean if a field has been set.

### GetOrderType

`func (o *SavedOrder) GetOrderType() OrderType`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *SavedOrder) GetOrderTypeOk() (*OrderType, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *SavedOrder) SetOrderType(v OrderType)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *SavedOrder) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetPrice

`func (o *SavedOrder) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *SavedOrder) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *SavedOrder) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *SavedOrder) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceLinkBasis

`func (o *SavedOrder) GetPriceLinkBasis() PriceLinkBasis`

GetPriceLinkBasis returns the PriceLinkBasis field if non-nil, zero value otherwise.

### GetPriceLinkBasisOk

`func (o *SavedOrder) GetPriceLinkBasisOk() (*PriceLinkBasis, bool)`

GetPriceLinkBasisOk returns a tuple with the PriceLinkBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceLinkBasis

`func (o *SavedOrder) SetPriceLinkBasis(v PriceLinkBasis)`

SetPriceLinkBasis sets PriceLinkBasis field to given value.

### HasPriceLinkBasis

`func (o *SavedOrder) HasPriceLinkBasis() bool`

HasPriceLinkBasis returns a boolean if a field has been set.

### GetPriceLinkType

`func (o *SavedOrder) GetPriceLinkType() PriceLinkType`

GetPriceLinkType returns the PriceLinkType field if non-nil, zero value otherwise.

### GetPriceLinkTypeOk

`func (o *SavedOrder) GetPriceLinkTypeOk() (*PriceLinkType, bool)`

GetPriceLinkTypeOk returns a tuple with the PriceLinkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceLinkType

`func (o *SavedOrder) SetPriceLinkType(v PriceLinkType)`

SetPriceLinkType sets PriceLinkType field to given value.

### HasPriceLinkType

`func (o *SavedOrder) HasPriceLinkType() bool`

HasPriceLinkType returns a boolean if a field has been set.

### GetQuantity

`func (o *SavedOrder) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *SavedOrder) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *SavedOrder) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *SavedOrder) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetReleaseTime

`func (o *SavedOrder) GetReleaseTime() string`

GetReleaseTime returns the ReleaseTime field if non-nil, zero value otherwise.

### GetReleaseTimeOk

`func (o *SavedOrder) GetReleaseTimeOk() (*string, bool)`

GetReleaseTimeOk returns a tuple with the ReleaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseTime

`func (o *SavedOrder) SetReleaseTime(v string)`

SetReleaseTime sets ReleaseTime field to given value.

### HasReleaseTime

`func (o *SavedOrder) HasReleaseTime() bool`

HasReleaseTime returns a boolean if a field has been set.

### GetRemainingQuantity

`func (o *SavedOrder) GetRemainingQuantity() float32`

GetRemainingQuantity returns the RemainingQuantity field if non-nil, zero value otherwise.

### GetRemainingQuantityOk

`func (o *SavedOrder) GetRemainingQuantityOk() (*float32, bool)`

GetRemainingQuantityOk returns a tuple with the RemainingQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingQuantity

`func (o *SavedOrder) SetRemainingQuantity(v float32)`

SetRemainingQuantity sets RemainingQuantity field to given value.

### HasRemainingQuantity

`func (o *SavedOrder) HasRemainingQuantity() bool`

HasRemainingQuantity returns a boolean if a field has been set.

### GetRequestedDestination

`func (o *SavedOrder) GetRequestedDestination() DestinationExchange`

GetRequestedDestination returns the RequestedDestination field if non-nil, zero value otherwise.

### GetRequestedDestinationOk

`func (o *SavedOrder) GetRequestedDestinationOk() (*DestinationExchange, bool)`

GetRequestedDestinationOk returns a tuple with the RequestedDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedDestination

`func (o *SavedOrder) SetRequestedDestination(v DestinationExchange)`

SetRequestedDestination sets RequestedDestination field to given value.

### HasRequestedDestination

`func (o *SavedOrder) HasRequestedDestination() bool`

HasRequestedDestination returns a boolean if a field has been set.

### GetSavedOrderId

`func (o *SavedOrder) GetSavedOrderId() float32`

GetSavedOrderId returns the SavedOrderId field if non-nil, zero value otherwise.

### GetSavedOrderIdOk

`func (o *SavedOrder) GetSavedOrderIdOk() (*float32, bool)`

GetSavedOrderIdOk returns a tuple with the SavedOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavedOrderId

`func (o *SavedOrder) SetSavedOrderId(v float32)`

SetSavedOrderId sets SavedOrderId field to given value.

### HasSavedOrderId

`func (o *SavedOrder) HasSavedOrderId() bool`

HasSavedOrderId returns a boolean if a field has been set.

### GetSavedTime

`func (o *SavedOrder) GetSavedTime() string`

GetSavedTime returns the SavedTime field if non-nil, zero value otherwise.

### GetSavedTimeOk

`func (o *SavedOrder) GetSavedTimeOk() (*string, bool)`

GetSavedTimeOk returns a tuple with the SavedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavedTime

`func (o *SavedOrder) SetSavedTime(v string)`

SetSavedTime sets SavedTime field to given value.

### HasSavedTime

`func (o *SavedOrder) HasSavedTime() bool`

HasSavedTime returns a boolean if a field has been set.

### GetSession

`func (o *SavedOrder) GetSession() Session`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *SavedOrder) GetSessionOk() (*Session, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *SavedOrder) SetSession(v Session)`

SetSession sets Session field to given value.

### HasSession

`func (o *SavedOrder) HasSession() bool`

HasSession returns a boolean if a field has been set.

### GetSpecialInstruction

`func (o *SavedOrder) GetSpecialInstruction() SpecialInstruction`

GetSpecialInstruction returns the SpecialInstruction field if non-nil, zero value otherwise.

### GetSpecialInstructionOk

`func (o *SavedOrder) GetSpecialInstructionOk() (*SpecialInstruction, bool)`

GetSpecialInstructionOk returns a tuple with the SpecialInstruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialInstruction

`func (o *SavedOrder) SetSpecialInstruction(v SpecialInstruction)`

SetSpecialInstruction sets SpecialInstruction field to given value.

### HasSpecialInstruction

`func (o *SavedOrder) HasSpecialInstruction() bool`

HasSpecialInstruction returns a boolean if a field has been set.

### GetStatus

`func (o *SavedOrder) GetStatus() OrderStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SavedOrder) GetStatusOk() (*OrderStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SavedOrder) SetStatus(v OrderStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SavedOrder) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDescription

`func (o *SavedOrder) GetStatusDescription() string`

GetStatusDescription returns the StatusDescription field if non-nil, zero value otherwise.

### GetStatusDescriptionOk

`func (o *SavedOrder) GetStatusDescriptionOk() (*string, bool)`

GetStatusDescriptionOk returns a tuple with the StatusDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDescription

`func (o *SavedOrder) SetStatusDescription(v string)`

SetStatusDescription sets StatusDescription field to given value.

### HasStatusDescription

`func (o *SavedOrder) HasStatusDescription() bool`

HasStatusDescription returns a boolean if a field has been set.

### GetStopPrice

`func (o *SavedOrder) GetStopPrice() float32`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *SavedOrder) GetStopPriceOk() (*float32, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *SavedOrder) SetStopPrice(v float32)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *SavedOrder) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.

### GetStopPriceLinkBasis

`func (o *SavedOrder) GetStopPriceLinkBasis() PriceLinkBasis`

GetStopPriceLinkBasis returns the StopPriceLinkBasis field if non-nil, zero value otherwise.

### GetStopPriceLinkBasisOk

`func (o *SavedOrder) GetStopPriceLinkBasisOk() (*PriceLinkBasis, bool)`

GetStopPriceLinkBasisOk returns a tuple with the StopPriceLinkBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPriceLinkBasis

`func (o *SavedOrder) SetStopPriceLinkBasis(v PriceLinkBasis)`

SetStopPriceLinkBasis sets StopPriceLinkBasis field to given value.

### HasStopPriceLinkBasis

`func (o *SavedOrder) HasStopPriceLinkBasis() bool`

HasStopPriceLinkBasis returns a boolean if a field has been set.

### GetStopPriceLinkType

`func (o *SavedOrder) GetStopPriceLinkType() PriceLinkType`

GetStopPriceLinkType returns the StopPriceLinkType field if non-nil, zero value otherwise.

### GetStopPriceLinkTypeOk

`func (o *SavedOrder) GetStopPriceLinkTypeOk() (*PriceLinkType, bool)`

GetStopPriceLinkTypeOk returns a tuple with the StopPriceLinkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPriceLinkType

`func (o *SavedOrder) SetStopPriceLinkType(v PriceLinkType)`

SetStopPriceLinkType sets StopPriceLinkType field to given value.

### HasStopPriceLinkType

`func (o *SavedOrder) HasStopPriceLinkType() bool`

HasStopPriceLinkType returns a boolean if a field has been set.

### GetStopPriceOffset

`func (o *SavedOrder) GetStopPriceOffset() float32`

GetStopPriceOffset returns the StopPriceOffset field if non-nil, zero value otherwise.

### GetStopPriceOffsetOk

`func (o *SavedOrder) GetStopPriceOffsetOk() (*float32, bool)`

GetStopPriceOffsetOk returns a tuple with the StopPriceOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPriceOffset

`func (o *SavedOrder) SetStopPriceOffset(v float32)`

SetStopPriceOffset sets StopPriceOffset field to given value.

### HasStopPriceOffset

`func (o *SavedOrder) HasStopPriceOffset() bool`

HasStopPriceOffset returns a boolean if a field has been set.

### GetStopType

`func (o *SavedOrder) GetStopType() StopType`

GetStopType returns the StopType field if non-nil, zero value otherwise.

### GetStopTypeOk

`func (o *SavedOrder) GetStopTypeOk() (*StopType, bool)`

GetStopTypeOk returns a tuple with the StopType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopType

`func (o *SavedOrder) SetStopType(v StopType)`

SetStopType sets StopType field to given value.

### HasStopType

`func (o *SavedOrder) HasStopType() bool`

HasStopType returns a boolean if a field has been set.

### GetTag

`func (o *SavedOrder) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *SavedOrder) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *SavedOrder) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *SavedOrder) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTaxLotMethod

`func (o *SavedOrder) GetTaxLotMethod() TaxLotMethod`

GetTaxLotMethod returns the TaxLotMethod field if non-nil, zero value otherwise.

### GetTaxLotMethodOk

`func (o *SavedOrder) GetTaxLotMethodOk() (*TaxLotMethod, bool)`

GetTaxLotMethodOk returns a tuple with the TaxLotMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxLotMethod

`func (o *SavedOrder) SetTaxLotMethod(v TaxLotMethod)`

SetTaxLotMethod sets TaxLotMethod field to given value.

### HasTaxLotMethod

`func (o *SavedOrder) HasTaxLotMethod() bool`

HasTaxLotMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


