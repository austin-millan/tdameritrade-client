# OptionChainUnderlying

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ask** | Pointer to **float32** |  | [optional] 
**AskSize** | Pointer to **float32** |  | [optional] 
**Bid** | Pointer to **float32** |  | [optional] 
**BidSize** | Pointer to **float32** |  | [optional] 
**Change** | Pointer to **float32** |  | [optional] 
**Close** | Pointer to **float32** |  | [optional] 
**Delayed** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ExchangeName** | Pointer to **string** |  | [optional] 
**FiftyTwoWeekHigh** | Pointer to **float32** |  | [optional] 
**FiftyTwoWeekLow** | Pointer to **float32** |  | [optional] 
**HighPrice** | Pointer to **float32** |  | [optional] 
**Last** | Pointer to **float32** |  | [optional] 
**LowPrice** | Pointer to **float32** |  | [optional] 
**Mark** | Pointer to **float32** |  | [optional] 
**MarkChange** | Pointer to **float32** |  | [optional] 
**MarkPercentChange** | Pointer to **float32** |  | [optional] 
**OpenPrice** | Pointer to **float32** |  | [optional] 
**PercentChange** | Pointer to **float32** |  | [optional] 
**QuoteTime** | Pointer to **float32** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**TotalVolume** | Pointer to **float32** |  | [optional] 
**TradeTime** | Pointer to **float32** |  | [optional] 

## Methods

### NewOptionChainUnderlying

`func NewOptionChainUnderlying() *OptionChainUnderlying`

NewOptionChainUnderlying instantiates a new OptionChainUnderlying object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionChainUnderlyingWithDefaults

`func NewOptionChainUnderlyingWithDefaults() *OptionChainUnderlying`

NewOptionChainUnderlyingWithDefaults instantiates a new OptionChainUnderlying object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsk

`func (o *OptionChainUnderlying) GetAsk() float32`

GetAsk returns the Ask field if non-nil, zero value otherwise.

### GetAskOk

`func (o *OptionChainUnderlying) GetAskOk() (*float32, bool)`

GetAskOk returns a tuple with the Ask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsk

`func (o *OptionChainUnderlying) SetAsk(v float32)`

SetAsk sets Ask field to given value.

### HasAsk

`func (o *OptionChainUnderlying) HasAsk() bool`

HasAsk returns a boolean if a field has been set.

### GetAskSize

`func (o *OptionChainUnderlying) GetAskSize() float32`

GetAskSize returns the AskSize field if non-nil, zero value otherwise.

### GetAskSizeOk

`func (o *OptionChainUnderlying) GetAskSizeOk() (*float32, bool)`

GetAskSizeOk returns a tuple with the AskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskSize

`func (o *OptionChainUnderlying) SetAskSize(v float32)`

SetAskSize sets AskSize field to given value.

### HasAskSize

`func (o *OptionChainUnderlying) HasAskSize() bool`

HasAskSize returns a boolean if a field has been set.

### GetBid

`func (o *OptionChainUnderlying) GetBid() float32`

GetBid returns the Bid field if non-nil, zero value otherwise.

### GetBidOk

`func (o *OptionChainUnderlying) GetBidOk() (*float32, bool)`

GetBidOk returns a tuple with the Bid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBid

`func (o *OptionChainUnderlying) SetBid(v float32)`

SetBid sets Bid field to given value.

### HasBid

`func (o *OptionChainUnderlying) HasBid() bool`

HasBid returns a boolean if a field has been set.

### GetBidSize

`func (o *OptionChainUnderlying) GetBidSize() float32`

GetBidSize returns the BidSize field if non-nil, zero value otherwise.

### GetBidSizeOk

`func (o *OptionChainUnderlying) GetBidSizeOk() (*float32, bool)`

GetBidSizeOk returns a tuple with the BidSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidSize

`func (o *OptionChainUnderlying) SetBidSize(v float32)`

SetBidSize sets BidSize field to given value.

### HasBidSize

`func (o *OptionChainUnderlying) HasBidSize() bool`

HasBidSize returns a boolean if a field has been set.

### GetChange

`func (o *OptionChainUnderlying) GetChange() float32`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *OptionChainUnderlying) GetChangeOk() (*float32, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *OptionChainUnderlying) SetChange(v float32)`

SetChange sets Change field to given value.

### HasChange

`func (o *OptionChainUnderlying) HasChange() bool`

HasChange returns a boolean if a field has been set.

### GetClose

`func (o *OptionChainUnderlying) GetClose() float32`

GetClose returns the Close field if non-nil, zero value otherwise.

### GetCloseOk

`func (o *OptionChainUnderlying) GetCloseOk() (*float32, bool)`

GetCloseOk returns a tuple with the Close field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClose

`func (o *OptionChainUnderlying) SetClose(v float32)`

SetClose sets Close field to given value.

### HasClose

`func (o *OptionChainUnderlying) HasClose() bool`

HasClose returns a boolean if a field has been set.

### GetDelayed

`func (o *OptionChainUnderlying) GetDelayed() bool`

GetDelayed returns the Delayed field if non-nil, zero value otherwise.

### GetDelayedOk

`func (o *OptionChainUnderlying) GetDelayedOk() (*bool, bool)`

GetDelayedOk returns a tuple with the Delayed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayed

`func (o *OptionChainUnderlying) SetDelayed(v bool)`

SetDelayed sets Delayed field to given value.

### HasDelayed

`func (o *OptionChainUnderlying) HasDelayed() bool`

HasDelayed returns a boolean if a field has been set.

### GetDescription

`func (o *OptionChainUnderlying) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OptionChainUnderlying) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OptionChainUnderlying) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OptionChainUnderlying) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExchangeName

