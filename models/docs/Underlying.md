# Underlying

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

### NewUnderlying

`func NewUnderlying() *Underlying`

NewUnderlying instantiates a new Underlying object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnderlyingWithDefaults

`func NewUnderlyingWithDefaults() *Underlying`

NewUnderlyingWithDefaults instantiates a new Underlying object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsk

`func (o *Underlying) GetAsk() float32`

GetAsk returns the Ask field if non-nil, zero value otherwise.

### GetAskOk

`func (o *Underlying) GetAskOk() (*float32, bool)`

GetAskOk returns a tuple with the Ask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsk

`func (o *Underlying) SetAsk(v float32)`

SetAsk sets Ask field to given value.

### HasAsk

`func (o *Underlying) HasAsk() bool`

HasAsk returns a boolean if a field has been set.

### GetAskSize

`func (o *Underlying) GetAskSize() float32`

GetAskSize returns the AskSize field if non-nil, zero value otherwise.

### GetAskSizeOk

`func (o *Underlying) GetAskSizeOk() (*float32, bool)`

GetAskSizeOk returns a tuple with the AskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskSize

`func (o *Underlying) SetAskSize(v float32)`

SetAskSize sets AskSize field to given value.

### HasAskSize

`func (o *Underlying) HasAskSize() bool`

HasAskSize returns a boolean if a field has been set.

### GetBid

`func (o *Underlying) GetBid() float32`

GetBid returns the Bid field if non-nil, zero value otherwise.

### GetBidOk

`func (o *Underlying) GetBidOk() (*float32, bool)`

GetBidOk returns a tuple with the Bid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBid

`func (o *Underlying) SetBid(v float32)`

SetBid sets Bid field to given value.

### HasBid

`func (o *Underlying) HasBid() bool`

HasBid returns a boolean if a field has been set.

### GetBidSize

`func (o *Underlying) GetBidSize() float32`

GetBidSize returns the BidSize field if non-nil, zero value otherwise.

### GetBidSizeOk

`func (o *Underlying) GetBidSizeOk() (*float32, bool)`

GetBidSizeOk returns a tuple with the BidSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidSize

`func (o *Underlying) SetBidSize(v float32)`

SetBidSize sets BidSize field to given value.

### HasBidSize

`func (o *Underlying) HasBidSize() bool`

HasBidSize returns a boolean if a field has been set.

### GetChange

`func (o *Underlying) GetChange() float32`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *Underlying) GetChangeOk() (*float32, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *Underlying) SetChange(v float32)`

SetChange sets Change field to given value.

### HasChange

`func (o *Underlying) HasChange() bool`

HasChange returns a boolean if a field has been set.

### GetClose

`func (o *Underlying) GetClose() float32`

GetClose returns the Close field if non-nil, zero value otherwise.

### GetCloseOk

`func (o *Underlying) GetCloseOk() (*float32, bool)`

GetCloseOk returns a tuple with the Close field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClose

`func (o *Underlying) SetClose(v float32)`

SetClose sets Close field to given value.

### HasClose

`func (o *Underlying) HasClose() bool`

HasClose returns a boolean if a field has been set.

### GetDelayed

`func (o *Underlying) GetDelayed() bool`

GetDelayed returns the Delayed field if non-nil, zero value otherwise.

### GetDelayedOk

`func (o *Underlying) GetDelayedOk() (*bool, bool)`

GetDelayedOk returns a tuple with the Delayed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayed

`func (o *Underlying) SetDelayed(v bool)`

SetDelayed sets Delayed field to given value.

### HasDelayed

`func (o *Underlying) HasDelayed() bool`

HasDelayed returns a boolean if a field has been set.

### GetDescription

`func (o *Underlying) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Underlying) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Underlying) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Underlying) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExchangeName

`func (o *Underlying) GetExchangeName() string`

GetExchangeName returns the ExchangeName field if non-nil, zero value otherwise.

### GetExchangeNameOk

`func (o *Underlying) GetExchangeNameOk() (*string, bool)`

GetExchangeNameOk returns a tuple with the ExchangeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeName

