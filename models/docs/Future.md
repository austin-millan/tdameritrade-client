# Future

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AskId** | Pointer to **string** |  | [optional] 
**AskPriceInDouble** | Pointer to **float32** |  | [optional] 
**BidId** | Pointer to **string** |  | [optional] 
**BidPriceInDouble** | Pointer to **float32** |  | [optional] 
**ChangeInDouble** | Pointer to **float32** |  | [optional] 
**ClosePriceInDouble** | Pointer to **float32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Exchange** | Pointer to **string** |  | [optional] 
**ExchangeName** | Pointer to **string** |  | [optional] 
**FutureActiveSymbol** | Pointer to **string** |  | [optional] 
**FutureExpirationDate** | Pointer to **string** |  | [optional] 
**FutureIsActive** | Pointer to **bool** |  | [optional] 
**FutureIsTradable** | Pointer to **bool** |  | [optional] 
**FutureMultiplier** | Pointer to **float32** |  | [optional] 
**FuturePercentChange** | Pointer to **float32** |  | [optional] 
**FuturePriceFormat** | Pointer to **string** |  | [optional] 
**FutureSettlementPrice** | Pointer to **float32** |  | [optional] 
**FutureTradingHours** | Pointer to **string** |  | [optional] 
**HighPriceInDouble** | Pointer to **float32** |  | [optional] 
**LastId** | Pointer to **string** |  | [optional] 
**LastPriceInDouble** | Pointer to **float32** |  | [optional] 
**LowPriceInDouble** | Pointer to **float32** |  | [optional] 
**Mark** | Pointer to **float32** |  | [optional] 
**OpenInterest** | Pointer to **float32** |  | [optional] 
**OpenPriceInDouble** | Pointer to **float32** |  | [optional] 
**Product** | Pointer to **string** |  | [optional] 
**SecurityStatus** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Tick** | Pointer to **float32** |  | [optional] 
**TickAmount** | Pointer to **float32** |  | [optional] 

## Methods

### NewFuture

`func NewFuture() *Future`

NewFuture instantiates a new Future object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFutureWithDefaults

`func NewFutureWithDefaults() *Future`

NewFutureWithDefaults instantiates a new Future object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAskId

`func (o *Future) GetAskId() string`

GetAskId returns the AskId field if non-nil, zero value otherwise.

### GetAskIdOk

`func (o *Future) GetAskIdOk() (*string, bool)`

GetAskIdOk returns a tuple with the AskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskId

`func (o *Future) SetAskId(v string)`

SetAskId sets AskId field to given value.

### HasAskId

`func (o *Future) HasAskId() bool`

HasAskId returns a boolean if a field has been set.

### GetAskPriceInDouble

`func (o *Future) GetAskPriceInDouble() float32`

GetAskPriceInDouble returns the AskPriceInDouble field if non-nil, zero value otherwise.

### GetAskPriceInDoubleOk

`func (o *Future) GetAskPriceInDoubleOk() (*float32, bool)`

GetAskPriceInDoubleOk returns a tuple with the AskPriceInDouble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskPriceInDouble

`func (o *Future) SetAskPriceInDouble(v float32)`

SetAskPriceInDouble sets AskPriceInDouble field to given value.

### HasAskPriceInDouble

`func (o *Future) HasAskPriceInDouble() bool`

HasAskPriceInDouble returns a boolean if a field has been set.

### GetBidId

`func (o *Future) GetBidId() string`

GetBidId returns the BidId field if non-nil, zero value otherwise.

### GetBidIdOk

`func (o *Future) GetBidIdOk() (*string, bool)`

GetBidIdOk returns a tuple with the BidId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidId

`func (o *Future) SetBidId(v string)`

SetBidId sets BidId field to given value.

### HasBidId

`func (o *Future) HasBidId() bool`

HasBidId returns a boolean if a field has been set.

### GetBidPriceInDouble

`func (o *Future) GetBidPriceInDouble() float32`

GetBidPriceInDouble returns the BidPriceInDouble field if non-nil, zero value otherwise.

### GetBidPriceInDoubleOk

`func (o *Future) GetBidPriceInDoubleOk() (*float32, bool)`

GetBidPriceInDoubleOk returns a tuple with the BidPriceInDouble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidPriceInDouble

`func (o *Future) SetBidPriceInDouble(v float32)`

SetBidPriceInDouble sets BidPriceInDouble field to given value.

### HasBidPriceInDouble

`func (o *Future) HasBidPriceInDouble() bool`

HasBidPriceInDouble returns a boolean if a field has been set.

### GetChangeInDouble

`func (o *Future) GetChangeInDouble() float32`

GetChangeInDouble returns the ChangeInDouble field if non-nil, zero value otherwise.

