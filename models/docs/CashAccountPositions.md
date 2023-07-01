# CashAccountPositions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgedQuantity** | Pointer to **float32** |  | [optional] 
**AveragePrice** | Pointer to **float32** |  | [optional] 
**CurrentDayProfitLoss** | Pointer to **float32** |  | [optional] 
**CurrentDayProfitLossPercentage** | Pointer to **float32** |  | [optional] 
**Instrument** | Pointer to [**InstrumentType**](InstrumentType.md) |  | [optional] 
**LongQuantity** | Pointer to **float32** |  | [optional] 
**MarketValue** | Pointer to **float32** |  | [optional] 
**SettledLongQuantity** | Pointer to **float32** |  | [optional] 
**SettledShortQuantity** | Pointer to **float32** |  | [optional] 
**ShortQuantity** | Pointer to **float32** |  | [optional] 

## Methods

### NewCashAccountPositions

`func NewCashAccountPositions() *CashAccountPositions`

NewCashAccountPositions instantiates a new CashAccountPositions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCashAccountPositionsWithDefaults

`func NewCashAccountPositionsWithDefaults() *CashAccountPositions`

NewCashAccountPositionsWithDefaults instantiates a new CashAccountPositions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgedQuantity

`func (o *CashAccountPositions) GetAgedQuantity() float32`

GetAgedQuantity returns the AgedQuantity field if non-nil, zero value otherwise.

### GetAgedQuantityOk

`func (o *CashAccountPositions) GetAgedQuantityOk() (*float32, bool)`

GetAgedQuantityOk returns a tuple with the AgedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgedQuantity

`func (o *CashAccountPositions) SetAgedQuantity(v float32)`

SetAgedQuantity sets AgedQuantity field to given value.

### HasAgedQuantity

`func (o *CashAccountPositions) HasAgedQuantity() bool`

HasAgedQuantity returns a boolean if a field has been set.

### GetAveragePrice

`func (o *CashAccountPositions) GetAveragePrice() float32`

GetAveragePrice returns the AveragePrice field if non-nil, zero value otherwise.

### GetAveragePriceOk

`func (o *CashAccountPositions) GetAveragePriceOk() (*float32, bool)`

GetAveragePriceOk returns a tuple with the AveragePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAveragePrice

`func (o *CashAccountPositions) SetAveragePrice(v float32)`

SetAveragePrice sets AveragePrice field to given value.

### HasAveragePrice

`func (o *CashAccountPositions) HasAveragePrice() bool`

HasAveragePrice returns a boolean if a field has been set.

### GetCurrentDayProfitLoss

`func (o *CashAccountPositions) GetCurrentDayProfitLoss() float32`

GetCurrentDayProfitLoss returns the CurrentDayProfitLoss field if non-nil, zero value otherwise.

### GetCurrentDayProfitLossOk

`func (o *CashAccountPositions) GetCurrentDayProfitLossOk() (*float32, bool)`

GetCurrentDayProfitLossOk returns a tuple with the CurrentDayProfitLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentDayProfitLoss

`func (o *CashAccountPositions) SetCurrentDayProfitLoss(v float32)`

SetCurrentDayProfitLoss sets CurrentDayProfitLoss field to given value.

### HasCurrentDayProfitLoss

`func (o *CashAccountPositions) HasCurrentDayProfitLoss() bool`

HasCurrentDayProfitLoss returns a boolean if a field has been set.

### GetCurrentDayProfitLossPercentage

`func (o *CashAccountPositions) GetCurrentDayProfitLossPercentage() float32`

GetCurrentDayProfitLossPercentage returns the CurrentDayProfitLossPercentage field if non-nil, zero value otherwise.

### GetCurrentDayProfitLossPercentageOk

`func (o *CashAccountPositions) GetCurrentDayProfitLossPercentageOk() (*float32, bool)`

GetCurrentDayProfitLossPercentageOk returns a tuple with the CurrentDayProfitLossPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentDayProfitLossPercentage

`func (o *CashAccountPositions) SetCurrentDayProfitLossPercentage(v float32)`

SetCurrentDayProfitLossPercentage sets CurrentDayProfitLossPercentage field to given value.

### HasCurrentDayProfitLossPercentage

