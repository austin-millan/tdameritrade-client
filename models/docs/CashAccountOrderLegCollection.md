# CashAccountOrderLegCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instruction** | Pointer to [**Instruction**](Instruction.md) |  | [optional] 
**Instrument** | Pointer to [**InstrumentType**](InstrumentType.md) |  | [optional] 
**LegId** | Pointer to **float32** |  | [optional] 
**OrderLegType** | Pointer to [**AssetType**](AssetType.md) |  | [optional] 
**PositionEffect** | Pointer to [**PositionEffect**](PositionEffect.md) |  | [optional] 
**Quantity** | Pointer to **float32** |  | [optional] 
**QuantityType** | Pointer to [**QuantityType**](QuantityType.md) |  | [optional] 

## Methods

### NewCashAccountOrderLegCollection

`func NewCashAccountOrderLegCollection() *CashAccountOrderLegCollection`

NewCashAccountOrderLegCollection instantiates a new CashAccountOrderLegCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCashAccountOrderLegCollectionWithDefaults

`func NewCashAccountOrderLegCollectionWithDefaults() *CashAccountOrderLegCollection`

NewCashAccountOrderLegCollectionWithDefaults instantiates a new CashAccountOrderLegCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstruction

`func (o *CashAccountOrderLegCollection) GetInstruction() Instruction`

GetInstruction returns the Instruction field if non-nil, zero value otherwise.

### GetInstructionOk

`func (o *CashAccountOrderLegCollection) GetInstructionOk() (*Instruction, bool)`

GetInstructionOk returns a tuple with the Instruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstruction

`func (o *CashAccountOrderLegCollection) SetInstruction(v Instruction)`

SetInstruction sets Instruction field to given value.

### HasInstruction

`func (o *CashAccountOrderLegCollection) HasInstruction() bool`

HasInstruction returns a boolean if a field has been set.

### GetInstrument

`func (o *CashAccountOrderLegCollection) GetInstrument() InstrumentType`

GetInstrument returns the Instrument field if non-nil, zero value otherwise.

### GetInstrumentOk

`func (o *CashAccountOrderLegCollection) GetInstrumentOk() (*InstrumentType, bool)`

GetInstrumentOk returns a tuple with the Instrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrument

`func (o *CashAccountOrderLegCollection) SetInstrument(v InstrumentType)`

SetInstrument sets Instrument field to given value.

### HasInstrument

`func (o *CashAccountOrderLegCollection) HasInstrument() bool`

HasInstrument returns a boolean if a field has been set.

### GetLegId

`func (o *CashAccountOrderLegCollection) GetLegId() float32`

GetLegId returns the LegId field if non-nil, zero value otherwise.

### GetLegIdOk

`func (o *CashAccountOrderLegCollection) GetLegIdOk() (*float32, bool)`

GetLegIdOk returns a tuple with the LegId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegId

`func (o *CashAccountOrderLegCollection) SetLegId(v float32)`

SetLegId sets LegId field to given value.

### HasLegId

`func (o *CashAccountOrderLegCollection) HasLegId() bool`

HasLegId returns a boolean if a field has been set.

### GetOrderLegType

`func (o *CashAccountOrderLegCollection) GetOrderLegType() AssetType`

GetOrderLegType returns the OrderLegType field if non-nil, zero value otherwise.

### GetOrderLegTypeOk

`func (o *CashAccountOrderLegCollection) GetOrderLegTypeOk() (*AssetType, bool)`

GetOrderLegTypeOk returns a tuple with the OrderLegType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderLegType

`func (o *CashAccountOrderLegCollection) SetOrderLegType(v AssetType)`

SetOrderLegType sets OrderLegType field to given value.

### HasOrderLegType

`func (o *CashAccountOrderLegCollection) HasOrderLegType() bool`

HasOrderLegType returns a boolean if a field has been set.

### GetPositionEffect

`func (o *CashAccountOrderLegCollection) GetPositionEffect() PositionEffect`

GetPositionEffect returns the PositionEffect field if non-nil, zero value otherwise.

### GetPositionEffectOk

`func (o *CashAccountOrderLegCollection) GetPositionEffectOk() (*PositionEffect, bool)`

GetPositionEffectOk returns a tuple with the PositionEffect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionEffect

`func (o *CashAccountOrderLegCollection) SetPositionEffect(v PositionEffect)`

SetPositionEffect sets PositionEffect field to given value.

### HasPositionEffect

`func (o *CashAccountOrderLegCollection) HasPositionEffect() bool`

HasPositionEffect returns a boolean if a field has been set.

### GetQuantity

`func (o *CashAccountOrderLegCollection) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CashAccountOrderLegCollection) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CashAccountOrderLegCollection) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *CashAccountOrderLegCollection) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetQuantityType

`func (o *CashAccountOrderLegCollection) GetQuantityType() QuantityType`

GetQuantityType returns the QuantityType field if non-nil, zero value otherwise.

### GetQuantityTypeOk

`func (o *CashAccountOrderLegCollection) GetQuantityTypeOk() (*QuantityType, bool)`

GetQuantityTypeOk returns a tuple with the QuantityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityType

`func (o *CashAccountOrderLegCollection) SetQuantityType(v QuantityType)`

SetQuantityType sets QuantityType field to given value.

### HasQuantityType

`func (o *CashAccountOrderLegCollection) HasQuantityType() bool`

HasQuantityType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