### GetChangeInDoubleOk

`func (o *Future) GetChangeInDoubleOk() (*float32, bool)`

GetChangeInDoubleOk returns a tuple with the ChangeInDouble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeInDouble

`func (o *Future) SetChangeInDouble(v float32)`

SetChangeInDouble sets ChangeInDouble field to given value.

### HasChangeInDouble

`func (o *Future) HasChangeInDouble() bool`

HasChangeInDouble returns a boolean if a field has been set.

### GetClosePriceInDouble

`func (o *Future) GetClosePriceInDouble() float32`

GetClosePriceInDouble returns the ClosePriceInDouble field if non-nil, zero value otherwise.

### GetClosePriceInDoubleOk

`func (o *Future) GetClosePriceInDoubleOk() (*float32, bool)`

GetClosePriceInDoubleOk returns a tuple with the ClosePriceInDouble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosePriceInDouble

`func (o *Future) SetClosePriceInDouble(v float32)`

SetClosePriceInDouble sets ClosePriceInDouble field to given value.

### HasClosePriceInDouble

`func (o *Future) HasClosePriceInDouble() bool`

HasClosePriceInDouble returns a boolean if a field has been set.

### GetDescription

`func (o *Future) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Future) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Future) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Future) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExchange

`func (o *Future) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *Future) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *Future) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *Future) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetExchangeName

`func (o *Future) GetExchangeName() string`

GetExchangeName returns the ExchangeName field if non-nil, zero value otherwise.

### GetExchangeNameOk

`func (o *Future) GetExchangeNameOk() (*string, bool)`

GetExchangeNameOk returns a tuple with the ExchangeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeName

`func (o *Future) SetExchangeName(v string)`

SetExchangeName sets ExchangeName field to given value.

### HasExchangeName

`func (o *Future) HasExchangeName() bool`

HasExchangeName returns a boolean if a field has been set.

### GetFutureActiveSymbol

`func (o *Future) GetFutureActiveSymbol() string`

GetFutureActiveSymbol returns the FutureActiveSymbol field if non-nil, zero value otherwise.

### GetFutureActiveSymbolOk

`func (o *Future) GetFutureActiveSymbolOk() (*string, bool)`

GetFutureActiveSymbolOk returns a tuple with the FutureActiveSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFutureActiveSymbol

`func (o *Future) SetFutureActiveSymbol(v string)`

SetFutureActiveSymbol sets FutureActiveSymbol field to given value.

### HasFutureActiveSymbol

`func (o *Future) HasFutureActiveSymbol() bool`

HasFutureActiveSymbol returns a boolean if a field has been set.

### GetFutureExpirationDate

`func (o *Future) GetFutureExpirationDate() string`

GetFutureExpirationDate returns the FutureExpirationDate field if non-nil, zero value otherwise.

### GetFutureExpirationDateOk

`func (o *Future) GetFutureExpirationDateOk() (*string, bool)`

GetFutureExpirationDateOk returns a tuple with the FutureExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFutureExpirationDate

`func (o *Future) SetFutureExpirationDate(v string)`

SetFutureExpirationDate sets FutureExpirationDate field to given value.

### HasFutureExpirationDate

`func (o *Future) HasFutureExpirationDate() bool`

HasFutureExpirationDate returns a boolean if a field has been set.

### GetFutureIsActive

`func (o *Future) GetFutureIsActive() bool`

GetFutureIsActive returns the FutureIsActive field if non-nil, zero value otherwise.

### GetFutureIsActiveOk

`func (o *Future) GetFutureIsActiveOk() (*bool, bool)`

GetFutureIsActiveOk returns a tuple with the FutureIsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFutureIsActive

`func (o *Future) SetFutureIsActive(v bool)`

SetFutureIsActive sets FutureIsActive field to given value.

### HasFutureIsActive

`func (o *Future) HasFutureIsActive() bool`

HasFutureIsActive returns a boolean if a field has been set.

### GetFutureIsTradable

`func (o *Future) GetFutureIsTradable() bool`

GetFutureIsTradable returns the FutureIsTradable field if non-nil, zero value otherwise.

### GetFutureIsTradableOk

`func (o *Future) GetFutureIsTradableOk() (*bool, bool)`

GetFutureIsTradableOk returns a tuple with the FutureIsTradable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFutureIsTradable

`func (o *Future) SetFutureIsTradable(v bool)`

SetFutureIsTradable sets FutureIsTradable field to given value.

### HasFutureIsTradable

`func (o *Future) HasFutureIsTradable() bool`

HasFutureIsTradable returns a boolean if a field has been set.

### GetFutureMultiplier

`func (o *Future) GetFutureMultiplier() float32`

