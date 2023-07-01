# TransactionTransactionItemInstrument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetType** | Pointer to [**AssetType**](AssetType.md) |  | [optional] 
**BondInterestRate** | Pointer to **float32** |  | [optional] 
**BondMaturityDate** | Pointer to **string** |  | [optional] 
**Cusip** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**OptionExpirationDate** | Pointer to **string** |  | [optional] 
**OptionStrikePrice** | Pointer to **float32** |  | [optional] 
**PutCall** | Pointer to [**PutOrCall**](PutOrCall.md) |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**UnderlyingSymbol** | Pointer to **string** |  | [optional] 

## Methods

### NewTransactionTransactionItemInstrument

`func NewTransactionTransactionItemInstrument() *TransactionTransactionItemInstrument`

NewTransactionTransactionItemInstrument instantiates a new TransactionTransactionItemInstrument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionTransactionItemInstrumentWithDefaults

`func NewTransactionTransactionItemInstrumentWithDefaults() *TransactionTransactionItemInstrument`

NewTransactionTransactionItemInstrumentWithDefaults instantiates a new TransactionTransactionItemInstrument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetType

`func (o *TransactionTransactionItemInstrument) GetAssetType() AssetType`

GetAssetType returns the AssetType field if non-nil, zero value otherwise.

### GetAssetTypeOk

`func (o *TransactionTransactionItemInstrument) GetAssetTypeOk() (*AssetType, bool)`

GetAssetTypeOk returns a tuple with the AssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetType

`func (o *TransactionTransactionItemInstrument) SetAssetType(v AssetType)`

SetAssetType sets AssetType field to given value.

### HasAssetType

`func (o *TransactionTransactionItemInstrument) HasAssetType() bool`

HasAssetType returns a boolean if a field has been set.

### GetBondInterestRate

`func (o *TransactionTransactionItemInstrument) GetBondInterestRate() float32`

GetBondInterestRate returns the BondInterestRate field if non-nil, zero value otherwise.

### GetBondInterestRateOk

`func (o *TransactionTransactionItemInstrument) GetBondInterestRateOk() (*float32, bool)`

GetBondInterestRateOk returns a tuple with the BondInterestRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondInterestRate

`func (o *TransactionTransactionItemInstrument) SetBondInterestRate(v float32)`

SetBondInterestRate sets BondInterestRate field to given value.

### HasBondInterestRate

`func (o *TransactionTransactionItemInstrument) HasBondInterestRate() bool`

HasBondInterestRate returns a boolean if a field has been set.

### GetBondMaturityDate

`func (o *TransactionTransactionItemInstrument) GetBondMaturityDate() string`

GetBondMaturityDate returns the BondMaturityDate field if non-nil, zero value otherwise.

### GetBondMaturityDateOk

`func (o *TransactionTransactionItemInstrument) GetBondMaturityDateOk() (*string, bool)`

GetBondMaturityDateOk returns a tuple with the BondMaturityDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondMaturityDate

`func (o *TransactionTransactionItemInstrument) SetBondMaturityDate(v string)`

SetBondMaturityDate sets BondMaturityDate field to given value.

### HasBondMaturityDate

`func (o *TransactionTransactionItemInstrument) HasBondMaturityDate() bool`

HasBondMaturityDate returns a boolean if a field has been set.

### GetCusip

`func (o *TransactionTransactionItemInstrument) GetCusip() string`

GetCusip returns the Cusip field if non-nil, zero value otherwise.

### GetCusipOk

`func (o *TransactionTransactionItemInstrument) GetCusipOk() (*string, bool)`

GetCusipOk returns a tuple with the Cusip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCusip

`func (o *TransactionTransactionItemInstrument) SetCusip(v string)`

SetCusip sets Cusip field to given value.

### HasCusip

`func (o *TransactionTransactionItemInstrument) HasCusip() bool`

HasCusip returns a boolean if a field has been set.

### GetDescription

`func (o *TransactionTransactionItemInstrument) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransactionTransactionItemInstrument) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransactionTransactionItemInstrument) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TransactionTransactionItemInstrument) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOptionExpirationDate

`func (o *TransactionTransactionItemInstrument) GetOptionExpirationDate() string`

GetOptionExpirationDate returns the OptionExpirationDate field if non-nil, zero value otherwise.

### GetOptionExpirationDateOk

`func (o *TransactionTransactionItemInstrument) GetOptionExpirationDateOk() (*string, bool)`

GetOptionExpirationDateOk returns a tuple with the OptionExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionExpirationDate

`func (o *TransactionTransactionItemInstrument) SetOptionExpirationDate(v string)`

SetOptionExpirationDate sets OptionExpirationDate field to given value.

### HasOptionExpirationDate

`func (o *TransactionTransactionItemInstrument) HasOptionExpirationDate() bool`

HasOptionExpirationDate returns a boolean if a field has been set.

### GetOptionStrikePrice

`func (o *TransactionTransactionItemInstrument) GetOptionStrikePrice() float32`

GetOptionStrikePrice returns the OptionStrikePrice field if non-nil, zero value otherwise.

### GetOptionStrikePriceOk

`func (o *TransactionTransactionItemInstrument) GetOptionStrikePriceOk() (*float32, bool)`

GetOptionStrikePriceOk returns a tuple with the OptionStrikePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionStrikePrice

`func (o *TransactionTransactionItemInstrument) SetOptionStrikePrice(v float32)`

SetOptionStrikePrice sets OptionStrikePrice field to given value.

### HasOptionStrikePrice

`func (o *TransactionTransactionItemInstrument) HasOptionStrikePrice() bool`

HasOptionStrikePrice returns a boolean if a field has been set.

### GetPutCall

`func (o *TransactionTransactionItemInstrument) GetPutCall() PutOrCall`

GetPutCall returns the PutCall field if non-nil, zero value otherwise.

### GetPutCallOk

`func (o *TransactionTransactionItemInstrument) GetPutCallOk() (*PutOrCall, bool)`

GetPutCallOk returns a tuple with the PutCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPutCall

`func (o *TransactionTransactionItemInstrument) SetPutCall(v PutOrCall)`

SetPutCall sets PutCall field to given value.

### HasPutCall

`func (o *TransactionTransactionItemInstrument) HasPutCall() bool`

HasPutCall returns a boolean if a field has been set.

### GetSymbol

`func (o *TransactionTransactionItemInstrument) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *TransactionTransactionItemInstrument) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *TransactionTransactionItemInstrument) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *TransactionTransactionItemInstrument) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetUnderlyingSymbol

`func (o *TransactionTransactionItemInstrument) GetUnderlyingSymbol() string`

GetUnderlyingSymbol returns the UnderlyingSymbol field if non-nil, zero value otherwise.

### GetUnderlyingSymbolOk

`func (o *TransactionTransactionItemInstrument) GetUnderlyingSymbolOk() (*string, bool)`

GetUnderlyingSymbolOk returns a tuple with the UnderlyingSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderlyingSymbol

`func (o *TransactionTransactionItemInstrument) SetUnderlyingSymbol(v string)`

SetUnderlyingSymbol sets UnderlyingSymbol field to given value.

### HasUnderlyingSymbol

`func (o *TransactionTransactionItemInstrument) HasUnderlyingSymbol() bool`

HasUnderlyingSymbol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


