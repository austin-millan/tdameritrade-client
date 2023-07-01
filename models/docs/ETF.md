# ETF

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Var52WkHigh** | Pointer to **float32** |  | [optional] 
**Var52WkLow** | Pointer to **float32** |  | [optional] 
**AskId** | Pointer to **string** |  | [optional] 
**AskPrice** | Pointer to **float32** |  | [optional] 
**AskSize** | Pointer to **float32** |  | [optional] 
**BidId** | Pointer to **string** |  | [optional] 
**BidPrice** | Pointer to **float32** |  | [optional] 
**BidSize** | Pointer to **float32** |  | [optional] 
**ClosePrice** | Pointer to **float32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Digits** | Pointer to **float32** |  | [optional] 
**DivAmount** | Pointer to **float32** |  | [optional] 
**DivDate** | Pointer to **string** |  | [optional] 
**DivYield** | Pointer to **float32** |  | [optional] 
**Exchange** | Pointer to **string** |  | [optional] 
**ExchangeName** | Pointer to **string** |  | [optional] 
**HighPrice** | Pointer to **float32** |  | [optional] 
**LastId** | Pointer to **string** |  | [optional] 
**LastPrice** | Pointer to **float32** |  | [optional] 
**LastSize** | Pointer to **float32** |  | [optional] 
**LowPrice** | Pointer to **float32** |  | [optional] 
**Marginable** | Pointer to **bool** |  | [optional] 
**Mark** | Pointer to **float32** |  | [optional] 
**NetChange** | Pointer to **float32** |  | [optional] 
**OpenPrice** | Pointer to **float32** |  | [optional] 
**PeRatio** | Pointer to **float32** |  | [optional] 
**QuoteTimeInLong** | Pointer to **float32** |  | [optional] 
**RegularMarketLastPrice** | Pointer to **float32** |  | [optional] 
**RegularMarketLastSize** | Pointer to **float32** |  | [optional] 
**RegularMarketNetChange** | Pointer to **float32** |  | [optional] 
**RegularMarketTradeTimeInLong** | Pointer to **float32** |  | [optional] 
**SecurityStatus** | Pointer to **string** |  | [optional] 
**Shortable** | Pointer to **bool** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**TotalVolume** | Pointer to **float32** |  | [optional] 
**TradeTimeInLong** | Pointer to **float32** |  | [optional] 
**Volatility** | Pointer to **float32** |  | [optional] 

## Methods

### NewETF

`func NewETF() *ETF`

NewETF instantiates a new ETF object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewETFWithDefaults

`func NewETFWithDefaults() *ETF`

NewETFWithDefaults instantiates a new ETF object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVar52WkHigh

`func (o *ETF) GetVar52WkHigh() float32`

GetVar52WkHigh returns the Var52WkHigh field if non-nil, zero value otherwise.

### GetVar52WkHighOk

`func (o *ETF) GetVar52WkHighOk() (*float32, bool)`

GetVar52WkHighOk returns a tuple with the Var52WkHigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar52WkHigh

`func (o *ETF) SetVar52WkHigh(v float32)`

SetVar52WkHigh sets Var52WkHigh field to given value.

### HasVar52WkHigh

`func (o *ETF) HasVar52WkHigh() bool`

HasVar52WkHigh returns a boolean if a field has been set.

### GetVar52WkLow

`func (o *ETF) GetVar52WkLow() float32`

GetVar52WkLow returns the Var52WkLow field if non-nil, zero value otherwise.

### GetVar52WkLowOk

`func (o *ETF) GetVar52WkLowOk() (*float32, bool)`

GetVar52WkLowOk returns a tuple with the Var52WkLow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar52WkLow

`func (o *ETF) SetVar52WkLow(v float32)`

SetVar52WkLow sets Var52WkLow field to given value.

### HasVar52WkLow

`func (o *ETF) HasVar52WkLow() bool`

HasVar52WkLow returns a boolean if a field has been set.

### GetAskId

`func (o *ETF) GetAskId() string`

GetAskId returns the AskId field if non-nil, zero value otherwise.

### GetAskIdOk

`func (o *ETF) GetAskIdOk() (*string, bool)`

GetAskIdOk returns a tuple with the AskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskId

`func (o *ETF) SetAskId(v string)`

SetAskId sets AskId field to given value.

### HasAskId

`func (o *ETF) HasAskId() bool`

HasAskId returns a boolean if a field has been set.

### GetAskPrice

`func (o *ETF) GetAskPrice() float32`

GetAskPrice returns the AskPrice field if non-nil, zero value otherwise.