`func (o *OptionChainUnderlying) GetExchangeName() string`

GetExchangeName returns the ExchangeName field if non-nil, zero value otherwise.

### GetExchangeNameOk

`func (o *OptionChainUnderlying) GetExchangeNameOk() (*string, bool)`

GetExchangeNameOk returns a tuple with the ExchangeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeName

`func (o *OptionChainUnderlying) SetExchangeName(v string)`

SetExchangeName sets ExchangeName field to given value.

### HasExchangeName

`func (o *OptionChainUnderlying) HasExchangeName() bool`

HasExchangeName returns a boolean if a field has been set.

### GetFiftyTwoWeekHigh

`func (o *OptionChainUnderlying) GetFiftyTwoWeekHigh() float32`

GetFiftyTwoWeekHigh returns the FiftyTwoWeekHigh field if non-nil, zero value otherwise.

### GetFiftyTwoWeekHighOk

`func (o *OptionChainUnderlying) GetFiftyTwoWeekHighOk() (*float32, bool)`

GetFiftyTwoWeekHighOk returns a tuple with the FiftyTwoWeekHigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiftyTwoWeekHigh

`func (o *OptionChainUnderlying) SetFiftyTwoWeekHigh(v float32)`

SetFiftyTwoWeekHigh sets FiftyTwoWeekHigh field to given value.

### HasFiftyTwoWeekHigh

`func (o *OptionChainUnderlying) HasFiftyTwoWeekHigh() bool`

HasFiftyTwoWeekHigh returns a boolean if a field has been set.

### GetFiftyTwoWeekLow

`func (o *OptionChainUnderlying) GetFiftyTwoWeekLow() float32`

GetFiftyTwoWeekLow returns the FiftyTwoWeekLow field if non-nil, zero value otherwise.

### GetFiftyTwoWeekLowOk

`func (o *OptionChainUnderlying) GetFiftyTwoWeekLowOk() (*float32, bool)`

GetFiftyTwoWeekLowOk returns a tuple with the FiftyTwoWeekLow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiftyTwoWeekLow

`func (o *OptionChainUnderlying) SetFiftyTwoWeekLow(v float32)`

SetFiftyTwoWeekLow sets FiftyTwoWeekLow field to given value.

### HasFiftyTwoWeekLow

`func (o *OptionChainUnderlying) HasFiftyTwoWeekLow() bool`

HasFiftyTwoWeekLow returns a boolean if a field has been set.

### GetHighPrice

`func (o *OptionChainUnderlying) GetHighPrice() float32`

GetHighPrice returns the HighPrice field if non-nil, zero value otherwise.

### GetHighPriceOk

`func (o *OptionChainUnderlying) GetHighPriceOk() (*float32, bool)`

GetHighPriceOk returns a tuple with the HighPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighPrice

`func (o *OptionChainUnderlying) SetHighPrice(v float32)`

SetHighPrice sets HighPrice field to given value.

### HasHighPrice

`func (o *OptionChainUnderlying) HasHighPrice() bool`

HasHighPrice returns a boolean if a field has been set.

### GetLast

`func (o *OptionChainUnderlying) GetLast() float32`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *OptionChainUnderlying) GetLastOk() (*float32, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *OptionChainUnderlying) SetLast(v float32)`

SetLast sets Last field to given value.

### HasLast

`func (o *OptionChainUnderlying) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetLowPrice

`func (o *OptionChainUnderlying) GetLowPrice() float32`

GetLowPrice returns the LowPrice field if non-nil, zero value otherwise.

### GetLowPriceOk

`func (o *OptionChainUnderlying) GetLowPriceOk() (*float32, bool)`

GetLowPriceOk returns a tuple with the LowPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowPrice

`func (o *OptionChainUnderlying) SetLowPrice(v float32)`

SetLowPrice sets LowPrice field to given value.

### HasLowPrice

`func (o *OptionChainUnderlying) HasLowPrice() bool`

HasLowPrice returns a boolean if a field has been set.

### GetMark

`func (o *OptionChainUnderlying) GetMark() float32`

GetMark returns the Mark field if non-nil, zero value otherwise.

### GetMarkOk

`func (o *OptionChainUnderlying) GetMarkOk() (*float32, bool)`

GetMarkOk returns a tuple with the Mark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMark

`func (o *OptionChainUnderlying) SetMark(v float32)`

SetMark sets Mark field to given value.

### HasMark

`func (o *OptionChainUnderlying) HasMark() bool`

HasMark returns a boolean if a field has been set.

### GetMarkChange

`func (o *OptionChainUnderlying) GetMarkChange() float32`

GetMarkChange returns the MarkChange field if non-nil, zero value otherwise.

### GetMarkChangeOk

`func (o *OptionChainUnderlying) GetMarkChangeOk() (*float32, bool)`

GetMarkChangeOk returns a tuple with the MarkChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkChange

`func (o *OptionChainUnderlying) SetMarkChange(v float32)`

SetMarkChange sets MarkChange field to given value.

### HasMarkChange

`func (o *OptionChainUnderlying) HasMarkChange() bool`

HasMarkChange returns a boolean if a field has been set.

### GetMarkPercentChange

`func (o *OptionChainUnderlying) GetMarkPercentChange() float32`

GetMarkPercentChange returns the MarkPercentChange field if non-nil, zero value otherwise.

### GetMarkPercentChangeOk

`func (o *OptionChainUnderlying) GetMarkPercentChangeOk() (*float32, bool)`

GetMarkPercentChangeOk returns a tuple with the MarkPercentChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkPercentChange

`func (o *OptionChainUnderlying) SetMarkPercentChange(v float32)`

SetMarkPercentChange sets MarkPercentChange field to given value.

### HasMarkPercentChange

`func (o *OptionChainUnderlying) HasMarkPercentChange() bool`

HasMarkPercentChange returns a boolean if a field has been set.

### GetOpenPrice

`func (o *OptionChainUnderlying) GetOpenPrice() float32`

GetOpenPrice returns the OpenPrice field if non-nil, zero value otherwise.

### GetOpenPriceOk

`func (o *OptionChainUnderlying) GetOpenPriceOk() (*float32, bool)`

GetOpenPriceOk returns a tuple with the OpenPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenPrice

`func (o *OptionChainUnderlying) SetOpenPrice(v float32)`

SetOpenPrice sets OpenPrice field to given value.

### HasOpenPrice

`func (o *OptionChainUnderlying) HasOpenPrice() bool`

HasOpenPrice returns a boolean if a field has been set.

### GetPercentChange

`func (o *OptionChainUnderlying) GetPercentChange() float32`

GetPercentChange returns the PercentChange field if non-nil, zero value otherwise.

### GetPercentChangeOk

`func (o *OptionChainUnderlying) GetPercentChangeOk() (*float32, bool)`

GetPercentChangeOk returns a tuple with the PercentChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentChange

`func (o *OptionChainUnderlying) SetPercentChange(v float32)`

SetPercentChange sets PercentChange field to given value.

### HasPercentChange

`func (o *OptionChainUnderlying) HasPercentChange() bool`

HasPercentChange returns a boolean if a field has been set.

### GetQuoteTime

`func (o *OptionChainUnderlying) GetQuoteTime() float32`

GetQuoteTime returns the QuoteTime field if non-nil, zero value otherwise.

### GetQuoteTimeOk

`func (o *OptionChainUnderlying) GetQuoteTimeOk() (*float32, bool)`

GetQuoteTimeOk returns a tuple with the QuoteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteTime

`func (o *OptionChainUnderlying) SetQuoteTime(v float32)`

SetQuoteTime sets QuoteTime field to given value.

### HasQuoteTime

`func (o *OptionChainUnderlying) HasQuoteTime() bool`

HasQuoteTime returns a boolean if a field has been set.

### GetSymbol

`func (o *OptionChainUnderlying) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *OptionChainUnderlying) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *OptionChainUnderlying) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *OptionChainUnderlying) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTotalVolume

`func (o *OptionChainUnderlying) GetTotalVolume() float32`

GetTotalVolume returns the TotalVolume field if non-nil, zero value otherwise.

### GetTotalVolumeOk

`func (o *OptionChainUnderlying) GetTotalVolumeOk() (*float32, bool)`

GetTotalVolumeOk returns a tuple with the TotalVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVolume

`func (o *OptionChainUnderlying) SetTotalVolume(v float32)`

SetTotalVolume sets TotalVolume field to given value.

### HasTotalVolume

`func (o *OptionChainUnderlying) HasTotalVolume() bool`

HasTotalVolume returns a boolean if a field has been set.

### GetTradeTime

`func (o *OptionChainUnderlying) GetTradeTime() float32`

GetTradeTime returns the TradeTime field if non-nil, zero value otherwise.

### GetTradeTimeOk

`func (o *OptionChainUnderlying) GetTradeTimeOk() (*float32, bool)`

GetTradeTimeOk returns a tuple with the TradeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeTime

`func (o *OptionChainUnderlying) SetTradeTime(v float32)`

SetTradeTime sets TradeTime field to given value.

### HasTradeTime

`func (o *OptionChainUnderlying) HasTradeTime() bool`

HasTradeTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


