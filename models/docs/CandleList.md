# CandleList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Candles** | Pointer to [**[]CandleListCandles**](CandleListCandles.md) |  | [optional] 
**Empty** | Pointer to **bool** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 

## Methods

### NewCandleList

`func NewCandleList() *CandleList`

NewCandleList instantiates a new CandleList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCandleListWithDefaults

`func NewCandleListWithDefaults() *CandleList`

NewCandleListWithDefaults instantiates a new CandleList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCandles

`func (o *CandleList) GetCandles() []CandleListCandles`

GetCandles returns the Candles field if non-nil, zero value otherwise.

### GetCandlesOk

`func (o *CandleList) GetCandlesOk() (*[]CandleListCandles, bool)`

GetCandlesOk returns a tuple with the Candles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandles

`func (o *CandleList) SetCandles(v []CandleListCandles)`

SetCandles sets Candles field to given value.

### HasCandles

`func (o *CandleList) HasCandles() bool`

HasCandles returns a boolean if a field has been set.

### GetEmpty

`func (o *CandleList) GetEmpty() bool`

GetEmpty returns the Empty field if non-nil, zero value otherwise.

### GetEmptyOk

`func (o *CandleList) GetEmptyOk() (*bool, bool)`

GetEmptyOk returns a tuple with the Empty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmpty

`func (o *CandleList) SetEmpty(v bool)`

SetEmpty sets Empty field to given value.

### HasEmpty

`func (o *CandleList) HasEmpty() bool`

HasEmpty returns a boolean if a field has been set.

### GetSymbol

`func (o *CandleList) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *CandleList) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *CandleList) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *CandleList) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