### GetAskPriceOk

`func (o *ETF) GetAskPriceOk() (*float32, bool)`

GetAskPriceOk returns a tuple with the AskPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskPrice

`func (o *ETF) SetAskPrice(v float32)`

SetAskPrice sets AskPrice field to given value.

### HasAskPrice

`func (o *ETF) HasAskPrice() bool`

HasAskPrice returns a boolean if a field has been set.

### GetAskSize

`func (o *ETF) GetAskSize() float32`

GetAskSize returns the AskSize field if non-nil, zero value otherwise.

### GetAskSizeOk

`func (o *ETF) GetAskSizeOk() (*float32, bool)`

GetAskSizeOk returns a tuple with the AskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskSize

`func (o *ETF) SetAskSize(v float32)`

SetAskSize sets AskSize field to given value.

### HasAskSize

`func (o *ETF) HasAskSize() bool`

HasAskSize returns a boolean if a field has been set.

### GetBidId

`func (o *ETF) GetBidId() string`

GetBidId returns the BidId field if non-nil, zero value otherwise.

### GetBidIdOk

`func (o *ETF) GetBidIdOk() (*string, bool)`

GetBidIdOk returns a tuple with the BidId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidId

`func (o *ETF) SetBidId(v string)`

SetBidId sets BidId field to given value.

### HasBidId

`func (o *ETF) HasBidId() bool`

HasBidId returns a boolean if a field has been set.

### GetBidPrice

`func (o *ETF) GetBidPrice() float32`

GetBidPrice returns the BidPrice field if non-nil, zero value otherwise.

### GetBidPriceOk

`func (o *ETF) GetBidPriceOk() (*float32, bool)`

GetBidPriceOk returns a tuple with the BidPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidPrice

`func (o *ETF) SetBidPrice(v float32)`

SetBidPrice sets BidPrice field to given value.

### HasBidPrice

`func (o *ETF) HasBidPrice() bool`

HasBidPrice returns a boolean if a field has been set.

### GetBidSize

`func (o *ETF) GetBidSize() float32`

GetBidSize returns the BidSize field if non-nil, zero value otherwise.

### GetBidSizeOk

`func (o *ETF) GetBidSizeOk() (*float32, bool)`

GetBidSizeOk returns a tuple with the BidSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidSize

`func (o *ETF) SetBidSize(v float32)`

SetBidSize sets BidSize field to given value.

### HasBidSize

`func (o *ETF) HasBidSize() bool`

HasBidSize returns a boolean if a field has been set.

### GetClosePrice

`func (o *ETF) GetClosePrice() float32`

GetClosePrice returns the ClosePrice field if non-nil, zero value otherwise.

### GetClosePriceOk

`func (o *ETF) GetClosePriceOk() (*float32, bool)`

GetClosePriceOk returns a tuple with the ClosePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosePrice

`func (o *ETF) SetClosePrice(v float32)`

SetClosePrice sets ClosePrice field to given value.

### HasClosePrice

`func (o *ETF) HasClosePrice() bool`

HasClosePrice returns a boolean if a field has been set.

### GetDescription

`func (o *ETF) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ETF) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ETF) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ETF) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDigits

`func (o *ETF) GetDigits() float32`

GetDigits returns the Digits field if non-nil, zero value otherwise.

### GetDigitsOk

`func (o *ETF) GetDigitsOk() (*float32, bool)`

GetDigitsOk returns a tuple with the Digits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigits

`func (o *ETF) SetDigits(v float32)`

SetDigits sets Digits field to given value.

### HasDigits

`func (o *ETF) HasDigits() bool`

HasDigits returns a boolean if a field has been set.

### GetDivAmount

`func (o *ETF) GetDivAmount() float32`

GetDivAmount returns the DivAmount field if non-nil, zero value otherwise.

### GetDivAmountOk

`func (o *ETF) GetDivAmountOk() (*float32, bool)`

GetDivAmountOk returns a tuple with the DivAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivAmount

`func (o *ETF) SetDivAmount(v float32)`

SetDivAmount sets DivAmount field to given value.

### HasDivAmount

`func (o *ETF) HasDivAmount() bool`

HasDivAmount returns a boolean if a field has been set.

### GetDivDate

`func (o *ETF) GetDivDate() string`

GetDivDate returns the DivDate field if non-nil, zero value otherwise.

### GetDivDateOk

`func (o *ETF) GetDivDateOk() (*string, bool)`

GetDivDateOk returns a tuple with the DivDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivDate

`func (o *ETF) SetDivDate(v string)`

SetDivDate sets DivDate field to given value.

### HasDivDate

`func (o *ETF) HasDivDate() bool`

HasDivDate returns a boolean if a field has been set.

### GetDivYield

`func (o *ETF) GetDivYield() float32`

GetDivYield returns the DivYield field if non-nil, zero value otherwise.

### GetDivYieldOk

`func (o *ETF) GetDivYieldOk() (*float32, bool)`

GetDivYieldOk returns a tuple with the DivYield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivYield

`func (o *ETF) SetDivYield(v float32)`

SetDivYield sets DivYield field to given value.

### HasDivYield

`func (o *ETF) HasDivYield() bool`

HasDivYield returns a boolean if a field has been set.

### GetExchange

`func (o *ETF) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *ETF) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *ETF) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *ETF) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetExchangeName

