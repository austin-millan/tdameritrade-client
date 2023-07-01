# Execution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityType** | Pointer to **string** |  | [optional] 
**ExecutionLegs** | Pointer to [**[]ExecutionExecutionLegs**](ExecutionExecutionLegs.md) |  | [optional] 
**ExecutionType** | Pointer to **string** |  | [optional] 
**OrderRemainingQuantity** | Pointer to **float32** |  | [optional] 
**Quantity** | Pointer to **float32** |  | [optional] 

## Methods

### NewExecution

`func NewExecution() *Execution`

NewExecution instantiates a new Execution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionWithDefaults

`func NewExecutionWithDefaults() *Execution`

NewExecutionWithDefaults instantiates a new Execution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivityType

`func (o *Execution) GetActivityType() string`

GetActivityType returns the ActivityType field if non-nil, zero value otherwise.

### GetActivityTypeOk

`func (o *Execution) GetActivityTypeOk() (*string, bool)`

GetActivityTypeOk returns a tuple with the ActivityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityType

`func (o *Execution) SetActivityType(v string)`

SetActivityType sets ActivityType field to given value.

### HasActivityType

`func (o *Execution) HasActivityType() bool`

HasActivityType returns a boolean if a field has been set.

### GetExecutionLegs

`func (o *Execution) GetExecutionLegs() []ExecutionExecutionLegs`

GetExecutionLegs returns the ExecutionLegs field if non-nil, zero value otherwise.

### GetExecutionLegsOk

`func (o *Execution) GetExecutionLegsOk() (*[]ExecutionExecutionLegs, bool)`

GetExecutionLegsOk returns a tuple with the ExecutionLegs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionLegs

`func (o *Execution) SetExecutionLegs(v []ExecutionExecutionLegs)`

SetExecutionLegs sets ExecutionLegs field to given value.

### HasExecutionLegs

`func (o *Execution) HasExecutionLegs() bool`

HasExecutionLegs returns a boolean if a field has been set.

### GetExecutionType

`func (o *Execution) GetExecutionType() string`

GetExecutionType returns the ExecutionType field if non-nil, zero value otherwise.

### GetExecutionTypeOk

`func (o *Execution) GetExecutionTypeOk() (*string, bool)`

GetExecutionTypeOk returns a tuple with the ExecutionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionType

`func (o *Execution) SetExecutionType(v string)`

SetExecutionType sets ExecutionType field to given value.

### HasExecutionType

`func (o *Execution) HasExecutionType() bool`

HasExecutionType returns a boolean if a field has been set.

### GetOrderRemainingQuantity

`func (o *Execution) GetOrderRemainingQuantity() float32`

GetOrderRemainingQuantity returns the OrderRemainingQuantity field if non-nil, zero value otherwise.

### GetOrderRemainingQuantityOk

`func (o *Execution) GetOrderRemainingQuantityOk() (*float32, bool)`

GetOrderRemainingQuantityOk returns a tuple with the OrderRemainingQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderRemainingQuantity

`func (o *Execution) SetOrderRemainingQuantity(v float32)`

SetOrderRemainingQuantity sets OrderRemainingQuantity field to given value.

### HasOrderRemainingQuantity

`func (o *Execution) HasOrderRemainingQuantity() bool`

HasOrderRemainingQuantity returns a boolean if a field has been set.

### GetQuantity

`func (o *Execution) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *Execution) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *Execution) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *Execution) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