`func (o *Underlying) SetExchangeName(v string)`

SetExchangeName sets ExchangeName field to given value.

### HasExchangeName

`func (o *Underlying) HasExchangeName() bool`

HasExchangeName returns a boolean if a field has been set.

### GetFiftyTwoWeekHigh

`func (o *Underlying) GetFiftyTwoWeekHigh() float32`

GetFiftyTwoWeekHigh returns the FiftyTwoWeekHigh field if non-nil, zero value otherwise.

### GetFiftyTwoWeekHighOk

`func (o *Underlying) GetFiftyTwoWeekHighOk() (*float32, bool)`

GetFiftyTwoWeekHighOk returns a tuple with the FiftyTwoWeekHigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiftyTwoWeekHigh

`func (o *Underlying) SetFiftyTwoWeekHigh(v float32)`

SetFiftyTwoWeekHigh sets FiftyTwoWeekHigh field to given value.

### HasFiftyTwoWeekHigh

`func (o *Underlying) HasFiftyTwoWeekHigh() bool`

HasFiftyTwoWeekHigh returns a boolean if a field has been set.

### GetFiftyTwoWeekLow

`func (o *Underlying) GetFiftyTwoWeekLow() float32`

GetFiftyTwoWeekLow returns the FiftyTwoWeekLow field if non-nil, zero value otherwise.

### GetFiftyTwoWeekLowOk

`func (o *Underlying) GetFiftyTwoWeekLowOk() (*float32, bool)`

GetFiftyTwoWeekLowOk returns a tuple with the FiftyTwoWeekLow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiftyTwoWeekLow

`func (o *Underlying) SetFiftyTwoWeekLow(v float32)`

SetFiftyTwoWeekLow sets FiftyTwoWeekLow field to given value.

### HasFiftyTwoWeekLow

`func (o *Underlying) HasFiftyTwoWeekLow() bool`

HasFiftyTwoWeekLow returns a boolean if a field has been set.

### GetHighPrice

`func (o *Underlying) GetHighPrice() float32`

GetHighPrice returns the HighPrice field if non-nil, zero value otherwise.

### GetHighPriceOk

`func (o *Underlying) GetHighPriceOk() (*float32, bool)`

GetHighPriceOk returns a tuple with the HighPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighPrice

`func (o *Underlying) SetHighPrice(v float32)`

SetHighPrice sets HighPrice field to given value.

### HasHighPrice

`func (o *Underlying) HasHighPrice() bool`

HasHighPrice returns a boolean if a field has been set.

### GetLast

`func (o *Underlying) GetLast() float32`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *Underlying) GetLastOk() (*float32, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *Underlying) SetLast(v float32)`

SetLast sets Last field to given value.

### HasLast

`func (o *Underlying) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetLowPrice

`func (o *Underlying) GetLowPrice() float32`

GetLowPrice returns the LowPrice field if non-nil, zero value otherwise.

### GetLowPriceOk

`func (o *Underlying) GetLowPriceOk() (*float32, bool)`

GetLowPriceOk returns a tuple with the LowPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowPrice

`func (o *Underlying) SetLowPrice(v float32)`

SetLowPrice sets LowPrice field to given value.

### HasLowPrice

`func (o *Underlying) HasLowPrice() bool`

HasLowPrice returns a boolean if a field has been set.

### GetMark

`func (o *Underlying) GetMark() float32`

GetMark returns the Mark field if non-nil, zero value otherwise.

### GetMarkOk

`func (o *Underlying) GetMarkOk() (*float32, bool)`

GetMarkOk returns a tuple with the Mark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMark

`func (o *Underlying) SetMark(v float32)`

SetMark sets Mark field to given value.

### HasMark

`func (o *Underlying) HasMark() bool`

HasMark returns a boolean if a field has been set.

### GetMarkChange

`func (o *Underlying) GetMarkChange() float32`

GetMarkChange returns the MarkChange field if non-nil, zero value otherwise.

### GetMarkChangeOk

`func (o *Underlying) GetMarkChangeOk() (*float32, bool)`