`func (o *ETF) GetExchangeName() string`

GetExchangeName returns the ExchangeName field if non-nil, zero value otherwise.

### GetExchangeNameOk

`func (o *ETF) GetExchangeNameOk() (*string, bool)`

GetExchangeNameOk returns a tuple with the ExchangeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeName

`func (o *ETF) SetExchangeName(v string)`

SetExchangeName sets ExchangeName field to given value.

### HasExchangeName

`func (o *ETF) HasExchangeName() bool`

HasExchangeName returns a boolean if a field has been set.

### GetHighPrice

`func (o *ETF) GetHighPrice() float32`

GetHighPrice returns the HighPrice field if non-nil, zero value otherwise.

### GetHighPriceOk

`func (o *ETF) GetHighPriceOk() (*float32, bool)`

GetHighPriceOk returns a tuple with the HighPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighPrice

`func (o *ETF) SetHighPrice(v float32)`

SetHighPrice sets HighPrice field to given value.

### HasHighPrice

`func (o *ETF) HasHighPrice() bool`

HasHighPrice returns a boolean if a field has been set.

### GetLastId

`func (o *ETF) GetLastId() string`

GetLastId returns the LastId field if non-nil, zero value otherwise.

### GetLastIdOk

`func (o *ETF) GetLastIdOk() (*string, bool)`

GetLastIdOk returns a tuple with the LastId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastId

`func (o *ETF) SetLastId(v string)`

SetLastId sets LastId field to given value.

### HasLastId

`func (o *ETF) HasLastId() bool`

HasLastId returns a boolean if a field has been set.

### GetLastPrice

`func (o *ETF) GetLastPrice() float32`

GetLastPrice returns the LastPrice field if non-nil, zero value otherwise.

### GetLastPriceOk

`func (o *ETF) GetLastPriceOk() (*float32, bool)`

GetLastPriceOk returns a tuple with the LastPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPrice

`func (o *ETF) SetLastPrice(v float32)`

SetLastPrice sets LastPrice field to given value.

### HasLastPrice

`func (o *ETF) HasLastPrice() bool`

HasLastPrice returns a boolean if a field has been set.

### GetLastSize

`func (o *ETF) GetLastSize() float32`

GetLastSize returns the LastSize field if non-nil, zero value otherwise.

### GetLastSizeOk

`func (o *ETF) GetLastSizeOk() (*float32, bool)`

GetLastSizeOk returns a tuple with the LastSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSize

`func (o *ETF) SetLastSize(v float32)`

SetLastSize sets LastSize field to given value.

### HasLastSize

`func (o *ETF) HasLastSize() bool`

HasLastSize returns a boolean if a field has been set.

### GetLowPrice

`func (o *ETF) GetLowPrice() float32`

GetLowPrice returns the LowPrice field if non-nil, zero value otherwise.

### GetLowPriceOk

`func (o *ETF) GetLowPriceOk() (*float32, bool)`

GetLowPriceOk returns a tuple with the LowPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowPrice

`func (o *ETF) SetLowPrice(v float32)`

SetLowPrice sets LowPrice field to given value.

### HasLowPrice

`func (o *ETF) HasLowPrice() bool`

HasLowPrice returns a boolean if a field has been set.

### GetMarginable

`func (o *ETF) GetMarginable() bool`

GetMarginable returns the Marginable field if non-nil, zero value otherwise.

### GetMarginableOk

`func (o *ETF) GetMarginableOk() (*bool, bool)`

GetMarginableOk returns a tuple with the Marginable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginable

`func (o *ETF) SetMarginable(v bool)`

SetMarginable sets Marginable field to given value.

### HasMarginable

`func (o *ETF) HasMarginable() bool`

HasMarginable returns a boolean if a field has been set.

### GetMark

