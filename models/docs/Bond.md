# Bond

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetType** | Pointer to [**AssetType**](AssetType.md) |  | [optional] 
**BondPrice** | Pointer to **float32** |  | [optional] 
**Cusip** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Exchange** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 

## Methods

### NewBond

`func NewBond() *Bond`

NewBond instantiates a new Bond object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBondWithDefaults

`func NewBondWithDefaults() *Bond`

NewBondWithDefaults instantiates a new Bond object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetType

`func (o *Bond) GetAssetType() AssetType`

GetAssetType returns the AssetType field if non-nil, zero value otherwise.

### GetAssetTypeOk

`func (o *Bond) GetAssetTypeOk() (*AssetType, bool)`

GetAssetTypeOk returns a tuple with the AssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetType

`func (o *Bond) SetAssetType(v AssetType)`

SetAssetType sets AssetType field to given value.

### HasAssetType

`func (o *Bond) HasAssetType() bool`

HasAssetType returns a boolean if a field has been set.

### GetBondPrice

`func (o *Bond) GetBondPrice() float32`

GetBondPrice returns the BondPrice field if non-nil, zero value otherwise.

### GetBondPriceOk

`func (o *Bond) GetBondPriceOk() (*float32, bool)`

GetBondPriceOk returns a tuple with the BondPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondPrice

`func (o *Bond) SetBondPrice(v float32)`

SetBondPrice sets BondPrice field to given value.

### HasBondPrice

`func (o *Bond) HasBondPrice() bool`

HasBondPrice returns a boolean if a field has been set.

### GetCusip

`func (o *Bond) GetCusip() string`

GetCusip returns the Cusip field if non-nil, zero value otherwise.

### GetCusipOk

`func (o *Bond) GetCusipOk() (*string, bool)`

GetCusipOk returns a tuple with the Cusip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCusip

`func (o *Bond) SetCusip(v string)`

SetCusip sets Cusip field to given value.

### HasCusip

`func (o *Bond) HasCusip() bool`

HasCusip returns a boolean if a field has been set.

### GetDescription

`func (o *Bond) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Bond) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Bond) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Bond) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExchange

`func (o *Bond) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *Bond) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *Bond) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *Bond) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetSymbol

`func (o *Bond) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Bond) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Bond) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *Bond) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