`func (o *CashAccountPositions) HasCurrentDayProfitLossPercentage() bool`

HasCurrentDayProfitLossPercentage returns a boolean if a field has been set.

### GetInstrument

`func (o *CashAccountPositions) GetInstrument() InstrumentType`

GetInstrument returns the Instrument field if non-nil, zero value otherwise.

### GetInstrumentOk

`func (o *CashAccountPositions) GetInstrumentOk() (*InstrumentType, bool)`

GetInstrumentOk returns a tuple with the Instrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrument

`func (o *CashAccountPositions) SetInstrument(v InstrumentType)`

SetInstrument sets Instrument field to given value.

### HasInstrument

`func (o *CashAccountPositions) HasInstrument() bool`

HasInstrument returns a boolean if a field has been set.

### GetLongQuantity

`func (o *CashAccountPositions) GetLongQuantity() float32`

GetLongQuantity returns the LongQuantity field if non-nil, zero value otherwise.

### GetLongQuantityOk

`func (o *CashAccountPositions) GetLongQuantityOk() (*float32, bool)`

GetLongQuantityOk returns a tuple with the LongQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongQuantity

`func (o *CashAccountPositions) SetLongQuantity(v float32)`

SetLongQuantity sets LongQuantity field to given value.

### HasLongQuantity

`func (o *CashAccountPositions) HasLongQuantity() bool`

HasLongQuantity returns a boolean if a field has been set.

### GetMarketValue

`func (o *CashAccountPositions) GetMarketValue() float32`

GetMarketValue returns the MarketValue field if non-nil, zero value otherwise.

### GetMarketValueOk

`func (o *CashAccountPositions) GetMarketValueOk() (*float32, bool)`

GetMarketValueOk returns a tuple with the MarketValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketValue

`func (o *CashAccountPositions) SetMarketValue(v float32)`

SetMarketValue sets MarketValue field to given value.

### HasMarketValue

`func (o *CashAccountPositions) HasMarketValue() bool`

HasMarketValue returns a boolean if a field has been set.

### GetSettledLongQuantity

`func (o *CashAccountPositions) GetSettledLongQuantity() float32`

GetSettledLongQuantity returns the SettledLongQuantity field if non-nil, zero value otherwise.

### GetSettledLongQuantityOk

`func (o *CashAccountPositions) GetSettledLongQuantityOk() (*float32, bool)`

GetSettledLongQuantityOk returns a tuple with the SettledLongQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettledLongQuantity

`func (o *CashAccountPositions) SetSettledLongQuantity(v float32)`

SetSettledLongQuantity sets SettledLongQuantity field to given value.

### HasSettledLongQuantity

`func (o *CashAccountPositions) HasSettledLongQuantity() bool`

HasSettledLongQuantity returns a boolean if a field has been set.

### GetSettledShortQuantity

`func (o *CashAccountPositions) GetSettledShortQuantity() float32`

GetSettledShortQuantity returns the SettledShortQuantity field if non-nil, zero value otherwise.

### GetSettledShortQuantityOk

`func (o *CashAccountPositions) GetSettledShortQuantityOk() (*float32, bool)`

GetSettledShortQuantityOk returns a tuple with the SettledShortQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettledShortQuantity

`func (o *CashAccountPositions) SetSettledShortQuantity(v float32)`

SetSettledShortQuantity sets SettledShortQuantity field to given value.

### HasSettledShortQuantity

`func (o *CashAccountPositions) HasSettledShortQuantity() bool`

HasSettledShortQuantity returns a boolean if a field has been set.

### GetShortQuantity

`func (o *CashAccountPositions) GetShortQuantity() float32`

GetShortQuantity returns the ShortQuantity field if non-nil, zero value otherwise.

### GetShortQuantityOk

`func (o *CashAccountPositions) GetShortQuantityOk() (*float32, bool)`

GetShortQuantityOk returns a tuple with the ShortQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortQuantity

`func (o *CashAccountPositions) SetShortQuantity(v float32)`

SetShortQuantity sets ShortQuantity field to given value.

### HasShortQuantity

`func (o *CashAccountPositions) HasShortQuantity() bool`

HasShortQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


