# Option

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetType** | Pointer to [**AssetType**](AssetType.md) |  | [optional] 
**Cusip** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**OptionDeliverables** | Pointer to [**[]OptionOptionDeliverables**](OptionOptionDeliverables.md) |  | [optional] 
**OptionMultiplier** | Pointer to **float32** |  | [optional] 
**PutCall** | Pointer to [**PutOrCall**](PutOrCall.md) |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**OptionType**](OptionType.md) |  | [optional] 
**UnderlyingSymbol** | Pointer to **string** |  | [optional] 

## Methods

### NewOption

`func NewOption() *Option`

NewOption instantiates a new Option object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionWithDefaults

`func NewOptionWithDefaults() *Option`

NewOptionWithDefaults instantiates a new Option object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetType

`func (o *Option) GetAssetType() AssetType`

GetAssetType returns the AssetType field if non-nil, zero value otherwise.

### GetAssetTypeOk

`func (o *Option) GetAssetTypeOk() (*AssetType, bool)`

GetAssetTypeOk returns a tuple with the AssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetType

`func (o *Option) SetAssetType(v AssetType)`

SetAssetType sets AssetType field to given value.

### HasAssetType

`func (o *Option) HasAssetType() bool`

HasAssetType returns a boolean if a field has been set.

### GetCusip

`func (o *Option) GetCusip() string`

GetCusip returns the Cusip field if non-nil, zero value otherwise.

### GetCusipOk

`func (o *Option) GetCusipOk() (*string, bool)`

GetCusipOk returns a tuple with the Cusip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCusip

`func (o *Option) SetCusip(v string)`

SetCusip sets Cusip field to given value.

### HasCusip

`func (o *Option) HasCusip() bool`

HasCusip returns a boolean if a field has been set.

### GetDescription

`func (o *Option) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Option) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Option) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Option) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOptionDeliverables

`func (o *Option) GetOptionDeliverables() []OptionOptionDeliverables`

GetOptionDeliverables returns the OptionDeliverables field if non-nil, zero value otherwise.

### GetOptionDeliverablesOk

`func (o *Option) GetOptionDeliverablesOk() (*[]OptionOptionDeliverables, bool)`

GetOptionDeliverablesOk returns a tuple with the OptionDeliverables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionDeliverables

`func (o *Option) SetOptionDeliverables(v []OptionOptionDeliverables)`

SetOptionDeliverables sets OptionDeliverables field to given value.

### HasOptionDeliverables

`func (o *Option) HasOptionDeliverables() bool`

HasOptionDeliverables returns a boolean if a field has been set.

### GetOptionMultiplier

`func (o *Option) GetOptionMultiplier() float32`

GetOptionMultiplier returns the OptionMultiplier field if non-nil, zero value otherwise.

### GetOptionMultiplierOk

`func (o *Option) GetOptionMultiplierOk() (*float32, bool)`

GetOptionMultiplierOk returns a tuple with the OptionMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionMultiplier

`func (o *Option) SetOptionMultiplier(v float32)`

SetOptionMultiplier sets OptionMultiplier field to given value.

### HasOptionMultiplier

`func (o *Option) HasOptionMultiplier() bool`

HasOptionMultiplier returns a boolean if a field has been set.

### GetPutCall

`func (o *Option) GetPutCall() PutOrCall`

GetPutCall returns the PutCall field if non-nil, zero value otherwise.

### GetPutCallOk

`func (o *Option) GetPutCallOk() (*PutOrCall, bool)`

GetPutCallOk returns a tuple with the PutCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPutCall

`func (o *Option) SetPutCall(v PutOrCall)`

SetPutCall sets PutCall field to given value.

### HasPutCall

`func (o *Option) HasPutCall() bool`

HasPutCall returns a boolean if a field has been set.

### GetSymbol

`func (o *Option) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Option) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Option) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *Option) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetType

`func (o *Option) GetType() OptionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Option) GetTypeOk() (*OptionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Option) SetType(v OptionType)`

SetType sets Type field to given value.

### HasType

`func (o *Option) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnderlyingSymbol

`func (o *Option) GetUnderlyingSymbol() string`

GetUnderlyingSymbol returns the UnderlyingSymbol field if non-nil, zero value otherwise.

### GetUnderlyingSymbolOk

`func (o *Option) GetUnderlyingSymbolOk() (*string, bool)`

GetUnderlyingSymbolOk returns a tuple with the UnderlyingSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderlyingSymbol

`func (o *Option) SetUnderlyingSymbol(v string)`

SetUnderlyingSymbol sets UnderlyingSymbol field to given value.

### HasUnderlyingSymbol

`func (o *Option) HasUnderlyingSymbol() bool`

HasUnderlyingSymbol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


