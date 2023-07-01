# OrderOrderLegCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instruction** | Pointer to [**Instruction**](Instruction.md) |  | [optional] 
**Instrument** | Pointer to [**InstrumentType**](InstrumentType.md) |  | [optional] 
**LegId** | Pointer to **float32** |  | [optional] 
**OrderLegType** | Pointer to [**OrderType**](OrderType.md) |  | [optional] 
**PositionEffect** | Pointer to [**PositionEffect**](PositionEffect.md) |  | [optional] 
**Quantity** | Pointer to **float32** |  | [optional] 
**QuantityType** | Pointer to [**QuantityType**](QuantityType.md) |  | [optional] 

## Methods

### NewOrderOrderLegCollection

`func NewOrderOrderLegCollection() *OrderOrderLegCollection`

NewOrderOrderLegCollection instantiates a new OrderOrderLegCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderOrderLegCollectionWithDefaults

`func NewOrderOrderLegCollectionWithDefaults() *OrderOrderLegCollection`

NewOrderOrderLegCollectionWithDefaults instantiates a new OrderOrderLegCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstruction

`func (o *OrderOrderLegCollection) GetInstruction() Instruction`

GetInstruction returns the Instruction field if non-nil, zero value otherwise.

### GetInstructionOk

`func (o *OrderOrderLegCollection) GetInstructionOk() (*Instruction, bool)`

GetInstructionOk returns a tuple with the Instruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstruction

`func (o *OrderOrderLegCollection) SetInstruction(v Instruction)`

SetInstruction sets Instruction field to given value.

### HasInstruction

`func (o *OrderOrderLegCollection) HasInstruction() bool`

HasInstruction returns a boolean if a field has been set.

### GetInstrument

`func (o *OrderOrderLegCollection) GetInstrument() InstrumentType`

GetInstrument returns the Instrument field if non-nil, zero value otherwise.

### GetInstrumentOk

`func (o *OrderOrderLegCollection) GetInstrumentOk() (*InstrumentType, bool)`

GetInstrumentOk returns a tuple with the Instrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrument

`func (o *OrderOrderLegCollection) SetInstrument(v InstrumentType)`

SetInstrument sets Instrument field to given value.

### HasInstrument

`func (o *OrderOrderLegCollection) HasInstrument() bool`

HasInstrument returns a boolean if a field has been set.

### GetLegId

`func (o *OrderOrderLegCollection) GetLegId() float32`

GetLegId returns the LegId field if non-nil, zero value otherwise.

### GetLegIdOk

`func (o *OrderOrderLegCollection) GetLegIdOk() (*float32, bool)`

GetLegIdOk returns a tuple with the LegId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegId

`func (o *OrderOrderLegCollection) SetLegId(v float32)`

SetLegId sets LegId field to given value.

### HasLegId

`func (o *OrderOrderLegCollection) HasLegId() bool`

HasLegId returns a boolean if a field has been set.

### GetOrderLegType

`func (o *OrderOrderLegCollection) GetOrderLegType() OrderType`

GetOrderLegType returns the OrderLegType field if non-nil, zero value otherwise.

### GetOrderLegTypeOk

`func (o *OrderOrderLegCollection) GetOrderLegTypeOk() (*OrderType, bool)`

GetOrderLegTypeOk returns a tuple with the OrderLegType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderLegType

`func (o *OrderOrderLegCollection) SetOrderLegType(v OrderType)`

SetOrderLegType sets OrderLegType field to given value.

### HasOrderLegType

`func (o *OrderOrderLegCollection) HasOrderLegType() bool`

HasOrderLegType returns a boolean if a field has been set.

### GetPositionEffect

`func (o *OrderOrderLegCollection) GetPositionEffect() PositionEffect`

GetPositionEffect returns the PositionEffect field if non-nil, zero value otherwise.

### GetPositionEffectOk

`func (o *OrderOrderLegCollection) GetPositionEffectOk() (*PositionEffect, bool)`

GetPositionEffectOk returns a tuple with the PositionEffect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionEffect

`func (o *OrderOrderLegCollection) SetPositionEffect(v PositionEffect)`

SetPositionEffect sets PositionEffect field to given value.

### HasPositionEffect

`func (o *OrderOrderLegCollection) HasPositionEffect() bool`

HasPositionEffect returns a boolean if a field has been set.

### GetQuantity

`func (o *OrderOrderLegCollection) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OrderOrderLegCollection) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OrderOrderLegCollection) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *OrderOrderLegCollection) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetQuantityType

`func (o *OrderOrderLegCollection) GetQuantityType() QuantityType`

GetQuantityType returns the QuantityType field if non-nil, zero value otherwise.

### GetQuantityTypeOk

`func (o *OrderOrderLegCollection) GetQuantityTypeOk() (*QuantityType, bool)`

GetQuantityTypeOk returns a tuple with the QuantityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityType

`func (o *OrderOrderLegCollection) SetQuantityType(v QuantityType)`

SetQuantityType sets QuantityType field to given value.

### HasQuantityType

`func (o *OrderOrderLegCollection) HasQuantityType() bool`

HasQuantityType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


