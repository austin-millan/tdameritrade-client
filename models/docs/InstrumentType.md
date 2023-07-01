# InstrumentType

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
**Type** | Pointer to **string** |  | [optional] 
**UnderlyingSymbol** | Pointer to **string** |  | [optional] 
**Factor** | Pointer to **float32** |  | [optional] 
**MaturityDate** | Pointer to **string** |  | [optional] 
**VariableRate** | Pointer to **float32** |  | [optional] 

## Methods

### NewInstrumentType

`func NewInstrumentType() *InstrumentType`

NewInstrumentType instantiates a new InstrumentType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstrumentTypeWithDefaults

`func NewInstrumentTypeWithDefaults() *InstrumentType`

NewInstrumentTypeWithDefaults instantiates a new InstrumentType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetType

`func (o *InstrumentType) GetAssetType() AssetType`

GetAssetType returns the AssetType field if non-nil, zero value otherwise.

### GetAssetTypeOk

`func (o *InstrumentType) GetAssetTypeOk() (*AssetType, bool)`

GetAssetTypeOk returns a tuple with the AssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetType

`func (o *InstrumentType) SetAssetType(v AssetType)`

SetAssetType sets AssetType field to given value.

### HasAssetType

`func (o *InstrumentType) HasAssetType() bool`

HasAssetType returns a boolean if a field has been set.

### GetCusip

`func (o *InstrumentType) GetCusip() string`

GetCusip returns the Cusip field if non-nil, zero value otherwise.

### GetCusipOk

`func (o *InstrumentType) GetCusipOk() (*string, bool)`

GetCusipOk returns a tuple with the Cusip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCusip

`func (o *InstrumentType) SetCusip(v string)`

SetCusip sets Cusip field to given value.

### HasCusip

`func (o *InstrumentType) HasCusip() bool`

HasCusip returns a boolean if a field has been set.

### GetDescription

`func (o *InstrumentType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InstrumentType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InstrumentType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InstrumentType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOptionDeliverables

`func (o *InstrumentType) GetOptionDeliverables() []OptionOptionDeliverables`

GetOptionDeliverables returns the OptionDeliverables field if non-nil, zero value otherwise.

### GetOptionDeliverablesOk

`func (o *InstrumentType) GetOptionDeliverablesOk() (*[]OptionOptionDeliverables, bool)`

GetOptionDeliverablesOk returns a tuple with the OptionDeliverables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionDeliverables

`func (o *InstrumentType) SetOptionDeliverables(v []OptionOptionDeliverables)`

SetOptionDeliverables sets OptionDeliverables field to given value.

### HasOptionDeliverables

`func (o *InstrumentType) HasOptionDeliverables() bool`

HasOptionDeliverables returns a boolean if a field has been set.

### GetOptionMultiplier

`func (o *InstrumentType) GetOptionMultiplier() float32`

GetOptionMultiplier returns the OptionMultiplier field if non-nil, zero value otherwise.

### GetOptionMultiplierOk

`func (o *InstrumentType) GetOptionMultiplierOk() (*float32, bool)`

GetOptionMultiplierOk returns a tuple with the OptionMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionMultiplier

`func (o *InstrumentType) SetOptionMultiplier(v float32)`

SetOptionMultiplier sets OptionMultiplier field to given value.

### HasOptionMultiplier

`func (o *InstrumentType) HasOptionMultiplier() bool`

HasOptionMultiplier returns a boolean if a field has been set.

### GetPutCall

`func (o *InstrumentType) GetPutCall() PutOrCall`

GetPutCall returns the PutCall field if non-nil, zero value otherwise.

### GetPutCallOk

`func (o *InstrumentType) GetPutCallOk() (*PutOrCall, bool)`

GetPutCallOk returns a tuple with the PutCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPutCall

`func (o *InstrumentType) SetPutCall(v PutOrCall)`

SetPutCall sets PutCall field to given value.

### HasPutCall

`func (o *InstrumentType) HasPutCall() bool`

HasPutCall returns a boolean if a field has been set.

### GetSymbol

`func (o *InstrumentType) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *InstrumentType) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *InstrumentType) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *InstrumentType) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetType

`func (o *InstrumentType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InstrumentType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InstrumentType) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InstrumentType) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnderlyingSymbol

`func (o *InstrumentType) GetUnderlyingSymbol() string`

GetUnderlyingSymbol returns the UnderlyingSymbol field if non-nil, zero value otherwise.

### GetUnderlyingSymbolOk

`func (o *InstrumentType) GetUnderlyingSymbolOk() (*string, bool)`

GetUnderlyingSymbolOk returns a tuple with the UnderlyingSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderlyingSymbol

`func (o *InstrumentType) SetUnderlyingSymbol(v string)`

SetUnderlyingSymbol sets UnderlyingSymbol field to given value.

### HasUnderlyingSymbol

`func (o *InstrumentType) HasUnderlyingSymbol() bool`

HasUnderlyingSymbol returns a boolean if a field has been set.

### GetFactor

`func (o *InstrumentType) GetFactor() float32`

GetFactor returns the Factor field if non-nil, zero value otherwise.

### GetFactorOk

`func (o *InstrumentType) GetFactorOk() (*float32, bool)`

GetFactorOk returns a tuple with the Factor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactor

`func (o *InstrumentType) SetFactor(v float32)`

SetFactor sets Factor field to given value.

### HasFactor

`func (o *InstrumentType) HasFactor() bool`

HasFactor returns a boolean if a field has been set.

### GetMaturityDate

`func (o *InstrumentType) GetMaturityDate() string`

GetMaturityDate returns the MaturityDate field if non-nil, zero value otherwise.

### GetMaturityDateOk

`func (o *InstrumentType) GetMaturityDateOk() (*string, bool)`

GetMaturityDateOk returns a tuple with the MaturityDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaturityDate

`func (o *InstrumentType) SetMaturityDate(v string)`

SetMaturityDate sets MaturityDate field to given value.

### HasMaturityDate

`func (o *InstrumentType) HasMaturityDate() bool`

HasMaturityDate returns a boolean if a field has been set.

### GetVariableRate

`func (o *InstrumentType) GetVariableRate() float32`

GetVariableRate returns the VariableRate field if non-nil, zero value otherwise.

### GetVariableRateOk

`func (o *InstrumentType) GetVariableRateOk() (*float32, bool)`

GetVariableRateOk returns a tuple with the VariableRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableRate

`func (o *InstrumentType) SetVariableRate(v float32)`

SetVariableRate sets VariableRate field to given value.

### HasVariableRate

`func (o *InstrumentType) HasVariableRate() bool`

HasVariableRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