GetMarkChangeOk returns a tuple with the MarkChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkChange

`func (o *Underlying) SetMarkChange(v float32)`

SetMarkChange sets MarkChange field to given value.

### HasMarkChange

`func (o *Underlying) HasMarkChange() bool`

HasMarkChange returns a boolean if a field has been set.

### GetMarkPercentChange

`func (o *Underlying) GetMarkPercentChange() float32`

GetMarkPercentChange returns the MarkPercentChange field if non-nil, zero value otherwise.

### GetMarkPercentChangeOk

`func (o *Underlying) GetMarkPercentChangeOk() (*float32, bool)`

GetMarkPercentChangeOk returns a tuple with the MarkPercentChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkPercentChange

`func (o *Underlying) SetMarkPercentChange(v float32)`

SetMarkPercentChange sets MarkPercentChange field to given value.

### HasMarkPercentChange

`func (o *Underlying) HasMarkPercentChange() bool`

HasMarkPercentChange returns a boolean if a field has been set.

### GetOpenPrice

`func (o *Underlying) GetOpenPrice() float32`

GetOpenPrice returns the OpenPrice field if non-nil, zero value otherwise.

### GetOpenPriceOk

`func (o *Underlying) GetOpenPriceOk() (*float32, bool)`

GetOpenPriceOk returns a tuple with the OpenPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenPrice

`func (o *Underlying) SetOpenPrice(v float32)`

SetOpenPrice sets OpenPrice field to given value.

### HasOpenPrice

`func (o *Underlying) HasOpenPrice() bool`

HasOpenPrice returns a boolean if a field has been set.

### GetPercentChange

`func (o *Underlying) GetPercentChange() float32`

GetPercentChange returns the PercentChange field if non-nil, zero value otherwise.

### GetPercentChangeOk

`func (o *Underlying) GetPercentChangeOk() (*float32, bool)`

GetPercentChangeOk returns a tuple with the PercentChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentChange

`func (o *Underlying) SetPercentChange(v float32)`

SetPercentChange sets PercentChange field to given value.

### HasPercentChange

`func (o *Underlying) HasPercentChange() bool`

HasPercentChange returns a boolean if a field has been set.

### GetQuoteTime

`func (o *Underlying) GetQuoteTime() float32`

GetQuoteTime returns the QuoteTime field if non-nil, zero value otherwise.

### GetQuoteTimeOk

`func (o *Underlying) GetQuoteTimeOk() (*float32, bool)`

GetQuoteTimeOk returns a tuple with the QuoteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteTime

`func (o *Underlying) SetQuoteTime(v float32)`

SetQuoteTime sets QuoteTime field to given value.

### HasQuoteTime

`func (o *Underlying) HasQuoteTime() bool`

HasQuoteTime returns a boolean if a field has been set.

### GetSymbol

`func (o *Underlying) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Underlying) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Underlying) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *Underlying) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTotalVolume

`func (o *Underlying) GetTotalVolume() float32`

GetTotalVolume returns the TotalVolume field if non-nil, zero value otherwise.

### GetTotalVolumeOk

`func (o *Underlying) GetTotalVolumeOk() (*float32, bool)`

GetTotalVolumeOk returns a tuple with the TotalVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVolume

`func (o *Underlying) SetTotalVolume(v float32)`

SetTotalVolume sets TotalVolume field to given value.

### HasTotalVolume

`func (o *Underlying) HasTotalVolume() bool`

HasTotalVolume returns a boolean if a field has been set.

### GetTradeTime

`func (o *Underlying) GetTradeTime() float32`

GetTradeTime returns the TradeTime field if non-nil, zero value otherwise.

### GetTradeTimeOk

`func (o *Underlying) GetTradeTimeOk() (*float32, bool)`

GetTradeTimeOk returns a tuple with the TradeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeTime

`func (o *Underlying) SetTradeTime(v float32)`

SetTradeTime sets TradeTime field to given value.

### HasTradeTime

`func (o *Underlying) HasTradeTime() bool`

HasTradeTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