GetFutureMultiplier returns the FutureMultiplier field if non-nil, zero value otherwise.

### GetFutureMultiplierOk

`func (o *Future) GetFutureMultiplierOk() (*float32, bool)`

GetFutureMultiplierOk returns a tuple with the FutureMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFutureMultiplier

`func (o *Future) SetFutureMultiplier(v float32)`

SetFutureMultiplier sets FutureMultiplier field to given value.

### HasFutureMultiplier

`func (o *Future) HasFutureMultiplier() bool`

HasFutureMultiplier returns a boolean if a field has been set.

### GetFuturePercentChange

`func (o *Future) GetFuturePercentChange() float32`

GetFuturePercentChange returns the FuturePercentChange field if non-nil, zero value otherwise.

### GetFuturePercentChangeOk

`func (o *Future) GetFuturePercentChangeOk() (*float32, bool)`

GetFuturePercentChangeOk returns a tuple with the FuturePercentChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuturePercentChange

`func (o *Future) SetFuturePercentChange(v float32)`

SetFuturePercentChange sets FuturePercentChange field to given value.

### HasFuturePercentChange

`func (o *Future) HasFuturePercentChange() bool`

HasFuturePercentChange returns a boolean if a field has been set.

### GetFuturePriceFormat

`func (o *Future) GetFuturePriceFormat() string`

GetFuturePriceFormat returns the FuturePriceFormat field if non-nil, zero value otherwise.

### GetFuturePriceFormatOk

`func (o *Future) GetFuturePriceFormatOk() (*string, bool)`

GetFuturePriceFormatOk returns a tuple with the FuturePriceFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuturePriceFormat

`func (o *Future) SetFuturePriceFormat(v string)`

SetFuturePriceFormat sets FuturePriceFormat field to given value.

### HasFuturePriceFormat

`func (o *Future) HasFuturePriceFormat() bool`

HasFuturePriceFormat returns a boolean if a field has been set.

### GetFutureSettlementPrice

`func (o *Future) GetFutureSettlementPrice() float32`

GetFutureSettlementPrice returns the FutureSettlementPrice field if non-nil, zero value otherwise.

### GetFutureSettlementPriceOk

`func (o *Future) GetFutureSettlementPriceOk() (*float32, bool)`

GetFutureSettlementPriceOk returns a tuple with the FutureSettlementPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFutureSettlementPrice

`func (o *Future) SetFutureSettlementPrice(v float32)`

SetFutureSettlementPrice sets FutureSettlementPrice field to given value.

### HasFutureSettlementPrice

`func (o *Future) HasFutureSettlementPrice() bool`

HasFutureSettlementPrice returns a boolean if a field has been set.

### GetFutureTradingHours

`func (o *Future) GetFutureTradingHours() string`

GetFutureTradingHours returns the FutureTradingHours field if non-nil, zero value otherwise.

### GetFutureTradingHoursOk

`func (o *Future) GetFutureTradingHoursOk() (*string, bool)`

GetFutureTradingHoursOk returns a tuple with the FutureTradingHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFutureTradingHours

`func (o *Future) SetFutureTradingHours(v string)`

SetFutureTradingHours sets FutureTradingHours field to given value.

### HasFutureTradingHours

`func (o *Future) HasFutureTradingHours() bool`

HasFutureTradingHours returns a boolean if a field has been set.

### GetHighPriceInDouble

`func (o *Future) GetHighPriceInDouble() float32`

GetHighPriceInDouble returns the HighPriceInDouble field if non-nil, zero value otherwise.

### GetHighPriceInDoubleOk

`func (o *Future) GetHighPriceInDoubleOk() (*float32, bool)`

GetHighPriceInDoubleOk returns a tuple with the HighPriceInDouble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighPriceInDouble

`func (o *Future) SetHighPriceInDouble(v float32)`

SetHighPriceInDouble sets HighPriceInDouble field to given value.

### HasHighPriceInDouble

`func (o *Future) HasHighPriceInDouble() bool`

HasHighPriceInDouble returns a boolean if a field has been set.

### GetLastId

`func (o *Future) GetLastId() string`

GetLastId returns the LastId field if non-nil, zero value otherwise.

### GetLastIdOk

`func (o *Future) GetLastIdOk() (*string, bool)`

GetLastIdOk returns a tuple with the LastId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastId

`func (o *Future) SetLastId(v string)`

SetLastId sets LastId field to given value.

### HasLastId

`func (o *Future) HasLastId() bool`

HasLastId returns a boolean if a field has been set.

### GetLastPriceInDouble

`func (o *Future) GetLastPriceInDouble() float32`

GetLastPriceInDouble returns the LastPriceInDouble field if non-nil, zero value otherwise.

