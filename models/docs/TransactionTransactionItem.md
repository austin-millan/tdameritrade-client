# TransactionTransactionItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **float32** |  | [optional] 
**Amount** | Pointer to **float32** |  | [optional] 
**Cost** | Pointer to **float32** |  | [optional] 
**Instruction** | Pointer to [**Instruction**](Instruction.md) |  | [optional] 
**Instrument** | Pointer to [**TransactionTransactionItemInstrument**](Transaction_transactionItem_instrument.md) |  | [optional] 
**ParentChildIndicator** | Pointer to **string** |  | [optional] 
**ParentOrderKey** | Pointer to **float32** |  | [optional] 
**PositionEffect** | Pointer to [**PositionEffect**](PositionEffect.md) |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 

## Methods

### NewTransactionTransactionItem

`func NewTransactionTransactionItem() *TransactionTransactionItem`

NewTransactionTransactionItem instantiates a new TransactionTransactionItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionTransactionItemWithDefaults

`func NewTransactionTransactionItemWithDefaults() *TransactionTransactionItem`

NewTransactionTransactionItemWithDefaults instantiates a new TransactionTransactionItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *TransactionTransactionItem) GetAccountId() float32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *TransactionTransactionItem) GetAccountIdOk() (*float32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *TransactionTransactionItem) SetAccountId(v float32)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *TransactionTransactionItem) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAmount

`func (o *TransactionTransactionItem) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionTransactionItem) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionTransactionItem) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *TransactionTransactionItem) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCost

`func (o *TransactionTransactionItem) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *TransactionTransactionItem) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *TransactionTransactionItem) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *TransactionTransactionItem) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetInstruction

`func (o *TransactionTransactionItem) GetInstruction() Instruction`

GetInstruction returns the Instruction field if non-nil, zero value otherwise.

### GetInstructionOk

`func (o *TransactionTransactionItem) GetInstructionOk() (*Instruction, bool)`

GetInstructionOk returns a tuple with the Instruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstruction

`func (o *TransactionTransactionItem) SetInstruction(v Instruction)`

SetInstruction sets Instruction field to given value.

### HasInstruction

`func (o *TransactionTransactionItem) HasInstruction() bool`

HasInstruction returns a boolean if a field has been set.

### GetInstrument

`func (o *TransactionTransactionItem) GetInstrument() TransactionTransactionItemInstrument`

GetInstrument returns the Instrument field if non-nil, zero value otherwise.

### GetInstrumentOk

`func (o *TransactionTransactionItem) GetInstrumentOk() (*TransactionTransactionItemInstrument, bool)`

GetInstrumentOk returns a tuple with the Instrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrument

`func (o *TransactionTransactionItem) SetInstrument(v TransactionTransactionItemInstrument)`

SetInstrument sets Instrument field to given value.

### HasInstrument

`func (o *TransactionTransactionItem) HasInstrument() bool`

HasInstrument returns a boolean if a field has been set.

### GetParentChildIndicator

`func (o *TransactionTransactionItem) GetParentChildIndicator() string`

GetParentChildIndicator returns the ParentChildIndicator field if non-nil, zero value otherwise.

### GetParentChildIndicatorOk

`func (o *TransactionTransactionItem) GetParentChildIndicatorOk() (*string, bool)`

GetParentChildIndicatorOk returns a tuple with the ParentChildIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentChildIndicator

`func (o *TransactionTransactionItem) SetParentChildIndicator(v string)`

SetParentChildIndicator sets ParentChildIndicator field to given value.

### HasParentChildIndicator

`func (o *TransactionTransactionItem) HasParentChildIndicator() bool`

HasParentChildIndicator returns a boolean if a field has been set.

### GetParentOrderKey

`func (o *TransactionTransactionItem) GetParentOrderKey() float32`

GetParentOrderKey returns the ParentOrderKey field if non-nil, zero value otherwise.

### GetParentOrderKeyOk

`func (o *TransactionTransactionItem) GetParentOrderKeyOk() (*float32, bool)`

GetParentOrderKeyOk returns a tuple with the ParentOrderKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentOrderKey

`func (o *TransactionTransactionItem) SetParentOrderKey(v float32)`

SetParentOrderKey sets ParentOrderKey field to given value.

### HasParentOrderKey

`func (o *TransactionTransactionItem) HasParentOrderKey() bool`

HasParentOrderKey returns a boolean if a field has been set.

### GetPositionEffect

`func (o *TransactionTransactionItem) GetPositionEffect() PositionEffect`

GetPositionEffect returns the PositionEffect field if non-nil, zero value otherwise.

### GetPositionEffectOk

`func (o *TransactionTransactionItem) GetPositionEffectOk() (*PositionEffect, bool)`

GetPositionEffectOk returns a tuple with the PositionEffect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionEffect

`func (o *TransactionTransactionItem) SetPositionEffect(v PositionEffect)`

SetPositionEffect sets PositionEffect field to given value.

### HasPositionEffect

`func (o *TransactionTransactionItem) HasPositionEffect() bool`

HasPositionEffect returns a boolean if a field has been set.

### GetPrice

`func (o *TransactionTransactionItem) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *TransactionTransactionItem) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *TransactionTransactionItem) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *TransactionTransactionItem) HasPrice() bool`

HasPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