`func (o *ETF) GetMark() float32`

GetMark returns the Mark field if non-nil, zero value otherwise.

### GetMarkOk

`func (o *ETF) GetMarkOk() (*float32, bool)`

GetMarkOk returns a tuple with the Mark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMark

`func (o *ETF) SetMark(v float32)`

SetMark sets Mark field to given value.

### HasMark

`func (o *ETF) HasMark() bool`

HasMark returns a boolean if a field has been set.

### GetNetChange

`func (o *ETF) GetNetChange() float32`

GetNetChange returns the NetChange field if non-nil, zero value otherwise.

### GetNetChangeOk

`func (o *ETF) GetNetChangeOk() (*float32, bool)`

GetNetChangeOk returns a tuple with the NetChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetChange

`func (o *ETF) SetNetChange(v float32)`

SetNetChange sets NetChange field to given value.

### HasNetChange

`func (o *ETF) HasNetChange() bool`

HasNetChange returns a boolean if a field has been set.

### GetOpenPrice

`func (o *ETF) GetOpenPrice() float32`

GetOpenPrice returns the OpenPrice field if non-nil, zero value otherwise.

### GetOpenPriceOk

`func (o *ETF) GetOpenPriceOk() (*float32, bool)`

GetOpenPriceOk returns a tuple with the OpenPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenPrice

`func (o *ETF) SetOpenPrice(v float32)`

SetOpenPrice sets OpenPrice field to given value.

### HasOpenPrice

`func (o *ETF) HasOpenPrice() bool`

HasOpenPrice returns a boolean if a field has been set.

### GetPeRatio

`func (o *ETF) GetPeRatio() float32`

GetPeRatio returns the PeRatio field if non-nil, zero value otherwise.

### GetPeRatioOk

`func (o *ETF) GetPeRatioOk() (*float32, bool)`

GetPeRatioOk returns a tuple with the PeRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeRatio

`func (o *ETF) SetPeRatio(v float32)`

SetPeRatio sets PeRatio field to given value.

### HasPeRatio

`func (o *ETF) HasPeRatio() bool`

HasPeRatio returns a boolean if a field has been set.

### GetQuoteTimeInLong

`func (o *ETF) GetQuoteTimeInLong() float32`

GetQuoteTimeInLong returns the QuoteTimeInLong field if non-nil, zero value otherwise.

### GetQuoteTimeInLongOk

`func (o *ETF) GetQuoteTimeInLongOk() (*float32, bool)`

GetQuoteTimeInLongOk returns a tuple with the QuoteTimeInLong field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteTimeInLong

`func (o *ETF) SetQuoteTimeInLong(v float32)`

SetQuoteTimeInLong sets QuoteTimeInLong field to given value.

### HasQuoteTimeInLong

`func (o *ETF) HasQuoteTimeInLong() bool`

HasQuoteTimeInLong returns a boolean if a field has been set.

### GetRegularMarketLastPrice

`func (o *ETF) GetRegularMarketLastPrice() float32`

GetRegularMarketLastPrice returns the RegularMarketLastPrice field if non-nil, zero value otherwise.

### GetRegularMarketLastPriceOk

`func (o *ETF) GetRegularMarketLastPriceOk() (*float32, bool)`

GetRegularMarketLastPriceOk returns a tuple with the RegularMarketLastPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegularMarketLastPrice

`func (o *ETF) SetRegularMarketLastPrice(v float32)`

SetRegularMarketLastPrice sets RegularMarketLastPrice field to given value.

### HasRegularMarketLastPrice

`func (o *ETF) HasRegularMarketLastPrice() bool`

HasRegularMarketLastPrice returns a boolean if a field has been set.

### GetRegularMarketLastSize

`func (o *ETF) GetRegularMarketLastSize() float32`

GetRegularMarketLastSize returns the RegularMarketLastSize field if non-nil, zero value otherwise.

### GetRegularMarketLastSizeOk

`func (o *ETF) GetRegularMarketLastSizeOk() (*float32, bool)`

GetRegularMarketLastSizeOk returns a tuple with the RegularMarketLastSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegularMarketLastSize

`func (o *ETF) SetRegularMarketLastSize(v float32)`

SetRegularMarketLastSize sets RegularMarketLastSize field to given value.

### HasRegularMarketLastSize

`func (o *ETF) HasRegularMarketLastSize() bool`

HasRegularMarketLastSize returns a boolean if a field has been set.

### GetRegularMarketNetChange

`func (o *ETF) GetRegularMarketNetChange() float32`