### GetLastPriceInDoubleOk

`func (o *Future) GetLastPriceInDoubleOk() (*float32, bool)`

GetLastPriceInDoubleOk returns a tuple with the LastPriceInDouble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPriceInDouble

`func (o *Future) SetLastPriceInDouble(v float32)`

SetLastPriceInDouble sets LastPriceInDouble field to given value.

### HasLastPriceInDouble

`func (o *Future) HasLastPriceInDouble() bool`

HasLastPriceInDouble returns a boolean if a field has been set.

### GetLowPriceInDouble

`func (o *Future) GetLowPriceInDouble() float32`

GetLowPriceInDouble returns the LowPriceInDouble field if non-nil, zero value otherwise.

### GetLowPriceInDoubleOk

`func (o *Future) GetLowPriceInDoubleOk() (*float32, bool)`

GetLowPriceInDoubleOk returns a tuple with the LowPriceInDouble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowPriceInDouble

`func (o *Future) SetLowPriceInDouble(v float32)`

SetLowPriceInDouble sets LowPriceInDouble field to given value.

### HasLowPriceInDouble

`func (o *Future) HasLowPriceInDouble() bool`

HasLowPriceInDouble returns a boolean if a field has been set.

### GetMark

`func (o *Future) GetMark() float32`

GetMark returns the Mark field if non-nil, zero value otherwise.

### GetMarkOk

`func (o *Future) GetMarkOk() (*float32, bool)`

GetMarkOk returns a tuple with the Mark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMark

`func (o *Future) SetMark(v float32)`

SetMark sets Mark field to given value.

### HasMark

`func (o *Future) HasMark() bool`

HasMark returns a boolean if a field has been set.

### GetOpenInterest

`func (o *Future) GetOpenInterest() float32`

GetOpenInterest returns the OpenInterest field if non-nil, zero value otherwise.

### GetOpenInterestOk

`func (o *Future) GetOpenInterestOk() (*float32, bool)`

GetOpenInterestOk returns a tuple with the OpenInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenInterest

`func (o *Future) SetOpenInterest(v float32)`

SetOpenInterest sets OpenInterest field to given value.

### HasOpenInterest

`func (o *Future) HasOpenInterest() bool`

HasOpenInterest returns a boolean if a field has been set.

### GetOpenPriceInDouble

`func (o *Future) GetOpenPriceInDouble() float32`

GetOpenPriceInDouble returns the OpenPriceInDouble field if non-nil, zero value otherwise.

### GetOpenPriceInDoubleOk

`func (o *Future) GetOpenPriceInDoubleOk() (*float32, bool)`

GetOpenPriceInDoubleOk returns a tuple with the OpenPriceInDouble field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenPriceInDouble

`func (o *Future) SetOpenPriceInDouble(v float32)`

SetOpenPriceInDouble sets OpenPriceInDouble field to given value.

### HasOpenPriceInDouble

`func (o *Future) HasOpenPriceInDouble() bool`

HasOpenPriceInDouble returns a boolean if a field has been set.

### GetProduct

`func (o *Future) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *Future) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *Future) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *Future) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetSecurityStatus

`func (o *Future) GetSecurityStatus() string`

GetSecurityStatus returns the SecurityStatus field if non-nil, zero value otherwise.

### GetSecurityStatusOk

`func (o *Future) GetSecurityStatusOk() (*string, bool)`

GetSecurityStatusOk returns a tuple with the SecurityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityStatus

`func (o *Future) SetSecurityStatus(v string)`

SetSecurityStatus sets SecurityStatus field to given value.

### HasSecurityStatus

`func (o *Future) HasSecurityStatus() bool`

HasSecurityStatus returns a boolean if a field has been set.

### GetSymbol

`func (o *Future) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Future) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Future) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *Future) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTick

`func (o *Future) GetTick() float32`

GetTick returns the Tick field if non-nil, zero value otherwise.

### GetTickOk

`func (o *Future) GetTickOk() (*float32, bool)`

GetTickOk returns a tuple with the Tick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTick

`func (o *Future) SetTick(v float32)`

SetTick sets Tick field to given value.

### HasTick

`func (o *Future) HasTick() bool`

HasTick returns a boolean if a field has been set.

### GetTickAmount

`func (o *Future) GetTickAmount() float32`

GetTickAmount returns the TickAmount field if non-nil, zero value otherwise.

### GetTickAmountOk

`func (o *Future) GetTickAmountOk() (*float32, bool)`

GetTickAmountOk returns a tuple with the TickAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickAmount

`func (o *Future) SetTickAmount(v float32)`

SetTickAmount sets TickAmount field to given value.

### HasTickAmount

`func (o *Future) HasTickAmount() bool`

HasTickAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