GetRegularMarketNetChange returns the RegularMarketNetChange field if non-nil, zero value otherwise.

### GetRegularMarketNetChangeOk

`func (o *ETF) GetRegularMarketNetChangeOk() (*float32, bool)`

GetRegularMarketNetChangeOk returns a tuple with the RegularMarketNetChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegularMarketNetChange

`func (o *ETF) SetRegularMarketNetChange(v float32)`

SetRegularMarketNetChange sets RegularMarketNetChange field to given value.

### HasRegularMarketNetChange

`func (o *ETF) HasRegularMarketNetChange() bool`

HasRegularMarketNetChange returns a boolean if a field has been set.

### GetRegularMarketTradeTimeInLong

`func (o *ETF) GetRegularMarketTradeTimeInLong() float32`

GetRegularMarketTradeTimeInLong returns the RegularMarketTradeTimeInLong field if non-nil, zero value otherwise.

### GetRegularMarketTradeTimeInLongOk

`func (o *ETF) GetRegularMarketTradeTimeInLongOk() (*float32, bool)`

GetRegularMarketTradeTimeInLongOk returns a tuple with the RegularMarketTradeTimeInLong field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegularMarketTradeTimeInLong

`func (o *ETF) SetRegularMarketTradeTimeInLong(v float32)`

SetRegularMarketTradeTimeInLong sets RegularMarketTradeTimeInLong field to given value.

### HasRegularMarketTradeTimeInLong

`func (o *ETF) HasRegularMarketTradeTimeInLong() bool`

HasRegularMarketTradeTimeInLong returns a boolean if a field has been set.

### GetSecurityStatus

`func (o *ETF) GetSecurityStatus() string`

GetSecurityStatus returns the SecurityStatus field if non-nil, zero value otherwise.

### GetSecurityStatusOk

`func (o *ETF) GetSecurityStatusOk() (*string, bool)`

GetSecurityStatusOk returns a tuple with the SecurityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityStatus

`func (o *ETF) SetSecurityStatus(v string)`

SetSecurityStatus sets SecurityStatus field to given value.

### HasSecurityStatus

`func (o *ETF) HasSecurityStatus() bool`

HasSecurityStatus returns a boolean if a field has been set.

### GetShortable

`func (o *ETF) GetShortable() bool`

GetShortable returns the Shortable field if non-nil, zero value otherwise.

### GetShortableOk

`func (o *ETF) GetShortableOk() (*bool, bool)`

GetShortableOk returns a tuple with the Shortable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortable

`func (o *ETF) SetShortable(v bool)`

SetShortable sets Shortable field to given value.

### HasShortable

`func (o *ETF) HasShortable() bool`

HasShortable returns a boolean if a field has been set.

### GetSymbol

`func (o *ETF) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ETF) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ETF) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *ETF) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTotalVolume

`func (o *ETF) GetTotalVolume() float32`

GetTotalVolume returns the TotalVolume field if non-nil, zero value otherwise.

### GetTotalVolumeOk

`func (o *ETF) GetTotalVolumeOk() (*float32, bool)`

GetTotalVolumeOk returns a tuple with the TotalVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVolume

`func (o *ETF) SetTotalVolume(v float32)`

SetTotalVolume sets TotalVolume field to given value.

### HasTotalVolume

`func (o *ETF) HasTotalVolume() bool`

HasTotalVolume returns a boolean if a field has been set.

### GetTradeTimeInLong

`func (o *ETF) GetTradeTimeInLong() float32`

GetTradeTimeInLong returns the TradeTimeInLong field if non-nil, zero value otherwise.

### GetTradeTimeInLongOk

`func (o *ETF) GetTradeTimeInLongOk() (*float32, bool)`

GetTradeTimeInLongOk returns a tuple with the TradeTimeInLong field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeTimeInLong

`func (o *ETF) SetTradeTimeInLong(v float32)`

SetTradeTimeInLong sets TradeTimeInLong field to given value.

### HasTradeTimeInLong

`func (o *ETF) HasTradeTimeInLong() bool`

HasTradeTimeInLong returns a boolean if a field has been set.

### GetVolatility

`func (o *ETF) GetVolatility() float32`

GetVolatility returns the Volatility field if non-nil, zero value otherwise.

### GetVolatilityOk

`func (o *ETF) GetVolatilityOk() (*float32, bool)`

GetVolatilityOk returns a tuple with the Volatility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolatility

`func (o *ETF) SetVolatility(v float32)`

SetVolatility sets Volatility field to given value.

### HasVolatility

`func (o *ETF) HasVolatility() bool`

HasVolatility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


